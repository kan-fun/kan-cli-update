package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func getOutputString(path string) (outputString string) {
	outputBytes, err := exec.Command(path, "-v").Output()
	if err != nil {
		panic(err)
	}

	outputString = strings.TrimSpace(string(outputBytes))

	return
}

func TestLinux(t *testing.T) {
	oldPath := "./test/linux/kan"
	newPath := "./kan"
	err := os.Rename(oldPath, newPath)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "kan-update version 0.0.0", getOutputString("./kan-update"))
	assert.Equal(t, "kan version 0.0.0", getOutputString("./kan"))

	_, err = exec.Command("./kan-update").Output()
	if err != nil {
		panic(err)
	}

	_, err = exec.Command("./kan-config", "init", "--access-key", "123", "--secret-key", "456").Output()
	if err != nil {
		panic(err)
	}

	newKanOptput := getOutputString("./kan")
	assert.Equal(t, true, strings.HasPrefix(newKanOptput, "kan version "))
	assert.NotEqual(t, "kan version 0.0.0", newKanOptput)

	newKanUpdateOptput := getOutputString("./kan-update")
	assert.Equal(t, true, strings.HasPrefix(newKanUpdateOptput, "kan-update version "))
	assert.NotEqual(t, "kan-update version 0.0.0", newKanUpdateOptput)
}
