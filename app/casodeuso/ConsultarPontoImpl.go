package casodeuso

import (
	"github.com/mvgv/lambda-relatorios/app/dominio"
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/repositorio"
)

type ConsultarPontoImpl struct {
	repositorio repositorio.PontoRepositorio
}

func NewConsultarPontoImpl(repositorio repositorio.PontoRepositorio) *ConsultarPontoImpl {
	return &ConsultarPontoImpl{
		repositorio: repositorio,
	}
}

func (c *ConsultarPontoImpl) ConsultarPontoDoMes(email, mes string) (*dominio.PontoDoDia, error) {

	registrosDoDia, err := c.repositorio.ConsultarPontoDoMes(email, mes)
	if err != nil {
		return nil, err
	}
	ponto := make([]dominio.Ponto, len(registrosDoDia.Registros))
	for i, registro := range registrosDoDia.Registros {
		ponto[i] = *dominio.NewPonto(registro.Timestamp, registro.Evento)
	}
	pontoDodia := dominio.NewPontoDoDia(email, ponto)

	return pontoDodia, nil
}
