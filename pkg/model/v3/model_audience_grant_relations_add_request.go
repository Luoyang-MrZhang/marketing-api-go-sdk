/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AudienceGrantRelationsAddRequest struct {
	AccountId      *int64            `json:"account_id,omitempty"`
	AudienceIdList *[]int64          `json:"audience_id_list,omitempty"`
	GrantType      AudienceGrantType `json:"grant_type,omitempty"`
	GrantSpec      *GrantSpec        `json:"grant_spec,omitempty"`
}
