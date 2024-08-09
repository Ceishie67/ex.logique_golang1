// Le but de cet exercice est de créer un motif grâce à des symboles comme * ou - par exemple

package main

import (
	"flag"
	"fmt"
	"strings"
)

func print_diamond(size int, motif string) {

	// Le motif a representer est le même
	// Mais pour cette partie suivant il faut modifier la taille du diamant
	// Mais aussi le symbole avec lesquels vous le représenter

	// Partie supérieure du diamant
	spaces := strings.Repeat(" ", (size - (size - 1)))
	stars := strings.Repeat(motif, size*2-1-(2*(size-(size-1)))-2)
	fmt.Println(spaces + stars)

	// Partie inférieure du diamant

	for i := 1; i < size; i++ {
		spaces := strings.Repeat(" ", i-1)
		stars := strings.Repeat(motif, size*2-1-(2*i))
		fmt.Println(spaces + stars)
	}
	//ajouter la gestion d'erreurs inférieur à 4
}

func main() {
	// Définir les flags de la ligne de commande
	size := flag.Int("size", 4, "taille du diamant")
	motif := flag.String("motif", "*", "motif à utiliser pour le diamant")

	// Alias pour les flags
	flag.IntVar(size, "s", 4, "taille du diamant")
	flag.StringVar(motif, "m", "*", "motif à utiliser pour le diamant")

	// Parse les flags
	flag.Parse()

	// Appeler la fonction print_diamond avec les valeurs des flags
	print_diamond(*size, *motif)
}
