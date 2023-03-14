/*
 * Section API
 *
 * Get edgey with the Section API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sectionio

type AccountUserUpdateResult struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	User *AccountUser `json:"user,omitempty"`
}
