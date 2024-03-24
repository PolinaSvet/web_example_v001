package cmd
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
	"strings"
)

const (
    january = iota + 1
    february
    march
    april
    may
    june
	july
	august
	september
	october
	november
	december
)

type Month struct {
	Number           int
	NameEn           string
	NameRu           string
	Image            string
	Describe         string
   }
   
type TypeMonthsFull struct {
	Months           []Month
   }

type TypeTask_001_01 struct {
	PageName         string
	PageDescribe     string
	PageTask         string
	TaskCode         string
   }

type TypeTask_001_02 struct {
	Values []        string
	Selected         string
	SelectedNameRu   string
	SelectedImage    string
	SelectedDescribe string
	PageName         string
	PageDescribe     string
	PageTask         string
   }

type TypeTask_001_03 struct {
	PageName         string
	PageDescribe     string
	PageTask         string
	Months           []Month
   }

var (
	monthsName = []string {"all", "january", "february", "march","april","may","june","july","august","september","october","november","december"}
)


func ServeHTTP_task_001(w http.ResponseWriter, r *http.Request) {
	defer handlePanic(w,r)

	dataInfo, err := loadData(MonthsInfo_json)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	if (r.URL.Path == Task_001_01_URL) {
		tmpl, err := template.ParseFiles(Task_001_01_Page, TemplatePage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		dataCode, err := LoadCodeData(Task_001_01_Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := TypeTask_001_01{PageName:     ReplaceTxtBr(ListTasksAll[0].Name, "; "), 
		                        PageDescribe: ReplaceTxtBr(ListTasksAll[0].Describe, " "),
		                        PageTask:     ListTasksAll[0].Task,
								TaskCode:     dataCode}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		return
	} 

	if (r.Method == "GET") && (r.URL.Path == Task_001_02_URL) {
		data := TypeTask_001_02{Values: monthsName}
		q_select := r.URL.Query().Get("selected")
		data = GetInfoMonthShort(q_select, dataInfo )
		
		tmpl := template.Must(template.ParseFiles(Task_001_02_Page, TemplatePage))
		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} 

	if (r.Method == "GET") && (r.URL.Path == Task_001_03_URL) {
		tmpl := template.Must(template.ParseFiles(Task_001_03_Page, TemplatePage))
		data := TypeTask_001_03{Months:       dataInfo.Months, 
			                    PageName:     ReplaceTxtBr(ListTasksAll[2].Name, "; "), 
								PageDescribe: ReplaceTxtBr(ListTasksAll[2].Describe, " "),
								PageTask:     ListTasksAll[2].Task}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}    

 }

 func ReplaceTxtBr(txt_ string, txtReplace_ string) string {
 	return strings.Replace(txt_, "<br/>", txtReplace_, -1)
 }

func GetAllMonthName() string {
	strValue:=""
	
	strValue =  fmt.Sprintf("Месяц № %d это %s;<br/>", january,   monthsName[january]) + 
			    fmt.Sprintf("Месяц № %d это %s;<br/>", february,  monthsName[february]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", march,     monthsName[march]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", april,     monthsName[april]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", may,       monthsName[may]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", june,      monthsName[june]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", july,      monthsName[july]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", august,    monthsName[august]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", september, monthsName[september]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", october,   monthsName[october]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", november,  monthsName[november]) + 
				fmt.Sprintf("Месяц № %d это %s;<br/>", december,  monthsName[december])

	return strValue
} 


func GetInfoMonthShort(q_select_ string, dataInfo_ TypeMonthsFull) (TypeTask_001_02) {
	data := TypeTask_001_02{Values:       monthsName,
		                    PageName:     ReplaceTxtBr(ListTasksAll[1].Name, "; "), 
		                    PageDescribe: ReplaceTxtBr(ListTasksAll[1].Describe, " "),
		                    PageTask:     ListTasksAll[1].Task}

	switch(q_select_) {
		case "all": 
			data.Selected         = q_select_
			data.SelectedNameRu   = q_select_
			data.SelectedImage    = MonthsInfo_image
			data.SelectedDescribe = GetAllMonthName()

		case "january", "february", "march","april","may","june","july","august","september","october","november","december": 
			dataMonth := GetInfoMonth(q_select_, dataInfo_) 
			data.Selected         = dataMonth.NameEn
			data.SelectedNameRu   = dataMonth.NameRu
			data.SelectedImage    = dataMonth.Image
			data.SelectedDescribe = dataMonth.Describe

		default: 
			data.Selected         = q_select_
			data.SelectedNameRu   = q_select_
			data.SelectedImage    = MonthsInfo_image
			data.SelectedDescribe = q_select_
	}

	return data
}

func GetInfoMonth(monthname_ string, dataMonths_ TypeMonthsFull) (Month) {
	data := Month{Number: 0, NameEn: "Error", NameRu: "Error", Image: MonthsInfo_image, Describe: "Error"}

	for index, _ := range dataMonths_.Months{
		if dataMonths_.Months[index].NameEn == monthname_{
			data = dataMonths_.Months[index]
			break
		}
	}

	return data
} 


func loadData(filename_ string) (TypeMonthsFull, error) {
	dataBytes, err := ioutil.ReadFile(filename_)
	if err != nil {
	 return TypeMonthsFull{}, fmt.Errorf("failed to read JSON file: %v", err)
	}
   
	var data TypeMonthsFull
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
	 return TypeMonthsFull{}, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}
   
	return data, nil
   }

   
func LoadCodeData(filename_ string) (string, error) {
	dataBytes, err := ioutil.ReadFile(filename_)
	if err != nil {
	 return "", fmt.Errorf("failed to read JSON file: %v", err)
	}
  
	return string(dataBytes), nil
}
