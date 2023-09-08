package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartAPI() {
	// Log file
	dateNow := time.Now()
	f, err := os.Create(fmt.Sprintf("%s.%d-%d-%d.log", viper.GetString("Microservice.Name"), dateNow.Day(), dateNow.Month(), dateNow.Year()))
	if err != nil {
		log.Fatalf("error al crear el archivo, err: %s", err.Error())
	}

	mw := io.MultiWriter(f)
	log.SetOutput(mw)
	logrus.SetOutput(mw)

	// Configuracion de apis
	var api Api

	err = api.initConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to init config: %s", err))
		return
	}

	go func() {
		err := api.startGRPCServer()
		if err != nil {
			log.Fatal(fmt.Errorf("failed to start gRPC server: %s", err))
		}
	}()

	go func() {
		err := api.startRESTServer()
		if err != nil {
			log.Fatal(fmt.Errorf("failed to start REST server: %s", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("signal caught. shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		logrus.Info("Server down.")
	}
}
