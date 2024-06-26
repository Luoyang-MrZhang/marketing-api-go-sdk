/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// WechatChannelsAuthType : 授权类型
type WechatChannelsAuthType string

// List of WechatChannelsAuthType
const (
	WechatChannelsAuthType_DEFAULT                     WechatChannelsAuthType = "DEFAULT"
	WechatChannelsAuthType_SAME_CORPORATION            WechatChannelsAuthType = "SAME_CORPORATION"
	WechatChannelsAuthType_OVER_CORPORATION            WechatChannelsAuthType = "OVER_CORPORATION"
	WechatChannelsAuthType_OVER_CORPORATION_SAME_GROUP WechatChannelsAuthType = "OVER_CORPORATION_SAME_GROUP"
	WechatChannelsAuthType_OVER_CORPORATION_EMPLOYMENT WechatChannelsAuthType = "OVER_CORPORATION_EMPLOYMENT"
)
