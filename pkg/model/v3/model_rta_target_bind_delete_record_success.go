/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 解绑成功的记录
type RtaTargetBindDeleteRecordSuccess struct {
	Id         *int64 `json:"Id,omitempty"`
	TargetType *int64 `json:"TargetType,omitempty"`
	IsMp       *int64 `json:"IsMp,omitempty"`
}
