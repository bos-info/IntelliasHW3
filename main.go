package main

import (
	"fmt"
	"math"
)

// не змінюй мене, бо Віктор буде сваритись!
const (
	totalBudget float64 = 23   // доступний нам бюджет на покупки
	applePrice  float64 = 5.99 // вартість одного яблука
	pearPrice   float64 = 7    // вартість однієї груші
)

func main() {
	//Зміни мене якщо хочеш
	var (
		quantityOfApples         float64 = 9 //кількість яблук для першої задачі
		quantityOfPears          float64 = 8 //кількість груш для першої задачі
		quantityOfApplesInBasket float64 = 2 //кількість яблук для четвертої задачі
		quantityOfPearsInBasket  float64 = 2 //кількість груш для четвертої задачі
	)
	// 1. Скільки грошей треба витратити, щоб купити quantityOfApples яблук та
	// quantityOfPears груш?
	cart1 := quantityOfApples*applePrice + quantityOfPears*pearPrice
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити %0.f яблук та %0.f груш? \n Відповідь: %.2f \n", quantityOfApples, quantityOfPears, cart1)
	// 2. Скільки груш ми можемо купити?
	countOfPears := (totalBudget - math.Mod(totalBudget, pearPrice)) / pearPrice
	fmt.Printf("2. Скільки груш ми можемо купити? \n Відповідь: %0.f \n", countOfPears)
	// 3. Скільки яблук ми можемо купити?
	countOfApple := (totalBudget - math.Mod(totalBudget, applePrice)) / applePrice
	fmt.Printf("3. Скільки яблук ми можемо купити? \n Відповідь: %0.f \n", countOfApple)
	// 4. Чи ми можемо купити quantityOfPearsInBasket груші та quantityOfApplesInBasket яблука?
	var resultOfCounting = "не можемо"
	if totalBudget >= quantityOfApplesInBasket*applePrice+quantityOfPearsInBasket*pearPrice {
		resultOfCounting = "можемо"
	}
	fmt.Printf("4. Чи ми можемо купити %0.f груші та %0.f яблука? \n Відповідь: Ми %s купити %0.f груші та  %0.f яблука \n  ", quantityOfPearsInBasket, quantityOfApplesInBasket, resultOfCounting, quantityOfPearsInBasket, quantityOfApplesInBasket)
}
