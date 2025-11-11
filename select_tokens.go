package main

import (
	"fmt"
	"time"

	"github.com/playwright-community/playwright-go"
)

func SelectTokens(page playwright.Page, token1 string, token2 string) error {
	// клик по выбору токена
	err := page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-token-field/oi-amount-input/div[1]/button`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика по токену: %w", err)
	}

	// клик по выбору сети
	err = page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div[1]/oi-chain-selector/button`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика по выбору сети: %w", err)
	}

	// выбор арбитрума
	err = page.Locator(`//*[@id="oi-select-item-11"]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика по арбитруму: %w", err)
	}

	// ввод 1 токена в поиск
	err = page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div/oi-search/div/oi-input/label/input`).Fill(token1)
	if err != nil {
		return fmt.Errorf("ошибка ввода названия токена: %w", err)
	}

	// ждем пока прогрузится результат поиска
	time.Sleep(time.Second * 3)

	token1ListLocator := page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[2]/oi-tokens-list/div`)

	token1Locator, err := getLocatorWithDivText(token1ListLocator, token1)
	if err != nil {
		return fmt.Errorf("ошибка поиска первого токена: %w", err)
	}

	// выбор 1 токена из списка
	err = token1Locator.Click()
	if err != nil {
		return fmt.Errorf("ошибка выбора первого токена: %w", err)
	}

	// select токен
	err = page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-swap-widget/oi-swap-form/oi-select-token/div[2]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика по выбору второго токена: %w", err)
	}

	// ввод второго токена в поиск
	err = page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[1]/div[1]/oi-search/div/oi-input/label`).Fill(token2)
	if err != nil {
		return fmt.Errorf("ошибка ввода второго токена: %w", err)
	}

	// ждем пока прогрузится результат поиска
	time.Sleep(time.Second * 3)

	token2ListLocator := page.Locator(`//html/body/app-root/app-shell/div/div/oi-page-simple-swap/div/div/oi-dialog/oi-token-picker-widget/div[2]/oi-tokens-list/div`)

	token2Locator, err := getLocatorWithDivText(token2ListLocator, token2)
	if err != nil {
		return fmt.Errorf("ошибка поиска второго токена: %w", err)
	}

	err = token2Locator.Click()
	if err != nil {
		return fmt.Errorf("ошибка выбора второго токена %s: %w", token2, err)
	}

	return nil
}
