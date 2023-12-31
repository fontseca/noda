package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"noda"
	"noda/data/transfer"
	"noda/service"
)

type AuthenticationHandler struct {
	s service.AuthenticationService
}

func NewAuthenticationHandler(s service.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{s}
}

func (h *AuthenticationHandler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	next := &transfer.UserCreation{}
	var err = parseRequestBody(w, r, next)
	if nil != err {
		noda.EmitError(w, noda.ErrMalformedRequest.Clone().SetDetails(err.Error()))
		return
	}
	err = next.Validate()
	if nil != err {
		noda.EmitError(w, noda.ErrBadRequest.Clone().SetDetails(err.Error()))
		return
	}
	insertedID, err := h.s.SignUp(next)
	if err != nil {
		var (
			a *noda.AggregateDetails
			e *noda.Error
		)
		if errors.As(err, &a) {
			noda.EmitError(w, noda.ErrPasswordRestrictions.Clone().SetDetails(a.Error()))
		} else if errors.As(err, &e) {
			noda.EmitError(w, e)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var payload = map[string]string{"user_id": insertedID.String()}
	data, err := json.Marshal(payload)
	if nil != err {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (h *AuthenticationHandler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	credentials := &transfer.UserCredentials{}
	var err = parseRequestBody(w, r, credentials)
	if nil != err {
		noda.EmitError(w, noda.ErrMalformedRequest.Clone().SetDetails(err.Error()))
		return
	}
	res, err := h.s.SignIn(credentials)
	if err != nil {
		var e *noda.Error
		if errors.As(err, &e) {
			switch {
			default:
				noda.EmitError(w, e)
			case errors.Is(e, noda.ErrUserNotFound):
				noda.EmitError(w, e.
					Clone().
					SetDetails("Could not find any user with the email %q.").
					FormatDetails(credentials.Email).
					SetHint("Verify email address or use another one."))
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	data, err := json.Marshal(res)
	if nil != err {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Write(data)
}
