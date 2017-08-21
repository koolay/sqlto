package cmd

import (
	"github.com/koolay/sqlto/config"
	"github.com/spf13/cobra"
)

// rootCmd root
var rootCmd = &cobra.Command{
	Use:   "sqlto",
	Short: "sqlto is a tool for exporting data to xls, json, csv files from a SQL.",
	Long:  `Support mysql, postegres, sqlite, mssql.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&config.Global.Output, "output", "o", "", "Output file eg. employees.xls")
	rootCmd.PersistentFlags().StringVarP(&config.Global.SQL, "sql", "l", "", "SQL. eg. select * from employee limit 20")
	rootCmd.PersistentFlags().BoolVarP(&config.Global.Pretty, "pretty", "p", false, "Pretty json output. It nly is affective for json.")
}

func initConfig() {

}
