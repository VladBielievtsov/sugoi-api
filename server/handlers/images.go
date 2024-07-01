package handlers

import (
	"net/http"
	"strconv"
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
	tag := r.URL.Query().Get("tag")
	character := r.URL.Query().Get("character")

	images, err := imagesService.GetImages(r, tag, character)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, images)
}

func GetImageByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"msg": "ID parameter is required"})
		return
	}

	image, err := imagesService.GetImageByID(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, image)
}

func GetImagesTags(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"msg": "ID parameter is required"})
		return
	}

	tags, err := imagesService.GetImagesTags(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, tags)
}

func GetRandomImages(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	queryLimit, _ := strconv.Atoi(q.Get("limit"))
	tag := r.URL.Query().Get("tag")
	character := r.URL.Query().Get("character")

	limit := 5
	if queryLimit > 0 {
		limit = queryLimit
	}

	images, errs := imagesService.GetRandomImages(limit, tag, character)
	if errs != nil {
		utils.JSONResponse(w, http.StatusNotFound, errs)
		return
	}
	utils.JSONResponse(w, http.StatusOK, images)
}

func DeleteImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"msg": "ID parameter is required"})
		return
	}

	image, err := imagesService.DeleteImage(id)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err["msg"])
		return
	}

	utils.JSONResponse(w, http.StatusOK, image)
}
