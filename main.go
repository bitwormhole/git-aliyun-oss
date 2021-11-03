package main

import (
	"github.com/bitwormhole/git-push2oss/push2oss"
	"github.com/bitwormhole/starter/application"
)

const (
	myName     = "github.com/bitwormhole/git-push2oss"
	myVersion  = "v1.0.0"
	myRevision = 1
)

func main() {

	mb := application.ModuleBuilder{}
	mb.Name(myName).Version(myVersion).Revision(myRevision)
	mod := mb.Create()

	err := push2oss.Run(mod)
	if err != nil {
		panic(err)
	}
}
