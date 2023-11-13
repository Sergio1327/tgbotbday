package gensql

import (
	"database/sql"
	"tgbot/internal/entity/global"

	"github.com/jmoiron/sqlx"
)

func Select[T any](tx *sqlx.Tx, sqlQuery string, params ...interface{}) ([]T, error) {
	data := make([]T, 0)

	err := tx.Select(&data, sqlQuery, params...)

	if err == nil && len(data) == 0 {
		err = sql.ErrNoRows
	}

	switch err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return nil, global.ErrNoData
	default:
		return nil, err
	}
}

func SelectNamed[T any](tx *sqlx.Tx, sqlQuery string, params map[string]interface{}) ([]T, error) {
	data := make([]T, 0)

	stmt, err := tx.PrepareNamed(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.Select(&data, params)
	if err != nil {
		return nil, err
	}

	if err == nil && len(data) == 0 {
		err = sql.ErrNoRows
	}

	switch err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return nil, global.ErrNoData
	default:
		return nil, err
	}
}

func SelectNamedStruct[T any, S any](tx *sqlx.Tx, sqlQuery string, s S) ([]T, error) {
	data := make([]T, 0)

	stmt, err := tx.PrepareNamed(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.Select(&data, &s)
	if err != nil {
		return nil, err
	}

	if err == nil && len(data) == 0 {
		err = sql.ErrNoRows
	}

	switch err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return nil, global.ErrNoData
	default:
		return nil, err
	}
}

func SelectListParam[T any, L comparable](tx *sqlx.Tx, sqlQuery string, list []L) ([]T, error) {
	data := make([]T, 0)

	query, args, err := sqlx.In(sqlQuery, list)
	if err != nil {
		return nil, err
	}

	query = tx.Rebind(query)

	err = tx.Select(&data, query, args...)
	if err != nil {
		return nil, err
	}

	if err == nil && len(data) == 0 {
		err = sql.ErrNoRows
	}

	switch err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return nil, global.ErrNoData
	default:
		return nil, err
	}
}
