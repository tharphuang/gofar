package command

import (
	"fmt"
	"github.com/TharpHuang/gofar/tools/text"
)

var warningMigrate = text.TrimLeft(`
Gofar migrate: no arguments input.
Use "gofar migrate [arguments]"
The arguments are:
	create	migrate all database migration 
	rollback	rollback the databases migration `)

func Migrate(arg string) {
	if arg == "" {
		fmt.Println(warningMigrate)
		return
	}

	switch arg {
	case "create":
		migrateAll()
	case "rollback":
		rollBack()
	default:
		fmt.Println(warningMigrate)
	}

}

func migrateAll() {
	fmt.Println("all files update")

}

func rollBack() {
	fmt.Println("all files rollback")

}
