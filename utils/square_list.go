package utils

import (
	"fmt"
	"strconv"
)

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
		List:   tmp,
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

			cond := strconv.Itoa(-focusedName) + " "
			for _, nextPoint := range nextPoints {
				nextName := next.Point2Name(nextPoint)
				cond += strconv.Itoa(nextName) + " "
			}
			cond += "0"
			retval = append(retval, cond)
		}
	}

	return retval
}

// 同一座標上で時間が別になっているものを集約したリストを返す
func (sl SquareList) ExtractGrid() [][]int {
	// 識別子の最大値
	identifierMax := len(sl.List) * sl.Length * sl.Length
	gridNum := sl.Length * sl.Length
	retval := make([][]int, gridNum)
	for i := 1; i <= identifierMax; i++ {
		retval[i%gridNum] = append(retval[i%gridNum], i)
	}
	return retval
}

// 複数回同じグリッドに移動しない制約
func (sl SquareList) GenGridCond() []string {
	grids := sl.ExtractGrid()
	retval := []string{}
	for _, v := range grids {
		fmt.Printf("inside v: %v\n", v)
		retval = append(retval, genLessThanOneVisit(v)...)
	}
	return retval
}

func genLessThanOneVisit(grids []int) []string {
	retval := []string{}
	gridLen := len(grids)
	for i := 0; i< gridLen; i++ {
		for j := i+1; j<gridLen; j++ {
			retval = append(retval, fmt.Sprintf("%v %v 0", -grids[i], -grids[j]))
		}
	}
	return retval
}

func (sl SquareList) GenCompleteCond() []string {
	retval := []string{}
	retval = append(retval, "c GridCond: 同じ場所に回ってこない")
	retval = append(retval, sl.GenGridCond()...)
	retval = append(retval, "c OneInTimeCond: 同じ時間に1つだけナイトが存在している")
	retval = append(retval, sl.GenOneInTimeCond()...)
	retval = append(retval, "c PossibleMoveCond: 不可能な動きをしない")
	retval = append(retval, sl.GenPossibleMoveCond()...)
	return retval
}

func (sl SquareList) VarNum() int {
	return sl.Length * sl.Length * len(sl.List)
}
