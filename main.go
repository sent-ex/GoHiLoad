package main

// filling unsorted slice with rand shit
// calling sort, passing slice pointer
//
import (
	"fmt"
	"math/rand"
	"time"
)

func swapIfLess(first, second *int) bool {
	if *first < *second {
		//		print(" swapping\n")
		tmp := *first
		*first = *second
		*second = tmp
		return true
	} else {
		//		print("\n")
		return false
	}
}

func printArray(Arr *[]int) {
	for i := 0; i < len(*Arr); i++ {
		println((*Arr)[i])
	}
}

func InsertionSort(Arr *[]int) bool {
	ArrayLen := len(*Arr)
	ThisTime := time.Now().Minute()
	for zoneIndex := 1; zoneIndex < ArrayLen; zoneIndex++ {
		if (*Arr)[zoneIndex] < (*Arr)[zoneIndex-1] {
			//	printArray(Arr)
			j := zoneIndex
			for (j > 0) && (swapIfLess(&(*Arr)[j], &(*Arr)[j-1])) {
				//println((*Arr)[j], ":", (*Arr)[j-1], " swapping")
				j--
			}
		}
		//println("------")
		if time.Now().Minute() != ThisTime {
			ThisTime = time.Now().Minute()
			println(ThisTime, "at: ", zoneIndex)
		}
	}
	return true
}
func IntBubbleSorting1(Arr *[]int) bool {

	if len(*Arr) < 2 {
		return false
	}
	var Swapped = true
	ThisTime := time.Now().Minute()
	for Swapped {
		//time.Sleep(time.Second * 2)
		//println("----")
		Swapped = false
		if ThisTime != time.Now().Minute() {
			println(ThisTime)
			ThisTime = time.Now().Minute()
		}
		for zoneIndex := 0; zoneIndex < len(*Arr)-1; zoneIndex++ {
			//	print((*Arr)[zoneIndex], ":", (*Arr)[zoneIndex]+1)
			if swapIfLess(&(*Arr)[zoneIndex+1], &(*Arr)[zoneIndex]) {
				Swapped = true
				//		print("swap\n")
				//	} else {
				//		print(" stay\n")
			}
			//	time.Sleep(time.Millisecond * 500)
		}
	}
	return true
}

func IntBubbleSorting(Arr *[]int) bool {
	//мин размер массива 2:
	if len(*Arr) < 2 {
		return false
	}
	// начинаем с 1 и дальше увеличиваем индекс конца зоны сортировки
	ThisTime := time.Now().Minute()

	for zoneIndex := 1; zoneIndex < len(*Arr); zoneIndex++ {
		if ThisTime != time.Now().Minute() {
			println(ThisTime, zoneIndex)
			ThisTime = time.Now().Minute()
		}
		//		time.Sleep(time.Second * 2)
		for sortIndex := zoneIndex; sortIndex > 0; sortIndex-- {
			//			print((*Arr)[sortIndex-1], ":", (*Arr)[sortIndex])
			if !swapIfLess(&(*Arr)[sortIndex], &(*Arr)[sortIndex-1]) {
				break
			}
			//			time.Sleep(time.Microsecond * 100)
		}
	}
	//второй индекс уменьшается с конца зоны сортировки до нуля
	//если исследуемый элемент больше соседнего, они меняются местами, самый легкий "всплывает"
	//
	return true
}

func main() {
	println("starting!")
	var ArraySize = 1000000 // будущий размер массива
	StartTime := time.Now()
	myArray := make([]int, 0) // создали пустой слайс под будущую сортировку

	counter := 0 // для элегантности делаем бесконечный цикл с выходом по каунтеру
	for {
		// заполняем инициализированный слайс случайныйми числами
		myArray = append(myArray, rand.Int())
		//myArray = append(myArray, rand.Intn(100))
		counter++
		if counter == ArraySize {
			break
		}
	}
	InsertionSort(&myArray)
	fmt.Println(ArraySize, "units sorting finished", time.Now().Sub(StartTime), "passed")

	//printArray(&myArray)
	//IntBubbleSorting1(&myArray) //передаем алгоритму сортировки указатель на массив
	//IntBubbleSorting(&myArray) //передаем алгоритму сортировки указатель на массив
	// (чтобы не плодить новый)

}
