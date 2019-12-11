package main

// Imports cobra et fmt
import "github.com/spf13/cobra"
import "database/sql"
import _"github.com/mattn/go-sqlite3"

const driver = "sqlite3"
const path = "commande.db"

// Création d'une db avec mysql
var devoirCmd = &cobra.Command{
	Use:   "devoir",
	Short: "Programme de commandes",
	Long: `Liste des commandes`,
	Run: func (cmd *cobra.Command, args []string) {
	db, _ := sql.Open(driver, path)
	query :=
			`
			CREATE TABLE IF NOT EXISTS documents (
			id INTEGER NOT NULL PRIMARY KEY,
			header TEXT,
			footer TEXT );
			CREATE TABLE IF NOT EXISTS lignes (
			id INTEGER NOT NULL PRIMARY KEY,
			document INT(3),
			tarif INT(255),
			descritpion TEXT,
			CONSTRAINT fk_ligne_document FOREIGN KEY (document) REFERENCE documents (id)
			ON UPDATE CASCADE ON DELETE CASCADE );
			`
	_, _ = db.Exec(query)
	}}

// Pour git push sur le lab je n'ai pas mon code eemi et via SSH un peu compliquer ne l'ayant jamais fait