package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	sqlitestore := NewSqliteStore(":memory:")

	id, err := sqlitestore.CreateEvent("test", "test")
	assert.NoError(t, err)
	assert.Equal(t, id, int64(1))

	id, err = sqlitestore.CreateEvent("test", "test")
	assert.NoError(t, err)
	assert.Equal(t, id, int64(2))

	id, err = sqlitestore.CreateEvent("", "test")
	assert.Error(t, err)
	assert.Equal(t, id, int64(0))

	id, err = sqlitestore.CreateEvent("test", "")
	assert.Error(t, err)
	assert.Equal(t, id, int64(0))
}

func TestCreateCursor(t *testing.T) {
	sqlitestore := NewSqliteStore(":memory:")

	id, err := sqlitestore.CreateCursor("test", "test")
	assert.NoError(t, err)
	assert.Equal(t, id, int64(1))

	id, err = sqlitestore.CreateCursor("test2", "test")
	assert.NoError(t, err)
	assert.Equal(t, id, int64(2))

	id, err = sqlitestore.CreateCursor("test2", "test")
	assert.Error(t, err)
	assert.Equal(t, id, int64(0))

	id, err = sqlitestore.CreateCursor("", "test")
	assert.Error(t, err)
	assert.Equal(t, id, int64(0))

	id, err = sqlitestore.CreateCursor("test3", "")
	assert.Error(t, err)
	assert.Equal(t, id, int64(0))
}

func TestGetNextEvent(t *testing.T) {
	sqlitestore := NewSqliteStore(":memory:")

	cursor_id, err := sqlitestore.CreateCursor("test", "test")
	assert.NoError(t, err)
	assert.Equal(t, cursor_id, int64(1))

	event_id, err := sqlitestore.CreateEvent("test", "test")
	assert.NoError(t, err)
	assert.Equal(t, event_id, int64(1))

	eventID, data, err := sqlitestore.GetNextEvent("test")
	assert.NoError(t, err)
	assert.Equal(t, eventID, int64(1))
	assert.Equal(t, data, "test")
}

func TestAckEvent(t *testing.T) {
	sqlitestore := NewSqliteStore(":memory:")

	cursor_id, err := sqlitestore.CreateCursor("test", "test")
	assert.NoError(t, err)
	assert.Equal(t, cursor_id, int64(1))

	_, err = sqlitestore.CreateEvent("test", "test 0")
	assert.NoError(t, err)

	_, err = sqlitestore.CreateEvent("test", "test 1")
	assert.NoError(t, err)

	eventID, data, err := sqlitestore.GetNextEvent("test")
	assert.NoError(t, err)
	assert.Equal(t, eventID, int64(1))
	assert.Equal(t, data, "test 0")

	err = sqlitestore.AckEvent("test", eventID)
	assert.NoError(t, err)

	eventID, data, err = sqlitestore.GetNextEvent("test")
	assert.NoError(t, err)
	assert.Equal(t, eventID, int64(2))
	assert.Equal(t, data, "test 1")

	err = sqlitestore.AckEvent("test", eventID)
	assert.NoError(t, err)

	eventID, data, err = sqlitestore.GetNextEvent("test")
	assert.Error(t, err)
}
