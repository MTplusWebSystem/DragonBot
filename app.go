package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"github.com/MTplusWebSystem/GoBotKit/botkit"
)
/*

gerar teste   [ php /opt/DragonCore/menu.php gerarteste $validade ]

criar usuário [php /opt/DragonCore/menu.php criaruser $validade $usuario $senha $limite ]

criar relatório [php /opt/DragonCore/menu.php relatoriouser]

alterar  validade [php /opt/DragonCore/menu.php alterardata $usuario $validade]

deletar usuario [ php /opt/DragonCore/menu.php deluser $usuario ]


*/
func main() {
	var token string
	var id int

	fmt.Print("Token:")
	fmt.Scan(&token)

	fmt.Print("ID:")
	fmt.Scan(&id)

	bot := botkit.BotInit{
		Token: token,
	}
	user := make([]string, 0)
	//newdate := make([]string, 0)
	for {
		if bot.ReceiveData(){
				go func() {
					bot.Handler("callback_query", func(event string) {
						if bot.ChatID == id {
							switch event {
								case "!GenTeste":
									bot.ForceReply("Quantas horas:")
								case "!CreateUser":
									bot.ForceReply("Usuário:")
								case "!relatorio":
									cmd := exec.Command("php", "/opt/DragonCore/menu.php", "relatoriouser")

									output, err := cmd.Output()
									if err != nil {
										fmt.Println("Erro ao executar o comando:", err)
										return
									}
									outputStr := string(output)
									bot.SendMessages(outputStr)
								case "!Deletar":
									bot.ForceReply("Nome do usuário:")
							}
						}else{
							bot.SendMessages("Não tem permição para Utilizar esse bot!")
						}

					})
				}()
				go func() { 
					bot.Handler("commands",func(event string) {
						if bot.ChatID == id {
							switch event {
								case "/start":
									bot.SendMessages(`
				--------------------------------
				SEJA BEM VINDO(a) AO BOT DragonCore
				--------------------------------
				para acessar o basta digitar /menu 
								`)
								case "/menu":
									layout := map[string]interface{}{
										"inline_keyboard": [][]map[string]interface{}{
											{
												{"text": "Criar Usúario", "callback_data": "!CreateUser"},
												{"text": "Gerar Teste", "callback_data": "!GenTeste"},
											},
											{
												{"text": "Alterar Senha", "callback_data": "!AlterarSH"},
												{"text": "Alterar Limite", "callback_data": "!AlterarLM"},
											},
											{
												{"text": "Alterar Data", "callback_data": "!AlterarDT"},
												{"text": "Remover", "callback_data": "!deleterar"},
											},
											{
												{"text": "Relatório", "callback_data": "!relatorio"},
											},
										},
									}
					
									bot.SendButton(`
					--------------------------------
					Menu de gerenciamento DragonCore
					--------------------------------
									`, layout)
							}
						}else{
							bot.SendMessages("Não tem permição para Utilizar esse bot!")
						}
						
					})
				}()
				go func() {
					bot.Handler("messages",func(event string) {
						if bot.ChatID == id {
							switch bot.ReplyMessageText {
								case "Usuário:":
									user = append(user, bot.Text)
									bot.ForceReply("Senha:")
                                case "Quantas horas:":
									horas, err := strconv.Atoi(bot.Text)
									if err != nil {
										fmt.Println("Erro ao converter horas:", err)
										return
									}
									calc := 60 * horas
									cmd := exec.Command("php", "/opt/DragonCore/menu.php", "gerarteste", strconv.Itoa(calc))

									output, err := cmd.Output()
									if err != nil {
										fmt.Println("Erro ao executar o comando:", err)
										return
									}
									outputStr := string(output)
									bot.SendMessages(outputStr)
                                case "Senha:":
									user = append(user, bot.Text)
									bot.ForceReply("Limite:")
                                case "Limite:":
									user = append(user, bot.Text)
									bot.ForceReply("Data em dias:")
                                case "Data em dias:":
									user = append(user, bot.Text)
									cmd := exec.Command("php", "/opt/DragonCore/menu.php", "criaruser", user[3], user[0], user[1] , user[2] )
									cmd.Run()
									bot.SendMessages("Usuario criados com sucesso")
                                case "Nome do Usuário:":
									cmd := exec.Command("php", "/opt/DragonCore/menu.php", "deluser", bot.Text)
									cmd.Run()
									bot.SendMessages("Usuário deletado com sucesso")
                            }
						} else {
							bot.SendMessages("Não tem permissão para utilizar esse bot!")
						}
					})
				}()
				time.Sleep( 1 * time.Second)
		}
	}
}
