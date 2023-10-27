package main

import "fmt"

// - - - - - - - - - - - - - - - - -
// Animal class.
type Animal struct {
	Age    int
	Name   string
	InCage bool
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
	a.InCage = false
}

// - - - - - - - - - - - - - - - - -
// Bear class.

type Bear struct {
	*Animal
}

func (b Bear) say() {
	fmt.Println("Bear's sound")
}

// - - - - - - - - - - - - - - - - -
// Turtle class.

type Turtle struct {
	*Animal
}

func (t Turtle) say() {
	fmt.Println("Turtle's sound")
}

// - - - - - - - - - - - - - - - - -
// Elephan class.

type Elephan struct {
	*Animal
}

func (e Elephan) say() {
	fmt.Println("Elephan's sound")
}

// - - - - - - - - - - - - - - - - -
// Tiger class.

type Tiger struct {
	*Animal
}

func (t Tiger) say() {
	fmt.Println("Tiger's sound")
}

// - - - - - - - - - - - - - - - - -
// Zebra class.

type Zebra struct {
	*Animal
}

func (z Zebra) say() {
	fmt.Println("Zebra's sound")
}

// - - - - - - - - - - - - - - - - -
// Cage class.

type Cage struct {
	Bear    *Bear
	Tiger   *Tiger
	Elephan *Elephan
	Turtle  *Turtle
	Zebra   *Zebra
}

// - - - - - - - - - - - - - - - - -
// AnimalKeeper class.

type AnimalKeeper struct {
}

func (ak *AnimalKeeper) catchZebra(z *Zebra) {
	z.InCage = true
}

func (ak *AnimalKeeper) catchTurtle(t *Turtle) {
	t.InCage = true
}

func (ak *AnimalKeeper) catchBear(b *Bear) {
	b.InCage = true
}

func (ak *AnimalKeeper) catchTiger(t *Tiger) {
	t.InCage = true
}

func (ak *AnimalKeeper) catchElephan(e *Elephan) {
	e.InCage = true
}

// - - - - - - - - - - - - - - - - -
// Entry point.

func main() {

	bear := &Bear{
		Animal: &Animal{
			Name:   "bear",
			Age:    10,
			InCage: true,
		},
	}

	tiger := &Tiger{
		Animal: &Animal{
			Name:   "tiger",
			Age:    20,
			InCage: true,
		},
	}

	elephan := &Elephan{
		Animal: &Animal{
			Name:   "elephan",
			Age:    25,
			InCage: true,
		},
	}

	turtle := &Turtle{
		Animal: &Animal{
			Name:   "turtle",
			Age:    100,
			InCage: true,
		},
	}

	zebra := &Zebra{
		Animal: &Animal{
			Name:   "zebra",
			Age:    5,
			InCage: true,
		},
	}

	// On this step all our animals are in the cage
	cage := &Cage{
		Bear:    bear,
		Tiger:   tiger,
		Elephan: elephan,
		Turtle:  turtle,
		Zebra:   zebra,
	}

	fmt.Println("Now looks like all animals are in the cage...")
	fmt.Println("Bear is in cage: ", cage.Bear.InCage)
	fmt.Println("Tiger is in cage: ", cage.Tiger.InCage)
	fmt.Println("Elephan is in cage: ", cage.Elephan.InCage)
	fmt.Println("Turtle is in cage: ", cage.Turtle.InCage)
	fmt.Println("Zebra is in cage: ", cage.Zebra.InCage)

	// Then something happens and all animals escaped...
	zebra.escape()

	tiger.escape()

	bear.escape()

	turtle.escape()

	elephan.escape()

	fmt.Println("Omg... Where are all our animals...")
	fmt.Println("Bear is in cage: ", cage.Bear.InCage)
	fmt.Println("Tiger is in cage: ", cage.Tiger.InCage)
	fmt.Println("Elephan is in cage: ", cage.Elephan.InCage)
	fmt.Println("Turtle is in cage: ", cage.Turtle.InCage)
	fmt.Println("Zebra is in cage: ", cage.Zebra.InCage)

	fmt.Println("We need to ask animalkeeper to help us!")

	animalKeeper := &AnimalKeeper{}

	animalKeeper.catchTiger(tiger)
	animalKeeper.catchBear(bear)
	animalKeeper.catchElephan(elephan)
	animalKeeper.catchZebra(zebra)
	animalKeeper.catchTurtle(turtle)

	fmt.Println("Omg... Where are all our animals...")
	fmt.Println("Bear is in cage: ", cage.Bear.InCage)
	fmt.Println("Tiger is in cage: ", cage.Tiger.InCage)
	fmt.Println("Elephan is in cage: ", cage.Elephan.InCage)
	fmt.Println("Turtle is in cage: ", cage.Turtle.InCage)
	fmt.Println("Zebra is in cage: ", cage.Zebra.InCage)

}
