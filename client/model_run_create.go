/*
 * Qase.io API
 *
 * You can use our API to access Qase.io API endpoints, which allows to retrieve information about entities stored in database and perform actions with them. The API is organized around REST.
 *
 * API version: 1.0.0
 * Contact: support@qase.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type RunCreate struct {
	Title           string   `json:"title"`
	Description     string   `json:"description,omitempty"`
	IncludeAllCases bool     `json:"include_all_cases,omitempty"`
	Cases           []int32  `json:"cases,omitempty"`
	EnvironmentId   int32    `json:"environment_id,omitempty"`
	MilestoneId     int32    `json:"milestone_id,omitempty"`
	PlanId          int32    `json:"plan_id,omitempty"`
	Tags            []string `json:"tags,omitempty"`
}