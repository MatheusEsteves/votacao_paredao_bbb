package models

import "time"

type Voto struct {
    ID            string    `json:"id" bson:"_id,omitempty"`
    Participante  string    `json:"participante" bson:"participante"`
    Timestamp     time.Time `json:"timestamp" bson:"timestamp"`
}