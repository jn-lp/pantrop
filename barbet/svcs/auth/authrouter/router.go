package authrouter

import (
  "github.com/gofiber/fiber/v2"
  "github.com/jn-lp/pantrop/barbet/svcs/auth"
  "github.com/jn-lp/pantrop/barbet/svcs/auth/authdto"
)

func Router(app fiber.Router, service auth.Service) {
  // app.Get("/", authmw.Protected(), getMe(service))
  app.Post("/", login(service))
}

// func getMe(service auth.Service) fiber.Handler {
//   return func(c *fiber.Ctx) error {
//     fetched, err := service.GetMe(c.Context())
//     if err != nil {
//       _ = c.JSON(&fiber.Map{
//         "status": false,
//         "error":  err,
//       })
//     }
//     return c.JSON(&fiber.Map{
//       "status": true,
//       "books":  fetched,
//     })
//   }
// }

func login(service auth.Service) fiber.Handler {
  return func(c *fiber.Ctx) error {
    var requestBody authdto.LoginRequest
    err := c.BodyParser(&requestBody)
    if err != nil {
      _ = c.JSON(&fiber.Map{
        "success": false,
        "error":   err,
      })
    }

    result, err := service.Login(c.Context(), &requestBody)
    return c.JSON(&fiber.Map{
      "status": result,
      "error":  err,
    })
  }
}
