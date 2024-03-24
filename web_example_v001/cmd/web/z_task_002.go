package cmd
import (
	"fmt"
	"net/http"
	"text/template"
	"encoding/json"
	calc "tracker/cmd/task/calc"
)

type TypeTask_002 struct {
	PageName         string
	PageDescribe     string
	PageTask         string
	TaskCode         string
	Values []        string
}

func processHandler_task_002_02_calc(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

    if r.Method == http.MethodPost {
        r.ParseForm()
        selection := r.Form.Get("selection")
        input1 := r.Form.Get("input1")
		input2 := r.Form.Get("input2")

		response:=""
		var number1, number2, result float64

		number1, err := calc.GetNumber(input1)
		if err != nil {
			response = fmt.Sprint(err)	
		}
		if response == ""{
			number2, err = calc.GetNumber(input2)
			if err != nil {
				response = fmt.Sprint(err)	
			}
	    }
		if response == ""{
			
			calculator := calc.NewCalculator()
			result, err = calculator.CalculateTwoArg(number1,number2,selection)
			
			if err != nil {
				response = fmt.Sprint(err)	
			}else{
				response = fmt.Sprintf("%v", result)
			}
	    }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(input1 + " " + selection + " " + input2 + " = " + response)
    }
}

func processHandler_task_002_03_calc(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

    if r.Method == http.MethodPost {
        r.ParseForm()
        input := r.Form.Get("input")

		response:=""

		calculator := calc.NewCalculator()
		result, err := calculator.CalculateExpression(input)

		if err != nil {
			response = fmt.Sprint(err)
		} else {
			response = input + " = " + fmt.Sprintf("%v", result)
		}

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func ServeHTTP_task_002(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)
		
	if (r.URL.Path == Task_002_01_URL) {
		tmpl, err := template.ParseFiles(Task_002_01_Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataCode, err := LoadCodeData(Task_002_01_Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_002{PageName:     ReplaceTxtBr(ListTasksAll[3].Name, "; "), 
		                        PageDescribe: ReplaceTxtBr(ListTasksAll[3].Describe, " "),
		                        PageTask:     ListTasksAll[3].Task,
							    TaskCode:     dataCode,
								Values:       []string {}}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	} 

	if (r.URL.Path == Task_002_02_URL) {
		tmpl, err := template.ParseFiles(Task_002_02_Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_002{PageName:     ReplaceTxtBr(ListTasksAll[4].Name, "; "), 
		                        PageDescribe: ReplaceTxtBr(ListTasksAll[4].Describe, " "),
		                        PageTask:     ListTasksAll[4].Task, 
								TaskCode:     "",
								Values:       []string {"+", "-", "/", "*"}}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	}

	if (r.URL.Path == Task_002_03_URL) {
		tmpl, err := template.ParseFiles(Task_002_03_Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_002{PageName:     ReplaceTxtBr(ListTasksAll[5].Name, "; "), 
		                        PageDescribe: ReplaceTxtBr(ListTasksAll[5].Describe, " "),
		                        PageTask:     ListTasksAll[5].Task, 
								TaskCode:     "",
								Values:       []string {}}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	}

 }






 