/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// UnionPackageSysStatus : 广告包渠道包状态
type UnionPackageSysStatus string

// List of UnionPackageSysStatus
const (
	UnionPackageSysStatus_PASSED          UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_PASSED"
	UnionPackageSysStatus_REVIEWING       UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_REVIEWING"
	UnionPackageSysStatus_DENIED          UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_DENIED"
	UnionPackageSysStatus_DENIED_AGAIN    UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_DENIED_AGAIN"
	UnionPackageSysStatus_REVIEWING_AGAIN UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_REVIEWING_AGAIN"
	UnionPackageSysStatus_ON_OFFLINE      UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_ON_OFFLINE"
	UnionPackageSysStatus_OFFLINE         UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_OFFLINE"
	UnionPackageSysStatus_DRAFT           UnionPackageSysStatus = "CHANNEL_PACKAGE_STATUS_DRAFT"
)
