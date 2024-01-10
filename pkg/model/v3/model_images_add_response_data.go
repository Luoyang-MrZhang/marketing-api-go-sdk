/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ImagesAddResponseData struct {
	ImageId        *string   `json:"image_id,omitempty"`
	ImageWidth     *int64    `json:"image_width,omitempty"`
	ImageHeight    *int64    `json:"image_height,omitempty"`
	ImageFileSize  *int64    `json:"image_file_size,omitempty"`
	ImageType      ImageType `json:"image_type,omitempty"`
	ImageSignature *string   `json:"image_signature,omitempty"`
	OuterImageId   *string   `json:"outer_image_id,omitempty"`
	PreviewUrl     *string   `json:"preview_url,omitempty"`
	Description    *string   `json:"description,omitempty"`
}
