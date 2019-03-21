package handlers

import (
	"context"

	"github.com/JayneJacobs/FullStackWebDev/kickstart/client/common"

	"go.isomorphicgo.org/go/isokit"
)

func MyProfileHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		println("Client-side My Profile Handler")
		m := make(map[string]string)
		m["PageTitle"] = "My Profile"
		env.TemplateSet.Render("profile_content", &isokit.RenderParams{Data: m, Disposition: isokit.PlacementReplaceInnerContents, Element: env.PrimaryContent})
	})
}
