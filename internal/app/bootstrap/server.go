package bootstrap

import (
	"fmt"

	"github.com/MatheusHenriq/GoApiLearn.git/internal/app/handlers/locations"
	repositories "github.com/MatheusHenriq/GoApiLearn.git/internal/infraestructure/repositories/locations"
	"github.com/gin-gonic/gin"
)

//se a função vai ser acessada fora do package, ela deve de ter letra maiúscula primeiro
//se a função tiver letra minuscula no começo, ela é "privada"

func StartServer() {
	e := gin.Default()
	configureRoutes(e)
	err := e.Run(":8001")
	//nil é a ausencia de valor
	if err != nil {
		panic(err)
	}
	fmt.Println("Server Started")
}

func configureRoutes(e *gin.Engine) {
	locationReponsitory := repositories.NewLocationRepository()
	locationHandler := locations.NewLocationHandler(locationReponsitory)

	g := e.Group("/api/v1")
	{
		g.GET("/states", locationHandler.GetAllStates)
	}
}
