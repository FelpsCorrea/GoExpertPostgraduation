/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/FelpsCorrea/GoExpertPostgratuation/Cobra-CLI/internal/database"
	"github.com/spf13/cobra"
)

func newCreateCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a category",
		RunE:  runCreate(categoryDb),
	}
}

// // createCmd represents the create command
// var createCmd = &cobra.Command{
// 	Use:   "create",
// 	Short: "A brief description of your command",
// 	Long:  `A longer description that spans multiple lines and likely contains examples`,
// 	RunE:  runCreate(GetCategoryDB(GetDb())),
// }

func runCreate(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDb.Create(name, description)

		if err != nil {
			return err
		}

		fmt.Println("Created!")
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(GetCategoryDB(GetDb()))

	categoryCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Category name")
	createCmd.Flags().StringP("description", "d", "", "Category description")
	createCmd.MarkFlagsRequiredTogether("name", "description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
