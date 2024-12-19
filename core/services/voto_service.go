package services

import (
    "errors"
    "time"
    "votacao-paredao-bbb/core/models"
    "votacao-paredao-bbb/core/ports"
)

type VotoService struct {
    VotoRepo ports.VotoRepository
}

func NovoVotoService(votoRepo ports.VotoRepository) *VotoService {
    return &VotoService{VotoRepo: votoRepo}
}

func (vs *VotoService) RegistrarVoto(participante string) error {
    voto := models.Voto{
        Participante: participante,
        Timestamp:    time.Now(),
    }

    return vs.VotoRepo.SalvarVoto(voto)
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

func (vs *VotoService) ObterVotosPorHora() (map[string]int, error) {
    votos, err := vs.VotoRepo.ObterVotos()
    if err != nil {
        return nil, err
    }

    resultadosPorHora := make(map[string]int)
    for _, voto := range votos {
        hora := voto.Timestamp.Format("2006-01-02 15")
        resultadosPorHora[hora]++
    }
    return resultadosPorHora, nil
}