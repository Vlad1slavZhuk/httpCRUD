package data

import (
	"encoding/json"
	"io"
)

type Car struct {
	//Brand string  `json:"brand"`
	Model string  `json:"model"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

var cars = map[uint64]*Car{
	1: {
		Model: "Mazda CX-5",
		Color: "Aqua",
		Price: 25000.00,
	},
	2: {
		Model: "Aston Martin One 77",
		Color: "Space Grey",
		Price: 80000.50,
	},
}

//TODO
// ToJSON serializes.
func (c *Car) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

//TODO
// FromJSON deserializes.
func (c *Car) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

// GetListCars return list cars.
func GetListCars() map[uint64]*Car {
	return cars
}

// GetCar returns the specified number from the cars list.
func GetCar(id uint64) (*Car, bool) {
	if _, ok := cars[id]; ok {
		return cars[id], true
	}
	return nil, false
}

//TODO
func AddCar(car *Car) bool {
	lastID := uint64(len(cars) + 1)

	if car.Model == "" || car.Color == "" || car.Price == 0.0 {
		return false
	}

	cars[lastID] = car
	return true
}

//TODO
func DeleteCar(id uint64) bool {
	if _, ok := cars[id]; !ok {
		return false
	}
	delete(cars, id)
	return true
}

//TODO
func UpdateCar(id uint64, c *Car) bool {
	_, ok := cars[id]
	if !ok {
		return false
	}

	cars[id] = c
	return true
}
