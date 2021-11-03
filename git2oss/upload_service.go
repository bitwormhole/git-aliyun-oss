package git2oss

import (
	"github.com/bitwormhole/starter/io/fs"
)

type UploadService interface {
	Upload() error
}

////////////////////////////////////////////////////////////////////////////////

type UploadServiceImpl struct {
	Context *Context
}

func (inst *UploadServiceImpl) Upload() error {

	ctx := inst.Context

	err := ctx.GitSer.Open()
	if err != nil {
		return err
	}

	err = ctx.OSSSer.Open()
	if err != nil {
		return err
	}

	err = ctx.ScanSer.Scan(inst.handleFile)
	if err != nil {
		return err
	}

	return nil
}

func (inst *UploadServiceImpl) handleFile(file fs.Path, spath string) error {
	return inst.Context.OSSSer.UploadFile(spath, file)
}
