package database

import (
	"context"
	"fmt"

	"github.com/crowdsecurity/crowdsec/pkg/csconfig"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent"
)

func NewClient(config *csconfig.DatabaseConfig) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", config.Path))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to sqlite: %v", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}
	return client, nil
}
