package mysql

import (
	"fmt"

	"github.com/wangzhezhe/BLG/Components/Factory/factorypattern/storagetype"
)

type Mysql struct {
}

func (m Mysql) List() {
	fmt.Println("list mysql info")
}

func (m Mysql) Create() (storagetype.StorageDriver, error) {
	fmt.Println("create mysql instance")
	instance := &Mysql{}
	return instance, nil
}
