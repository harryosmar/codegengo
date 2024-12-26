package errors

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
	"net/http"
	"regexp"
)

func ErrorHandlerFunc(ctx context.Context, _ *runtime.ServeMux, _ runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	code, message, details, httpStatus := getErrorCode(err)

	// Customize the error response
	customError := map[string]interface{}{
		"code":    code,
		"message": message,
	}

	if details != nil {
		customError["details"] = details
	}

	// Set the response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	// Write the response
	if err = json.NewEncoder(w).Encode(customError); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getErrorCode(err error) (code string, message string, details interface{}, httpStatus int) {
	// Convert the error to a gRPC status
	grpcStatus, ok := status.FromError(err)
	if !ok {
		return "ERR999", "Unknown error", nil, http.StatusInternalServerError
	}

	s := grpcStatus.Message()
	re := regexp.MustCompile(`^\[(.*)\]\s+(.*)$`)
	matches := re.FindAllStringSubmatch(s, -1)

	if len(matches) == 1 && len(matches[0]) == 3 {
		var d interface{}
		d = grpcStatus.Details()
		grpcDetails := grpcStatus.Details()
		if len(grpcDetails) < 2 {
			d = grpcDetails[0]
		}

		return matches[0][1], matches[0][2], d, runtime.HTTPStatusFromCode(grpcStatus.Code())
	}

	return "ERR999", s, nil, http.StatusInternalServerError
}
