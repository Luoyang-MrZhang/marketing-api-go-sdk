/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 选择按钮结构
type ChosenButtonStruct struct {
	LeftButton  *ChosenBaseButtonStruct `json:"left_button,omitempty"`
	RightButton *ChosenBaseButtonStruct `json:"right_button,omitempty"`
}
