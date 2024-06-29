package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"sugoi-api/services"
	"sugoi-api/utils"

	"github.com/go-chi/chi/v5"
)

var imagesService = services.NewImagesService()

func StoreImage(w http.ResponseWriter, r *http.Request) {
	img, err := imagesService.CreateImage(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err)
	}

	utils.JSONResponse(w, http.StatusCreated, &img)
}

func GetImages(w http.ResponseWriter, r *http.Request) {
	images, err := imagesService.GetImages(r)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, images)
}

func GetImageByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	image, err := imagesService.GetImageByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, image)
}

func GetRandomImages(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		Limit int `json:"limit"`
	}

	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil && err != io.EOF {
		utils.JSONResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	limit := 5
	if requestBody.Limit > 0 {
		limit = requestBody.Limit
	}

	images, errs := imagesService.GetRandomImages(limit)
	if errs != nil {
		utils.JSONResponse(w, http.StatusNotFound, errs)
		return
	}
	utils.JSONResponse(w, http.StatusOK, images)
}
