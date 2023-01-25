package internal

import (
	"github.com/spf13/viper"
)

var useEtcd			= false
var useConfigFile	= false

type Config struct {
	SECRET string
}

func NewData() (Config, error) {
	salida := Config{}

	// Archivo de configuración
	if(useConfigFile) {
		viper.SetConfigType("env")
		viper.SetConfigName("app")
		viper.AddConfigPath("..")

		err := viper.ReadInConfig()
		if err != nil {
			salida.SECRET = "error"
			return salida, &errorGenerateConfig{" fallo en la lectura del archivo .env"}
		}
	} else {
		salida.SECRET = "prueba"
	}

	// Variables de entorno
	viper.AutomaticEnv()

	// Etcd
	if(useEtcd) {
		viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/tests.json")
		viper.SetConfigType("json")
		err := viper.ReadRemoteConfig()
		if err != nil {
			salida.SECRET = "error"
			return salida, &errorGenerateConfig{" fallo en la lectura con etcd"}
		}
	}

	err := viper.Unmarshal(&salida)
	if err != nil {
		return salida, &errorGenerateConfig{" error haciendo unmarshal "}
	}

	return salida, nil
}
