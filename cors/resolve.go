package cors

import (
	"net/http"
	"strconv"
	"strings"
)

func ResolveCors(w http.ResponseWriter, corsConfig *CORSConfig) {
	if len(corsConfig.AllowedOrigins) > 0 {
		w.Header().Set("Access-Control-Allow-Origin", strings.Join(corsConfig.AllowedOrigins, ", "))
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
