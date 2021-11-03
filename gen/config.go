package gen

import "github.com/bitwormhole/starter/application"

func ExportConfigGitOSS(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
