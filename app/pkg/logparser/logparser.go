package logparser

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Parse() {
	lines, requestA := 0, 0
	f, err := os.Open("request.log")
	if err != nil {
		fmt.Print("There has been an error!: ", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines++
		// filter request a
		line := scanner.Bytes()
		if len(line) <= 30 || line[30] != 'A' {
			continue
		}
		if !bytes.Equal(line[22:], []byte("REQUEST-A")) {
			continue
		}
		requestA++
		request := string(line)

		// handle request a
		fmt.Println(request)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines, requestA)
}
