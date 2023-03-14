/*
 * Section API
 *
 * Get edgey with the Section API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sectionio

type Proxy struct {
	Href string `json:"href"`
	Name string `json:"name"`
	Image string `json:"image"`
	ProxyType *ProxyType `json:"proxyType"`
}