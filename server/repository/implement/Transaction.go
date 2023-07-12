package implement

import (
	"context"
	"database/sql"
	"server/core/infra/repository"
)

var _ repository.ITransaction = &Transaction{}

type Transaction struct {
	db *sql.DB
	tx *sql.Tx
}

func NewTransaction() (*Transaction, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}

	return &Transaction{
		db: db,
	}, nil
}
func (r *Transaction) Begin() error {
	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}
	r.tx = tx
	return nil
}

func (r *Transaction) Commit() error {
	err := r.tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r *Transaction) Rollback() error {
	err := r.tx.Rollback()
	if err != nil {
		return err
	}
	return nil
}
