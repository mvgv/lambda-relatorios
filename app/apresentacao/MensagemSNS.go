package apresentacao

type MensagemSNS struct {
	Email string `json:"email"`
	Mes   string `json:"mes"`
}

func NewMensagemSNS(email, mes string) *MensagemSNS {
	return &MensagemSNS{
		Email: email,
		Mes:   mes,
	}
}
