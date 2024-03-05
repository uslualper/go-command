package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}
