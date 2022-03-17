/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 小说章节信息
type Chapter struct {
	Subtitle  *string   `json:"subtitle,omitempty"`
	Chapterid *int64    `json:"chapterid,omitempty"`
	Text      *[]string `json:"text,omitempty"`
}