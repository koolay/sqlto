package cmd

import (
	"fmt"
	"os"

	"github.com/koolay/sqlto/config"
	"github.com/koolay/sqlto/db"
	"github.com/koolay/sqlto/output"
)

func execExport(driver string) {
	da := db.NewDB(driver)
	rows, err := da.Query(&config.Global)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if len(rows) == 0 {
		fmt.Println("no data")
		os.Exit(0)
	}
	if err := output.Export(&config.Global, rows); err != nil {
		fmt.Println(err)
	}
}

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
	rootCmd.AddCommand(pgCmd)
}
