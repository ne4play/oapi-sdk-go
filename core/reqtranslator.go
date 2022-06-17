package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type ReqTranslator struct {
}

func (translator *ReqTranslator) translate(ctx context.Context, input interface{}, tokenType AccessTokenType, config *Config, httpMethod, httpPath string, option *RequestOption) (*http.Request, error) {
	paths, queries, body := translator.parseInput(input, option)
	if _, ok := body.(*Formdata); ok {
		option.FileUpload = true
	}

	contentType, rawBody, err := translator.payload(body)
	if err != nil {
		return nil, err
	}

	fullURL, err := translator.getFullReqUrl(config.Domain, httpPath, paths, queries)
	if err != nil {
		return nil, err
	}

	req, err := translator.newHTTPRequest(ctx, httpMethod, fullURL, contentType, rawBody, tokenType, option, config)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func authorizationToHeader(req *http.Request, token string) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
}

func (translator *ReqTranslator) newHTTPRequest(ctx context.Context,
	httpMethod, url, contentType string, body []byte,
	accessTokenType AccessTokenType, option *RequestOption, config *Config) (*http.Request, error) {
	httpRequest, err := http.NewRequestWithContext(ctx, httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if option.RequestId != "" {
		httpRequest.Header.Add(customRequestId, option.RequestId)
	}
	for k, vs := range option.Header {
		for _, v := range vs {
			httpRequest.Header.Add(k, v)
		}
	}
	httpRequest.Header.Set(userAgentHeader, userAgent())
	if contentType != "" {
		httpRequest.Header.Set(contentTypeHeader, contentType)
	}
	switch accessTokenType {
	case AccessTokenTypeApp:
		appAccessToken := option.AppAccessToken
		if config.EnableTokenCache {
			appAccessToken, err = tokenManager.getAppAccessToken(ctx, config)
			if err != nil {
				return nil, err
			}
		}
		authorizationToHeader(httpRequest, appAccessToken)

	case AccessTokenTypeTenant:
		tenantAccessToken := option.TenantAccessToken
		if config.EnableTokenCache {
			tenantAccessToken, err = tokenManager.getTenantAccessToken(ctx, config, option.TenantKey)
			if err != nil {
				return nil, err
			}
		}
		authorizationToHeader(httpRequest, tenantAccessToken)

	case AccessTokenTypeUser:
		authorizationToHeader(httpRequest, option.UserAccessToken)
	}

	if err != nil {
		return nil, err
	}

	err = translator.signHelpdeskAuthToken(httpRequest, option.NeedHelpDeskAuth, config.HelpdeskAuthToken)
	if err != nil {
		return nil, err
	}
	return httpRequest, nil
}

func (translator *ReqTranslator) signHelpdeskAuthToken(rawRequest *http.Request, needHelpDeskAuth bool, authToken string) error {
	if needHelpDeskAuth {
		if authToken == "" {
			return errors.New("help desk API, please set the helpdesk information of lark.App")
		}
		rawRequest.Header.Set("X-Lark-Helpdesk-Authorization", authToken)
	}
	return nil
}

func (translator *ReqTranslator) getFullReqUrl(domain string, httpPath string, paths, queries map[string]interface{}) (string, error) {
	// path
	var pathSegs []string
	for _, p := range strings.Split(httpPath, "/") {
		if strings.Index(p, ":") == 0 {
			varName := p[1:]
			v, ok := paths[varName]
			if !ok {
				return "", fmt.Errorf("http path:%s, name: %s, not found value", httpPath, varName)
			}
			val := fmt.Sprint(v)
			if val == "" {
				return "", fmt.Errorf("http path:%s, name: %s, value is empty", httpPath, varName)
			}
			pathSegs = append(pathSegs, val)
			continue
		}
		pathSegs = append(pathSegs, p)
	}
	newPath := strings.Join(pathSegs, "/")
	if strings.Index(newPath, "http") != 0 {
		if strings.Index(newPath, "/open-apis") == 0 {
			newPath = fmt.Sprintf("%s%s", domain, newPath)
		} else {
			newPath = fmt.Sprintf("%s/open-apis/%s", domain, newPath)
		}
	}
	// query
	query := make(url.Values)
	for k, v := range queries {
		sv := reflect.ValueOf(v)
		if sv.Kind() == reflect.Slice || sv.Kind() == reflect.Array {
			for i := 0; i < sv.Len(); i++ {
				query.Add(k, fmt.Sprint(sv.Index(i)))
			}
		} else {
			query.Set(k, fmt.Sprint(v))
		}
	}
	if len(query) > 0 {
		newPath = fmt.Sprintf("%s?%s", newPath, query.Encode())
	}
	return newPath, nil
}

func (translator *ReqTranslator) payload(body interface{}) (string, []byte, error) {
	if fd, ok := body.(*Formdata); ok {
		return fd.content()
	}
	contentType := defaultContentType
	if body == nil {
		return contentType, nil, nil
	}
	bs, err := json.Marshal(body)
	return contentType, bs, err
}

func NewFormdata() *Formdata {
	return &Formdata{}
}

func (fd *Formdata) AddField(field string, val interface{}) *Formdata {
	if fd.fields == nil {
		fd.fields = map[string]interface{}{}
	}
	fd.fields[field] = val
	return fd
}

func (fd *Formdata) AddFile(field string, r io.Reader) *Formdata {
	return fd.AddField(field, r)
}

func (fd *Formdata) content() (string, []byte, error) {
	if fd.data != nil {
		return fd.data.contentType, fd.data.content, nil
	}
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key, val := range fd.fields {
		if r, ok := val.(io.Reader); ok {
			part, err := writer.CreateFormFile(key, "unknown-file")
			if err != nil {
				return "", nil, err
			}
			_, err = io.Copy(part, r)
			if err != nil {
				return "", nil, err
			}
			continue
		}
		err := writer.WriteField(key, fmt.Sprint(val))
		if err != nil {
			return "", nil, err
		}
	}
	contentType := writer.FormDataContentType()
	err := writer.Close()
	if err != nil {
		return "", nil, err
	}
	fd.data = &struct {
		content     []byte
		contentType string
	}{content: buf.Bytes(), contentType: contentType}
	return fd.data.contentType, fd.data.content, nil
}

type Formdata struct {
	fields map[string]interface{}
	data   *struct {
		content     []byte
		contentType string
	}
}

func (translator *ReqTranslator) parseInput(input interface{}, option *RequestOption) (map[string]interface{}, map[string]interface{}, interface{}) {
	if input == nil {
		return nil, nil, nil
	}
	if _, ok := input.(*Formdata); ok {
		return nil, nil, input
	}
	var hasHTTPTag bool
	paths, queries := map[string]interface{}{}, map[string]interface{}{}
	vv := reflect.ValueOf(input)
	vt := reflect.TypeOf(input)
	if vt.Kind() == reflect.Ptr {
		vv = vv.Elem()
		vt = vt.Elem()
	}
	if vt.Kind() != reflect.Struct {
		return nil, nil, input
	}
	var body interface{}
	for i := 0; i < vt.NumField(); i++ {
		fieldValue := vv.Field(i)
		fieldType := vt.Field(i)
		if path, ok := fieldType.Tag.Lookup("path"); ok {
			hasHTTPTag = true
			if path != "" && !isEmptyValue(fieldValue) {
				paths[path] = reflect.Indirect(fieldValue).Interface()
			}
			continue
		}
		if query, ok := fieldType.Tag.Lookup("query"); ok {
			hasHTTPTag = true
			if query != "" && !isEmptyValue(fieldValue) {
				queries[query] = reflect.Indirect(fieldValue).Interface()
			}
			continue
		}
		if _, ok := fieldType.Tag.Lookup("body"); ok {
			hasHTTPTag = true
			body = fieldValue.Interface()
		}
	}
	if !hasHTTPTag {
		body = input
		if option.FileUpload {
			body = toFormdata(input)
		}
		return nil, nil, body
	}
	if body != nil {
		if option.FileUpload {
			body = toFormdata(body)
		}
	}
	return paths, queries, body
}

func toFormdata(body interface{}) *Formdata {
	formdata := &Formdata{}
	v := reflect.ValueOf(body)
	t := reflect.TypeOf(body)
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		if isEmptyValue(fieldValue) {
			continue
		}
		if fieldName := fieldType.Tag.Get("json"); fieldName != "" {
			fieldName = strings.TrimSuffix(fieldName, ",omitempty")
			formdata.AddField(fieldName, reflect.Indirect(fieldValue).Interface())
		}
	}
	return formdata
}