package upload

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UploadHandler struct {
	uploadServices *UploadServices
}

func NewUploadHandler(s *UploadServices) *UploadHandler {
	return &UploadHandler{uploadServices: s}
}

func (h *UploadHandler) RegisterRoutes(r chi.Router) {
	r.Post("/", h.UploadFile)

	// s.Router.Post("/upload", uploadHandler.UploadFile)
}

// UploadFile godoc
// @Summary Upload a file
// @Description Upload an image/file and return public URL
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /upload [post]
func (h *UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath, err := h.uploadServices.SaveFile(file, header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": map[string]string{
			"url": filePath,
		},
	})
}
