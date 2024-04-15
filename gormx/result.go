package gormx

import (
	"errors"
	"github.com/gophero/goal/errorx"
	"gorm.io/gorm"
)

var (
	InsertError = errors.New("insert failed, returned row count is unexpected 0")
	UpdateError = errors.New("update failed, returned row count is unexpected 0")
	DeleteError = errors.New("delete failed, returned row count is unexpected 0")
	QueryError  = errors.New("query failed, returned row count is unexpected 0")
)

func InsertResult[T any](r *gorm.DB, e T) (T, error) {
	if r.Error != nil {
		return e, r.Error
	}
	if r.RowsAffected < 1 {
		return e, InsertError
	}
	return e, nil
}

func UpdateResult[T any](r *gorm.DB, e T) (T, error) {
	if r.Error != nil {
		return e, r.Error
	}
	if r.RowsAffected < 1 {
		return e, UpdateError
	}
	return e, nil
}

func DeleteResult[T any](r *gorm.DB) (bool, error) {
	if r.Error != nil {
		return false, r.Error
	}
	if r.RowsAffected < 1 {
		return false, DeleteError
	}
	return true, nil
}

func QueryResult[T any](r *gorm.DB, t T) (T, error) {
	if r.Error != nil {
		return t, r.Error
	}
	return t, nil
}

func QueryResultFilterNotFound[T any](r *gorm.DB, t T) T {
	if r.Error != nil {
		if errors.Is(r.Error, gorm.ErrRecordNotFound) {
			return t
		}
	}
	errorx.Throw(r.Error)
	return t
}
