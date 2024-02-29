/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package login

import (
	"fmt"

	"github.com/spf13/cobra"
)

type loginCmd struct{}

func NewLoginCmd() *cobra.Command {
	l := &loginCmd{}

	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login to your Pennywise account",
		PreRun: func(_ *cobra.Command, _ []string) {
			// cmd.logger.Warn("WARNING: This command is in alpha version and may have some bugs and breaking changes in future releases.")
			yellow := "\033[33m"
			reset := "\033[0m"
			text := "WARNING: This command is in alpha version and may have some bugs and breaking changes in future releases."
			fmt.Printf("%s%s%s\n", yellow, text, reset)
		},
		RunE: l.run,
	}

	return loginCmd
}

func (l *loginCmd) run(cmd *cobra.Command, args []string) error {
	fmt.Println("login called")
	return nil
}
