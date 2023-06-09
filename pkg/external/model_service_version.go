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
// ServiceVersion Version information for a service
type ServiceVersion struct {
	Service ServiceVersionService `json:"service,omitempty"`
	Api ServiceVersionApi `json:"api,omitempty"`
	Db ServiceVersionDb `json:"db,omitempty"`
}
