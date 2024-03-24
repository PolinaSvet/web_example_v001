package cmd

import (
	"encoding/json"
	"net/http"
	"text/template"
	"tracker/cmd/task/graphBfs"
	"tracker/cmd/task/graphDfs"
	"tracker/cmd/task/treeCustom"
)

type TypeTask007 struct {
	PageName     string
	PageDescribe string
	PageTask     string
	TaskCode     string
	//SvgTreeView      string
}

type TypeTask007tree struct {
	ValueSvg string
}

var rootNode *treeCustom.Node
var graphBfsMap graphBfs.Graph
var graphDfsMap graphDfs.Graph

func processHandlerTask007treeCreate(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.Method == http.MethodPost {
		type RequestData struct {
			ListInputbox int `json:"listInputbox"`
		}

		var requestData RequestData

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ""
		svgTree := ""
		rootNode, svgTree = treeCustom.CreateElementRenderTreeSVG(40)

		tmpl, err := template.ParseFiles(Task007Tree, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask007tree{ValueSvg: svgTree}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := resp
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func processHandlerTask007treeAction(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.Method == http.MethodPost {
		type RequestData struct {
			ListInputbox int `json:"listInputbox"`
			ListAction   int `json:"listAction"`
		}

		var requestData RequestData

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ""
		svgTree := ""
		switch requestData.ListAction {
		case 1:
			rootNode, svgTree = treeCustom.InsertElementRenderTreeSVG(rootNode, requestData.ListInputbox)
		case 2:
			rootNode, svgTree = treeCustom.DeleteElementRenderTreeSVG(rootNode, requestData.ListInputbox)
		case 3:
			rootNode, svgTree = treeCustom.FindElementRenderTreeSVG(rootNode, requestData.ListInputbox)
		case 4:
			rootNode, svgTree = treeCustom.PrintElementRenderTreeSVG(rootNode)
		}

		tmpl, err := template.ParseFiles(Task007Tree, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask007tree{ValueSvg: svgTree}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := resp
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func processHandlerTask007graphBfsAction(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.Method == http.MethodPost {
		type RequestData struct {
			NumElementGraphBfsAVal int `json:"numElementGraphBfsAVal"`
			NumElementGraphBfsBVal int `json:"numElementGraphBfsBVal"`
			ActionVal              int `json:"actionVal"`
		}

		var requestData RequestData

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ""
		svgTree := ""
		switch requestData.ActionVal {
		case 1:
			graphBfsMap, svgTree = graphBfs.CreateElementRenderGraphSVG(10)
		case 2:
			svgTree = graphBfs.FindElementRenderGraphSVG(graphBfsMap, requestData.NumElementGraphBfsAVal, requestData.NumElementGraphBfsBVal)
		}

		tmpl, err := template.ParseFiles(Task007Tree, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask007tree{ValueSvg: svgTree}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := resp
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func processHandlerTask007graphDfsAction(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.Method == http.MethodPost {
		type RequestData struct {
			NumElementGraphDfsAVal int `json:"numElementGraphDfsAVal"`
			NumElementGraphDfsBVal int `json:"numElementGraphDfsBVal"`
			ActionVal              int `json:"actionVal"`
		}

		var requestData RequestData

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ""
		svgTree := ""
		switch requestData.ActionVal {
		case 1:
			graphDfsMap, svgTree = graphDfs.CreateElementRenderGraphSVG(10)
		case 2:
			svgTree = graphDfs.FindElementRenderGraphSVG(graphDfsMap, requestData.NumElementGraphDfsAVal, requestData.NumElementGraphDfsBVal)
		}

		tmpl, err := template.ParseFiles(Task007Tree, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask007tree{ValueSvg: svgTree}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := resp
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func ServeHTTPTask007(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w, r)

	if r.URL.Path == Task007URL {

		tmpl, err := template.ParseFiles(Task007Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask007{PageName: ReplaceTxtBr(ListTasksAll[10].Name, "; "),
			PageDescribe: ReplaceTxtBr(ListTasksAll[10].Describe, " "),
			PageTask:     ListTasksAll[10].Task,
			TaskCode:     ""}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

}
