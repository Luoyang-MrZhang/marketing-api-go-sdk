/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MaterialLabelsAddRequest struct {
	AccountId *int64         `json:"account_id,omitempty"`
	Labels    *[]CreateLabel `json:"labels,omitempty"`
}
