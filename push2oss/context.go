package push2oss

import (
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
)

type Context struct {
	ScanSer   ScanService
	GitSer    GitService
	OSSSer    OSSService
	UploadSer UploadService

	Pool lang.ReleasePool

	PWD           fs.Path
	GitConfigFile fs.Path
	GitRepoDir    fs.Path
	GitWorkingDir fs.Path

	OSSURL string

	BucketParams   OSSBucketParams
	GitConfigProps collection.Properties
}
