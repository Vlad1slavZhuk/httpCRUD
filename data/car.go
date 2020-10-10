package data

import (
	"encoding/json"
	"io"
	"sync"
)

type Car struct {
	ID    uint    `json:"id"`
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Color string  `json:"color"`
	Price float64 `json:"price"`
}

//For Test - add 2 car.
var (
	cars = []*Car{
		&Car{
			ID:    1,
			Brand: "Mazda",
			Model: "CX-5",
			Color: "Aqua",
			Price: 25000.00,
		},
		&Car{
			ID:    2,
			Brand: "Aston Martin",
			Model: "One 77",
			Color: "Space Grey",
			Price: 80000.50,
		},
	}
	rwm    sync.Mutex
	lastID uint = 3
)

// ToJSON serializes.
func (c *Car) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

// FromJSON deserializes.
func (c *Car) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}

// GetListCars return list cars.
func GetListCars() []*Car {
	rwm.Lock()
	defer rwm.Unlock()
	return cars
}

// GetCar - returns the specified number from the cars list.
//
// If ok - return car, true. Otherwise - return nil, false
func GetCar(id uint) (*Car, bool) {
	rwm.Lock()
	defer rwm.Unlock()
	for _, car := range cars {
		if car.ID == id {
			return car, true
		}
	}
	return nil, false
}

// AddCar - add Car
//
// If empty data - return false. Otherwise - return true
func AddCar(car *Car) bool {
	rwm.Lock()
	defer rwm.Unlock()
	if len(cars) > 0 {
		lastID = cars[len(cars)-1].ID + 1
	} else {
		lastID = 1
	}

	if car.Brand == "" || car.Model == "" || car.Color == "" || car.Price == 0.0 {
		return false
	}

	car.ID = lastID
	cars = append(cars, car)
	return true
}

// DeleteCar - deletes car by ID
func DeleteCar(id uint) bool {
	rwm.Lock()
	defer rwm.Unlock()
	isFind := false
	for i, car := range cars {
		if car.ID == id {
			if i == len(cars)-1 {
				cars[i] = nil
				cars = cars[:i]
			} else {
				cars = append(cars[:i], cars[i+1:]...)
			}
			isFind = true
			break
		}
	}

	if !isFind {
		return false
	}

	var index uint = 1
	for _, car := range cars {
		if car.ID != index {
			car.ID = index
			index++
		} else {
			index++
		}
	}

	return true
}

// UpdateCar - update car by ID
func UpdateCar(id uint, c *Car) bool {
	rwm.Lock()
	defer rwm.Unlock()
	for _, car := range cars {
		if car.ID == id {
			if len(c.Brand) != 0 || c.Brand != "" {
				car.Brand = c.Brand
			}
			if len(c.Model) != 0 || c.Model != "" {
				car.Model = c.Model
			}
			if len(c.Color) != 0 || c.Color != "" {
				car.Color = c.Color
			}
			if c.Price >= 1000.00 {
				car.Price = c.Price
			}
			return true
		}
	}
	return false
}
