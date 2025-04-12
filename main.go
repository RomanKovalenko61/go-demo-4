package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login string
	password string
	url string
}

func (acc *account) generatePassword(n int) {
	pass := make([]rune, n)
	runes := []rune("abcdef123456")
	for index := range pass {
		pass[index] = runes[rand.IntN(len(runes))]
	}
	acc.password = string(pass)
}

func (acc account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

// 1. Если логина нет, то ошибка
// 2. Если пароля нет, то генерим
func newAccount(login, password, urlString string) (*account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := account{
		url: urlString,
		login: login,
		password: password,
	}
	if password == "" {
		newAcc.generatePassword(10)
	}
	return &newAcc, nil
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат Login или URL")
		return
	}
	//myAccount.generatePassword(10)
	myAccount.outputPassword()
}

func promptData(prompt string) string{
	fmt.Print(prompt + " : ")
	var res string
	fmt.Scanln(&res)
	return res
}