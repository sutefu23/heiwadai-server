// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/user/UserReport.proto

package userconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	user "server/api/v1/user"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// UserReportControllerName is the fully-qualified name of the UserReportController service.
	UserReportControllerName = "server.user.UserReportController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserReportControllerSendProcedure is the fully-qualified name of the UserReportController's Send
	// RPC.
	UserReportControllerSendProcedure = "/server.user.UserReportController/Send"
)

// UserReportControllerClient is a client for the server.user.UserReportController service.
type UserReportControllerClient interface {
	Send(context.Context, *connect_go.Request[user.UserReportRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewUserReportControllerClient constructs a client for the server.user.UserReportController
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserReportControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) UserReportControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userReportControllerClient{
		send: connect_go.NewClient[user.UserReportRequest, emptypb.Empty](
			httpClient,
			baseURL+UserReportControllerSendProcedure,
			opts...,
		),
	}
}

// userReportControllerClient implements UserReportControllerClient.
type userReportControllerClient struct {
	send *connect_go.Client[user.UserReportRequest, emptypb.Empty]
}

// Send calls server.user.UserReportController.Send.
func (c *userReportControllerClient) Send(ctx context.Context, req *connect_go.Request[user.UserReportRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.send.CallUnary(ctx, req)
}

// UserReportControllerHandler is an implementation of the server.user.UserReportController service.
type UserReportControllerHandler interface {
	Send(context.Context, *connect_go.Request[user.UserReportRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewUserReportControllerHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserReportControllerHandler(svc UserReportControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	userReportControllerSendHandler := connect_go.NewUnaryHandler(
		UserReportControllerSendProcedure,
		svc.Send,
		opts...,
	)
	return "/server.user.UserReportController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserReportControllerSendProcedure:
			userReportControllerSendHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserReportControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedUserReportControllerHandler struct{}

func (UnimplementedUserReportControllerHandler) Send(context.Context, *connect_go.Request[user.UserReportRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.UserReportController.Send is not implemented"))
}
