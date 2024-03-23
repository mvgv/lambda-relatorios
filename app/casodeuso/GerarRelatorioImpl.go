package casodeuso

import (
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/dto"
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/mensagens"
)

type GerarRelatorioImpl struct {
	producer mensagens.Produtor
}

func NewGerarRelatorioImpl(producer mensagens.Produtor) *GerarRelatorioImpl {
	return &GerarRelatorioImpl{
		producer: producer,
	}
}

func (g *GerarRelatorioImpl) GerarRelatorioMensal(listaPontos []string, horas string) error {
	mensagem := dto.NewSolicitacaoRelatorio(listaPontos, horas)
	err := g.producer.EnviarMensagem(mensagem)
	if err != nil {
		return err
	}
	return nil
}
