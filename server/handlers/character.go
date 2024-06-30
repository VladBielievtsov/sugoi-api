package handlers

import (
	"net/http"
	"sugoi-api/services"
	"sugoi-api/utils"
)

var charactersService = services.NewCharactersService()

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	tag, err := charactersService.CreateCharacter(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
	}

	utils.JSONResponse(w, http.StatusCreated, &tag)
}

func GetCharacters(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	gender := q.Get("gender")
	species := q.Get("species")

	characters, err := charactersService.GetCharacters(name, gender, species)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, characters)
}
