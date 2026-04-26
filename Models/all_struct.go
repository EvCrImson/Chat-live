package Models

type Usuario struct {
	Name  string
	Senha string
}

type Users struct {
	User_id       int
	Username      string
	Password_hash string
}

type Dados struct {
	Password_hash string
	User_id       string
}

type Mensagens struct {
	Id_mensagens          int
	Mensagem              string
	Mensagem_enviado_por  int
	Mensagem_recebida_por int
}

type Mensagem_para_criar struct {
	Mensagem              string
	Mensagem_recebida_por int
}

type Refressrequest struct {
	Refresstoken string
}

type Mensagens_para_atualizar struct {
	Id_mensagens          int
	Mensagem_para_atulizar string
	Mensagem_enviado_por  int
	Mensagem_recebida_por int
	Mensagem_antiga 	  string
}


type Deletar_usuario struct{
	User_id string
	Usuario string
}