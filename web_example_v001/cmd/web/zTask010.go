package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"tracker/cmd/task/chTask1"
	"tracker/cmd/task/chTask2"
	"tracker/cmd/task/chTask3"
	"tracker/cmd/task/chTask4"
	"tracker/cmd/task/chTask5"
	"tracker/cmd/task/chTask6"
)

type TypeTask010 struct {
	PageName     string
	PageDescribe string
	PageTask     string
	TaskCode     string
}

func processHandlerTask010Start(w http.ResponseWriter, r *http.Request) {
	channelNum := r.URL.Query().Get("channel")
	par1Str := r.URL.Query().Get("par1")
	par2Str := r.URL.Query().Get("par2")
	var par1, par2 int

	if i, err := strconv.Atoi(par1Str); err == nil {
		par1 = i
	}
	if i, err := strconv.Atoi(par2Str); err == nil {
		par2 = i
	}

	if *flagTask10 == false {
		*flagTask10 = true
	} else {
		*flagTask10 = false
	}

	switch channelNum {
	case "1":
		go chTask1.StartTask(channelTask10)
	case "2":
		go chTask2.StartTask(channelTask10)
	case "3":
		go chTask3.StartTask(channelTask10, flagTask10)
	case "4":
		go chTask4.StartTask(channelTask10, flagTask10)
	case "5":
		go chTask5.StartTask(channelTask10)
	case "6":
		go chTask6.StartTask(channelTask10, par1, par2) //100,5
	}

	fmt.Printf("Start %v %v %v: \n", channelNum, par1, par2)
}

func processHandlerTask010Stop(w http.ResponseWriter, r *http.Request) {
	channelNum := r.URL.Query().Get("channel")
	if channelNum == "1" {
		stopChannelTask10 <- true
	}
	fmt.Printf("Stop %v: \n", channelNum)
}

func processHandlerTask010RefreshData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == Task010URLData {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for {
			select {
			case <-r.Context().Done():
				fmt.Println("Client closed the connection")
				return
			case msg := <-channelTask10:
				fmt.Fprintf(w, "data: %s\n\n", msg)
				w.(http.Flusher).Flush()
			case <-stopChannelTask10:
				fmt.Println("Client closed the connection")
				return
			}
		}
	}
}

func ServeHTTPTask010(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == Task010URL {

		defer func() {
			*flagTask10 = true
		}()

		tmpl, err := template.ParseFiles(Task010Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask010{PageName: ReplaceTxtBr(ListTasksAll[13].Name, "; "),
			PageDescribe: ReplaceTxtBr(ListTasksAll[13].Describe, " "),
			PageTask:     ListTasksAll[13].Task,
			TaskCode:     ""}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

}
