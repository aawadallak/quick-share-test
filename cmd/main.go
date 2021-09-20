package main

import (
	"we/config"
	"we/infra/repository"
	"we/server"
	"we/usecases"
)

func main() {

	config.Init()
	config.InitLogger()
	s := server.NewServer()

	svc := usecases.NewService(&repository.FileRepository{}, &repository.DatabaseRepository{})

	go svc.CleanTempFiles()

	s.Run()

}
