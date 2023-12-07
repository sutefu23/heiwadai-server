// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: v1/admin/Auth.proto

package adminconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	admin "server/api/v1/admin"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AuthControllerName is the fully-qualified name of the AuthController service.
	AuthControllerName = "server.admin.AuthController"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AuthControllerRegisterProcedure is the fully-qualified name of the AuthController's Register RPC.
	AuthControllerRegisterProcedure = "/server.admin.AuthController/Register"
	// AuthControllerSignUpProcedure is the fully-qualified name of the AuthController's SignUp RPC.
	AuthControllerSignUpProcedure = "/server.admin.AuthController/SignUp"
	// AuthControllerSignOutProcedure is the fully-qualified name of the AuthController's SignOut RPC.
	AuthControllerSignOutProcedure = "/server.admin.AuthController/SignOut"
	// AuthControllerRefreshProcedure is the fully-qualified name of the AuthController's Refresh RPC.
	AuthControllerRefreshProcedure = "/server.admin.AuthController/Refresh"
	// AuthControllerSignInProcedure is the fully-qualified name of the AuthController's SignIn RPC.
	AuthControllerSignInProcedure = "/server.admin.AuthController/SignIn"
	// AuthControllerResetPasswordMailProcedure is the fully-qualified name of the AuthController's
	// ResetPasswordMail RPC.
	AuthControllerResetPasswordMailProcedure = "/server.admin.AuthController/ResetPasswordMail"
	// AuthControllerUpdatePasswordProcedure is the fully-qualified name of the AuthController's
	// UpdatePassword RPC.
	AuthControllerUpdatePasswordProcedure = "/server.admin.AuthController/UpdatePassword"
	// AuthControllerUpdateEmailProcedure is the fully-qualified name of the AuthController's
	// UpdateEmail RPC.
	AuthControllerUpdateEmailProcedure = "/server.admin.AuthController/UpdateEmail"
)

// AuthControllerClient is a client for the server.admin.AuthController service.
type AuthControllerClient interface {
	Register(context.Context, *connect_go.Request[admin.AdminRegisterRequest]) (*connect_go.Response[emptypb.Empty], error)
	SignUp(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error)
	SignOut(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	Refresh(context.Context, *connect_go.Request[admin.AdminRefreshTokenRequest]) (*connect_go.Response[admin.AdminAuthResponse], error)
	SignIn(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AdminAuthResponse], error)
	ResetPasswordMail(context.Context, *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdatePassword(context.Context, *connect_go.Request[admin.UpdatePasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdateEmail(context.Context, *connect_go.Request[admin.UpdateEmailRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAuthControllerClient constructs a client for the server.admin.AuthController service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthControllerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthControllerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authControllerClient{
		register: connect_go.NewClient[admin.AdminRegisterRequest, emptypb.Empty](
			httpClient,
			baseURL+AuthControllerRegisterProcedure,
			opts...,
		),
		signUp: connect_go.NewClient[admin.AdminAuthRequest, emptypb.Empty](
			httpClient,
			baseURL+AuthControllerSignUpProcedure,
			opts...,
		),
		signOut: connect_go.NewClient[emptypb.Empty, emptypb.Empty](
			httpClient,
			baseURL+AuthControllerSignOutProcedure,
			opts...,
		),
		refresh: connect_go.NewClient[admin.AdminRefreshTokenRequest, admin.AdminAuthResponse](
			httpClient,
			baseURL+AuthControllerRefreshProcedure,
			opts...,
		),
		signIn: connect_go.NewClient[admin.AdminAuthRequest, admin.AdminAuthResponse](
			httpClient,
			baseURL+AuthControllerSignInProcedure,
			opts...,
		),
		resetPasswordMail: connect_go.NewClient[admin.ResetPasswordRequest, emptypb.Empty](
			httpClient,
			baseURL+AuthControllerResetPasswordMailProcedure,
			opts...,
		),
		updatePassword: connect_go.NewClient[admin.UpdatePasswordRequest, emptypb.Empty](
			httpClient,
			baseURL+AuthControllerUpdatePasswordProcedure,
			opts...,
		),
		updateEmail: connect_go.NewClient[admin.UpdateEmailRequest, emptypb.Empty](
			httpClient,
			baseURL+AuthControllerUpdateEmailProcedure,
			opts...,
		),
	}
}

// authControllerClient implements AuthControllerClient.
type authControllerClient struct {
	register          *connect_go.Client[admin.AdminRegisterRequest, emptypb.Empty]
	signUp            *connect_go.Client[admin.AdminAuthRequest, emptypb.Empty]
	signOut           *connect_go.Client[emptypb.Empty, emptypb.Empty]
	refresh           *connect_go.Client[admin.AdminRefreshTokenRequest, admin.AdminAuthResponse]
	signIn            *connect_go.Client[admin.AdminAuthRequest, admin.AdminAuthResponse]
	resetPasswordMail *connect_go.Client[admin.ResetPasswordRequest, emptypb.Empty]
	updatePassword    *connect_go.Client[admin.UpdatePasswordRequest, emptypb.Empty]
	updateEmail       *connect_go.Client[admin.UpdateEmailRequest, emptypb.Empty]
}

// Register calls server.admin.AuthController.Register.
func (c *authControllerClient) Register(ctx context.Context, req *connect_go.Request[admin.AdminRegisterRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.register.CallUnary(ctx, req)
}

// SignUp calls server.admin.AuthController.SignUp.
func (c *authControllerClient) SignUp(ctx context.Context, req *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.signUp.CallUnary(ctx, req)
}

// SignOut calls server.admin.AuthController.SignOut.
func (c *authControllerClient) SignOut(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return c.signOut.CallUnary(ctx, req)
}

// Refresh calls server.admin.AuthController.Refresh.
func (c *authControllerClient) Refresh(ctx context.Context, req *connect_go.Request[admin.AdminRefreshTokenRequest]) (*connect_go.Response[admin.AdminAuthResponse], error) {
	return c.refresh.CallUnary(ctx, req)
}

// SignIn calls server.admin.AuthController.SignIn.
func (c *authControllerClient) SignIn(ctx context.Context, req *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AdminAuthResponse], error) {
	return c.signIn.CallUnary(ctx, req)
}

// ResetPasswordMail calls server.admin.AuthController.ResetPasswordMail.
func (c *authControllerClient) ResetPasswordMail(ctx context.Context, req *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.resetPasswordMail.CallUnary(ctx, req)
}

// UpdatePassword calls server.admin.AuthController.UpdatePassword.
func (c *authControllerClient) UpdatePassword(ctx context.Context, req *connect_go.Request[admin.UpdatePasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updatePassword.CallUnary(ctx, req)
}

// UpdateEmail calls server.admin.AuthController.UpdateEmail.
func (c *authControllerClient) UpdateEmail(ctx context.Context, req *connect_go.Request[admin.UpdateEmailRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updateEmail.CallUnary(ctx, req)
}

// AuthControllerHandler is an implementation of the server.admin.AuthController service.
type AuthControllerHandler interface {
	Register(context.Context, *connect_go.Request[admin.AdminRegisterRequest]) (*connect_go.Response[emptypb.Empty], error)
	SignUp(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error)
	SignOut(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error)
	Refresh(context.Context, *connect_go.Request[admin.AdminRefreshTokenRequest]) (*connect_go.Response[admin.AdminAuthResponse], error)
	SignIn(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AdminAuthResponse], error)
	ResetPasswordMail(context.Context, *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdatePassword(context.Context, *connect_go.Request[admin.UpdatePasswordRequest]) (*connect_go.Response[emptypb.Empty], error)
	UpdateEmail(context.Context, *connect_go.Request[admin.UpdateEmailRequest]) (*connect_go.Response[emptypb.Empty], error)
}

// NewAuthControllerHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthControllerHandler(svc AuthControllerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	authControllerRegisterHandler := connect_go.NewUnaryHandler(
		AuthControllerRegisterProcedure,
		svc.Register,
		opts...,
	)
	authControllerSignUpHandler := connect_go.NewUnaryHandler(
		AuthControllerSignUpProcedure,
		svc.SignUp,
		opts...,
	)
	authControllerSignOutHandler := connect_go.NewUnaryHandler(
		AuthControllerSignOutProcedure,
		svc.SignOut,
		opts...,
	)
	authControllerRefreshHandler := connect_go.NewUnaryHandler(
		AuthControllerRefreshProcedure,
		svc.Refresh,
		opts...,
	)
	authControllerSignInHandler := connect_go.NewUnaryHandler(
		AuthControllerSignInProcedure,
		svc.SignIn,
		opts...,
	)
	authControllerResetPasswordMailHandler := connect_go.NewUnaryHandler(
		AuthControllerResetPasswordMailProcedure,
		svc.ResetPasswordMail,
		opts...,
	)
	authControllerUpdatePasswordHandler := connect_go.NewUnaryHandler(
		AuthControllerUpdatePasswordProcedure,
		svc.UpdatePassword,
		opts...,
	)
	authControllerUpdateEmailHandler := connect_go.NewUnaryHandler(
		AuthControllerUpdateEmailProcedure,
		svc.UpdateEmail,
		opts...,
	)
	return "/server.admin.AuthController/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AuthControllerRegisterProcedure:
			authControllerRegisterHandler.ServeHTTP(w, r)
		case AuthControllerSignUpProcedure:
			authControllerSignUpHandler.ServeHTTP(w, r)
		case AuthControllerSignOutProcedure:
			authControllerSignOutHandler.ServeHTTP(w, r)
		case AuthControllerRefreshProcedure:
			authControllerRefreshHandler.ServeHTTP(w, r)
		case AuthControllerSignInProcedure:
			authControllerSignInHandler.ServeHTTP(w, r)
		case AuthControllerResetPasswordMailProcedure:
			authControllerResetPasswordMailHandler.ServeHTTP(w, r)
		case AuthControllerUpdatePasswordProcedure:
			authControllerUpdatePasswordHandler.ServeHTTP(w, r)
		case AuthControllerUpdateEmailProcedure:
			authControllerUpdateEmailHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAuthControllerHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthControllerHandler struct{}

func (UnimplementedAuthControllerHandler) Register(context.Context, *connect_go.Request[admin.AdminRegisterRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.Register is not implemented"))
}

func (UnimplementedAuthControllerHandler) SignUp(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.SignUp is not implemented"))
}

func (UnimplementedAuthControllerHandler) SignOut(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.SignOut is not implemented"))
}

func (UnimplementedAuthControllerHandler) Refresh(context.Context, *connect_go.Request[admin.AdminRefreshTokenRequest]) (*connect_go.Response[admin.AdminAuthResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.Refresh is not implemented"))
}

func (UnimplementedAuthControllerHandler) SignIn(context.Context, *connect_go.Request[admin.AdminAuthRequest]) (*connect_go.Response[admin.AdminAuthResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.SignIn is not implemented"))
}

func (UnimplementedAuthControllerHandler) ResetPasswordMail(context.Context, *connect_go.Request[admin.ResetPasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.ResetPasswordMail is not implemented"))
}

func (UnimplementedAuthControllerHandler) UpdatePassword(context.Context, *connect_go.Request[admin.UpdatePasswordRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.UpdatePassword is not implemented"))
}

func (UnimplementedAuthControllerHandler) UpdateEmail(context.Context, *connect_go.Request[admin.UpdateEmailRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("server.admin.AuthController.UpdateEmail is not implemented"))
}
