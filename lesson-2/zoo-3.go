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
	Name   string
	Locked bool
}

// - - - - - - - - - - - - - - - - -
// Animal class.
type Animal struct {
	*Cage
	Name string
}

func (a Animal) isLocked() bool {
	fmt.Printf("The %s is in the cage: ", a.Name)
	return a.Cage.Locked
}

func (a *Animal) escape() {
	a.Cage.Locked = false
	fmt.Printf("The %s escaped\n", a.Name)
}

// - - - - - - - - - - - - - - - - -
// AnimalKeeper class.

type AnimalKeeper struct {
	TigerCage *Cage
	BearCage  *Cage
}

func (ak *AnimalKeeper) catch(a *Animal) {
	a.Cage.Locked = true
	fmt.Printf("The animal keeper caught %s\n", a.Name)
}

// - - - - - - - - - - - - - - - - -
// Entry point.

func main() {
	bearCage := Cage{
		Name:   "bear",
		Locked: true,
	}

	tigerCage := Cage{
		Name:   "tiger",
		Locked: true,
	}

	animalKeeper := AnimalKeeper{
		TigerCage: &tigerCage,
		BearCage:  &bearCage,
	}

	bear := &Animal{
		Name: "bear",
		Cage: &bearCage,
	}

	tiger := &Animal{
		Name: "tiger",
		Cage: &tigerCage,
	}

	// Bear
	fmt.Println(bear.isLocked())
	bear.escape()
	fmt.Println(bear.isLocked())

	// Tiger
	fmt.Println(tiger.isLocked())
	tiger.escape()
	fmt.Println(tiger.isLocked())

	animalKeeper.catch(bear)
	fmt.Println(bear.isLocked())

}
