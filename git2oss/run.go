package git2oss

import (
	"fmt"
	"os"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
)

func Run(mod application.Module) error {

	fmt.Println("Git >>>>> OSS (" + mod.GetVersion() + ")")

	ctx := &Context{}

	pool := lang.CreateReleasePool()
	ctx.Pool = pool
	defer pool.Release()

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	ctx.PWD = fs.Default().GetPath(pwd)

	ctx.GitSer = &GitServiceImpl{Context: ctx}
	ctx.OSSSer = &OSSServiceImpl{Context: ctx}
	ctx.ScanSer = &ScanServiceImpl{Context: ctx}
	ctx.UploadSer = &UploadServiceImpl{Context: ctx}

	err = ctx.UploadSer.Upload()
	if err != nil {
		return err
	}

	fmt.Println("已完成。")
	return nil
}
