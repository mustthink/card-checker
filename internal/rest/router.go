package rest

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/mustthink/card-checker/internal/card"
)

func getRouter(log *logrus.Logger) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		"/validate",
		validator(log.WithField("handler", "/validate")),
	)

	return router
}

func validator(logger *logrus.Entry) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		const op = "rest.handler.Validator"

		body, err := io.ReadAll(request.Body)
		if err != nil {
			logger.Errorf("%s: %s", op, err.Error())
			writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		c, err := card.NewFromBody(body)
		if err != nil {
			logger.Errorf("%s: %s", op, err.Error())
			writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := c.Validate(); err != nil {
			logger.Errorf("%s: %s", op, err.Error())
			writer.Write([]byte(err.Error()))
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		writer.WriteHeader(http.StatusOK)
	}
}
