package servicos

import (
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/dto"
)

type ServicoCliente interface {
	GetCliente(email string) (*dto.ClienteRequisicao, error)
}
