/*
 * Section API
 *
 * Get edgey with the Section API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package sectionio

// A JSONPatch document as defined by RFC 6902
type Rfc6902Operation struct {
	// The operation to be performed
	Op string `json:"op"`
	// A JSON-Pointer
	Path string `json:"path"`
	// The value to be used within the operations.
	Value interface{} `json:"value,omitempty"`
	// A string containing a JSON Pointer value.
	From string `json:"from,omitempty"`
}
