package tgbotpackage

import (
	"github.com/Syfaro/telegram-bot-api"
	"github.com/go-martini/martini"
	"log"
)
type TGBot struct {
	bot     *tgbotapi.BotAPI
	uconfig tgbotapi.UpdateConfig
	uchan tgbotapi.UpdatesChannel
}

var main_bot TGBot

func updateScan(){
	for update := range main_bot.uchan {
		if update.Message == nil {
			continue
		}
		log.Println(update)
		log.Println(update.Message)
	}
}

func InitBot(token string)(err error){

	main_bot.bot,err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	main_bot.uconfig = tgbotapi.NewUpdate(0)
	main_bot.uchan, err = main_bot.bot.GetUpdatesChan(main_bot.uconfig)
	main_bot.uconfig.Timeout=60
	go updateScan()
	return err
}

func (bot *TGBot)SendMessage(message string, chan_id int64) {
	msg := tgbotapi.NewMessage(chan_id,message)
	bot.bot.Send(msg)
}

func Middleware(ctx martini.Context) {

	ctx.Map(&main_bot)

	ctx.Next()

}







