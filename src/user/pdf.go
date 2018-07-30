package user

import (
	"bytes"
	"fmt"

	"../courseconfig"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/russross/blackfriday"
)

func (us *User) GeneratePdf() {
	fmt.Println("AAH")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Load Posts, Markdownify and add to document
	//TODO: Überschrift
	cfg := courseconfig.GetConfig()
	//Generate Header
	text := []byte(fmt.Sprintf("<h1>%s - Abgaben %s %s %v</h1>\n\n",
		cfg.Course_name, us.Vorname, us.Nachname, us.Matrikel))

	Posts, _ := us.GetAllPosts()
	md := blackfriday.MarkdownCommon(append(text, Posts...))
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(md))
	page.FooterRight.Set("[page]")
	page.MinimumFontSize.Set(20)
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		fmt.Println(err)
		return
	}
	// Write buffer contents to file on disk
	filename := fmt.Sprintf("./%v_%s_%s.pdf",
		us.Matrikel,
		cfg.Course_name,
		us.Nachname)

	err = pdfg.WriteFile(filename)
	if err != nil {
		fmt.Println(err)
		//If fails, crete file with save name FURCHTBARER CODE - TODO
		pdfg.WriteFile(fmt.Sprintf("./%v_posts.pdf", us.Matrikel))
	}

	fmt.Println("Created Post History PDF for user", us.Matrikel)
}
