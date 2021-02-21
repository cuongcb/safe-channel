package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishWatch(t *testing.T) {
	ch := New()
	tcs := []struct {
		input interface{}
		want  interface{}
	}{
		{"a", "a"},
		{"abc", "abc"},
		{"", ""},
	}

	go func() {
		for _, tc := range tcs {
			ch.Publish(tc.input)
		}
		ch.Close()
	}()

	i := 0
	for got := range ch.Watch() {
		assert.EqualValues(t, got, tcs[i].want, "expected %v, got %v", tcs[i].want, got)
		i++
	}
}

func TestWriteOnClosedChannel(t *testing.T) {
	ch := New()
	ch.Close()

	writeOnClosedFunc := func(t *testing.T, str string) {
		defer func() {
			err := recover()
			assert.Nil(t, err, "panic: %s", err)
		}()

		ch.Publish(str)
	}

	t.Run("write x", func(t *testing.T) {
		writeOnClosedFunc(t, "x")
	})

	t.Run("write y", func(t *testing.T) {
		writeOnClosedFunc(t, "y")
	})

	t.Run("write z", func(t *testing.T) {
		writeOnClosedFunc(t, "z")
	})

	// cannot read here
	t.Run("watch-on-closed", func(t *testing.T) {
		_, ok := <-ch.Watch()
		assert.Equal(t, ok, false, "[1] channel is open? %t", ok)
		assert.Equal(t, ok, false, "[2] channel is open? %t", ok)
	})
}

func TestCloseOnClosedChannel(t *testing.T) {
	ch := New()
	ch.Close()

	t.Run("reclose", func(t *testing.T) {
		defer func(t *testing.T) {
			err := recover()
			assert.Nil(t, err, "panic: %s", err)
		}(t)

		// reclose, no panic
		ch.Close()
		ch.Close()
	})
}
