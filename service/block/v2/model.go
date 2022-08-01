// Package block code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkblock

import (
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 生成枚举值

// 生成数据类型

type Entity struct {
	BlockId     *string `json:"block_id,omitempty"`
	Title       *string `json:"title,omitempty"`
	BlockTypeId *string `json:"block_type_id,omitempty"`
	SourceData  *string `json:"source_data,omitempty"`
	SourceMeta  *string `json:"source_meta,omitempty"`
	Version     *int64  `json:"version,omitempty,string"`
	SourceLink  *string `json:"source_link,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	Preview     *string `json:"preview,omitempty"`
	I18nSummay  *string `json:"i18n_summay,omitempty"`
	I18nPreview *string `json:"i18n_preview,omitempty"`
	Owner       *string `json:"owner,omitempty"`
	Extra       *string `json:"extra,omitempty"`
}

// builder开始
type EntityBuilder struct {
	blockId         string
	blockIdFlag     bool
	title           string
	titleFlag       bool
	blockTypeId     string
	blockTypeIdFlag bool
	sourceData      string
	sourceDataFlag  bool
	sourceMeta      string
	sourceMetaFlag  bool
	version         int64
	versionFlag     bool
	sourceLink      string
	sourceLinkFlag  bool
	summary         string
	summaryFlag     bool
	preview         string
	previewFlag     bool
	i18nSummay      string
	i18nSummayFlag  bool
	i18nPreview     string
	i18nPreviewFlag bool
	owner           string
	ownerFlag       bool
	extra           string
	extraFlag       bool
}

func NewEntityBuilder() *EntityBuilder {
	builder := &EntityBuilder{}
	return builder
}

func (builder *EntityBuilder) BlockId(blockId string) *EntityBuilder {
	builder.blockId = blockId
	builder.blockIdFlag = true
	return builder
}
func (builder *EntityBuilder) Title(title string) *EntityBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}
func (builder *EntityBuilder) BlockTypeId(blockTypeId string) *EntityBuilder {
	builder.blockTypeId = blockTypeId
	builder.blockTypeIdFlag = true
	return builder
}
func (builder *EntityBuilder) SourceData(sourceData string) *EntityBuilder {
	builder.sourceData = sourceData
	builder.sourceDataFlag = true
	return builder
}
func (builder *EntityBuilder) SourceMeta(sourceMeta string) *EntityBuilder {
	builder.sourceMeta = sourceMeta
	builder.sourceMetaFlag = true
	return builder
}
func (builder *EntityBuilder) Version(version int64) *EntityBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}
func (builder *EntityBuilder) SourceLink(sourceLink string) *EntityBuilder {
	builder.sourceLink = sourceLink
	builder.sourceLinkFlag = true
	return builder
}
func (builder *EntityBuilder) Summary(summary string) *EntityBuilder {
	builder.summary = summary
	builder.summaryFlag = true
	return builder
}
func (builder *EntityBuilder) Preview(preview string) *EntityBuilder {
	builder.preview = preview
	builder.previewFlag = true
	return builder
}
func (builder *EntityBuilder) I18nSummay(i18nSummay string) *EntityBuilder {
	builder.i18nSummay = i18nSummay
	builder.i18nSummayFlag = true
	return builder
}
func (builder *EntityBuilder) I18nPreview(i18nPreview string) *EntityBuilder {
	builder.i18nPreview = i18nPreview
	builder.i18nPreviewFlag = true
	return builder
}
func (builder *EntityBuilder) Owner(owner string) *EntityBuilder {
	builder.owner = owner
	builder.ownerFlag = true
	return builder
}
func (builder *EntityBuilder) Extra(extra string) *EntityBuilder {
	builder.extra = extra
	builder.extraFlag = true
	return builder
}

func (builder *EntityBuilder) Build() *Entity {
	req := &Entity{}
	if builder.blockIdFlag {
		req.BlockId = &builder.blockId

	}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	if builder.blockTypeIdFlag {
		req.BlockTypeId = &builder.blockTypeId

	}
	if builder.sourceDataFlag {
		req.SourceData = &builder.sourceData

	}
	if builder.sourceMetaFlag {
		req.SourceMeta = &builder.sourceMeta

	}
	if builder.versionFlag {
		req.Version = &builder.version

	}
	if builder.sourceLinkFlag {
		req.SourceLink = &builder.sourceLink

	}
	if builder.summaryFlag {
		req.Summary = &builder.summary

	}
	if builder.previewFlag {
		req.Preview = &builder.preview

	}
	if builder.i18nSummayFlag {
		req.I18nSummay = &builder.i18nSummay

	}
	if builder.i18nPreviewFlag {
		req.I18nPreview = &builder.i18nPreview

	}
	if builder.ownerFlag {
		req.Owner = &builder.owner

	}
	if builder.extraFlag {
		req.Extra = &builder.extra

	}
	return req
}

// builder结束

type Message struct {
	Body     *string  `json:"body,omitempty"`
	Version  *int64   `json:"version,omitempty,string"`
	BlockId  *string  `json:"block_id,omitempty"`
	Resource *string  `json:"resource,omitempty"`
	OpenIds  []string `json:"open_ids,omitempty"`
}

// builder开始
type MessageBuilder struct {
	body         string
	bodyFlag     bool
	version      int64
	versionFlag  bool
	blockId      string
	blockIdFlag  bool
	resource     string
	resourceFlag bool
	openIds      []string
	openIdsFlag  bool
}

func NewMessageBuilder() *MessageBuilder {
	builder := &MessageBuilder{}
	return builder
}

func (builder *MessageBuilder) Body(body string) *MessageBuilder {
	builder.body = body
	builder.bodyFlag = true
	return builder
}
func (builder *MessageBuilder) Version(version int64) *MessageBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}
func (builder *MessageBuilder) BlockId(blockId string) *MessageBuilder {
	builder.blockId = blockId
	builder.blockIdFlag = true
	return builder
}
func (builder *MessageBuilder) Resource(resource string) *MessageBuilder {
	builder.resource = resource
	builder.resourceFlag = true
	return builder
}
func (builder *MessageBuilder) OpenIds(openIds []string) *MessageBuilder {
	builder.openIds = openIds
	builder.openIdsFlag = true
	return builder
}

func (builder *MessageBuilder) Build() *Message {
	req := &Message{}
	if builder.bodyFlag {
		req.Body = &builder.body

	}
	if builder.versionFlag {
		req.Version = &builder.version

	}
	if builder.blockIdFlag {
		req.BlockId = &builder.blockId

	}
	if builder.resourceFlag {
		req.Resource = &builder.resource

	}
	if builder.openIdsFlag {
		req.OpenIds = builder.openIds
	}
	return req
}

// builder结束

// 生成请求和响应结果类型，以及请求对象的Builder构造器

type CreateEntityReqBodyBuilder struct {
	title           string
	titleFlag       bool
	blockTypeId     string
	blockTypeIdFlag bool
	sourceData      string
	sourceDataFlag  bool
	sourceMeta      string
	sourceMetaFlag  bool
	version         int64
	versionFlag     bool
	sourceLink      string
	sourceLinkFlag  bool
	owner           string
	ownerFlag       bool
	extra           string
	extraFlag       bool
	i18nSummary     string
	i18nSummaryFlag bool
	i18nPreview     string
	i18nPreviewFlag bool
	summary         string
	summaryFlag     bool
	preview         string
	previewFlag     bool
}

// 生成body的New构造器
func NewCreateEntityReqBodyBuilder() *CreateEntityReqBodyBuilder {
	builder := &CreateEntityReqBodyBuilder{}
	return builder
}

// 1.2 生成body的builder属性方法
func (builder *CreateEntityReqBodyBuilder) Title(title string) *CreateEntityReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) BlockTypeId(blockTypeId string) *CreateEntityReqBodyBuilder {
	builder.blockTypeId = blockTypeId
	builder.blockTypeIdFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) SourceData(sourceData string) *CreateEntityReqBodyBuilder {
	builder.sourceData = sourceData
	builder.sourceDataFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) SourceMeta(sourceMeta string) *CreateEntityReqBodyBuilder {
	builder.sourceMeta = sourceMeta
	builder.sourceMetaFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) Version(version int64) *CreateEntityReqBodyBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) SourceLink(sourceLink string) *CreateEntityReqBodyBuilder {
	builder.sourceLink = sourceLink
	builder.sourceLinkFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) Owner(owner string) *CreateEntityReqBodyBuilder {
	builder.owner = owner
	builder.ownerFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) Extra(extra string) *CreateEntityReqBodyBuilder {
	builder.extra = extra
	builder.extraFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) I18nSummary(i18nSummary string) *CreateEntityReqBodyBuilder {
	builder.i18nSummary = i18nSummary
	builder.i18nSummaryFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) I18nPreview(i18nPreview string) *CreateEntityReqBodyBuilder {
	builder.i18nPreview = i18nPreview
	builder.i18nPreviewFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) Summary(summary string) *CreateEntityReqBodyBuilder {
	builder.summary = summary
	builder.summaryFlag = true
	return builder
}
func (builder *CreateEntityReqBodyBuilder) Preview(preview string) *CreateEntityReqBodyBuilder {
	builder.preview = preview
	builder.previewFlag = true
	return builder
}

// 1.3 生成body的build方法
func (builder *CreateEntityReqBodyBuilder) Build() *CreateEntityReqBody {
	req := &CreateEntityReqBody{}
	if builder.titleFlag {
		req.Title = &builder.title
	}
	if builder.blockTypeIdFlag {
		req.BlockTypeId = &builder.blockTypeId
	}
	if builder.sourceDataFlag {
		req.SourceData = &builder.sourceData
	}
	if builder.sourceMetaFlag {
		req.SourceMeta = &builder.sourceMeta
	}
	if builder.versionFlag {
		req.Version = &builder.version
	}
	if builder.sourceLinkFlag {
		req.SourceLink = &builder.sourceLink
	}
	if builder.ownerFlag {
		req.Owner = &builder.owner
	}
	if builder.extraFlag {
		req.Extra = &builder.extra
	}
	if builder.i18nSummaryFlag {
		req.I18nSummary = &builder.i18nSummary
	}
	if builder.i18nPreviewFlag {
		req.I18nPreview = &builder.i18nPreview
	}
	if builder.summaryFlag {
		req.Summary = &builder.summary
	}
	if builder.previewFlag {
		req.Preview = &builder.preview
	}
	return req
}

// 上传文件path开始
type CreateEntityPathReqBodyBuilder struct {
	title           string
	titleFlag       bool
	blockTypeId     string
	blockTypeIdFlag bool
	sourceData      string
	sourceDataFlag  bool
	sourceMeta      string
	sourceMetaFlag  bool
	version         int64
	versionFlag     bool
	sourceLink      string
	sourceLinkFlag  bool
	owner           string
	ownerFlag       bool
	extra           string
	extraFlag       bool
	i18nSummary     string
	i18nSummaryFlag bool
	i18nPreview     string
	i18nPreviewFlag bool
	summary         string
	summaryFlag     bool
	preview         string
	previewFlag     bool
}

func NewCreateEntityPathReqBodyBuilder() *CreateEntityPathReqBodyBuilder {
	builder := &CreateEntityPathReqBodyBuilder{}
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) Title(title string) *CreateEntityPathReqBodyBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) BlockTypeId(blockTypeId string) *CreateEntityPathReqBodyBuilder {
	builder.blockTypeId = blockTypeId
	builder.blockTypeIdFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) SourceData(sourceData string) *CreateEntityPathReqBodyBuilder {
	builder.sourceData = sourceData
	builder.sourceDataFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) SourceMeta(sourceMeta string) *CreateEntityPathReqBodyBuilder {
	builder.sourceMeta = sourceMeta
	builder.sourceMetaFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) Version(version int64) *CreateEntityPathReqBodyBuilder {
	builder.version = version
	builder.versionFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) SourceLink(sourceLink string) *CreateEntityPathReqBodyBuilder {
	builder.sourceLink = sourceLink
	builder.sourceLinkFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) Owner(owner string) *CreateEntityPathReqBodyBuilder {
	builder.owner = owner
	builder.ownerFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) Extra(extra string) *CreateEntityPathReqBodyBuilder {
	builder.extra = extra
	builder.extraFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) I18nSummary(i18nSummary string) *CreateEntityPathReqBodyBuilder {
	builder.i18nSummary = i18nSummary
	builder.i18nSummaryFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) I18nPreview(i18nPreview string) *CreateEntityPathReqBodyBuilder {
	builder.i18nPreview = i18nPreview
	builder.i18nPreviewFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) Summary(summary string) *CreateEntityPathReqBodyBuilder {
	builder.summary = summary
	builder.summaryFlag = true
	return builder
}
func (builder *CreateEntityPathReqBodyBuilder) Preview(preview string) *CreateEntityPathReqBodyBuilder {
	builder.preview = preview
	builder.previewFlag = true
	return builder
}

func (builder *CreateEntityPathReqBodyBuilder) Build() (*CreateEntityReqBody, error) {
	req := &CreateEntityReqBody{}
	if builder.titleFlag {
		req.Title = &builder.title
	}
	if builder.blockTypeIdFlag {
		req.BlockTypeId = &builder.blockTypeId
	}
	if builder.sourceDataFlag {
		req.SourceData = &builder.sourceData
	}
	if builder.sourceMetaFlag {
		req.SourceMeta = &builder.sourceMeta
	}
	if builder.versionFlag {
		req.Version = &builder.version
	}
	if builder.sourceLinkFlag {
		req.SourceLink = &builder.sourceLink
	}
	if builder.ownerFlag {
		req.Owner = &builder.owner
	}
	if builder.extraFlag {
		req.Extra = &builder.extra
	}
	if builder.i18nSummaryFlag {
		req.I18nSummary = &builder.i18nSummary
	}
	if builder.i18nPreviewFlag {
		req.I18nPreview = &builder.i18nPreview
	}
	if builder.summaryFlag {
		req.Summary = &builder.summary
	}
	if builder.previewFlag {
		req.Preview = &builder.preview
	}
	return req, nil
}

// 上传文件path结束

// 1.4 生成请求的builder结构体
type CreateEntityReqBuilder struct {
	apiReq *larkcore.ApiReq
	body   *CreateEntityReqBody
}

// 生成请求的New构造器
func NewCreateEntityReqBuilder() *CreateEntityReqBuilder {
	builder := &CreateEntityReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *CreateEntityReqBuilder) Body(body *CreateEntityReqBody) *CreateEntityReqBuilder {
	builder.body = body
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *CreateEntityReqBuilder) Build() *CreateEntityReq {
	req := &CreateEntityReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.body
	return req
}

type CreateEntityReqBody struct {
	Title       *string `json:"title,omitempty"`
	BlockTypeId *string `json:"block_type_id,omitempty"`
	SourceData  *string `json:"source_data,omitempty"`
	SourceMeta  *string `json:"source_meta,omitempty"`
	Version     *int64  `json:"version,omitempty,string"`
	SourceLink  *string `json:"source_link,omitempty"`
	Owner       *string `json:"owner,omitempty"`
	Extra       *string `json:"extra,omitempty"`
	I18nSummary *string `json:"i18n_summary,omitempty"`
	I18nPreview *string `json:"i18n_preview,omitempty"`
	Summary     *string `json:"summary,omitempty"`
	Preview     *string `json:"preview,omitempty"`
}

type CreateEntityReq struct {
	apiReq *larkcore.ApiReq
	Body   *CreateEntityReqBody `body:""`
}

type CreateEntityRespData struct {
	Entity *Entity `json:"entity,omitempty"`
}

type CreateEntityResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *CreateEntityRespData `json:"data"`
}

func (resp *CreateEntityResp) Success() bool {
	return resp.Code == 0
}

// 1.4 生成请求的builder结构体
type UpdateEntityReqBuilder struct {
	apiReq *larkcore.ApiReq
	entity *Entity
}

// 生成请求的New构造器
func NewUpdateEntityReqBuilder() *UpdateEntityReqBuilder {
	builder := &UpdateEntityReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *UpdateEntityReqBuilder) BlockId(blockId string) *UpdateEntityReqBuilder {
	builder.apiReq.PathParams.Set("block_id", fmt.Sprint(blockId))
	return builder
}
func (builder *UpdateEntityReqBuilder) Entity(entity *Entity) *UpdateEntityReqBuilder {
	builder.entity = entity
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *UpdateEntityReqBuilder) Build() *UpdateEntityReq {
	req := &UpdateEntityReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	req.apiReq.Body = builder.entity
	return req
}

type UpdateEntityReq struct {
	apiReq *larkcore.ApiReq
	Entity *Entity `body:""`
}

type UpdateEntityResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
}

func (resp *UpdateEntityResp) Success() bool {
	return resp.Code == 0
}

// 1.4 生成请求的builder结构体
type CreateMessageReqBuilder struct {
	apiReq  *larkcore.ApiReq
	message *Message
}

// 生成请求的New构造器
func NewCreateMessageReqBuilder() *CreateMessageReqBuilder {
	builder := &CreateMessageReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 1.5 生成请求的builder属性方法
func (builder *CreateMessageReqBuilder) Message(message *Message) *CreateMessageReqBuilder {
	builder.message = message
	return builder
}

// 1.5 生成请求的builder的build方法
func (builder *CreateMessageReqBuilder) Build() *CreateMessageReq {
	req := &CreateMessageReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.Body = builder.message
	return req
}

type CreateMessageReq struct {
	apiReq  *larkcore.ApiReq
	Message *Message `body:""`
}

type CreateMessageResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
}

func (resp *CreateMessageResp) Success() bool {
	return resp.Code == 0
}

// 生成消息事件结构体

// 生成请求的builder构造器
// 1.1 生成body的builder结构体