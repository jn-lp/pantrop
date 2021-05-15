package usersrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jn-lp/pantrop/barbet"
	authmw2 "github.com/jn-lp/pantrop/barbet/svcs/auth/authmw"
	"github.com/jn-lp/pantrop/barbet/svcs/users"
)

func Router(app fiber.Router, service users.Service) {
  app.Get("/", listUsers(service))
  app.Post("/", createUser(service))
  app.Put("/", authmw2.Protected(), updateUser(service))
  app.Get("/:username", getUser(service))
  app.Delete("/:username", authmw2.Protected(), deleteUser(service))
}

func listUsers(service users.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    fetched, err := service.ListUsers(c.Context())
    if err != nil {
      _ = c.JSON(&fiber.Map{
        "status": false,
        "error":  err,
      })
    }
    return c.JSON(&fiber.Map{
      "status": true,
      "books":  fetched,
    })
  }
}

func createUser(service users.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    var requestBody barbet.User
    err := c.BodyParser(&requestBody)
    if err != nil {
      _ = c.JSON(&fiber.Map{
        "success": false,
        "error":   err,
      })
    }
    result, dbErr := service.CreateUser(c.Context(), &requestBody)
    return c.JSON(&fiber.Map{
      "status": result,
      "error":  dbErr,
    })
  }
}

func updateUser(service users.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    var requestBody barbet.User
    err := c.BodyParser(&requestBody)
    if err != nil {
      _ = c.JSON(&fiber.Map{
        "success": false,
        "error":   err,
      })
    }
    result, dbErr := service.UpdateUser(c.Context(), &requestBody)
    return c.JSON(&fiber.Map{
      "status": result,
      "error":  dbErr,
    })
  }
}

func getUser(service users.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    user, dbErr := service.GetUser(c.Context(), c.Params("username"))
    if dbErr != nil {
      _ = c.JSON(&fiber.Map{
        "status": false,
        "error":  dbErr,
      })
    }

    return c.JSON(&fiber.Map{
      "status": true,
      "result": user,
    })
  }
}

func deleteUser(service users.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    dbErr := service.DeleteUser(c.Context(), c.Params("username"))
    if dbErr != nil {
      _ = c.JSON(&fiber.Map{
        "status": false,
        "error":  dbErr,
      })
    }
    return c.JSON(&fiber.Map{
      "status":  false,
      "message": "deleted successfully",
    })
  }
}
