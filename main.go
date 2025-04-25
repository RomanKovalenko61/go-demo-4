package main

import (
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("___ Мненджер паролей ___")
	vault := account.GetVault(files.NewJsonDB("data.json"))
Menu:
	for {
		variant := promptData([]string{
			"1 - Создать аккаунт",
			"2 - Найти аккаунт",
			"3 - Удалить аккаунт",
			"4 - выход",
			"Выберите действие",
		})
		switch variant {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат Login или URL")
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска аккаунта"})
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		output.PrintError("Аккаунтов не найдено")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска аккаунта"})
	result := vault.DeleteAccountByUrl(url)
	if result {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func promptData[T any](prompt []T) string {
	for indx, el := range prompt {
		if indx == len(prompt)-1 {
			fmt.Printf("%v: ", el)
		} else {
			fmt.Println(el)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
