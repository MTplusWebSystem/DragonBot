package main

import (
	"fmt"
	"os/exec"

	"github.com/MTplusWebSystem/GoBotKit/botkit"
)

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


	for {
		if bot.ReceiveData(){
				go func() {
					bot.Handler("callback_query", func(event string) {
						if bot.ChatID == id {
							if event == "!GenTeste"{
								bot.ForceReplyToMessage(bot.QueryMessageID,"Quantas horas:")
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
											{"text": "Criar Usúario", "callback_data": "!suporte"},
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
								bot.SendMessages(fmt.Sprintf("Teste de %s horas",bot.Text))
								cmd := exec.Command("php /opt/DragonCore/menu.php gerarteste "+bot.Text)
								cmd.Run()
							}
						}else{
							bot.SendMessages("Não tem permição para Utilizar esse bot!")
						}
					})
				}()
			 
		}
	}
}
