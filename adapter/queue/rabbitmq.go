package queue

import (
    "encoding/json"
    "github.com/streadway/amqp"
    "votacao-paredao-bbb/core/models"
	"votacao-paredao-bbb/core/ports"
    "log"
)

type RabbitMQ struct {
	VotoRepo ports.VotoRepository
    Channel *amqp.Channel
    Queue   amqp.Queue
}

func NovoRabbitMQ(amqpURI, queueName string, votoRepo ports.VotoRepository) (*RabbitMQ, error) {
    log.Printf("Iniciando criação do canal para o RabbitMQ")

    conn, err := amqp.Dial(amqpURI)
    if err != nil {
        return nil, err
    }

    channel, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    log.Println("Conectado ao RabbitMQ com sucesso!")

    queue, err := channel.QueueDeclare(
        queueName,
        false, 
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return nil, err
    }

    return &RabbitMQ{VotoRepo: votoRepo, Channel: channel, Queue: queue}, nil
}

func (r *RabbitMQ) EnfileirarVoto(voto models.Voto) error {
    votoJSON, err := json.Marshal(voto)
    if err != nil {
        return err
    }

    err = r.Channel.Publish(
        "",
        r.Queue.Name,
        false,
        false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        votoJSON,
        },
    )
    return err
}

func (r *RabbitMQ) ConsumirFila() error {
    msgs, err := r.Channel.Consume(
        r.Queue.Name,
        "", 
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Printf("Erro ao tentar consumir fila")
        return err
    }

    go func() {
        for msg := range msgs {
            var voto models.Voto
            err := json.Unmarshal(msg.Body, &voto)

            if err != nil {
                log.Printf("Erro ao processar voto: %v", err)
                continue
            }

            log.Printf("Novo voto recebido para o participante %v", voto.Participante)
            
            err = r.VotoRepo.SalvarVoto(voto)
            if err != nil {
                log.Printf("Erro ao salvar voto na base: %v", err)
                continue
            }
        }
    }()

    return nil
}