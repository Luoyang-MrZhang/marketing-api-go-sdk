/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type VideosGetListStruct struct {
	VideoId                  *int64            `json:"video_id,omitempty"`
	Width                    *int64            `json:"width,omitempty"`
	Height                   *int64            `json:"height,omitempty"`
	VideoFrames              *int64            `json:"video_frames,omitempty"`
	VideoFps                 *float64          `json:"video_fps,omitempty"`
	VideoCodec               *string           `json:"video_codec,omitempty"`
	VideoBitRate             *int64            `json:"video_bit_rate,omitempty"`
	AudioCodec               *string           `json:"audio_codec,omitempty"`
	AudioBitRate             *int64            `json:"audio_bit_rate,omitempty"`
	FileSize                 *int64            `json:"file_size,omitempty"`
	Type_                    MediaType         `json:"type,omitempty"`
	Signature                *string           `json:"signature,omitempty"`
	SystemStatus             MediaSystemStatus `json:"system_status,omitempty"`
	Description              *string           `json:"description,omitempty"`
	PreviewUrl               *string           `json:"preview_url,omitempty"`
	KeyFrameImageUrl         *string           `json:"key_frame_image_url,omitempty"`
	CreatedTime              *int64            `json:"created_time,omitempty"`
	LastModifiedTime         *int64            `json:"last_modified_time,omitempty"`
	VideoProfileName         *string           `json:"video_profile_name,omitempty"`
	AudioSampleRate          *int64            `json:"audio_sample_rate,omitempty"`
	MaxKeyframeInterval      *int64            `json:"max_keyframe_interval,omitempty"`
	MinKeyframeInterval      *int64            `json:"min_keyframe_interval,omitempty"`
	SampleAspectRatio        *string           `json:"sample_aspect_ratio,omitempty"`
	AudioProfileName         *string           `json:"audio_profile_name,omitempty"`
	ScanType                 *string           `json:"scan_type,omitempty"`
	ImageDurationMillisecond *int64            `json:"image_duration_millisecond,omitempty"`
	AudioDurationMillisecond *int64            `json:"audio_duration_millisecond,omitempty"`
	SourceType               MediaSourceType   `json:"source_type,omitempty"`
	ProductCatalogId         *string           `json:"product_catalog_id,omitempty"`
	ProductOuterId           *string           `json:"product_outer_id,omitempty"`
	SourceReferenceId        *string           `json:"source_reference_id,omitempty"`
	OwnerAccountId           *string           `json:"owner_account_id,omitempty"`
	Status                   MediaStatusType   `json:"status,omitempty"`
}
