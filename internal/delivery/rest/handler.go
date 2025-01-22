package rest

import "github.com/takumifahri/GO-Restaurant-Practice/internal/usecase/resto"

type handler struct {
	RestoUsecase resto.Usecase
}

func NewHandler( RestoUsecase resto.Usecase ) *handler {
	return &handler{
		RestoUsecase: RestoUsecase,
	}
}