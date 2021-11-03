// (todo:gen2.template)
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package gen

import (
	config0x003853 "github.com/bitwormhole/git-aliyun-oss/config"
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
)

func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()

	// component: com0-config0x003853.ConfigBoot
	cominfobuilder.Next()
	cominfobuilder.ID("com0-config0x003853.ConfigBoot").Class("").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComConfigBoot{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComConfigBoot : the factory of component: com0-config0x003853.ConfigBoot
type comFactory4pComConfigBoot struct {
	mPrototype *config0x003853.ConfigBoot
}

func (inst *comFactory4pComConfigBoot) init() application.ComponentFactory {

	inst.mPrototype = inst.newObject()
	return inst
}

func (inst *comFactory4pComConfigBoot) newObject() *config0x003853.ConfigBoot {
	return &config0x003853.ConfigBoot{}
}

func (inst *comFactory4pComConfigBoot) castObject(instance application.ComponentInstance) *config0x003853.ConfigBoot {
	return instance.Get().(*config0x003853.ConfigBoot)
}

func (inst *comFactory4pComConfigBoot) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst *comFactory4pComConfigBoot) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst *comFactory4pComConfigBoot) AfterService() application.ComponentAfterService {
	return inst
}

func (inst *comFactory4pComConfigBoot) Init(instance application.ComponentInstance) error {
	return inst.castObject(instance).Init()
}

func (inst *comFactory4pComConfigBoot) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst *comFactory4pComConfigBoot) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	return nil
}
