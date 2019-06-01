package main

import (
	"net/http"
)

type authHandler struct {
	next http.Handler
}
