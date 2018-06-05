// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/bottos-project/magiccube/service/data/proto/data.proto

/*
Package data is a generated protocol buffer package.

It is generated from these files:
	github.com/bottos-project/magiccube/service/data/proto/data.proto

It has these top-level messages:
	AllRequest
	Slice
	Url
	Node
	Ip
	Hash
	FileCheckRequest
	FileCheckResponse
	GetFileUploadURLRequest
	GetFileUploadURLResponse
	GetFileSliceUploadURLRequest
	GetFileSliceUploadURLResponse
	GetFileStorageNodeRequest
	GetFileStorageNodeResponse
	PutFileRequest
	PutFileResponse
	GetUploadProgressRequest
	GetUploadProgressResponse
	DownloadFileRequest
	DownloadFileResponse
	ComposeFileRequest
	ComposeFileResponse
	GetFileStorageURLRequest
	GetFileStorageURLResponse
	GetStorageIPRequest
	GetStorageIPResponse
	GetFileDownloadURLRequest
	GetFileDownloadURLResponse
*/

package data

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Data service

type DataClient interface {
	FileCheck(ctx context.Context, in *FileCheckRequest, opts ...client.CallOption) (*FileCheckResponse, error)
	GetFileUploadURL(ctx context.Context, in *GetFileUploadURLRequest, opts ...client.CallOption) (*GetFileUploadURLResponse, error)
	GetFileSliceUploadURL(ctx context.Context, in *GetFileSliceUploadURLRequest, opts ...client.CallOption) (*GetFileSliceUploadURLResponse, error)
	GetUploadProgress(ctx context.Context, in *GetUploadProgressRequest, opts ...client.CallOption) (*GetUploadProgressResponse, error)
	GetFileStorageNode(ctx context.Context, in *GetFileStorageNodeRequest, opts ...client.CallOption) (*GetFileStorageNodeResponse, error)
	PutFile(ctx context.Context, in *PutFileRequest, opts ...client.CallOption) (*PutFileResponse, error)
	GetFileStorageURL(ctx context.Context, in *GetFileStorageURLRequest, opts ...client.CallOption) (*GetFileStorageURLResponse, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...client.CallOption) (*DownloadFileResponse, error)
	ComposeFile(ctx context.Context, in *ComposeFileRequest, opts ...client.CallOption) (*ComposeFileResponse, error)
	GetStorageIP(ctx context.Context, in *GetStorageIPRequest, opts ...client.CallOption) (*GetStorageIPResponse, error)
	GetFileDownloadURL(ctx context.Context, in *GetFileDownloadURLRequest, opts ...client.CallOption) (*GetFileDownloadURLResponse, error)
}

type dataClient struct {
	c           client.Client
	serviceName string
}

func NewDataClient(serviceName string, c client.Client) DataClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "data"
	}
	return &dataClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *dataClient) FileCheck(ctx context.Context, in *FileCheckRequest, opts ...client.CallOption) (*FileCheckResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.FileCheck", in)
	out := new(FileCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetFileUploadURL(ctx context.Context, in *GetFileUploadURLRequest, opts ...client.CallOption) (*GetFileUploadURLResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.GetFileUploadURL", in)
	out := new(GetFileUploadURLResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetFileSliceUploadURL(ctx context.Context, in *GetFileSliceUploadURLRequest, opts ...client.CallOption) (*GetFileSliceUploadURLResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.GetFileSliceUploadURL", in)
	out := new(GetFileSliceUploadURLResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetUploadProgress(ctx context.Context, in *GetUploadProgressRequest, opts ...client.CallOption) (*GetUploadProgressResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.GetUploadProgress", in)
	out := new(GetUploadProgressResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetFileStorageNode(ctx context.Context, in *GetFileStorageNodeRequest, opts ...client.CallOption) (*GetFileStorageNodeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.GetFileStorageNode", in)
	out := new(GetFileStorageNodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) PutFile(ctx context.Context, in *PutFileRequest, opts ...client.CallOption) (*PutFileResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.PutFile", in)
	out := new(PutFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetFileStorageURL(ctx context.Context, in *GetFileStorageURLRequest, opts ...client.CallOption) (*GetFileStorageURLResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.GetFileStorageURL", in)
	out := new(GetFileStorageURLResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...client.CallOption) (*DownloadFileResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.DownloadFile", in)
	out := new(DownloadFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) ComposeFile(ctx context.Context, in *ComposeFileRequest, opts ...client.CallOption) (*ComposeFileResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.ComposeFile", in)
	out := new(ComposeFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetStorageIP(ctx context.Context, in *GetStorageIPRequest, opts ...client.CallOption) (*GetStorageIPResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.GetStorageIP", in)
	out := new(GetStorageIPResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataClient) GetFileDownloadURL(ctx context.Context, in *GetFileDownloadURLRequest, opts ...client.CallOption) (*GetFileDownloadURLResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Data.GetFileDownloadURL", in)
	out := new(GetFileDownloadURLResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Data service

type DataHandler interface {
	FileCheck(context.Context, *FileCheckRequest, *FileCheckResponse) error
	GetFileUploadURL(context.Context, *GetFileUploadURLRequest, *GetFileUploadURLResponse) error
	GetFileSliceUploadURL(context.Context, *GetFileSliceUploadURLRequest, *GetFileSliceUploadURLResponse) error
	GetUploadProgress(context.Context, *GetUploadProgressRequest, *GetUploadProgressResponse) error
	GetFileStorageNode(context.Context, *GetFileStorageNodeRequest, *GetFileStorageNodeResponse) error
	PutFile(context.Context, *PutFileRequest, *PutFileResponse) error
	GetFileStorageURL(context.Context, *GetFileStorageURLRequest, *GetFileStorageURLResponse) error
	DownloadFile(context.Context, *DownloadFileRequest, *DownloadFileResponse) error
	ComposeFile(context.Context, *ComposeFileRequest, *ComposeFileResponse) error
	GetStorageIP(context.Context, *GetStorageIPRequest, *GetStorageIPResponse) error
	GetFileDownloadURL(context.Context, *GetFileDownloadURLRequest, *GetFileDownloadURLResponse) error
}

func RegisterDataHandler(s server.Server, hdlr DataHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Data{hdlr}, opts...))
}

type Data struct {
	DataHandler
}

func (h *Data) FileCheck(ctx context.Context, in *FileCheckRequest, out *FileCheckResponse) error {
	return h.DataHandler.FileCheck(ctx, in, out)
}

func (h *Data) GetFileUploadURL(ctx context.Context, in *GetFileUploadURLRequest, out *GetFileUploadURLResponse) error {
	return h.DataHandler.GetFileUploadURL(ctx, in, out)
}

func (h *Data) GetFileSliceUploadURL(ctx context.Context, in *GetFileSliceUploadURLRequest, out *GetFileSliceUploadURLResponse) error {
	return h.DataHandler.GetFileSliceUploadURL(ctx, in, out)
}

func (h *Data) GetUploadProgress(ctx context.Context, in *GetUploadProgressRequest, out *GetUploadProgressResponse) error {
	return h.DataHandler.GetUploadProgress(ctx, in, out)
}

func (h *Data) GetFileStorageNode(ctx context.Context, in *GetFileStorageNodeRequest, out *GetFileStorageNodeResponse) error {
	return h.DataHandler.GetFileStorageNode(ctx, in, out)
}

func (h *Data) PutFile(ctx context.Context, in *PutFileRequest, out *PutFileResponse) error {
	return h.DataHandler.PutFile(ctx, in, out)
}

func (h *Data) GetFileStorageURL(ctx context.Context, in *GetFileStorageURLRequest, out *GetFileStorageURLResponse) error {
	return h.DataHandler.GetFileStorageURL(ctx, in, out)
}

func (h *Data) DownloadFile(ctx context.Context, in *DownloadFileRequest, out *DownloadFileResponse) error {
	return h.DataHandler.DownloadFile(ctx, in, out)
}

func (h *Data) ComposeFile(ctx context.Context, in *ComposeFileRequest, out *ComposeFileResponse) error {
	return h.DataHandler.ComposeFile(ctx, in, out)
}

func (h *Data) GetStorageIP(ctx context.Context, in *GetStorageIPRequest, out *GetStorageIPResponse) error {
	return h.DataHandler.GetStorageIP(ctx, in, out)
}

func (h *Data) GetFileDownloadURL(ctx context.Context, in *GetFileDownloadURLRequest, out *GetFileDownloadURLResponse) error {
	return h.DataHandler.GetFileDownloadURL(ctx, in, out)
}
