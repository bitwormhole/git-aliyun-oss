package main

// import (
// 	"embed"

// 	"github.com/bitwormhole/git-aliyun-oss/gen"
// 	"github.com/bitwormhole/starter"
// 	startercli "github.com/bitwormhole/starter-cli"
// 	"github.com/bitwormhole/starter/application"
// 	"github.com/bitwormhole/starter/collection"
// )

// const (
// 	theModuleName = "github.com/bitwormhole/git-aliyun-oss"
// 	theModuleVer  = "v0.0.1"
// 	theModuleRev  = 1
// )

// //go:embed src/main/resources
// var theMainRes embed.FS

// func theModule() application.Module {
// 	mb := application.ModuleBuilder{}
// 	mb.Name(theModuleName).Version(theModuleVer).Revision(theModuleRev)
// 	mb.Resources(collection.LoadEmbedResources(&theMainRes, "src/main/resources"))
// 	mb.OnMount(gen.ExportConfigGitOSS)

// 	mb.Dependency(starter.Module())
// 	mb.Dependency(startercli.Module())
// 	mb.Dependency(startercli.ModuleWithBasicCommands())

// 	return mb.Create()
// }
