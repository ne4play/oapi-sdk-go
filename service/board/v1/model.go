// Package board code generated by oapi sdk gen
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

package larkboard

import (
	"io"

	"io/ioutil"

	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

type CompositeShape struct {
	Type *string `json:"type,omitempty"` // 基础图形的具体类型
}

type CompositeShapeBuilder struct {
	type_    string // 基础图形的具体类型
	typeFlag bool
}

func NewCompositeShapeBuilder() *CompositeShapeBuilder {
	builder := &CompositeShapeBuilder{}
	return builder
}

// 基础图形的具体类型
//
// 示例值：
func (builder *CompositeShapeBuilder) Type(type_ string) *CompositeShapeBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

func (builder *CompositeShapeBuilder) Build() *CompositeShape {
	req := &CompositeShape{}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	return req
}

type Connector struct {
	StartObject *ConnectorAttachedObject `json:"start_object,omitempty"` // 连线连接的起点图形
	EndObject   *ConnectorAttachedObject `json:"end_object,omitempty"`   // 连线连接的终点图形
	Captions    *ConnectorCaption        `json:"captions,omitempty"`     // 连线文本
}

type ConnectorBuilder struct {
	startObject     *ConnectorAttachedObject // 连线连接的起点图形
	startObjectFlag bool
	endObject       *ConnectorAttachedObject // 连线连接的终点图形
	endObjectFlag   bool
	captions        *ConnectorCaption // 连线文本
	captionsFlag    bool
}

func NewConnectorBuilder() *ConnectorBuilder {
	builder := &ConnectorBuilder{}
	return builder
}

// 连线连接的起点图形
//
// 示例值：
func (builder *ConnectorBuilder) StartObject(startObject *ConnectorAttachedObject) *ConnectorBuilder {
	builder.startObject = startObject
	builder.startObjectFlag = true
	return builder
}

// 连线连接的终点图形
//
// 示例值：
func (builder *ConnectorBuilder) EndObject(endObject *ConnectorAttachedObject) *ConnectorBuilder {
	builder.endObject = endObject
	builder.endObjectFlag = true
	return builder
}

// 连线文本
//
// 示例值：
func (builder *ConnectorBuilder) Captions(captions *ConnectorCaption) *ConnectorBuilder {
	builder.captions = captions
	builder.captionsFlag = true
	return builder
}

func (builder *ConnectorBuilder) Build() *Connector {
	req := &Connector{}
	if builder.startObjectFlag {
		req.StartObject = builder.startObject
	}
	if builder.endObjectFlag {
		req.EndObject = builder.endObject
	}
	if builder.captionsFlag {
		req.Captions = builder.captions
	}
	return req
}

type ConnectorAttachedObject struct {
	Id *string `json:"id,omitempty"` // 连接图形的 id
}

type ConnectorAttachedObjectBuilder struct {
	id     string // 连接图形的 id
	idFlag bool
}

func NewConnectorAttachedObjectBuilder() *ConnectorAttachedObjectBuilder {
	builder := &ConnectorAttachedObjectBuilder{}
	return builder
}

// 连接图形的 id
//
// 示例值：o1:1
func (builder *ConnectorAttachedObjectBuilder) Id(id string) *ConnectorAttachedObjectBuilder {
	builder.id = id
	builder.idFlag = true
	return builder
}

func (builder *ConnectorAttachedObjectBuilder) Build() *ConnectorAttachedObject {
	req := &ConnectorAttachedObject{}
	if builder.idFlag {
		req.Id = &builder.id

	}
	return req
}

type ConnectorCaption struct {
	Data []*Text `json:"data,omitempty"` // 文本
}

type ConnectorCaptionBuilder struct {
	data     []*Text // 文本
	dataFlag bool
}

func NewConnectorCaptionBuilder() *ConnectorCaptionBuilder {
	builder := &ConnectorCaptionBuilder{}
	return builder
}

// 文本
//
// 示例值：
func (builder *ConnectorCaptionBuilder) Data(data []*Text) *ConnectorCaptionBuilder {
	builder.data = data
	builder.dataFlag = true
	return builder
}

func (builder *ConnectorCaptionBuilder) Build() *ConnectorCaption {
	req := &ConnectorCaption{}
	if builder.dataFlag {
		req.Data = builder.data
	}
	return req
}

type DepartmentId struct {
	DepartmentId     *string `json:"department_id,omitempty"`      //
	OpenDepartmentId *string `json:"open_department_id,omitempty"` //
}

type DepartmentIdBuilder struct {
	departmentId         string //
	departmentIdFlag     bool
	openDepartmentId     string //
	openDepartmentIdFlag bool
}

func NewDepartmentIdBuilder() *DepartmentIdBuilder {
	builder := &DepartmentIdBuilder{}
	return builder
}

// 示例值：
func (builder *DepartmentIdBuilder) DepartmentId(departmentId string) *DepartmentIdBuilder {
	builder.departmentId = departmentId
	builder.departmentIdFlag = true
	return builder
}

// 示例值：
func (builder *DepartmentIdBuilder) OpenDepartmentId(openDepartmentId string) *DepartmentIdBuilder {
	builder.openDepartmentId = openDepartmentId
	builder.openDepartmentIdFlag = true
	return builder
}

func (builder *DepartmentIdBuilder) Build() *DepartmentId {
	req := &DepartmentId{}
	if builder.departmentIdFlag {
		req.DepartmentId = &builder.departmentId

	}
	if builder.openDepartmentIdFlag {
		req.OpenDepartmentId = &builder.openDepartmentId

	}
	return req
}

type Image struct {
	Token *string `json:"token,omitempty"` // 图片 token
}

type ImageBuilder struct {
	token     string // 图片 token
	tokenFlag bool
}

func NewImageBuilder() *ImageBuilder {
	builder := &ImageBuilder{}
	return builder
}

// 图片 token
//
// 示例值：EeSHb3qs9oSBXoxvw33bqtOsczb
func (builder *ImageBuilder) Token(token string) *ImageBuilder {
	builder.token = token
	builder.tokenFlag = true
	return builder
}

func (builder *ImageBuilder) Build() *Image {
	req := &Image{}
	if builder.tokenFlag {
		req.Token = &builder.token

	}
	return req
}

type MindMap struct {
	ParentId *string `json:"parent_id,omitempty"` // 思维导图父节点 id ，为空表示是思维导图的根节点
}

type MindMapBuilder struct {
	parentId     string // 思维导图父节点 id ，为空表示是思维导图的根节点
	parentIdFlag bool
}

func NewMindMapBuilder() *MindMapBuilder {
	builder := &MindMapBuilder{}
	return builder
}

// 思维导图父节点 id ，为空表示是思维导图的根节点
//
// 示例值：z1:1
func (builder *MindMapBuilder) ParentId(parentId string) *MindMapBuilder {
	builder.parentId = parentId
	builder.parentIdFlag = true
	return builder
}

func (builder *MindMapBuilder) Build() *MindMap {
	req := &MindMap{}
	if builder.parentIdFlag {
		req.ParentId = &builder.parentId

	}
	return req
}

type Section struct {
	Title *string `json:"title,omitempty"` // 分区标题
}

type SectionBuilder struct {
	title     string // 分区标题
	titleFlag bool
}

func NewSectionBuilder() *SectionBuilder {
	builder := &SectionBuilder{}
	return builder
}

// 分区标题
//
// 示例值：分区
func (builder *SectionBuilder) Title(title string) *SectionBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

func (builder *SectionBuilder) Build() *Section {
	req := &Section{}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	return req
}

type Style struct {
	FillOpacity   *float64 `json:"fill_opacity,omitempty"`   // 填充透明度
	BorderStyle   *string  `json:"border_style,omitempty"`   // 边框样式
	BorderWidth   *string  `json:"border_width,omitempty"`   // 边框宽度
	BorderOpacity *float64 `json:"border_opacity,omitempty"` // 边框透明度
	HFlip         *bool    `json:"h_flip,omitempty"`         // 水平翻折
	VFlip         *bool    `json:"v_flip,omitempty"`         // 垂直翻折
}

type StyleBuilder struct {
	fillOpacity       float64 // 填充透明度
	fillOpacityFlag   bool
	borderStyle       string // 边框样式
	borderStyleFlag   bool
	borderWidth       string // 边框宽度
	borderWidthFlag   bool
	borderOpacity     float64 // 边框透明度
	borderOpacityFlag bool
	hFlip             bool // 水平翻折
	hFlipFlag         bool
	vFlip             bool // 垂直翻折
	vFlipFlag         bool
}

func NewStyleBuilder() *StyleBuilder {
	builder := &StyleBuilder{}
	return builder
}

// 填充透明度
//
// 示例值：50
func (builder *StyleBuilder) FillOpacity(fillOpacity float64) *StyleBuilder {
	builder.fillOpacity = fillOpacity
	builder.fillOpacityFlag = true
	return builder
}

// 边框样式
//
// 示例值：
func (builder *StyleBuilder) BorderStyle(borderStyle string) *StyleBuilder {
	builder.borderStyle = borderStyle
	builder.borderStyleFlag = true
	return builder
}

// 边框宽度
//
// 示例值：
func (builder *StyleBuilder) BorderWidth(borderWidth string) *StyleBuilder {
	builder.borderWidth = borderWidth
	builder.borderWidthFlag = true
	return builder
}

// 边框透明度
//
// 示例值：50
func (builder *StyleBuilder) BorderOpacity(borderOpacity float64) *StyleBuilder {
	builder.borderOpacity = borderOpacity
	builder.borderOpacityFlag = true
	return builder
}

// 水平翻折
//
// 示例值：false
func (builder *StyleBuilder) HFlip(hFlip bool) *StyleBuilder {
	builder.hFlip = hFlip
	builder.hFlipFlag = true
	return builder
}

// 垂直翻折
//
// 示例值：false
func (builder *StyleBuilder) VFlip(vFlip bool) *StyleBuilder {
	builder.vFlip = vFlip
	builder.vFlipFlag = true
	return builder
}

func (builder *StyleBuilder) Build() *Style {
	req := &Style{}
	if builder.fillOpacityFlag {
		req.FillOpacity = &builder.fillOpacity

	}
	if builder.borderStyleFlag {
		req.BorderStyle = &builder.borderStyle

	}
	if builder.borderWidthFlag {
		req.BorderWidth = &builder.borderWidth

	}
	if builder.borderOpacityFlag {
		req.BorderOpacity = &builder.borderOpacity

	}
	if builder.hFlipFlag {
		req.HFlip = &builder.hFlip

	}
	if builder.vFlipFlag {
		req.VFlip = &builder.vFlip

	}
	return req
}

type Table struct {
	Meta  *TableMeta   `json:"meta,omitempty"`  // 元信息
	Title *string      `json:"title,omitempty"` // 标题
	Cells []*TableCell `json:"cells,omitempty"` // 单元格列表
}

type TableBuilder struct {
	meta      *TableMeta // 元信息
	metaFlag  bool
	title     string // 标题
	titleFlag bool
	cells     []*TableCell // 单元格列表
	cellsFlag bool
}

func NewTableBuilder() *TableBuilder {
	builder := &TableBuilder{}
	return builder
}

// 元信息
//
// 示例值：
func (builder *TableBuilder) Meta(meta *TableMeta) *TableBuilder {
	builder.meta = meta
	builder.metaFlag = true
	return builder
}

// 标题
//
// 示例值：表格
func (builder *TableBuilder) Title(title string) *TableBuilder {
	builder.title = title
	builder.titleFlag = true
	return builder
}

// 单元格列表
//
// 示例值：
func (builder *TableBuilder) Cells(cells []*TableCell) *TableBuilder {
	builder.cells = cells
	builder.cellsFlag = true
	return builder
}

func (builder *TableBuilder) Build() *Table {
	req := &Table{}
	if builder.metaFlag {
		req.Meta = builder.meta
	}
	if builder.titleFlag {
		req.Title = &builder.title

	}
	if builder.cellsFlag {
		req.Cells = builder.cells
	}
	return req
}

type TableCell struct {
	RowIndex  *int                `json:"row_index,omitempty"`  // 行下标，从 1 开始
	ColIndex  *int                `json:"col_index,omitempty"`  // 列下标，从 1 开始
	MergeInfo *TableCellMergeInfo `json:"merge_info,omitempty"` // 单元格合并信息
	Children  []string            `json:"children,omitempty"`   // 单元格包含的子节点 id
	Text      *Text               `json:"text,omitempty"`       // 单元格内文字
}

type TableCellBuilder struct {
	rowIndex      int // 行下标，从 1 开始
	rowIndexFlag  bool
	colIndex      int // 列下标，从 1 开始
	colIndexFlag  bool
	mergeInfo     *TableCellMergeInfo // 单元格合并信息
	mergeInfoFlag bool
	children      []string // 单元格包含的子节点 id
	childrenFlag  bool
	text          *Text // 单元格内文字
	textFlag      bool
}

func NewTableCellBuilder() *TableCellBuilder {
	builder := &TableCellBuilder{}
	return builder
}

// 行下标，从 1 开始
//
// 示例值：1
func (builder *TableCellBuilder) RowIndex(rowIndex int) *TableCellBuilder {
	builder.rowIndex = rowIndex
	builder.rowIndexFlag = true
	return builder
}

// 列下标，从 1 开始
//
// 示例值：1
func (builder *TableCellBuilder) ColIndex(colIndex int) *TableCellBuilder {
	builder.colIndex = colIndex
	builder.colIndexFlag = true
	return builder
}

// 单元格合并信息
//
// 示例值：
func (builder *TableCellBuilder) MergeInfo(mergeInfo *TableCellMergeInfo) *TableCellBuilder {
	builder.mergeInfo = mergeInfo
	builder.mergeInfoFlag = true
	return builder
}

// 单元格包含的子节点 id
//
// 示例值：
func (builder *TableCellBuilder) Children(children []string) *TableCellBuilder {
	builder.children = children
	builder.childrenFlag = true
	return builder
}

// 单元格内文字
//
// 示例值：
func (builder *TableCellBuilder) Text(text *Text) *TableCellBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

func (builder *TableCellBuilder) Build() *TableCell {
	req := &TableCell{}
	if builder.rowIndexFlag {
		req.RowIndex = &builder.rowIndex

	}
	if builder.colIndexFlag {
		req.ColIndex = &builder.colIndex

	}
	if builder.mergeInfoFlag {
		req.MergeInfo = builder.mergeInfo
	}
	if builder.childrenFlag {
		req.Children = builder.children
	}
	if builder.textFlag {
		req.Text = builder.text
	}
	return req
}

type TableCellMergeInfo struct {
	RowSpan *int `json:"row_span,omitempty"` // 从当前行索引起被合并的连续行数
	ColSpan *int `json:"col_span,omitempty"` // 从当前列索引起被合并的连续列数
}

type TableCellMergeInfoBuilder struct {
	rowSpan     int // 从当前行索引起被合并的连续行数
	rowSpanFlag bool
	colSpan     int // 从当前列索引起被合并的连续列数
	colSpanFlag bool
}

func NewTableCellMergeInfoBuilder() *TableCellMergeInfoBuilder {
	builder := &TableCellMergeInfoBuilder{}
	return builder
}

// 从当前行索引起被合并的连续行数
//
// 示例值：2
func (builder *TableCellMergeInfoBuilder) RowSpan(rowSpan int) *TableCellMergeInfoBuilder {
	builder.rowSpan = rowSpan
	builder.rowSpanFlag = true
	return builder
}

// 从当前列索引起被合并的连续列数
//
// 示例值：2
func (builder *TableCellMergeInfoBuilder) ColSpan(colSpan int) *TableCellMergeInfoBuilder {
	builder.colSpan = colSpan
	builder.colSpanFlag = true
	return builder
}

func (builder *TableCellMergeInfoBuilder) Build() *TableCellMergeInfo {
	req := &TableCellMergeInfo{}
	if builder.rowSpanFlag {
		req.RowSpan = &builder.rowSpan

	}
	if builder.colSpanFlag {
		req.ColSpan = &builder.colSpan

	}
	return req
}

type TableMeta struct {
	RowNum *int `json:"row_num,omitempty"` // 行数
	ColNum *int `json:"col_num,omitempty"` // 列数
}

type TableMetaBuilder struct {
	rowNum     int // 行数
	rowNumFlag bool
	colNum     int // 列数
	colNumFlag bool
}

func NewTableMetaBuilder() *TableMetaBuilder {
	builder := &TableMetaBuilder{}
	return builder
}

// 行数
//
// 示例值：3
func (builder *TableMetaBuilder) RowNum(rowNum int) *TableMetaBuilder {
	builder.rowNum = rowNum
	builder.rowNumFlag = true
	return builder
}

// 列数
//
// 示例值：3
func (builder *TableMetaBuilder) ColNum(colNum int) *TableMetaBuilder {
	builder.colNum = colNum
	builder.colNumFlag = true
	return builder
}

func (builder *TableMetaBuilder) Build() *TableMeta {
	req := &TableMeta{}
	if builder.rowNumFlag {
		req.RowNum = &builder.rowNum

	}
	if builder.colNumFlag {
		req.ColNum = &builder.colNum

	}
	return req
}

type Text struct {
	Text            *string `json:"text,omitempty"`             // 文字内容
	FontWeight      *string `json:"font_weight,omitempty"`      // 文字字重
	FontSize        *int    `json:"font_size,omitempty"`        // 文字大小
	HorizontalAlign *string `json:"horizontal_align,omitempty"` // 水平对齐
	VerticalAlign   *string `json:"vertical_align,omitempty"`   // 垂直对齐
}

type TextBuilder struct {
	text                string // 文字内容
	textFlag            bool
	fontWeight          string // 文字字重
	fontWeightFlag      bool
	fontSize            int // 文字大小
	fontSizeFlag        bool
	horizontalAlign     string // 水平对齐
	horizontalAlignFlag bool
	verticalAlign       string // 垂直对齐
	verticalAlignFlag   bool
}

func NewTextBuilder() *TextBuilder {
	builder := &TextBuilder{}
	return builder
}

// 文字内容
//
// 示例值：文字内容
func (builder *TextBuilder) Text(text string) *TextBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

// 文字字重
//
// 示例值：regular
func (builder *TextBuilder) FontWeight(fontWeight string) *TextBuilder {
	builder.fontWeight = fontWeight
	builder.fontWeightFlag = true
	return builder
}

// 文字大小
//
// 示例值：14
func (builder *TextBuilder) FontSize(fontSize int) *TextBuilder {
	builder.fontSize = fontSize
	builder.fontSizeFlag = true
	return builder
}

// 水平对齐
//
// 示例值：
func (builder *TextBuilder) HorizontalAlign(horizontalAlign string) *TextBuilder {
	builder.horizontalAlign = horizontalAlign
	builder.horizontalAlignFlag = true
	return builder
}

// 垂直对齐
//
// 示例值：
func (builder *TextBuilder) VerticalAlign(verticalAlign string) *TextBuilder {
	builder.verticalAlign = verticalAlign
	builder.verticalAlignFlag = true
	return builder
}

func (builder *TextBuilder) Build() *Text {
	req := &Text{}
	if builder.textFlag {
		req.Text = &builder.text

	}
	if builder.fontWeightFlag {
		req.FontWeight = &builder.fontWeight

	}
	if builder.fontSizeFlag {
		req.FontSize = &builder.fontSize

	}
	if builder.horizontalAlignFlag {
		req.HorizontalAlign = &builder.horizontalAlign

	}
	if builder.verticalAlignFlag {
		req.VerticalAlign = &builder.verticalAlign

	}
	return req
}

type WhiteboardNode struct {
	Id             *string         `json:"id,omitempty"`              // 节点 id
	Type           *string         `json:"type,omitempty"`            // 节点图形类型，目前创建节点仅支持创建图片、文本、基础图形等类型，读取到不支持创建的图形时只返回一些基础信息，如 id、type、text、style 等
	ParentId       *string         `json:"parent_id,omitempty"`       // 父节点 id
	Children       []string        `json:"children,omitempty"`        // 子节点
	X              *float64        `json:"x,omitempty"`               // 图形相对画布的 x 轴位置信息（存在父容器时为相对父容器的坐标，父容器为组合图形 group 时，坐标是穿透的），单位为 px
	Y              *float64        `json:"y,omitempty"`               // 图形相对画布的 y 轴位置信息（存在父容器时为相对父容器的坐标，父容器为组合图形 group 时，坐标是穿透的），单位为 px
	Angle          *float64        `json:"angle,omitempty"`           // 图形旋转角度
	Width          *float64        `json:"width,omitempty"`           // 图形宽度，单位为 px
	Height         *float64        `json:"height,omitempty"`          // 图形高度，单位为 px
	Text           *Text           `json:"text,omitempty"`            // 图形内文字
	Style          *Style          `json:"style,omitempty"`           // 图形样式
	Image          *Image          `json:"image,omitempty"`           // 图片
	CompositeShape *CompositeShape `json:"composite_shape,omitempty"` // 基础图形属性
	Connector      *Connector      `json:"connector,omitempty"`       // 连线属性
	Section        *Section        `json:"section,omitempty"`         // 分区属性
	Table          *Table          `json:"table,omitempty"`           // 表格属性
	MindMap        *MindMap        `json:"mind_map,omitempty"`        // 思维导图属性
}

type WhiteboardNodeBuilder struct {
	id                 string // 节点 id
	idFlag             bool
	type_              string // 节点图形类型，目前创建节点仅支持创建图片、文本、基础图形等类型，读取到不支持创建的图形时只返回一些基础信息，如 id、type、text、style 等
	typeFlag           bool
	parentId           string // 父节点 id
	parentIdFlag       bool
	children           []string // 子节点
	childrenFlag       bool
	x                  float64 // 图形相对画布的 x 轴位置信息（存在父容器时为相对父容器的坐标，父容器为组合图形 group 时，坐标是穿透的），单位为 px
	xFlag              bool
	y                  float64 // 图形相对画布的 y 轴位置信息（存在父容器时为相对父容器的坐标，父容器为组合图形 group 时，坐标是穿透的），单位为 px
	yFlag              bool
	angle              float64 // 图形旋转角度
	angleFlag          bool
	width              float64 // 图形宽度，单位为 px
	widthFlag          bool
	height             float64 // 图形高度，单位为 px
	heightFlag         bool
	text               *Text // 图形内文字
	textFlag           bool
	style              *Style // 图形样式
	styleFlag          bool
	image              *Image // 图片
	imageFlag          bool
	compositeShape     *CompositeShape // 基础图形属性
	compositeShapeFlag bool
	connector          *Connector // 连线属性
	connectorFlag      bool
	section            *Section // 分区属性
	sectionFlag        bool
	table              *Table // 表格属性
	tableFlag          bool
	mindMap            *MindMap // 思维导图属性
	mindMapFlag        bool
}

func NewWhiteboardNodeBuilder() *WhiteboardNodeBuilder {
	builder := &WhiteboardNodeBuilder{}
	return builder
}

// 节点 id
//
// 示例值：o1:1
func (builder *WhiteboardNodeBuilder) Id(id string) *WhiteboardNodeBuilder {
	builder.id = id
	builder.idFlag = true
	return builder
}

// 节点图形类型，目前创建节点仅支持创建图片、文本、基础图形等类型，读取到不支持创建的图形时只返回一些基础信息，如 id、type、text、style 等
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Type(type_ string) *WhiteboardNodeBuilder {
	builder.type_ = type_
	builder.typeFlag = true
	return builder
}

// 父节点 id
//
// 示例值：o1:1
func (builder *WhiteboardNodeBuilder) ParentId(parentId string) *WhiteboardNodeBuilder {
	builder.parentId = parentId
	builder.parentIdFlag = true
	return builder
}

// 子节点
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Children(children []string) *WhiteboardNodeBuilder {
	builder.children = children
	builder.childrenFlag = true
	return builder
}

// 图形相对画布的 x 轴位置信息（存在父容器时为相对父容器的坐标，父容器为组合图形 group 时，坐标是穿透的），单位为 px
//
// 示例值：100
func (builder *WhiteboardNodeBuilder) X(x float64) *WhiteboardNodeBuilder {
	builder.x = x
	builder.xFlag = true
	return builder
}

// 图形相对画布的 y 轴位置信息（存在父容器时为相对父容器的坐标，父容器为组合图形 group 时，坐标是穿透的），单位为 px
//
// 示例值：100
func (builder *WhiteboardNodeBuilder) Y(y float64) *WhiteboardNodeBuilder {
	builder.y = y
	builder.yFlag = true
	return builder
}

// 图形旋转角度
//
// 示例值：100
func (builder *WhiteboardNodeBuilder) Angle(angle float64) *WhiteboardNodeBuilder {
	builder.angle = angle
	builder.angleFlag = true
	return builder
}

// 图形宽度，单位为 px
//
// 示例值：100
func (builder *WhiteboardNodeBuilder) Width(width float64) *WhiteboardNodeBuilder {
	builder.width = width
	builder.widthFlag = true
	return builder
}

// 图形高度，单位为 px
//
// 示例值：100
func (builder *WhiteboardNodeBuilder) Height(height float64) *WhiteboardNodeBuilder {
	builder.height = height
	builder.heightFlag = true
	return builder
}

// 图形内文字
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Text(text *Text) *WhiteboardNodeBuilder {
	builder.text = text
	builder.textFlag = true
	return builder
}

// 图形样式
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Style(style *Style) *WhiteboardNodeBuilder {
	builder.style = style
	builder.styleFlag = true
	return builder
}

// 图片
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Image(image *Image) *WhiteboardNodeBuilder {
	builder.image = image
	builder.imageFlag = true
	return builder
}

// 基础图形属性
//
// 示例值：
func (builder *WhiteboardNodeBuilder) CompositeShape(compositeShape *CompositeShape) *WhiteboardNodeBuilder {
	builder.compositeShape = compositeShape
	builder.compositeShapeFlag = true
	return builder
}

// 连线属性
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Connector(connector *Connector) *WhiteboardNodeBuilder {
	builder.connector = connector
	builder.connectorFlag = true
	return builder
}

// 分区属性
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Section(section *Section) *WhiteboardNodeBuilder {
	builder.section = section
	builder.sectionFlag = true
	return builder
}

// 表格属性
//
// 示例值：
func (builder *WhiteboardNodeBuilder) Table(table *Table) *WhiteboardNodeBuilder {
	builder.table = table
	builder.tableFlag = true
	return builder
}

// 思维导图属性
//
// 示例值：
func (builder *WhiteboardNodeBuilder) MindMap(mindMap *MindMap) *WhiteboardNodeBuilder {
	builder.mindMap = mindMap
	builder.mindMapFlag = true
	return builder
}

func (builder *WhiteboardNodeBuilder) Build() *WhiteboardNode {
	req := &WhiteboardNode{}
	if builder.idFlag {
		req.Id = &builder.id

	}
	if builder.typeFlag {
		req.Type = &builder.type_

	}
	if builder.parentIdFlag {
		req.ParentId = &builder.parentId

	}
	if builder.childrenFlag {
		req.Children = builder.children
	}
	if builder.xFlag {
		req.X = &builder.x

	}
	if builder.yFlag {
		req.Y = &builder.y

	}
	if builder.angleFlag {
		req.Angle = &builder.angle

	}
	if builder.widthFlag {
		req.Width = &builder.width

	}
	if builder.heightFlag {
		req.Height = &builder.height

	}
	if builder.textFlag {
		req.Text = builder.text
	}
	if builder.styleFlag {
		req.Style = builder.style
	}
	if builder.imageFlag {
		req.Image = builder.image
	}
	if builder.compositeShapeFlag {
		req.CompositeShape = builder.compositeShape
	}
	if builder.connectorFlag {
		req.Connector = builder.connector
	}
	if builder.sectionFlag {
		req.Section = builder.section
	}
	if builder.tableFlag {
		req.Table = builder.table
	}
	if builder.mindMapFlag {
		req.MindMap = builder.mindMap
	}
	return req
}

type DownloadAsImageWhiteboardReqBuilder struct {
	apiReq *larkcore.ApiReq
}

func NewDownloadAsImageWhiteboardReqBuilder() *DownloadAsImageWhiteboardReqBuilder {
	builder := &DownloadAsImageWhiteboardReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 画板唯一标识
//
// 示例值：Ru8nwrWFOhEmaFbEU2VbPRsHcxb
func (builder *DownloadAsImageWhiteboardReqBuilder) WhiteboardId(whiteboardId string) *DownloadAsImageWhiteboardReqBuilder {
	builder.apiReq.PathParams.Set("whiteboard_id", fmt.Sprint(whiteboardId))
	return builder
}

func (builder *DownloadAsImageWhiteboardReqBuilder) Build() *DownloadAsImageWhiteboardReq {
	req := &DownloadAsImageWhiteboardReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	return req
}

type DownloadAsImageWhiteboardReq struct {
	apiReq *larkcore.ApiReq
}

type DownloadAsImageWhiteboardResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	File     io.Reader `json:"-"`
	FileName string    `json:"-"`
}

func (resp *DownloadAsImageWhiteboardResp) Success() bool {
	return resp.Code == 0
}

func (resp *DownloadAsImageWhiteboardResp) WriteFile(fileName string) error {
	bs, err := ioutil.ReadAll(resp.File)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, bs, 0666)
	if err != nil {
		return err
	}
	return nil
}

type ListWhiteboardNodeReqBuilder struct {
	apiReq *larkcore.ApiReq
}

func NewListWhiteboardNodeReqBuilder() *ListWhiteboardNodeReqBuilder {
	builder := &ListWhiteboardNodeReqBuilder{}
	builder.apiReq = &larkcore.ApiReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	return builder
}

// 画板唯一标识
//
// 示例值：Ru8nwrWFOhEmaFbEU2VbPRsHcxb
func (builder *ListWhiteboardNodeReqBuilder) WhiteboardId(whiteboardId string) *ListWhiteboardNodeReqBuilder {
	builder.apiReq.PathParams.Set("whiteboard_id", fmt.Sprint(whiteboardId))
	return builder
}

func (builder *ListWhiteboardNodeReqBuilder) Build() *ListWhiteboardNodeReq {
	req := &ListWhiteboardNodeReq{}
	req.apiReq = &larkcore.ApiReq{}
	req.apiReq.PathParams = builder.apiReq.PathParams
	return req
}

type ListWhiteboardNodeReq struct {
	apiReq *larkcore.ApiReq
}

type ListWhiteboardNodeRespData struct {
	Nodes []*WhiteboardNode `json:"nodes,omitempty"` // 查询结果
}

type ListWhiteboardNodeResp struct {
	*larkcore.ApiResp `json:"-"`
	larkcore.CodeError
	Data *ListWhiteboardNodeRespData `json:"data"` // 业务数据
}

func (resp *ListWhiteboardNodeResp) Success() bool {
	return resp.Code == 0
}
