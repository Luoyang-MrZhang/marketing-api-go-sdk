/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意规格和投放权限数据结构
type AdcreativeTemplateListStructAdpermit struct {
	CreativeTemplateId          *int64            `json:"creative_template_id,omitempty"`
	CreativeTemplateStyle       *string           `json:"creative_template_style,omitempty"`
	CreativeTemplateAppellation *string           `json:"creative_template_appellation,omitempty"`
	CreativeSampleImage         *string           `json:"creative_sample_image,omitempty"`
	SiteSet                     SiteSetDefinition `json:"site_set,omitempty"`
}
