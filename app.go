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

	for {
		if bot.ReceiveData(){
				go func() {
					bot.Handler("callback_query", func(event string) {
						if bot.ChatID == id {
							if event == "!GenTeste"{
								bot.ForceReply("Quantas horas:")
							}else if event == "!CreateUser"{
								bot.ForceReply("Usuário:")
							}else if event == "!relatorio"{
								cmd := exec.Command("php", "/opt/DragonCore/menu.php", "relatoriouser")

								output, err := cmd.Output()
								if err != nil {
									fmt.Println("Erro ao executar o comando:", err)
									return
								}
								outputStr := string(output)
								bot.SendMessages(outputStr)
							}
						}else{
							bot.SendMessages("Não tem permição para Utilizar esse bot!")
						}

					})
				}()
				go func() { 
					bot.Handler("commands",func(event string) {
						if bot.ChatID == id {
							if event == "/start"{

								bot.SendMessages(`
				--------------------------------
				SEJA BEM VINDO(a) AO BOT DragonCore
				--------------------------------
				para acessar o basta digitar /menu 
								`)
							} else if event == "/menu"{
								layout := map[string]interface{}{
									"inline_keyboard": [][]map[string]interface{}{
										{
											{"text": "Criar Usúario", "callback_data": "!CreateUser"},
											{"text": "Gerar Teste", "callback_data": "!GenTeste"},
										},
										{
											{"text": "Alterar Senha", "callback_data": "!suporte"},
											{"text": "Alterar Limite", "callback_data": "!painel"},
										},
										{
											{"text": "Alterar Data", "callback_data": "!suporte"},
											{"text": "Remover", "callback_data": "!painel"},
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
							
							if bot.ReplyMessageText == "Quantas horas:"{
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
							}else if bot.ReplyMessageText == "Usuário:"{
								user = append(user, bot.Text)
								bot.ForceReply("Senha:")
							}else if bot.ReplyMessageText == "Senha:"{
								user = append(user, bot.Text)
                                bot.ForceReply("Limite:")
							}else if bot.ReplyMessageText == "Limite:"{
								user = append(user, bot.Text)
                                bot.ForceReply("Data:")
							}else if bot.ReplyMessageText == "Data:"{
								
								user = append(user, bot.Text)
								fmt.Println(user)
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
