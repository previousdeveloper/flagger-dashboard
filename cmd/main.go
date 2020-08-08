package main

import (
	"github.com/previousdeveloper/flagger-dashboard/pkg/client"
	"github.com/previousdeveloper/flagger-dashboard/pkg/controller"
	"github.com/previousdeveloper/flagger-dashboard/pkg/server"
)

func main() {

	sClient := client.NewK8sClient()

	exampleController := controller.CanaryController{K8sClient: sClient}
	newServer := server.NewServer(&exampleController)
	newServer.Start()
}
