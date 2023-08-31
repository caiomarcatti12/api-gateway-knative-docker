package cors

import (
	"net/http"
	"strconv"
	"strings"
)

func ResolveCors(w http.ResponseWriter, r *http.Request, corsConfig *CORSConfig) {
	origin := r.Header.Get("Origin")

	isAllowed := false

	for _, allowedOrigin := range corsConfig.AllowedOrigins {
		if allowedOrigin == origin {
			isAllowed = true
			break
		}
	}

	if isAllowed {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}

	if len(corsConfig.AllowedMethods) > 0 {
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(corsConfig.AllowedMethods, ", "))
	}
	if len(corsConfig.AllowedHeaders) > 0 {
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(corsConfig.AllowedHeaders, ", "))
	}
	if len(corsConfig.ExposedHeaders) > 0 {
		w.Header().Set("Access-Control-Expose-Headers", strings.Join(corsConfig.ExposedHeaders, ", "))
	}
	if corsConfig.AllowCredentials {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	if corsConfig.MaxAge > 0 {
		w.Header().Set("Access-Control-Max-Age", strconv.Itoa(corsConfig.MaxAge))
	}
}
