// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/koolay/sqlto/config"
	"github.com/spf13/cobra"
)

// mysqlCmd represents the mysql command
var pgCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Export records from postgres",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execExport("postgres")
	},
}

func init() {
	pgCmd.PersistentFlags().StringVarP(&config.Global.Source, "source", "s", "", "format: postgres://pqgotest:password@localhost/pqgotest?sslmode=disable")
}
