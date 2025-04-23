package main

import (
	"demo/password/account"
	"fmt"
)

func main() {
	// 1. Создать аккаун
	// 2. Найти аккаунт
	// 3. Удалить аккаунт
	// 4. Выход
	fmt.Println("___ Мненджер паролей ___")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
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

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат Login или URL")
		return
	}
	vault := account.GetVault()
	vault.AddAccount(*myAccount)
}

func findAccount() {

}

func deleteAccount() {

}

func promptData(prompt string) string {
	fmt.Print(prompt + " : ")
	var res string
	fmt.Scanln(&res)
	return res
}
