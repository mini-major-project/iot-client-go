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

// ArduinoDashboard media type (default view)
type ArduinoDashboard struct {
	// The configuration of the dashboard
	Configuration map[string]interface{} `json:"configuration"`
	// The created at of the dashboard
	CreatedAt time.Time `json:"created_at,omitempty"`
	// The id of the dashboard
	Id string `json:"id,omitempty"`
	// The friendly name of the property
	Name string `json:"name"`
	// The type of the dashboard
	Type string `json:"type"`
	// The updated at of the dashboard
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The owner of the dashboard
	UserId string `json:"user_id,omitempty"`
}
