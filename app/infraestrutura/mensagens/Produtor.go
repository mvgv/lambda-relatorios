package mensagens

import (
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/dto"
)

type Produtor interface {
	EnviarMensagem(mensagem *dto.SolicitacaoRelatorio) error
}
