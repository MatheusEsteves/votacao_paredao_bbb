package main

import (
    "log"
    "votacao-paredao-bbb/adapter/db"
    "votacao-paredao-bbb/adapter/queue"
    "votacao-paredao-bbb/adapter/router"
    "votacao-paredao-bbb/core/services"
)

func main() {
    client, err := db.ConectarMongoDB("mongodb://localhost:27017")
    if err != nil {
        log.Fatal(err)
    }

    rabbitMQ, err := queue.NovoRabbitMQ("amqp://guest:guest@localhost:5672/", "votos")
    if err != nil {
        log.Fatal(err)
    }

    votoRepo := db.NovoVotoMongoRepository(client)
    votoService := services.NovoVotoService(votoRepo, rabbitMQ)

    r := router.SetupRouter(votoService)
    r.Run(":8080")
}