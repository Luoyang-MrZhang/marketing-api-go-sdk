/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PromotedObjectAuthorizationAddRequest struct {
	AccountId          *int64             `json:"account_id,omitempty"`
	PromotedObjectType PromotedObjectType `json:"promoted_object_type,omitempty"`
	PromotedObjectName *string            `json:"promoted_object_name,omitempty"`
	AuthSpec           *AuthSpec          `json:"auth_spec,omitempty"`
	AuthBeginDate      *int64             `json:"auth_begin_date,omitempty"`
	AuthTtl            *int64             `json:"auth_ttl,omitempty"`
}
