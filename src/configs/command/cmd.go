package command

import ("github.com/spf13/cobra"
		"carRent/src/configs/serve"
		database "carRent/src/database")

var initCommand = cobra.Command{
	Use: "dBackend Golang",
	Short: "api with Golang",
}

func init() {
	initCommand.AddCommand(serve.ServeCmd)
	initCommand.AddCommand(database.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}