/*
Copyright © 2021 Wolfgres info@wolfgres.com

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
	"AsturDB/internal/AsturDB"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	nameStressTest string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a stress test",
	Long:  `Start a stress test for a database PostgreSQL`,
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//startCmd.PersistentFlags().String(&host, "host", "h", "Host server PostgreSQL")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.Flags().StringVarP(&nameStressTest, "name", "n", "wolfgres_db", "Name strest test")
}

func start() {
	log.Info("************************************************")
	log.Info("*** AsturDB by Wolfgres - Postgres Enterprise ****")
	log.Info("************************************************")
	log.Info("")
	log.Info("Start de Stress test Name Stress Test: ", nameStressTest)

	log.Debug("Database name set: ", nameStressTest)
	AsturDB.RunStressTest(nameStressTest)
}
