package product

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductInput struct {
	title      string   `bson:"title" json:"title"`
	Decription string   `bson:"description" json:"description"`
	URL        []string `json:"url" bson:"url"`
	Price      float64  `bson:"price" json:"price"`
	Discount   float64  `bson:"discount" json:"discount"`
	Tag        string   `bson:"tag" json:"tag"`
	IsPopular  bool     `bson:"isPopular" json:"isPopular"`
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

// @Summary Get all products
// @Description Retrieve all product items
// @Tags products
// @Produce json
// @Success 200 {array} Product
// @Failure 500 {string} string "Internal Server Error"
// @Router /products [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	carousels, err := h.service.GetProducts(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(carousels)
}

// @Summary Create a new product
// @Description Add a new product item
// @Tags products
// @Accept json
// @Produce json
// @Param product body ProductInput true "Product object"
// @Success 201 {object} Product
// @Failure 400 {string} string "Bad Request"
// @Router /products [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var c Product

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProduct(r.Context(), &c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

// @Summary Update product by ID
// @Description Update an existing product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body Product true "Product object"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var c Product
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateProduct(r.Context(), id, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "updated successfully"})
}

// @Summary Delete product by ID
// @Description Delete a product item
// @Tags products
// @Param id path string true "Product ID"
// @Success 204 {string} string "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProduct(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
