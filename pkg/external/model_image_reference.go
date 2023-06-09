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
// ImageReference A summary of an image identity, including digest, id (if available), and any tags known to have ever been mapped to the digest
type ImageReference struct {
	// The image digest
	Digest string `json:"digest,omitempty"`
	// The image id if available
	Id string `json:"id,omitempty"`
	// Timestamp, in rfc3339 format, indicating when the image state became 'analyzed' in Nextlinux Engine.
	AnalyzedAt string `json:"analyzed_at,omitempty"`
	TagHistory []TagEntry `json:"tag_history,omitempty"`
}
