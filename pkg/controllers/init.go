package controllers

import (
	"net/http"

	"github.com/sebastian009w/go_template/app/database"
	"github.com/sebastian009w/go_template/pkg/utils"
)

type Username struct {
	ID    int
	Name  string
	Email string
}

func Init(w http.ResponseWriter, r *http.Request) {
	conn := database.MySQL()

	array, err := conn.Query("SELECT * FROM migo;")

	username := Username{}
	arrayUsername := []Username{}

	for array.Next() {
		var id int
		var name, email string

		err = array.Scan(&id, &name, &email)

		if err != nil {
			panic(err.Error())
		}
		username.ID = id
		username.Name = name
		username.Email = email

		arrayUsername = append(arrayUsername, username)
	}
	if err != nil {
		panic(err.Error())
	}

	utils.Templates.ExecuteTemplate(w, "index", nil)

}
