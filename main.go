package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	HashType string
	HashOk   bool
	HashStr  string `json:"HashStr" xml:"HashStr" form:"HashStr"`
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetSHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetSHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetSHA512Hash(text string) string {
	hash := sha512.Sum512([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	p := new(App)

	app := fiber.New(fiber.Config{
		ServerHeader: "HashWeb Server",
		AppName:      "HashWeb v0.1",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index.html", fiber.Map{"HashOk": false, "HashType": "md5"})
	})

	app.Get("/logo.svg", func(c *fiber.Ctx) error {
		return c.SendFile("./logo.svg")
	})

	app.Get("/md5", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		return c.Render("index.html", fiber.Map{"HashStr": GetMD5Hash(p.HashStr), "HashOk": false, "HashType": "md5"})
	})

	app.Post("/md5", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.Render("index.html", fiber.Map{"HashStr": GetMD5Hash(p.HashStr), "HashOk": true, "HashType": "md5"})
	})

	app.Get("/sha256", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		return c.Render("index.html", fiber.Map{"HashStr": GetSHA256Hash(p.HashStr), "HashOk": false, "HashType": "sha256"})
	})

	app.Post("/sha256", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.Render("index.html", fiber.Map{"HashStr": GetSHA256Hash(p.HashStr), "HashOk": true, "HashType": "sha256"})
	})

	app.Get("/sha1", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		return c.Render("index.html", fiber.Map{"HashStr": GetSHA1Hash(p.HashStr), "HashOk": false, "HashType": "sha1"})
	})

	app.Post("/sha1", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.Render("index.html", fiber.Map{"HashStr": GetSHA1Hash(p.HashStr), "HashOk": true, "HashType": "sha1"})
	})

	app.Get("/sha512", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		return c.Render("index.html", fiber.Map{"HashStr": GetSHA512Hash(p.HashStr), "HashOk": false, "HashType": "sha512"})
	})

	app.Post("/sha512", func(c *fiber.Ctx) error {
		c.Accepts("text/html")
		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.Render("index.html", fiber.Map{"HashStr": GetSHA512Hash(p.HashStr), "HashOk": true, "HashType": "sha512"})
	})

	app.Listen(":3000")
}
