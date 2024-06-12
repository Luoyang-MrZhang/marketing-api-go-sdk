/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// RTA策略绑定信息
type RtaTargetBindListRecord struct {
	Id            *int64  `json:"Id,omitempty"`
	OuterTargetId *string `json:"OuterTargetId,omitempty"`
	TargetType    *int64  `json:"TargetType,omitempty"`
	IsMp          *int64  `json:"IsMp,omitempty"`
	UId           *int64  `json:"UId,omitempty"`
	OptPlatform   *int64  `json:"OptPlatform,omitempty"`
}
