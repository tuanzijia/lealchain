package app

import (
	"fmt"
	"lealchain/utils/log"
	"lealchain/utils/redis"
	"sync"

	goRedis "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

var (
	app  *App
	once sync.Once
)

type handle func()

type App struct {
	redis       redis.Client
	logicHandle map[string]handle
}

func newChainApp() *App {
	log.InstanceLogObject("log", "", zapcore.Level(viper.GetInt(FlagNameLogLevel)))

	return &App{redis: redis.InstanceRedisClint(&goRedis.UniversalOptions{
		Addrs:    viper.GetStringSlice(FlagNameRedisAddress),
		DB:       viper.GetInt(FlagNameRedisDB),
		Password: viper.GetString(FlagNameRedisPassword),
	})}
}

func InstanceApp() *App {
	once.Do(func() {
		app = newChainApp()
	})

	return app
}

func (a *App) RegisterLogicHandler(logicName string, h handle) {
	if _, exists := a.logicHandle[logicName]; exists {
		panic(fmt.Errorf("逻辑处理器已经存在:%s", logicName))
	}
	a.logicHandle[logicName] = h
}
