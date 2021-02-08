package middleware

import (
	"github.com/rs/zerolog"
	"github.com/urfave/negroni"
	"github.com/vibhugarg123/book-my-show/appcontext"
	"github.com/vibhugarg123/book-my-show/utils"
	"net/http"
	"os"
)

func MDCMiddleware() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		appcontext.Logger = zerolog.New(os.Stdout).
			With().Timestamp().Logger().
			With().Str("request-method", r.Method).Logger().
			With().Str("request-url", r.URL.String()).Logger().
			With().Int64("request-header-size", utils.HeaderSize(r.Header)).Logger().
			With().Str("user-agent", r.UserAgent()).Logger().
			With().Str("referer", r.Referer()).Logger().
			With().Str("proto", r.Proto).Logger().
			With().Str("remote-ip", utils.IpFromHostPort(r.RemoteAddr)).Logger()
		next(rw, r)
	})
}
