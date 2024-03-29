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
	println ("starting!")
	var ArraySize = 2000000 // будущий размер массива
	StartTime := time.Now()
	myArray := make([]int, 0) // создали пустой слайс под будущую сортировку

	counter := 0 // для элегантности делаем бесконечный цикл с выходом по каунтеру
	for {
		// заполняем инициализированный слайс случайныйми числами
		myArray = append(myArray, rand.Int())
		counter++
		if counter == ArraySize {
			break
		}
	}
	IntBubbleSorting(&myArray) //передаем алгоритму сортировки указатель на массив
	// (чтобы не плодить новый)
	fmt.Println(ArraySize, " sorting finished", time.Now().Sub(StartTime), "passed")

}
