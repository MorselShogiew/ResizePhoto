package middleware

import (
	"net/http"
	"strconv"
)

type ContextKey string

const (
	contextKeyRequestID  ContextKey = "requestID"
	contextKeyEmployeeID ContextKey = "employeeID"
	contextKeyPerms      ContextKey = "permissions"
	contextKeyPositionID ContextKey = "position_id"
	contextKeySu         ContextKey = "su"
	contextKeyPhone      ContextKey = "phone"
)

func GetReqID(r *http.Request) string {
	id := r.Context().Value(contextKeyRequestID)
	if ret, ok := id.(string); ok {
		return ret
	}

	return "0"
}

func GetEmployeeID(r *http.Request) int64 {
	id := r.Context().Value(contextKeyEmployeeID)
	ret, ok := id.(string)
	if !ok {
		return 0
	}

	employeeID, err := strconv.ParseInt(ret, 10, 64)
	if err != nil {
		return 0
	}

	return employeeID
}
