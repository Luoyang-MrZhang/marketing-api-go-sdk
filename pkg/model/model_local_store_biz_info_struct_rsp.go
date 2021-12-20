/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 门店经营信息
type LocalStoreBizInfoStructRsp struct {
	PoiId              *string                `json:"poi_id,omitempty"`
	CustomerProfile    *CustomerProfileStruct `json:"customer_profile,omitempty"`
	CustomerPerCost    *int64                 `json:"customer_per_cost,omitempty"`
	FirstCategoryId    *int64                 `json:"first_category_id,omitempty"`
	SecondCategoryId   *int64                 `json:"second_category_id,omitempty"`
	ThirdCategoryId    *int64                 `json:"third_category_id,omitempty"`
	FourthCategoryId   *int64                 `json:"fourth_category_id,omitempty"`
	FirstCategoryName  *string                `json:"first_category_name,omitempty"`
	SecondCategoryName *string                `json:"second_category_name,omitempty"`
	ThirdCategoryName  *string                `json:"third_category_name,omitempty"`
	FourthCategoryName *string                `json:"fourth_category_name,omitempty"`
	PeakPeriod         *[]PeakPeriod          `json:"peak_period,omitempty"`
}