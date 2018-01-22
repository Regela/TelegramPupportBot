package route

import (
"github.com/martini-contrib/render"
"TelegramSupportBot/session"
"TelegramSupportBot/models"
	"net/http"
	//"net/url"
	//"fmt"
	//"TelegramSupportBot/settings"
	"TelegramSupportBot/tgbot"
	"TelegramSupportBot/settings"
)

func IndexHandler(rnd render.Render,session *session.Session){
	test := models.BaseModel{}
	rnd.HTML(200, "index", test)
}

func SupportSendHandler(rnd render.Render,session *session.Session,r *http.Request, bot *tgbotpackage.TGBot){

	//req, err := http.PostForm("https://api.telegram.org/bot"+settings.BotId()+"/sendMessage",
	//	url.Values{"chat_id": {settings.ChannelId()}, "text": {r.FormValue("message")+"\n"+r.FormValue("name")+" ("+r.FormValue("tel")+")"}})
	//
	//fmt.Println(req)

	bot.SendMessage(r.FormValue("message")+"\n"+r.FormValue("first_name")+" "+r.FormValue("last_name") +" ("+r.FormValue("tel")+")\nBrowser: "+r.UserAgent() +" IP:"+ r.RemoteAddr,settings.ChannelId())

	rnd.HTML(200, "success", nil)
}