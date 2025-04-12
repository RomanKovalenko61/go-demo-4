package main

import (
	"fmt"
	"math/rand/v2"
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

func main() {
	login := promptData("Введите логин")
	url := promptData("Введите URL")

	myAccount := account{
		url: url,
		login: login,
	}
	myAccount.generatePassword(10)
	myAccount.outputPassword()
}

func promptData(prompt string) string{
	fmt.Print(prompt + " : ")
	var res string
	fmt.Scan(&res)
	return res
}