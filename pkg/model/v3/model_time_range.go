/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 时间范围
type TimeRange struct {
	StartTime *int64   `json:"start_time,omitempty"`
	EndTime   *int64   `json:"end_time,omitempty"`
	TimeType  TimeType `json:"time_type,omitempty"`
}
