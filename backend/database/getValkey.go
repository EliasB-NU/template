package database

import (
	"fmt"
	"github.com/valkey-io/valkey-go"
	"log"
	"template/backend/config"
)

// GetValkey returns a valkey client, used for caching
func GetValkey(cfg *config.Config) valkey.Client {
	client, err := valkey.NewClient(valkey.ClientOption{
		Username:    cfg.Database.Valkey.User,
		Password:    cfg.Database.Valkey.Password,
		InitAddress: []string{fmt.Sprintf("%s:%d", cfg.Database.Valkey.Host, cfg.Database.Valkey.Port)},
		SelectDB:    cfg.Database.Valkey.Database,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	return client
}
