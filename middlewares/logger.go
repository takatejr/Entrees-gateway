package middlewares

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Log struct {
	time time.Time
	url  string
}

type ResLog struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

func ResponseLog(msg ResLog) {
	log.Println(ResLog{
		Status:  msg.Status,
		Message: msg.Message,
	})
}

func RequestLogger(next http.Handler, url string) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Print(
			Log{
				time: time.Time{},
				url:  url,
			})

		next.ServeHTTP(res, req)
	})
}
