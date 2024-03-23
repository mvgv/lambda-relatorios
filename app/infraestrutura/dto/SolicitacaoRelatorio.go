package dto

type SolicitacaoRelatorio struct {
	Registros []string `json:"registros"`
	Horas     string   `json:"horas"`
}

func NewSolicitacaoRelatorio(registros []string, horas string) *SolicitacaoRelatorio {
	return &SolicitacaoRelatorio{
		Registros: registros,
		Horas:     horas,
	}
}
