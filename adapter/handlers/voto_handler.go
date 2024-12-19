package handlers

import (
    "log"
    "github.com/gin-gonic/gin"
    "votacao-paredao-bbb/core/services"
)

type VotoHandler struct {
    VotoService *services.VotoService
}

func (vh *VotoHandler) RegistrarVoto(c *gin.Context) {
    var request struct {
        Participante string `json:"participante"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": "Participante é obrigatório"})
        return
    }

    err := vh.VotoService.RegistrarVoto(request.Participante)
    if err != nil {
        c.JSON(500, gin.H{"error": "Erro ao registrar voto"})
        return
    }

    c.JSON(200, gin.H{"message": "Voto registrado com sucesso"})
}

func (vh *VotoHandler) ObterResultados(c *gin.Context) {
    resultados, err := vh.VotoService.ObterResultadosGeral()
    if err != nil {
        log.Printf("Erro ao obter resultados : %v", err)
        c.JSON(500, gin.H{"error": "Erro ao obter resultados"})
        return
    }
    c.JSON(200, resultados)
}