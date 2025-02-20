// Code generated by Lark OpenAPI.

package larkmoments

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"net/http"
)

type V1 struct {
	Comment        *comment        // comment
	Post           *post           // post
	PostStatistics *postStatistics // post_statistics
	Reaction       *reaction       // reaction
}

func New(config *larkcore.Config) *V1 {
	return &V1{
		Comment:        &comment{config: config},
		Post:           &post{config: config},
		PostStatistics: &postStatistics{config: config},
		Reaction:       &reaction{config: config},
	}
}

type comment struct {
	config *larkcore.Config
}
type post struct {
	config *larkcore.Config
}
type postStatistics struct {
	config *larkcore.Config
}
type reaction struct {
	config *larkcore.Config
}

// Get
//
// -
//
// - 官网API文档链接:https://open.feishu.cn/api-explorer?from=op_doc_tab&apiName=get&project=moments&resource=post&version=v1
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/momentsv1/get_post.go
func (p *post) Get(ctx context.Context, req *GetPostReq, options ...larkcore.RequestOptionFunc) (*GetPostResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/moments/v1/posts/:post_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, p.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetPostResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp, p.config)
	if err != nil {
		return nil, err
	}
	return resp, err
}
