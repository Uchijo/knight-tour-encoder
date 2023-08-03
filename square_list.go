package main

import "fmt"

type SquareList []Square

// lengthは辺の長さ
func NewSquareList(length int) SquareList {
	tmp := []Square{}
	for i := 1; i <= length*length; i++ {
		tmp = append(tmp, NewSquare(i, length))
	}

	return tmp
}

func (sl SquareList) PrintSquares() {
	for _, v := range sl {
		v.PrintSquare()
		fmt.Println()
	}
}

func (sl SquareList) GenOneInTimeCond() []string {
	retval := []string{}
	for _, v := range sl {
		retval = append(retval, v.GenOneInTimeCondition()...)
	}
	return retval
}
