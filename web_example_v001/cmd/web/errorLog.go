package cmd
import (
	//"fmt"
	"net/http"
	"text/template"
	"tracker/cmd/task/errorhandler"
)

type errorLogType struct {
	PageName         string
	Values           []errorhandler.ErrorResponse
}

func ServeHTTPerrorLog(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

	if (r.URL.Path == ErrorLogURL) {

		tmpl, err := template.ParseFiles(ErrorLogPage, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataCode, err := errorhandler.LoadErrorFromJSONFile()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := errorLogType{PageName:     "ErrorLog", 
							 Values:        dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	} 

}



 