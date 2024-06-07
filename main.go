package main

import (
	"fmt"
)

// type functions
type Mod func(m *map[string]int32)
type Pattern func(mods ...Mod)

const (
	colorStart  = "\033[%sm"
	colorFinish = "\033[0m"
)

const (
	Red = iota + 31
	Green
	Yellow
	Blue
)

func sizeFunc(i int32) Mod {
	return func(m *map[string]int32) {
		(*m)["size"] = i
	}
}

func charFunc(r rune) Mod {
	return func(m *map[string]int32) {
		(*m)["char"] = r
	}
}

func colorFunc(i int32) Mod {
	return func(m *map[string]int32) {
		(*m)["color"] = i
	}
}

func main() {
	sandglass(diagonalHourglass, charFunc('*'), sizeFunc(15), colorFunc(Red))
	sandglass(upperHalfFullSandglass, charFunc('*'), sizeFunc(15), colorFunc(Yellow))
	sandglass(hollowSandglass, charFunc('*'), sizeFunc(15), colorFunc(Blue))
}

func sandglass(p Pattern, mods ...Mod) {
	p(mods...)
}

func assignparam(m *map[string]int32, mods ...Mod) {
	for _, mod := range mods {
		mod(m)
	}
}

func hollowSandglass(mods ...Mod) {
	m := make(map[string]int32)
	assignparam(&m, mods...)

	startColor := fmt.Sprintf(fmt.Sprintf("%s", colorStart), fmt.Sprintf("%d", m["color"]))
	var numSpaces int32 = 2
	var numChars int32 = m["size"]
	var i int32
	var j int32

	for i = 1; i <= m["size"]; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j == numChars {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
			} else if i == 1 || i == m["size"] {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
			} else {
				fmt.Print(" ")
			}
		}
		if i <= m["size"]/2 {
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
	m := make(map[string]int32)
	assignparam(&m, mods...)

	startColor := fmt.Sprintf(fmt.Sprintf("%s", colorStart), fmt.Sprintf("%d", m["color"]))

	var numSpaces int32 = 2
	var numChars int32 = m["size"]
	var i int32
	var j int32

	for i = 1; i <= m["size"]/2; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j == numChars {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
			} else if i == 1 || i == m["char"] {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
			} else {
				fmt.Print(" ")
			}
		}
		numSpaces++
		numChars -= 2
		fmt.Println()
	}

	for i = 1; i <= m["size"]/2+1; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			fmt.Print(startColor + string(m["char"]) + colorFinish)
		}
		numSpaces--
		numChars += 2
		fmt.Println()
	}
	fmt.Println()
}

// diagonal sandglass
func diagonalHourglass(mods ...Mod) {
	m := make(map[string]int32)
	assignparam(&m, mods...)

	startColor := fmt.Sprintf(fmt.Sprintf("%s", colorStart), fmt.Sprintf("%d", m["color"]))

	var numSpaces int32 = 2
	var numChars int32 = m["size"]
	var i int32
	var j int32

	for i = 1; i <= m["size"]/2; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j > numChars/2 {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
			} else if i == 1 {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
			} else {
				fmt.Print(" ")
			}
		}
		numSpaces++
		numChars -= 2
		fmt.Println()
	}

	for i = 1; i <= m["size"]/2+1; i++ {
		for j = 1; j <= numSpaces; j++ {
			fmt.Print(" ")
		}
		for j = 1; j <= numChars; j++ {
			if j == 1 || j <= numChars/2+1 || j == numChars {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
			} else if i == 1 || i == m["size"]/2+1 {
				fmt.Print(startColor + string(m["char"]) + colorFinish)
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
