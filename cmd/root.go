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
	"AsturDB/pkg/wolfgres"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "asturdb",
	Short: "AsturDB is a WolfgresTool to create a stress test to wolfgres_db",
	Long: `*********************
    AsturDB by Wolfgres - Postgres Enterprise
*********************
AsturDB is a WolfgresTool to run stress test to database in PostgreSQL.
The objetive is create a wolfgres_db (or your like name it) 
to size like you want, this project run many users simuntanius 
in the stress test. 

--------------------------------
www.wolfgres.com
Wolfgres - Postgres Enterprise
Live PostgreSQL - 2022
--------------------------------`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.AsturDB.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Verificar
	//viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		dir, file := path.Split(cfgFile)
		viper.AddConfigPath(dir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(file)

	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".AsturDB" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".AsturDB")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//log.Debug("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Fatal("Could not read configuration file")
		os.Exit(1)
	}

	//fmt.Println(viper.Get("log.log_format"))
	if viper.Get("log.log_format") == "text" {
		log.SetFormatter(&prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		})
	}

	exePath := wolfgres.GetExecPath()
	logPath := exePath + wolfgres.PathSeparator + filepath.Dir(fmt.Sprintf("%v", viper.Get("log.log_file")))
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		log.Error(err)
		log.Info("Create log directory in ", logPath)
		err := os.Mkdir(logPath, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if viper.Get("log.log_format") == "json" {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})
	}

	logFile := exePath + wolfgres.PathSeparator + fmt.Sprintf("%v", viper.Get("log.log_file"))
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)

	// Only log the warning severity or above.
	logLevel := wolfgres.GetLogLevel(fmt.Sprintf("%v", viper.Get("log.log_level")))
	log.SetLevel(logLevel)
	//log.SetLevel(log.InfoLevel)
	//fmt.Println(viper.Get("log.log_format"))
}
