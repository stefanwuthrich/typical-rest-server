package extension

import (
	"fmt"

	"github.com/typical-go/typical-rest-server/typical/task/database"
	"gopkg.in/urfave/cli.v1"
)

// DatabaseExtension provide standard command for database operation like create, drop, migrate, rollback, generate migration
type DatabaseExtension struct {
	ActionTrigger
	Extension
}

// Setup database extension
func (e *DatabaseExtension) Setup() error {
	return fmt.Errorf("not implemented")
}

//Command for database extension
func (e *DatabaseExtension) Command() cli.Command {
	return cli.Command{
		Name:      "database",
		ShortName: "db",
		Subcommands: []cli.Command{
			{Name: "create", Usage: "Create New Database", Action: e.Invoke(database.Create)},
			{Name: "drop", Usage: "Drop Database", Action: e.Invoke(database.Drop)},
			{Name: "migrate", Usage: "Migrate Database", Action: e.Invoke(database.Migrate)},
			{Name: "rollback", Usage: "Rollback Database", Action: e.Invoke(database.Rollback)},
		},
	}
}
