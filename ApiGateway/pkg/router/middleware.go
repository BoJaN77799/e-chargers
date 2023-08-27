package router

import (
	"ApiGateway/pkg/handlers"
	"ApiGateway/pkg/utils"
	"net/http"
	"strings"
)

var publicRouteParts = []string{
	"/api/auth/login",
	"/api/auth/register",
	"/api/chargers",
	"/api/chargers/search",
	"/api/chargers/closest",
	"/api/recensions/charger",
}

func isRoutePublic(route string) bool {
	for _, publicRoutePart := range publicRouteParts {
		if strings.Contains(route, publicRoutePart) {
			return true
		}
	}
	return false
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		if isRoutePublic(r.URL.String()) {
			next.ServeHTTP(w, r)
			return
		}

		accessToken, err := utils.ExtractAccessTokenFromHeader(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		err = utils.IsAccessTokenValid(accessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		_, err = handlers.DoRequestWithToken(r, http.MethodGet, utils.BaseUserServicePath.Next().Host+"/auth", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isRoleInRequiredRoles(requiredRoles []string, role string) bool {
	for _, requiredRole := range requiredRoles {
		if strings.EqualFold(requiredRole, role) {
			return true
		}
	}
	return false
}

func authorizationMiddleware(next http.HandlerFunc, requiredRoles []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := utils.GetRoleFromToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if role == Administrator {
			next(w, r)
		}

		if !isRoleInRequiredRoles(requiredRoles, role) {
			http.Error(w, "Required scope: ["+strings.Join(requiredRoles, " ,")+"]", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
