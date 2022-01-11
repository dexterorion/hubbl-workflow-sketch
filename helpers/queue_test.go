package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q := StartEmptyQueue()
	_, err := q.Pop()

	require.Equal(t, ErrPopingEmptyQueue, err)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	require.Equal(t, int64(5), q.Size())
	require.Equal(t, "[1 2 3 4 5]", q.String())
	require.Equal(t, []interface{}{1, 2, 3, 4, 5}, q.AsArray())

	el, _ := q.Pop()
	require.Equal(t, 1, el)
	require.Equal(t, int64(4), q.Size())
	require.Equal(t, "[2 3 4 5]", q.String())
	require.Equal(t, []interface{}{2, 3, 4, 5}, q.AsArray())

	q.Push(1)
	require.Equal(t, int64(5), q.Size())
	require.Equal(t, "[2 3 4 5 1]", q.String())
	require.Equal(t, []interface{}{2, 3, 4, 5, 1}, q.AsArray())
}
