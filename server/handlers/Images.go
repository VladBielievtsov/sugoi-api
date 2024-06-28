package handlers

import (
	"net/http"
	"sugoi-api/services"
	"sugoi-api/utils"
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
	images, err := imagesService.GetImages()
	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, err)
	}

	utils.JSONResponse(w, http.StatusOK, images)
}
