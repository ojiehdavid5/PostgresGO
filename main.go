package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"github.com/chuks/PostgresGO/storage"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type repository struct {
	DB *gorm.DB
}
type Repository interface {
	CreateBook(c *fiber.Ctx) error
	DeleteBook(c *fiber.Ctx) error
	GetBookByID(c *fiber.Ctx) error
	GetAllBooks(c *fiber.Ctx) error
}

func NewRepository(db *gorm.DB) Repository {

	return &repository{
		DB: db,
	}
}

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	r := NewRepository(db)
	api := app.Group("/api")
	api.Post("/create-books", r.CreateBook)
	api.Delete("/delete-books/:id", r.DeleteBook)
	api.Get("/get-books/:id", r.GetBookByID)
	api.Get("/books", r.GetAllBooks)
}

func (r *repository) CreateBook(c *fiber.Ctx) error {
	panic("implement me")
}
func (r *repository) DeleteBook(c *fiber.Ctx) error {
	panic("implement")
}
func (r *repository) GetBookByID(c *fiber.Ctx) error {
	panic("implement")
}
func (r *repository) GetAllBooks(c *fiber.Ctx) error {
	panic("implement")
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		fmt.Println("Error loading .env file")
	}

	db, err := storage.NewConnection()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error connecting to database")
	}

	app := fiber.New()

	SetupRoutes(app, db)
	app.Listen(":8080")
}
