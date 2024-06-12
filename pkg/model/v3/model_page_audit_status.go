/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// PageAuditStatus : 蹊径落地页状态
type PageAuditStatus string

// List of PageAuditStatus
const (
	PageAuditStatus_EDITING  PageAuditStatus = "LANDING_PAGE_STATUS_EDITING"
	PageAuditStatus_PENDING  PageAuditStatus = "LANDING_PAGE_STATUS_PENDING"
	PageAuditStatus_APPROVED PageAuditStatus = "LANDING_PAGE_STATUS_APPROVED"
	PageAuditStatus_REJECTED PageAuditStatus = "LANDING_PAGE_STATUS_REJECTED"
	PageAuditStatus_DELETED  PageAuditStatus = "LANDING_PAGE_STATUS_DELETED"
)
