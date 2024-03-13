package main

import (
	"log"
	"net/http"

	"github.com/FelpsCorrea/GoExpertPostgraduation/API/configs"
	_ "github.com/FelpsCorrea/GoExpertPostgraduation/API/docs"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/entity"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/infra/database"
	"github.com/FelpsCorrea/GoExpertPostgraduation/API/internal/infra/webserver/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title GoExpertPostgraduation API
// @version 1.0
// @description API para gerenciamento de produtos e usuários
// @termsOfService http://swagger.io/terms/

// @contact.name Felipe Correa
// @contact.url http://github.com/FelpsCorrea
// @contact.email dev.felipecls@gmail.com

// @license.name Dev License
// @license.url http://github.com/FelpsCorrea

// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	db.AutoMigrate(&entity.User{}, &entity.Product{})

	productDB := database.NewProduct(db)

	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)

	userHandler := handlers.NewUserHandler(userDB)

	// Mux criado pelo Chi
	r := chi.NewRouter()

	// Inserir parâmetros no context
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn))

	// Para aparecer log dos requests no terminal do servidor
	r.Use(middleware.Logger)

	// middleware personalizado
	// r.Use(LogRequest)

	// Não deixa a aplicação cair mesmo se der um panic
	r.Use(middleware.Recoverer)
	// github.com/go-chi/chi
	// go install github.com/swaggo/swag/cmd/swag@latest
	// swag init -g cmd/server/main.go

	// Agrupamento de rotas
	r.Route("/products", func(r chi.Router) {

		// Middleware para validar token
		// Pega o contexto
		r.Use(jwtauth.Verifier(configs.TokenAuth))

		// Valida o token
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Get("/", productHandler.GetProducts)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	http.ListenAndServe(":8080", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
