package repositorio

import "github.com/mvgv/lambda-relatorios/app/infraestrutura/dto"

type PontoRepositorio interface {
	ConsultarPontoDoMes(email, mes string) (*dto.PontoDoDiaEntidade, error)
}
