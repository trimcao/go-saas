package cache

import (
	"bytes"
	"encoding/gob"
)

// LogRequest adds a new item to the list of pending requests to be logged
func LogRequest(v interface{}) (int64, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(v); err != nil {
		return 0, err
	}

	if _, err := rc.RPush("reqlog", buf.String()).Result(); err != nil {
		return 0, err
	}
	return 0, nil
}

// DequeueRequests returns all pending requests ready to be inserted into the database
func DequeueRequests() ([]string, error) {
	return rc.LRange("reqlog", 0, -1).Result()
}
