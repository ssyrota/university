package db

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "gorm.io/driver/postgres"
)

type SqlDB struct {
	*sqlx.DB
	Wildcard string
}

func (db *SqlDB) RebindQuery(baseQuery string, data map[string]any) (string, []interface{}, error) {
	query, args, err := sqlx.Named(baseQuery, data)
	if err != nil {
		return "", nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return "", nil, err
	}
	query = db.Rebind(query)
	return query, args, nil
}

func (db *SqlDB) Close() error {
	return db.DB.Close()
}

var pgWildcard = "%"

func NewWithPg(connStr string) (*SqlDB, error) {
	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxIdleTime(1 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Minute)

	return &SqlDB{db, pgWildcard}, nil
}
