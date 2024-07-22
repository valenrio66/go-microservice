package controller

import (
	"fmt"
	"go-microservice/config"
	"go-microservice/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// GetBooks mengambil semua buku dari database
func GetBooks(c *fiber.Ctx) error {
	var books []model.Book
	result := config.DB.Find(&books)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"StatusCode": fiber.StatusInternalServerError,
			"Message":    "Database error",
			"Data":       nil,
		})
	}
	return c.JSON(fiber.Map{
		"StatusCode": fiber.StatusOK,
		"Message":    "Books retrieved successfully",
		"Data":       books,
	})
}

func GetBooksById(c *fiber.Ctx) error {
	var books model.Book
	id := c.Params("books_id")
	result := config.DB.Find(&books, "books_id = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"StatusCode": fiber.StatusInternalServerError,
			"Message":    "Database error",
			"Data":       nil,
		})
	}
	return c.JSON(fiber.Map{
		"StatusCode": fiber.StatusOK,
		"Message":    "Books retrieved successfully",
		"Data":       books,
	})
}

// CreateBook menambahkan buku baru ke database
func CreateBook(c *fiber.Ctx) error {
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"StatusCode": fiber.StatusBadRequest,
			"Message":    "Review your input",
			"Data":       nil,
		})
	}
	result := config.DB.Create(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"StatusCode": fiber.StatusInternalServerError,
			"Message":    "Database error",
			"Data":       nil,
		})
	}
	return c.JSON(fiber.Map{
		"StatusCode": fiber.StatusCreated,
		"Message":    "Book added successfully",
		"Data":       book,
	})
}

// UpdateBook memperbarui buku yang ada berdasarkan id
func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("books_id")
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"StatusCode": fiber.StatusBadRequest,
			"Message":    "Review your input",
			"Data":       nil,
		})
	}
	result := config.DB.Model(&model.Book{}).Where("books_id = ?", id).Updates(book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"StatusCode": fiber.StatusInternalServerError,
			"Message":    "Database error",
			"Data":       nil,
		})
	}
	return c.JSON(fiber.Map{
		"StatusCode": fiber.StatusOK,
		"Message":    "Book updated successfully",
		"Data":       book,
	})
}

// DeleteBook menghapus buku dari database
func DeleteBook(c *fiber.Ctx) error {
	idStr := c.Params("books_id") // Mengambil 'books_id' dari query parameter
	fmt.Println("books_id:", idStr)
	id, err := strconv.Atoi(idStr)
	fmt.Println("books_id converted:", id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"StatusCode": fiber.StatusBadRequest,
			"Message":    "Invalid book ID",
			"Data":       nil,
		})
	}
	result := config.DB.Delete(&model.Book{}, "books_id = ?", id)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"StatusCode": fiber.StatusInternalServerError,
			"Message":    "Database error",
			"Data":       nil,
		})
	}
	return c.JSON(fiber.Map{
		"StatusCode": fiber.StatusNoContent,
		"Message":    "Book deleted successfully",
		"Data":       nil,
	})
}
