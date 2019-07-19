package typictx

import (
	"go.uber.org/dig"
	"gopkg.in/urfave/cli.v1"
)

// Module of typical-go application
type Module struct {
	Name         string
	ShortName    string
	Usage        string
	ConfigPrefix string
	Config       interface{}

	Commands     []Command
	Constructors []interface{}

	OpenFunc  interface{}
	CloseFunc interface{}

	SideEffects        []string
	AppSideEffects     []string
	DevToolSideEffects []string
}

// Inject dependencies for the module
func (m *Module) Inject(container *dig.Container) {
	for _, constructor := range m.Constructors {
		container.Provide(constructor)
	}
	container.Provide(m.OpenFunc)
	return
}

// Invoke the function for CLI command
func (m *Module) Invoke(invokeFunc interface{}) interface{} {
	return func(ctx *cli.Context) error {
		container := dig.New()
		container.Provide(ctx.Args) // NOTE: inject cli arguments
		m.Inject(container)

		return container.Invoke(invokeFunc)
	}
}
