package geek_demo

import (
	"fmt"
	"strings"
	"testing"
)

var (
	t []int
)

const (
	a = iota
	b = "str"
	c
	d = iota
	e
	f
)
const (
	g = iota
	h
	i
)

func TestConstValue(t *testing.T) {
	fmt.Println(mainm)

}

func TestSolveNQueens2(t *testing.T) {
	res := solveNQueens2(4)
	for _, re := range res {
		fmt.Println(strings.Join(re, ""))
	}
}
