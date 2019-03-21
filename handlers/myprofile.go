package handlers

import (
	"net/http"

	"go.isomorphicgo.org/go/isokit"

	"github.com/JayneJacobs/FullStackWebDev/kickstart/common"
)

func MyProfileHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		m := make(map[string]string)
		m["PageTitle"] = "My Profile"
		env.TemplateSet.Render("profile_page", &isokit.RenderParams{Writer: w, Data: m})
	})
}
