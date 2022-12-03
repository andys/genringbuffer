package genringbuffer

import "time"

// Ringbuffer is a fixed length FIFO queue that drops new items when full and allows non-blocking polling for new items with optional timeout
type Ringbuffer[T any] chan T

// NewRingbuffer creates a new ringbuffer with the given fixed number of items.
func NewRingBuffer[T any](size int) Ringbuffer[T] {
	return Ringbuffer[T](make(chan T, size))
}

// Offer adds an item to the ringbuffer. If the buffer is full, it returns immediately. Returns true if added.
func (r Ringbuffer[T]) Offer(input T) bool {
	select {
	case r <- input:
		return true
	default:
		return false
	}
}

// Poll gets an item from the ringbuffer, or blocks until one is available, until dur duration passed. Returns true if an item was returned.
func (r Ringbuffer[T]) Poll(dur time.Duration) (obj T, res bool) {
	select {
	case obj = <-r:
		res = true
	case <-time.After(dur):
		res = false
	}
	return
}

// Get gets an item from the ringbuffer, or blocks until one is available.
func (r Ringbuffer[T]) Get() T {
	return <-r
}

// Len returns the number of items in the ringbuffer.
func (r Ringbuffer[T]) Len() int {
	return len(r)
}

// Cap returns the capacity of the ringbuffer.
func (r Ringbuffer[T]) Cap() int {
	return cap(r)
}

// Close closes the ringbuffer.
func (r Ringbuffer[T]) Close() {
	close(r)
}
