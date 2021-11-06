package controller
import (
	"text/template"
	"net/http"
	"os"
)

type View struct {
	View string
	Writer http.ResponseWriter
}

type Page struct {
	Title string
}

func (v *View) Load(data interface{}) error {
	fileLocation := "templates/" + v.View + ".blade.html"
	if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
		t := template.Must(template.ParseFiles(fileLocation))
		err = t.Execute(v.Writer, data)
		return err
	} else {
		v.Writer.Write([]byte("Error: template file not found"))
		return err
	}
}