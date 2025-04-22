package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"reflect"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	login    string `json:"login" xml:"test"`
	password string
	url      string
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account   // acc account
}

func (acc *Account) generatePassword(n int) {
	pass := make([]rune, n)
	runes := []rune("abcdef123456")
	for index := range pass {
		pass[index] = runes[rand.IntN(len(runes))]
	}
	acc.password = string(pass)
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
	fmt.Println(acc.login, acc.password, acc.url)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	field, _ := reflect.TypeOf(newAcc).FieldByName("login")
	fmt.Println(string(field.Tag))
	if password == "" {
		newAcc.generatePassword(10)
	}
	return &newAcc, nil
}