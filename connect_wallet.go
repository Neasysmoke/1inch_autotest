package main

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

// подключение metaMask к 1inch
func ConnectWallet(page playwright.Page) error {
	_, err := page.Goto("https://1inch.com/swap")
	if err != nil {
		return fmt.Errorf("ошибка открытия 1inch: %w", err)
	}

	err = page.WaitForLoadState()
	if err != nil {
		return fmt.Errorf("ошибка ожидания загрузки страницы: %w", err)
	}

	err = page.Locator(`//html/body/app-root/app-shell/div/app-header/div[1]/div/div[3]/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	err = page.Locator(`//*[@id="cdk-overlay-0"]/oi-sidebar/div/app-wallet-connection-dialog/oi-wallet-list/div[1]/oi-wallet-cell[3]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	err = page.Locator(`//*[@id="cdk-overlay-0"]/oi-sidebar/div/app-wallet-connection-dialog/oi-select-network/div/div[1]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика: %w", err)
	}

	return nil
}
