/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
import (
	"time"
)
// BatchQueryRequestMediaV1 struct for BatchQueryRequestMediaV1
type BatchQueryRequestMediaV1 struct {
	// From timestamp
	From time.Time `json:"from"`
	// Resolution in seconds
	Interval int64 `json:"interval,omitempty"`
	// Query
	Q string `json:"q"`
	// Max of values
	SeriesLimit int64 `json:"series_limit,omitempty"`
	// To timestamp
	To time.Time `json:"to"`
}
