package logparser

import (
	"fmt"
	"io/ioutil"
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
	cmd := exec.Command("tail", "-n", lines, fullPath)
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

func convertAddressToString(i []interface{}) string {
	var s string
	for index, v := range i {
		if index == len(i)-1 {
			s += fmt.Sprintf("%v", v)
		} else {
			s += fmt.Sprintf("%v.", v)
		}
	}
	return s
}
