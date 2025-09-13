package main

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type Account struct {
	ID      int
	Balance float64
}

type Transaction struct {
	ID            int
	FromAccountID int
	ToAccountID   int
	Amount        float64
}

func createTowTables() {
	GlobalDB.AutoMigrate(&Account{}, &Transaction{})
}

func insertAccount() {
	GlobalDB.Create(&[]Account{{Balance: 999}, {Balance: 100}})
}

var mu sync.Mutex

func transferFunds(fromID, toID int, amount float64) error {
	mu.Lock()
	defer mu.Unlock()

	return GlobalDB.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account
		if err := tx.First(&fromAccount, fromID).Error; err != nil {
			return err
		}
		if err := tx.First(&toAccount, toID).Error; err != nil {
			return err
		}
		if fromAccount.Balance < amount {
			return errors.New("余额不足")
		}
		fromAccount.Balance -= amount
		toAccount.Balance += amount

		if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance).Error; err != nil {
			return err
		}
		if err := tx.Model(&toAccount).Update("balance", toAccount.Balance).Error; err != nil {
			return err
		}
		if err := tx.Create(&Transaction{FromAccountID: fromID, ToAccountID: toID, Amount: amount}).Error; err != nil {
			return err
		}
		return nil
	})
}
