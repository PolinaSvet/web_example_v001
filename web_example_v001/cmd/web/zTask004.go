package cmd
import (
	"fmt"
	"net/http"
	"text/template"
	"encoding/json"
	auto "tracker/cmd/task/auto"
)

type TypeTask_004 struct {
	PageName         string
	PageDescribe     string
	PageTask         string
	TaskCode         string
	ValuesView       []string
}

type TypeTask_004_list struct {
	Values           []auto.ItemsHtml
}



func processHandlerTask004Show(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

    if r.Method == http.MethodPost {

		// Декодирование JSON-объекта и конвертация в тип UnitType
        type RequestData struct {
			ViewPhone string   `json:"viewPhone"`
			Brand     auto.UnitType `json:"brand"`
	    }

	    var requestData RequestData
	    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	    }
		
		tmpl, err := template.ParseFiles(Task004PList, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataCode, err := auto.LoadAutoFromJSONFile(AutoInfoJson, requestData.Brand) //auto.Inch //auto.CM
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_004_list{ Values:  dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response:= fmt.Sprintf("Габариты автомобилей выводятся в %v", requestData.Brand)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func ServeHTTPTask004(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

	if (r.URL.Path == Task004URL) {

		//v1:=5
		//v2:=0
		//fmt.Printf("1: %v\n", v1/v2)

		tmpl, err := template.ParseFiles(Task004Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_004{PageName:     ReplaceTxtBr(ListTasksAll[7].Name, "; "), 
		                     PageDescribe: ReplaceTxtBr(ListTasksAll[7].Describe, " "),
		                     PageTask:     ListTasksAll[7].Task,
							 TaskCode:     "",
							 ValuesView:   auto.UnitTypeArr}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	} 

}



 