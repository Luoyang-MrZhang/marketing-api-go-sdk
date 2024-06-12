/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创建页面返回数据结构
type CreatePageResponse struct {
	LandingPageId *int64  `json:"landingPageId,omitempty"`
	Id            *string `json:"id,omitempty"`
	Code          *int64  `json:"code,omitempty"`
	Message       *string `json:"message,omitempty"`
}
