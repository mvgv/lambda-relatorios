package casodeuso

import "github.com/mvgv/lambda-relatorios/app/dominio"

type CalcularHorasTrabalhadas interface {
	CalcularHorasTrabalhadasNoDia(listaPontos []dominio.Ponto) (string, error)
}
