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

func main() {
	login := promptData("Введите логин")
	password := generatePassword(10)
	url := promptData("Введите URL")

	myAccount := account{
		password: password,
		url: url,
		login: login,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string{
	fmt.Print(prompt + " : ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)
}

func generatePassword(n int) string{
	pass := make([]rune, n)
	runes := []rune("abcdef123456")
	for index := range pass {
		pass[index] = runes[rand.IntN(len(runes))]
	}
	return string(pass)
}