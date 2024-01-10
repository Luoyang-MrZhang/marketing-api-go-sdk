/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 渠道包信息
type ChannelPackageStruct struct {
	ChannelId        *string               `json:"channel_id,omitempty"`
	ChannelName      *string               `json:"channel_name,omitempty"`
	SystemStatus     UnionPackageSysStatus `json:"system_status,omitempty"`
	CreatedTime      *int64                `json:"created_time,omitempty"`
	LastModifiedTime *int64                `json:"last_modified_time,omitempty"`
}
