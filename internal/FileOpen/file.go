package FileOpen

import (
	"bufio"
	"os"
)

func OpenFile(patch string) []string {
	file, err := os.Open(patch)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
