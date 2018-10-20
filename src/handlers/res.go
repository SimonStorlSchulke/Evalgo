package handlers

import (
	"fmt"
	"net/http"
)

//Planned to be used as File Explorer - currently unused
func HandleRessources(w http.ResponseWriter, r *http.Request) {

	empty, err := isEmpty("./coursedata/ressources")
	if empty {
		fmt.Fprintf(w, "No Ressources")
		return
	}
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

}
