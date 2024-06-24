package main

import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/he-he-heng/toktok-backend/ent"
	"github.com/he-he-heng/toktok-backend/utils"
)

func main() {
	dsn := "root:qwe123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	q := &Querier{client: client}
	app := fiber.New()

	app.Post("/api/v1/auth/code/phone-issue", func(c *fiber.Ctx) error {
		var body PhoneRequest
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		a, err := q.GetAttempt(context.Background(), body.Phone)
		if ent.IsNotFound(err) {
			newAttempt, err := q.CreateAttempt(context.Background(), &ent.Attempt{
				Phone:     body.Phone,
				Cnt:       0,
				Timestamp: time.Now(),
				Break:     false,
				Authcode:  9999999999999,
			})
			if err != nil {
				return err
			}

			a = newAttempt
		} else if err != nil {
			return err
		}

		if a.Timestamp.Add(time.Hour * 24).Before(time.Now()) {
			updatedAttempt, err := a.Update().
				SetCnt(0).
				SetBreak(false).
				SetPhone(body.Phone).
				SetTimestamp(time.Now()).
				SetAuthcode(99999999).
				Save(context.Background())
			if err != nil {
				return err
			}

			a = updatedAttempt
		}

		if a.Break {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if a.Cnt >= 5 {
			a.Update().
				SetCnt(0).
				SetBreak(true).
				SetPhone(body.Phone).
				SetAuthcode(99999999).
				SetTimestamp(time.Now())

			return c.SendStatus(fiber.StatusBadRequest)
		}

		a.Cnt++
		a.Authcode = utils.RandomDigit()

		_, err = q.UpdateAttempt(context.Background(), body.Phone, a)
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{"authcode": a.Authcode})
	})

	app.Post("/api/v1/auth/code/phone-verify", func(c *fiber.Ctx) error {
		var body VertifyReqeus
		if err := c.BodyParser(&body); err != nil {
			return err
		}

		a, err := q.GetAttempt(context.Background(), body.Phone)
		if err != nil {
			return err
		}

		if a.Timestamp.Add(time.Minute * 5).Before(time.Now()) {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if a.Authcode != body.Authcode {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.SendString("ok~")
	})

	app.Listen(":8080")
}
