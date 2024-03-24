package cmd

import (
	"encoding/json"
	"net/http"
	"text/template"
	"tracker/cmd/task/bankClient"
)

type TypeTask009 struct {
	PageName     string
	PageDescribe string
	PageTask     string
	TaskCode     string
}

type TypeTask_009_test struct {
	Values bankClient.BankClientImpl
	Mess   string
}

func processHandlerTask009Action(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.Method == http.MethodPost {
		type RequestData struct {
			InputAmount int    `json:"inputAmount"`
			ActionVal   string `json:"actionVal"`
		}

		var requestData RequestData

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bankClientData, bankClientMess := bankClient.BankClientOperation(requestData.ActionVal, requestData.InputAmount)

		tmpl, err := template.ParseFiles(Task009PList, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_009_test{Values: bankClientData,
			Mess: bankClientMess}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := ""
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func ServeHTTPTask009(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.URL.Path == Task009URL {

		tmpl, err := template.ParseFiles(Task009Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask009{PageName: ReplaceTxtBr(ListTasksAll[12].Name, "; "),
			PageDescribe: ReplaceTxtBr(ListTasksAll[12].Describe, " "),
			PageTask:     ListTasksAll[12].Task,
			TaskCode:     ""}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

}
