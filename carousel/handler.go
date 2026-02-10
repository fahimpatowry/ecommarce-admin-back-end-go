package carousel

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarouselInput struct {
    URL      string `json:"url" bson:"url"`
    Slug     string `json:"slug" bson:"slug"`
    IsActive bool   `json:"isActive" bson:"isActive"`
}

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

// func NewHandler(s *Service) *Handler {
// 	return &Handler{service: s}
// }

func (h *Handler) RegisterRoutes(r chi.Router){
	r.Get("/", h.GetAll)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}

// @Summary Get all carousels
// @Description Retrieve all carousel items
// @Tags carousels
// @Produce json
// @Success 200 {array} Carousel
// @Failure 500 {string} string "Internal Server Error"
// @Router /carousels [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	carousels, err := h.service.GetCarousels(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(carousels)
}

// @Summary Create a new carousel
// @Description Add a new carousel item
// @Tags carousels
// @Accept json
// @Produce json
// @Param carousel body CarouselInput true "Carousel object"
// @Success 201 {object} Carousel
// @Failure 400 {string} string "Bad Request"
// @Router /carousels [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request){
	var c Carousel

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCarousel(r.Context(), &c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// @Summary Update carousel by ID
// @Description Update an existing carousel
// @Tags carousels
// @Accept json
// @Produce json
// @Param id path string true "Carousel ID"
// @Param carousel body Carousel true "Carousel object"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /carousels/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var c Carousel
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil{
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCarousel(r.Context(), id, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "updated successfully"})
}

// @Summary Delete carousel by ID
// @Description Delete a carousel item
// @Tags carousels
// @Param id path string true "Carousel ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /carousels/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCarousel(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}