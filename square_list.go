package main

import "fmt"

type SquareList struct {
	List   []Square
	Length int
}

// lengthは辺の長さ
func NewSquareList(length int) SquareList {
	tmp := []Square{}
	for i := 1; i <= length*length; i++ {
		tmp = append(tmp, NewSquare(i, length))
	}

	return SquareList{
		List: tmp,
		Length: length,
	}
}

func (sl SquareList) PrintSquares() {
	for _, v := range sl.List {
		v.PrintSquare()
		fmt.Println()
	}
}

func (sl SquareList) GenOneInTimeCond() []string {
	retval := []string{}
	for _, v := range sl.List {
		retval = append(retval, v.GenOneInTimeCondition()...)
	}
	return retval
}

func (sl SquareList) GenPossibleMoveCond() []string {
	retval := []string{}
	squareNum := len(sl.List)

	// 現在と次の状態を見る
	for i, current := range sl.List {
		if i+1 == squareNum {
			break
		}
		next := sl.List[i+1]

		// 各座標を見る
		allPoints := current.AllPoints()
		for _, focusedPoint := range allPoints {
			focusedName := current.Point2Name(focusedPoint)
			nextPoints := current.NextPoints(focusedPoint)

			for _, nextPoint := range nextPoints {
				nextName := next.Point2Name(nextPoint)
				cond := fmt.Sprintf("%v %v 0", -focusedName, -nextName)
				retval = append(retval, cond)
			}
		}
	}

	return retval
}
