package interfaces

import "fmt"

type Franchise interface {
	PrepareFood() string
	GetMenu() string
	GetLocation() string
}
type McDonalds struct {
	Menu           string
	FoodPrepResult string
	Location       string
}

func (m McDonalds) PrepareFood() string {
	return m.FoodPrepResult
}
func (m McDonalds) GetMenu() string {
	return m.Menu
}
func (m McDonalds) GetLocation() string {
	return m.Location
}

type KFC struct {
	Menu           string
	FoodPrepResult string
	Location       string
}

func (k KFC) PrepareFood() string {
	return k.FoodPrepResult
}
func (k KFC) GetMenu() string {
	return k.Menu
}
func (k KFC) GetLocation() string {
	return k.Location
}

type BurgerKing struct {
	Menu           string
	FoodPrepResult string
	Location       string
}

func (b BurgerKing) PrepareFood() string {
	return b.FoodPrepResult
}
func (b BurgerKing) GetMenu() string {
	return b.Menu
}
func (b BurgerKing) GetLocation() string {
	return b.Location
}
func Operate(fr Franchise) {
	fmt.Println(fr.GetLocation())
	fmt.Println(fr.GetMenu())
	fmt.Println(fr.PrepareFood())
}