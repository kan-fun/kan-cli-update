package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/kardianos/osext"
)

type infoStruct struct {
	fullName string
	version  string
}

func getCurrentDir() (dir string, err error) {
	dir, err = osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}

	return
}

func getInfoFromExec(filename string) (key string, info *infoStruct) {
	dir, err := getCurrentDir()
	if err != nil {
		panic(err)
	}

	outputBytes, err := exec.Command(path.Join(dir, filename), "-v").Output()
	if err != nil {
		panic(err)
	}

	outputString := string(outputBytes)

	t := strings.Split(outputString, " ")

	info = &infoStruct{
		filename,
		t[2][:len(t[2])-1],
	}

	key = t[0]
	return
}

func getUpdateInfos() (result map[string]*infoStruct) {
	osString := runtime.GOOS

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.kan-fun.com/bin", nil)
	if err != nil {
		panic(err)
	}

	q := req.URL.Query()
	q.Add("platform", osString)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic("Network Error")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	respString := string(respBody)

	filenames := strings.Split(respString, "\n")

	result = make(map[string]*infoStruct)
	for _, filename := range filenames {
		t := strings.Split(filename, "_")
		result[t[0]] = &infoStruct{
			filename,
			t[1],
		}
	}

	return
}

func getCurrentInfos() (result map[string]*infoStruct) {
	dir, err := getCurrentDir()
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	result = make(map[string]*infoStruct)
	for _, object := range files {
		if !object.IsDir() {
			filename := object.Name()
			if strings.HasPrefix(filename, "kan") {
				key, info := getInfoFromExec(filename)
				result[key] = info
			}
		}
	}

	return
}

func getBinary(fullName string) (reader io.ReadCloser, err error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://bin.kan-fun.com/%s/%s", runtime.GOOS, fullName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		panic("Network Error")
	}

	reader = resp.Body

	return
}
