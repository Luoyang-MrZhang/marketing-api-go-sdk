/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProductItemsBatchUpdateRequest struct {
	AccountId         *int64               `json:"account_id,omitempty"`
	CatalogId         *int64               `json:"catalog_id,omitempty"`
	ProductUpdateList *[]ProductUpdateItem `json:"product_update_list,omitempty"`
}
