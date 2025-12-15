package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	galons    uint8
	ownerInfo owner // gasEngine.ownerInfo.name
	int             // default value 0, gasEngine.int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type engine interface {
	milesLeft() uint8
}

// type gasEngine struct {
// 	mpg       uint8
// 	galons    uint8
// 	owner // gasEngine.name
// }

type owner struct {
	name string
}

// with func (e gasEngine) we define a method for gasEngine type
func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.galons
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

func canMakeIt(e engine, distance uint8) bool {
	return e.milesLeft() >= distance
}

/*
Structs in Go:
	- Custom data types that group related fields
	- Defined using the 'type' and 'struct' keywords
	- Fields can be accessed and modified using dot notation
	- Support for methods, enabling behavior associated with the struct
*/
func structs() {
	var myGasEngine gasEngine = gasEngine{mpg: 30, galons: 10, ownerInfo: owner{name: "John"}, int: 10}
	myGasEngine.mpg = 20
	var myGasEngine2 gasEngine = gasEngine{30, 10, owner{"John"}, 10} // assigned by order of declaration

	fmt.Println(myGasEngine2)
	fmt.Printf("%v's gas engine can go %v miles\n", myGasEngine2.ownerInfo.name, myGasEngine2.mpg*myGasEngine2.galons)

	//anonymous struct definition and instantiation
	var myElectricEngine = struct {
		batteryCapacity uint16
		ownerInfo       owner
	}{400, owner{"Alice"}}

	fmt.Println(myElectricEngine)

	var myGasEngine3 gasEngine = gasEngine{mpg: 25, galons: 8, ownerInfo: owner{name: "Bob"}, int: 5}
	fmt.Printf("%v's gas engine can go %v miles\n", myGasEngine3.ownerInfo.name, myGasEngine3.milesLeft())

	var distance uint8 = 200

	var myGasEngine4 gasEngine = gasEngine{mpg: 30, galons: 10, ownerInfo: owner{name: "John"}, int: 10}
	var myElectricEngine2 electricEngine = electricEngine{mpkwh: 4, kwh: 50}

	if canMakeIt(myGasEngine4, distance) {
		fmt.Printf("%v's gas engine can make it %v miles\n", myGasEngine4.ownerInfo.name, distance)
	} else {
		fmt.Printf("%v's gas engine cannot make it %v miles\n", myGasEngine4.ownerInfo.name, distance)
	}

	if canMakeIt(myElectricEngine2, distance) {
		fmt.Printf("The electric engine can make it %v miles\n", distance)
	} else {
		fmt.Printf("The electric engine cannot make it %v miles\n", distance)
	}
}
