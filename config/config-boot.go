package config

import (
	"github.com/bitwormhole/git-aliyun-oss/git2oss"
	"github.com/bitwormhole/starter-cli/cli"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/vlog"
)

type ConfigBoot struct {
	markup.Component `initMethod:"Init"`

	Context application.Context `inject:"context"`
	CLI     cli.ClientFactory   `inject:"#cli-client-factory"`
}

func (inst *ConfigBoot) Init() error {

	vlog.Warn("boot ... todo")

	// ctx := inst.Context
	// client := inst.CLI.CreateClient(ctx)
	// client.Execute("git-aliyun-oss", os.Args)

	return git2oss.Run()

	// return nil
}
