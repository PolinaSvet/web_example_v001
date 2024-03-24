package mapEqualFind

import (
	"crypto/sha1"
	"fmt"
	"strings"
	"time"
)

func splitStr(value string) []string {

	parts := strings.Fields(value)

	var result []string
	for _, part := range parts {
		if len(part) > 0 {
			result = append(result, part)
		}
	}

	return result
}

func findCommonElements(str1 string, str2 string) map[string]int {
	commonElements := make(map[string]int)

	arr1 := splitStr(str1)
	arr2 := splitStr(str2)

	// Создаем map для первого массива, где ключи - элементы массива
	// и значения - количество повторений каждого элемента
	countMap := make(map[string]int)
	for _, elem := range arr1 {
		countMap[elem]++
	}

	// Проверяем второй массив на наличие общих элементов
	for _, elem := range arr2 {
		if countMap[elem] > 0 {
			commonElements[elem]++
		}
	}

	//fmt.Printf("%v: %v: %v: \n", arr1, arr2, commonElements)

	return commonElements
}

func getInformation(value map[string]int) (string, string) {
	var keys string
	for key := range value {
		keys += key + ","
	}

	// Генерация уникального хеша на основе времени и ключей
	currentTime := time.Now().String()
	hash := sha1.New()
	hash.Write([]byte(currentTime + keys))
	hashResult := hash.Sum(nil)

	return keys, fmt.Sprintf("%x", hashResult)
}

func FindElements(str1 string, str2 string) (string, string) {

	mapFind := findCommonElements(str1, str2)

	commonElements, commonElementsHash := getInformation(mapFind)
	//fmt.Printf("%v: %v: %v: \n", mapFind, commonElements, commonElementsHash)

	return commonElements, commonElementsHash
}
