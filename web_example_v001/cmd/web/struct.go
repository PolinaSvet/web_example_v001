package cmd

import "net/http"

type NeuteredFileSystem struct {
	fs http.FileSystem
}

type Err struct {
	Status int
	Text   string
}

type ListTasks struct {
	Image          string              
	Name           string              
	Describe       string
	Link           string 
	PageHTML       string
	Task           string          
}

type httpHandler struct{
    message string
    pageName string
}

var (
	ListTasksAll     []ListTasks
)

