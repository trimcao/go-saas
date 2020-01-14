// +build !mgo

package model

import "database/sql"

type Connection = sql.DB
type Key = int64

func Open(options ...string) (*sql.DB, error) {
	conn, err := sql.Open(options[0], options[1])
	if err != nil {
		return nil, err
	} else if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
