/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type UnionPositionPackagesUpdateResponse struct {
	Code      *int64                                   `json:"code,omitempty"`
	Message   *string                                  `json:"message,omitempty"`
	MessageCn *string                                  `json:"message_cn,omitempty"`
	Errors    *[]ApiErrorStruct                        `json:"errors,omitempty"`
	Data      *UnionPositionPackagesUpdateResponseData `json:"data,omitempty"`
}
