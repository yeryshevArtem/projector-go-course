package main

import "fmt"

type color string

const (
	ColorReset color = "\033[0m"
	ColorGreen color = "\u001b[32m"
)

func (c color) String() string {
	return string(c)
}

func textInColor(t string, c color) {
	fmt.Printf("%s%s\n"+ColorReset.String(), c, t)
}

// - - - - - - - - - - - - - - - - -
// Cage class.

type Cage struct {
	Locked bool
}

// - - - - - - - - - - - - - - - - -
// Animal class.
type Animal struct {
	Cage
	Age  int
	Name string
}

func (a Animal) getAge() int {
	return a.Age
}

func (a *Animal) setAge(age int) {
	a.Age = age
}

func (a Animal) getName() string {
	return a.Name
}

func (a *Animal) escape() {
	a.Cage.Locked = false
}

// - - - - - - - - - - - - - - - - -
// AnimalKeeper class.

type AnimalKeeper struct {
}

func (ak *AnimalKeeper) catch(a *Animal) {
	a.Cage.Locked = true
	fmt.Printf("The %s was caught\n", a.Name)
}

// - - - - - - - - - - - - - - - - -
// Entry point.

func main() {

	bear := &Animal{
		Name: "bear",
		Age:  10,
		Cage: Cage{
			Locked: true,
		},
	}

	tiger := &Animal{
		Name: "tiger",
		Age:  20,
		Cage: Cage{
			Locked: true,
		},
	}

	elephan := &Animal{
		Name: "turtle",
		Age:  100,
		Cage: Cage{
			Locked: true,
		},
	}

	turtle := &Animal{
		Name: "elephan",
		Age:  100,
		Cage: Cage{
			Locked: true,
		},
	}

	zebra := &Animal{
		Name: "zebra",
		Age:  5,
		Cage: Cage{
			Locked: true,
		},
	}

	textInColor(`Now looks like all animals are in the cage...`, ColorGreen)
	fmt.Println("Bear is in the cage: ", bear.Cage.Locked)
	fmt.Println("Tiger is in the cage: ", tiger.Cage.Locked)
	fmt.Println("Elephan is in the cage: ", elephan.Cage.Locked)
	fmt.Println("Turtle is in the cage: ", turtle.Cage.Locked)
	fmt.Println("Zebra is in the cage: ", zebra.Cage.Locked)

	// Then something happens and all animals escaped...
	zebra.escape()

	tiger.escape()

	bear.escape()

	turtle.escape()

	elephan.escape()

	fmt.Printf("Hello, my name is %s\n", elephan.getName())

	fmt.Printf("Hello, my name is %s\n", turtle.getName())

	fmt.Printf("Hello, my name is %s\n", bear.getName())

	fmt.Printf("Hello, my name is %s\n", tiger.getName())

	fmt.Printf("Hello, my name is %s\n", zebra.getName())

	textInColor(`Omg... Where are all our animals???`, ColorGreen)
	fmt.Println("Bear is in the cage: ", bear.Cage.Locked)
	fmt.Println("Tiger is in the cage: ", tiger.Cage.Locked)
	fmt.Println("Elephan is in the cage: ", elephan.Cage.Locked)
	fmt.Println("Turtle is in the cage: ", turtle.Cage.Locked)
	fmt.Println("Zebra is in the cage: ", zebra.Cage.Locked)

	textInColor(`We need to ask animal keeper to help us!`, ColorGreen)

	textInColor(`Animal keeper is here!`, ColorGreen)

	animalKeeper := AnimalKeeper{}

	animalKeeper.catch(elephan)
	animalKeeper.catch(tiger)
	animalKeeper.catch(turtle)
	animalKeeper.catch(bear)
	animalKeeper.catch(zebra)

	textInColor(`Seems like all animals were caught!`, ColorGreen)

	fmt.Println("Bear is in the cage: ", bear.Cage.Locked)
	fmt.Println("Tiger is in the cage: ", tiger.Cage.Locked)
	fmt.Println("Elephan is in the cage: ", elephan.Cage.Locked)
	fmt.Println("Turtle is in the cage: ", turtle.Cage.Locked)
	fmt.Println("Zebra is in the cage: ", zebra.Cage.Locked)

}
