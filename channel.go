package channel

import "sync/atomic"

// Channel wraps the built-in channel and provide safe operations beyond it
type Channel struct {
	state int32
	ch    chan interface{}
}

// New creates a new Channel
func New() *Channel {
	return &Channel{
		state: 0,
		ch:    make(chan interface{}),
	}
}

// NewBufferedChannel creates a buffered Channel
func NewBufferedChannel(n int) *Channel {
	return &Channel{
		state: 0,
		ch:    make(chan interface{}, n),
	}
}

// Publish sends an event to channel
// If channel is closed, Publish does nothing
func (c *Channel) Publish(x interface{}) {
	if atomic.CompareAndSwapInt32(&c.state, 0, 0) {
		c.ch <- x
	}
}

// Watch keeps listening on a channel
func (c *Channel) Watch() <-chan interface{} {
	return c.ch
}

// Close closes an active channel
// This method can be called many times without risks
func (c *Channel) Close() {
	if atomic.CompareAndSwapInt32(&c.state, 0, 1) {
		close(c.ch)
	}
}
