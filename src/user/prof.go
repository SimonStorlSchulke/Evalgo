package user

type Prof struct {
	Vorname  string `json:"vorname"`
	Nachname string `json:"nachname"`
}

func NewProf(Vorname, Nachname string) Prof {
	return Prof{Vorname, Nachname}
}

/*
//Returns Folderpath to Studentdata as string
func (pr Prof) getPath() string {
	return filepath.Join(".", "Userdata", "Students", fmt.Sprintf("%v", st.Matrikel))
}

func (pr Prof) Unregister() {
	os.RemoveAll(pr.getPath())
	fmt.Printf("Unregistered Student %s %s at %s\n", pr.Vorname, pr.Nachname, st.getPath())
}

TODO: Register(), Unregister()
*/
