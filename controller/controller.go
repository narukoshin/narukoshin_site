package controller
import (
	"compress/gzip"
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
	// Setting headers to load the page instead of file downloading
	v.Writer.Header().Add("Content-Type", "text/html;charset=utf-8")
	v.Writer.Header().Add("Content-Encoding", "gzip")
	// Preparing writer for gzip writer
	gz := gzip.NewWriter(v.Writer)

	fileLocation := "templates/" + v.View + ".blade.html"
	defer gz.Close()
	if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
		t := template.Must(template.ParseFiles(fileLocation))
		err = t.Execute(gz, data)
		return err
	} else {
		gz.Write([]byte("Error: template file not found"))
		return err
	}
}