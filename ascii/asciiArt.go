package ascii

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArrt(text []string, bannerName string) string {
	if bannerName != "shadow" && bannerName != "thinkertoy" && bannerName != "standard" && bannerName != "jacky" && bannerName != "ascii3D" {
		return ""
	}

	file, err := os.ReadFile(bannerName + ".txt")
	if err != nil {
		return ""
	}

	bannerContent := strings.ReplaceAll(string(file), "\r\n", "\n")
	banner := strings.Split(bannerContent, "\n")

	// text := strings.Split(text, "\\n")

	if len(text) > 0 && !hasNonEmptyLine(text) {
		text = text[1:]
	}
	res := ""
	for _, word := range text {
		if word == "" {
			fmt.Println()
		} else {
			res += printArt(word, banner)
		}
	}
	return res
}

func printArt(word string, banner []string) string {
	for _, r := range word {
		if r < 32 || r > 126 {
			return ""
		}
	}
	var str string
	for i := 0; i < 9; i++ {
		for _, r := range word {
			output := banner[9*(int(r)-32)+i]
			str += output
		}
		str = str + "\n"
	}
	return str
}

func hasNonEmptyLine(lines []string) bool {
	for _, line := range lines {
		if line != "" {
			return true
		}
	}
	return false
}
