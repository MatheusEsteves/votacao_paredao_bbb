package router

import (
    "github.com/gin-gonic/gin"
    "votacao-paredao-bbb/adapter/handlers"
    "votacao-paredao-bbb/core/services"
)

func SetupRouter(votoService *services.VotoService) *gin.Engine {
    router := gin.Default()
    votoHandler := &handlers.VotoHandler{VotoService: votoService}

    router.POST("/voto", votoHandler.RegistrarVoto)
    router.GET("/resultados", votoHandler.ObterResultados)

    return router
}