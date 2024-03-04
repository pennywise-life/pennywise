/*
Copyright © 2024 Azharuddin Mohammed <azharuddinmohammed998@gmail.com>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/pennywise-life/pennywise/cmd/login"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

const PENNYWISE = `
░▒▓███████▓▒░░▒▓████████▓▒░▒▓███████▓▒░░▒▓███████▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓███████▓▒░▒▓████████▓▒░ 
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        
░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        
░▒▓███████▓▒░░▒▓██████▓▒░ ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓██████▓▒░░▒▓██████▓▒░   
░▒▓█▓▒░      ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
░▒▓█▓▒░      ░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
░▒▓█▓▒░      ░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░    ░▒▓█████████████▓▒░░▒▓█▓▒░▒▓███████▓▒░░▒▓████████▓▒░ 
`

func newRootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Version: "v0.0.1",
		Use:     "pennywise",
		Aliases: []string{"pw"},
		Short:   "Pennywise is a CLI tool for managing your personal finance.",
		Long:    getPennywiseDescription(),
	}
	return rootCmd
}

func buildRootCommand() *cobra.Command {
	rootCmd := newRootCommand()
	rootCmd.AddCommand(login.NewLoginCmd())
	return rootCmd
}

func Execute() {
	rootCmd := buildRootCommand()
	handleExecutionError(rootCmd.Execute())
}

func init() {
	rootCmd := buildRootCommand()
	cobra.OnInitialize(initConfig)
	setupRootCmdFlags(rootCmd)
}

func getPennywiseDescription() string {
	return lipgloss.NewStyle().
		Padding(0, 2).
		Foreground(lipgloss.Color("#FFB0B0")).
		Background(lipgloss.Color("#0C2D57")).
		SetString(PENNYWISE).
		Render("\n") + "\n" +
		lipgloss.NewStyle().
			Width(123).
			Padding(1, 2).
			Align(lipgloss.Left).
			Bold(true).
			Background(lipgloss.Color("#0C2D57")).
			Foreground(lipgloss.Color("#EFECEC")).
			Render("Pennywise is a CLI tool for managing your personal finance. It helps you to track your expenses and income. Also you can create reports, analyze your financial data and much more using Pennywise Dashboard.")
}

func handleExecutionError(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func setupRootCmdFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pennywise.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".pennywise")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
