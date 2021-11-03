package git2oss

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/util"
)

const (
	ConfigOSSBucket        = "oss.{{id}}.bucket"
	ConfigOSSEndpoint      = "oss.{{id}}.endpoint"
	ConfigOSSAccessKeyID   = "oss.{{id}}.accesskeyid"
	ConfigOSSAccessKeyFile = "oss.{{id}}.accesskeyfile" // AccessKeySecret 的内容保存在keyfile中
)

type OSSBucketParams struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
}

type OSSService interface {
	Open() error
	UploadFile(objectName string, file fs.Path) error
}

////////////////////////////////////////////////////////////////////////////////

type OSSServiceImpl struct {
	Context *Context
	client  *oss.Client
	bucket  *oss.Bucket
}

func (inst *OSSServiceImpl) Open() error {

	err := inst.loadOSSConfig()
	if err != nil {
		return err
	}

	err = inst.openBucket()
	if err != nil {
		return err
	}

	return nil
}

func (inst *OSSServiceImpl) loadOSSConfig() error {

	id := "default"
	errs := lang.NewErrorCollector()
	ctx := inst.Context

	// read config props
	keyFilePath := inst.getConfigValue(id, ConfigOSSAccessKeyFile, errs)
	ctx.BucketParams.Bucket = inst.getConfigValue(id, ConfigOSSBucket, errs)
	ctx.BucketParams.Endpoint = inst.getConfigValue(id, ConfigOSSEndpoint, errs)
	ctx.BucketParams.AccessKeyID = inst.getConfigValue(id, ConfigOSSAccessKeyID, errs)
	err := errs.LastError()
	if err != nil {
		return err
	}

	// load accesskeySecret
	keyFile := fs.Default().GetPath(keyFilePath)
	keyData, err := keyFile.GetIO().ReadText(nil)
	if err != nil {
		return err
	}
	ctx.BucketParams.AccessKeySecret = strings.TrimSpace(keyData)
	return nil
}

func (inst *OSSServiceImpl) getConfigValue(id string, field string, ec lang.ErrorCollector) string {
	const token = "{{id}}"
	key := strings.ReplaceAll(field, token, id)
	props := inst.Context.GitConfigProps
	value, err := props.GetPropertyRequired(key)
	if err != nil {
		ec.Append(err)
		return ""
	}
	return value
}

func (inst *OSSServiceImpl) openBucket() error {

	p := inst.Context.BucketParams

	const shortKeyLen = 8
	shortKeyID := p.AccessKeyID
	if len(shortKeyID) > shortKeyLen {
		shortKeyID = "***" + shortKeyID[len(shortKeyID)-shortKeyLen:]
	}

	fmt.Println("oss.endpoint    = ", p.Endpoint)
	fmt.Println("oss.bucket      = ", p.Bucket)
	fmt.Println("oss.accessKeyID = ", shortKeyID)

	client, err := oss.New(p.Endpoint, p.AccessKeyID, p.AccessKeySecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(p.Bucket)
	if err != nil {
		return err
	}

	inst.client = client
	inst.bucket = bucket
	return nil
}

func (inst *OSSServiceImpl) UploadFile(objectName string, file fs.Path) error {

	logger := &strings.Builder{}
	logger.WriteString("上传文件 ")
	logger.WriteString(objectName)
	logger.WriteString(" ......")
	for logger.Len() < 80 {
		logger.WriteRune('.')
	}
	defer func() {
		fmt.Println(logger.String())
	}()

	sum1, err := inst.fetchObjectSum(objectName)
	if err == nil && sum1 != nil {
		sum2, err := inst.computeMD5sum(file)
		if err != nil {
			return err
		}
		if bytes.Compare(sum1, sum2) == 0 {
			logger.WriteString("... 文件已存在，跳过!")
			return nil
		}
	}

	input, err := file.GetIO().OpenReader(nil)
	if err != nil {
		return err
	}
	defer input.Close()

	err = inst.bucket.PutObject(objectName, input)
	if err != nil {
		return err
	}

	logger.WriteString("... [OK]")
	return nil
}

func (inst *OSSServiceImpl) computeMD5sum(file fs.Path) ([]byte, error) {
	hash := md5.New()
	input, err := file.GetIO().OpenReader(nil)
	if err != nil {
		return nil, err
	}
	defer input.Close()
	err = util.PumpStream(input, hash, nil)
	if err != nil {
		return nil, err
	}
	sum := hash.Sum([]byte{})
	return sum[:], nil
}

func (inst *OSSServiceImpl) fetchObjectSum(name string) ([]byte, error) {
	props, err := inst.bucket.GetObjectDetailedMeta(name)
	if err != nil {
		return nil, err
	}
	key := "Content-Md5"
	value := props.Get(key)
	if value == "" {
		return nil, errors.New("no object.meta: " + name + "#" + key)
	}
	return base64.StdEncoding.DecodeString(value)
}
