/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type BarrageGetListStruct struct {
	Id           *int64    `json:"id,omitempty"`
	Text         *string   `json:"text,omitempty"`
	ReviewStatus SysStatus `json:"review_status,omitempty"`
}
