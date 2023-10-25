// Bruce Wayne bio

package main

import (
	"fmt"
)

type (
	age             int
	firstName       string
	lastName        string
	superPower      string
	city            string
	height          int
	weight          int
	side            string
	role            string
	phrase          string
	printFullNameFn func(f firstName, l lastName)
	printPhysiqueFn func(h height, w weight)
)

const (
	CharacterWeight    weight    = 95
	CharacterHeight    height    = 192
	CharacterFirstName firstName = "Bruce"
	CharacterLastName  lastName  = "Wayne"
	CharacterRole      role      = "protector"
	CharacterSide      side      = "Good"
	CharacterCity      city      = "Gotham City"
	FavouritePhrase    phrase    = "Yes! I'm Batman!"
)

func printFullName(first firstName, last lastName) {
	fullName := string(first) + " " + string(last) + "."
	fmt.Println(fullName)
}

func printPhysique(h height, w weight) {
	fmt.Printf("He is %d meters talls and weights is %d kilograms.\n", h, w)
}

func printFullBio(first firstName, last lastName, fullNameFn printFullNameFn, w weight, h height, physiqueFn printPhysiqueFn) {
	fullNameFn(first, last)
	physiqueFn(h, w)
}

func printTitle(characterCity city, characterSide side, characterRole role) {
	fmt.Printf("He is the superhero %s of %s.\n", characterRole, characterCity)
	fmt.Printf("Side: %s\n", characterSide)
}

func printFavouritePhrase(characterPhrase phrase) {
	var characterLogo rune = '🦇'
	fmt.Println(characterPhrase, characterLogo)
}

func main() {
	// This is just for example function composition
	printFullBio(CharacterFirstName, CharacterLastName, printFullName, CharacterWeight, CharacterHeight, printPhysique)
	printTitle(CharacterCity, CharacterSide, CharacterRole)
	printFavouritePhrase(FavouritePhrase)
}
