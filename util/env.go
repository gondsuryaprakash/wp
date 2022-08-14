package util

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetConfigValue(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in reading config File", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Not present")
	}
	return value
}
