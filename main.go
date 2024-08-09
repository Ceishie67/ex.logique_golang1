// Le but de cet exercice est de créer un motif grâce à des symboles comme * ou - par exemple

package main

import "fmt"

func main() {
	// main excute la fonction print_diamond
	print_diamond()

}

func print_diamond() {
	// Le motif a représenter dans cet exercice est un diamant
	// le code doit afficher dans le terminal :
	//	*****
	// ********
	//	*****
	//	 ***
	//	  *
	fmt.Println(" *****\n********\n *****\n  ***\n   *")

}
