package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	output.PrintError(1)
	output.PrintError("sdsd")
	fmt.Println("___ Мненджер паролей ___")
	vault := account.GetVault(files.NewJsonDB("data.json"))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите действие: 1 - Создать акк 2 - Найти акк 3 - Удалить акк 4 - выход")
	fmt.Scanln(&variant)
	return variant
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат Login или URL")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска аккаунта")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска аккаунта")
	result := vault.DeleteAccountByUrl(url)
	if result {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + " : ")
	var res string
	fmt.Scanln(&res)
	return res
}
