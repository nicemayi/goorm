package session

import (
	"database/sql"
	"strings"
)

type Sessoin struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB) *