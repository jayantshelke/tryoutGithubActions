package httpserver

import (
	"ProjectIdeas/monolith/api"
	"context"
	"encoding/json"
	"net/http"
)

type publicParamHandler struct {
	//puParams map[string]string
	//queryParams map[string]string
	//headerParams map[string]string
	innerH http.Handler
}

func (p publicParamHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// read public parameters here and passdown via setting them in the context
	p.innerH.ServeHTTP(writer, request)
}

type erroringHandler struct {
	InnerH Servicer
}

func (e erroringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// check for returning errors
	_ = e.InnerH.serve(context.Background(), writer, request)

	//do something with this error and then return
}

type loggingHandler struct {
	innerH http.Handler
}

func (l loggingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	l.innerH.ServeHTTP(writer, request)
}

type Servicer interface {
	serve(ctx context.Context, writer http.ResponseWriter, request *http.Request) error
}

type userGetHandler struct {
	a api.APIer
}

func newUserGetHandler(ctx context.Context, a api.APIer) http.Handler {
	return loggingHandler{
		publicParamHandler{
			erroringHandler{
				userGetHandler{
					a: a,
				},
			},
		},
	}
}

func (u userGetHandler) serve(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	// get the user either with the id that's passed down the context
	// or with the email. Check which one is passed
	ctx = request.Context() // this needs to be first set in the precursor handlers.

	user, err := u.a.UserByID(ctx, `testId`)
	if err != nil {
		return err
	}

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(data)

	return nil
}

type userDeleteHandler struct {
	a api.APIer
}

func newUserDeleteHandler(ctx context.Context, a api.APIer) http.Handler {
	return loggingHandler{
		publicParamHandler{
			erroringHandler{
				userDeleteHandler{
					a: a,
				},
			},
		},
	}
}

func (u userDeleteHandler) serve(ctx context.Context, writer http.ResponseWriter, request *http.Request) error {
	// get the user either with the id that's passed down the context
	// or with the email. Check which one is passed
	ctx = request.Context() // this needs to be first set in the precursor handlers.

	if err := u.a.DeleteUserByID(ctx, `testId`); err != nil {
		return err
	}

	writer.WriteHeader(http.StatusOK)
	return nil
}
