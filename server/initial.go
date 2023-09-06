package server

import (
	"fmt"
	"gc/database/server/variables"
	"log"
	"os"

	"github.com/spf13/viper"
)

func init() {
	if _, err := os.Stat("./settings.yaml"); os.IsNotExist(err) {
		err := os.WriteFile("./settings.yaml", []byte(variables.Setting), 0644)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}

	//Inicializamos la lectura del archivo de configuracion
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("No es posible leer el archivo setting %v\n", err)
		os.Exit(1)
	}
}
