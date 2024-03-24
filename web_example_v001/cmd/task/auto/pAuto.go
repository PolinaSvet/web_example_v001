package auto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	aBMWPhoto           string = "/ui/static/img/auto/bmw.jpg"
	aMercedesPhoto      string = "/ui/static/img/auto/mercedes.jpg"
	aDodgePhoto         string = "/ui/static/img/auto/dodge.jpg"

	aBMWTypeStr         string = "BMW"
	aMercedesTypeStr    string = "Mercedes"
	aDodgeTypeStr       string = "Dodge"

)

type nameFuncFunc func(string, string, float64, float64, float64, int, int) Auto

var nameFuncCar = []string{aBMWTypeStr, aMercedesTypeStr, aDodgeTypeStr}

// Создаем словарь функций, где ключами являются их имена
var	functions = map[string]nameFuncFunc{
		aBMWTypeStr:      NewAutoBMW,
		aMercedesTypeStr: NewAutoMercedes,
		aDodgeTypeStr:    NewAutoDodge,
}

type UnitType string

const (
  Inch    UnitType = "inch"
  CM      UnitType = "cm"
  Def     UnitType = "def"
)

var UnitTypeArr = []string{
	string(Def),
	string(Inch),
	string(CM),
}

type Unit struct {
  Value float64
  T     UnitType
}

func dimensionsToString(d Dimensions, typeView UnitType) string {
	length  := d.Length().Get(typeView)
	width   := d.Width().Get(typeView)
	height  := d.Height().Get(typeView)
	unit    := fmt.Sprintf("%v", typeView)
	if typeView == Def {
        unit = d.TypeUint()
	}
   
	return fmt.Sprintf("%.2f x %.2f x %.2f %v", length, width, height, unit)
}


func (u Unit) Get(t UnitType) (float64) {
	value := u.Value

	if t != u.T {
		// выполнить конвертирование значения в заданный тип
		switch t {
			case Inch:
				value *= 0.393701 // 1 см = 0.393701 дюйма
			case CM:
				value *= 2.54 // 1 дюйм = 2.54 см
		}
	}
 	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
	TypeUint() string
}
   
type InchesDimensions struct {
	LengthInches float64
	WidthInches  float64
	HeightInches float64
}
   
func (d InchesDimensions) Length() Unit {
	return Unit{
		Value: d.LengthInches,
		T:     Inch,
	}
}
   
func (d InchesDimensions) Width() Unit {
	return Unit{
		Value: d.WidthInches,
		T:     Inch,
	}
}
   
func (d InchesDimensions) Height() Unit {
	return Unit{
		Value: d.HeightInches,
		T:     Inch,
	}
}

func (d InchesDimensions) TypeUint() string {
	return string(Inch) 
}
   
type CMDimensions struct {
	LengthCM float64
	WidthCM  float64
	HeightCM float64
}
   
func (d CMDimensions) Length() Unit {
	return Unit{
		Value: d.LengthCM,
		T:     CM,
	}
}
   
func (d CMDimensions) Width() Unit {
	return Unit{
		Value: d.WidthCM,
		T:     CM,
	}
}
   
func (d CMDimensions) Height() Unit {
	return Unit{
		Value: d.HeightCM,
		T:     CM,
	}
}

func (d CMDimensions) TypeUint() string {
	return string(CM)
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
	Photo() string
}
   
type BMW struct {
	BrandVal      string
	ModelVal      string
	DimensionsVal Dimensions
	MaxSpeedVal   int
	EnginePowerVal int
}
   
func (b BMW) Brand() string {
	return b.BrandVal
}
   
func (b BMW) Model() string {
	return b.ModelVal
}
   
func (b BMW) Dimensions() Dimensions {
	return b.DimensionsVal
}
   
func (b BMW) MaxSpeed() int {
	return b.MaxSpeedVal
}
   
func (b BMW) EnginePower() int {
	return b.EnginePowerVal
}

func (a BMW) Photo() string {
	return aBMWPhoto
}

// создает новый экземпляр структуры 
func NewAutoBMW(brandPar, modelPar string, lengthPar, widthPar, heightPar float64, maxSpeedPar, enginePowerPar int) Auto {
	dimensionsCMDPar := CMDimensions{
		LengthCM: lengthPar,
		WidthCM:  widthPar,
		HeightCM: heightPar,
	}
	return BMW{
		BrandVal: brandPar,
		ModelVal: modelPar,
		DimensionsVal:  dimensionsCMDPar,
		MaxSpeedVal: maxSpeedPar,
		EnginePowerVal: enginePowerPar,
	}
}
   
type Mercedes struct {
	BrandVal      string
	ModelVal      string
	DimensionsVal Dimensions
	MaxSpeedVal   int
	EnginePowerVal int
}
   
func (m Mercedes) Brand() string {
	return m.BrandVal
}
   
func (m Mercedes) Model() string {
	return m.ModelVal
}
   
func (m Mercedes) Dimensions() Dimensions {
	return m.DimensionsVal
}
   
func (m Mercedes) MaxSpeed() int {
	return m.MaxSpeedVal
}
   
func (m Mercedes) EnginePower() int {
	return m.EnginePowerVal
}

func (a Mercedes) Photo() string {
	return aMercedesPhoto
}

// создает новый экземпляр структуры 
func NewAutoMercedes(brandPar, modelPar string, lengthPar, widthPar, heightPar float64, maxSpeedPar, enginePowerPar int) Auto {
	dimensionsCMDPar := CMDimensions{
		LengthCM: lengthPar,
		WidthCM:  widthPar,
		HeightCM: heightPar,
	}
	return Mercedes{
		BrandVal: brandPar,
		ModelVal: modelPar,
		DimensionsVal:  dimensionsCMDPar,
		MaxSpeedVal: maxSpeedPar,
		EnginePowerVal: enginePowerPar,
	}
}
   
type Dodge struct {
	BrandVal      string
	ModelVal      string
	DimensionsVal InchesDimensions
	MaxSpeedVal   int
	EnginePowerVal int
}
   
func (d Dodge) Brand() string {
	return d.BrandVal
}
   
func (d Dodge) Model() string {
	return d.ModelVal
}
   
func (d Dodge) Dimensions() Dimensions {
	return d.DimensionsVal
}
   
func (d Dodge) MaxSpeed() int {
	return d.MaxSpeedVal
}
   
func (d Dodge) EnginePower() int {
	return d.EnginePowerVal
}

func (a Dodge) Photo() string {
	return aDodgePhoto
}

// создает новый экземпляр структуры 
func NewAutoDodge(brandPar, modelPar string, lengthPar, widthPar, heightPar float64, maxSpeedPar, enginePowerPar int) Auto {
	dimensionsCMDPar := InchesDimensions{
		LengthInches: lengthPar,
		WidthInches:  widthPar,
		HeightInches: heightPar,
	}
	return Dodge{
		BrandVal: brandPar,
		ModelVal: modelPar,
		DimensionsVal:  dimensionsCMDPar,
		MaxSpeedVal: maxSpeedPar,
		EnginePowerVal: enginePowerPar,
	}
}

// создает новый экземпляр структуры 
func newAuto(brandPar, modelPar string, lengthPar, widthPar, heightPar float64, maxSpeedPar, enginePowerPar int)(Auto, error) {
	return executeFunction(brandPar, nameFuncCar, functions, brandPar, modelPar, lengthPar, widthPar, heightPar, maxSpeedPar, enginePowerPar)
}

func executeFunction(name string, nameFuncCar []string, functions map[string]nameFuncFunc, brandPar, modelPar string, lengthPar, widthPar, heightPar float64, maxSpeedPar, enginePowerPar int)(Auto, error) {
	// Проверяем, есть ли имя функции в массиве
	if fn, ok := functions[name]; ok {
	    return fn(brandPar, modelPar, lengthPar, widthPar, heightPar, maxSpeedPar, enginePowerPar),nil // Вызываем функцию с переданным аргументом
	} else {
		return nil, fmt.Errorf("бренд '%s' не реализован. Доступны следующие бренды: %s", brandPar, strings.Join(nameFuncCar, ", ")) 
	}
}

// структура для json
type JsonAuto struct {
	Brand       string     `json:"brand"`
	Model       string     `json:"model"`
	Length      float64    `json:"length"`
	Width       float64    `json:"width"`
	Height      float64    `json:"height"`
	MaxSpeed    int        `json:"maxSpeed"`
	EnginePower int        `json:"enginePower"`
}

type ItemsHtml struct {
	Brand       string    
	Model       string     
	Dimensions  string    
	MaxSpeed    string     
	EnginePower string    
	Photo       string             
}

func LoadAutoFromJSONFile(filePath string, typeView UnitType) ([]ItemsHtml, error) {
	jsonAuto:= []JsonAuto{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {	 return nil, err }

	err = json.Unmarshal(data, &jsonAuto)
	if err != nil {	 return nil, err }

	listItemsHtml := make([]ItemsHtml, 0)
	for _, t := range jsonAuto {
		
		auto_, err := newAuto(t.Brand, t.Model, t.Length , t.Width , t.Height , t.MaxSpeed , t.EnginePower)
		if err != nil {	 return nil, err }

		listItemsHtml = append(listItemsHtml,getListItemsHtmlByAuto(auto_, typeView))

	}

	return listItemsHtml, nil
}

func getListItemsHtmlByAuto(auto_ Auto, typeView UnitType) (ItemsHtml) {
	var listItemsHtml ItemsHtml

	listItemsHtml.Brand       = "Brand: " + auto_.Brand()
	listItemsHtml.Model       = "Model: " + auto_.Model()
	listItemsHtml.Dimensions  = "Dimensions: " + dimensionsToString(auto_.Dimensions(), typeView)
	listItemsHtml.MaxSpeed    = "MaxSpeed: " + strconv.Itoa(auto_.MaxSpeed())
	listItemsHtml.EnginePower = "EnginePower: " + strconv.Itoa(auto_.EnginePower())
	listItemsHtml.Photo       = auto_.Photo()
   
	return listItemsHtml
}




