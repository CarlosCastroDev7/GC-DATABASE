package api

import (
	"fmt"
)

func StartAPI() {
	// Log file
	// dateNow := time.Now()
	// f, err := os.Create(fmt.Sprintf("%s.%d-%d-%d.log", viper.GetString("Microservice.Name"), dateNow.Day(), dateNow.Month(), dateNow.Year()))
	// if err != nil {
	// 	log.Fatalf("error al crear el archivo, err: %s", err.Error())
	// }

	// mw := io.MultiWriter(f)
	// log.SetOutput(mw)
	// logrus.SetOutput(mw)

	var api Api

	api.InitConfig()

	fmt.Println(api)

	// go func() {
	// 	err := protocol.ApiGRPC(configApi)
	// 	if err != nil {
	// 		log.Fatal(fmt.Errorf("failed to start gRPC server: %s", err))
	// 	}
	// }()

	// // go func() {
	// // 	err := startRESTServer(restAddress, grpcAddress, certFile, keyFile)
	// // 	if err != nil {
	// // 		log.Fatal(fmt.Errorf("Failed to start REST server: %s", err))
	// // 	}
	// // }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit
	// logrus.Info("signal caught. shutting down...")

	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	// select {
	// case <-ctx.Done():
	// 	logrus.Info("Server down.")
	// }
}
