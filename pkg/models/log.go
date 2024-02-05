package audit

import (
	"time"
)

const (
	ENTITY_USER     = "USER"
	ENTITY_COMPUTER = "COMPUTER"

	ACTION_CREATE   = "CREATE"
	ACTION_GET      = "GET"
	ACTION_UPDATE   = "UPDATE"
	ACTION_DELETE   = "DELETE"
	ACTION_REGISTER = "REGISTER"
	ACTION_LOGIN    = "LOGIN"
)

type LogItem struct {
	Entity    string    `json:"entity"`
	Action    string    `json:"action"`
	EntityID  int64     `json:"entity_id"`
	Timestamp time.Time `json:"timestamp"`
}
