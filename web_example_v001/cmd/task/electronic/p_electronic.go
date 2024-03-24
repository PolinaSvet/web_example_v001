package electronic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

const (
	ApplePhonePhoto    string = "/ui/static/img/phone/applePhone.jpg"
	AndroidPhonePhoto  string = "/ui/static/img/phone/androidPhone.jpg"
 	RadioPhonePhoto    string = "/ui/static/img/phone/radioPhone.jpg"

	ApplePhoneType     string = "smartphone"
	AndroidPhoneType   string = "smartphone"
	RadioPhoneType     string = "station"

	ApplePhoneOS       string = "iOS"
	AndroidPhoneOS     string = "Android"
	RadioPhoneOS       string = ""

	ApplePhoneBrand    string = "apple"
	AndroidPhoneBrand  string = ""
	RadioPhoneBrand    string = ""

	ApplePhone         TypePhone = "applePhone"
	AndroidPhone       TypePhone = "androidPhone"
 	RadioPhone         TypePhone = "radioPhone"

)

var ViewPhoneArr = []string{
	string(AndroidPhone),
	string(ApplePhone),
	string(RadioPhone),
}

// Phone интерфейс представляет общие методы для телефонов
type Phone interface {
	Brand() string
	Model() string
	Type()  string
	Photo() string
}

// StationPhone интерфейс представляет метод для стационарных телефонов
type StationPhone interface {
	ButtonsCount() int
}
   

// Smartphone интерфейс представляет метод для смартфонов
type Smartphone interface {
	OS() string  //название операционной системы
}

// applePhone структура представляет телефоны Apple
type applePhone struct {
	brand   string 
	model   string 
}
   
func (a applePhone) Brand() string {
	return a.brand
}
   
func (a applePhone) Model() string {
	return a.model
}
   
func (a applePhone) Type() string {
	return ApplePhoneType
}
   
func (a applePhone) OS() string {
	return ApplePhoneOS
}

func (a applePhone) Photo() string {
	return ApplePhonePhoto
}
   
// NewApplePhone создает новый экземпляр структуры applePhone
func NewApplePhone(model string) Phone {
	return applePhone{
	 brand: ApplePhoneBrand,
	 model: model,
	}
}
   
// androidPhone структура представляет телефоны Android
type androidPhone struct {
	brand   string 
	model   string 
}
   
func (a androidPhone) Brand() string {
	return a.brand
}
   
func (a androidPhone) Model() string {
	return a.model
}
   
func (a androidPhone) Type() string {
	return AndroidPhoneType
}
   
func (a androidPhone) OS() string {
	return AndroidPhoneOS
}

func (a androidPhone) Photo() string {
	return AndroidPhonePhoto
}
   
// NewAndroidPhone создает новый экземпляр структуры androidPhone
func NewAndroidPhone(brand, model string) Phone {
	return androidPhone{
	 brand: brand,
	 model: model,
	}
}
   
// radioPhone структура представляет радиотелефоны
type radioPhone struct {
	brand   string 
	model   string 
	buttons int    
}
  
func (r radioPhone) Brand() string {
	return r.brand
}
   
func (r radioPhone) Model() string {
	return r.model
}
   
func (r radioPhone) Type() string {
	return RadioPhoneType 
}
   
func (r radioPhone) ButtonsCount() int {
	return r.buttons
}

func (a radioPhone) Photo() string {
	return RadioPhonePhoto
}
   
// NewRadioPhone создает новый экземпляр структуры radioPhone
func NewRadioPhone(brand, model string, buttons int) Phone {
	return radioPhone{
	 brand:   brand,
	 model:   model,
	 buttons: buttons,
	}
}

type TypePhone string

// NewPhoneByType создает новый экземпляр по типу телефона
func NewPhoneByType(typePhone TypePhone, brand, model string, buttons int) (Phone, error) {
	switch typePhone {
	case ApplePhone:
		return NewApplePhone(model), nil
	case AndroidPhone:
		return NewAndroidPhone(brand,model), nil
	case RadioPhone:
		return NewRadioPhone(brand,model,buttons), nil
	default:
	 	return nil, fmt.Errorf("некорректный тип телефона. Возможны следующие типы: applePhone, androidPhone, radioPhone)")
	}
}


//преобразование string to TypePhone
func ParseTypePhone(typeStr string) (TypePhone, error) {
	switch typeStr {
	case string(ApplePhone): 
	 	return ApplePhone, nil
	case string(AndroidPhone): 
	 	return AndroidPhone, nil
	case string(RadioPhone): 
	 	return RadioPhone, nil
	default:
	 	return "", fmt.Errorf("неверное значение типа телефона: %s", typeStr)
	}
}

// структура для json
type JsonPhone struct {
	Brand   string `json:"brand"`
	Model   string `json:"model"`
	Buttons int    `json:"buttons"`
	View    string `json:"view"`
}

type ItemsHtml struct {
	Brand    string              
	Model    string              
	Type     string
	Photo    string 
	OS       string
	Buttons  string          
}

func LoadPhoneFromJSONFile(filePath string) ([]ItemsHtml, error) {
	jsonPhone:= []JsonPhone{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {	 return nil, err }

	err = json.Unmarshal(data, &jsonPhone)
	if err != nil {	 return nil, err }

	listItemsHtml := make([]ItemsHtml, 0)
	for _, t := range jsonPhone {
		phoneType, err := ParseTypePhone(t.View)
		if err != nil {	 return nil, err }
		
		phone_, err := NewPhoneByType(phoneType, t.Brand, t.Model, t.Buttons)
		if err != nil {	 return nil, err }

		listItemsHtml = append(listItemsHtml,GetListItemsHtmlByPhone(phone_))

	}

	return listItemsHtml, nil
}

// Функция для загрузки данных из JSON-файла и создания экземпляров Phone
func GetListItemsHtmlByPhone(phone_ Phone) (ItemsHtml) {
	var listItemsHtml ItemsHtml

	listItemsHtml.Brand = "Brand: " + phone_.Brand()
	listItemsHtml.Model = "Model: " + phone_.Model()
	listItemsHtml.Type  = "Type: "  + phone_.Type()
	listItemsHtml.Photo = phone_.Photo()
	
	if stationPhone, ok := phone_.(StationPhone); ok {
		listItemsHtml.Buttons = "Buttons: " + strconv.Itoa(stationPhone.ButtonsCount()) 
	}
   
	if smartphone, ok := phone_.(Smartphone); ok {
		listItemsHtml.OS = "OS: " + smartphone.OS()
	}
   
	return listItemsHtml
}

func AddPhoneToJSONFile(phone JsonPhone, filename string) error {

	jsonPhone:= []JsonPhone{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {	 
		return fmt.Errorf(fmt.Sprint(err))	  
	}

	err = json.Unmarshal(data, &jsonPhone)
	if err != nil {	 
		return fmt.Errorf(fmt.Sprint(err))	  
	}

	//fmt.Printf("1: %v\n", phone)

    jsonPhone = append(jsonPhone, phone)

    if len(jsonPhone) > 30 {
        jsonPhone = jsonPhone[len(jsonPhone)-30:] // Оставляем только последние 20 записей
    }

	rawDataOut, err := json.MarshalIndent(&jsonPhone, "", "  ")
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err))	 
	}

	err = ioutil.WriteFile(filename, rawDataOut, 0)
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err))	 
	}

    return nil
}


// NewPhoneByType создает новый экземпляр по типу телефона
func NewPhoneByTypeStr(typePhone, brand, model, buttons, filename string) (error) {

	typePhoneVal, err := ParseTypePhone(typePhone) 
	if err != nil {	 
		return fmt.Errorf(fmt.Sprint(err))	 
	}

	buttonsVal:=0
	if (typePhoneVal == RadioPhone){
		buttons_, err := strconv.Atoi(buttons)
		if err != nil {	 
			return fmt.Errorf(fmt.Sprint(err))	 
		}
		if (buttons_ > 0) && (buttons_ < 1000){
			buttonsVal = buttons_
		}else{
			return fmt.Errorf("неверное количество кнопок: %s", buttons)
		}
	}

	if len(model)<1{
		return fmt.Errorf("ошибка ввод модели: %s", model)
	}

	if (len(brand)<1) && ((typePhoneVal == RadioPhone) || (typePhoneVal == AndroidPhone)){
		return fmt.Errorf("ошибка ввода бренда: %s", brand)
	}


	jsonPhoneVal :=JsonPhone {
		Brand:    brand,
		Model:    model,
		Buttons:  buttonsVal,
		View:     typePhone}

	err = AddPhoneToJSONFile(jsonPhoneVal, filename)
	if err != nil {	 
		return fmt.Errorf(fmt.Sprint(err))	 
	}

	return nil	

}

