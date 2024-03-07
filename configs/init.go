package configs

import (
	"os"

	"github.com/pennywise-life/pennywise/configs/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgFile string

func InitConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			logrus.Errorf("Failed to get home directory: %v", err)
			os.Exit(1)
		}
		// check if config file exists
		if _, err := os.Stat(home + "/.pennywise/.env.local"); os.IsNotExist(err) {
			logrus.Info("Config file not found. Creating a new one...")
			// create config file
			// check if ~/.pennywise config file exists
			if _, err := os.Stat(home + "/.pennywise"); os.IsNotExist(err) {
				// create ~/.pennywise directory
				err := os.Mkdir(home+"/.pennywise", 0755)
				if err != nil {
					logrus.Errorf("Failed to create config directory: %v", err)
					os.Exit(1)
				}
			}
			// create config.yaml file in ~/.pennywise
			// check if ~/.pennywise/.env.local file exists
			if _, err := os.Stat(home + "/.pennywise/.env.local"); os.IsNotExist(err) {
				// create ~/.pennywise/.env.local file
				_, err = os.Create(home + "/.pennywise/.env.local")
				if err != nil {
					logrus.Errorf("Failed to create config file: %v", err)
					os.Exit(1)
				}
			}
		}
		envPath := home + "/.pennywise/.env.local"
		viper.AddConfigPath(home + "/.pennywise")
		viper.SetConfigType("env")
		viper.SetConfigFile(envPath)

		// set env variables
		viper.AutomaticEnv()

		// read config file
		err = viper.ReadInConfig()
		if err != nil {
			logrus.Errorf("Failed to read config file: %v", err)
			os.Exit(1)
		}

		err = viper.Unmarshal(&models.Env{})
		if err != nil {
			logrus.Errorf("Failed to unmarshal config file: %v", err)
			os.Exit(1)
		}

		// set config file
		cfgFile = viper.ConfigFileUsed()
	}
}

func GetConfig() *viper.Viper {
	return viper.GetViper()
}
