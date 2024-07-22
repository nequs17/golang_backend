package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// getImageHandler handles the image retrieval
// @Summary Get an uploaded image
// @Description Retrieve the uploaded image
// @Produce  image/png
// @Success 200 {string} string "Image retrieved successfully"
// @Failure 404 {string} string "Image not found"
// @Router /api/test/image [get]
func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	imageData, err := ioutil.ReadFile("uploaded_image.png")
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(imageData)
}

// getImageBase64Handler handles the image retrieval and encodes it in base64
// @Summary Get an uploaded image in base64
// @Description Retrieve the uploaded image and return it encoded in base64
// @Produce  json
// @Success 200 {object} ImagePayload "Base64 encoded image"
// @Failure 404 {string} string "Image not found"
// @Router /api/sockets/image/base64 [get]
func GetImageBase64Handler(w http.ResponseWriter, r *http.Request) {
	imageData, err := ioutil.ReadFile("uploaded_image.png")
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}

	encodedImage := base64.StdEncoding.EncodeToString(imageData)
	payload := ImagePayload{
		Image: encodedImage,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
