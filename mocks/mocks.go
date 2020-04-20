package mocks

import (
	"time"

	"github.com/vimeda/goengine"
	"github.com/vimeda/goengine/metadata"
)

// In order to make sure that we have the same mocks we can regenerate them using `go generate`
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=mocks -destination event_store.go -mock_names EventStore=EventStore github.com/vimeda/goengine EventStore
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=mocks -destination event_stream.go -mock_names EventStream=EventStream github.com/vimeda/goengine EventStream
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=mocks -destination message_payload_factory.go -mock_names MessagePayloadFactory=MessagePayloadFactory github.com/vimeda/goengine MessagePayloadFactory
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=mocks -destination message_payload_converter.go -mock_names MessagePayloadConverter=MessagePayloadConverter github.com/vimeda/goengine MessagePayloadConverter
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=mocks -destination message_payload_resolver.go -mock_names MessagePayloadResolver=MessagePayloadResolver github.com/vimeda/goengine MessagePayloadResolver
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=mocks -destination query.go -mock_names Query=Query github.com/vimeda/goengine Query
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=mocks -destination message.go -mock_names Message=Message github.com/vimeda/goengine Message
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=aggregate -destination aggregate/aggregate.go -mock_names Root=Root github.com/vimeda/goengine/aggregate Root
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=aggregate -destination aggregate/aggregate_another.go -mock_names Root=AnotherRoot github.com/vimeda/goengine/aggregate Root
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=sql -destination driver/sql/execer.go -mock_names Execer=Execer github.com/vimeda/goengine/driver/sql Execer
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=sql -destination driver/sql/queryer.go -mock_names Queryer=Queryer github.com/vimeda/goengine/driver/sql Queryer
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=sql -destination driver/sql/persistence_strategy.go -mock_names PersistenceStrategy=PersistenceStrategy github.com/vimeda/goengine/driver/sql PersistenceStrategy
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=sql -destination driver/sql/projection_state_serialization.go -mock_names ProjectionStateSerialization=ProjectionStateSerialization github.com/vimeda/goengine/driver/sql ProjectionStateSerialization
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package=sql -destination driver/sql/message_factory.go -mock_names MessageFactory=MessageFactory github.com/vimeda/goengine/driver/sql MessageFactory
//go:generate go run ../vendor/github.com/golang/mock/mockgen/ -package sql -destination driver/sql/notification_queue.go -mock_names NotificationQueuer=NotificationQueuer github.com/vimeda/goengine/driver/sql NotificationQueuer

var _ goengine.Message = &DummyMessage{}

// DummyMessage a simple goengine.Message implementation used for testing
type DummyMessage struct {
	uuid      goengine.UUID
	payload   interface{}
	metadata  metadata.Metadata
	createdAt time.Time
}

// NewDummyMessage returns a new DummyMessage
func NewDummyMessage(id goengine.UUID, payload interface{}, meta metadata.Metadata, time time.Time) *DummyMessage {
	return &DummyMessage{
		id,
		payload,
		meta,
		time,
	}
}

// UUID returns the identifier of this message
func (d *DummyMessage) UUID() goengine.UUID {
	return d.uuid
}

// CreatedAt returns the created time of the message
func (d *DummyMessage) CreatedAt() time.Time {
	return d.createdAt
}

// Payload returns the payload of the message
func (d *DummyMessage) Payload() interface{} {
	return d.payload
}

// Metadata return the message metadata
func (d *DummyMessage) Metadata() metadata.Metadata {
	return d.metadata
}

// WithMetadata Returns new instance of the message with key and value added to metadata
func (d *DummyMessage) WithMetadata(key string, value interface{}) goengine.Message {
	newMessage := *d
	newMessage.metadata = metadata.WithValue(d.metadata, key, value)

	return &newMessage
}
