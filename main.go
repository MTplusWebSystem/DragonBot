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

alterar validade [php /opt/DragonCore/menu.php alterardata $usuario $validade]

alterar limite [php /opt/DragonCore/menu.php uplimit $usuario $limite]

alterar senhs [php /opt/DragonCore/menu.php uppass $usuario $senha ]

deletar usuario [ php /opt/DragonCore/menu.php deluser $usuario ]


*/
type DataStorage struct{
	NewData	bool
	NewPass bool
	NewLimiter bool
}
func main() {
	var token string
	var id int

	fmt.Print("Token:")
	fmt.Scan(&token)

	fmt.Print("ID:")
	fmt.Scan(&id)
	dataStore := DataStorage{}
	bot := botkit.BotInit{
		Token: token,
	}

	user := make([]string, 0)
	newdate := make([]string, 0)
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
								case "!AlterarLM":
									bot.ForceReply("Nome do Usuário:")
									dataStore.NewLimiter = true
								case "!AlterarDT":
									bot.ForceReply("Nome do Usuário:")
									dataStore.NewData = true
								case "!AlterarSH":
									bot.ForceReply("Nome do Usuário:")
                                    dataStore.NewPass = true
								case "!Relatorio":
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
												{"text": "Remover", "callback_data": "!Deletar"},
											},
											{
												{"text": "Relatório", "callback_data": "!Relatorio"},
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
									if dataStore.NewData == true{
										newdate = append(newdate, bot.Text)
                                        bot.ForceReply("Nova data:")
										dataStore.NewData = false
									}else if dataStore.NewLimiter == true{
										newdate = append(newdate, bot.Text)
                                        bot.ForceReply("Novo limite:")
										dataStore.NewLimiter = false
									}else if dataStore.NewPass == true{
										newdate = append(newdate, bot.Text)
                                        bot.ForceReply("Nova senha:")
                                        dataStore.NewPass = false
									}else {
										cmd := exec.Command("php", "/opt/DragonCore/menu.php", "deluser", bot.Text)
										cmd.Run()
										bot.SendMessages("Usuário deletado com sucesso")
									}

								case "Nova data:":
									cmd := exec.Command("php", "/opt/DragonCore/menu.php", "alterardata",newdate[0], bot.Text)
									cmd.Run()
									bot.SendMessages("Data alterada com sucesso")
								case "Novo limite:":
									cmd := exec.Command("php", "/opt/DragonCore/menu.php", "uplimit",newdate[0], bot.Text)
									cmd.Run()
									bot.SendMessages("Limite alterado com sucesso")
								case "Nova senha:":
									cmd := exec.Command("php", "/opt/DragonCore/menu.php", "uppass",newdate[0], bot.Text)
									cmd.Run()
									bot.SendMessages("Senha alterada com sucesso")
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
