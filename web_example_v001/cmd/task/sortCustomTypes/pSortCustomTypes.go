package sortCustomTypes

import (
	"fmt"
	"math/rand"
	"strings"
	"tracker/cmd/task/errorhandler"
)


// bubbleSortInterruptAsc(arr []int, ascending bool)        //1) Сортировка пузырьком
// selectionSortBidirectionalAsc(arr []int, ascending bool) //2) Сортировка выбором
// insertionSortAsc(arr []int, ascending bool)              //3) Сортировка вставками
// mergeSortAsc(arr []int, ascending bool)                  //4) Сортировка слиянием
// quickSortAsc(arr []int, ascending bool)                  //5) Быстрая сортировка

type ItemNameData struct {
	Name     string     
	Data     string     
}

//1) алгоритм сортировки пузырьком
//=========================================

//1.1) обычный по возрастанию
func BubbleSort(ar []int) {
	n := len(ar)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            // Если текущий элемент больше следующего, меняем их местами
            if ar[j] > ar[j+1] {
                ar[j], ar[j+1] = ar[j+1], ar[j]
            }
        }
    }
}

//1.2) улучшенный алгоритм пузырьковой сортировки по возрастанию
func BubbleSortInterrupt(ar []int) {
	n := len(ar)
    for i := 0; i < n-1; i++ {
        swapped := false

        for j := 0; j < n-i-1; j++ {
            if ar[j] > ar[j+1] {
                ar[j], ar[j+1] = ar[j+1], ar[j]
                swapped = true
            }
        }

        if !swapped {
            break
        }
    }
}

//1.3) функция сортировки для сортировки элементов в порядке убывания
func BubbleSortInterruptReversed(ar []int) {
	n := len(ar)
    for i := 0; i < n-1; i++ {
        swapped := false

        for j := 0; j < n-i-1; j++ {
            if ar[j] < ar[j+1] {
                ar[j], ar[j+1] = ar[j+1], ar[j]
                swapped = true
            }
        }

        if !swapped {
            break
        }
    }
}

//1.4) функция сортировки для сортировки элементов по возрастанию с использованием рекурсивной функции
func BubbleSortRecursive(ar []int) {
    n := len(ar)

    // Базовый случай, массив уже отсортирован
    if n <= 1 {
        return
    }

    for i := 0; i < n-1; i++ {
        if ar[i] > ar[i+1] {
            ar[i], ar[i+1] = ar[i+1], ar[i]
        }
    }

    // Рекурсивный вызов для сортировки оставшейся части массива
    BubbleSortRecursive(ar[:n-1])
}

//1.5) функция сортировки для сортировки элементов по возрастанию или по убыванию
func BubbleSortInterruptAsc(arr []int, ascending bool) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
		swapped := false
        for j := 0; j < n-i-1; j++ {
            if ascending {
                if arr[j] > arr[j+1] {
                    arr[j], arr[j+1] = arr[j+1], arr[j] // обмен значениями
					swapped = true
                }
            } else {
                if arr[j] < arr[j+1] {
                    arr[j], arr[j+1] = arr[j+1], arr[j] // обмен значениями
					swapped = true
                }
            }
        }
		
		if !swapped {
            break
        }
    }
}

//2) Сортировка выбором
//=========================================

//2.1) Сортировка выбором "слева направо"
func SelectionSort(ar []int) {
    n := len(ar)

    for i := 0; i < n-1; i++ {
        minIndex := i

        for j := i + 1; j < n; j++ {
            if ar[j] < ar[minIndex] {
                minIndex = j
            }
        }

        if minIndex != i {
            ar[i], ar[minIndex] = ar[minIndex], ar[i]
        }
    }
}

//2.2) Сортировка выбором "справа налево"
func SelectionSortByMax(ar []int) {
    n := len(ar)

    for i := n - 1; i > 0; i-- {
        maxIndex := i

        for j := i - 1; j >= 0; j-- {
            if ar[j] > ar[maxIndex] {
                maxIndex = j
            }
        }

        if maxIndex != i {
            ar[i], ar[maxIndex] = ar[maxIndex], ar[i]
        }
    }
}

//2.3) Двунаправленная сортировка выбором
func SelectionSortBidirectional(ar []int) {
    n := len(ar)
    left, right := 0, n-1

    for left < right {
        minIndex, maxIndex := left, right

        for i := left; i <= right; i++ {
            if ar[i] < ar[minIndex] {
                minIndex = i
            }
            if ar[i] > ar[maxIndex] {
                maxIndex = i
            }
        }

        if minIndex != left {
            ar[left], ar[minIndex] = ar[minIndex], ar[left]
        }

        // Обработка случая, когда максимальный элемент
        // был на позиции left и был перемещен на minIndex
        if maxIndex == left {
            maxIndex = minIndex
        }

        if maxIndex != right {
            ar[right], ar[maxIndex] = ar[maxIndex], ar[right]
        }

        left++
        right--
    }
}

//2.4) Двунаправленная сортировка выбором для сортировки элементов по возрастанию или по убыванию
func SelectionSortBidirectionalAsc(ar []int, ascending bool) {
    n := len(ar)
    left, right := 0, n-1

    for left < right {
        minIndex, maxIndex := left, right

        for i := left; i <= right; i++ {
            if ascending && ar[i] < ar[minIndex] {
                minIndex = i
            } else if !ascending && ar[i] > ar[minIndex] {
                minIndex = i
            }
            if ascending && ar[i] > ar[maxIndex] {
                maxIndex = i
            } else if !ascending && ar[i] < ar[maxIndex] {
                maxIndex = i
            }
        }

        if minIndex != left {
            ar[left], ar[minIndex] = ar[minIndex], ar[left]
        }

        // Обработка случая, когда максимальный элемент
        // был на позиции left и был перемещен на minIndex
        if maxIndex == left {
            maxIndex = minIndex
        }

        if maxIndex != right {
            ar[right], ar[maxIndex] = ar[maxIndex], ar[right]
        }

        left++
        right--
    }
}

//3) Сортировка вставками
//=========================================

//3.1)
func InsertionSort(ar []int) {
    n := len(ar)

    for i := 1; i < n; i++ {
        key := ar[i]
        j := i - 1

        for j >= 0 && ar[j] > key {
            ar[j+1] = ar[j]
            j--
        }

        ar[j+1] = key
    }
}

//3.2) сортировка в обе стороны
func InsertionSortAsc(arr []int, ascending bool) {
    n := len(arr)
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1
        if ascending {
            for j >= 0 && arr[j] > key {
                arr[j+1] = arr[j]
                j--
            }
        } else {
            for j >= 0 && arr[j] < key {
                arr[j+1] = arr[j]
                j--
            }
        }
        arr[j+1] = key
    }
}

//4) Сортировка слиянием
//=========================================
//4.1) сортировка
func MergeSort(ar []int) []int {
    if len(ar) <= 1 {
        return ar
    }

    mid := len(ar) / 2
    left := MergeSort(ar[:mid])
    right := MergeSort(ar[mid:])

    return merge(left, right)
}

func merge(left, right []int) []int {
    result := make([]int, 0)

    for len(left) > 0 && len(right) > 0 {
        if left[0] <= right[0] {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }

    result = append(result, left...)
    result = append(result, right...)

    return result
}


//4.2) сортировка в обе стороны
func MergeSortAsc(arr []int, ascending bool) []int {
    n := len(arr)
    if n <= 1 {
        return arr
    }
    mid := n / 2
    left := MergeSortAsc(arr[:mid], ascending)
    right := MergeSortAsc(arr[mid:], ascending)
    return mergeAsc(left, right, ascending)
}

func mergeAsc(left, right []int, ascending bool) []int {
    result := make([]int, 0)
    i, j := 0, 0
    for i < len(left) && j < len(right) {
        if ascending {
            if left[i] < right[j] {
                result = append(result, left[i])
                i++
            } else {
                result = append(result, right[j])
                j++
            }
        } else {
            if left[i] > right[j] {
                result = append(result, left[i])
                i++
            } else {
                result = append(result, right[j])
                j++
            }
        }
    }
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)
    return result
}

//5) Быстрая сортировка
//=========================================

//5.1)
func QuickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar) - 1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	QuickSort(ar[:left])
	QuickSort(ar[left + 1:])

	return
}

//5.2) сортировка в обе стороны

func QuickSortAsc(ar []int, ascending bool) {
	if len(ar) < 2 {
	 return
	}
   
	left, right := 0, len(ar)-1
	pivotIndex := rand.Int() % len(ar)
   
	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]
   
	for i := 0; i < len(ar); i++ {
	 if (ar[i] < ar[right] && ascending) || (ar[i] > ar[right] && !ascending) {
	  ar[i], ar[left] = ar[left], ar[i]
	  left++
	 }
	}
   
	ar[left], ar[right] = ar[right], ar[left]
   
	QuickSortAsc(ar[:left], ascending)
	QuickSortAsc(ar[left+1:], ascending)
   }

func GenerateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}

func arrayToData(arr []int, nameSort string) (ItemNameData) {
	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = fmt.Sprint(num)
	}

	var itemNameDataVar ItemNameData
	itemNameDataVar.Name = nameSort
	itemNameDataVar.Data = strings.Join(strArr, ", ")

	return itemNameDataVar 
}



func SortArrayGetList(max, size int, ascending bool)([]ItemNameData, error) {
	
	if ((max<0)||(max>1000000)){
		return nil, errorhandler.PrintError(fmt.Errorf("max д.б. в пределе от 1 до 1000000"))//fmt.Errorf("max д.б. в пределе от 1 до 1000000")
	}
	if ((size<0)||(size>1000)){
		return nil, errorhandler.PrintError(fmt.Errorf("size д.б. в пределе от 1 до 1000"))//fmt.Errorf("size д.б. в пределе от 1 до 1000")
	}	
	
	var itemText ItemNameData
	listNameData := make([]ItemNameData, 0)
	arrayInit := GenerateSlice(max, size)
	ar := make([]int, len(arrayInit))
	
	listNameData = append(listNameData,arrayToData(arrayInit,"Исходный массив:"))
	
	itemText.Name = "Результату сортировки:"
	itemText.Data = "=============================================="
	listNameData = append(listNameData,itemText)

	//1.1
	copy(ar, arrayInit)
	BubbleSort(ar)
	listNameData = append(listNameData,arrayToData(ar,"1.1) bubbleSort"))
	//1.2
	copy(ar, arrayInit)
	BubbleSortInterrupt(ar)
	listNameData = append(listNameData,arrayToData(ar,"1.2) bubbleSortInterrupt"))
	//1.3
	copy(ar, arrayInit)
	BubbleSortInterruptReversed(ar)
	listNameData = append(listNameData,arrayToData(ar,"1.3) bubbleSortInterruptReversed"))
	//1.4
	copy(ar, arrayInit)
	BubbleSortRecursive(ar)
	listNameData = append(listNameData,arrayToData(ar,"1.4) bubbleSortRecursive"))
	//2.1
	copy(ar, arrayInit)
	SelectionSort(ar)
	listNameData = append(listNameData,arrayToData(ar,"2.1) selectionSort"))
	//2.2
	copy(ar, arrayInit)
	SelectionSortByMax(ar)
	listNameData = append(listNameData,arrayToData(ar,"2.2) selectionSortByMax"))
	//2.3
	copy(ar, arrayInit)
	SelectionSortBidirectional(ar)
	listNameData = append(listNameData,arrayToData(ar,"2.3) selectionSortBidirectional"))
	//3.1
	copy(ar, arrayInit)
	InsertionSort(ar)
	listNameData = append(listNameData,arrayToData(ar,"3.1) insertionSort"))
	//4.1
	copy(ar, arrayInit)
	arMegre := MergeSort(ar)
	listNameData = append(listNameData,arrayToData(arMegre,"4.1) mergeSort"))
	//5.1
	copy(ar, arrayInit)
	QuickSort(ar)
	listNameData = append(listNameData,arrayToData(ar,"5.1) quickSort"))

	itemText.Name = "Результату сортировки по возрастанию/убыванию:"
	itemText.Data = "=============================================="
	listNameData = append(listNameData,itemText)

	//1.5
	copy(ar, arrayInit)
	BubbleSortInterruptAsc(ar,ascending)
	listNameData = append(listNameData,arrayToData(ar,"1.5) bubbleSortInterruptAsc"))
	//2.4
	copy(ar, arrayInit)
	SelectionSortBidirectionalAsc(ar,ascending)
	listNameData = append(listNameData,arrayToData(ar,"2.4) selectionSortBidirectionalAsc"))
	//3.2
	copy(ar, arrayInit)
	InsertionSortAsc(ar,ascending)
	listNameData = append(listNameData,arrayToData(ar,"3.2) insertionSortAsc"))
	//4.2
	copy(ar, arrayInit)
	arMegreAsc := MergeSortAsc(ar,ascending)
	listNameData = append(listNameData,arrayToData(arMegreAsc,"4.2) mergeSortAsc"))
	//5.2
	copy(ar, arrayInit)
	QuickSortAsc(ar,ascending)
	listNameData = append(listNameData,arrayToData(ar,"5.2) quickSortAsc"))

	//fmt.Println(ar)
	//fmt.Printf("%v", listNameData)

	return listNameData, nil
}



