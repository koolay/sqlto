package cmd

import (
	"fmt"
	"os"
)

// Exec execute main command
func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(mysqlCmd)
}
