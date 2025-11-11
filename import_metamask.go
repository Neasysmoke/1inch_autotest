package main

import (
	"fmt"
	"os"

	"github.com/playwright-community/playwright-go"
)

func ImportMetamask(page playwright.Page, seed string) error {
	// создать новый кошелек
	err := page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	// у меня есть существующий кошелек
	err = page.Locator(`//html/body/div[3]/div[2]/div/section/div/button[3]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	// ввод сид-фразы
	err = page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[1]/div[4]/form/div/div[1]/div/textarea`).PressSequentially(seed)
	if err != nil {
		return fmt.Errorf("ошибка ввода seed: %w", err)
	}

	// продолжить
	err = page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	password := os.Getenv("METAMASK_PASSWORD")

	// ввод пароля
	err = page.Locator(`//*[@id="create-password-new"]`).Fill(password)
	if err != nil {
		return fmt.Errorf("ошибка ввода пароля: %w", err)
	}

	// подтверждение пароля
	err = page.Locator(`//*[@id="create-password-confirm"]`).Fill(password)
	if err != nil {
		return fmt.Errorf("ошибка подтверждения пароля: %w", err)
	}

	// галочка соглашения
	err = page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/form/div[1]/div[4]/label/span[1]/input`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика чекбокса: %w", err)
	}

	// создать пароль
	err = page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/form/div[2]/button`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	// чекбокс об аналитике
	err = page.Locator(`//*[@id="metametrics-opt-in"]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика чекбокса: %w", err)
	}

	// продолжить
	err = page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[4]/button`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	// выполнено
	err = page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[2]/button`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	return nil
}
