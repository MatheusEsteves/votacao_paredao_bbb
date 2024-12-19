package ports

import (
    "votacao-paredao-bbb/core/models"
)

type VotoRepository interface {
    SalvarVoto(voto models.Voto) error
    ObterVotos() ([]models.Voto, error)
}