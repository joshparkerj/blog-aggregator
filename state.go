package main

import (
	"github.com/joshparkerj/blog-aggregator/internal/config"
	"github.com/joshparkerj/blog-aggregator/internal/database"
)

type State struct {
	Configuration *config.Config
	DB            *database.Queries
}
