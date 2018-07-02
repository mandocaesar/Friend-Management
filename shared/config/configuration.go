package config

import (
	"github.com/golang/glog"
	"github.com/spf13/viper"
)

//Configuration : Struct to hold configuration object
type Configuration struct {
	Database DatabaseConfiguration
	Server   ServerConfiguration
}

//New : Create new Configuration object/instance
func New() (*Configuration, error) {
	viper.SetConfigName("default")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		glog.Fatal("Failed to load configuration file", err)
		return nil, err
	}

	viper.SetConfigName(".env")
	if err := viper.MergeInConfig(); err != nil {
		glog.Warningf("Failed to load custom configuration : %s", err)
	}

	cfg := new(Configuration)
	if err := viper.Unmarshal(cfg); err != nil {
		glog.Fatalf("Failed to deserialize config struct: %s", err)
		return nil, err
	}

	return cfg, nil
}
