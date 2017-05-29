package redis

import (
	"fmt"

	"github.com/wangzhezhe/BLG/Components/Factory/factorypattern/storagetype"
)

type Redis struct {
}

func (r Redis) List() {
	fmt.Println("list redis info")
}

func (r Redis) Create() (storagetype.StorageDriver, error) {
	fmt.Println("create redis instance")
	instance := &Redis{}
	return instance, nil
}
