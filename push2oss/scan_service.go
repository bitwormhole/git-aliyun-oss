package push2oss

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter/io/fs"
)

type ScanHandlerFn func(file fs.Path, spath string) error

type ScanService interface {
	Scan(h ScanHandlerFn) error
}

type ScanServiceImpl struct {
	Context *Context
}

func (inst *ScanServiceImpl) Scan(h ScanHandlerFn) error {
	params := &inst.Context.BucketParams
	dir := inst.Context.GitWorkingDir
	if params.LocalRoot != "" {
		dir = dir.GetHref(params.LocalRoot)
	}
	return inst.scanDir(dir, params.RemoteRoot, 0, h)
}

func (inst *ScanServiceImpl) scanDir(dir fs.Path, spath string, depth int, h ScanHandlerFn) error {

	if !dir.IsDir() {
		return errors.New("the path is not a directory, path=" + dir.Path())
	}

	if depth > 99 {
		return errors.New("the path is too deep, path=" + dir.Path())
	}

	if dir.Name() == ".git" {
		return nil // skip
	}

	if len(spath) > 0 {
		if !strings.HasSuffix(spath, "/") {
			spath = spath + "/"
		}
	}

	list := dir.ListItems()

	for _, item := range list {
		path2 := spath + item.Name()
		if item.IsDir() {
			err := inst.scanDir(item, path2, depth+1, h)
			if err != nil {
				return err
			}
		} else if item.IsFile() {
			err := h(item, path2)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
