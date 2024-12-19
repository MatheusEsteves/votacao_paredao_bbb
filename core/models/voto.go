package models

import "time"

type Voto struct {
    Participante  string    `json:"participante" bson:"participante"`
    Timestamp     time.Time `json:"timestamp" bson:"timestamp"`
}