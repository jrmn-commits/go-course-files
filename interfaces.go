package main

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	name string
	breed string
}

type Cat struct {
	name string
	color string
	numberofTeeth int
}

func main () {
	dog := Dog{
		name: "Fido",
        breed: "Labrador",
	}

	PrintInfo(&dog)


	cat := Cat{
        name: "Neko",
        color: "Black",
        numberofTeeth: 4,
    }

	PrintInfo(&cat)

func PrintInfo(a Animal) {
	fmt.Println(a.Says())
    fmt.Println(a.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Woof!"
}

func (d *Dog) NumberOfLegs() int {
    return 4
}

func (c *Cat) Says() string {
    return "Meow!"
}

func (c *Cat) NumberOfLegs() int {
    return 4
}