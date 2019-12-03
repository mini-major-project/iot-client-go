/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot
// ThingSketch ThingSketchPayload describes a sketch of a thing
type ThingSketch struct {
	// The wifi password
	Pass string `json:"pass,omitempty"`
	// The autogenerated sketch version
	SketchVersion string `json:"sketch_version,omitempty"`
	// The wifi ssid to connect
	Ssid string `json:"ssid,omitempty"`
}
