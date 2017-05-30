package factorypattern

import (
	"testing"

	"github.com/wangzhezhe/BLG/Components/Factory/factorypattern/influx"
	"github.com/wangzhezhe/BLG/Components/Factory/factorypattern/mysql"
	"github.com/wangzhezhe/BLG/Components/Factory/factorypattern/redis"
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
			return
		}
		sdriver.List()
	}

}
