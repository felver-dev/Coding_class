package main

import "fmt"

func main() {
	// Affiche deux variables en les concatenant
	var name = "Jules"
	var firstName = "Felver"

	fmt.Println("Hello", name+" "+firstName)

	// LES TYPES DES VARIABLES :
	/*
		Un type qualifie la nature de la donnée passée à une variable.
		Ex :
			var age = 24
			var prenom = "Jules"
			var PI = 3.14

			Nous avons les entiers, les chaines de caractères, les booléens (True ou false), les nombres flottants : Nombres à virgules

			var isAdult = false

			demander à l'utilisateur de rentrer son âge. Si son âge est strictement inférieur à 18 ans alors isAdult = true ,
	*/

	// LES ENTIERS : INTEGERS | INT
	var age = 25
	// var years = "ans"
	fmt.Println("I am", age, "years old")

	// Il existe des entiers signés et les entiers non
	/*
		int/uint : 8bits 16bits, b32 bits ou de 64 bits
		Ex : soit int32 ou int64  ou encore uint32 ou uint64
	*/

	var numberOfWives uint
	numberOfWives = 4

	// var numberOfWives = 4
	var PI float32
	PI = 3.14

	var universityName string
	universityName = "FSSA"

	var isStudent bool
	isStudent = true

	fmt.Println("Nombre des femmes : ", numberOfWives)
	fmt.Println("La valeur de PI est : ", PI)
	fmt.Println("Je suis étudiant à la : ", universityName)
	fmt.Println(isStudent, "I am a student")

	// LES OPERATEURS : Sont des signaux mathématiques nous permettant de faire des calculs ou de la logique
	/*
		L'addition : +
		La soustraction : -
		La multiplication : *
		La division : /
		Le module : %

		Les opérateurs logiques
		> : strictement supérieur
		>= :  supérieur ou égal
		< : strictement inférieur
		<= : inférieur ou égal
		== : Egalité
		!== : L'inégalité ou la différence
		&& : ET logique          AND
		|| : Ou logique 		 OR
		! : Le contraire logique NOT
	*/

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("*************************************************")
	fmt.Println("*                LES OPERATEURS                 *")
	fmt.Println("*************************************************")

	var x int
	var y int

	var A int
	var S int
	var D int
	var M int
	var Mod int

	x = 10
	y = 20
	A = x + y
	S = x - y
	D = y / x
	M = x * y
	Mod = y % x

	fmt.Println("La valeur de A est : ", A)
	fmt.Println("La valeur de S est : ", S)
	fmt.Println("La valeur de D est : ", D)
	fmt.Println("La valeur de M est : ", M)
	fmt.Println("La valeur de Mod est : ", Mod)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	fmt.Println("*************************************************")
	fmt.Println("*           LES STRUCTURES DE CONTROLES         *")
	fmt.Println("*************************************************")
	// If, else, else if, switch

	var isAdult bool

	// Example : Miguel -> 20 ans et n'est pas marié
	// var nameOfUser string
	// nameOfUser = "Miguel"
	// isAdult = false
	isAdult = false
	var miguelNumberOfWives uint8
	// miguelNumberOfWives = 0
	// miguelNumberOfWives = 1
	miguelNumberOfWives = 1

	// if isAdult == true {
	// 	if miguelNumberOfWives == 0 {
	// 		fmt.Println(nameOfUser, "n'est pas marié")
	// 	} else if miguelNumberOfWives == 1 {
	// 		fmt.Println(nameOfUser, "est monogame")
	// 	} else {
	// 		fmt.Println(nameOfUser, "est polygame")
	// 	}
	// } else {
	// 	fmt.Println(nameOfUser, "tu n'est pas adulte, rentre étudier tes leçons 😒")
	// }

	// alt gr + 6 : le pipe
	if isAdult && miguelNumberOfWives == 0 {
		fmt.Println("Hey, Miguel! il est temps de te marier 😁")
	} else if isAdult && miguelNumberOfWives == 1 {
		fmt.Println("Super félicitations. Vous faites un très beau couple 😍")
		fmt.Println("Vous êtes conviés à la réunion")
	} else if isAdult && miguelNumberOfWives >= 2 {
		fmt.Println("Cher ami, la polygamie est contraire à la Parole ❌")
	} else if !isAdult {
		fmt.Println("Vous n'est pas conviés à la réunion")
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}
