/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 任务所需条件
type TaskSpec struct {
	UpdateUnionPositionPackageSpec                   *[]UpdateUnionPositionPackageItem                   `json:"update_union_position_package_spec,omitempty"`
	UpdateExcludeUnionPositionPackageSpec            *[]UpdateExcludeUnionPositionPackageItem            `json:"update_exclude_union_position_package_spec,omitempty"`
	UpdateDeepConversionBehaviorBidSpec              *[]UpdateDeepConversionBehaviorBidItem              `json:"update_deep_conversion_behavior_bid_spec,omitempty"`
	DeleteAdgroupSpec                                *[]DeleteAdgroupItem                                `json:"delete_adgroup_spec,omitempty"`
	UpdateAdgroupDeepConversionWorthRateSpec         *[]UpdateAdgroupDeepConversionWorthRateItem         `json:"update_adgroup_deep_conversion_worth_rate_spec,omitempty"`
	TargetingsShareSpec                              *[]TargetingsShareItem                              `json:"targetings_share_spec,omitempty"`
	UpdateAdgroupConfiguredStatusSpec                *[]UpdateAdgroupConfiguredStatusItem                `json:"update_adgroup_configured_status_spec,omitempty"`
	UpdateAdgroupDailyBudgetSpec                     *[]UpdateAdgroupDailyBudgetItem                     `json:"update_adgroup_daily_budget_spec,omitempty"`
	UpdateAdgroupAutoAcquisitionSpec                 *[]UpdateAdgroupAutoAcquisitionItem                 `json:"update_adgroup_auto_acquisition_spec,omitempty"`
	UpdateAdgroupDeepConversionWorthAdvancedRateSpec *[]UpdateAdgroupDeepConversionWorthAdvancedRateItem `json:"update_adgroup_deep_conversion_worth_advanced_rate_spec,omitempty"`
	UpdateDeepConversionBehaviorAdvancedBidSpec      *[]UpdateDeepConversionBehaviorAdvancedBidItem      `json:"update_deep_conversion_behavior_advanced_bid_spec,omitempty"`
	ReplyFinderObjectCommentSpec                     *[]ReplyFinderObjectCommentItem                     `json:"reply_finder_object_comment_spec,omitempty"`
	DeleteFinderObjectCommentSpec                    *[]DeleteFinderObjectCommentItem                    `json:"delete_finder_object_comment_spec,omitempty"`
	UpdateFinderObjectCommentFlagSpec                *[]UpdateFinderObjectCommentFlagItem                `json:"update_finder_object_comment_flag_spec,omitempty"`
	UpdateAdgroupTimeSpec                            *[]UpdateAdgroupTimeItem                            `json:"update_adgroup_time_spec,omitempty"`
	UpdateAdgroupDateSpec                            *[]UpdateAdgroupDateItem                            `json:"update_adgroup_date_spec,omitempty"`
	UpdateAdgroupBidAmountSpec                       *[]UpdateAdgroupBidAmountItem                       `json:"update_adgroup_bid_amount_spec,omitempty"`
	UpdateAdgroupBindRtaPolicySpec                   *[]UpdateAdgroupBindRtaPolicyItem                   `json:"update_adgroup_bind_rta_policy_spec,omitempty"`
	UpdateAdcreativeObjectCommentFlagSpec            *[]UpdateAdcreativeObjectCommentFlagItem            `json:"update_adcreative_object_comment_flag_spec,omitempty"`
}
