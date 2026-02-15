package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"

	// health "initial_project/internal/handlers/healthCheck"
	carousel "initial_project/carousel"
	health "initial_project/handlers"
	seasonalOffer "initial_project/seasonalOffer"
	"initial_project/upload"
)

type Server struct {
	Router *chi.Mux
	DB     *mongo.Database
}

func NewServer(db *mongo.Database) *Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	return &Server{
		Router: r,
		DB:     db,
	}
}

// Add routes here
func (s *Server) Routes() {
	s.Router.Get("/health", health.HealthCheck)

	// --- REGISTER CAROUSEL ROUTES ---
	carouselRepo := carousel.NewRepository(s.DB)
	carouselService := carousel.NewService(carouselRepo)
	carouselHandler := carousel.NewHandler(carouselService)

	// all carousel routes under /carousels
	s.Router.Route("/carousels", carouselHandler.RegisterRoutes)

	// --- REGISTER SEASONAL OFFER ROUTES ---
	seasonalOfferRepo := seasonalOffer.NewRepository(s.DB)
	seasonalOfferService := seasonalOffer.NewService(seasonalOfferRepo)
	seasonalOfferHandler := seasonalOffer.NewHandler(seasonalOfferService)

	// all carousel routes under /seasonalOffer
	s.Router.Route("/seasonalOffers", seasonalOfferHandler.RegisterRoutes)

	// --- REGISTER UPLOAD ROUTES ---
	uploadService := upload.NewUploadService()
	uploadHandler := upload.NewUploadHandler(uploadService)

	// all upload routes under /carousels
	s.Router.Route("/upload", uploadHandler.RegisterRoutes)

	// Serve static files from ./public
	fs := http.FileServer(http.Dir("./public"))
	s.Router.Handle("/public/*", http.StripPrefix("/public", fs))

	// Swagger route with dynamic URL
	port := os.Getenv("PORT")
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/doc.json", port)

	s.Router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(swaggerURL), // <- dynamically sets swagger.json URL
	))
}

func (s *Server) Start(port string) error {
	fmt.Println("Server running on port", port)
	return http.ListenAndServe(":"+port, s.Router)
}
