package ops

import (
	"litevents/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrCreateCursor(t *testing.T) {
	s := store.NewSqliteStore(":memory:")

	id, err := GetOrCreateCursor(s, "test", "test")
	assert.Nil(t, err)
	assert.Equal(t, id, int64(1))
}
