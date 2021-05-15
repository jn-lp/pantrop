package router

import (
  "log"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/monitor"
  "github.com/jn-lp/pantrop/barbet"
  "github.com/jn-lp/pantrop/barbet/svcs/auth"
  "github.com/jn-lp/pantrop/barbet/svcs/auth/authrouter"
  "github.com/jn-lp/pantrop/barbet/svcs/auth/authsvc"
  "github.com/jn-lp/pantrop/barbet/svcs/trips/tripssvc"
  "github.com/jn-lp/pantrop/barbet/svcs/users/userssvc"

  "gorm.io/gorm"

  "github.com/jn-lp/pantrop/barbet/svcs/trips"
  "github.com/jn-lp/pantrop/barbet/svcs/trips/tripsrepo"
  "github.com/jn-lp/pantrop/barbet/svcs/trips/tripsrouter"
  "github.com/jn-lp/pantrop/barbet/svcs/users"
  "github.com/jn-lp/pantrop/barbet/svcs/users/usersrepo"
  "github.com/jn-lp/pantrop/barbet/svcs/users/usersrouter"
)

func Setup(app *fiber.App, dialector gorm.Dialector) {
  app.Get("/dashboard", monitor.New())

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World 👋!")
  })

  var usersService users.Service
  {
    db, err := gorm.Open(dialector, &gorm.Config{PrepareStmt: true})
    if err != nil {
      log.Fatal("Users DB Connection Error $s", err)
    }
    err = db.AutoMigrate(&barbet.User{})
    if err != nil {
      log.Fatal("Users DB Migration Error $s", err)
    }

    repo := usersrepo.New(db)
    usersService = userssvc.New(repo)
    routes := app.Group("/v1/users")
    usersrouter.Router(routes, usersService)
  }

  var tripsService trips.Service
  {
    db, err := gorm.Open(dialector, &gorm.Config{PrepareStmt: true})
    if err != nil {
      log.Fatal("Trips DB Connection Error $s", err)
    }
    err = db.AutoMigrate(&barbet.Trip{})
    if err != nil {
      log.Fatal("Trips DB Migration Error $s", err)
    }

    repo := tripsrepo.New(db)
    tripsService = tripssvc.New(repo)
    routes := app.Group("/v1/trips")
    tripsrouter.Router(routes, tripsService)
  }

  var authService auth.Service
  {
    authService = authsvc.New(usersService)
    routes := app.Group("/v1/auth")
    authrouter.Router(routes, authService)
  }

  app.Use(func(c *fiber.Ctx) error {
    return c.SendStatus(404)
  })
}
