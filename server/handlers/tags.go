package handlers

import (
	"net/http"
	"strings"
	"sugoi-api/services"
	"sugoi-api/utils"

	"github.com/go-chi/chi/v5"
)

var tagsService = services.NewTagsService()

func CreateTag(w http.ResponseWriter, r *http.Request) {
	tag, err := tagsService.CreateTag(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
	}

	utils.JSONResponse(w, http.StatusCreated, &tag)
}

func GetTagByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	name = strings.Join(strings.Split(name, "-"), " ")
	name = utils.CapitalizeWords(name)

	tag, err := tagsService.GetTagByName(name)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
	}

	utils.JSONResponse(w, http.StatusOK, tag)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := tagsService.GetTags()
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, tags)
}
