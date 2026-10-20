package hscontrol

import (
	"net/http"
	"strings"

	"github.com/juanfont/headscale/hscontrol/templates"
	"github.com/juanfont/headscale/hscontrol/types"
	"github.com/rs/zerolog/log"
)

// The issance way to join a network: the registration page takes the customer's issance key (a valid
// pre-auth key of their user) and registers the waiting device to that user, through the same path as
// `tower auth register`. Only someone holding the key can register a device, so the POST cannot be forged
// by a link, and a GET still registers nothing.

// issanceTower is the running server, set in NewHeadscale.
var issanceTower *Headscale

func issanceRegisterRender(w http.ResponseWriter, status int, authID types.AuthID, message string, done bool) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(templates.IssanceRegister("/register/"+authID.String(), message, done).Render()))
}

func issanceRegisterPost(w http.ResponseWriter, req *http.Request) {
	authID, err := authIDFromRequest(req)
	if err != nil {
		httpError(w, err)
		return
	}
	if err := req.ParseForm(); err != nil {
		issanceRegisterRender(w, http.StatusBadRequest, authID, "The form could not be read. Try again.", false)
		return
	}
	key := strings.TrimSpace(req.PostFormValue("key"))
	pak, err := issanceTower.state.GetPreAuthKey(key)
	if err != nil || pak.Validate() != nil || pak.User == nil {
		log.Info().Str("auth_id", authID.String()).Msg("issance key rejected on the registration page")
		issanceRegisterRender(w, http.StatusForbidden, authID, "That issance key is not valid. Check it and try again.", false)
		return
	}
	node, nodeChange, err := issanceTower.state.HandleNodeFromAuthPath(authID, types.UserID(pak.User.ID), nil, "issance-key")
	if err != nil {
		issanceTower.Change(nodeChange)
		issanceRegisterRender(w, http.StatusNotFound, authID, "This request expired. Open the Tailscale app and sign in again.", false)
		return
	}
	routeChange, err := issanceTower.state.AutoApproveRoutes(node)
	issanceTower.Change(nodeChange, routeChange)
	if err != nil {
		log.Error().Err(err).Msg("auto approving routes after issance key registration")
	}
	issanceRegisterRender(w, http.StatusOK, authID, "", true)
}
