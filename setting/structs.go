package setting

import (
	"encoding/json"
	"time"
)

type SuccessBody struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorBody struct {
	Status   string `json:"status"`
	ErrorMsg string `json:"errorMsg"`
	Error    string `json:"error"`
}

type KeyJSON struct {
	Address   string      `json:"address"`
	SecretKey string      `json:"secretkey"`
	Extension interface{} `json:"extension"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Mobile    string    `json:"mobile"`
	Timestamp time.Time `json:"timestamp"`
}

func (obj *User) Marshal() ([]byte, error) {
	return json.Marshal(obj)
}

func (obj *User) Unmarshal(data []byte) error {
	return json.Unmarshal(data, obj)
}
