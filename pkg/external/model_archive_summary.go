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
// ArchiveSummary A summarization of the available archives, a place to for long-term storage of audit, analysis, or other data to remove it from the system's working set but keep it available.
type ArchiveSummary struct {
	Images AnalysisArchiveSummary `json:"images,omitempty"`
	Rules AnalysisArchiveRulesSummary `json:"rules,omitempty"`
}
