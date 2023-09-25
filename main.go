package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := godotenv.Load(".env.local")

	if err != nil {
		fmt.Println("Gagal load file env")
	}

	fmt.Println("Host from godotenv : ", os.Getenv("DB_HOST"))
	fmt.Println("Port from godotenv : ", os.Getenv("DB_PORT"))

	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")

	err = viper.ReadInConfig()

	if err != nil {
		fmt.Println("Gagal load file env dari viper")
	}

	fmt.Println("Host from viper : ", viper.GetString("DB_HOST"))
	fmt.Println("Port from viper : ", viper.GetInt("DB_PORT"))
}
