package franchise

import "fmt"

type Franchise interface {
 PrepareFood() string
  GetMenu() string
  GetLocation() string
}

type McDonalds struct { 
  Location string
  Menu  string
  Preparation string
}
func (m McDonalds) PrepareFood() string {
  return m.Preparation
}
func (m McDonalds) GetMenu() string {
  return m.Menu
} 
func (m McDonalds) GetLocation() string {
  return m.Location
}
type KFC struct {
  Location string
  Menu  string
  Preparation string
}

func (k KFC) PrepareFood() string {
  return k.Preparation
} 
func (k KFC) GetMenu() string {
  return k.Menu
} 
func (k KFC) GetLocation() string {
  return k.Location
}

type BurgerKing struct {
  Location string
  Menu  string
  Preparation string
} 
func (b BurgerKing) PrepareFood() string {
  return b.Preparation
}        
func (b BurgerKing) GetMenu() string {
  return b.Menu
}
func (b BurgerKing) GetLocation() string {
  return b.Location
}

func Operate(fr Franchise) {
  fmt.Println("Franchise Location:", fr.GetLocation())
  fmt.Println("Menu Offered:", fr.GetMenu())
  fmt.Println("Food Preparation Details:", fr.PrepareFood())
}
