package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
)

func main() {

	log := log.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("ошибка загрузки .env: %v", err)
	} // загрузка и проверка на ошибку

	seed := os.Getenv("METAMASK_SEED")
	if seed == "" {
		log.Fatal("METAMASK_SEED не найден в .env")
	}

	localPath := os.Getenv("LOCAL_PATH")
	if localPath == "" {
		log.Fatal("LOCAL_PATH не найден")
	}

	token1 := os.Getenv("TOKEN_1")
	if token1 == "" {
		log.Fatal("TOKEN_1 не найден")
	}

	token2 := os.Getenv("TOKEN_2")
	if token2 == "" {
		log.Fatal("TOKEN_2 не найден")
	}

	userDataPath := localPath + "/pw_userdata"

	pw, pwErr := playwright.Run()
	if pwErr != nil {
		log.Fatalf("не удалось запустить Playwright: %v", pwErr)
	} // проверка запуска

	defer pw.Stop()

	err = os.RemoveAll(userDataPath)
	if err != nil {
		log.Fatalf("ошибка очистки userdata до запуска браузера: %v", err)
	}

	// запуск браузера с расширением метамаск
	browser, err := pw.Chromium.LaunchPersistentContext(userDataPath, playwright.BrowserTypeLaunchPersistentContextOptions{
		Headless: playwright.Bool(false),
		Args: []string{
			"--disable-extensions-except=./metamask",
			"--load-extension=./metamask",
			"--disable-web-security",
			"--no-sandbox",
			"--allow-file-access-from-files",
		},
	})
	if err != nil {
		log.Fatalf("ошибка запуска браузера: %v", err)
	}

	log.Println("браузер запущен!")

	defer func() {
		err = browser.Close()
		if err != nil {
			log.Fatalf("ошибка закрытия браузера: %v", err)
		}

		err = os.RemoveAll(userDataPath)
		if err != nil {
			log.Fatalf("ошибка очистки userdata после запуска браузера: %v", err)
		}
	}()

	// ждем пока браузер откроет страницу с расширением
	time.Sleep(time.Second * 3)

	var metamaskPage playwright.Page

	pages := browser.Pages()
	for _, page := range pages {
		if strings.HasPrefix(page.URL(), "chrome-extension://") {
			metamaskPage = page
		}
	}

	if metamaskPage == nil {
		log.Fatalf("страница с метамаском не найдена")
	}

	// импортируем кошелек metaMask по seed-фразе
	err = ImportMetamask(metamaskPage, seed)
	if err != nil {
		log.Fatalf("ошибка импорта MetaMask: %v", err)
	}

	log.Println("кошелёк импортирован!")

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("ошибка создания страницы: %v", err)
	}

	// подключаем кошелек metaMask к 1inch
	err = ConnectWallet(page)
	if err != nil {
		log.Fatalf("ошибка подключения кошелька: %v", err)
	}

	log.Println("запрос на подключение кошелька..............")

	// ждем пока браузер откроет окно с подтверждением
	time.Sleep(time.Second * 3)

	// подтверждаем подключение кошелька в окне metaMask
	confirmPage := browser.Pages()[len(browser.Pages())-1]
	err = AcceptConnectWallet(confirmPage)
	if err != nil {
		log.Fatalf("ошибка подключения кошелька: %v", err)
	}

	log.Println("кошелёк подключен!")

	// выбираем токены для свапа
	err = SelectTokens(page, token1, token2)
	if err != nil {
		log.Fatalf("ошибка выбора токена: %v", err)
	}

	log.Println("токены выбраны!")

	log.Println("кейс выполнен успешно! для дальнейших шагов нужны деньги :(")
	log.Println("(через 15 сек окно браузера закроется)")
	time.Sleep(15 * time.Second)
}

func getLocatorWithDivText(source playwright.Locator, text string) (playwright.Locator, error) {

	listFiltered, err := source.Locator("div").Filter(playwright.LocatorFilterOptions{
		HasText: text,
	}).All()
	if err != nil {
		return nil, fmt.Errorf("ошибка поиска объектов: %w", err)
	}

	var result playwright.Locator

	for _, elem := range listFiltered {
		textContents, err := elem.AllTextContents()
		if err != nil {
			return nil, fmt.Errorf("ошибка получение текста: %w", err)
		}

		if slices.Contains(textContents, text) {
			result = elem
			break
		}
	}

	if result == nil {
		return nil, fmt.Errorf("div с текстом %s не найден", text)
	}
	return result, nil
}
