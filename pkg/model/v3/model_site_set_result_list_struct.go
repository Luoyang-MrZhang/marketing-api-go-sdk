/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 版位粒度审核结果
type SiteSetResultListStruct struct {
	SiteSet                 SiteSetDefinition                    `json:"site_set,omitempty"`
	SystemStatus            ReviewResultStatus                   `json:"system_status,omitempty"`
	RejectMessage           *string                              `json:"reject_message,omitempty"`
	ElementRejectDetailInfo *[]ElementRejectDetailInfoListStruct `json:"element_reject_detail_info,omitempty"`
}
