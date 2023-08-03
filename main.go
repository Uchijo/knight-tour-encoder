package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/uchijo/knight-tour-encoder/utils"
)

func main() {
	args := os.Args
	if len(args) <= 2 {
		panic("insufficient arguments; usage: `command size_of_grid output`")
	}
	filename := args[2]
	length, err := strconv.Atoi(args[1])
	if err != nil {
		panic(err)
	}

	squareList := utils.NewSquareList(length)
	// cond := squareList.GenGridCond()
	// for _, v := range cond {
	// 	fmt.Printf("%v\n", v)
	// }
	cond := squareList.GenCompleteCond()
	varNum := squareList.VarNum()
	err = saveFile(filename, varNum, cond)
}

func saveFile(filename string, varNum int, cond []string) error {
	clauseNum := len(cond) - 3
	lines := fmt.Sprintf("p cnf %v %v\n", varNum, clauseNum)
	for _, v := range cond {
		lines += v + "\n"
	}

	fmt.Println(lines)
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(lines)
	if err != nil {
		panic(err)
	}
	return nil
}
