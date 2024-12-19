package main

import (
    "os"
    "log"
    "votacao-paredao-bbb/adapter/db"
    "votacao-paredao-bbb/adapter/queue"
    "votacao-paredao-bbb/adapter/router"
    "votacao-paredao-bbb/core/services"
)

func main() {
    log.Printf("Inicialização da app golang")

    rabbitmqURL := os.Getenv("RABBITMQ_URL")
	mongoURL := os.Getenv("MONGO_URL")

    client, err := db.ConectarMongoDB(mongoURL)
    if err != nil {
        log.Fatal(err)
    }

    votoRepo := db.NovoVotoMongoRepository(client)
    rabbitMQ, err := queue.NovoRabbitMQ(rabbitmqURL, "votos", votoRepo)
    if err != nil {
        log.Fatal(err)
    }

    err = rabbitMQ.ConsumirFila()
    if err != nil {
        log.Fatal(err)
    }

    votoService := services.NovoVotoService(votoRepo, rabbitMQ)

    r := router.SetupRouter(votoService)
    r.Run(":8080")
}