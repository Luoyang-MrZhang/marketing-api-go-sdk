/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 链路节点信息
type ConversionLinkAssetNodeStruct struct {
	ConversionLinkNodeId     *int64             `json:"conversion_link_node_id,omitempty"`
	ConversionLinkNodeName   *string            `json:"conversion_link_node_name,omitempty"`
	ConversionLinkNodeIndex  *int64             `json:"conversion_link_node_index,omitempty"`
	ConversionLinkActionType ActionType         `json:"conversion_link_action_type,omitempty"`
	CarrierId                *int64             `json:"carrier_id,omitempty"`
	CarrierName              *string            `json:"carrier_name,omitempty"`
	OgId                     *[]int64           `json:"og_id,omitempty"`
	OgList                   *[]NodeOgStruct    `json:"og_list,omitempty"`
	RoiOgId                  *[]int64           `json:"roi_og_id,omitempty"`
	RoiOgList                *[]NodeRoiOgStruct `json:"roi_og_list,omitempty"`
	ConversionInfo           *[]ConvInfoStruct  `json:"conversion_info,omitempty"`
	OgStatus                 *int64             `json:"og_status,omitempty"`
	NodeStatus               *int64             `json:"node_status,omitempty"`
}
