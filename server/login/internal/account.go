package internal

import (
	"fmt"
	"server/db"
)

type Account struct {
	PlayerID  uint   `gorm:"primary_key"`
	AccountID string `gorm:"not null;unique"`
	Password  string
}

func getAccountByAccountID(accountID string) *Account {

	var account Account
	mysql := db.Mysql()
	err := mysql.Where("Account = ?", accountID).Limit(1).Find(&account).Error
	if nil != err {
		fmt.Println(err)
		return nil
	}
	fmt.Println("password:", account.Password)
	return &account
}

func creatAccountByAccountIDAndPassword(accountID string, password string) *Account {
	mysql := db.Mysql()
	var account = Account{AccountID: accountID, Password: password}
	err := mysql.Create(&account).Error
	if nil != err {
		return nil
	}

	return &account
}
