package services

import (
    "time"
    "votacao-paredao-bbb/core/models"
    "votacao-paredao-bbb/core/ports"
)

type VotoService struct {
    VotoRepo ports.VotoRepository
    VotoQueue ports.VotoQueue
}

func NovoVotoService(votoRepo ports.VotoRepository, votoQueue ports.VotoQueue) *VotoService {
    return &VotoService{VotoRepo: votoRepo, VotoQueue: votoQueue,}
}

func (vs *VotoService) RegistrarVoto(participante string) error {
    voto := models.Voto{
        Participante: participante,
        Timestamp:    time.Now(),
    }

    err := vs.VotoQueue.EnfileirarVoto(voto)
    if err != nil {
        return err
    }

    return nil
}

func (vs *VotoService) ObterResultadosGeral() (map[string]int, error) {
    votos, err := vs.VotoRepo.ObterVotos()
    if err != nil {
        return nil, err
    }

    resultados := make(map[string]int)
    for _, voto := range votos {
        resultados[voto.Participante]++
    }
    return resultados, nil
}