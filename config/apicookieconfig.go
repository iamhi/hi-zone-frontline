package config

import (
	"os"
	"strconv"
)

const default_domain = "localhost"
const default_httponly = false
const default_secure = false
const default_maxage = 172800
const default_path = "/"

var api_cookie_config = ApiCookieConfig{}

type ApiCookieConfig struct {
	Domain   string
	HttpOnly bool
	Secure   bool
	MaxAge   int
	Path     string
}

func initApiCookieConfig() {
	domain := os.Getenv("COOKIE_DOMAIN")

	if domain == "" {
		domain = default_domain
	}

	path := os.Getenv("COOKIE_PATH")

	if path == "" {
		path = default_path
	}

	httpOnly, err := strconv.ParseBool(os.Getenv("COOKIE_HTTPONLY"))

	if err != nil {
		httpOnly = default_httponly
	}

	secure, err := strconv.ParseBool(os.Getenv("COOKIE_SECURE"))

	if err != nil {
		secure = default_secure
	}

	maxAge, err := strconv.Atoi(os.Getenv("COOKIE_MAXAGE"))

	if err != nil {
		maxAge = default_maxage
	}

	api_cookie_config = ApiCookieConfig{
		Domain:   domain,
		HttpOnly: httpOnly,
		Secure:   secure,
		MaxAge:   maxAge,
		Path:     path}
}

func GetApiCookieConfig() ApiCookieConfig {
	return api_cookie_config
}
