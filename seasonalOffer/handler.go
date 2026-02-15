package seasonalOffer

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SeasonalOfferInput struct {
	URL      string `json:"url" bson:"url"`
	Slug     string `json:"slug" bson:"slug"`
	IsActive bool   `json:"isActive" bson:"isActive"`
	Position int    `bson:"position" json:"position"`
}

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
	r.Get("/", h.GetAll)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}

// @Summary Get all seasonalOffers
// @Description Retrieve all seasonalOffer items
// @Tags seasonalOffers
// @Produce json
// @Success 200 {array} SeasonalOffer
// @Failure 500 {string} string "Internal Server Error"
// @Router /seasonalOffers [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	SeasonalOffers, err := h.service.GetSeasonalOffers(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(SeasonalOffers)
}

// @Summary Create a new seasonalOffer
// @Description Add a new seasonalOffer item
// @Tags seasonalOffers
// @Accept json
// @Produce json
// @Param seasonalOffer body SeasonalOffer true "SeasonalOffer object"
// @Success 201 {object} SeasonalOffer
// @Failure 400 {string} string "Bad Request"
// @Router /SeasonalOffers [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var c SeasonalOffer

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateSeasonalOffer(r.Context(), &c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// @Summary Update seasonalOffer by ID
// @Description Update an existing seasonalOffer
// @Tags seasonalOffers
// @Accept json
// @Produce json
// @Param id path string true "SeasonalOffer ID"
// @Param seasonalOffer body SeasonalOffer true "SeasonalOffer object"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /seasonalOffers/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var c SeasonalOffer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateSeasonalOffer(r.Context(), id, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "updated successfully"})
}

// @Summary Delete seasonalOffer by ID
// @Description Delete a seasonalOffer item
// @Tags seasonalOffers
// @Param id path string true "SeasonalOffer ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /seasonalOffers/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteSeasonalOffer(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
