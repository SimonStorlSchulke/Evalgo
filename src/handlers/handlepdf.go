package handlers

import (
	"fmt"
	"net/http"
)

func HandlePdf(w http.ResponseWriter, r *http.Request) {
	us, err := studentFromURL(r)
	if err != nil {
		fmt.Fprint(w, "error reading matrikel")
		return
	}

	//Check Permission
	if !checkViewPermission(us, r) {
		fmt.Fprintf(w, "Permission Denied")
		return
	}

	us.GeneratePdf()
	fmt.Fprint(w, "generated pdf")
}
