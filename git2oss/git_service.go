package git2oss

import (
	"errors"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
)

type GitService interface {
	Open() error
}

type GitServiceImpl struct {
	Context *Context
}

func (inst *GitServiceImpl) Open() error {
	dotgit, err := inst.findDotGit()
	if err != nil {
		return err
	}
	inst.Context.GitConfigFile = dotgit.GetChild("config")
	inst.Context.GitRepoDir = dotgit
	inst.Context.GitWorkingDir = dotgit.Parent()
	return inst.loadConfigProperties()
}

func (inst *GitServiceImpl) loadConfigProperties() error {

	file := inst.Context.GitConfigFile

	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	inst.Context.GitConfigProps = props
	return nil
}

func (inst *GitServiceImpl) findDotGit() (fs.Path, error) {
	pwd := inst.Context.PWD
	for p := pwd; p != nil; p = p.Parent() {
		dotgit := p.GetChild(".git")
		if dotgit.IsDir() {
			return dotgit, nil
		}
	}
	return nil, errors.New("cannot find '.git' in path of " + pwd.Path())
}
