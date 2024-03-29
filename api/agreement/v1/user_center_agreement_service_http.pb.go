// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.4
// source: user_center_agreement_service.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationServiceAddContent = "/agreement.Service/AddContent"
const OperationServiceAddScene = "/agreement.Service/AddScene"
const OperationServiceDeleteContent = "/agreement.Service/DeleteContent"
const OperationServiceDeleteScene = "/agreement.Service/DeleteScene"
const OperationServiceGetContent = "/agreement.Service/GetContent"
const OperationServiceGetSceneByKeyword = "/agreement.Service/GetSceneByKeyword"
const OperationServicePageContent = "/agreement.Service/PageContent"
const OperationServicePageScene = "/agreement.Service/PageScene"
const OperationServiceUpdateContent = "/agreement.Service/UpdateContent"
const OperationServiceUpdateScene = "/agreement.Service/UpdateScene"

type ServiceHTTPServer interface {
	// AddContent 新增协议内容
	AddContent(context.Context, *AddContentRequest) (*AddContentReply, error)
	// AddScene 新增协议场景
	AddScene(context.Context, *AddSceneRequest) (*AddSceneReply, error)
	// DeleteContent 删除指定的协议内容
	DeleteContent(context.Context, *DeleteContentRequest) (*emptypb.Empty, error)
	// DeleteScene 删除协议场景
	DeleteScene(context.Context, *DeleteSceneRequest) (*emptypb.Empty, error)
	// GetContent 获取协议的详细内容
	GetContent(context.Context, *GetContentRequest) (*Content, error)
	// GetSceneByKeyword 获取指定的协议场景以及协议场景包含的协议列表
	GetSceneByKeyword(context.Context, *GetSceneByKeywordRequest) (*Scene, error)
	// PageContent 获取分页协议
	PageContent(context.Context, *PageContentRequest) (*PageContentReply, error)
	// PageScene 获取分页协议场景
	PageScene(context.Context, *PageSceneRequest) (*PageSceneReply, error)
	// UpdateContent 更新协议内容
	UpdateContent(context.Context, *UpdateContentRequest) (*emptypb.Empty, error)
	// UpdateScene 更新协议场景
	UpdateScene(context.Context, *UpdateSceneRequest) (*emptypb.Empty, error)
}

func RegisterServiceHTTPServer(s *http.Server, srv ServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/user-center/v1/agreement/contents", _Service_PageContent0_HTTP_Handler(srv))
	r.GET("/user-center/client/v1/agreement/content", _Service_GetContent0_HTTP_Handler(srv))
	r.GET("/user-center/v1/agreement/content", _Service_GetContent1_HTTP_Handler(srv))
	r.POST("/user-center/v1/agreement/content", _Service_AddContent0_HTTP_Handler(srv))
	r.PUT("/user-center/v1/agreement/content", _Service_UpdateContent0_HTTP_Handler(srv))
	r.DELETE("/user-center/v1/agreement/content", _Service_DeleteContent0_HTTP_Handler(srv))
	r.GET("/user-center/v1/agreement/scenes", _Service_PageScene0_HTTP_Handler(srv))
	r.GET("/user-center/client/v1/agreement/scene", _Service_GetSceneByKeyword0_HTTP_Handler(srv))
	r.POST("/user-center/v1/agreement/scene", _Service_AddScene0_HTTP_Handler(srv))
	r.PUT("/user-center/v1/agreement/scene", _Service_UpdateScene0_HTTP_Handler(srv))
	r.DELETE("/user-center/v1/agreement/scene", _Service_DeleteScene0_HTTP_Handler(srv))
}

func _Service_PageContent0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageContentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServicePageContent)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.PageContent(ctx, req.(*PageContentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageContentReply)
		return ctx.Result(200, reply)
	}
}

func _Service_GetContent0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetContentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetContent)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetContent(ctx, req.(*GetContentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Content)
		return ctx.Result(200, reply)
	}
}

func _Service_GetContent1_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetContentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetContent)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetContent(ctx, req.(*GetContentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Content)
		return ctx.Result(200, reply)
	}
}

func _Service_AddContent0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddContentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAddContent)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.AddContent(ctx, req.(*AddContentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddContentReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateContent0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateContentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateContent)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateContent(ctx, req.(*UpdateContentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Service_DeleteContent0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteContentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceDeleteContent)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteContent(ctx, req.(*DeleteContentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Service_PageScene0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PageSceneRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServicePageScene)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.PageScene(ctx, req.(*PageSceneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PageSceneReply)
		return ctx.Result(200, reply)
	}
}

func _Service_GetSceneByKeyword0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSceneByKeywordRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceGetSceneByKeyword)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.GetSceneByKeyword(ctx, req.(*GetSceneByKeywordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Scene)
		return ctx.Result(200, reply)
	}
}

func _Service_AddScene0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddSceneRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceAddScene)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.AddScene(ctx, req.(*AddSceneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddSceneReply)
		return ctx.Result(200, reply)
	}
}

func _Service_UpdateScene0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSceneRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceUpdateScene)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.UpdateScene(ctx, req.(*UpdateSceneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Service_DeleteScene0_HTTP_Handler(srv ServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSceneRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationServiceDeleteScene)
		h := ctx.Middleware(func(ctx context.Context, req any) (any, error) {
			return srv.DeleteScene(ctx, req.(*DeleteSceneRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type ServiceHTTPClient interface {
	AddContent(ctx context.Context, req *AddContentRequest, opts ...http.CallOption) (rsp *AddContentReply, err error)
	AddScene(ctx context.Context, req *AddSceneRequest, opts ...http.CallOption) (rsp *AddSceneReply, err error)
	DeleteContent(ctx context.Context, req *DeleteContentRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	DeleteScene(ctx context.Context, req *DeleteSceneRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetContent(ctx context.Context, req *GetContentRequest, opts ...http.CallOption) (rsp *Content, err error)
	GetSceneByKeyword(ctx context.Context, req *GetSceneByKeywordRequest, opts ...http.CallOption) (rsp *Scene, err error)
	PageContent(ctx context.Context, req *PageContentRequest, opts ...http.CallOption) (rsp *PageContentReply, err error)
	PageScene(ctx context.Context, req *PageSceneRequest, opts ...http.CallOption) (rsp *PageSceneReply, err error)
	UpdateContent(ctx context.Context, req *UpdateContentRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	UpdateScene(ctx context.Context, req *UpdateSceneRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type ServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewServiceHTTPClient(client *http.Client) ServiceHTTPClient {
	return &ServiceHTTPClientImpl{client}
}

func (c *ServiceHTTPClientImpl) AddContent(ctx context.Context, in *AddContentRequest, opts ...http.CallOption) (*AddContentReply, error) {
	var out AddContentReply
	pattern := "/user-center/v1/agreement/content"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceAddContent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) AddScene(ctx context.Context, in *AddSceneRequest, opts ...http.CallOption) (*AddSceneReply, error) {
	var out AddSceneReply
	pattern := "/user-center/v1/agreement/scene"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceAddScene))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) DeleteContent(ctx context.Context, in *DeleteContentRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user-center/v1/agreement/content"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceDeleteContent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) DeleteScene(ctx context.Context, in *DeleteSceneRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user-center/v1/agreement/scene"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceDeleteScene))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GetContent(ctx context.Context, in *GetContentRequest, opts ...http.CallOption) (*Content, error) {
	var out Content
	pattern := "/user-center/v1/agreement/content"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetContent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) GetSceneByKeyword(ctx context.Context, in *GetSceneByKeywordRequest, opts ...http.CallOption) (*Scene, error) {
	var out Scene
	pattern := "/user-center/client/v1/agreement/scene"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServiceGetSceneByKeyword))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) PageContent(ctx context.Context, in *PageContentRequest, opts ...http.CallOption) (*PageContentReply, error) {
	var out PageContentReply
	pattern := "/user-center/v1/agreement/contents"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServicePageContent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) PageScene(ctx context.Context, in *PageSceneRequest, opts ...http.CallOption) (*PageSceneReply, error) {
	var out PageSceneReply
	pattern := "/user-center/v1/agreement/scenes"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationServicePageScene))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateContent(ctx context.Context, in *UpdateContentRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user-center/v1/agreement/content"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateContent))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ServiceHTTPClientImpl) UpdateScene(ctx context.Context, in *UpdateSceneRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/user-center/v1/agreement/scene"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationServiceUpdateScene))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
