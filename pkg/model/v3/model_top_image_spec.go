/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 顶部图片组件元素
type TopImageSpec struct {
	ImageIdList *string `json:"image_id_list,omitempty"`
	Width       *int64  `json:"width,omitempty"`
	Height      *int64  `json:"height,omitempty"`
	AdLocation  *string `json:"ad_location,omitempty"`
	OuterStyle  *int64  `json:"outer_style,omitempty"`
}
