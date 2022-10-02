/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			todoName := args[0]
			description, _ := cmd.Flags().GetString("description")
			deadline, _ := cmd.Flags().GetString("deadline")
			todo := Todo{
				Name:        todoName,
				Description: description,
				Deadline:    deadline,
			}

			todos.Todos = append(todos.Todos, todo)
			viper.Set("todos", todos.Todos)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			fmt.Println("Enter just one todo ")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	today := time.Now()
	tomorrow := today.AddDate(0, 0, 1)
	addCmd.Flags().String("description", "", "Todo description")
	addCmd.Flags().String("deadline", tomorrow.String(), "Deadline for the todo")
}
