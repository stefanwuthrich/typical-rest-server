package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-rest-server/internal/infra"
)

func init() {
	typapp.AppendCtor(
		&typapp.Constructor{
			Name: "",
			Fn: func() (*infra.App, error) {
				var cfg infra.App
				if err := envconfig.Process("APP", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
		&typapp.Constructor{
			Name: "",
			Fn: func() (*infra.Redis, error) {
				var cfg infra.Redis
				if err := envconfig.Process("REDIS", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
		&typapp.Constructor{
			Name: "",
			Fn: func() (*infra.Pg, error) {
				var cfg infra.Pg
				if err := envconfig.Process("PG", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
	)
}