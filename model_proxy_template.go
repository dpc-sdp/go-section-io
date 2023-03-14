/*
 * Section API
 *
 * Get edgey with the Section API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sectionio

type ProxyTemplate struct {
	Name string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Image string `json:"image,omitempty"`
	InitialStateUri string `json:"initialStateUri,omitempty"`
}
