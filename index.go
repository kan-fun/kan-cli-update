package main

import (
	"fmt"
	"os"

	"github.com/inconshreveable/go-update"
	"github.com/urfave/cli/v2"
)

func f() (err error) {
	updateInfo := getUpdateInfos()
	currentInfo := getCurrentInfos()

	for uk, uv := range updateInfo {
		info := currentInfo[uk]

		if info == nil {
			println("Not Existed ", uk)
			continue
		}

		version := info.version
		if version != uv.version {
			fmt.Printf("🔧 Current %s is %s\n", uk, version)
			fmt.Printf("🔧 Newest %s is %s\n", uk, uv.version)

			fmt.Printf("🔧 Updating %s\n", uk)

			reader, err := getBinary(uv.fullName)
			defer reader.Close()
			if err != nil {
				panic(err)
			}

			var options update.Options

			if uk != "kan-update" {
				options = update.Options{
					TargetPath: info.fullName,
				}
			}

			err = update.Apply(reader, options)
			if err != nil {
				panic(err)
			}

			if err := os.Chmod(info.fullName, 755); err != nil {
				panic(err)
			}

			fmt.Printf("✅ Update %s\n", uk)
		}
	}

	return
}

func index(c *cli.Context) (err error) {
	return f()
}
