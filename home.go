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
	return nil
}

func (v *Car) Return() error {
	v.rented = false
	if !v.rented {
		return errors.New("Автомобиль не арендован\n")
	}
	
	return nil
}

func (v *Car) Details() error {
	fmt.Printf("Детали автомобиля: %s, ID: %d, арендован: %v\n", v.brand, v.Id, v.rented)
	return nil
}

func (v *Bike) Rent() error {
	if v.rented {
		return errors.New("Велосипед уже арендован\n")
	}
	
	v.rented = true
	return nil
}

func (v *Bike) Return() error {
	if !v.rented {
		return errors.New("Велосипед не арендован\n")
	}
	v.rented = false
	fmt.Printf("Возврат велосипеда: %s\n", v.brand)
	return nil
}

func (v *Bike) Details() error {
	fmt.Printf("Детали велосипеда: %s, ID: %d, арендован: %v\n", v.brand, v.Id, v.rented)
	return nil
}

func (v *Skate) Rent() error {
	if v.rented {
		return errors.New("Скейты уже арендованы\n")
	}
	v.rented = true
	return nil
}

func (v *Skate) Return() error {
	if !v.rented {
		return errors.New("Скейты не арендованы\n")
	}
	v.rented = false 
	fmt.Printf("Возврат Скейта: %s\n", v.brand)
	return nil
}

func (v *Skate) Details() error {
	fmt.Printf("Детали Скейтов: %s, ID: %d, арендование: %v\n", v.brand, v.Id, v.rented)
	return nil
}

func main() {

	var vehicles []Vehicle
	var command int

	for {
		fmt.Println("\nВведите команду!: ")
		fmt.Println("1. Посмотреть транспортные средства!")
		fmt.Println("2. Арендовать транспортное средство!")
		fmt.Println("3. Вернуть транспортное средство!")
		fmt.Println("4. Добавить новое транспортное средство!")
		fmt.Println("5. Выход!")

		_, err := fmt.Scanln(&command)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		switch command {
		case 1:
			if len(vehicles) == 0 {
				fmt.Println("Нет доступных транспортных средств.")
			} else {
				for _, vehicle := range vehicles {
					vehicle.Details()
				}
			}

		case 2:
			fmt.Println("Введите ID транспортного средства для аренды:")
			var id int
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

		case 3:
			fmt.Println("Введите ID транспортного средства для возврата:")
			var id int
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

		case 4:
			fmt.Println("Выберите тип транспортного средства для добавления:")
			fmt.Println("1. Добавить автомобиль")
			fmt.Println("2. Добавить велосипед")
			fmt.Println("3. Добавить скейт")
			var vehicleType int
			_, err := fmt.Scanln(&vehicleType)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}

			var brand string
			fmt.Println("Введите бренд транспортного средства:")
			_, err = fmt.Scanln(&brand)
			if err != nil {
				fmt.Println("Ошибка ввода:", err)
				continue
			}

			id := len(vehicles) + 1

			switch vehicleType {
			case 1:
				vehicles = append(vehicles, &Car{brand: brand, Id: id, rented: false})
			case 2:
				vehicles = append(vehicles, &Bike{brand: brand, Id: id, rented: false})
			case 3:
				vehicles = append(vehicles, &Skate{brand: brand, Id: id, rented: false})
			default:
				fmt.Println("Неверный выбор!")
				continue
			}
			fmt.Println("Транспортное средство успешно добавлено.")

		case 5:
			fmt.Println("Выход из ПРОГРАММЫ!!!!!")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
