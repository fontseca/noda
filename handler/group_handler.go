package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"noda"
	"noda/data/transfer"
	"noda/service"
)

type GroupHandler struct {
	s *service.GroupService
}

func NewGroupHandler(service *service.GroupService) *GroupHandler {
	return &GroupHandler{service}
}

func (h *GroupHandler) HandleGroupCreation(w http.ResponseWriter, r *http.Request) {
	var group = new(transfer.GroupCreation)
	var err = decodeJSONRequestBody(w, r, group)
	if nil != err {
		noda.EmitError(w, noda.ErrMalformedRequest.Clone().SetDetails(err.Error()))
		return
	}
	err = group.Validate()
	if nil != err {
		noda.EmitError(w, noda.ErrBadRequest.Clone().SetDetails(err.Error()))
		return
	}
	userID, _ := extractUserPayload(r)
	insertedID, err := h.s.SaveGroup(userID, group)
	if nil != err {
		var e *noda.Error
		if errors.As(err, &e) {
			noda.EmitError(w, e)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	var result = map[string]string{"insertedID": insertedID}
	data, err := json.Marshal(result)
	if nil != err {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (h *GroupHandler) HandleRetrieveGroupByID(w http.ResponseWriter, r *http.Request) {
	userID, _ := extractUserPayload(r)
	groupID, err := parsePathParameterToUUID(r, "group_id")
	if nil != err {
		var e *noda.Error
		if errors.As(err, &e) {
			noda.EmitError(w, e)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	group, err := h.s.FindGroupByID(userID, groupID)
	if nil != err {
		var e *noda.Error
		if errors.As(err, &e) {
			noda.EmitError(w, e)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	data, err := json.Marshal(group)
	if nil != err {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (h *GroupHandler) HandleGroupsRetrieval(w http.ResponseWriter, r *http.Request) {
	userID, _ := extractUserPayload(r)
	pagination := parsePagination(w, r)
	if nil == pagination {
		return
	}
	groups, err := h.s.FindGroups(userID, pagination, "", "")
	if nil != err {
		var e *noda.Error
		if errors.As(err, &e) {
			noda.EmitError(w, e)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	data, err := json.Marshal(groups)
	if nil != err {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (h *GroupHandler) HandleGroupUpdate(w http.ResponseWriter, r *http.Request) {
	var up = new(transfer.GroupUpdate)
	err := decodeJSONRequestBody(w, r, up)
	if nil != err {
		noda.EmitError(w, noda.ErrMalformedRequest.Clone().SetDetails(err.Error()))
		return
	}
	userID, _ := extractUserPayload(r)
	groupID, err := parsePathParameterToUUID(r, "group_id")
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ok, err := h.s.UpdateGroup(userID, groupID, up)
	if nil != err {
		var e *noda.Error
		if errors.As(err, &e) {
			noda.EmitError(w, e)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	if ok {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var (
		scheme = "http://"
		host   = r.Host
		path   = fmt.Sprintf("/me/groups/%s", groupID)
	)
	if nil != r.TLS {
		scheme = "https://"
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s%s", scheme, host, path))
	w.WriteHeader(http.StatusSeeOther)
}

func (h *GroupHandler) HandleGroupDeletion(w http.ResponseWriter, r *http.Request) {
	groupID, err := parsePathParameterToUUID(r, "group_id")
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userID, _ := extractUserPayload(r)
	_, err = h.s.DeleteGroup(userID, groupID)
	if nil != err {
		var e *noda.Error
		if errors.As(err, &e) {
			noda.EmitError(w, e)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}