/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信搜一搜广告品牌官方区原始编辑器表单数据
type WxSearchBrandAreaOriginData struct {
	ZoneConfig *WxSearchBrandAreaZoneConfig              `json:"zone_config,omitempty"`
	Items      *[]WxSearchBrandAreaOriginDataItemsStruct `json:"items,omitempty"`
}