package images

import (
	"encoding/json"
	"fmt"
	"net/http"

	"weproov/internal/helper"

	"github.com/gorilla/mux"
)



type ImagesRoutes struct {
		r *mux.Router
		imgs *ImageServices

}

type image struct {
	ImagekeyName string
}

type getImageResponse struct {
	ImageUrl string

}

func NewImagesRoutes( r *mux.Router, imgs *ImageServices) *ImagesRoutes {
	return &ImagesRoutes{
		r : r,
		imgs : imgs,
	}
}

func (i *ImagesRoutes) Routes() {
	i.r.HandleFunc("/image", i.createNewImageHandler).Methods("POST")
	i.r.HandleFunc("/image", i.deleteImageHandler).Methods("DELETE")
	i.r.HandleFunc("/image", i.getImageHandler).Methods("GET")

}

func (i *ImagesRoutes) createNewImageHandler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("data-type")
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	defer file.Close()

	keyName, err := i.imgs.CreateImage(handler.Filename, headerContentTtype, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(keyName)
	resp := &image{
		ImagekeyName : keyName,
	}

	json.NewEncoder(w).Encode(resp)

}

func (i *ImagesRoutes) deleteImageHandler(w http.ResponseWriter, r *http.Request) {
	imgResponse := new(image)
	err := helper.ParseIncomingInput(w, r, imgResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = i.imgs.DeleteImage(imgResponse.ImagekeyName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)

}

func (i *ImagesRoutes) getImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("name")
	if len(imageName) == 0 {
		http.Error(w, "empty string", http.StatusBadRequest)
		return
	}

	imageUrl, err := i.imgs.GetImageByKey(imageName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := &getImageResponse{
		ImageUrl : imageUrl,
	}
	json.NewEncoder(w).Encode(resp)
}


