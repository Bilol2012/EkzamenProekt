package main

import (
	"errors"
	"fmt"
)

type Car struct {
	brand  string
	Id     int
	rented bool
}

type Bike struct {
	brand  string
	Id     int
	rented bool
}

type Skate struct {
	brand  string
	Id     int
	rented bool
}

type Vehicle interface {
	Rent() error
	Return() error
	Details() error
}

func (v *Car) Rent() error {
	if v.rented {
		return errors.New("Aвтомобиль уже арендован\n")
	}
	v.rented = true
	fmt.Printf("Аренда автомобиля: %s\n", v.brand)
	return nil
}
//-----------------------------------------------------------
func (v *Car) Return() error {
	if !v.rented {
		return errors.New("Автомобиль не арендован\n")
	}
	v.rented = false
	fmt.Printf("Возврат автомобиля: %s\n", v.brand)
	return nil
}
//-------------------------------------------------------------
func (v *Car) Details() error {
	fmt.Printf("Детали автомобиля: %s, ID: %d, арендован: %v\n", v.brand, v.Id, v.rented)
	return nil
}
//--------------------------------------------------------
func (v *Bike) Rent() error {
	if v.rented {
		return errors.New("Велосипед уже арендован\n")
	}
	v.rented = true
	fmt.Printf("Аренда велосипеда: %n\n", v.brand)
	return nil
}
//---------------------------------------------------------
func (v *Bike) Return() error {
	if !v.rented {
		return errors.New("Велосипед не арендован\n")
	}
	v.rented = false
	fmt.Printf("Возврат велосипеда: %s\n", v.brand)
	return nil
}
//---------------------------------------------------
func (v *Bike) Details() error {
	fmt.Printf("Детали велосипеда: %s, ID: %d, арендован: %v\n", v.brand, v.Id, v.rented)
	return nil
}
//---------------------------------------------------------
func (v *Skate) Rent() error {
	if v.rented {
		return errors.New("Скейты уже арендованы\n")
	}
	v.rented = true
	fmt.Printf("Аренда Скейтов: %s\n", v.brand)
	return nil
}
//---------------------------------------------------------
func (v *Skate) Return() error {
	if !v.rented {
		return errors.New("Скейты не арендованы\n")
	}
	v.rented = false
	fmt.Printf("Возврат Скейта\n: %s", v.brand)
	return nil
}
//----------------------------------------------------------
func (v *Skate) Details() error {
	fmt.Printf("Детали Скейтов: %s, ID: %d, арендован: %v\n", v.brand, v.Id, v.rented)
	return nil
}
//----------------------------------------------------------------
func main() {
	vehicles := []Vehicle{
		&Car{brand: "Toyota", Id: 1, rented: false},
		&Bike{brand: "BMX", Id: 2, rented: false},
		&Skate{brand: "Bulls", Id: 3, rented: false},
	}
//-----------------------------------------------------
	var command int
	for {
		fmt.Println("\nВведите команду!:")
		fmt.Println("1. Посмотреть доступные транспортные средства!")
		fmt.Println("2. Арендовать транспортное средство!")
		fmt.Println("3. Вернуть транспортное средство!")
		fmt.Println("4. Выход!")
		for range []int{0} {
			_, err := fmt.Scanln(&command)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}
		}
//-----------------------------------------------------------------------------------------------------
		switch command {
		case 1:
			for _, vehicle := range vehicles {
				vehicle.Details()
				}

			case 2:
				fmt.Println("Введите ID транспортного средства для аренды:")
				var id int
				for range []int{0} {
					_, err := fmt.Scanln(&id)
					if err != nil {
					fmt.Println("Ошибка ввода:", err)
					continue
				}

				if id <= 0 || id > len(vehicles) {
					fmt.Println("Неверный ID")
					continue
				}

				err = vehicles[id-1].Rent()
				if err != nil {
					fmt.Println("Ошибка аренды:", err)
				}
				break
			}

		case 3:
			fmt.Println("Введите ID транспортного средства для возврата:")
			var id int
			for range []int{0} {
				_, err := fmt.Scanln(&id)
				if err != nil {
					fmt.Println("Ошибка ввода:", err)
					continue
				}

				if id <= 0 || id > len(vehicles) {
					fmt.Println("Неверный ID")
					continue
				}

				err = vehicles[id-1].Return()
				if err != nil {
					fmt.Println("Ошибка возврата:", err)
				}
				break 
			}

		case 4:
			fmt.Println("Выход из ПРОГРАММЫ!!!!!")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

