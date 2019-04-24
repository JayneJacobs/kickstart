package main

import (
	"strings"

	"github.com/JayneJacobs/FullStackWebDev/kickstart/client/handlers"

	"github.com/JayneJacobs/FullStackWebDev/kickstart/client/common"

	"go.isomorphicgo.org/go/isokit"
	"honnef.co/go/js/dom"
)

// D is a shortcut for the dom functions
var D = dom.GetWindow().Document().(dom.HTMLDocument)

func initializeEventHandlers(env *common.Env) {
	println("location: ", env.Window.Location().Href)
	l := strings.Split(env.Window.Location().Href, "/")
	pageName := l[len(l)-1]

	switch pageName {

	}
}

func run() {
	println("Kickstart is on the client-side, thanks to GopherJS!")

	templateSetChannel := make(chan *isokit.TemplateSet)
	go isokit.FetchTemplateBundle(templateSetChannel)
	ts := <-templateSetChannel

	env := common.Env{}
	env.TemplateSet = ts
	env.Window = dom.GetWindow()
	env.Document = dom.GetWindow().Document()
	env.PrimaryContent = env.Document.GetElementByID("primaryContent")

	r := isokit.NewRouter()
	r.Handle("/feed", handlers.FeedHandler(&env))
	r.Handle("/friends", handlers.FriendsHandler(&env))
	r.Handle("/profile", handlers.MyProfileHandler(&env))
	r.Listen()

	initializeEventHandlers(&env)

}

func main() {
	switch readyState := D.ReadyState(); readyState {
	case "loading":
		D.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go run()
		})
	case "interactive", "complete":
		run()
	default:
		println("Unexpected document.ReadyState value!")
	}
}
