package cmd

import "github.com/spf13/cobra"
import "fmt"

// Gestionnaire de commandes disponibles
var devoirCmd = &cobra.Command{
	Use:   "devoir",
	Short: "Programme de commandes",
	Long: `Liste des commandes`,
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("Commande disponibles :")
	},
}

// Première commande qui affiche un
var devoir1Cmd = &cobra.Command{
	Use:   "un",
	Short: "Progamme devoir un",
	Long: `Voir le programme devoir un`,
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("un")
	},
}

// Deuxième commande qui affiche deux
var devoir2Cmd = &cobra.Command{
	Use:   "deux",
	Short: "Programme devoir un",
	Long: `Voir le programme devoir deux`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deux")
	},
}

func init() {
	devoirCmd.AddCommand(devoir1Cmd)
	devoirCmd.AddCommand(devoir2Cmd)
}

func Execute() {
	devoirCmd.Execute()
}

func main () {
	Execute()
}