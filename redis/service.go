// Copyright 2020 orivil.com. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found at https://mit-license.org.

package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/orivil/cfg"
	"github.com/orivil/service"
)

type Service struct {
	configService   *cfg.Service
	self            service.Provider
	configNamespace string
}

func (s *Service) New(ctn *service.Container) (value interface{}, err error) {
	var envs cfg.Env
	envs, err = s.configService.Get(ctn)
	if err != nil {
		return nil, err
	}
	env := &Env{}
	err = envs.UnmarshalSub(s.configNamespace, env)
	if err != nil {
		panic(err)
	}
	var client *redis.Client
	client, err = env.Init()
	if err != nil {
		return nil, err
	}
	ctn.OnClose(client.Close)
	return client, nil
}

func (s *Service) Get(ctn *service.Container) (*redis.Client, error) {
	c, err := ctn.Get(&s.self)
	if err != nil {
		return nil, err
	} else {
		return c.(*redis.Client), nil
	}
}

func NewService(configNamespace string, configService *cfg.Service) *Service {
	s := &Service{
		configService:   configService,
		configNamespace: configNamespace,
	}
	s.self = s
	return s
}
