package main

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"html/template"
	"TelegramSupportBot/route"
	"TelegramSupportBot/session"
	"TelegramSupportBot/settings"
	"TelegramSupportBot/tgbot"
)

func unescape(x string) interface{} {
	return template.HTML(x)
}

func main()  {



	fmt.Println("Listening on port :"+settings.Port())

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	tgbotpackage.InitBot(settings.BotId())
	m.Use(session.Middleware)
	m.Use(tgbotpackage.Middleware)
	m.Use(render.Renderer(render.Options{
		Directory:  "template",                          // Specify what path to load the templates from.
		Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html", ".tpl"},  // Specify extensions to load for templates.
		Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
	}))


	//handlers
	m.Get("/",route.IndexHandler)
	m.Post("/",route.SupportSendHandler)
	m.Use(martini.Static("template/js", martini.StaticOptions{Prefix: "js"}))
	m.Use(martini.Static("template/css", martini.StaticOptions{Prefix: "css"}))
	m.Use(martini.Static("template/resources", martini.StaticOptions{Prefix: "res"}))
	//
	//m.Get("/", routes.IndexHandler)
	//m.Get("/login", routes.GetLoginHandler)
	//m.Get("/logout", routes.LogoutHandler)
	//m.Post("/login", routes.PostLoginHandler)
	//m.Get("/write", routes.WriteHandler)
	//m.Get("/edit/:id", routes.EditHandler)
	//m.Get("/view/:id", routes.ViewHandler)
	//m.Get("/delete/:id", routes.DeleteHandler)
	//m.Post("/SavePost", routes.SavePostHandler)
	//m.Post("/gethtml", routes.GetHtmlHandler)

	m.RunOnAddr(":"+settings.Port())


}