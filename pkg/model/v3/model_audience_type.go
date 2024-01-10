/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// AudienceType : 人群类型
type AudienceType string

// List of AudienceType
const (
	AudienceType_CUSTOMER_FILE AudienceType = "CUSTOMER_FILE"
	AudienceType_LOOKALIKE     AudienceType = "LOOKALIKE"
	AudienceType_USER_ACTION   AudienceType = "USER_ACTION"
	AudienceType_KEYWORD       AudienceType = "KEYWORD"
	AudienceType_AD            AudienceType = "AD"
	AudienceType_COMBINE       AudienceType = "COMBINE"
	AudienceType_LABEL         AudienceType = "LABEL"
)
