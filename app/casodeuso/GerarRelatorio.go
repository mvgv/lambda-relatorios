package casodeuso

type GerarRelatorio interface {
	GerarRelatorioMensal(listaPontos []string, horas string) error
}
