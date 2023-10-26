package main

import "fmt"

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
		fmt.Println("The bear was caught!")
		return c.Bear
		break
	case "tiger":
		return c.Tiger
		break
	case "elephan":
		return c.Elephan
		break
	case "turtle":
		return c.Turtle
		break
	case "zebra":
		return c.Zebra
		break
	}
	return false
}

// We will indicate if the animal keeper catch the animal not doesn't
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

type Animal struct {
	Age int
}

func (a *Animal) getAge() int {
	return a.Age
}

type Bear struct {
	Animal
}

type Turtle struct {
	Animal
}

type Elephan struct {
	Animal
}

type Tiger struct {
	Animal
}

type Zebra struct {
	Animal
}

func main() {

	animalKeeper := AnimalKeeper{
		Cage: Cage{
			Bear:    false,
			Tiger:   false,
			Elephan: false,
			Turtle:  false,
			Zebra:   false,
		},
	}
	// fmt.Println("Seems like all animals are lost...", animalKeeper)
	animalKeeper.catchAnimal("bear")
	// fmt.Println("Good catch! I found the bear", animalKeeper)

	// animalKeeper.printAllAnimals()

	fmt.Println(animalKeeper.isAnimalInCage("bear"))

	fmt.Println(animalKeeper.isAnimalInCage("tiger"))

}
