/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 线索数据信息结构
type LeadsInfoListStruct struct {
	VideoChannelId   *string `json:"video_channel_id,omitempty"`
	VideoChannelName *string `json:"video_channel_name,omitempty"`
	Date             *int64  `json:"date,omitempty"`
	DealerId         *int64  `json:"dealer_id,omitempty"`
	DealerName       *string `json:"dealer_name,omitempty"`
	LeadsCnt         *int64  `json:"leads_cnt,omitempty"`
	FormLeadsCnt     *int64  `json:"form_leads_cnt,omitempty"`
	WecomLeadsCnt    *int64  `json:"wecom_leads_cnt,omitempty"`
	ConsultLeadsCnt  *int64  `json:"consult_leads_cnt,omitempty"`
	AdLeadsCnt       *int64  `json:"ad_leads_cnt,omitempty"`
	NaturalLeadsCnt  *int64  `json:"natural_leads_cnt,omitempty"`
}
