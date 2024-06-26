/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type ConversionLinksGetListStruct struct {
	ConversionLinkId     *int64                      `json:"conversion_link_id,omitempty"`
	MarketingGoal        *[]string                   `json:"marketing_goal,omitempty"`
	MarketingTargetType  *string                     `json:"marketing_target_type,omitempty"`
	MarketingCarrierType MarketingCarrierType        `json:"marketing_carrier_type,omitempty"`
	GoldLinkType         GoldLinkType                `json:"gold_link_type,omitempty"`
	LandingPageAccess    *LandingPageAccess          `json:"landing_page_access,omitempty"`
	ConversionLinkNodes  *[]ConversionLinkNodeStruct `json:"conversion_link_nodes,omitempty"`
}
