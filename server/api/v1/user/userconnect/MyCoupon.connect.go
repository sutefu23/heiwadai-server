// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/user/MyCoupon.proto

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
	// MyCouponControllerName is the fully-qualified name of the MyCouponController service.
	MyCouponControllerName = "server.user.MyCouponController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MyCouponControllerGetDetailProcedure is the fully-qualified name of the MyCouponController's
	// GetDetail RPC.
	MyCouponControllerGetDetailProcedure = "/server.user.MyCouponController/GetDetail"
	// MyCouponControllerGetListProcedure is the fully-qualified name of the MyCouponController's
	// GetList RPC.
	MyCouponControllerGetListProcedure = "/server.user.MyCouponController/GetList"
	// MyCouponControllerUseProcedure is the fully-qualified name of the MyCouponController's Use RPC.
	MyCouponControllerUseProcedure = "/server.user.MyCouponController/Use"
)

// MyCouponControllerClient is a client for the server.user.MyCouponController service.
type MyCouponControllerClient interface {
	GetDetail(context.Context, *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error)
	GetList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.MyCouponsResponse], error)
	Use(context.Context, *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewMyCouponControllerClient constructs a client for the server.user.MyCouponController service.
// By default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped
// responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMyCouponControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MyCouponControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &myCouponControllerClient{
		getDetail: connect_go.NewClient[user.CouponIDRequest, shared.Coupon](
			httpClient,
			baseURL+MyCouponControllerGetDetailProcedure,
			opts...,
		),
		getList: connect_go.NewClient[emptypb.Empty, user.MyCouponsResponse](
			httpClient,
			baseURL+MyCouponControllerGetListProcedure,
			opts...,
		),
		use: connect_go.NewClient[user.CouponIDRequest, emptypb.Empty](
			httpClient,
			baseURL+MyCouponControllerUseProcedure,
			opts...,
		),
	}
}

// myCouponControllerClient implements MyCouponControllerClient.
type myCouponControllerClient struct {
	getDetail *connect_go.Client[user.CouponIDRequest, shared.Coupon]
	getList   *connect_go.Client[emptypb.Empty, user.MyCouponsResponse]
	use       *connect_go.Client[user.CouponIDRequest, emptypb.Empty]
}

// GetDetail calls server.user.MyCouponController.GetDetail.
func (c *myCouponControllerClient) GetDetail(ctx context.Context, req *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error) {
	return c.getDetail.CallUnary(ctx, req)
}

// GetList calls server.user.MyCouponController.GetList.
func (c *myCouponControllerClient) GetList(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.MyCouponsResponse], error) {
	return c.getList.CallUnary(ctx, req)
}

// Use calls server.user.MyCouponController.Use.
func (c *myCouponControllerClient) Use(ctx context.Context, req *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.use.CallUnary(ctx, req)
}

// MyCouponControllerHandler is an implementation of the server.user.MyCouponController service.
type MyCouponControllerHandler interface {
	GetDetail(context.Context, *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error)
	GetList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.MyCouponsResponse], error)
	Use(context.Context, *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewMyCouponControllerHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMyCouponControllerHandler(svc MyCouponControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	myCouponControllerGetDetailHandler := connect_go.NewUnaryHandler(
		MyCouponControllerGetDetailProcedure,
		svc.GetDetail,
		opts...,
	)
	myCouponControllerGetListHandler := connect_go.NewUnaryHandler(
		MyCouponControllerGetListProcedure,
		svc.GetList,
		opts...,
	)
	myCouponControllerUseHandler := connect_go.NewUnaryHandler(
		MyCouponControllerUseProcedure,
		svc.Use,
		opts...,
	)
	return "/server.user.MyCouponController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MyCouponControllerGetDetailProcedure:
			myCouponControllerGetDetailHandler.ServeHTTP(w, r)
		case MyCouponControllerGetListProcedure:
			myCouponControllerGetListHandler.ServeHTTP(w, r)
		case MyCouponControllerUseProcedure:
			myCouponControllerUseHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMyCouponControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedMyCouponControllerHandler struct{}

func (UnimplementedMyCouponControllerHandler) GetDetail(context.Context, *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[shared.Coupon], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.MyCouponController.GetDetail is not implemented"))
}

func (UnimplementedMyCouponControllerHandler) GetList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[user.MyCouponsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.MyCouponController.GetList is not implemented"))
}

func (UnimplementedMyCouponControllerHandler) Use(context.Context, *connect_go.Request[user.CouponIDRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.user.MyCouponController.Use is not implemented"))
}
