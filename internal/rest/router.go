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
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		c, err := card.NewFromBody(body)
		if err != nil {
			logger.Errorf("%s: %s", op, err.Error())
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
			return
		}

		if err := c.Validate(); err != nil {
			logger.Errorf("%s: %s", op, err.Error())
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err.Error()))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("card valid"))
	}
}
