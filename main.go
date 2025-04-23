package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

type Repository struct {
	DB *gorm.DB

}
func (r *Repository) SetupRoutes(app *fiber.App) {
	app.Get("/users", r.GetUsers)
	app.Post("/users", r.CreateUser)
	app.Put("/users/:id", r.UpdateUser)
	app.Delete("/users/:id", r.DeleteUser)
}

func main() {

	 err :=godotenv.Load(".env")

	 if err != nil {
		log.Fatal(err)
		fmt.Println("Error loading .env file")
	 }

	 r:=Repository{
		DB:db,
	 }
	 app := fiber.New()


	 r.SetupRoutes(app)
	 app.listen(":8080")
}