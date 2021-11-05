package controller
import (
	"text/template"
	"net/http"
	"os"

	"logging"
)

type View struct {
	View string
	Writer http.ResponseWriter
}

type Page struct {
	Title string
}

func (v *View) Load(data interface{}) {
	fileLocation := "templates/" + v.View + ".blade.html"
	if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
		t := template.Must(template.ParseFiles(fileLocation))
		err = t.Execute(v.Writer, data)
		if err != nil {
			logging.Save(err)
		}
	} else {
		logging.Save(err)
	}
}