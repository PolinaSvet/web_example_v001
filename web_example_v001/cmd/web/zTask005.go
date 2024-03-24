package cmd
import (
	//"fmt"
	"net/http"
	"text/template"
	"encoding/json"
	"tracker/cmd/task/criminal"
)

type TypeTask005 struct {
	PageName         string
	PageDescribe     string
	PageTask         string
	TaskCode         string
	Values           map[string]criminal.Man
}

type TypeTask005list struct {
	Values           map[string]criminal.Man
}


func processHandlerTask005Show(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

    if r.Method == http.MethodPost  {

        type RequestData struct {
			ListSuspects     string   `json:"listSuspects"`
	    }

		var requestData RequestData
	    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	    }

		dataCode, resp := criminal.FindCriminalMan(requestData.ListSuspects)
		response := resp	 
		
		tmpl, err := template.ParseFiles(Task005PList, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask005list{ Values:  dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")	
		json.NewEncoder(w).Encode(response)

    }
}

func ServeHTTPTask005(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

	if (r.URL.Path == Task005URL) {

		tmpl, err := template.ParseFiles(Task005Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataCode, err := criminal.LoadFromJSONFile(CriminalInfoJson)

		data := TypeTask005{PageName:     ReplaceTxtBr(ListTasksAll[8].Name, "; "), 
		                     PageDescribe: ReplaceTxtBr(ListTasksAll[8].Describe, " "),
		                     PageTask:     ListTasksAll[8].Task,
							 TaskCode:     "",
							 Values:       dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	} 

}



 