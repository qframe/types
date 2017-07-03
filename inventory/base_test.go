package inventory

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


const (
	inv1 = `{"time": 0, "subject":"node1", "object": "node2", "action": "connected", "tags": {}}`
)

func TestNewBaseFromJson(t *testing.T) {
	b, err := NewBaseFromJson(inv1)
	assert.NoError(t, err)
	assert.Equal(t, "node1", b.Subject)
}

func TestSplitUnixNano(t *testing.T) {
	now := 1257894000000000011
	s, n := SplitUnixNano(int64(now))
	assert.Equal(t, 1257894000, s)
	assert.Equal(t, 11, n)
}
