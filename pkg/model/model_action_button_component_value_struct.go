/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 行动按钮结构
type ActionButtonComponentValueStruct struct {
	ButtonText *string                 `json:"button_text,omitempty"`
	JumpInfo   *[]LandingPageStructure `json:"jump_info,omitempty"`
}