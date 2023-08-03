package main

import "fmt"

func main() {
	squareList := NewSquareList(4)
	squareList.PrintSquares()

	fmt.Println(" ======== hogehoge~ ======== ")

	cond := squareList.GenPossibleMoveCond()
	for _, v := range cond {
		fmt.Printf("%v\n", v)
	}
}
