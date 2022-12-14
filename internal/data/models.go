package data

import (
	"database/sql"
	"errors"
)

var ErrRecodeNotFound = errors.New("record not found")
var ErrEditConflict = errors.New("edit conflict")

type Models struct {
	Movies MovieModel
	Users  UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db},
	}
}
