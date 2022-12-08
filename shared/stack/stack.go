package stack

import (
	"math"
)

type node struct {
  value string
  prev *node
}

type Stack struct {
  length int
  head *node
}

func (s *Stack) Push(item string)  {
  nd := &node{value: item}
  s.length += 1

  if s.head == nil {
    s.head = nd
    return
  }

  nd.prev = s.head
  s.head = nd
}

func (s *Stack) Pop() string {
  s.length = int(math.Max(0, float64(s.length - 1)))

  if s.length == 0 {
    head := s.head
    s.head = nil

    if head == nil {
      return ""
    }

    return head.value
  }

  head := s.head
  s.head = s.head.prev

  return head.value
}

func (s Stack) Pick() string {
  if s.head == nil {
    return ""
  }

  return s.head.value
}

func NewStack() Stack {
  return Stack{}
}
