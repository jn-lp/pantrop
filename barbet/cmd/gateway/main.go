package main

import (
  "flag"
  "log"
  "os"
  "time"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cache"
  "github.com/gofiber/fiber/v2/middleware/compress"
  "github.com/gofiber/fiber/v2/middleware/cors"
  "github.com/gofiber/fiber/v2/middleware/limiter"
  "github.com/gofiber/fiber/v2/middleware/logger"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "github.com/jn-lp/pantrop/barbet/router"

  "github.com/jn-lp/pantrop/barbet/pkg/signal"
)

var (
  port = flag.String("port", ":8000", "Port to listen on")
)

func main() {
  flag.Parse()

  // if err := godotenv.Load(".env"); err != nil {
  //   log.Fatalf("Error loading .env file")
  // }

  app := fiber.New(fiber.Config{
    Prefork:     true,
    BodyLimit:   4 * 1024 * 1024,
    Concurrency: 256 * 1024,
  })

  app.Use(compress.New(compress.Config{
    Next: func(c *fiber.Ctx) bool {
      return c.Path() == "/dashboard"
    },
    Level: compress.LevelBestSpeed,
  }))

  app.Use(cache.New(cache.Config{
    Next: func(c *fiber.Ctx) bool {
      return c.Path() == "/dashboard" || c.Query("refresh") == "true"
    },
    Expiration:   30 * time.Minute,
    CacheControl: true,
    KeyGenerator: func(c *fiber.Ctx) string {
      return c.Path()
    },
    // Storage: sqlite3.New(),
  }))

  file, err := os.OpenFile("./requests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
  if err != nil {
    log.Fatalf("error opening file: %v", err)
  }
  defer func(file *os.File) {
    err := file.Close()
    if err != nil {

    }
  }(file)

  app.Use(logger.New(logger.Config{
    Format:       "[${time}] ${pid} ${status} – ${latency} ${method} ${path}\n",
    TimeFormat:   "02-Jan-2006",
    TimeZone:     "Local",
    TimeInterval: time.Second,
    Output:       file,
  }))

  app.Use(limiter.New(limiter.Config{
    Next: func(c *fiber.Ctx) bool {
      return c.IP() == "127.0.0.1"
    },
    Max:        5,
    Expiration: 30 * time.Second,
    KeyGenerator: func(c *fiber.Ctx) string {
      return "key"
    },
    // Storage: sqlite3.New(),
  }))

  app.Use(cors.New())
  // app.Use(recover.New())

  router.Setup(app)

  go func() {
    if err := app.Listen(*port); err != nil {
      log.Panic(err)
    }
  }()

  signal.WaitForTermination()
  if err = app.Shutdown(); err != nil {
    log.Panic(err)
  }
}
