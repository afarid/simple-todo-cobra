/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			todoName := args[0]
			for index, todo := range todos.Todos {
				if todo.Name == todoName {
					// check if --descriptions is passed as flag or not
					if cmd.Flags().Lookup("description").Changed {
						description, _ := cmd.Flags().GetString("description")
						todos.Todos[index].Description = description
					}
					// check if --deadline is passed as flag or not
					if cmd.Flags().Lookup("deadline").Changed {
						deadline, _ := cmd.Flags().GetString("deadline")
						todos.Todos[index].Deadline = deadline

					}
				}

			}
			viper.Set("todos", todos.Todos)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("description", "", "Todo description")
	updateCmd.Flags().String("deadline", "", "Deadline for the todo")
}
