/*
LogCom API

LogCom Swagger documentation

API version: 1.2.21
Contact: laborit@blutspende.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logcomapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// ConsoleLogApiService ConsoleLogApi service
type ConsoleLogApiService service

type ApiCreateConsoleLogV1IntRequest struct {
	ctx context.Context
	ApiService *ConsoleLogApiService
	model *CreateConsoleLogRequestDTO
}

// The console log DTO
func (r ApiCreateConsoleLogV1IntRequest) Model(model CreateConsoleLogRequestDTO) ApiCreateConsoleLogV1IntRequest {
	r.model = &model
	return r
}

func (r ApiCreateConsoleLogV1IntRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateConsoleLogV1IntExecute(r)
}

/*
CreateConsoleLogV1Int Create console log

Creates a new console log

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateConsoleLogV1IntRequest
*/
func (a *ConsoleLogApiService) CreateConsoleLogV1Int(ctx context.Context) ApiCreateConsoleLogV1IntRequest {
	return ApiCreateConsoleLogV1IntRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConsoleLogApiService) CreateConsoleLogV1IntExecute(r ApiCreateConsoleLogV1IntRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsoleLogApiService.CreateConsoleLogV1Int")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/int/console-logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.model == nil {
		return nil, reportError("model is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.model
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetConsoleLogsV1Request struct {
	ctx context.Context
	ApiService *ConsoleLogApiService
	page *int32
	pageSize *int32
	filter *string
	direction *string
	sort *string
}

// The desired page number
func (r ApiGetConsoleLogsV1Request) Page(page int32) ApiGetConsoleLogsV1Request {
	r.page = &page
	return r
}

// The desired number of items per page
func (r ApiGetConsoleLogsV1Request) PageSize(pageSize int32) ApiGetConsoleLogsV1Request {
	r.pageSize = &pageSize
	return r
}

// The search term
func (r ApiGetConsoleLogsV1Request) Filter(filter string) ApiGetConsoleLogsV1Request {
	r.filter = &filter
	return r
}

// The sorting direction
func (r ApiGetConsoleLogsV1Request) Direction(direction string) ApiGetConsoleLogsV1Request {
	r.direction = &direction
	return r
}

// The sorting parameter
func (r ApiGetConsoleLogsV1Request) Sort(sort string) ApiGetConsoleLogsV1Request {
	r.sort = &sort
	return r
}

func (r ApiGetConsoleLogsV1Request) Execute() (*ConsoleLogListPageResponse, *http.Response, error) {
	return r.ApiService.GetConsoleLogsV1Execute(r)
}

/*
GetConsoleLogsV1 Get console logs

Gets all console log

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConsoleLogsV1Request
*/
func (a *ConsoleLogApiService) GetConsoleLogsV1(ctx context.Context) ApiGetConsoleLogsV1Request {
	return ApiGetConsoleLogsV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ConsoleLogListPageResponse
func (a *ConsoleLogApiService) GetConsoleLogsV1Execute(r ApiGetConsoleLogsV1Request) (*ConsoleLogListPageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsoleLogListPageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConsoleLogApiService.GetConsoleLogsV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/console-logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.page == nil {
		return localVarReturnValue, nil, reportError("page is required and must be specified")
	}
	if *r.page < 0 {
		return localVarReturnValue, nil, reportError("page must be greater than 0")
	}
	if r.pageSize == nil {
		return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
	}
	if *r.pageSize < 0 {
		return localVarReturnValue, nil, reportError("pageSize must be greater than 0")
	}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
	}
	localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
