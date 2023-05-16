/*
 * Nextlinux Engine API Server
 *
 * This is the Nextlinux Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.16
 * Contact: nurmi@nextlinux.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package external
// DeleteImageResponse Image deletion response containing status and details
type DeleteImageResponse struct {
	Digest string `json:"digest"`
	// Current status of the image deletion
	Status string `json:"status"`
	Detail string `json:"detail,omitempty"`
}