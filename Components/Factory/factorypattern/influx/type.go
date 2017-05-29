package influx

import (
	"fmt"

	"github.com/wangzhezhe/BLG/Components/Factory/factorypattern/storagetype"
)

type Influx struct {
}

func (i Influx) List() {
	fmt.Println("list influx info")
}

func (i Influx) Create() (storagetype.StorageDriver, error) {
	fmt.Println("create influx instance")
	instance := &Influx{}
	return instance, nil
}
