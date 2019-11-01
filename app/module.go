package app

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typiobj"
	"github.com/typical-go/typical-rest-server/app/config"
	"github.com/urfave/cli"
)

// Module of application
func Module() interface{} {
	return applicationModule{
		Configuration: typiobj.Configuration{
			Prefix: "APP",
			Spec:   &config.Config{},
		},
	}
}

type applicationModule struct {
	typiobj.Configuration
}

func (m applicationModule) CommandLine() cli.Command {
	return cli.Command{
		Name: "route", Description: "Print available API Routes", Action: Route,
	}
}

func (m applicationModule) Prepare() []interface{} {
	return []interface{}{
		Routes,
		Middlewares,
	}
}

func (m applicationModule) Run() interface{} {
	return Start
}

func (m applicationModule) Provide() []interface{} {
	return []interface{}{
		m.loadConfig,
	}
}

func (m applicationModule) loadConfig() (cfg *config.Config, err error) {
	err = m.Configuration.Load()
	cfg = m.Configuration.Spec.(*config.Config)
	return
}