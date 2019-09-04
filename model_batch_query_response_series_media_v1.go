/*
 * Iot API
 *
 * Collection of all public API endpoints.
 *
 * API version: 2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iot

type BatchQueryResponseSeriesMediaV1 struct {
	// Aggregation type
	Aggr string `json:"aggr"`
	// Query
	Expression string `json:"expression"`
	// Metric name
	Metric string `json:"metric"`
}
