// 引用样例
// @author MoGuQAQ
// @version 1.0.0

package main

import (
	"sample/core"
	"sample/global"
)

func main() {
	global.Log = core.InitLogger()
	global.Log.Infof("Info样例,int %d,string %s,float %f,char %c", 1, "1", 1.0, '1')
	global.Log.Warnf("Warn样例")
	global.Log.Errorf("Error样例")
	global.Log.Infoln("不带format的输出样例")
	test()
	global.Log.Fatalf("Fatal样例")
}

func test() {
	if global.Log == nil {
		global.Log = core.InitLogger()
	}
	global.Log.Infof("在其他函数内引用logger")
}
