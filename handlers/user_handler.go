package handlers

import "github.com/gofiber/fiber/v2"

func Greet(c *fiber.Ctx) error {

	return c.SendFile("./public/index.html")
}

func Hello(c *fiber.Ctx) error {
	name := c.Query("name", "World")
	return c.SendString("Hello, " + name + "!")
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"status":      "success",
		"message":     "User fetched successfully",
		"status_code": 200,
		"user_id":     id,
	})
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(c *fiber.Ctx) error {
	var data User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "failed",
			"status_code": 400,
			"message":     "Cannot parse JSON",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":      "success",
		"status_code": 201,
		"message":     "User created successfully",
		"user":        data,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var data User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "failed",
			"status_code": 400,
			"message":     "Cannot parse JSON",
		})
	}
	return c.JSON(fiber.Map{
		"status":      "success",
		"status_code": 200,
		"message":     "User updated successfully",
		"user_id":     id,
		"user":        data,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"status":      "success",
		"status_code": 200,
		"message":     "User deleted successfully",
		"user_id":     id,
	})
}
