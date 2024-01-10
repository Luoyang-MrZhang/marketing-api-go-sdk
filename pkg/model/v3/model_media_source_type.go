/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// MediaSourceType : 视频来源
type MediaSourceType string

// List of MediaSourceType
const (
	MediaSourceType_UNSUPPORTED     MediaSourceType = "SOURCE_TYPE_UNSUPPORTED"
	MediaSourceType_LOCAL           MediaSourceType = "SOURCE_TYPE_LOCAL"
	MediaSourceType_API             MediaSourceType = "SOURCE_TYPE_API"
	MediaSourceType_VIDEO_MAKER_XSJ MediaSourceType = "SOURCE_TYPE_VIDEO_MAKER_XSJ"
	MediaSourceType_TCC             MediaSourceType = "SOURCE_TYPE_TCC"
	MediaSourceType_DERIVE          MediaSourceType = "SOURCE_TYPE_DERIVE"
	MediaSourceType_DERIVATION      MediaSourceType = "SOURCE_TYPE_DERIVATION"
)
