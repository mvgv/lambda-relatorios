package casodeuso

import "github.com/mvgv/lambda-relatorios/app/dominio"

type ConsultarPonto interface {
	ConsultarPontoDoMes(email, mes string) (*dominio.PontoDoDia, error)
}
