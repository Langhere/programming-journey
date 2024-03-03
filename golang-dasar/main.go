package main

import "fmt"

func beyondHello() {
	x := 2
	y := 4
	fmt.Println("Value y : ", y)
	fmt.Println("Value x : ", x)
	result := learnMultiple(x, y)
	fmt.Println("sum : ", result)
	tesString()
	tesArray()
	fmt.Println()
	createStart()
	arrayNoExplisit()
	learnSlice()
	learnMap()
	learnIfElse()
	learnSwitch()
	learnFor()
}

func tesString() {
	string1 := `this can use
literal blablabala "yuhu",
i use 3 paragraph
	`
	fmt.Println(string1)
}

func tesArray() {
	var array [4]int
	s := []int{1, 2, 3}
	array[0] = 1
	fmt.Println(array[0])
	fmt.Println("before s append", s[0], s[1], s[2])
	fmt.Println("After s append")
	s = append(s, 4, 5, 6, 7, 8)
	fmt.Print("{ ")
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i])
		fmt.Print(", ")
	}
	fmt.Print("}")
	fmt.Println()
	fmt.Println("s1 is copying from s")
	s1 := s
	fmt.Print("{ ")
	for i := 0; i < len(s1); i++ {
		fmt.Print(s1[i])
		fmt.Print(", ")
	}
	fmt.Print("}")
}

func arrayNoExplisit() {
	names := [...]string{"mae", "suro", "amalia"}
	fmt.Println(names)
}

func learnSlice() {
	fmt.Println("Learn Slice")
	names := [...]string{"mae", "suro", "amalia", "surok"}
	slice1 := names[1:2]
	slice2 := names[:]
	// var slice []string := names[:] -> ini bentuk manual
	slice3 := names[:3]
	slice4 := names[2:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println("Learn Append Slice")
	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	fmt.Println(days)
	slicedays1 := days[5:]
	fmt.Println(slicedays1)
	slicedays1[0] = "sabtu baru"
	slicedays1[1] = "minggu baru"
	fmt.Println(slicedays1)
	slicedays2 := append(slicedays1, "libur baru")
	fmt.Println(slicedays2)

	fmt.Println("Learn Make Slace")
	newSlice := make([]string, 2, 5)
	newSlice[0] = "mae"
	newSlice[1] = "suro"
	newSlice2 := append(newSlice, "amalia")
	newSlice3 := append(newSlice, "surok")
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println(newSlice3)

	// copy slice
	fmt.Println("Learn copy slice...")
	fromSlice := newSlice
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println(fromSlice)

}

func learnMap() {
	fmt.Println("Learn Map")
	// [type key]type value
	map1 := map[string]string{
		"name": "Mae",
		"age":  "20",
		"city": "Jakarta",
	}

	fmt.Println(map1)
	fmt.Println(map1["name"])
	fmt.Println("Learn Make structure Map")
	// kalau mau map nya kosong dulu
	var map2 map[string]string = map[string]string{}
	map2["name"] = "Maesuro"
	map2["age"] = "20"
	map2["city"] = "mataram"
	fmt.Println(map2)
	delete(map2, "city")
	fmt.Println(map2)

	person := make(map[string]string)
	person["name"] = "Mae"
	person["age"] = "28"
	person["city"] = "lotim"
	fmt.Println(person)

}

func learnIfElse() {
	name := "maesuro amalia"
	if cond := len(name); cond > 100 {
		fmt.Println("To Long")
	} else {
		fmt.Print("Enough")
	}
}

func learnSwitch() {
	fmt.Println()
	name := "Gilang"
	switch name {
	case "Gilang":
		fmt.Println("Mae")
	case "Suro":
		fmt.Println("Suro")
	case "Amalia":
		fmt.Println("Amalia")
	default:
		fmt.Println("Not Found")
	}

	switch cond := len(name); cond > 2 {
	case true:
		fmt.Println("To Long")
	case false:
		fmt.Print("Enough")
	default:
		fmt.Println("Not Found")
	}
	fmt.Println()

}

func learnFor() {
	names := []string{"eko", "kurniawan"}
	for key, value := range names {
		fmt.Println("key: ", key, ": "+value)
	}
	fmt.Println("Without key")
	for _, value := range names {
		fmt.Println(value)
	}

	// break and continue
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke - ", i)
	}
	fmt.Println("continue")
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("perulangan ke - ", i)
	}
}

// learn func
func helloName(name string) string {
	helo := "hello" + name
	return helo
}

func getName() (firstName, middleName, lastname string) {
	firstName = "Anugrah"
	middleName = "Gilang"
	lastname = "Ramadhan"
	return firstName, middleName, lastname
}

// variadic func
func sumAll(num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

func createStart() {
	fmt.Println("Learn Create Triangle")
	y := "*"
	for i := 1; i < 7; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(y, " ")
		}
		fmt.Println()
	}
}

//learn func as paramater
func callName(name string, sanitation func(name string) string) string {
	sanitationname := sanitation(name)
	return "hello" + sanitationname
}

func sanitation(name string) string {
	if name == "anj" || name == "ajg" || name == "anjing" {
		return ".."
	} else {
		return name
	}
}

// type func
type Filter func(string) string

func callName2(name string, filter Filter) string {
	filteredName := filter(name)
	return "hello" + filteredName
}

func FilterSan(name string) string {
	if name == "anj" || name == "ajg" || name == "anjing" {
		return ".."
	} else {
		return name
	}
}

// anonymous func

type Sanis func(string) bool

func callSan(name string, sanis Sanis) {
	if sanis(name) {
		fmt.Println("Your Blocked")
	} else {
		fmt.Println(name)
	}
}

// defer func -> mirip await pada javascritp
func first() {
	fmt.Println("selesai memanggil fungsi")
}

func sec() {
	defer first()
	fmt.Println("memanggil fungsi")
}

// panic -> menghentikan program namun defer ttep dijalankan
func cekError(cond bool) {
	defer first()
	if cond {
		panic("error")
	}
}

// reovery -> mengambil error panic dan program tetap berjalan

func defercek() {
	fmt.Println("selesai")
	message := recover()
	fmt.Println("pesan error : ", message)
}

func cekError2(cond bool) {
	defer defercek()
	if cond {
		panic("error bruh")
	}
}

// struct

type Custom struct {
	name, call, city string
	age              int
}

// learn method(fungsi yang menempel pada struct)
func (custom Custom) sayHello(name string) {
	fmt.Println("hello", name, "my name is", custom.name)

}

// learn interface(unknow function-> fungsi yg blm dbuat di awal)

type Person interface {
	callmyName() string
}

type Student struct {
	name string
}

func (student Student) callmyName() string {
	return "my name is " + student.name
}

// interface kosong -> bisa ngembaliin return value apa aja
func ups() any {
	// return 1
	return "yes hello"
}

// asertion -> change type data interface
func tesAssert() any {
	return "ups"
}

// for pointer
type kota struct {
	name, city string
}

// func pointer and method pointer

func tesPointer(name *string) {
	*name = "sisi"
}

type funcPointer struct {
	name string
}

func (funcpoint *funcPointer) married() {
	funcpoint.name = "Mr " + funcpoint.name
}

func learnMultiple(x int, y int) int {
	return x + y
}

func main() {
	beyondHello()
	result := helloName("eko")
	fmt.Println(result)
	fmt.Println(helloName("wahyu"))
	fmt.Println(getName())
	count := sumAll(10, 10, 20, 20, 10, 20, 10, 20)
	fmt.Println(count)
	// if punya slice mau di jadikan var argument
	number := []int{10, 10, 10, 10, 10, 10, 10}
	fmt.Println(sumAll(number...))
	// call func as paramater func
	filter := sanitation
	fmt.Println(callName("gilang", filter))
	filter2 := FilterSan
	fmt.Println(callName2("ajg", filter2))
	// anonymous
	anon := func(name string) bool {
		return name == "anjing"
	}
	callSan("anjing", anon)

	// cara 2
	fmt.Println("cara 2")
	callSan("yoi", func(name string) bool {
		return name == "anjing"
	})

	// defer
	sec()
	// cekError(true)
	cekError2(true)

	// struct
	var obj1 Custom
	obj1.name = "Mae"
	obj1.age = 20
	obj1.call = "suro"
	obj1.city = "Jakarta"
	fmt.Println(obj1)
	// another way inisialitation
	nezuko := Custom{
		name: "nezuko",
		call: "nez",
		age:  20,
		city: "Jakarta",
	}
	fmt.Println(nezuko)
	budi := Custom{"tanjiro", "tanji", "indo", 20}
	fmt.Println(budi)

	// call function with method
	fmt.Println("this is method")
	budi.sayHello("agus")

	// learn interface
	fmt.Println("this is interface")
	tamio := Student{"tamio"}
	fmt.Println(tamio.callmyName())

	// interface kosong
	fmt.Println("this is interface kosong")
	var return1 any = ups()
	fmt.Println(return1)
	fmt.Println("this is interface kosong tes assertions")
	var ran = tesAssert()
	var ran2 = ran.(string)
	fmt.Println(ran2)
	// lebih aman dgn switch jika assertions
	fmt.Println("this is interface kosong tes assertions using switch")
	switch value := return1.(type) {
	case string:
		fmt.Println(value)
	default:
		fmt.Println("cannot assert")

	}

	// pointer

	fmt.Println("learn pointer pass by value ")
	value1 := kota{"dinda", "aceh"}
	value2 := value1

	value2.city = "jakarta"
	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println("learn pointer pass by reference")
	value3 := kota{"miska", "subangan"}
	value4 := &value3

	value4.city = "jakarta"
	fmt.Println(value3)
	fmt.Println(value4)
	fmt.Println(&value3)
	fmt.Println(&value4)
	fmt.Println("learn pointer operator without *")
	address1 := kota{"ani", "jakarta"}
	address2 := &address1
	address2.city = "bandar lampung"
	fmt.Println(address1) //changhe  city from jakarta to bandar lampung in address1 and 2
	address2 = &kota{"dinda", "lombok"}
	address2.city = "bandar"
	fmt.Println(address2) // its a new pointer for address 2
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println("learn pointer operator *")
	address3 := kota{"ani", "jakarta"}
	address4 := &address3
	address4.city = "bandar lampung"
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println("After operator *")
	*address4 = kota{"yuhu", "jakarjakar"}
	fmt.Println(address3)
	fmt.Println(address4)
	address4.city = "link"
	fmt.Println(address3)
	fmt.Println(address4)

	var cekfuncpointer = "joko"
	fmt.Println(cekfuncpointer)
	tesPointer(&cekfuncpointer)
	fmt.Println(cekfuncpointer)

	// method pointer
	didi := funcPointer{"sina"}
	fmt.Println(didi.name)
	didi.married()
	fmt.Println(didi.name)

	// access modifier
	//kalo diawal capital artinya bisa diakses di package lain, sebaliknya

}
