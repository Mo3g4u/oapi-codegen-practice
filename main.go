package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"hoge/gen"
)

type Server struct{}

type Response struct {
	Message string `json:"message"`
}

func (s Server) FindPets(ctx echo.Context, params gen.FindPetsParams) error {
	return ctx.JSON(http.StatusOK, Response{
		Message: "oapi-codegen: FindPets",
	})
}

func (s Server) AddPet(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Response{
		Message: "oapi-codegen: AddPet",
	})
}

func (s Server) DeletePet(ctx echo.Context, id int64) error {
	return ctx.JSON(http.StatusOK, Response{
		Message: "oapi-codegen: DeletePet",
	})
}

func (s Server) FindPetByID(ctx echo.Context, id int64) error {
	return ctx.JSON(http.StatusOK, Response{
		Message: "oapi-codegen: FindPetByID",
	})
}

func main() {
	e := echo.New()
	s := Server{}

	gen.RegisterHandlers(e, s)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
