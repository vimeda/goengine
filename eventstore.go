package goengine

import (
	"context"

	"github.com/vimeda/goengine/metadata"
)

type (
	// StreamName is the name of an event stream
	StreamName string

	// EventStream is the result of an event store query. Its cursor starts before the first row
	// of the result set. Use Next to advance through the results:
	EventStream interface {
		// Next prepares the next result for reading.
		// It returns true on success, or false if there is no next result row or an error happened while preparing it.
		// Err should be consulted to distinguish between the two cases.
		Next() bool

		// Err returns the error, if any, that was encountered during iteration.
		Err() error

		// Close closes the EventStream, preventing further enumeration. If Next is called
		// and returns false and there are no further result sets,
		// result of Err. Close is idempotent and does not affect the result of Err.
		Close() error

		// Message returns the current message and it's number within the EventStream.
		Message() (Message, int64, error)
	}

	// EventStore an interface describing an event store
	EventStore interface {
		ReadOnlyEventStore

		// Create creates an event stream
		Create(ctx context.Context, streamName StreamName) error

		// AppendTo appends the provided messages to the stream
		AppendTo(ctx context.Context, streamName StreamName, streamEvents []Message) error
	}

	// ReadOnlyEventStore an interface describing a readonly event store
	ReadOnlyEventStore interface {
		// HasStream returns true if the stream exists
		HasStream(ctx context.Context, streamName StreamName) bool

		// Load returns a list of events based on the provided conditions
		Load(ctx context.Context, streamName StreamName, fromNumber int64, count *uint, metadataMatcher metadata.Matcher) (EventStream, error)
	}
)

// ReadEventStream reads the entire event stream and returns it's content as a slice.
// The main purpose of the function is for testing and debugging.
func ReadEventStream(stream EventStream) ([]Message, []int64, error) {
	var messages []Message
	var messageNumbers []int64
	for stream.Next() {
		msg, msgNumber, err := stream.Message()
		if err != nil {
			return nil, nil, err
		}

		messages = append(messages, msg)
		messageNumbers = append(messageNumbers, msgNumber)
	}

	if err := stream.Err(); err != nil {
		return nil, nil, err
	}

	return messages, messageNumbers, nil
}
