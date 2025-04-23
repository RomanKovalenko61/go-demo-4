package main

import (
	"demo/password/account"
	"demo/password/files"
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
			searchAccount()
		case 3:
			removeAccount()
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
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func searchAccount() {

}

func removeAccount() {

}

func promptData(prompt string) string {
	fmt.Print(prompt + " : ")
	var res string
	fmt.Scanln(&res)
	return res
}
