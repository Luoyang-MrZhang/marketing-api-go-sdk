/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProductCategoriesListGetResponseData struct {
	List     *[]CategoryItem `json:"list,omitempty"`
	PageInfo *Conf           `json:"page_info,omitempty"`
}
