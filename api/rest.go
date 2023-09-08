package api

import (
	"context"
	"fmt"
	"gc/database/app/grpc/proto"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (api *Api) startRESTServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	creds, err := credentials.NewClientTLSFromFile(api.CertFile, "")
	if err != nil {
		return fmt.Errorf("cloud not load TLS certificate: %s", err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	err = proto.RegisterServiceDBHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", api.IP, api.GRPC.Port), opts)
	if err != nil {
		return fmt.Errorf("could not register service Ping: %s", err)
	}

	log.Printf("Start HTTP/1.1 with TLS REST server on %s:%d", api.IP, api.Rest.Port)
	if err := http.ListenAndServeTLS(fmt.Sprintf("%s:%d", api.IP, api.Rest.Port), api.CertFile, api.KeyFile, mux); err != nil {
		return err
	}

	return nil
}
