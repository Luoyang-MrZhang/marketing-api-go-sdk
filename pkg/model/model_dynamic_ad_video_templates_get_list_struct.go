/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type DynamicAdVideoTemplatesGetListStruct struct {
	TemplateId           *int64               `json:"template_id,omitempty"`
	TemplateName         *string              `json:"template_name,omitempty"`
	TemplateType         VideoTemplateType    `json:"template_type,omitempty"`
	ProductCatalogId     *int64               `json:"product_catalog_id,omitempty"`
	AdcreativeTemplateId *int64               `json:"adcreative_template_id,omitempty"`
	CoverImageUrl        *string              `json:"cover_image_url,omitempty"`
	IntroVideoUrl        *string              `json:"intro_video_url,omitempty"`
	DeliveryVideoUrl     *string              `json:"delivery_video_url,omitempty"`
	SubTemplateList      *[]SubTemplateStruct `json:"sub_template_list,omitempty"`
}
