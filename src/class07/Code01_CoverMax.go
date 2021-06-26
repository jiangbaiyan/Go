package class07

import (
	"github.com/jiangbaiyan/go/src/class06"
	"sort"
)

// Line 给定很多线段，每个线段都有两个数[start, end]，
// 表示线段开始位置和结束位置，左右都是闭区间
// 规定：
// 1）线段的开始和结束位置一定都是整数值
// 2）线段重合区域的长度必须>=1
// 返回线段最多重合区域中，包含了几条线段
type Line struct {
	start int
	end   int
}

type Lines []Line

func (l Lines) Less(i, j int) bool {
	return l[i].start < l[j].start
}

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// MaxCover 最大线段重合长度
func MaxCover(m [][]int) int {
	lines := make(Lines, 0)
	for i := 0; i < len(m); i++ {
		lines[i] = Line{
			start: m[i][0],
			end:   m[i][1],
		}
	}
	sort.Sort(lines)
	heap := &class06.MyHeap{}
	ans := 0
	for i := 0; i < lines.Len(); i++ {
		for !heap.IsEmpty() && heap.Peek() <= lines[i].start {
			heap.Pop()
		}
		heap.Push(lines[i].end)
		ans = max(ans, heap.Size())
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
