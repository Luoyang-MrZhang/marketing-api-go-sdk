/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 视频号列表
type WechatChannelsAccountStruct struct {
	WechatChannelsAccountId   *string `json:"wechat_channels_account_id,omitempty"`
	WechatChannelsAccountName *string `json:"wechat_channels_account_name,omitempty"`
	CreatedTime               *int64  `json:"created_time,omitempty"`
	LastModifiedTime          *int64  `json:"last_modified_time,omitempty"`
}
