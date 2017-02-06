package factorypattern

import (
	"errors"
	"fmt"

	"github.com/BLG/Components/Factory/factorypattern/storagetype"
)

type StorageDriverFactory interface {
	//other storage functions...
	Create() (storagetype.StorageDriver, error)
}

var driverFactories = make(map[string]StorageDriverFactory)

func Create(name string) (storagetype.StorageDriver, error) {
	driverFactory, ok := driverFactories[name]
	if !ok {
		return nil, errors.New("invalid parameter: " + name)
	}
	return driverFactory.Create()
}

func Register(name string, factory StorageDriverFactory) {
	if factory == nil {
		panic("Must not provide nil StorageDriverFactory")
	}
	_, registered := driverFactories[name]
	if registered {
		panic(fmt.Sprintf("StorageDriverFactory named %s already registered", name))
	}

	driverFactories[name] = factory
}
