package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	inputStrByte, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	inputStr := string(inputStrByte)

	x, y := counter(string(inputStr))
	strX := convertNbr(x)
	strY := convertNbr(y)

	str := []string{}

	quads := []struct {
		raidName string
		check    func(int, int) string
	}{
		{"[quadA]", QuadA},
		{"[quadB]", QuadB},
		{"[quadC]", QuadC},
		{"[quadD]", QuadD},
		{"[quadE]", QuadE},
	}

	// Loop through each quad
	for _, quad := range quads {
		if quad.check(x, y) == inputStr {
			str = append(str, quad.raidName+" "+"["+strX+"]"+" "+"["+strY+"]"+" ")
		}
	}
	if len(str) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	for i := 0; i <= len(str)-2; i++ {
		fmt.Print(str[i] + "||" + " ")
	}
	fmt.Println(str[len(str)-1])

	// Print the input string
	fmt.Println("Input received:")
	fmt.Print(string(inputStr))
}

func convertNbr(n int) string {
	res := ""
	sign := ""

	if n < 0 {
		sign = "-"
		n = -n
	}

	if n == 0 {
		return "0"
	}

	for n > 0 {
		res = string(48+n%10) + res
		n /= 10
	}

	return sign + res
}

func counter(output string) (x, y int) {
	countX := 0
	countY := 0
	flag := true
	for _, s := range output {
		if s == '\n' {
			countY++
			flag = false
		} else {
			if flag {
				countX++
			}
		}
	}
	return countX, countY
}

func QuadA(x, y int) string {
	return quad(x, y, "o", "o", "o", "o", "-", "-", "|", "|")
}

func QuadB(x, y int) string {
	return quad(x, y, "/", "\\", "\\", "/", "*", "*", "*", "*")
}

func QuadC(x, y int) string {
	return quad(x, y, "A", "A", "C", "C", "B", "B", "B", "B")
}

func QuadD(x, y int) string {
	return quad(x, y, "A", "C", "A", "C", "B", "B", "B", "B")
}

func QuadE(x, y int) string {
	return quad(x, y, "A", "C", "C", "A", "B", "B", "B", "B")
}

func quad(x, y int, topLeft, topRight, bottomLeft, bottomRight, topBorder, bottomBorder, leftBorder, rightBorder string) string {
	if x < 1 || y < 1 {
		return "Error"
	}

	res := ""

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			// top left
			if i == 0 && j == 0 {
				res += topLeft
				continue
			}
			// top right
			if i == 0 && j == x-1 {
				res += topRight
				continue
			}
			// bottom left
			if i == y-1 && j == 0 {
				res += bottomLeft
				continue
			}
			// bottom right
			if i == y-1 && j == x-1 {
				res += bottomRight
				continue
			}
			// top border
			if i == 0 && j < x-1 {
				res += topBorder
				continue
			}
			// bottom border
			if i == y-1 && j < x-1 {
				res += bottomBorder
				continue
			}
			// left border
			if i < y-1 && j == 0 {
				res += leftBorder
				continue
			}
			// right border
			if i < y-1 && j == x-1 {
				res += rightBorder
				continue
			}

			res += " "
		}
		res += "\n"
	}

	return res
}
