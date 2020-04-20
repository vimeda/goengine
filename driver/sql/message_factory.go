package sql

import (
	"database/sql"

	"github.com/vimeda/goengine"
)

// MessageFactory reconstruct messages from the database
type MessageFactory interface {
	// CreateEventStream reconstructs the message from the provided rows
	CreateEventStream(rows *sql.Rows) (goengine.EventStream, error)
}
