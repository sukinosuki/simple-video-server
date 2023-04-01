package core

import (
	"golang.org/x/time/rate"
	"time"
)

var limit = rate.Every(1000 * time.Millisecond)

var Limiter = rate.NewLimiter(limit, 10)
