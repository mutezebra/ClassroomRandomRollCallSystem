package cache

import (
	. "github.com/bytedance/mockey"
	"github.com/mutezebra/ClassroomRandomRollCallSystem/config"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func testInit(t *testing.T) {
	host, port, password, network, db := "127.0.0.1", "6379", "123456", "tcp", 0
	redis := config.Redis{
		Host:     host,
		Port:     port,
		Database: db,
		Network:  network,
		Password: password,
	}
	conf := &config.Config{Redis: &redis}
	MockValue(&config.Conf).To(conf)
	InitCache()
}

func TestInitCache(t *testing.T) {
	testInit(t)
	Convey("Test Redis Init", t, func() {
		So(RedisClient, ShouldNotBeNil)
	})
}
