// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package redis_test

import (
	"context"
	"fmt"
	"github.com/orivil/cfg"
	"github.com/orivil/memory/redis"
	"github.com/orivil/service"
	"time"
)

var config = `
# redis配置
[redis]
# 地址
addr = ""
# 密码
password = ""
# 数据库
db = 1
`

func ExampleNewService() {
	cfgService := cfg.NewService(cfg.NewMemoryStorageService(config))
	redisService := redis.NewService("redis", cfgService)
	container := service.NewContainer()
	defer container.Close() // The Close function will close redis.Client
	client, err := redisService.Get(container)
	if err != nil {
		panic(err)
	}
	client.Set(context.Background(), "key", "value", time.Second*2)
	var v1 string
	v1, err = client.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(v1) // value
}
