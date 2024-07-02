package handlers

import (
	"net/http"
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

func GetTagByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"msg": "ID parameter is required"})
		return
	}

	tag, err := tagsService.GetTagByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
	}

	utils.JSONResponse(w, http.StatusOK, tag)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	tags, err := tagsService.GetTags(name)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, tags)
}
