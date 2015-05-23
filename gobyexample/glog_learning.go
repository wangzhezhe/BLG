//具体参见 github.com/golang/glog 具体使用golang来输出相关的日志信息 还有问题
package main

import (
	"github.com/golang/glog"
)

func main() {

	glog.Info("Prepare to repel boarders")

	//glog.Fatalf("Initialization failed: %s", err)

	if glog.V(2) {
		glog.Info("Starting transaction...")
	}

	//glog.V(2).Infoln("Processed", nItems, "elements")

}
