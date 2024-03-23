package controlador

import (
	"fmt"

	"github.com/mvgv/lambda-relatorios/app/casodeuso"
)

type ConsultaPontoController struct {
	consultarClienteUC         casodeuso.ConsultarCliente
	consultarPontoUC           casodeuso.ConsultarPonto
	calcularHorasTrabalhadasUC casodeuso.CalcularHorasTrabalhadas
	gerarRelatorioUC           casodeuso.GerarRelatorio
}

func NewConsultaPontoController(consultarClienteUC casodeuso.ConsultarCliente,
	consultarPontoUC casodeuso.ConsultarPonto, calcularHorasTrabalhadasUC casodeuso.CalcularHorasTrabalhadas,
	gerarRelatorioUC casodeuso.GerarRelatorio) *ConsultaPontoController {
	return &ConsultaPontoController{
		consultarClienteUC:         consultarClienteUC,
		consultarPontoUC:           consultarPontoUC,
		calcularHorasTrabalhadasUC: calcularHorasTrabalhadasUC,
		gerarRelatorioUC:           gerarRelatorioUC,
	}
}

func (c *ConsultaPontoController) Handle(email, mes string) error {
	fmt.Println("ConsultaPontoController.Handle()")
	ponto, err := c.consultarPontoUC.ConsultarPontoDoMes(email, mes)
	if err != nil {
		return err
	}
	listaPontos := make([]string, len(ponto.Registros))
	for i, registro := range ponto.Registros {
		listaPontos[i] = registro.Horario
	}

	horasTrabalhas, _ := c.calcularHorasTrabalhadasUC.CalcularHorasTrabalhadasNoDia(ponto.Registros)

	err = c.gerarRelatorioUC.GerarRelatorioMensal(listaPontos, horasTrabalhas)

	if err != nil {
		return err
	}

	return nil
}
