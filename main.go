package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)
type Book struct {	

	Author string `json:"author"`	
	Title string `json:"title"`	
	Publisher string `json:"publisher"`


}

type Repository struct {
	DB *gorm.DB

}
func (r *Repository) SetupRoutes(app *fiber.App) {
	api:= app.Group("/api")
	api.Post("/create-books", r.CreateBook)
	api.Delete("/delete-books/:id", r.DeleteBook)
	api.Get("/get-books/:id", r.GetBookByID)
	api.Get("/books", r.GetAllBooks)
}

func main() {

	 err :=godotenv.Load(".env")

	 if err != nil {
		log.Fatal(err)
		fmt.Println("Error loading .env file")
	 }

	 db,err:=storage.NewConnection()
	 if err != nil {	
		log.Fatal(err)
		fmt.Println("Error connecting to database")	
	 }

	 r:=Repository{
		DB:db,
	 }
	 app := fiber.New()


	 r.SetupRoutes(app)
	 app.Listen(":8080")
}