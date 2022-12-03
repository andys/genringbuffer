package genringbuffer

import (
	"testing"

	"github.com/zeebo/assert"
)

func TestNewRingBuffer(t *testing.T) {
	r := NewRingBuffer[int](10)
	assert.Equal(t, r.Cap(), 10)
	assert.Equal(t, r.Len(), 0)
}

func TestOffer(t *testing.T) {
	r := NewRingBuffer[int](1)
	assert.Equal(t, r.Offer(1), true)
	assert.Equal(t, r.Offer(2), false)
}

func TestPoll(t *testing.T) {
	r := NewRingBuffer[int](1)
	r.Offer(1)
	obj, res := r.Poll(0)
	assert.Equal(t, res, true)
	assert.Equal(t, obj, 1)

	obj, res = r.Poll(0)
	assert.Equal(t, res, false)
	assert.Equal(t, obj, 0)
}

func TestGet(t *testing.T) {
	r := NewRingBuffer[int](1)
	r.Offer(1)
	obj := r.Get()
	assert.Equal(t, obj, 1)
}

func TestLen(t *testing.T) {
	r := NewRingBuffer[int](1)
	assert.Equal(t, r.Len(), 0)
	r.Offer(1)
	assert.Equal(t, r.Len(), 1)
	r.Get()
	assert.Equal(t, r.Len(), 0)
}

func TestCap(t *testing.T) {
	r := NewRingBuffer[int](10)
	assert.Equal(t, r.Cap(), 10)
}

func TestClose(t *testing.T) {
	r := NewRingBuffer[int](10)
	r.Close()
}
