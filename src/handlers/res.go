package handlers

import (
	"net/http"
)

//Planned to be used as File Explorer - currently unused
func HandleRessources(w http.ResponseWriter, r *http.Request) {

	empty, err := isEmpty("./coursedata/ressources")
	if empty {
		WriteMsg(w, MsgNoRessources)
		return
	}
	if err != nil {
		WriteError(w, "Ressources folder cannot be read", err)
		return
	}

}
