package git2oss

import (
	"errors"

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
	dir := inst.Context.GitWorkingDir
	return inst.scanDir(dir, "", 0, h)
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

	list := dir.ListItems()

	for _, item := range list {
		sep := ""
		if depth > 0 {
			sep = "/"
		}
		path2 := spath + sep + item.Name()
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
