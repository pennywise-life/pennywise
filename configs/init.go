package configs

import (
	"os"

	"github.com/pennywise-life/pennywise/configs/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgFile string

// InitConfig initializes Viper config.
// If a config file is found, read it in.
func InitConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			logrus.WithError(err).Errorf("Failed to get home directory")
			os.Exit(1)
		}

		// Search config in ~/.pennywise directory with name ".env.local".
		viper.AddConfigPath(home + "/.pennywise")
		viper.SetConfigType("env")
		viper.SetConfigName(".env.local")

		// Create config file if not found.
		if _, err := os.Stat(home + "/.pennywise/.env.local"); os.IsNotExist(err) {
			logrus.Info("Config file not found. Creating a new one...")
			// Create config directory if not exists.
			if _, err := os.Stat(home + "/.pennywise"); os.IsNotExist(err) {
				err = os.Mkdir(home+"/.pennywise", 0755)
				if err != nil {
					logrus.WithError(err).Errorf("Failed to create config directory")
					os.Exit(1)
				}
			}
			// Create config file.
			f, err := os.Create(home + "/.pennywise/.env.local")
			if err != nil {
				logrus.WithError(err).Errorf("Failed to create config file")
				os.Exit(1)
			}
			f.Close()
		}
	}

	// Set environment variables.
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Errorf("Failed to read config file")
		os.Exit(1)
	}

	// Unmarshal the configuration into models.Env.
	if err := viper.Unmarshal(&models.Env{}); err != nil {
		logrus.WithError(err).Errorf("Failed to unmarshal config file")
		os.Exit(1)
	}

	// Set config file used by Viper.
	cfgFile = viper.ConfigFileUsed()
}

func GetConfig() *viper.Viper {
	return viper.GetViper()
}
