package main

import "fmt"

// - - - - - - - - - - - - - - - - -
// Cage class.

type Cage struct {
	Bear    bool
	Tiger   bool
	Elephan bool
	Turtle  bool
	Zebra   bool
}

func (c Cage) printAllAnimals() {
	fmt.Printf("Bear: %t, Tiger: %t, Elephan: %t, Turtle: %t, Zebra: %t\n", c.Bear, c.Tiger, c.Elephan, c.Turtle, c.Zebra)
}

func (c Cage) isAnimalInCage(animal string) bool {
	switch animal {
	case "bear":
		return c.Bear
	case "tiger":
		return c.Tiger
	case "elephan":
		return c.Elephan
	case "turtle":
		return c.Turtle
	case "zebra":
		return c.Zebra
	}
	return false
}

// - - - - - - - - - - - - - - - - -
// AnimalKeeper class.

type AnimalKeeper struct {
	Cage
}

func (k *AnimalKeeper) catchAnimal(animal string) {
	switch animal {
	case "bear":
		k.Bear = true
		fmt.Println("The bear was caught!")
		break
	case "tiger":
		k.Tiger = true
		fmt.Println("The tiger was caught!")
		break
	case "elephan":
		k.Elephan = true
		fmt.Println("The elephan was caught!")
		break
	case "turtle":
		k.Turtle = true
		fmt.Println("The turtle was caught!")
		break
	case "zebra":
		k.Zebra = true
		fmt.Println("The zebra was caught!")
	}
}

// - - - - - - - - - - - - - - - - -
// Animal class.
type Animal struct {
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

// - - - - - - - - - - - - - - - - -
// Bear class.

type Bear struct {
	Animal
	Type string
}

func (b Bear) say() {
	fmt.Println("Bear's sound")
}

// - - - - - - - - - - - - - - - - -
// Turtle class.

type Turtle struct {
	Animal
	Type string
}

func (t Turtle) say() {
	fmt.Println("Turtle's sound")
}

// - - - - - - - - - - - - - - - - -
// Elephan class.

type Elephan struct {
	Animal
	Type string
}

func (e Elephan) say() {
	fmt.Println("Elephan's sound")
}

// - - - - - - - - - - - - - - - - -
// Tiger class.

type Tiger struct {
	Animal
	Type string
}

func (t Tiger) say() {
	fmt.Println("Tiger's sound")
}

// - - - - - - - - - - - - - - - - -
// Zebra class.

type Zebra struct {
	Animal
	Type string
}

func (z Zebra) say() {
	fmt.Println("Zebra's sound")
}

// - - - - - - - - - - - - - - - - -
// Entry point.

func main() {

	// tiger instance
	tiger := Tiger{
		Animal: Animal{
			Name: "tiger",
			Age:  20,
		},
		Type: "predatory",
	}

	// bear instance

	bear := Bear{
		Animal: Animal{
			Name: "bear",
			Age:  10,
		},
		Type: "predatory",
	}

	fmt.Println(bear)

	// zebra instance

	zebra := Zebra{
		Animal: Animal{
			Name: "zebra",
			Age:  5,
		},
		Type: "equids",
	}

	fmt.Println(zebra)

	// turtle instance

	turtle := Turtle{
		Animal: Animal{
			Name: "turtle",
			Age:  100,
		},
		Type: "reptile",
	}

	fmt.Println(turtle)

	// elephan instance

	elephan := Elephan{
		Animal: Animal{
			Name: "elephan",
			Age:  25,
		},
		Type: "proboscis",
	}

	fmt.Println(elephan)

	// animal keepr instance

	animalKeeper := AnimalKeeper{
		Cage: Cage{
			Bear:    false,
			Tiger:   false,
			Elephan: false,
			Turtle:  false,
			Zebra:   false,
		},
	}
	fmt.Println("Seems like all animals are lost...", animalKeeper)
	animalKeeper.printAllAnimals()

	// Let's try to catch all animals

	animalKeeper.catchAnimal(bear.getName())

	bear.say()

	fmt.Println("the bear is in the zoo:", animalKeeper.isAnimalInCage(bear.getName()))

	fmt.Println("the tiger is in tho zoo:", animalKeeper.isAnimalInCage(tiger.getName()))

	animalKeeper.catchAnimal(tiger.getName())
	tiger.say()

	animalKeeper.catchAnimal(zebra.getName())
	zebra.say()

	animalKeeper.catchAnimal(elephan.getName())
	elephan.say()

	animalKeeper.catchAnimal(turtle.getName())
	turtle.say()

	fmt.Println("Now I think the mission is completed and all animals are in the zoo! Let's double check...")

	animalKeeper.printAllAnimals()

}
