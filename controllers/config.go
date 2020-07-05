package controllers

import (
	"context"

	"github.com/lmduc/contest-stats/database"
	"github.com/sirupsen/logrus"
)

// Config holds common configurations
type Config struct {
	Ctx    context.Context
	Logger *logrus.Entry
	DB     database.Database
}
