/*
 * operationcontainer.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package restsdk

type OperationContainerAddress struct {
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Zipcode int32  `json:"zipcode,omitempty"`
	Street  string `json:"street,omitempty"`
}
