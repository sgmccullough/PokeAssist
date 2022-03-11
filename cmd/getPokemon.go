/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/sgmccullough/pokeapi-go"
	"github.com/spf13/cobra"
)

// getPokemonCmd represents the getPokemon command
var getPokemonCmd = &cobra.Command{
	Use:   "getPokemon",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getPokemon called")
		l, err := pokeapi.Pokemon(strings.ToLower(args[0]))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(l)
	},
}

func init() {
	rootCmd.AddCommand(getPokemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPokemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPokemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
