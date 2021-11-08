/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AccountPublicKey struct {
	Index int32 `json:"index"`
	PublicKey string `json:"public_key"`
	SigningAlgorithm *SigningAlgorithm `json:"signing_algorithm"`
	HashingAlgorithm *HashingAlgorithm `json:"hashing_algorithm"`
	SequenceNumber int32 `json:"sequence_number"`
	Weight int32 `json:"weight"`
	Revoked bool `json:"revoked"`
}
