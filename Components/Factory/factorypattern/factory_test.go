package factorypattern

import (
	"testing"

	"github.com/BLG/Components/Factory/factorypattern/influx"
	"github.com/BLG/Components/Factory/factorypattern/mysql"
	"github.com/BLG/Components/Factory/factorypattern/redis"
)

func TestRegisterandCreate(t *testing.T) {

	Register("mysql", &mysql.Mysql{})
	Register("influx", &influx.Influx{})
	Register("redis", &redis.Redis{})

	strlist := []string{"mysql", "redis", "influx", "others"}

	for _, item := range strlist {
		sdriver, err := Create(item)
		if err != nil {
			t.Error(err)
		}
		sdriver.List()
	}

}
