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
// ContentFilesResponse File content listings from images
type ContentFilesResponse struct {
	ImageDigest string `json:"imageDigest,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Content []ContentFilesResponseContent `json:"content,omitempty"`
}