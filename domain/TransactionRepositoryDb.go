package domain

import (
	"database/sql"
	"github.com/djedjethai/bankingLib/errs"
	"github.com/djedjethai/bankingLib/logger"
	"github.com/djedjethai/bankingSqlx/dto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}

func (db TransactionRepositoryDb) GetBalance(an string) (float64, *errs.AppError) {
	balanceSql := "SELECT amount FROM accounts WHERE account_id = ?"

	var balance float64
	if err := db.client.Get(&balance, balanceSql, an); err != nil {
		if err == sql.ErrNoRows {
			return balance, errs.NewBadRequestError("Unknow bank account")
		} else {
			logger.Error("err while scanning balance" + err.Error())
			return balance, errs.NewInternalServerError("Unexpected database err")
		}
	}

	return balance, nil
}

func (db TransactionRepositoryDb) RunTransaction(t Transaction, balance float64) (*dto.NewTransactionResponse, *errs.AppError) {

	// startting the db transaction block
	tx, err := db.client.Begin()
	if err != nil {
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	var newBalance float64
	// insert bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.TransactionType == "withdrawal" {
		newBalance = balance - t.Amount
		_, err = tx.Exec(`UPDATE accounts SET amount =  amount - ? WHERE account_id = ?`, t.Amount, t.AccountId)
	} else {
		newBalance = balance + t.Amount
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? WHERE account_id = ?`, t.Amount, t.AccountId)
	}

	// in case of any err we rollback the transaction
	if err != nil {
		tx.Rollback()
		logger.Error("Error while running transaction" + err.Error())
		return nil, errs.NewInternalServerError("Unexpected server error")
	}

	// commit the transaction if all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commmiting the transaction" + err.Error())
		return nil, errs.NewInternalServerError("Unexpected server error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while guetting last transactionId" + err.Error())
		return nil, errs.NewInternalServerError("Unexpected server error")
	}

	//
	transactionId := strconv.FormatInt(id, 10)
	tr := dto.NewTransactionResponse{
		Amount:        newBalance,
		TransactionId: transactionId,
	}

	return &tr, nil
}

// func (db TransactionRepositoryDb) UpdateAccountAmount(amt float64, aid string) *errs.AppError {
// 	updateSql := "UPDATE accounts SET amount = ? WHERE account_id = ?"
//
// 	_, err := db.client.Exec(updateSql, amt, aid)
// 	if err != nil {
// 		logger.Error("err while updating account" + err.Error())
// 		return errs.NewInternalServerError("Unexpected database error")
// 	}
//
// 	return nil
// }

// =====

// func (db TransactionRepositoryDb) UpdateTransactionTable(t Transaction) (string, *errs.AppError) {
// 	transactionSql := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"
//
// 	var transactionId string
// 	result, err := db.client.Exec(transactionSql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
// 	if err != nil {
// 		logger.Error("Error while updating transaction table" + err.Error())
// 		return transactionId, errs.NewInternalServerError("Unexpected server error")
// 	}
//
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		logger.Error("Error while guetting last transactionId" + err.Error())
// 		return transactionId, errs.NewInternalServerError("Unexpected server error")
// 	}
//
// 	transactionId = strconv.FormatInt(id, 10)
// 	return transactionId, nil
// }
//
// func (db TransactionRepositoryDb) UpdateAccountAmount(amt float64, aid string) *errs.AppError {
// 	updateSql := "UPDATE accounts SET amount = ? WHERE account_id = ?"
//
// 	_, err := db.client.Exec(updateSql, amt, aid)
// 	if err != nil {
// 		logger.Error("err while updating account" + err.Error())
// 		return errs.NewInternalServerError("Unexpected database error")
// 	}
//
// 	return nil
// }
