package middleware


import (
	"context"
	"net/http"
	"notes-api/internal/auth"
	"strings"
)

type Authmiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *Authmiddleware {
	return &Authmiddleware{
		jwtService: jwtService,
	}
}

func (a *Authmiddleware) Middleware(next http.Handler) http.Handler {
	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// extract the token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// check for bearer token format
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := parts[1]

		// validate the token
		claims, err := a.jwtService.ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// add user info to the request context
		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// helper function to get user ID from context
func GetUserIDFromContext(ctx context.Context) string {
    if userID, ok := ctx.Value("userID").(string); ok {
        return userID
    }
    return ""
}