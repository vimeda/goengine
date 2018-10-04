package postgres

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/hellofresh/goengine/eventstore"
	"github.com/hellofresh/goengine/messaging"
)

var (
	// ErrEmptyStreamName error on empty stream name
	ErrEmptyStreamName = errors.New("stream name cannot be empty")
	// ErrNoPayloadConverter error on no payload converter provided
	ErrNoPayloadConverter = errors.New("payload converter should be provided")
)

// SingleStreamStrategy struct represents eventstore with single stream
type SingleStreamStrategy struct {
	converter eventstore.PayloadConverter
}

// NewPostgresStrategy is the constructor postgres for PersistenceStrategy interface
func NewPostgresStrategy(converter eventstore.PayloadConverter) (eventstore.PersistenceStrategy, error) {
	if converter == nil {
		return nil, ErrNoPayloadConverter
	}

	return &SingleStreamStrategy{converter: converter}, nil
}

// CreateSchema returns a valid set of SQL statements to create the event store tables and indexes
func (s *SingleStreamStrategy) CreateSchema(tableName string) []string {
	uniqueIndexName := quoteIdentifier(fmt.Sprintf(`%s_unique_index___aggregate_type__aggregate_id__aggregate_version`, tableName))
	indexName := quoteIdentifier(fmt.Sprintf(`%s_index__aggregate_type__aggregate_id`, tableName))
	tableName = quoteIdentifier(tableName)

	statements := make([]string, 3)
	statements[0] = fmt.Sprintf(
		`CREATE TABLE %s (
    no BIGSERIAL,
    event_id UUID NOT NULL,
    event_name VARCHAR(100) NOT NULL,
    payload JSON NOT NULL,
    metadata JSONB NOT NULL,
    created_at TIMESTAMP(6) NOT NULL,
    PRIMARY KEY (no),
    CONSTRAINT aggregate_version_not_null CHECK ((metadata->>'_aggregate_version') IS NOT NULL),
    CONSTRAINT aggregate_type_not_null CHECK ((metadata->>'_aggregate_type') IS NOT NULL),
    CONSTRAINT aggregate_id_not_null CHECK ((metadata->>'_aggregate_id') IS NOT NULL),
    UNIQUE (event_id)
);`,
		tableName,
	)
	statements[1] = fmt.Sprintf(
		`CREATE UNIQUE INDEX %s ON %s
((metadata->>'_aggregate_type'), (metadata->>'_aggregate_id'), (metadata->>'_aggregate_version'));`,
		uniqueIndexName,
		tableName,
	)
	statements[2] = fmt.Sprintf(
		`CREATE INDEX %s ON %s
((metadata->>'_aggregate_type'), (metadata->>'_aggregate_id'), no);`,
		indexName,
		tableName,
	)
	return statements
}

// ColumnNames returns the columns that need to be inserted into the table in the correct order
func (s *SingleStreamStrategy) ColumnNames() []string {
	return []string{"event_id", "event_name", "payload", "metadata", "created_at"}
}

// PrepareData transforms a slice of messaging into a flat interface slice with the correct column order
func (s *SingleStreamStrategy) PrepareData(messages []messaging.Message) ([]interface{}, error) {
	var out = make([]interface{}, 0, len(messages)*5) // optimization for the number of columns
	for _, msg := range messages {
		payloadType, payloadData, err := s.converter.ConvertPayload(msg.Payload())
		if err != nil {
			return nil, err
		}

		meta, err := json.Marshal(msg.Metadata())
		if err != nil {
			return nil, err
		}

		out = append(out,
			msg.UUID(),
			payloadType,
			payloadData,
			meta,
			msg.CreatedAt(),
		)
	}
	return out, nil
}

// GenerateTableName returns a valid table name for postgres
func (s *SingleStreamStrategy) GenerateTableName(streamName eventstore.StreamName) (string, error) {
	if len(streamName) == 0 {
		return "", ErrEmptyStreamName
	}

	// remove underscore at the end
	regLastUnderScores, err := regexp.Compile("_+$")
	if err != nil {
		return "", err
	}

	// remove not allowed symbols
	regNotAllowed, err := regexp.Compile("[^a-z0-9_]+")
	if err != nil {
		return "", err
	}

	name := strings.ToLower(string(streamName))
	name = regNotAllowed.ReplaceAllString(name, "")
	name = regLastUnderScores.ReplaceAllString(name, "")
	return fmt.Sprintf("events_%s", name), nil
}
