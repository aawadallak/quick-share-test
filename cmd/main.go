package main

import (
	"quick_share/config"
	"quick_share/infra/repository"
	"quick_share/server"
	"quick_share/usecases"
)

func main() {

	config.Init()
	config.InitLogger()
	s := server.NewServer()

	svc := usecases.NewService(&repository.FileRepository{}, &repository.DatabaseRepository{})

	go svc.CleanTempFiles()

	s.Run()

}
