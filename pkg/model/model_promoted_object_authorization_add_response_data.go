/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type PromotedObjectAuthorizationAddResponseData struct {
	PromotedObjectId *string `json:"promoted_object_id,omitempty"`
	QrCodeUrl        *string `json:"qr_code_url,omitempty"`
	Description      *string `json:"description,omitempty"`
	ExpiredTime      *int64  `json:"expired_time,omitempty"`
	Agreement        *string `json:"agreement,omitempty"`
}
