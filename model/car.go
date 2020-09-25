package model

type Car struct {
	Model string  `json:"model"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

var cars = map[uint64]*Car{
	1: {
		Model: "Mazda",
		Color: "Red",
		Price: 25000.00,
	},
	2: {
		Model: "Opel",
		Color: "Blue",
		Price: 15000.68,
	},
}

func GetListCars() map[uint64]*Car {
	return cars
}

func GetCar(id uint64) (*Car, bool) {
	if _, ok := cars[id]; ok {
		return cars[id], true
	}
	return nil, false
}

func AddCar(car *Car) {
	lastID := uint64(len(cars) + 1)
	cars[lastID] = car
}

func DeleteCar(id uint64) bool {
	if _, ok := cars[id]; !ok {
		return false
	}
	delete(cars, id)
	return true
}
