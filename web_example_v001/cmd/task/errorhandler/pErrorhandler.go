package errorhandler

import (
 "encoding/json"
 "time"
 "fmt"
 "io/ioutil"
 "runtime"
 "sort"
)

const (
	eTypePanic string = "Panic"
	eTypeError string = "Error"

	coutMess = 100
	ErrorLogJson      string = "ui/static/json/errorLog.json"
)
var TypeError = []string{eTypePanic, eTypeError}

type ErrorResponse struct {
 Timestamp time.Time `json:"timestamp"`
 Type      string    `json:"type"`
 Message   string    `json:"message"`
}

func logAndSaveError(errorLog error,filename string, typeError string) error {
   
	jsonError:= []ErrorResponse{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {	 
		return fmt.Errorf(fmt.Sprint(err)) 
	}
   
	err = json.Unmarshal(data, &jsonError)
	if err != nil {	 
		return fmt.Errorf(fmt.Sprint(err))
	}
   
	errResponse := ErrorResponse{
		Timestamp: time.Now(),
		Type:      typeError,
		Message:   errorLog.Error(),
	}
   
	jsonError = append(jsonError, errResponse)
   
	if len(jsonError) > coutMess {
		jsonError = jsonError[len(jsonError)-coutMess:] 
	}
   
	rawDataOut, err := json.MarshalIndent(&jsonError, "", "  ")
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err))	 
	}
   
	err = ioutil.WriteFile(filename, rawDataOut, 0)
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err))	 
	}
   
	return nil
}
   
func PrintError(errLog error) error {
	if pc, filename, _, ok := runtime.Caller(1); ok {
		fname := runtime.FuncForPC(pc).Name()
		//filename := path.Base(fname)
		_ = logAndSaveError(fmt.Errorf("Ошибка возникла в файле %s, в функции %s\n: %v ", filename, fname, errLog), ErrorLogJson, eTypeError)
	}
	return errLog
}

func PrintErrorPanic(errLog error) {
	if pc, filename, _, ok := runtime.Caller(1); ok {
		fname := runtime.FuncForPC(pc).Name()
		//filename := path.Base(fname)
		_ = logAndSaveError(fmt.Errorf("Ошибка возникла в файле %s, в функции %s\n: %v ", filename, fname, errLog), ErrorLogJson, eTypePanic)
	}
}


func LoadErrorFromJSONFile() ([]ErrorResponse, error) {
	val:= []ErrorResponse{}
	data, err := ioutil.ReadFile(ErrorLogJson)
	if err != nil {	 return nil, err }

	err = json.Unmarshal(data, &val)
	if err != nil {	 return nil, err }

	// Отсортировать записи в порядке возрастания Timestamp
	//sort.Slice(val, func(i, j int) bool {
	//	return val[i].Timestamp.Before(val[j].Timestamp)
	//})

	// Отсортировать записи в порядке убывания Timestamp
	sort.Slice(val, func(i, j int) bool {
		return val[i].Timestamp.After(val[j].Timestamp)
	})

	return val, nil
}
