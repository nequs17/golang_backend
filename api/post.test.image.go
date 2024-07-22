package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ImagePayload struct {
	Image string `json:"image"`
}

// uploadHandler handles the image upload
// @Summary Upload an image
// @Description Upload an image in base64 format
// @Accept  json
// @Produce  json
// @Param   image body ImagePayload true "Base64 encoded image"
// @Success 200 {string} string "Image uploaded successfully"
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Unable to save the image"
// @Router /api/test/upload [post]
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var payload ImagePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	imageData, err := base64.StdEncoding.DecodeString(payload.Image)
	if err != nil {
		http.Error(w, "Invalid base64 string", http.StatusBadRequest)
		return
	}

	err = ioutil.WriteFile("uploaded_image.png", imageData, 0644)
	if err != nil {
		http.Error(w, "Unable to save the image", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Image uploaded successfully")
}
