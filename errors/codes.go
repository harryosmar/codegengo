package errors

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CodeErr int

type CodeErrEntity struct {
	Code    string
	Status  codes.Code
	Message string
}

const (
	ErrGeneral CodeErr = iota
)

var codeErrMap = map[CodeErr]CodeErrEntity{
	ErrGeneral: {Code: "ERR999", Status: codes.InvalidArgument, Message: "Terjadi kesalahan."},
}

func GetCodeErrMap[T CodeErrEntity](k CodeErr) T {
	return T(codeErrMap[k])
}

func (c CodeErr) GRPCStatus() *status.Status {
	return status.New(codeErrMap[c].Status, c.Error())
}

func (c CodeErr) Error() string {
	codeErrEntity := codeErrMap[c]
	return fmt.Sprintf("[%s] %s", codeErrEntity.Code, codeErrEntity.Message)
}

type CodeErrEntityWithDetails struct {
	codeErr CodeErr
	Details []proto.Message
}

func NewCodeErrEntityWithDetails(codeErr CodeErr, details ...proto.Message) *CodeErrEntityWithDetails {
	return &CodeErrEntityWithDetails{codeErr: codeErr, Details: details}
}

func (c CodeErrEntityWithDetails) Error() string {
	return c.codeErr.Error()
}

func (c CodeErrEntityWithDetails) GRPCStatus() *status.Status {
	grpcStatus := c.codeErr.GRPCStatus()
	if withDetails, err := grpcStatus.WithDetails(c.Details...); err == nil {
		return withDetails
	}
	return grpcStatus
}
