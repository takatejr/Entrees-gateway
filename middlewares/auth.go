package middlewares

import (
	"encoding/json"
	"net/http"
)

func CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Token")

		if token == "" {
			resLog := ResLog{
				Status:  http.StatusUnauthorized,
				Message: "Not Authorized",
			}

			odp, _ := json.Marshal(resLog)

			res.WriteHeader(http.StatusUnauthorized)
			_, err := res.Write(odp)

			if err != nil {
				return
			}

			ResponseLog(resLog)

			return
		}

		next.ServeHTTP(res, req)
	})
}
