package utils

import (
	"back-usm/internals/order/core/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func Payment(ctx *fiber.Ctx) error {
	var order domain.Order
	err := ctx.BodyParser(&order)
	if err != nil {
		StatusError("400", "POST", "Payment")
		return ctx.Status(400).JSON("Invalid order")
	}

	stripe.Key = GetEnvVar("STRIPE_KEY")

	_, err = charge.New(&stripe.ChargeParams{
		Amount:       stripe.Int64(order.Total),
		Currency:     stripe.String(string(stripe.CurrencyCLP)),
		Source:       &stripe.SourceParams{Token: stripe.String("tok_visa")},
		ReceiptEmail: stripe.String(order.CustomerEmail),
	})

	if err != nil {
		StatusError("400", "POST", "Payment")
		return ctx.Status(400).JSON(err)
	}

	StatusOk("200", "POST", "Payment")
	return ctx.Status(200).JSON("Payment successful")
}
