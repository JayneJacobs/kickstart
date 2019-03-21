package handlers

import (
	"context"

	"github.com/JayneJacobs/FullStackWebDev/kickstart/client/common"

	"go.isomorphicgo.org/go/isokit"
)

func FeedHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		println("Client-side Feed Handler")
		m := make(map[string]string)
		m["PageTitle"] = "Feed"
		env.TemplateSet.Render("feed_content", &isokit.RenderParams{Data: m, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})
	})
}
