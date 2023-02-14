package main

import (
	"net/http"
	"database/sql"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	CodUser  int    `json:"codUser"`
	Username string `json:"username"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	// GET /:usuario
	r.GET("/api/:usuario", func (c *gin.Context) {
		var u Usuario

		err := db.QueryRow("SELECT cod_user, username FROM TodoUsuarios WHERE username = ?;",
		c.Param("usuario")).Scan(&u.CodUser, &u.Username)
		if err != nil {
			log.Fatal(err)
		}

		c.IndentedJSON(http.StatusOK, u)
	})

	// POST /
	// UPDATE /:usuario/:tarefa
	// UPDATE /:usuario
	// DELETE /:usuario/:tarefa
	// DELETE /:usuario

	r.Run(":4000")

}