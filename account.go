package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account   // acc account
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

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		newAcc.generatePassword(10)
	}
	return &newAcc, nil
}