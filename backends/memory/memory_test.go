package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testKey   = "hello"
	testValue = "world"
)

func TestGet(t *testing.T) {
	s := NewServer()
	s.store.Store(testKey, []byte(testValue))

	response, err := s.get(testKey)
	assert.Nil(t, err)
	assert.Equal(t, testValue, string(response))
}

func TestSet(t *testing.T) {
	s := NewServer()
	err := s.set(testKey, []byte(testValue))
	assert.Nil(t, err)
	result, ok := s.store.Load(testKey)
	assert.True(t, ok)

	b, _ := result.([]byte)
	assert.Equal(t, testValue, string(b))
}

func TestDel(t *testing.T) {
	s := NewServer()
	s.store.Store(testKey, []byte(testValue))

	response, err := s.get(testKey)
	assert.Nil(t, err)
	assert.Equal(t, testValue, string(response))

	err = s.del(testKey)
	assert.Nil(t, err)

	_, ok := s.store.Load(testKey)
	assert.False(t, ok)
}
