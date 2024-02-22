/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回信息
type ResponseStruct struct {
	HttpCode *int64          `json:"http_code,omitempty"`
	Headers  *[]HeaderStruct `json:"headers,omitempty"`
	Body     *string         `json:"body,omitempty"`
}
