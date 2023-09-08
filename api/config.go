package api

import (
	"fmt"

	"github.com/spf13/viper"
)

type Api struct {
	IP         string
	Production bool
	CertFile   string
	KeyFile    string
	Rest       RESTApi
	GRPC       GRPCApi
}

type RESTApi struct {
	Port int
}

type GRPCApi struct {
	Port int
}

func (api *Api) initConfig() error {

	api.IP = "127.0.0.1"
	api.Production = viper.GetBool("Production.Status")
	api.CertFile = viper.GetString("Production.FullPATH")
	api.KeyFile = viper.GetString("Production.PrivPATH")
	api.Rest.Port = viper.GetInt("Microservice.Port")
	api.GRPC.Port = viper.GetInt("ProtoGRPC.Port")

	if api.GRPC.Port == 0 || api.Rest.Port == 0 {
		return fmt.Errorf("no se registro puerto GRPC o REST en el archivo settings.yaml")
	}

	if api.Production {
		if api.CertFile == "" || api.KeyFile == "" {
			return fmt.Errorf("faltan las llaves para un estado productivo de la api, sobre el archivo settings.yaml")
		}
	}

	return nil
}
