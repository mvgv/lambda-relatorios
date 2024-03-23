package repositorio

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mvgv/lambda-relatorios/app/infraestrutura/dto"
)

type PontoRepositorioDynamoImpl struct {
	svc *dynamodb.DynamoDB
}

func NewPontoRepositorioDynamoImpl() *PontoRepositorioDynamoImpl {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	return &PontoRepositorioDynamoImpl{svc: dynamodb.New(sess)}
}

func (p *PontoRepositorioDynamoImpl) ConsultarPontoDoMes(email, mes string) (*dto.PontoDoDiaEntidade, error) {
	t, err := time.Parse("2006-01", mes)
	if err != nil {
		log.Fatalf("Erro ao analisar o mês: %v", err)
	}

	inicioMes := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).Format("2006-01-02T15:04:05")
	fimMes := time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 0, t.Location()).Format("2006-01-02T15:04:05")

	input := &dynamodb.QueryInput{
		TableName: aws.String("RegistrosPonto"),
		KeyConditions: map[string]*dynamodb.Condition{
			"email": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(email),
					},
				},
			},
			"timestamp": {
				ComparisonOperator: aws.String("BETWEEN"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(inicioMes),
					},
					{
						S: aws.String(fimMes),
					},
				},
			},
		},
	}

	result, err := p.svc.Query(input)
	if err != nil {
		return nil, err
	}

	registros := make([]dto.PontoEntidade, len(result.Items))

	for i, item := range result.Items {
		registros[i] = *dto.NewPontoEntidade(*item["email"].S, *item["timestamp"].S, *item["evento"].S)
	}

	pontoDodia := dto.NewPontoDoDiaEntidade(*result.Items[0]["email"].S, registros)
	fmt.Println("Ponto do dia: ", pontoDodia)
	return pontoDodia, nil

}
