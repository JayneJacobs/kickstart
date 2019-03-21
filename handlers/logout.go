package handlers

import (
	"net/http"

	"github.com/JayneJacobs/FullStackWebDev/kickstart/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
