// +build mem

package model

type Connection = bool
type Key = int64

func Open(options ...string) (bool, error) {
	conn := true
	return conn, nil
}
