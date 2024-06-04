package main

import (
	"fmt"
)

// type functions
type Mod func() map[int32]string
type Pattern func(mods ...Mod)

const (
	colorStart        = "\033[%sm"
	colorFinish       = "\033[0m"
	Red         int32 = 31
	Green       int32 = 32
	Yellow      int32 = 33
	Blue        int32 = 34
)

// global variables
var size int32 = 15
var color int32 = 32
var char rune = '*'

func sizeFunc() map[int32]string {
	m := make(map[int32]string)
	m[15] = "size"
	return m
}

func charFunc() map[rune]string {
	m := make(map[rune]string)
	m['*'] = "char"
	return m
}

func colorFunc() map[int32]string {
	m := make(map[int32]string)
	m[Blue] = "color"
	return m
}

func main() {
	sandglass(diagonalHourglass, charFunc, sizeFunc, colorFunc)
}

func sandglass(p Pattern, mods ...Mod) {
	for _, m := range mods {
		p(m)
	}
}

// hollow sandglass
func hollowSandglass(mods ...Mod) {
	for _, m := range mods {
		validateParameter(m)
	}

	startColor := fmt.Sprintf(fmt.Sprintf("%s", colorStart), fmt.Sprintf("%d", color))

	var numSpaces int32 = 2
	var numChars int32 = size
	var i int32
	var j int32

	for i = 1; i <= size; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j == numChars {
				fmt.Print(startColor + string(char) + colorFinish)
			} else if i == 1 || i == size {
				fmt.Print(startColor + string(char) + colorFinish)
			} else {
				fmt.Print(" ")
			}
		}
		if i <= size/2 {
			numSpaces++
			numChars -= 2
		} else {
			numSpaces--
			numChars += 2
		}
		fmt.Println()
	}
}

// half full sandglass
func upperHalfFullSandglass(mods ...Mod) {
	for _, m := range mods {
		validateParameter(m)
	}

	startColor := fmt.Sprintf(fmt.Sprintf("%s", colorStart), fmt.Sprintf("%d", color))

	var numSpaces int32 = 2
	var numChars int32 = size
	var i int32
	var j int32

	for i = 1; i <= size/2; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j == numChars {
				fmt.Print(startColor + string(char) + colorFinish)
			} else if i == 1 || i == size {
				fmt.Print(startColor + string(char) + colorFinish)
			} else {
				fmt.Print(" ")
			}
		}
		numSpaces++
		numChars -= 2
		fmt.Println()
	}

	for i = 1; i <= size/2+1; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			fmt.Print(startColor + string(char) + colorFinish)
		}
		numSpaces--
		numChars += 2
		fmt.Println()
	}
	fmt.Println()
}

// diagonal sandglass
func diagonalHourglass(mods ...Mod) {
	var size int32 = 15
	var color int32 = 32
	var char rune = '*'

	for _, m := range mods {
		validateParameter(m)
	}

	startColor := fmt.Sprintf(fmt.Sprintf("%s", colorStart), fmt.Sprintf("%d", color))

	var numSpaces int32 = 2
	var numChars int32 = size
	var i int32
	var j int32

	for i = 1; i <= size/2; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j > numChars/2 {
				fmt.Print(startColor + string(char) + colorFinish)
			} else if i == 1 {
				fmt.Print(startColor + string(char) + colorFinish)
			} else {
				fmt.Print(" ")
			}
		}
		numSpaces++
		numChars -= 2
		fmt.Println()
	}

	for i = 1; i <= size/2+1; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j <= numChars/2+1 || j == numChars {
				fmt.Print(startColor + string(char) + colorFinish)
			} else if i == 1 || i == size/2+1 {
				fmt.Print(startColor + string(char) + colorFinish)
			} else {
				fmt.Print(" ")
			}
		}
		numSpaces--
		numChars += 2
		fmt.Println()
	}
	fmt.Println()
}

// param validator
func validateParameter(m Mod) {
	for key, val := range m() {
		if val == "size" {
			size = key
		} else if val == "color" {
			color = key
		} else if val == "char" {
			char = key
		}
	}
}
