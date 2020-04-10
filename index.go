package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"github.com/inconshreveable/go-update"
	"github.com/urfave/cli/v2"
)

func mkFullname(uk string) string {
	osString := runtime.GOOS

	if osString == "linux" {
		return uk
	}

	panic("Not match any OS")
}

func getPanicProgram() string {
	return "exit 1"
}

func f() (err error) {
	updateInfo := getUpdateInfos()
	currentInfo := getCurrentInfos()

	dir, err := getCurrentDir()
	if err != nil {
		panic(err)
	}

	for uk, uv := range updateInfo {
		info := currentInfo[uk]

		if info == nil {
			fullname := mkFullname(uk)
			fmt.Printf("😁 Getting %s...\n", uk)

			fullPath := path.Join(dir, fullname)

			panicProgramBytes := []byte(getPanicProgram())
			err := ioutil.WriteFile(fullPath, panicProgramBytes, 0755)
			if err != nil {
				return err
			}

			reader, err := getBinary(uv.fullName)
			defer reader.Close()
			if err != nil {
				return err
			}

			var options update.Options

			options = update.Options{
				TargetPath: fullPath,
			}

			err = update.Apply(reader, options)
			if err != nil {
				return err
			}

			continue
		}

		version := info.version
		if version != uv.version {
			fmt.Printf("😺 Current %s is %s\n", uk, version)
			fmt.Printf("😼 Newest %s is %s\n", uk, uv.version)

			fmt.Printf("🔧 Updating %s\n", uk)

			reader, err := getBinary(uv.fullName)
			defer reader.Close()
			if err != nil {
				return err
			}

			var options update.Options

			if uk != "kan-update" {
				options = update.Options{
					TargetPath: info.fullName,
				}
			}

			err = update.Apply(reader, options)
			if err != nil {
				return err
			}

			fmt.Printf("✅ Update %s\n", uk)
		}
	}

	return
}

func index(c *cli.Context) (err error) {
	return f()
}
