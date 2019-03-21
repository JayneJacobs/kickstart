package handlers

import (
	"context"

	"github.com/JayneJacobs/FullStackWebDev/kickstart/client/common"

	"go.isomorphicgo.org/go/isokit"
)

func FriendsHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		println("Client-side Friends Handler")
		m := make(map[string]string)
		m["PageTitle"] = "Friends"
		env.TemplateSet.Render("friends_content", &isokit.RenderParams{Data: m, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})

	})
}
