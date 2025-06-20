/*
OpenPay API

super charge your subscription management.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package getopenpay

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// PaymentIntentsAPIService PaymentIntentsAPI service
type PaymentIntentsAPIService service

type ApiGetPaymentIntentRequest struct {
	ctx context.Context
	ApiService *PaymentIntentsAPIService
	paymentIntentId string
	expand *[]string
}

func (r ApiGetPaymentIntentRequest) Expand(expand []string) ApiGetPaymentIntentRequest {
	r.expand = &expand
	return r
}

func (r ApiGetPaymentIntentRequest) Execute() (*PaymentIntentExternal, *http.Response, error) {
	return r.ApiService.GetPaymentIntentExecute(r)
}

/*
GetPaymentIntent Get Payment Intent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentIntentId
 @return ApiGetPaymentIntentRequest
*/
func (a *PaymentIntentsAPIService) GetPaymentIntent(ctx context.Context, paymentIntentId string) ApiGetPaymentIntentRequest {
	return ApiGetPaymentIntentRequest{
		ApiService: a,
		ctx: ctx,
		paymentIntentId: paymentIntentId,
	}
}

// Execute executes the request
//  @return PaymentIntentExternal
func (a *PaymentIntentsAPIService) GetPaymentIntentExecute(r ApiGetPaymentIntentRequest) (*PaymentIntentExternal, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntentExternal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.GetPaymentIntent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents/{payment_intent_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_intent_id"+"}", url.PathEscape(parameterValueToString(r.paymentIntentId, "paymentIntentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expand != nil {
		t := *r.expand
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "expand", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "expand", t, "form", "multi")
		}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiListPaymentIntentsRequest struct {
	ctx context.Context
	ApiService *PaymentIntentsAPIService
	paymentIntentQueryParams *PaymentIntentQueryParams
}

func (r ApiListPaymentIntentsRequest) PaymentIntentQueryParams(paymentIntentQueryParams PaymentIntentQueryParams) ApiListPaymentIntentsRequest {
	r.paymentIntentQueryParams = &paymentIntentQueryParams
	return r
}

func (r ApiListPaymentIntentsRequest) Execute() (*ListResponsePaymentIntentExternal, *http.Response, error) {
	return r.ApiService.ListPaymentIntentsExecute(r)
}

/*
ListPaymentIntents List Payment Intents

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListPaymentIntentsRequest
*/
func (a *PaymentIntentsAPIService) ListPaymentIntents(ctx context.Context) ApiListPaymentIntentsRequest {
	return ApiListPaymentIntentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListResponsePaymentIntentExternal
func (a *PaymentIntentsAPIService) ListPaymentIntentsExecute(r ApiListPaymentIntentsRequest) (*ListResponsePaymentIntentExternal, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListResponsePaymentIntentExternal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.ListPaymentIntents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.paymentIntentQueryParams == nil {
		return localVarReturnValue, nil, reportError("paymentIntentQueryParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.paymentIntentQueryParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiUpdatePaymentIntentRequest struct {
	ctx context.Context
	ApiService *PaymentIntentsAPIService
	paymentIntentId string
	updatePaymentIntent *UpdatePaymentIntent
}

func (r ApiUpdatePaymentIntentRequest) UpdatePaymentIntent(updatePaymentIntent UpdatePaymentIntent) ApiUpdatePaymentIntentRequest {
	r.updatePaymentIntent = &updatePaymentIntent
	return r
}

func (r ApiUpdatePaymentIntentRequest) Execute() (*PaymentIntentExternal, *http.Response, error) {
	return r.ApiService.UpdatePaymentIntentExecute(r)
}

/*
UpdatePaymentIntent Update Payment Intent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param paymentIntentId
 @return ApiUpdatePaymentIntentRequest
*/
func (a *PaymentIntentsAPIService) UpdatePaymentIntent(ctx context.Context, paymentIntentId string) ApiUpdatePaymentIntentRequest {
	return ApiUpdatePaymentIntentRequest{
		ApiService: a,
		ctx: ctx,
		paymentIntentId: paymentIntentId,
	}
}

// Execute executes the request
//  @return PaymentIntentExternal
func (a *PaymentIntentsAPIService) UpdatePaymentIntentExecute(r ApiUpdatePaymentIntentRequest) (*PaymentIntentExternal, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaymentIntentExternal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentIntentsAPIService.UpdatePaymentIntent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-intents/{payment_intent_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_intent_id"+"}", url.PathEscape(parameterValueToString(r.paymentIntentId, "paymentIntentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updatePaymentIntent == nil {
		return localVarReturnValue, nil, reportError("updatePaymentIntent is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updatePaymentIntent
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
