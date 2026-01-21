package interfaces


type Franchise interface {
	PrepareFood() string
	GetMenu() []string
	GetLocation() string
}


type McDonalds struct{}

func (m McDonalds) PrepareFood() string {
	return "Preparing Big Mac and Fries"
}

func (m McDonalds) GetMenu() []string {
	return []string{"Big Mac", "Pizza", "Fries", "Coke"}
}

func (m McDonalds) GetLocation() string {
	return "Hyderabad"
}


type KFC struct{}

func (k KFC) PrepareFood() string {
	return "Frying crispy chicken"
}

func (k KFC) GetMenu() []string {
	return []string{"Fried Chicken", "Chicken Burger", "Popcorn Chicken"}
}

func (k KFC) GetLocation() string {
	return "Bangalore"
}


type BurgerKing struct{}

func (b BurgerKing) PrepareFood() string {
	return " Grilling Whopper"
}

func (b BurgerKing) GetMenu() []string {
	return []string{"Whopper", "Cheeseburger", "Onion Rings"}
}

func (b BurgerKing) GetLocation() string {
	return "Chennai"
}
