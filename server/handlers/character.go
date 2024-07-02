package handlers

import (
	"encoding/json"
	"net/http"
	"sugoi-api/services"
	"sugoi-api/types"
	"sugoi-api/utils"

	"github.com/go-chi/chi/v5"
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

func GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"msg": "ID parameter is required"})
		return
	}

	character, err := charactersService.GetCharacterByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
	}

	utils.JSONResponse(w, http.StatusOK, character)
}

func UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"msg": "ID parameter is required"})
		return
	}

	var req types.CreateCharacterBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	character, err := charactersService.UpdateCharacter(id, req.Name, req.Description, req.Gender, req.Species)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, character)
}
