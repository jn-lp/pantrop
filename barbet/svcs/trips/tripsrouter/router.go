package tripsrouter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jn-lp/pantrop/barbet"
	authmw2 "github.com/jn-lp/pantrop/barbet/svcs/auth/authmw"
	"github.com/jn-lp/pantrop/barbet/svcs/trips"
)

func Router(app fiber.Router, service trips.Service) {
  app.Get("/", listTrips(service))
  app.Post("/", authmw2.Protected(), createTrip(service))
  app.Put("/", authmw2.Protected(), updateTrip(service))
  app.Get("/:tripID", getTrip(service))
  app.Delete("/:tripID", authmw2.Protected(), deleteTrip(service))
}

func listTrips(service trips.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    fetched, err := service.ListTrips(c.Context())
    if err != nil {
      _ = c.JSON(&fiber.Map{
        "status": false,
        "error":  err,
      })
    }
    return c.JSON(&fiber.Map{
      "status": true,
      "trips":  fetched,
    })
  }
}

func createTrip(service trips.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    var requestBody barbet.Trip
    err := c.BodyParser(&requestBody)
    if err != nil {
      _ = c.JSON(&fiber.Map{
        "success": false,
        "error":   err,
      })
    }
    result, dbErr := service.CreateTrip(c.Context(), &requestBody)
    return c.JSON(&fiber.Map{
      "status": result,
      "error":  dbErr,
    })
  }
}

func updateTrip(service trips.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    var requestBody barbet.Trip
    err := c.BodyParser(&requestBody)
    if err != nil {
      _ = c.JSON(&fiber.Map{
        "success": false,
        "error":   err,
      })
    }
    result, dbErr := service.UpdateTrip(c.Context(), &requestBody)
    return c.JSON(&fiber.Map{
      "status": result,
      "error":  dbErr,
    })
  }
}

func getTrip(service trips.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    user, dbErr := service.GetTrip(c.Context(), c.Params("tripID"))
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

func deleteTrip(service trips.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    dbErr := service.DeleteTrip(c.Context(), c.Params("tripID"))
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
