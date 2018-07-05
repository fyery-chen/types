package client

const (
	BusinessQuotaSpecType             = "businessQuotaSpec"
	BusinessQuotaSpecFieldBusinessId  = "businessId"
	BusinessQuotaSpecFieldDescription = "description"
	BusinessQuotaSpecFieldDisplayName = "displayName"
	BusinessQuotaSpecFieldNodeCount   = "nodeCount"
)

type BusinessQuotaSpec struct {
	BusinessId  string `json:"businessId,omitempty" yaml:"businessId,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	NodeCount   int64  `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
}
