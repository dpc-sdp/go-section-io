/*
 * Section API
 *
 * Get edgey with the Section API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sectionio

type EnvironmentCreateRequest struct {
	Name string `json:"name"`
	SourceEnvironmentName string `json:"source_environment_name"`
	DomainName string `json:"domain_name"`
}
