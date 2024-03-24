package sortCustomTypesTest

import (
    "math/rand"
    "testing"
    "fmt"
    "tracker/cmd/task/sortCustomTypes"
    "strconv"
    "encoding/json"
	"io/ioutil"
    "time"
   )

type ItemNameDataTest struct {
	Name       string  `json:"name"`
	Test01     string  `json:"test01"`
    Test02     string  `json:"test02"`
    Test03     string  `json:"test03"`
    Test04     string  `json:"test04"`
    Test05     string  `json:"test05"`
    Test06     string  `json:"test06"`
    Test07     string  `json:"test07"`
    Allocs     string  `json:"allocs"` 
    TimeRun    string  `json:"timeRun"`
}

var (
	ArrName = []string {"1.1) BubbleSort", "1.2) BubbleSortInterrupt", "1.3) BubbleSortInterruptReversed", "1.4) BubbleSortRecursive","2.1) SelectionSort","2.2) SelectionSortByMax","2.3) SelectionSortBidirectional","3.1) InsertionSort","4.1) MergeSort","5.1) QuickSort"}
)


func calculateBenchmarkResult(result testing.BenchmarkResult) (string, int64) {
    // Получение значения result.AllocsPerOp
    allocsPerOp := result.AllocsPerOp()
    // Вычисление количества разрядов в целой части
    intPart := int(float64(result.T.Nanoseconds()) / float64(result.N))
    intPartStr := strconv.Itoa(intPart)
    intPartLength := len(intPartStr)
    intPartLengthStr := strconv.Itoa(intPartLength)
    //intPartLengthStrV := "["+intPartLengthStr+"]"+intPartStr +" ns/op; "+ "[-]"+strconv.FormatInt(allocsPerOp, 10) +" allocs/op"
    intPartLengthStrV := "["+intPartLengthStr+"]"+intPartStr +"; "+ "[-]"+strconv.FormatInt(allocsPerOp, 10) +""
    if (allocsPerOp>0){
        //intPartLengthStrV = "["+intPartLengthStr+"]"+intPartStr +" ns/op; "+ "[+]"+strconv.FormatInt(allocsPerOp, 10) +" allocs/op"
        intPartLengthStrV = "["+intPartLengthStr+"]"+intPartStr +"; "+ "[+]"+strconv.FormatInt(allocsPerOp, 10) +""
    }
 
    return intPartLengthStrV, allocsPerOp
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}

func generateSliceMinMax(size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = i
	}

	return ar
}

func generateSliceMaxMin(size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = size-i
	}

	return ar
}

func BenchmarkSortArrayBubbleSort(tBubbleSort int) func(b *testing.B) {
    return func(b *testing.B) {
        fmt.Printf("++++: %v\n", tBubbleSort)
        b.StopTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
		    ar := generateSlice(10, 10)
            b.StartTimer()
			sortCustomTypes.BubbleSort(ar)
			b.StopTimer()
		}
    }
}


//1.1) обычный по возрастанию
func BenchmarkSortArrayBubbleSort01(tSort int) func(b *testing.B) {
	return func(b *testing.B) {
		   b.StopTimer()
		   b.ReportAllocs()
		   for i := 0; i < b.N; i++ {
			   ar := generateSlice(10, 10)
               switch tSort {
                    case 0:
                        b.StartTimer()
                        sortCustomTypes.BubbleSort(ar)
                        b.StopTimer()
                    case 1:
                        b.StartTimer()
                        sortCustomTypes.BubbleSortInterrupt(ar)
                        b.StopTimer()
                    case 2:
                        b.StartTimer()
                        sortCustomTypes.BubbleSortInterruptReversed(ar)
                        b.StopTimer()
                    case 3:
                        b.StartTimer()
                        sortCustomTypes.BubbleSortRecursive(ar)
                        b.StopTimer()
                    case 4:
                        b.StartTimer()
                        sortCustomTypes.SelectionSort(ar)
                        b.StopTimer()
                    case 5:
                        b.StartTimer()
                        sortCustomTypes.SelectionSortByMax(ar)
                        b.StopTimer()
                    case 6:
                        b.StartTimer()
                        sortCustomTypes.SelectionSortBidirectional(ar)
                        b.StopTimer()
                    case 7:
                        b.StartTimer()
                        sortCustomTypes.InsertionSort(ar)
                        b.StopTimer()
                    case 8:
                        b.StartTimer()
                        _ = sortCustomTypes.MergeSort(ar)
                        b.StopTimer()
                    case 9:
                        b.StartTimer()
                        sortCustomTypes.QuickSort(ar)
                        b.StopTimer()
                }
		   }
	   }
}
func BenchmarkSortArrayBubbleSort02(tSort int) func(b *testing.B)   {
    return func(b *testing.B) {
        b.StopTimer()
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            ar := generateSlice(100, 1000)
            switch tSort {
                 case 0:
                     b.StartTimer()
                     sortCustomTypes.BubbleSort(ar)
                     b.StopTimer()
                 case 1:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterrupt(ar)
                     b.StopTimer()
                 case 2:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterruptReversed(ar)
                     b.StopTimer()
                 case 3:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortRecursive(ar)
                     b.StopTimer()
                 case 4:
                     b.StartTimer()
                     sortCustomTypes.SelectionSort(ar)
                     b.StopTimer()
                 case 5:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortByMax(ar)
                     b.StopTimer()
                 case 6:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortBidirectional(ar)
                     b.StopTimer()
                 case 7:
                     b.StartTimer()
                     sortCustomTypes.InsertionSort(ar)
                     b.StopTimer()
                 case 8:
                     b.StartTimer()
                     _ = sortCustomTypes.MergeSort(ar)
                     b.StopTimer()
                 case 9:
                     b.StartTimer()
                     sortCustomTypes.QuickSort(ar)
                     b.StopTimer()
             }
        }
    }
}
func BenchmarkSortArrayBubbleSort03(tSort int) func(b *testing.B) {
    return func(b *testing.B) {
        b.StopTimer()
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            ar := generateSlice(100, 100000)
            switch tSort {
                 case 0:
                     b.StartTimer()
                     sortCustomTypes.BubbleSort(ar)
                     b.StopTimer()
                 case 1:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterrupt(ar)
                     b.StopTimer()
                 case 2:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterruptReversed(ar)
                     b.StopTimer()
                 case 3:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortRecursive(ar)
                     b.StopTimer()
                 case 4:
                     b.StartTimer()
                     sortCustomTypes.SelectionSort(ar)
                     b.StopTimer()
                 case 5:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortByMax(ar)
                     b.StopTimer()
                 case 6:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortBidirectional(ar)
                     b.StopTimer()
                 case 7:
                     b.StartTimer()
                     sortCustomTypes.InsertionSort(ar)
                     b.StopTimer()
                 case 8:
                     b.StartTimer()
                     _ = sortCustomTypes.MergeSort(ar)
                     b.StopTimer()
                 case 9:
                     b.StartTimer()
                     sortCustomTypes.QuickSort(ar)
                     b.StopTimer()
             }
        }
    }
}
func BenchmarkSortArrayBubbleSort04(tSort int) func(b *testing.B)   {
    return func(b *testing.B) {
        b.StopTimer()
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            ar := generateSlice(10000, 100000)
            switch tSort {
                 case 0:
                     b.StartTimer()
                     sortCustomTypes.BubbleSort(ar)
                     b.StopTimer()
                 case 1:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterrupt(ar)
                     b.StopTimer()
                 case 2:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterruptReversed(ar)
                     b.StopTimer()
                 case 3:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortRecursive(ar)
                     b.StopTimer()
                 case 4:
                     b.StartTimer()
                     sortCustomTypes.SelectionSort(ar)
                     b.StopTimer()
                 case 5:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortByMax(ar)
                     b.StopTimer()
                 case 6:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortBidirectional(ar)
                     b.StopTimer()
                 case 7:
                     b.StartTimer()
                     sortCustomTypes.InsertionSort(ar)
                     b.StopTimer()
                 case 8:
                     b.StartTimer()
                     _ = sortCustomTypes.MergeSort(ar)
                     b.StopTimer()
                 case 9:
                     b.StartTimer()
                     sortCustomTypes.QuickSort(ar)
                     b.StopTimer()
             }
        }
    }
}
func BenchmarkSortArrayBubbleSort05(tSort int) func(b *testing.B)   {
    return func(b *testing.B) {
        b.StopTimer()
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            ar := generateSlice(1000000, 100000)
            switch tSort {
                 case 0:
                     b.StartTimer()
                     sortCustomTypes.BubbleSort(ar)
                     b.StopTimer()
                 case 1:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterrupt(ar)
                     b.StopTimer()
                 case 2:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterruptReversed(ar)
                     b.StopTimer()
                 case 3:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortRecursive(ar)
                     b.StopTimer()
                 case 4:
                     b.StartTimer()
                     sortCustomTypes.SelectionSort(ar)
                     b.StopTimer()
                 case 5:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortByMax(ar)
                     b.StopTimer()
                 case 6:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortBidirectional(ar)
                     b.StopTimer()
                 case 7:
                     b.StartTimer()
                     sortCustomTypes.InsertionSort(ar)
                     b.StopTimer()
                 case 8:
                     b.StartTimer()
                     _ = sortCustomTypes.MergeSort(ar)
                     b.StopTimer()
                 case 9:
                     b.StartTimer()
                     sortCustomTypes.QuickSort(ar)
                     b.StopTimer()
             }
        }
    }
}
func BenchmarkSortArrayBubbleSort06(tSort int) func(b *testing.B)   {
    return func(b *testing.B) {
        b.StopTimer()
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            ar := generateSliceMaxMin(100000)
            switch tSort {
                 case 0:
                     b.StartTimer()
                     sortCustomTypes.BubbleSort(ar)
                     b.StopTimer()
                 case 1:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterrupt(ar)
                     b.StopTimer()
                 case 2:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterruptReversed(ar)
                     b.StopTimer()
                 case 3:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortRecursive(ar)
                     b.StopTimer()
                 case 4:
                     b.StartTimer()
                     sortCustomTypes.SelectionSort(ar)
                     b.StopTimer()
                 case 5:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortByMax(ar)
                     b.StopTimer()
                 case 6:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortBidirectional(ar)
                     b.StopTimer()
                 case 7:
                     b.StartTimer()
                     sortCustomTypes.InsertionSort(ar)
                     b.StopTimer()
                 case 8:
                     b.StartTimer()
                     _ = sortCustomTypes.MergeSort(ar)
                     b.StopTimer()
                 case 9:
                     b.StartTimer()
                     sortCustomTypes.QuickSort(ar)
                     b.StopTimer()
             }
        }
    }
}
func BenchmarkSortArrayBubbleSort07(tSort int) func(b *testing.B)   {
    return func(b *testing.B) {
        b.StopTimer()
        b.ReportAllocs()
        for i := 0; i < b.N; i++ {
            ar := generateSliceMinMax(100000)
            switch tSort {
                 case 0:
                     b.StartTimer()
                     sortCustomTypes.BubbleSort(ar)
                     b.StopTimer()
                 case 1:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterrupt(ar)
                     b.StopTimer()
                 case 2:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortInterruptReversed(ar)
                     b.StopTimer()
                 case 3:
                     b.StartTimer()
                     sortCustomTypes.BubbleSortRecursive(ar)
                     b.StopTimer()
                 case 4:
                     b.StartTimer()
                     sortCustomTypes.SelectionSort(ar)
                     b.StopTimer()
                 case 5:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortByMax(ar)
                     b.StopTimer()
                 case 6:
                     b.StartTimer()
                     sortCustomTypes.SelectionSortBidirectional(ar)
                     b.StopTimer()
                 case 7:
                     b.StartTimer()
                     sortCustomTypes.InsertionSort(ar)
                     b.StopTimer()
                 case 8:
                     b.StartTimer()
                     _ = sortCustomTypes.MergeSort(ar)
                     b.StopTimer()
                 case 9:
                     b.StartTimer()
                     sortCustomTypes.QuickSort(ar)
                     b.StopTimer()
             }
        }
    }
}

func makeTestListTest(numberSort int)(ItemNameDataTest, error) {

    startTime := time.Now()

    var itemText ItemNameDataTest

    if ((numberSort<0)||(numberSort>9)){
		return itemText, fmt.Errorf("numberSort д.б. в пределе от 0 до 9")
	}

    var allocsSum, allocsSumOp int64
    allocsSum  = 0
    allocsSumOp = 0
    itemText.Name = ArrName[numberSort]

    bubbleSortBenchmark01 := BenchmarkSortArrayBubbleSort01(numberSort)
    itemText.Test01, allocsSumOp = calculateBenchmarkResult(testing.Benchmark(bubbleSortBenchmark01))
    allocsSum += allocsSumOp

    bubbleSortBenchmark02 := BenchmarkSortArrayBubbleSort02(numberSort)
    itemText.Test02, allocsSumOp = calculateBenchmarkResult(testing.Benchmark(bubbleSortBenchmark02))
    allocsSum += allocsSumOp

    bubbleSortBenchmark03 := BenchmarkSortArrayBubbleSort03(numberSort)
    itemText.Test03, allocsSumOp = calculateBenchmarkResult(testing.Benchmark(bubbleSortBenchmark03))
    allocsSum += allocsSumOp

    bubbleSortBenchmark04 := BenchmarkSortArrayBubbleSort04(numberSort)
    itemText.Test04, allocsSumOp = calculateBenchmarkResult(testing.Benchmark(bubbleSortBenchmark04))
    allocsSum += allocsSumOp

    bubbleSortBenchmark05 := BenchmarkSortArrayBubbleSort05(numberSort)
    itemText.Test05, allocsSumOp = calculateBenchmarkResult(testing.Benchmark(bubbleSortBenchmark05))
    allocsSum += allocsSumOp

    bubbleSortBenchmark06 := BenchmarkSortArrayBubbleSort06(numberSort)
    itemText.Test06, allocsSumOp = calculateBenchmarkResult(testing.Benchmark(bubbleSortBenchmark06))
    allocsSum += allocsSumOp

    bubbleSortBenchmark07 := BenchmarkSortArrayBubbleSort07(numberSort)
    itemText.Test07, allocsSumOp = calculateBenchmarkResult(testing.Benchmark(bubbleSortBenchmark07))
    allocsSum += allocsSumOp
    
    itemText.Allocs = "[-]"+strconv.FormatInt(allocsSum, 10)
    if (allocsSum>0){
        itemText.Allocs = "[+]"+strconv.FormatInt(allocsSum, 10)
    }

    elapsedTime := time.Since(startTime)
    itemText.TimeRun = formatMinutesToMMSS(elapsedTime.Minutes())

    return itemText, nil

}

func formatMinutesToMMSS(minutes float64) string {
    duration := time.Duration(minutes * float64(time.Minute))
    timeValue := time.Time{}.Add(duration)
   
    return timeValue.Format("04:05")
}

func SortArrayGetListTest(tSort int)([]ItemNameDataTest, error) {

    listNameData := make([]ItemNameDataTest, 0)

    itemText, err := makeTestListTest(tSort)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

	return listNameData, nil
}

func fullArrayGetListTest()([]ItemNameDataTest, error) {

    listNameData := make([]ItemNameDataTest, 0)
    //1.1) BubbleSort - 0: Время работы: 246.394063 секунд, 4.1065677166666665 минут
    itemText, err := makeTestListTest(0)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //1.2) BubbleSortInterrupt - 1: Время работы: 213.6463188 секунд, 3.56077198 минут
    itemText, err = makeTestListTest(1)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //1.3) BubbleSortInterruptReversed - 2: Время работы: 221.0996433 секунд, 3.6849940549999998 минут
    itemText, err = makeTestListTest(2)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //1.4) BubbleSortRecursive - 3: Время работы: 204.763584 секунд, 3.4127264 минут
    itemText, err = makeTestListTest(3)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //2.1) SelectionSort - 4: Время работы: 203.978512 секунд, 3.399641866666667 минут
    itemText, err = makeTestListTest(4)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //2.2) SelectionSortByMax - 5: Время работы: 197.8855593 секунд, 3.298092655 минут
    itemText, err = makeTestListTest(5)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //2.3) SelectionSortBidirectional - 6: Время работы: 141.4990308 секунд, 2.3583171800000002 минут
    itemText, err = makeTestListTest(6)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //3.1) InsertionSort - 7: Время работы: 228.0015193 секунд, 3.8000253216666664 минут
    itemText, err = makeTestListTest(7)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //4.1) MergeSort - 8: Время работы: 34.337702 секунд, 0.5722950333333333 минут
    itemText, err = makeTestListTest(8)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

    //5.1) QuickSort - 9: Время работы: 107.9532208 секунд, 1.7992203466666665 минут
    itemText, err = makeTestListTest(9)
    if err != nil {
        return nil, fmt.Errorf(fmt.Sprint(err))	
    }   
    listNameData = append(listNameData,itemText)

	return listNameData, nil
}

func SaveToJSONFile(data[]ItemNameDataTest, filename string) error {

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err))
	}

	err = ioutil.WriteFile(filename, jsonData, 0)
	if err != nil {
		return fmt.Errorf(fmt.Sprint(err))	 
	}

    return nil
}

func LoadFromJSONFile(filePath string) ([]ItemNameDataTest, error) {
	itemNameDataTest:= []ItemNameDataTest{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {	 return nil, err }

	err = json.Unmarshal(data, &itemNameDataTest)
	if err != nil {	 return nil, err }

	return itemNameDataTest, nil
}

