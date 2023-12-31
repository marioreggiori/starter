// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file blog/v1/posts.proto (package blog.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreatePostRequest, CreatePostResponse, DeletePostRequest, DeletePostResponse, ListPostsRequest, ListPostsResponse } from "./posts_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service blog.v1.BlogService
 */
export const BlogService = {
  typeName: "blog.v1.BlogService",
  methods: {
    /**
     * @generated from rpc blog.v1.BlogService.CreatePost
     */
    createPost: {
      name: "CreatePost",
      I: CreatePostRequest,
      O: CreatePostResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc blog.v1.BlogService.ListPosts
     */
    listPosts: {
      name: "ListPosts",
      I: ListPostsRequest,
      O: ListPostsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc blog.v1.BlogService.DeletePost
     */
    deletePost: {
      name: "DeletePost",
      I: DeletePostRequest,
      O: DeletePostResponse,
      kind: MethodKind.Unary,
    },
  }
};

