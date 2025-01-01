package main

import (
	"onion/internal/infrastructure/config"
	errs "onion/internal/infrastructure/error"
	"onion/internal/infrastructure/logger"
	"onion/internal/presentation/router"
	"onion/internal/usecase"
	"onion/internal/usecase/articles"
)

func main() {
	config := config.New()
	if err := config.Setup(); err != nil {
		panic(err)
	}

	var l *logger.Logger
	if config.Server.IsLocal() {
		l = logger.New(logger.NewDebugHandler())
	} else {
		l = logger.New(logger.NewJSONHandler())
	}

	rt := router.New(config.Server.Port, l)
	baseUsecase := usecase.NewBaseUsecase(l, errs.New())
	rt.AddFetchArticlesHandler(
		articles.NewFetchArticlesUsecase(baseUsecase),
	)
	if err := rt.Run(); err != nil {
		panic(err)
	}
}
