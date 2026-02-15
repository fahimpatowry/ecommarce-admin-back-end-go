package category

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryInput struct {
	URL  string `bson:"url" json:"url"`
	Slug string `bson:"slug" json:"slug"`
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

// @Summary Get all categorys
// @Description Retrieve all category items
// @Tags categorys
// @Produce json
// @Success 200 {array} Category
// @Failure 500 {string} string "Internal Server Error"
// @Router /categorys [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	categorys, err := h.service.GetCategorys(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categorys)
}

// @Summary Create a new category
// @Description Add a new category item
// @Tags categorys
// @Accept json
// @Produce json
// @Param category body CategoryInput true "Category object"
// @Success 201 {object} Category
// @Failure 400 {string} string "Bad Request"
// @Router /categorys [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var c Category

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateCategory(r.Context(), &c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// @Summary Update category by ID
// @Description Update an existing category
// @Tags categorys
// @Accept json
// @Produce json
// @Param id path string true "category ID"
// @Param category body Category true "Category object"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /categorys/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var c Category
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateCategory(r.Context(), id, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "updated successfully"})
}

// @Summary Delete category by ID
// @Description Delete a category item
// @Tags categorys
// @Param id path string true "category ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /categorys/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteCategory(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
