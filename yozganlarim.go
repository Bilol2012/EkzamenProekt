// package main

// import (
// 	"fmt"
// )

// // type Student struct {
// // 	Name  string
// // 	Age   int
// // 	Grade float64
// // }

// // func main() {
// // 	students := []Student{
// // 		{"Ali", 20, 85.0},
// // 		{"Oyna", 21, 90.0},
// // 		{"John", 22, 78.0},
// // 		{"Zarina", 19, 88.0},
// // 	}

// // 	for {
// // 		fmt.Println("Tanlov:")
// // 		fmt.Println("1. Talaba ma'lumotlarini ko'rsatish")
// // 		fmt.Println("2. Barcha talabalar ro'yxatini ko'rsatish")
// // 		fmt.Println("3. Orta bahoni ko'rsatish")
// // 		fmt.Println("4. Chiqish")

// // 		var sum int
// // 		fmt.Println("Tanlovingizni kiriting (1-4): ")
// // 		fmt.Scanln(&sum)
// // 		switch sum {
// // 		case 1:
// // 			var studentIndex int
// // 			fmt.Println("Talaba raqami: ", len(students))
// // 			fmt.Scanln(&studentIndex)

// //	if studentIndex >= 1 && studentIndex <= len(students) {
// //		student := students[studentIndex-1]
// //		fmt.Println("Talaba:",student.Name, "Yosh:", student.Age," Baho:", student.Grade)
// //	} else {
// //
// //		fmt.Println("Noto'g'ri talaba raqami")
// //	}
// type Student struct {
// 	Name  string
// 	Age   int
// 	Id    int
// 	Grade map[string]int
// }

// func main() {
// 	students := make([]Student, 0)
// 	// var studentCount int
// 	for {
// 		fmt.Println("1. Talaba qo'shish")
// 		fmt.Println("2. Barcha talabalar ro'yxatini ko'rsatish")
// 		fmt.Println("3. Orta bahoni ko'rsatish")
// 		fmt.Println("4. Chiqish")

// 		var sum int
// 		fmt.Println("Tanlovingizni kiriting (1-5): ")
// 		fmt.Scanln(&sum)

// 		switch sum {
// 		case 1:
// 			var name string
// 			var age int
// 			grade := make(map[string]int)
// 			var id int
// 			fmt.Println("Talaba id kiriting: ")
// 			fmt.Scanln(&id)

// 			fmt.Println("Talaba ismini kiriting: ")
// 			fmt.Scanln(&name)
// 			fmt.Println("Talaba yoshini kiriting: ")
// 			fmt.Scanln(&age)

// 			var input int
// 			fmt.Scanln(&input)
// 			for i := 0; i < input; i++ {
// 				var fan string
// 				var baho int
// 				fmt.Println("fanni kiriting")
// 				fmt.Scanln(&fan)
// 				fmt.Println("bahoni kiriting")
// 				fmt.Scanln(&baho)
// 				grade[fan] = baho
// 			}

// 			// studentCount++
// 			students = append(students, Student{Id: id, Name: name, Age: age, Grade: grade})
// 			fmt.Println("Talaba qo'shildi.")
// 		case 2:
// 			fmt.Println("Barcha talabalar ro'yxati:")
// 			for index, student := range students {
// 				fmt.Println(index+1, student.Name)
// 			}
// 		case 3:
// 			var aydi, total int
// 			fmt.Println("aydi kiriting")
// 			fmt.Scanln(&aydi)
// 			g := map[string]int{}
// 			for _, student := range students {
// 				// totalGrade += student.Grade
// 				if aydi == student.Id {
// 					g = student.Grade
// 					fmt.Println(g)
// 					break
// 				}
// 			}
// 			for _, v := range g {
// 				total += v
// 			}

// 			ortanchaGrade := total / len(g)
// 			fmt.Println("Ortancha baho: ", ortanchaGrade)
// 		case 4:
// 			fmt.Println("Dasturdan chiqdingiz")
// 			return

// 		default:
// 			fmt.Println("Noto'g'ri tanlov, iltimos qaytadan urinib ko'ring.")
// 		}

		//  var num1, num2 float64
		//  var operator string

		//  fmt.Println("Введите первое число: ")
		//  fmt.Scanln(&num1)

		//  fmt.Println("Введите оператор (+, -, *, /): ")
		//  fmt.Scanln(&operator)

		//  fmt.Print("Введите второе число: ")
		//  fmt.Scanln(&num2)

		//  if operator == "+" {
		//   fmt.Println("Результат:", num1+num2)
		//  } else if operator == "-" {
		//   fmt.Println("Результат:", num1-num2)
		//  } else if operator == "*" {
		//   fmt.Println("Результат:", num1*num2)
		//  } else if operator == "/" {
		//   if num2 != 0 {
		//    fmt.Println("Результат:", num1/num2)
		//   }else {
		//   fmt.Println("Ошибка: неверный оператор")
		//  }
		//  }

		// var yil string

		// fmt.Println("1ta faslni yozing: ")
		// fmt.Scanln(&yil)

		// switch yil {
		// 	case "qish":
		// 		fmt.Println("Issiq kiym")
		// 	case "bahor":
		// 		fmt.Println("Jinsi")
		// 	case "yoz":
		// 		fmt.Println("Shorti futbolka")
		// 	case "kuz":
		// 		fmt.Println("Issiq kiym")
		// 	default:
		// 		fmt.Println("Notog'ri!")
		// }

		// a, b := 0, 1

		// for i := 0; i < 12; i++ {
		//     fmt.Println(a)
		//     a, b = b, a+b
		// }
		// for i :=1; i < 100; i++ {
		// 	if i%2 != 0{

		// 		fmt.Println(i)
		// 	}
		// }
		// i := 0
		// for i < 100 {
		// 	if i%2 != 0 {

		// 		fmt.Println(i)
		// 	}
		// 	i++
		// }
		// name := "rasul"
		// for name =="rasul"{
		// 	for i:=0; i<=10;i++{
		// 		fmt.Println(i)
		// 		if i == 7 {
		// 			name="qwerty"
		// 		}
		// 	}
		// }

		// for i := 1; i <= 100; i++ {

		//     if i%3 == 0 && i%5 == 0 {
		//         fmt.Println("fithbase")
		//     } else if i%3 == 0 {
		//         fmt.Println("fith")
		//     } else if i%5 == 0 {
		//         fmt.Println("base")
		//     } else {
		//         fmt.Println(i)
		//     }
		// }

		//     for i := 1; i <= 100; i++ {
		//         switch {
		//         case i%3 == 0 && i%5 == 0:
		//             fmt.Println("fithbase")
		//         case i%3 == 0:
		//             fmt.Println("fith")
		//         case i%5 == 0:
		//             fmt.Println("base")
		//         default:
		//             fmt.Println(i)
		// 		}

		// }
		// for i := 1; i <= 12; i++ {
		// 	switch i {
		// 	case 1:
		// 		fmt.Println("Январ")
		// 	case 2:
		// 		fmt.Println("Февраль")
		// 	case 3:
		// 		fmt.Println("Март")
		// 	case 4:
		// 		fmt.Println("Апрель")
		// 	case 5:
		// 		fmt.Println("Май")
		// 	case 6:
		// 		fmt.Println("Июнь")
		// 	case 7:
		// 		fmt.Println("Июль")
		// 	case 8:
		// 		fmt.Println("Август")
		// 	case 9:
		// 		fmt.Println("Сентябрь")
		// 	case 10:
		// 		fmt.Println("Октябрь")
		// 	case 11:
		// 		fmt.Println("Ноябрь")
		// 	case 12:
		// 		fmt.Println("Декабрь")
		// 	default:
		// 		fmt.Println("Номаълум ой")
		// 	}
		// }
		// //______-----------------------------------------------------
		// var op string
		//  var num1, num2 float64
		//  maxIterations := 5

		//  for i := 0; i < maxIterations; i++ {
		//   fmt.Println("Kalkulyator:")
		//   fmt.Println("Birinchi sonni kiriting: ")
		//   fmt.Scanln(&num1)
		//   fmt.Println("Ikkinchi sonni kiriting: ")
		//   fmt.Scanln(&num2)
		//   fmt.Println("Amalni kiriting (+, -, *, / yoki 'exit' chiqish uchun): ")
		//   fmt.Scanln(&op)

		//   if op == "exit" {
		//    fmt.Println("Dastur tugatildi.")
		//    break
		//   }

		//   switch op {
		//   case "+":
		//    fmt.Println("Natija: ", num1+num2)
		//   case "-":
		//    fmt.Println("Natija: ", num1-num2)
		//   case "*":
		//    fmt.Println("Natija: ", num1*num2)
		//   case "/":
		//    if num2 != 0 {
		//     fmt.Println("Natija: ", num1/num2)
		//    } else {
		//     fmt.Println("0 ga bo‘lish mumkin emas!")
		//    }
		//   default:
		//    fmt.Println("Noto‘g‘ri amal kiritildi.")
		//   }

		//   if i == maxIterations-1 {
		//    fmt.Println("Maksimal ishlash soniga yetildi. Dastur tugatildi.")
		//   }
		//  }
		//  // ---------------------------------------------------------------------------------
		//     // привет!
		// 	//------------------------------------------------------
		// 	var n int
		//     fmt.Println("Сон киритинг: ")
		//     fmt.Scanln(&n)

		//     oddSum := 0
		//     evenSum := 0

		//     for i := 1; i <= n; i++ {
		//         switch i % 2 {
		//         case 0:
		//             evenSum += i
		//         case 1:
		//             oddSum += i
		//         }
		//     }

		//     fmt.Println("Ток сонлар суммаси: ", oddSum)
		//     fmt.Println("Жуфт сонлар суммаси: ", evenSum)

		// fmt.Println("hello world")

		//------------------------------------------------------- slice and array ----------------------------------------------------------------

		// -------------------------------------------------uyga vazifa---------------------------
		// numbers := [5]int{10, 20, 30, 40, 50}
		// sum := 0
		// for _, number := range numbers {
		// 	fmt.Println("Index", number)
		// 	sum = sum + number
		// }
		// fmt.Println("Array elementlarining summasi:", sum)
		// //-----------------
		// fmt.Println("-----------------------------------2-----------------------")
		// //---------------------------------------------------
		// slice := numbers[1:4]

		// max := slice[0]

		// for _, number := range slice {
		// 	if number > max {
		// 		max = number
		// 	}
		// }

		// fmt.Println("Slice ichidagi 1chi indexdan 4gacham bolganidan eng katta qiymat:", max)
		// //----------------------------------------------------
		// fmt.Println("--------------------------------------------------3-----------------------")
		// //--------------------------------------------------------
		// var numbersNot []int

		// var n int
		// fmt.Println("Nechta son kiritmoqchisiz? ")
		// fmt.Scanln(&n)

		// for i := 0; i <= n; i++ {
		// 	var number int
		// 	fmt.Println("Sonni kiriting : ", i+1)
		// 	fmt.Scanln(&number)
		// 	numbersNot = append(numbersNot, number)
		// }

		// oddCount := 0
		// evenCount := 0

		// for _, number := range numbersNot {
		// 	if number%2 == 0 {
		// 		evenCount++
		// 	} else {
		// 		oddCount++
		// 	}
		// }

		// fmt.Println("Juft sonlar soni:", evenCount)
		// fmt.Println("Toq sonlar soni:", oddCount)
		// //---------------------------------------------------------------------------
		// fmt.Println("--------------------------------------------------------------4------------")
		// //-----------------------------------------------------------------------

		// var numbersIs []int

		// var nil int
		// fmt.Println("Son Kiriting (eki sonlar) ")
		// fmt.Scanln(&nil)

		// for i := 0; i <= nil; i++ {
		// 	var number int
		// 	fmt.Println("Sonni kiriting : ", i+1)
		// 	fmt.Scanln(&number)
		// 	numbersIs = append(numbersIs, number)
		// }

		// fmt.Println("Musbat sonlar:")

		// for _, number := range numbersIs {
		// 	if number >= 0 {
		// 		fmt.Println(number)
		// 	}
		// }
		// /////////0-----------------------------------------------------------
		// fmt.Println("------------------------------------------------------5-----------------------------")

		// numUp := []int{10, 20, 30, 40, 50}

		// run := make([]int, len(numUp))

		// for i, number := range numUp {
		// 	run[len(numUp)-1-i] = number
		// }

		// fmt.Println("Asl slice:", numUp)
		// fmt.Println("Teskari slice:", run)

		//----------------------------------------------------------------------------------

		// nums := []int{4, 1, 2, 5,   6, 8, 10,10}
		// m := make(map[int]int)
		// for _, v := range nums {
		// 	m[v]++
		// }
		// for k, v := range m {
		// 	if v == 2 {
		//        fmt.Println(k)
		// 	   break
		// 	}
		// }
		//--------------------------------------------------------------------------

		// var input, output string
		// fmt.Println("Matn kiriting: ")
		// fmt.Scanln(&input)

		// m := make(map[string]int)
		// for _, v := range input {
		// 	m[string(v)]++
		// }
		// fmt.Println(m)

		// //------------------------------------------------------------------------

		// fmt.Scanln(&input)
		// arr := make([]string, 0, len(input))
		// for _, s := range input {
		// 	arr = append(arr, string(s))
		// }
		// for i := len(arr) - 1; i >= 0; i-- {
		// 	output += arr[i]
		// 	fmt.Println(output)
		// }
		// fmt.Println(output)

		//-------------------------
		// var person Person

		// fmt.Println("Oquvchini ismi:")
		// fmt.Scanln(&person.Ism)
		// fmt.Println("O'quvchini yoshi")
		// fmt.Scanln(&person.Age)
		// fmt.Println("O'quvchini sinfi")
		// fmt.Scanln(&person.Sinf)

		// fmt.Println("Oquvchining ismi", person.Ism, "Oquvchinig yoshi", person.Age, "O'quvchini sinfi", person.Sinf)

		// m:= map[int] string{
		// 	1:"Ism"
		// 	2:"Yosh"F
		// 	3:"Sinf"

		// }
		// v,ok:=m[2]
		// fmt.Println("value:", v,"Ex",ok)

		//-----------------------------------------------------------------------------------------------------------------------

		/* Ikkinchi task */

		// var input, output string
		// // Hello😁
		// fmt.Scanln(&input)

		// arr := make([]string, 0, len(input))
		// for _, s := range input {
		// 	arr = append(arr, string(s))

		// 	fmt.Println(arr)
		// }

		// for i := len(arr) - 1; i >= 0; i-- {
		// 	output += arr[i] // output = output + arr[i] bilan bir xil

		// 	fmt.Println(output)
		// }

		// fmt.Println(output)

		/* Uchinchi task */
		// type Cat struct {
		// 	Name string
		// 	Color string
		// 	Voice string
		// }

		// cat := Cat{}

		// fmt.Scanln(&cat.Name)
		// fmt.Scanln(&cat.Color)
		// fmt.Scanln(&cat.Voice)

		// fmt.Println("My cat's name is ", cat.Name, " he is ", cat.Color, " and he says ", cat.Voice)

		// sayHello("Rasul")

		// var v, t string = "v", "t"
		// var s = "s"
		// n := 123
		// var (
		// 	a, b int
		// 	c string
		// )

		// fmt.Printf("%T, %T, %T", v, t, s) Variable qaysi typeda ekanini aniqlash yo'li

		// if 5 == 5 || 2 == 5 {

		// } else if 2 == 2 {

		// } else {}

		// v, t := 1, 2
		// switch {
		// case v == 1:
		// 	fmt.Println(v)
		// case t == 2:
		// 	fmt.Println(t)
		// default:
		// 	fmt.Println("this is default")
		// }

		/*
			int, int8, int16, int32, int64
			uint, uint8, uint16, uint32, uint64
			float, float32, float64
			bool
			string
			byte
			rune == int32
			[1]int, []int --> arr := make([]int, 0, 10) --> a := append(arr, n, 1, 2, 4)
			m := make(map[int]string) --> v, ok := m[12]
			struct
		*/

		// for i := 0; i < 10; i++ {}
		// for 5 == 5 {}
		// for i, v := range arr {}
		// func sayHello(name string) {
		// 	fmt.Println("Hello, ", name, "!")

		//--------------------------------------- ЭКЗАМЕН --------------------------------

// 	}

// }
// package main

// import "fmt"

// func BankAccount()(func(float64), func(float64), func(), func()) {
// 	balance := 0.0

// 	deposit := func(amount float64) {
// 		if amount > 0 {
// 			balance += amount
// 			fmt.Println("mablag' qo'shildi. Hozirgi hisob: ", amount, balance)
// 		} else {
// 			fmt.Println("Hisobni to'ldirish uchun miqdor musbat bo'lishi kerak.")
// 		}
// 	}

// 	mablagYechish:= func(amount float64) {
// 		if amount > 0 && amount <= balance {
// 			balance -= amount
// 			fmt.Println("mablag' yechildi. Hozirgi hisob: ", amount, balance)
// 		} else if amount > balance {
// 			fmt.Println("Yetarli mablag' mavjud emas.")
// 		} else {
// 			fmt.Println("Yechish miqdori musbat bo'lishi kerak.")
// 		}
// 	}

// 	mablagniKorsatish := func() {
// 		fmt.Println("Hozirgi hisobdagi mablag'i: ", balance)
// 	}

// 	tozalash := func() {
// 		balance = 0
// 		fmt.Println("Hisob 0 ga tushirildi.")
// 	}

// 	return deposit, mablagYechish,  mablagniKorsatish, tozalash
// }

// func main() {
// 	deposit, withdraw, showBalance, resetBalance := BankAccount()

// 	deposit(1000)    
// 	withdraw(250)    
// 	showBalance()    
// 	resetBalance()   
// 	showBalance()    
// }

	// // func BankAccount() (deposit func(float64), withdraw func(float64), showBalance func(), resetBalance func()) {
	// // 	balance := 0.0
	
	// // 	deposit = func(amount float64) {
	// // 		if amount > 0 {
	// // 			balance += amount
	// // 			fmt.Println(amount, "miqdorda mablag' qo'shildi. Hozirgi hisob:", balance)
	// // 		} else {
	// // 			fmt.Println("Hisobni to'ldirish uchun miqdor musbat bo'lishi kerak.")
	// // 		}
	// // 	}
	
	// // 	withdraw = func(amount float64) {
	// // 		if amount > 0 && amount <= balance {
	// // 			balance -= amount
	// // 			fmt.Println(amount, "miqdorda mablag' yechildi. Hozirgi hisob:", balance)
	// // 		} else if amount > balance {
	// // 			fmt.Println("Yeterli mablag' mavjud emas.")
	// // 		} else {
	// // 			fmt.Println("Yechish miqdori musbat bo'lishi kerak.")
	// // 		}
	// // 	}
	
	// // 	showBalance = func() {
	// // 		fmt.Println("Hozirgi hisobdagi mablag':", balance)
	// // 	}
	
	// // 	resetBalance = func() {
	// // 		defer fmt.Println("Hisob 0 ga tushirildi.")
	// // 		balance = 0
	// // 	}
	// // 	return
	// // }
	
	// // func main() {
	// // 	deposit, withdraw, showBalance, resetBalance := BankAccount()
	
	// //		deposit(1000)
	// //		withdraw(250)
	// //		showBalance()
	// //		resetBalance()
	// //		showBalance()
	// //	}
	// )
	
	// // type Mish struct {
	// // 	Name  string
	// // 	Color string
	// // 	// eyes  string
	// // }
	
	// // func (c Mish) AAA() {
	// // 	fmt.Println(c.Name, "AAA!!!")
	// // }
	
	// // func (c *Mish) ChangeColor(color string) {
	// // 	c.Color = color
	// // 	// c.eyes = color
	// // }
	
	// // func main() {
	// // 	mish := Mish{
	// // 		Name:  "Джери",
	// // 		Color: "Коричневый",
	// // 		// eyes:  "Зеленные",
	// // 	}
	// // 	fmt.Println(mish)
	
	// // 	mish.AAA()
	// // 	mish.ChangeColor("Черный")
	
	// //		fmt.Println(mish)
	// //	}
	
	// // -----------------------------------------------uy ishi------------------------------------------------------
	// type Mish struct {
	// 	Name  string
	// 	Color string
	// 	Eyes  string
	// }
	
	// func (c Mish) AAA() {
	// 	fmt.Println(c.Name, "Поменяйтесь!")
	// }
	
	// func (c *Mish) ChangeColor(color string) {
	// 	c.Color = color
	// }
	
	// func (c *Mish) ChangeEyes(eyeColor string) {
	// 	c.Eyes = eyeColor
	// }
	
	// func main() {
	// 	mish := Mish{
	// 		Name:  "Джери:",
	// 		Color: "Коричневый,",
	// 		Eyes:  "Глаза: Зеленые",
	// 	}
	// 	fmt.Println(mish)
	// 	mish.AAA()
	// 	mish.ChangeColor("Черный,")
	// 	mish.ChangeEyes(" Глаза: Синие")
	// 	fmt.Println(mish)
	// }

	//------------------------------------------------------------------------------------------------------------
	// type Customer interface {
//  CalculateReward(purchases []float64) float64
// }

// type RegularCustomer struct{}

// func (r RegularCustomer) CalculateReward(purchases []float64) float64 {
//  total := sum(purchases)
//  return total * 0.01 
// }

// type PremiumCustomer struct{}

// func (p PremiumCustomer) CalculateReward(purchases []float64) float64 {
//  total := sum(purchases)
//  return total * 0.05 
// }

// type VIPCustomer struct{}

// func (v VIPCustomer) CalculateReward(purchases []float64) float64 {
//  total := sum(purchases)
//  return total * 0.10
// }

// func sum(numbers []float64) float64 {
//  total := 0.0
//  for _, number := range numbers {
//   total += number
//  }
//  return total
// }

// func main() {
//  customers := map[string]struct {
//   customerType Customer
//   purchases    []float64
//  }{
//   "Ali":   {RegularCustomer{}, []float64{100, 200, 300}},
//   "Bilol":     {PremiumCustomer{}, []float64{400, 500}},
//   "Abu": {VIPCustomer{}, []float64{1000, 2000, 3000}},
//  }

//  for name, data := range customers {
//   reward := data.customerType.CalculateReward(data.purchases)
//   fmt.Printf(\n"%s uchun umumiy cashback: %d", name, reward)
//  }
// }

//--------------------------------------------------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
	// m := make(map[string]int)
	// m := map[string]int{
	// 	"Biloliddin": 0,
	// 	"Abdulloh": 1,
	// }


	// m["Abubakr"] = 991234567
	// m["Abubakr"] = 991234568

	// // name := m["Abubakr"]

	// // fmt.Println(m["adfadsf"])
	// // fmt.Println(m["adfadsadff"])

	// v, ok := m["asfasdfsadf"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("No such key")
	// }

// 	c := map[string]Customer {
// 		"Abdulloh": Vip{
// 			Purchases: []int{1500, 1000},
// 		},
// 		"Biloliddin": Premium{
// 			Purchases: []int{1000, 500},
// 		},
// 		"Abubakr": Regular{
// 			Purchases: []int{10000, 2400},
// 		},
// 	}

// 	for k, v := range c {
// 		fmt.Printf("\nCustomer: %s, got %d amount as cashback", k, v.CalculateReward())
// 	}

// }

// type Customer interface {
// 	CalculateReward() int
// }

// type Regular struct {
// 	Purchases []int
// }

// func (r Regular) CalculateReward() int {
// 	sum := 0
// 	for _, v := range r.Purchases {
// 		sum += v
// 	}

// 	return int(float32(sum)*0.01)
// }

// type Premium struct {
// 	Purchases []int
// }

// func (p Premium) CalculateReward() int {
// 	sum := 0
// 	for _, v := range p.Purchases {
// 		sum += v
// 	}

// 	return int(float32(sum)*0.05)
// }

// type Vip struct {
// 	Purchases []int
// }

// func (v Vip) CalculateReward() int {
// 	sum := 0
// 	for _, p := range v.Purchases {
// 		sum += p
// 	}

// 	return int(float32(sum)*0.1)
// }

//------------------------------------------------------------------------------

// m := map[string]int{
	// 	"Biloliddin": 0,
	// 	"Xusan":      1,
	// 	"Abu":        2,
	// }
	// m["Xusan"] = 950888889
	// m["Biloliddin"] = 935775500
	// m["Abu"] = 99808822034
	// fmt.Println(m)
	// v, ok := m["Abu"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("No such key")
	// }

//-----------------------------------------------------------------------------


// type Shape interface {
// 	Avea() int
// }

// type Treugolnik struct {
// 	calcAll []int
// }

// type Kvadrat struct {
// 	calcAll []int
// }

// type Krug struct {
// 	calcAll []int
// }

// func (t Treugolnik) Avea() int {
// 	sum := 0
// 	for _, v := range t.calcAll {
// 		sum += v * 3 / 2
// 	}
// 	return sum
// }

// func (k Kvadrat) Avea() int {
// 	sum := 0
// 	for _, v := range k.calcAll {
// 		sum += v * v
// 	}
// 	return sum
// }

// func (c Krug) Avea() int {
// 	sum := 0
// 	for _, v := range c.calcAll {
// 		sum += int(3.14 * float64(v) * float64(v))
// 	}
// 	return sum
// }

// func main() {
// 	treugolnik := Treugolnik{calcAll: []int{1, 2, 3}}
// 	kvadrat := Kvadrat{calcAll: []int{4, 5, 6}}
// 	krug := Krug{calcAll: []int{7, 8, 9}}

// 	figures := []Shape{treugolnik, kvadrat, krug}

// 	for _, figure := range figures {
// 		switch f := figure.(type) {
// 		case Treugolnik:
// 			fmt.Println("Треугольник, площадь:", f.Avea())
// 		case Kvadrat:
// 			fmt.Println("Квадрат, площадь:", f.Avea())
// 		case Krug:
// 			fmt.Println("Круг, площадь:", f.Avea())
// 		default:
// 			fmt.Println("Неизвестная фигура")
// 		}
// 	}
// }
//-------------------------------------------------------------------------------
// type Shape interface {
	// 	Area() float64
	// }
	
	// type Treugolnik struct {
	// 	sides []float64
	// }
	
	// type Kvadrat struct {
	// 	sides []float64
	// }
	
	// type Krug struct {
	// 	radii []float64
	// }
	
	// func (t Treugolnik) Area() float64 {
	// 	sum := 0.0
	// 	for _, v := range t.sides {
	// 		sum += 3 * v
	// 	}
	// 	return sum
	// }
	
	// func (k Kvadrat) Area() float64 {
	// 	sum := 0.0
	// 	for _, v := range k.sides {
	// 		sum += 4 * v
	// 	}
	// 	return sum
	// }
	
	// func (c Krug) Area() float64 {
	// 	sum := 0.0
	// 	for _, v := range c.radii {
	// 		sum += 2 * 3.14 * v
	// 	}
	// 	return sum
	// }
	
	// func main() {
	// 	var input int
	// 	fmt.Println("1. Uchburchak perimetrini hisoblash")
	// 	fmt.Println("2. Kvadrat perimetrini hisoblash")
	// 	fmt.Println("3. Aylana perimetrini hisoblash")
	// 	fmt.Print("Tanlang: ")
	// 	fmt.Scanln(&input)
	
	// 	switch input {
	// 	case 1:
	// 		treugolnik := Treugolnik{}
	// 		fmt.Println("Uchburchakning 3 ta tomonini kiriting:")
	// 		for i := 0; i < 3; i++ {
	// 			var side float64
	// 			fmt.Scanln(&side)
	// 			treugolnik.sides = append(treugolnik.sides, side)
	// 		}
	// 		fmt.Println("Uchburchakning umumiy perimetri: ", treugolnik.Area())
	
	// 	case 2:
	// 		kvadrat := Kvadrat{}
	// 		fmt.Println("Kvadratning 1 ta ravniy tomonini kiriting:")
	// 		for i := 0; i < 1; i++ {
	// 			var side float64
	// 			fmt.Scanln(&side)
	// 			kvadrat.sides = append(kvadrat.sides, side)
	// 		}
	// 		fmt.Println("Kvadratning umumiy perimetri: ", kvadrat.Area())
	
	// 	case 3:
	// 		krug := Krug{}
	// 		fmt.Println("Krugni 1 ta radiusini kiriting:")
	// 		for i := 0; i < 1; i++ {
	// 			var radius float64
	// 			fmt.Scanln(&radius)
	// 			krug.radii = append(krug.radii, radius)
	// 		}
	// 		fmt.Println("Krugning umumiy perimetri: ", krug.Area())
	
	// 	default:
	// 		fmt.Println("Noto'g'ri!!!")
	// 	}
	// }

//--------------------------------------------------------------------------------
// //---------------ERRORS:-----------------
// package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var (
// 		a, b int
// 		op   string
// 	)
// 	fmt.Scanln(&a)
// 	fmt.Scanln(&op)
// 	fmt.Scanln(&b)
// 	res, err := calculate(a, b, op)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(res)
// }


// type MyError struct {
// 	Time        time.Time
// 	Description string
// }

// func (e MyError) Error() string {
// 	return fmt.Sprintf("%s: %s", e.Time.String(), e.Description)
// }

// func calculate(a, b int, op string) (int, error) {
// 	switch op {
// 	case "+":
// 		return a + b, nil
// 	case "-":	 
// 		return a - b, nil
// 	case "*":
// 		return a * b, nil
// 	case "/":
// 		if b == 0 {
// 			// return 0, MyError{Time: time.Now(), Description: "error"}
// 			return 0, errors.New("error")
// 		}
// 		return a / b, nil
// 	default:
// 		// return 0, MyError{Time: time.Now(), Description: "Invalid operator"}
// 		return 0, errors.New("Invalid Operator")
// 	}
// }
//----------------------------------------------------------------
// package main

// import "fmt"

// type Employee interface {
//     GetSalary() int
// }

// type MonthlyEmployee struct {
//     salary int
// }
// type ContractEmployee struct {
// 	contractValue int
// }
// type HourlyEmploye struct {
// 	hourlyRate int
// 	hoursWorked int
// }

// func (me MonthlyEmployee) GetSalary() int {
//     return me.salary
// }


// func (he HourlyEmploye) GetSalary() int {
//     return he.hourlyRate * he.hoursWorked
// }


// func (ce ContractEmployee) GetSalary() int {
//     return ce.contractValue
// }

// func main() {
//     me := MonthlyEmployee{salary: 5000}

//     he := HourlyEmploye{hourlyRate: 20, hoursWorked: 160}

//     ce := ContractEmployee{contractValue: 10000}

//     workers := []Employee{me, he, ce}

//     for i, worker := range workers {
//         fmt.Println("Ishchi:  Maosh: ", i+1, worker.GetSalary())
//     }
// }

//-----------------------------------------------------------------------


// func main() {
// 	category := CreateCategory("Electronics")

// 	product1 := CreateProduct("Phone", 699.99, 10)
// 	product2 := CreateProduct("Laptop", 999.99, 5)
// 	product3 := CreateProduct("Tablet", 399.99, 15)

// 	err := category.AddProduct(product1)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	err = category.AddProduct(product2)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	err = category.AddProduct(product3)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	category.Display()

// 	updatedProduct := CreateProduct("Phone", 649.99, 8)
// 	err = category.UpdateProduct(updatedProduct)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	category.Display()

// 	err = category.DeleteProduct("Laptop")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	category.Display()

// }


//-------------------------------------------------------------------------------------









































// Product va Category structlari bo’ladi; Product: Name, Price, Quantity; Category: Name, Products array.
// 2. Yangi product yaratish uchun CreateProduct funksiyasi yoziladi.
// 3. Yangi category yaratish uchun CreateCategory funksiyasi yoziladi.
// 4. Category uchun “AddProduct() error” metodi yoziladi. Categoryga yangi product qo’shish vazifasini bajaradi.
// 5. Category uchun “UpdateProduct() error” metodi yoziladi. Categorydagi productni o’zgartirish vazifasini bajaradi. Har ikkisida error qaytariladi
// 6. Displayabe interface yaratiladi va ichida Display() metodi yoziladi. Har ikki struct uchun Display() metodi implement qilinadi. Bu metod Product uchun porductni sifatlarini ekranga chiqarib beradi, Category uchun category nomi va unga tegishli productlarni chiqarib beradi.

// Bonus: Category uchun “DeleteProduct() error” metodi yoziladi. Categorydan ma’lum productni o’chirish vazifasini bajaradi.
// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Product struct {
// 	Name     string
// 	Price    float64
// 	Quantity int
// }

// type Category struct {
// 	Name     string
// 	Products []Product
// }

// type Displayable interface {
// 	Display()
// }

// func (p Product) Display() {
// 	fmt.Printf("Название продукта: %s, Цена: %.2f, Количество: %d\n", p.Name, p.Price, p.Quantity)
// }

// func (c Category) Display() {
// 	fmt.Printf("Категория: %s\n", c.Name)
// 	fmt.Println("Продукты:")
// 	for _, product := range c.Products {
// 		product.Display()
// 	}
// }

// func CreateProduct(name string, price float64, quantity int) Product {
// 	return Product{Name: name, Price: price, Quantity: quantity}
// }

// func CreateCategory(name string) Category {
// 	return Category{Name: name, Products: []Product{}}
// }

// func (c *Category) AddProduct(product Product) error {
// 	for _, p := range c.Products {
// 		if p.Name == product.Name {
// 			return errors.New("продукт с таким именем уже существует в категории")
// 		}
// 	}
// 	c.Products = append(c.Products, product)
// 	return nil
// }

// func (c *Category) UpdateProduct(product Product) error {
// 	for i, p := range c.Products {
// 		if p.Name == product.Name {
// 			c.Products[i] = product
// 			return nil
// 		}
// 	}
// 	return errors.New("продукт не найден в категории")
// }

// func main() {
// 	var categories []Category

// 	for {
// 		var choice string
// 		fmt.Println("1) Создать продукт")
// 		fmt.Println("2) Создать категорию")
// 		fmt.Println("3) Добавить продукт в категорию")
// 		fmt.Println("4) Обновить продукт в категории")
// 		fmt.Println("5) Все данные")
// 		fmt.Println("6) Выйти")
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case "1":
// 			var name string
// 			var price float64
// 			var quantity int
// 			fmt.Print("Название продукта: ")
// 			fmt.Scanln(&name)
// 			fmt.Print("Цена продукта: ")
// 			fmt.Scanln(&price)
// 			fmt.Print("Количество продукта: ")
// 			fmt.Scanln(&quantity)
// 			product := CreateProduct(name, price, quantity)
// 			product.Display()
// 			fmt.Println("Продукт создан.")
// 		case "2":
// 			var categoryName string
// 			fmt.Print("Название категории: ")
// 			fmt.Scanln(&categoryName)
// 			category := CreateCategory(categoryName)
// 			categories = append(categories, category)
// 			category.Display()
// 			fmt.Println("Категория создана.")
// 		case "3":
// 			var categoryName, productName string
// 			var price float64
// 			var quantity int
// 			fmt.Print("Название категории: ")
// 			fmt.Scanln(&categoryName)
// 			var category *Category
// 			for i := range categories {
// 				if categories[i].Name == categoryName {
// 					category = &categories[i]
// 					break
// 				}
// 			}
// 			if category == nil {
// 				fmt.Println("Категория не найдена!")
// 				break
// 			}
// 			fmt.Print("Название продукта: ")
// 			fmt.Scanln(&productName)
// 			fmt.Print("Цена продукта: ")
// 			fmt.Scanln(&price)
// 			fmt.Print("Количество продукта: ")
// 			fmt.Scanln(&quantity)
// 			product := CreateProduct(productName, price, quantity)
// 			err := category.AddProduct(product)
// 			if err != nil {
// 				fmt.Println("Ошибка:", err)
// 			} else {
// 				fmt.Println("Продукт добавлен")
// 			}
// 		case "4":
// 			var categoryName, productName string
// 			var price float64
// 			var quantity int
// 			fmt.Print("Название категории: ")
// 			fmt.Scanln(&categoryName)
// 			var category *Category
// 			for i := range categories {
// 				if categories[i].Name == categoryName {
// 					category = &categories[i]
// 					break
// 				}
// 			}
// 			if category == nil {
// 				fmt.Println("Категория не найдена!")
// 				break
// 			}
// 			fmt.Print("Название продукта для обновления: ")
// 			fmt.Scanln(&productName)
// 			fmt.Print("Новая цена продукта: ")
// 			fmt.Scanln(&price)
// 			fmt.Print("Новое количество продукта: ")
// 			fmt.Scanln(&quantity)
// 			product := CreateProduct(productName, price, quantity)
// 			err := category.UpdateProduct(product)
// 			if err != nil {
// 				fmt.Println("Ошибка:", err)
// 			} else {
// 				fmt.Println("Продукт обновлен")
// 			}
// 		case "5":
// 			if len(categories) == 0 {
// 				fmt.Println("Нет категорий для отображения.")
// 			} else {
// 				for _, category := range categories {
// 					category.Display()
// 				}
// 			}
// 		case "6":
// 			return
// 		default:
// 			fmt.Println("Неверный выбор!")
// 		}
// 	}
// }
//-------------------------------------------------------

// type Book struct {
// 	Title          string
// 	Author         string
// 	ISBN           string
// 	AvailableCopies int
// }

// type Library struct {
// 	Books map[string]*Book
// }

// var ErrBookNotFound = errors.New("Книга не найдена!")
// var ErrNoCopiesAvailable = errors.New("Нет доступных копий!")

// func (l *Library) AddBook(book *Book) {
// 	if l.Books == nil {
// 		l.Books = make(map[string]*Book)
// 	}
// 	l.Books[book.ISBN] = book
// }

// func (l *Library) CheckOutBook(isbn string) error {
// 	book, exists := l.Books[isbn]
// 	if !exists {
// 		return ErrBookNotFound
// 	}
// 	if book.AvailableCopies == 0 {
// 		return ErrNoCopiesAvailable
// 	}
// 	book.AvailableCopies--
// 	return nil
// }

// func (l *Library) ReturnBook(isbn string) error {
// 	book, exists := l.Books[isbn]
// 	if !exists {
// 		return ErrBookNotFound
// 	}
// 	book.AvailableCopies++
// 	return nil
// }

// func (l *Library) ListAvailableBooks() {
// 	fmt.Println("Доступные книги:")
// 	i := 1
// 	for _, book := range l.Books {
// 		if book.AvailableCopies > 0 {
// 			fmt.Printf("%d. %s - Доступные копии: %d\n", i, book.Title, book.AvailableCopies)
// 			i++
// 		}
// 	}
// }

// type User struct {
// 	Name            string
// 	CheckedOutBooks map[string]*Book
// }

// func (u *User) CheckOutBook(l *Library, isbn string) error {
// 	err := l.CheckOutBook(isbn)
// 	if err != nil {
// 		return err
// 	}
// 	book := l.Books[isbn]
// 	if u.CheckedOutBooks == nil {
// 		u.CheckedOutBooks = make(map[string]*Book)
// 	}
// 	u.CheckedOutBooks[isbn] = book
// 	fmt.Printf("Успешно взяли книгу %s\n", book.Title)
// 	return nil
// }

// func (u *User) ReturnBook(l *Library, isbn string) error {
// 	book, exists := u.CheckedOutBooks[isbn]
// 	if !exists {
// 		return ErrBookNotFound
// 	}
// 	err := l.ReturnBook(isbn)
// 	if err != nil {
// 		return err
// 	}
// 	delete(u.CheckedOutBooks, isbn)
// 	fmt.Printf("Успешно вернули книгу %s\n", book.Title)
// 	return nil
// }

// func addBookFromUser(l *Library) {
// 	var title, author, isbn string
// 	var availableCopies int

// 	fmt.Println("Введите данные о книге:")
// 	fmt.Print("Название книги: ")
// 	fmt.Scanln(&title)
// 	fmt.Print("Автор книги: ")
// 	fmt.Scanln(&author)
// 	fmt.Print("ISBN книги: ")
// 	fmt.Scanln(&isbn)
// 	fmt.Print("Количество доступных копий: ")
// 	fmt.Scanln(&availableCopies)

// 	book := &Book{
// 		Title:           title,
// 		Author:          author,
// 		ISBN:            isbn,
// 		AvailableCopies: availableCopies,
// 	}
// 	l.AddBook(book)
// 	fmt.Println("Книга добавлена в библиотеку!")
// }

// func main() {
// 	lib := &Library{}
// 	user1 := &User{Name: "Biloliddin"}

// 	for {
// 		fmt.Println("\nВведите команду:")
// 		fmt.Println("1. Добавить книгу (add)")
// 		fmt.Println("2. Посмотреть доступные книги (list)")
// 		fmt.Println("3. Взять книгу (checkout)")
// 		fmt.Println("4. Вернуть книгу (return)")
// 		fmt.Println("5. Выйти (exit)")

// 		var command string
// 		fmt.Scanln(&command)

// 		switch command {
// 		case "add":
// 			addBookFromUser(lib)
// 		case "list":
// 			lib.ListAvailableBooks()
// 		case "checkout":
// 			fmt.Println("Введите ISBN книги для взятия: ")
// 			var isbnToCheckout string
// 			fmt.Scanln(&isbnToCheckout)
// 			if err := user1.CheckOutBook(lib, isbnToCheckout); err != nil {
// 				fmt.Println("Ошибка:", err)
// 			}
// 		case "return":
// 			fmt.Println("Введите ISBN книги для возврата: ")
// 			var isbnToReturn string
// 			fmt.Scanln(&isbnToReturn)
// 			if err := user1.ReturnBook(lib, isbnToReturn); err != nil {
// 				fmt.Println("Ошибка:", err)
// 			}
// 		case "exit":
// 			fmt.Println("Выход из программы.")
// 			return
// 		default:
// 			fmt.Println("Неверная команда. Попробуйте снова.")
// 		}
// 	}
// }

//-----------------------------------------------------------------------------------------

	

































































































































































































































