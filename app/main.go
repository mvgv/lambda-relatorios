package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mvgv/lambda-relatorios/app/apresentacao"
	"github.com/mvgv/lambda-relatorios/app/casodeuso"
	"github.com/mvgv/lambda-relatorios/app/controlador"
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/mensagens"
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/repositorio"
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/servicos"
)

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) {
	clienteService := servicos.NewServicoClienteImpl()             //falta implementar
	pontoRepository := repositorio.NewPontoRepositorioDynamoImpl() //falta implementar
	messageSNS := mensagens.NewProdutorSNSImpl()                   //falta implementar

	consultarClienteUC := casodeuso.NewConsultarClienteImpl(clienteService)
	gerarRelatorioUC := casodeuso.NewGerarRelatorioImpl(messageSNS)
	consultarPontoUC := casodeuso.NewConsultarPontoImpl(pontoRepository)
	calcularHorasTrabalhadasUC := casodeuso.NewCalcularHorasTrabalhadasImpl()

	controller := controlador.NewConsultaPontoController(consultarClienteUC, consultarPontoUC, calcularHorasTrabalhadasUC, gerarRelatorioUC)

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		var mensagem apresentacao.MensagemSNS
		err := json.Unmarshal([]byte(snsRecord.Message), &mensagem)
		if err != nil {
			log.Printf("Error unmarshalling SNS message: %s", err)
			continue
		}

		controller.Handle(mensagem.Email, mensagem.Mes)
		log.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
