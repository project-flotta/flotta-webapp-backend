package logparser

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
)

const LogDir = "./tmp/"

func ReadLogFileRaw(dirPath, lines string) (string, error) {
	// list files in directory to get the latest file
	file, err := getLatestModifiedFile(LogDir + dirPath)
	if err != nil {
		return "", fmt.Errorf("error getting latest modified file: %v", err.Error())
	}

	// full path to log file
	fullPath := LogDir + dirPath + "/" + file

	// read number of line n from the end of log file
	cmd := exec.Command("tail", "-S", "-n", lines, fullPath)
	res, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing tail command: %v", err.Error())
	}
	return string(res), nil
}

func getLatestModifiedFile(dir string) (string, error) {
	// get files in directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	// sort files by modified time
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime().Before(files[j].ModTime())
	})

	// return latest file
	// if file is dir return latest file in dir (recursive)
	if files[0].IsDir() {
		return getLatestModifiedFile(dir + "/" + files[0].Name())
	}
	return files[0].Name(), nil
}

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
