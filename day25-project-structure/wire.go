package main

import (
	"go-basics/day25-project-structure/handler"
	"go-basics/day25-project-structure/repository"
	"go-basics/day25-project-structure/service"

	"go-basics/day25-project-structure/config"
)

func wireHandlers() *handler.Handlers {
	return &handler.Handlers{
		User: handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(config.DB))),
		Post: handler.NewPostHandler(service.NewPostService(repository.NewPostRepository(config.DB))),
	}
}
