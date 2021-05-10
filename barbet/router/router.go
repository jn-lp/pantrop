package router

import (
  "log"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/monitor"

  "gorm.io/driver/sqlite"
  "gorm.io/gorm"

  "github.com/jn-lp/pantrop/barbet/svcs/trips"
  "github.com/jn-lp/pantrop/barbet/svcs/trips/tripsrepo"
  "github.com/jn-lp/pantrop/barbet/svcs/trips/tripsrouter"
  "github.com/jn-lp/pantrop/barbet/svcs/users"
  "github.com/jn-lp/pantrop/barbet/svcs/users/usersrepo"
  "github.com/jn-lp/pantrop/barbet/svcs/users/usersrouter"
)

func Setup(app *fiber.App) {
  app.Get("/dashboard", monitor.New())

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World 👋!")
  })

  {
    db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{PrepareStmt: true})
    if err != nil {
      log.Fatal("Users DB Connection Error $s", err)
    }
    // db.AutoMigrate(&barbet.User{})

    repo := usersrepo.New(db)
    service := users.New(repo)
    routes := app.Group("/v1/users")
    usersrouter.Router(routes, service)
  }

  {
    db, err := gorm.Open(sqlite.Open("trips.db"), &gorm.Config{PrepareStmt: true})
    if err != nil {
      log.Fatal("Trips DB Connection Error $s", err)
    }
    // db.AutoMigrate(&barbet.Trip{})

    repo := tripsrepo.New(db)
    service := trips.New(repo)
    routes := app.Group("/v1/trips")
    tripsrouter.Router(routes, service)
  }

  app.Use(func(c *fiber.Ctx) error {
    return c.SendStatus(404)
  })
}
