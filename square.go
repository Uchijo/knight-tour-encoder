package main

import (
	"fmt"
	"strconv"
)

type Square struct {
	Matrix [][]int
	Length int
	Start  int
	Time   int
}

func (s Square) PrintSquare() {
	fmt.Printf("[info] time: %v\n", s.Time)
	for i := 0; i < s.Length; i++ {
		for j := 0; j < s.Length; j++ {
			fmt.Printf("%5v", s.Matrix[i][j])
		}
		fmt.Println()
	}
}

// 次に移動しうる場所
func (s Square) NextPoints(current Point) []Point {
	return current.MovablePoints(s.Length)
}

// 座標からそれに対応するリテラルを取り出す
func (s Square) Point2Name(point Point) int {
	return s.Matrix[point.X][point.Y]
}

func (s Square) Name2Point(name int) Point {
	for i, outer := range s.Matrix {
		for j, inner := range outer {
			if inner == name {
				return Point{X: i, Y: j}
			}
		}
	}
	panic("something went wrong")
}

func (s Square) AllPoints() []Point {
	retval := []Point{}
	for i, outer := range s.Matrix {
		for j := range outer {
			retval = append(retval, Point{X: i, Y: j})
		}
	}
	return retval
}

// 1以上かつ1以下ならok
func (s Square) GenOneInTimeCondition() []string {
	tmp := []string{
		s.genMoreThanOneCond(),
	}
	tmp = append(tmp, s.genLessThanOneCond()...)
	return tmp
}

// 1以上
func (s Square) genMoreThanOneCond() string {
	tmp := ""
	end := s.Start + s.Length*s.Length
	for i := s.Start; i < end; i++ {
		tmp += strconv.Itoa(i) + " "
	}
	tmp += "0"

	return tmp
}

func (s Square) AllNames() []int {
	tmp := []int{}
	end := s.Start + s.Length*s.Length
	for i := s.Start; i < end; i++ {
		tmp = append(tmp, i)
	}
	return tmp
}

// 1以下
// 任意に2個選んで否定にしてorで繋いだ物が真 -> これらペアがすべて真なら1以下
// 2個選んでandを取ってそのペアをorで繋ぐ -> 真になる場合、2個以上真のリテラルが存在する。これにド・モルガンを適用
func (s Square) genLessThanOneCond() []string {
	names := s.AllNames()
	namesLen := len(names)
	if namesLen == 0 {
		panic("something went wrong")
	}

	start := s.Start
	end := start + s.Length * s.Length

	clauses := []string{}
	for i := start; i<end; i++ {
		for j := i + 1; j<end; j++ {
			line := fmt.Sprintf("%v %v 0", -i, -j)
			clauses = append(clauses, line)
		}
	}

	return clauses
}

// timeは時間、lengthは一辺の長さ
// timeは1始まり
func NewSquare(time, length int) Square {
	matrix := [][]int{}
	// startは1始まり
	start := length*length*(time-1) + 1
	current := start

	// matrixを埋める
	for i := 0; i < length; i++ {
		matrix = append(matrix, []int{})
		for j := 0; j < length; j++ {
			matrix[i] = append(matrix[i], current)
			current += 1
		}
	}

	return Square{
		Matrix: matrix,
		Length: length,
		Start:  start,
		Time:   time,
	}
}
