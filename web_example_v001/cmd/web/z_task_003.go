package cmd
import (
	"fmt"
	"net/http"
	"text/template"
	"encoding/json"
	electronic "tracker/cmd/task/electronic"
)

type TypeTask_003 struct {
	PageName         string
	PageDescribe     string
	PageTask         string
	TaskCode         string
	ValuesView       []string
}

type TypeTask_003_list struct {
	Values           []electronic.ItemsHtml
}



func processHandler_task_003_01_calc(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

    if r.Method == http.MethodPost {
        r.ParseForm()
        el_viewPhone := r.Form.Get("viewPhone")
		el_brand := r.Form.Get("brand")
		el_model := r.Form.Get("model")
		el_buttons := r.Form.Get("buttons")

		response:="Телефон добавлен"
		err := electronic.NewPhoneByTypeStr(el_viewPhone, el_brand, el_model, el_buttons, PhoneInfo_json) 
		if err != nil {	 
			response = fmt.Sprint(err)	 
		}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode("")
		
		tmpl, err := template.ParseFiles(Task_003_01_PList, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataCode, err := electronic.LoadPhoneFromJSONFile(PhoneInfo_json)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_003_list{ Values:  dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(response)
    }
}

func processHandler_task_003_01_show(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)
	
    if r.Method == http.MethodPost {
        
		r.ParseForm()
		
		tmpl, err := template.ParseFiles(Task_003_01_PList, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataCode, err := electronic.LoadPhoneFromJSONFile(PhoneInfo_json)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_003_list{ Values:  dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response:="Каталог обнавлен"
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func ServeHTTP_task_003(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

	if (r.URL.Path == Task_003_01_URL) {

		tmpl, err := template.ParseFiles(Task_003_01_Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_003{PageName:     ReplaceTxtBr(ListTasksAll[6].Name, "; "), 
		                     PageDescribe: ReplaceTxtBr(ListTasksAll[6].Describe, " "),
		                     PageTask:     ListTasksAll[6].Task,
							 TaskCode:     "",
							 ValuesView:   electronic.ViewPhoneArr}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	} 

}


 