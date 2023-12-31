// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: blog/v1/posts.proto

package blogv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	http "net/http"
	v1 "starter/apis/blog/v1"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// BlogServiceName is the fully-qualified name of the BlogService service.
	BlogServiceName = "blog.v1.BlogService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BlogServiceCreatePostProcedure is the fully-qualified name of the BlogService's CreatePost RPC.
	BlogServiceCreatePostProcedure = "/blog.v1.BlogService/CreatePost"
	// BlogServiceListPostsProcedure is the fully-qualified name of the BlogService's ListPosts RPC.
	BlogServiceListPostsProcedure = "/blog.v1.BlogService/ListPosts"
	// BlogServiceDeletePostProcedure is the fully-qualified name of the BlogService's DeletePost RPC.
	BlogServiceDeletePostProcedure = "/blog.v1.BlogService/DeletePost"
)

// BlogServiceClient is a client for the blog.v1.BlogService service.
type BlogServiceClient interface {
	CreatePost(context.Context, *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error)
	ListPosts(context.Context, *connect.Request[v1.ListPostsRequest]) (*connect.Response[v1.ListPostsResponse], error)
	DeletePost(context.Context, *connect.Request[v1.DeletePostRequest]) (*connect.Response[v1.DeletePostResponse], error)
}

// NewBlogServiceClient constructs a client for the blog.v1.BlogService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBlogServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) BlogServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &blogServiceClient{
		createPost: connect.NewClient[v1.CreatePostRequest, v1.CreatePostResponse](
			httpClient,
			baseURL+BlogServiceCreatePostProcedure,
			opts...,
		),
		listPosts: connect.NewClient[v1.ListPostsRequest, v1.ListPostsResponse](
			httpClient,
			baseURL+BlogServiceListPostsProcedure,
			opts...,
		),
		deletePost: connect.NewClient[v1.DeletePostRequest, v1.DeletePostResponse](
			httpClient,
			baseURL+BlogServiceDeletePostProcedure,
			opts...,
		),
	}
}

// blogServiceClient implements BlogServiceClient.
type blogServiceClient struct {
	createPost *connect.Client[v1.CreatePostRequest, v1.CreatePostResponse]
	listPosts  *connect.Client[v1.ListPostsRequest, v1.ListPostsResponse]
	deletePost *connect.Client[v1.DeletePostRequest, v1.DeletePostResponse]
}

// CreatePost calls blog.v1.BlogService.CreatePost.
func (c *blogServiceClient) CreatePost(ctx context.Context, req *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error) {
	return c.createPost.CallUnary(ctx, req)
}

// ListPosts calls blog.v1.BlogService.ListPosts.
func (c *blogServiceClient) ListPosts(ctx context.Context, req *connect.Request[v1.ListPostsRequest]) (*connect.Response[v1.ListPostsResponse], error) {
	return c.listPosts.CallUnary(ctx, req)
}

// DeletePost calls blog.v1.BlogService.DeletePost.
func (c *blogServiceClient) DeletePost(ctx context.Context, req *connect.Request[v1.DeletePostRequest]) (*connect.Response[v1.DeletePostResponse], error) {
	return c.deletePost.CallUnary(ctx, req)
}

// BlogServiceHandler is an implementation of the blog.v1.BlogService service.
type BlogServiceHandler interface {
	CreatePost(context.Context, *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error)
	ListPosts(context.Context, *connect.Request[v1.ListPostsRequest]) (*connect.Response[v1.ListPostsResponse], error)
	DeletePost(context.Context, *connect.Request[v1.DeletePostRequest]) (*connect.Response[v1.DeletePostResponse], error)
}

// NewBlogServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBlogServiceHandler(svc BlogServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	blogServiceCreatePostHandler := connect.NewUnaryHandler(
		BlogServiceCreatePostProcedure,
		svc.CreatePost,
		opts...,
	)
	blogServiceListPostsHandler := connect.NewUnaryHandler(
		BlogServiceListPostsProcedure,
		svc.ListPosts,
		opts...,
	)
	blogServiceDeletePostHandler := connect.NewUnaryHandler(
		BlogServiceDeletePostProcedure,
		svc.DeletePost,
		opts...,
	)
	return "/blog.v1.BlogService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BlogServiceCreatePostProcedure:
			blogServiceCreatePostHandler.ServeHTTP(w, r)
		case BlogServiceListPostsProcedure:
			blogServiceListPostsHandler.ServeHTTP(w, r)
		case BlogServiceDeletePostProcedure:
			blogServiceDeletePostHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBlogServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBlogServiceHandler struct{}

func (UnimplementedBlogServiceHandler) CreatePost(context.Context, *connect.Request[v1.CreatePostRequest]) (*connect.Response[v1.CreatePostResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("blog.v1.BlogService.CreatePost is not implemented"))
}

func (UnimplementedBlogServiceHandler) ListPosts(context.Context, *connect.Request[v1.ListPostsRequest]) (*connect.Response[v1.ListPostsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("blog.v1.BlogService.ListPosts is not implemented"))
}

func (UnimplementedBlogServiceHandler) DeletePost(context.Context, *connect.Request[v1.DeletePostRequest]) (*connect.Response[v1.DeletePostResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("blog.v1.BlogService.DeletePost is not implemented"))
}
