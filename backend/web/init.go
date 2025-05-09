package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valkey-io/valkey-go"
	"gorm.io/gorm"
	"template/backend/config"
)

type API struct {
	PSQL   *gorm.DB
	Valkey valkey.Client

	CFG *config.Config
}

// InitWeb Creates the fiber instance and all api endpoints.
func InitWeb(psql *gorm.DB, valkey valkey.Client, cfg *config.Config) {
	var (
		a = API{
			PSQL:   psql,
			Valkey: valkey,
			CFG:    cfg,
		}

		backend = fiber.New(fiber.Config{
			ServerHeader: "template:fiber",
			AppName:      "template",
		})

		api = fiber.New(fiber.Config{})
	)

	backend.Use("/api", api)

	// API
	// Login
	api.Post("/login", a.login) // <- Email & Password & DeviceID | -> Token | To login into the site
}
