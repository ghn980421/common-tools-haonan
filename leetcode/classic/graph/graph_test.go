package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_200(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := buildGraph("[[\"1\",\"1\",\"1\",\"1\",\"0\"],[\"1\",\"1\",\"0\",\"1\",\"0\"],[\"1\",\"1\",\"0\",\"0\",\"0\"],[\"0\",\"0\",\"0\",\"0\",\"0\"]]")
		assert.Equal(t, 1, numIslands(case1))
	})

	t.Run("case2", func(t *testing.T) {
		case2 := buildGraph("[[\"1\",\"1\",\"0\",\"0\",\"0\"],[\"1\",\"1\",\"0\",\"0\",\"0\"],[\"0\",\"0\",\"1\",\"0\",\"0\"],[\"0\",\"0\",\"0\",\"1\",\"1\"]]")
		assert.Equal(t, 3, numIslands(case2))
	})

	t.Run("case3", func(t *testing.T) {
		case3 := buildGraph("[[\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"0\",\"1\",\"1\"],[\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\"],[\"1\",\"0\",\"1\",\"1\",\"1\",\"0\",\"0\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\"],[\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"0\",\"1\",\"1\",\"1\",\"1\",\"0\",\"0\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"],[\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\",\"1\"]]")
		assert.Equal(t, 1, numIslands(case3))
	})

}

func Test_130(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		case1 := buildGraph("[[\"X\",\"X\",\"X\",\"X\"],[\"X\",\"O\",\"O\",\"X\"],[\"X\",\"X\",\"O\",\"X\"],[\"X\",\"O\",\"X\",\"X\"]]")
		solve(case1)
		assert.Equal(t, buildGraph("[[\"X\",\"X\",\"X\",\"X\"],[\"X\",\"X\",\"X\",\"X\"],[\"X\",\"X\",\"X\",\"X\"],[\"X\",\"O\",\"X\",\"X\"]]"), case1)
	})

	t.Run("case2", func(t *testing.T) {
		case2 := buildGraph("[[\"O\",\"O\",\"O\",\"O\",\"X\",\"X\"],[\"O\",\"O\",\"O\",\"O\",\"O\",\"O\"],[\"O\",\"X\",\"O\",\"X\",\"O\",\"O\"],[\"O\",\"X\",\"O\",\"O\",\"X\",\"O\"],[\"O\",\"X\",\"O\",\"X\",\"O\",\"O\"],[\"O\",\"X\",\"O\",\"O\",\"O\",\"O\"]]")
		solve(case2)
		assert.Equal(t, buildGraph("[[\"O\",\"O\",\"O\",\"O\",\"X\",\"X\"],[\"O\",\"O\",\"O\",\"O\",\"O\",\"O\"],[\"O\",\"X\",\"O\",\"X\",\"O\",\"O\"],[\"O\",\"X\",\"O\",\"O\",\"X\",\"O\"],[\"O\",\"X\",\"O\",\"X\",\"O\",\"O\"],[\"O\",\"X\",\"O\",\"O\",\"O\",\"O\"]]"),
			case2)
	})
}
