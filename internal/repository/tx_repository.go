package repository

import (
	"context"
	"fmt"

	"gitlab.com/abiewardani/scaffold/internal/system"
	"gitlab.com/abiewardani/scaffold/internal/system/connection"
	"gitlab.com/abiewardani/scaffold/pkg/logger"
)

type Tfunc func(ctx context.Context) error

type TxRepository interface {
	Do(context.Context, Tfunc) error
	Begin(ctx context.Context) context.Context
	Exec(ctx context.Context, err error) error
}

type txRepository struct {
	conn connection.Connection
}

func NewTxRepository(sys system.System) TxRepository {
	return &txRepository{
		conn: sys.Conn,
	}
}

func (c *txRepository) Begin(ctx context.Context) context.Context {
	tx := c.conn.DB().Master().Begin()
	ctx = c.conn.DB().AddTx(ctx, tx)

	return ctx
}

func (c *txRepository) Exec(ctx context.Context, err error) error {
	tx := c.conn.DB().GetTx(ctx)
	if tx == nil {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		// if error, rollback
		if errRollback := tx.Rollback().Error; errRollback != nil {
			logger.E(fmt.Sprintf("Error Rollback Transaction: %v", errRollback.Error()))
		}

		return err
	}

	// if no error, commit
	if errCommit := tx.Commit().Error; errCommit != nil {
		logger.E(fmt.Sprintf("Error Commit Transaction: %v", errCommit.Error()))
	}

	return nil
}

func (c *txRepository) Do(ctx context.Context, tFunc Tfunc) error {
	var err error
	// begin transaction
	tx := c.conn.DB().Master().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// run callback
	err = tFunc(c.conn.DB().AddTx(ctx, tx))
	if err != nil {
		// if error, rollback
		if errRollback := tx.Rollback().Error; errRollback != nil {
			logger.E(fmt.Sprintf("Error Rollback Transaction: %v", errRollback.Error()))
		}

		return err
	}
	// if no error, commit
	if errCommit := tx.Commit().Error; errCommit != nil {
		logger.E(fmt.Sprintf("Error Commit Transaction: %v", errCommit.Error()))
	}
	return nil
}
