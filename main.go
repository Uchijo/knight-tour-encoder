package main

import "fmt"

func main() {
	squareList := NewSquareList(4)
	squareList.PrintSquares()

	fmt.Println(" ======== hogehoge~ ======== ")

	cond := squareList.GenOneInTimeCond()
	for _, v := range cond {
		fmt.Printf("%v\n", v)
	}
}
