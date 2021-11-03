package main

import (
	"github.com/bitwormhole/git-aliyun-oss/git2oss"
)

func main() {
	// i := starter.InitApp()
	// i.Use(theModule())
	// i.Run()

	err := git2oss.Run()
	if err != nil {
		panic(err)
	}
}
