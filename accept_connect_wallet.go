package main

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

// подключене metaMask к 1inch - подтверждение подключения
func AcceptConnectWallet(page playwright.Page) error {
	err := page.Locator(`//*[@id="app-content"]/div/div/div/div[2]/div/div[3]/div/div/button[2]`).Click()
	if err != nil {
		return fmt.Errorf("ошибка клика по кнопке - подключить: %w", err)
	}
	return nil
}
