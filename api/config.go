package api

import (
	"fmt"

	"github.com/spf13/viper"
)

type Api struct {
	IP         string
	Production bool
	FullPATH   string
	PrivPATH   string
	Rest       RESTApi
	GRPC       GRPCApi
}

type RESTApi struct {
	Port int
}

type GRPCApi struct {
	Port int
}

func (Api) InitConfig() error {

	config := Api{
		IP:         "127.0.0.1",
		Production: viper.GetBool("Production.Status"),
		FullPATH:   viper.GetString("Production.FullPATH"),
		PrivPATH:   viper.GetString("Production.PrivPATH"),
		Rest:       RESTApi{Port: viper.GetInt("Microservice.Port")},
		GRPC:       GRPCApi{Port: viper.GetInt("ProtoGRPC.Port")},
	}

	if config.GRPC.Port == 0 || config.Rest.Port == 0 {
		return fmt.Errorf("no se registro puerto en el archivo settings.yaml")
	}

	if config.Production {
		if config.FullPATH == "" || config.PrivPATH == "" {
			return fmt.Errorf("faltan las llaves para un estado productivo de la api, sobre el archivo settings.yaml")
		}
	}

	return nil
}

// func startGRPCServer(address, certFile, keyFile string) error {
// 	lis, err := net.Listen("tcp", address)
// 	if err != nil {
// 		return fmt.Errorf("failed to listen: %v", err)
// 	}

// 	s := h.Server{}

// 	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
// 	if err != nil {
// 		return fmt.Errorf("could not load TLS keys: %s", err)
// 	}

// 	opts := []grpc.ServerOption{grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor)}

// 	grpcServer := grpc.NewServer(opts...)

// 	dbproto.RegisterDBServiceServer(grpcServer, &s)

// 	log.Printf("Start HTTP/2 gRPC server on: %s", address)
// 	if err := grpcServer.Serve(lis); err != nil {
// 		return fmt.Errorf("failed to serve: %s", err)
// 	}

// 	return nil
// }
