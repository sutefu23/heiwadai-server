// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/user/Store.proto

package userconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	shared "server/api/v1/shared"
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
	// StoreControllerName is the fully-qualified name of the StoreController service.
	StoreControllerName = "server.user.StoreController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StoreControllerGetByIDProcedure is the fully-qualified name of the StoreController's GetByID RPC.
	StoreControllerGetByIDProcedure = "/server.user.StoreController/GetByID"
	// StoreControllerGetAllProcedure is the fully-qualified name of the StoreController's GetAll RPC.
	StoreControllerGetAllProcedure = "/server.user.StoreController/GetAll"
	// StoreControllerGetStayablesProcedure is the fully-qualified name of the StoreController's
	// GetStayables RPC.
	StoreControllerGetStayablesProcedure = "/server.user.StoreController/GetStayables"
	// StoreControllerGetStayableByIDProcedure is the fully-qualified name of the StoreController's
	// GetStayableByID RPC.
	StoreControllerGetStayableByIDProcedure = "/server.user.StoreController/GetStayableByID"
)

// StoreControllerClient is a client for the server.user.StoreController service.
type StoreControllerClient interface {
	GetByID(context.Context, *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.Store], error)
	GetAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.Stores], error)
	GetStayables(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.StayableStores], error)
	GetStayableByID(context.Context, *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.StayableStore], error)
}

// NewStoreControllerClient constructs a client for the server.user.StoreController service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStoreControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) StoreControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &storeControllerClient{
		getByID: connect_go.NewClient[user.SoreIDRequest, shared.Store](
			httpClient,
			baseURL+StoreControllerGetByIDProcedure,
			opts...,
		),
		getAll: connect_go.NewClient[emptypb.Empty, shared.Stores](
			httpClient,
			baseURL+StoreControllerGetAllProcedure,
			opts...,
		),
		getStayables: connect_go.NewClient[emptypb.Empty, shared.StayableStores](
			httpClient,
			baseURL+StoreControllerGetStayablesProcedure,
			opts...,
		),
		getStayableByID: connect_go.NewClient[user.SoreIDRequest, shared.StayableStore](
			httpClient,
			baseURL+StoreControllerGetStayableByIDProcedure,
			opts...,
		),
	}
}

// storeControllerClient implements StoreControllerClient.
type storeControllerClient struct {
	getByID         *connect_go.Client[user.SoreIDRequest, shared.Store]
	getAll          *connect_go.Client[emptypb.Empty, shared.Stores]
	getStayables    *connect_go.Client[emptypb.Empty, shared.StayableStores]
	getStayableByID *connect_go.Client[user.SoreIDRequest, shared.StayableStore]
}

// GetByID calls server.user.StoreController.GetByID.
func (c *storeControllerClient) GetByID(ctx context.Context, req *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.Store], error) {
	return c.getByID.CallUnary(ctx, req)
}

// GetAll calls server.user.StoreController.GetAll.
func (c *storeControllerClient) GetAll(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.Stores], error) {
	return c.getAll.CallUnary(ctx, req)
}

// GetStayables calls server.user.StoreController.GetStayables.
func (c *storeControllerClient) GetStayables(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.StayableStores], error) {
	return c.getStayables.CallUnary(ctx, req)
}

// GetStayableByID calls server.user.StoreController.GetStayableByID.
func (c *storeControllerClient) GetStayableByID(ctx context.Context, req *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.StayableStore], error) {
	return c.getStayableByID.CallUnary(ctx, req)
}

// StoreControllerHandler is an implementation of the server.user.StoreController service.
type StoreControllerHandler interface {
	GetByID(context.Context, *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.Store], error)
	GetAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.Stores], error)
	GetStayables(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.StayableStores], error)
	GetStayableByID(context.Context, *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.StayableStore], error)
}

// NewStoreControllerHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStoreControllerHandler(svc StoreControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	storeControllerGetByIDHandler := connect_go.NewUnaryHandler(
		StoreControllerGetByIDProcedure,
		svc.GetByID,
		opts...,
	)
	storeControllerGetAllHandler := connect_go.NewUnaryHandler(
		StoreControllerGetAllProcedure,
		svc.GetAll,
		opts...,
	)
	storeControllerGetStayablesHandler := connect_go.NewUnaryHandler(
		StoreControllerGetStayablesProcedure,
		svc.GetStayables,
		opts...,
	)
	storeControllerGetStayableByIDHandler := connect_go.NewUnaryHandler(
		StoreControllerGetStayableByIDProcedure,
		svc.GetStayableByID,
		opts...,
	)
	return "/server.user.StoreController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StoreControllerGetByIDProcedure:
			storeControllerGetByIDHandler.ServeHTTP(w, r)
		case StoreControllerGetAllProcedure:
			storeControllerGetAllHandler.ServeHTTP(w, r)
		case StoreControllerGetStayablesProcedure:
			storeControllerGetStayablesHandler.ServeHTTP(w, r)
		case StoreControllerGetStayableByIDProcedure:
			storeControllerGetStayableByIDHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStoreControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedStoreControllerHandler struct{}

func (UnimplementedStoreControllerHandler) GetByID(context.Context, *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.Store], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.StoreController.GetByID is not implemented"))
}

func (UnimplementedStoreControllerHandler) GetAll(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.Stores], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.StoreController.GetAll is not implemented"))
}

func (UnimplementedStoreControllerHandler) GetStayables(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[shared.StayableStores], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.StoreController.GetStayables is not implemented"))
}

func (UnimplementedStoreControllerHandler) GetStayableByID(context.Context, *connect_go.Request[user.SoreIDRequest]) (*connect_go.Response[shared.StayableStore], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.StoreController.GetStayableByID is not implemented"))
}
