// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mgmt/v1alpha1/transformer.proto

package mgmtv1alpha1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// TransformersServiceName is the fully-qualified name of the TransformersService service.
	TransformersServiceName = "mgmt.v1alpha1.TransformersService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TransformersServiceGetSystemTransformersProcedure is the fully-qualified name of the
	// TransformersService's GetSystemTransformers RPC.
	TransformersServiceGetSystemTransformersProcedure = "/mgmt.v1alpha1.TransformersService/GetSystemTransformers"
	// TransformersServiceGetUserDefinedTransformersProcedure is the fully-qualified name of the
	// TransformersService's GetUserDefinedTransformers RPC.
	TransformersServiceGetUserDefinedTransformersProcedure = "/mgmt.v1alpha1.TransformersService/GetUserDefinedTransformers"
	// TransformersServiceGetUserDefinedTransformerByIdProcedure is the fully-qualified name of the
	// TransformersService's GetUserDefinedTransformerById RPC.
	TransformersServiceGetUserDefinedTransformerByIdProcedure = "/mgmt.v1alpha1.TransformersService/GetUserDefinedTransformerById"
	// TransformersServiceCreateUserDefinedTransformerProcedure is the fully-qualified name of the
	// TransformersService's CreateUserDefinedTransformer RPC.
	TransformersServiceCreateUserDefinedTransformerProcedure = "/mgmt.v1alpha1.TransformersService/CreateUserDefinedTransformer"
	// TransformersServiceDeleteUserDefinedTransformerProcedure is the fully-qualified name of the
	// TransformersService's DeleteUserDefinedTransformer RPC.
	TransformersServiceDeleteUserDefinedTransformerProcedure = "/mgmt.v1alpha1.TransformersService/DeleteUserDefinedTransformer"
	// TransformersServiceUpdateUserDefinedTransformerProcedure is the fully-qualified name of the
	// TransformersService's UpdateUserDefinedTransformer RPC.
	TransformersServiceUpdateUserDefinedTransformerProcedure = "/mgmt.v1alpha1.TransformersService/UpdateUserDefinedTransformer"
	// TransformersServiceIsTransformerNameAvailableProcedure is the fully-qualified name of the
	// TransformersService's IsTransformerNameAvailable RPC.
	TransformersServiceIsTransformerNameAvailableProcedure = "/mgmt.v1alpha1.TransformersService/IsTransformerNameAvailable"
	// TransformersServiceValidateUserJavascriptCodeProcedure is the fully-qualified name of the
	// TransformersService's ValidateUserJavascriptCode RPC.
	TransformersServiceValidateUserJavascriptCodeProcedure = "/mgmt.v1alpha1.TransformersService/ValidateUserJavascriptCode"
)

// TransformersServiceClient is a client for the mgmt.v1alpha1.TransformersService service.
type TransformersServiceClient interface {
	GetSystemTransformers(context.Context, *connect.Request[v1alpha1.GetSystemTransformersRequest]) (*connect.Response[v1alpha1.GetSystemTransformersResponse], error)
	GetUserDefinedTransformers(context.Context, *connect.Request[v1alpha1.GetUserDefinedTransformersRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformersResponse], error)
	GetUserDefinedTransformerById(context.Context, *connect.Request[v1alpha1.GetUserDefinedTransformerByIdRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformerByIdResponse], error)
	CreateUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.CreateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.CreateUserDefinedTransformerResponse], error)
	DeleteUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.DeleteUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.DeleteUserDefinedTransformerResponse], error)
	UpdateUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.UpdateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.UpdateUserDefinedTransformerResponse], error)
	IsTransformerNameAvailable(context.Context, *connect.Request[v1alpha1.IsTransformerNameAvailableRequest]) (*connect.Response[v1alpha1.IsTransformerNameAvailableResponse], error)
	ValidateUserJavascriptCode(context.Context, *connect.Request[v1alpha1.ValidateUserJavascriptCodeRequest]) (*connect.Response[v1alpha1.ValidateUserJavascriptCodeResponse], error)
}

// NewTransformersServiceClient constructs a client for the mgmt.v1alpha1.TransformersService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTransformersServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) TransformersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &transformersServiceClient{
		getSystemTransformers: connect.NewClient[v1alpha1.GetSystemTransformersRequest, v1alpha1.GetSystemTransformersResponse](
			httpClient,
			baseURL+TransformersServiceGetSystemTransformersProcedure,
			opts...,
		),
		getUserDefinedTransformers: connect.NewClient[v1alpha1.GetUserDefinedTransformersRequest, v1alpha1.GetUserDefinedTransformersResponse](
			httpClient,
			baseURL+TransformersServiceGetUserDefinedTransformersProcedure,
			opts...,
		),
		getUserDefinedTransformerById: connect.NewClient[v1alpha1.GetUserDefinedTransformerByIdRequest, v1alpha1.GetUserDefinedTransformerByIdResponse](
			httpClient,
			baseURL+TransformersServiceGetUserDefinedTransformerByIdProcedure,
			opts...,
		),
		createUserDefinedTransformer: connect.NewClient[v1alpha1.CreateUserDefinedTransformerRequest, v1alpha1.CreateUserDefinedTransformerResponse](
			httpClient,
			baseURL+TransformersServiceCreateUserDefinedTransformerProcedure,
			opts...,
		),
		deleteUserDefinedTransformer: connect.NewClient[v1alpha1.DeleteUserDefinedTransformerRequest, v1alpha1.DeleteUserDefinedTransformerResponse](
			httpClient,
			baseURL+TransformersServiceDeleteUserDefinedTransformerProcedure,
			opts...,
		),
		updateUserDefinedTransformer: connect.NewClient[v1alpha1.UpdateUserDefinedTransformerRequest, v1alpha1.UpdateUserDefinedTransformerResponse](
			httpClient,
			baseURL+TransformersServiceUpdateUserDefinedTransformerProcedure,
			opts...,
		),
		isTransformerNameAvailable: connect.NewClient[v1alpha1.IsTransformerNameAvailableRequest, v1alpha1.IsTransformerNameAvailableResponse](
			httpClient,
			baseURL+TransformersServiceIsTransformerNameAvailableProcedure,
			opts...,
		),
		validateUserJavascriptCode: connect.NewClient[v1alpha1.ValidateUserJavascriptCodeRequest, v1alpha1.ValidateUserJavascriptCodeResponse](
			httpClient,
			baseURL+TransformersServiceValidateUserJavascriptCodeProcedure,
			opts...,
		),
	}
}

// transformersServiceClient implements TransformersServiceClient.
type transformersServiceClient struct {
	getSystemTransformers         *connect.Client[v1alpha1.GetSystemTransformersRequest, v1alpha1.GetSystemTransformersResponse]
	getUserDefinedTransformers    *connect.Client[v1alpha1.GetUserDefinedTransformersRequest, v1alpha1.GetUserDefinedTransformersResponse]
	getUserDefinedTransformerById *connect.Client[v1alpha1.GetUserDefinedTransformerByIdRequest, v1alpha1.GetUserDefinedTransformerByIdResponse]
	createUserDefinedTransformer  *connect.Client[v1alpha1.CreateUserDefinedTransformerRequest, v1alpha1.CreateUserDefinedTransformerResponse]
	deleteUserDefinedTransformer  *connect.Client[v1alpha1.DeleteUserDefinedTransformerRequest, v1alpha1.DeleteUserDefinedTransformerResponse]
	updateUserDefinedTransformer  *connect.Client[v1alpha1.UpdateUserDefinedTransformerRequest, v1alpha1.UpdateUserDefinedTransformerResponse]
	isTransformerNameAvailable    *connect.Client[v1alpha1.IsTransformerNameAvailableRequest, v1alpha1.IsTransformerNameAvailableResponse]
	validateUserJavascriptCode    *connect.Client[v1alpha1.ValidateUserJavascriptCodeRequest, v1alpha1.ValidateUserJavascriptCodeResponse]
}

// GetSystemTransformers calls mgmt.v1alpha1.TransformersService.GetSystemTransformers.
func (c *transformersServiceClient) GetSystemTransformers(ctx context.Context, req *connect.Request[v1alpha1.GetSystemTransformersRequest]) (*connect.Response[v1alpha1.GetSystemTransformersResponse], error) {
	return c.getSystemTransformers.CallUnary(ctx, req)
}

// GetUserDefinedTransformers calls mgmt.v1alpha1.TransformersService.GetUserDefinedTransformers.
func (c *transformersServiceClient) GetUserDefinedTransformers(ctx context.Context, req *connect.Request[v1alpha1.GetUserDefinedTransformersRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformersResponse], error) {
	return c.getUserDefinedTransformers.CallUnary(ctx, req)
}

// GetUserDefinedTransformerById calls
// mgmt.v1alpha1.TransformersService.GetUserDefinedTransformerById.
func (c *transformersServiceClient) GetUserDefinedTransformerById(ctx context.Context, req *connect.Request[v1alpha1.GetUserDefinedTransformerByIdRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformerByIdResponse], error) {
	return c.getUserDefinedTransformerById.CallUnary(ctx, req)
}

// CreateUserDefinedTransformer calls
// mgmt.v1alpha1.TransformersService.CreateUserDefinedTransformer.
func (c *transformersServiceClient) CreateUserDefinedTransformer(ctx context.Context, req *connect.Request[v1alpha1.CreateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.CreateUserDefinedTransformerResponse], error) {
	return c.createUserDefinedTransformer.CallUnary(ctx, req)
}

// DeleteUserDefinedTransformer calls
// mgmt.v1alpha1.TransformersService.DeleteUserDefinedTransformer.
func (c *transformersServiceClient) DeleteUserDefinedTransformer(ctx context.Context, req *connect.Request[v1alpha1.DeleteUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.DeleteUserDefinedTransformerResponse], error) {
	return c.deleteUserDefinedTransformer.CallUnary(ctx, req)
}

// UpdateUserDefinedTransformer calls
// mgmt.v1alpha1.TransformersService.UpdateUserDefinedTransformer.
func (c *transformersServiceClient) UpdateUserDefinedTransformer(ctx context.Context, req *connect.Request[v1alpha1.UpdateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.UpdateUserDefinedTransformerResponse], error) {
	return c.updateUserDefinedTransformer.CallUnary(ctx, req)
}

// IsTransformerNameAvailable calls mgmt.v1alpha1.TransformersService.IsTransformerNameAvailable.
func (c *transformersServiceClient) IsTransformerNameAvailable(ctx context.Context, req *connect.Request[v1alpha1.IsTransformerNameAvailableRequest]) (*connect.Response[v1alpha1.IsTransformerNameAvailableResponse], error) {
	return c.isTransformerNameAvailable.CallUnary(ctx, req)
}

// ValidateUserJavascriptCode calls mgmt.v1alpha1.TransformersService.ValidateUserJavascriptCode.
func (c *transformersServiceClient) ValidateUserJavascriptCode(ctx context.Context, req *connect.Request[v1alpha1.ValidateUserJavascriptCodeRequest]) (*connect.Response[v1alpha1.ValidateUserJavascriptCodeResponse], error) {
	return c.validateUserJavascriptCode.CallUnary(ctx, req)
}

// TransformersServiceHandler is an implementation of the mgmt.v1alpha1.TransformersService service.
type TransformersServiceHandler interface {
	GetSystemTransformers(context.Context, *connect.Request[v1alpha1.GetSystemTransformersRequest]) (*connect.Response[v1alpha1.GetSystemTransformersResponse], error)
	GetUserDefinedTransformers(context.Context, *connect.Request[v1alpha1.GetUserDefinedTransformersRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformersResponse], error)
	GetUserDefinedTransformerById(context.Context, *connect.Request[v1alpha1.GetUserDefinedTransformerByIdRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformerByIdResponse], error)
	CreateUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.CreateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.CreateUserDefinedTransformerResponse], error)
	DeleteUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.DeleteUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.DeleteUserDefinedTransformerResponse], error)
	UpdateUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.UpdateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.UpdateUserDefinedTransformerResponse], error)
	IsTransformerNameAvailable(context.Context, *connect.Request[v1alpha1.IsTransformerNameAvailableRequest]) (*connect.Response[v1alpha1.IsTransformerNameAvailableResponse], error)
	ValidateUserJavascriptCode(context.Context, *connect.Request[v1alpha1.ValidateUserJavascriptCodeRequest]) (*connect.Response[v1alpha1.ValidateUserJavascriptCodeResponse], error)
}

// NewTransformersServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTransformersServiceHandler(svc TransformersServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	transformersServiceGetSystemTransformersHandler := connect.NewUnaryHandler(
		TransformersServiceGetSystemTransformersProcedure,
		svc.GetSystemTransformers,
		opts...,
	)
	transformersServiceGetUserDefinedTransformersHandler := connect.NewUnaryHandler(
		TransformersServiceGetUserDefinedTransformersProcedure,
		svc.GetUserDefinedTransformers,
		opts...,
	)
	transformersServiceGetUserDefinedTransformerByIdHandler := connect.NewUnaryHandler(
		TransformersServiceGetUserDefinedTransformerByIdProcedure,
		svc.GetUserDefinedTransformerById,
		opts...,
	)
	transformersServiceCreateUserDefinedTransformerHandler := connect.NewUnaryHandler(
		TransformersServiceCreateUserDefinedTransformerProcedure,
		svc.CreateUserDefinedTransformer,
		opts...,
	)
	transformersServiceDeleteUserDefinedTransformerHandler := connect.NewUnaryHandler(
		TransformersServiceDeleteUserDefinedTransformerProcedure,
		svc.DeleteUserDefinedTransformer,
		opts...,
	)
	transformersServiceUpdateUserDefinedTransformerHandler := connect.NewUnaryHandler(
		TransformersServiceUpdateUserDefinedTransformerProcedure,
		svc.UpdateUserDefinedTransformer,
		opts...,
	)
	transformersServiceIsTransformerNameAvailableHandler := connect.NewUnaryHandler(
		TransformersServiceIsTransformerNameAvailableProcedure,
		svc.IsTransformerNameAvailable,
		opts...,
	)
	transformersServiceValidateUserJavascriptCodeHandler := connect.NewUnaryHandler(
		TransformersServiceValidateUserJavascriptCodeProcedure,
		svc.ValidateUserJavascriptCode,
		opts...,
	)
	return "/mgmt.v1alpha1.TransformersService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TransformersServiceGetSystemTransformersProcedure:
			transformersServiceGetSystemTransformersHandler.ServeHTTP(w, r)
		case TransformersServiceGetUserDefinedTransformersProcedure:
			transformersServiceGetUserDefinedTransformersHandler.ServeHTTP(w, r)
		case TransformersServiceGetUserDefinedTransformerByIdProcedure:
			transformersServiceGetUserDefinedTransformerByIdHandler.ServeHTTP(w, r)
		case TransformersServiceCreateUserDefinedTransformerProcedure:
			transformersServiceCreateUserDefinedTransformerHandler.ServeHTTP(w, r)
		case TransformersServiceDeleteUserDefinedTransformerProcedure:
			transformersServiceDeleteUserDefinedTransformerHandler.ServeHTTP(w, r)
		case TransformersServiceUpdateUserDefinedTransformerProcedure:
			transformersServiceUpdateUserDefinedTransformerHandler.ServeHTTP(w, r)
		case TransformersServiceIsTransformerNameAvailableProcedure:
			transformersServiceIsTransformerNameAvailableHandler.ServeHTTP(w, r)
		case TransformersServiceValidateUserJavascriptCodeProcedure:
			transformersServiceValidateUserJavascriptCodeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTransformersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTransformersServiceHandler struct{}

func (UnimplementedTransformersServiceHandler) GetSystemTransformers(context.Context, *connect.Request[v1alpha1.GetSystemTransformersRequest]) (*connect.Response[v1alpha1.GetSystemTransformersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.GetSystemTransformers is not implemented"))
}

func (UnimplementedTransformersServiceHandler) GetUserDefinedTransformers(context.Context, *connect.Request[v1alpha1.GetUserDefinedTransformersRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.GetUserDefinedTransformers is not implemented"))
}

func (UnimplementedTransformersServiceHandler) GetUserDefinedTransformerById(context.Context, *connect.Request[v1alpha1.GetUserDefinedTransformerByIdRequest]) (*connect.Response[v1alpha1.GetUserDefinedTransformerByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.GetUserDefinedTransformerById is not implemented"))
}

func (UnimplementedTransformersServiceHandler) CreateUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.CreateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.CreateUserDefinedTransformerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.CreateUserDefinedTransformer is not implemented"))
}

func (UnimplementedTransformersServiceHandler) DeleteUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.DeleteUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.DeleteUserDefinedTransformerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.DeleteUserDefinedTransformer is not implemented"))
}

func (UnimplementedTransformersServiceHandler) UpdateUserDefinedTransformer(context.Context, *connect.Request[v1alpha1.UpdateUserDefinedTransformerRequest]) (*connect.Response[v1alpha1.UpdateUserDefinedTransformerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.UpdateUserDefinedTransformer is not implemented"))
}

func (UnimplementedTransformersServiceHandler) IsTransformerNameAvailable(context.Context, *connect.Request[v1alpha1.IsTransformerNameAvailableRequest]) (*connect.Response[v1alpha1.IsTransformerNameAvailableResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.IsTransformerNameAvailable is not implemented"))
}

func (UnimplementedTransformersServiceHandler) ValidateUserJavascriptCode(context.Context, *connect.Request[v1alpha1.ValidateUserJavascriptCodeRequest]) (*connect.Response[v1alpha1.ValidateUserJavascriptCodeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mgmt.v1alpha1.TransformersService.ValidateUserJavascriptCode is not implemented"))
}
