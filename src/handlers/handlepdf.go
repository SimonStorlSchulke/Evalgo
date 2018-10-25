package handlers

import (
	"fmt"
	"net/http"
)

func HandlePdf(w http.ResponseWriter, r *http.Request) {
	us, err := studentFromURL(r)
	if err != nil {
		WriteError(w, "cannot read matrikel", err)
		return
	}

	//Check Permission
	if !checkViewPermission(us, r) {
		WriteMsg(w, MsgPermissionDenied)
		return
	}

	us.GeneratePdf()
	fmt.Fprint(w, "generated pdf")
}
