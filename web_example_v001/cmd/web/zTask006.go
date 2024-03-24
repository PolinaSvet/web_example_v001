package cmd
import (
	//"fmt"
	"net/http"
	"text/template"
	"encoding/json"
	"tracker/cmd/task/sortCustomTypes"
	"tracker/cmd/task/sortCustomTypesTest"
)

type TypeTask006 struct {
	PageName         string
	PageDescribe     string
	PageTask         string
	TaskCode         string
	Values           []string
	ValuesData       []sortCustomTypesTest.ItemNameDataTest
}

type TypeTask_006_list struct {
	Values           []sortCustomTypes.ItemNameData
}

type TypeTask_006_test struct {
	Values           []sortCustomTypesTest.ItemNameDataTest
}

func processHandlerTask006ShowSortArrayTest(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

    if r.Method == http.MethodPost  {
        type RequestData struct {
			ListInputboxTypeSort  int    `json:"listInputboxTypeSort"`
	    }

		var requestData RequestData

	    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	    }

		resp := ""
		dataCode, err := sortCustomTypesTest.SortArrayGetListTest(requestData.ListInputboxTypeSort)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		tmpl, err := template.ParseFiles(Task006PTest, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_006_test{ Values:  dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response:= resp
		w.Header().Set("Content-Type", "application/json")	
		json.NewEncoder(w).Encode(response)

    }
}

func processHandlerTask006ShowSortArray(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

    if r.Method == http.MethodPost  {
        type RequestData struct {
			ListInputboxSize  int    `json:"listInputboxSize"`
			ListInputboxMax   int    `json:"listInputboxMax"`
			ListInputboxAsc   bool   `json:"listInputboxAsc"` 
	    }

		var requestData RequestData

	    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	    }


		resp := ""
		dataCode, err := sortCustomTypes.SortArrayGetList( requestData.ListInputboxMax, requestData.ListInputboxSize, requestData.ListInputboxAsc)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			//resp = fmt.Sprintf("%v", err)
		}
		

		tmpl, err := template.ParseFiles(Task006PList, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_006_list{ Values:  dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response:= resp
		w.Header().Set("Content-Type", "application/json")	
		json.NewEncoder(w).Encode(response)

    }
}




func ServeHTTPTask006(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

	if (r.URL.Path == Task006URL) {

		tmpl, err := template.ParseFiles(Task006Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataCode, err := sortCustomTypesTest.LoadFromJSONFile(SortCustomJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask006{PageName:      ReplaceTxtBr(ListTasksAll[9].Name, "; "), 
		                     PageDescribe: ReplaceTxtBr(ListTasksAll[9].Describe, " "),
		                     PageTask:     ListTasksAll[9].Task,
							 TaskCode:     "",
							 Values:       sortCustomTypesTest.ArrName,
							 ValuesData:   dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	} 

}



 