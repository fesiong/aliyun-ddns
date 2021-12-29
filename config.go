package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func initPath() {
	sep := string(os.PathSeparator)
	ExecPath, _ = os.Getwd()
	pathArray := strings.Split(ExecPath, "/")
	if strings.Contains(ExecPath, "\\") {
		pathArray = strings.Split(ExecPath, "\\")
	}

	for i, v := range pathArray {
		if v == "ddns" {
			ExecPath = strings.Join(pathArray[:i+1], "/")
			break
		}
	}
	length := utf8.RuneCountInString(ExecPath)
	lastChar := ExecPath[length-1:]
	if lastChar != sep {
		ExecPath = ExecPath + sep

	}
	fmt.Println(ExecPath)
}

func initJSON() {
	rawConfig, err := ioutil.ReadFile(fmt.Sprintf("%sconfig.json", ExecPath))
	if err != nil {
		fmt.Println("Invalid Config: ", err.Error())
		os.Exit(-1)
	}

	if err := json.Unmarshal(rawConfig, &JsonData); err != nil {
		fmt.Println("Invalid Config: ", err.Error())
		os.Exit(-1)
	}

	log.Println(JsonData)
}

var ExecPath string
var JsonData aliConfig

func init() {
	//log.SetFlags( log.Ldate | log.Ltime | log.Llongfile)
	initPath()
	initJSON()
}