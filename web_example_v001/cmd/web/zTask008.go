package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"
	"tracker/cmd/task/mapEqualFind"
	"tracker/cmd/task/mapEqualSave"
)

type TypeTask008 struct {
	PageName     string
	PageDescribe string
	PageTask     string
	TaskCode     string
}

var memCache *mapEqualSave.InMemoryCache

func processHandlerTask008mapRefreshData(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == Task008URLData {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for {

			select {
			case <-r.Context().Done():
				fmt.Println("Client closed the connection")
				return
			default:
				if memCache != nil {
					if memCache.IsCachedEntries() {
						data := memCache.PrintCachedEntriesStr()
						fmt.Fprintf(w, "data: %s\n\n", data)
						w.(http.Flusher).Flush()
					} else {
						data := time.Now().Format(time.Stamp) + ": хэш пуст, нет записей для хранения !!!"
						fmt.Fprintf(w, "data: %s\n\n", data)
						w.(http.Flusher).Flush()
						return
					}
				}
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func processHandlerTask008mapEqualAction(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.Method == http.MethodPost {
		type RequestData struct {
			InputString1Val string `json:"inputString1Val"`
			InputString2Val string `json:"inputString2Val"`
			ActionVal       int    `json:"actionVal"`
		}

		var requestData RequestData

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ""
		respHash := ""

		switch requestData.ActionVal {
		case 1:
			resp, respHash = mapEqualFind.FindElements(requestData.InputString1Val, requestData.InputString2Val)
			if memCache == nil {
				memCache = mapEqualSave.NewInMemoryCache(30 * time.Second)
				go memCache.StartDeletionCycle()
			}
			memCache.Set(respHash, resp)
		}

		response := resp
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func ServeHTTPTask008(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.URL.Path == Task008URL {

		tmpl, err := template.ParseFiles(Task008Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask007{PageName: ReplaceTxtBr(ListTasksAll[11].Name, "; "),
			PageDescribe: ReplaceTxtBr(ListTasksAll[11].Describe, " "),
			PageTask:     ListTasksAll[11].Task,
			TaskCode:     ""}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

}
