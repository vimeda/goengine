package sql

import (
	"context"

	"github.com/vimeda/goengine"
	"github.com/vimeda/goengine/metadata"
)

// ReadOnlyEventStore an interface describing a readonly event store that supports providing a SQL conn
type ReadOnlyEventStore interface {
	// LoadWithConnection returns a eventstream based on the provided constraints using the provided Queryer
	LoadWithConnection(ctx context.Context, conn Queryer, streamName goengine.StreamName, fromNumber int64, count *uint, metadataMatcher metadata.Matcher) (goengine.EventStream, error)
}
