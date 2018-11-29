package client

const (
	GoogleKubernetesEngineConfigType                                    = "googleKubernetesEngineConfig"
	GoogleKubernetesEngineConfigFieldClusterIpv4Cidr                    = "clusterIpv4Cidr"
	GoogleKubernetesEngineConfigFieldCredential                         = "credential"
	GoogleKubernetesEngineConfigFieldDescription                        = "description"
	GoogleKubernetesEngineConfigFieldDiskSizeGb                         = "diskSizeGb"
	GoogleKubernetesEngineConfigFieldEnableAlphaFeature                 = "enableAlphaFeature"
	GoogleKubernetesEngineConfigFieldEnableHTTPLoadBalancing            = "enableHttpLoadBalancing"
	GoogleKubernetesEngineConfigFieldEnableHorizontalPodAutoscaling     = "enableHorizontalPodAutoscaling"
	GoogleKubernetesEngineConfigFieldEnableKubernetesDashboard          = "enableKubernetesDashboard"
	GoogleKubernetesEngineConfigFieldEnableLegacyAbac                   = "enableLegacyAbac"
	GoogleKubernetesEngineConfigFieldEnableMasterAuthorizedNetwork      = "enableMasterAuthorizedNetwork"
	GoogleKubernetesEngineConfigFieldEnableNetworkPolicyConfig          = "enableNetworkPolicyConfig"
	GoogleKubernetesEngineConfigFieldEnableNodepoolAutoscaling          = "enableNodepoolAutoscaling"
	GoogleKubernetesEngineConfigFieldEnablePrivateEndpoint              = "enablePrivateEndpoint"
	GoogleKubernetesEngineConfigFieldEnablePrivateNodes                 = "enablePrivateNodes"
	GoogleKubernetesEngineConfigFieldEnableStackdriverLogging           = "enableStackdriverLogging"
	GoogleKubernetesEngineConfigFieldEnableStackdriverMonitoring        = "enableStackdriverMonitoring"
	GoogleKubernetesEngineConfigFieldIPPolicyClusterIPV4CidrBlock       = "ipPolicyClusterIpv4CidrBlock"
	GoogleKubernetesEngineConfigFieldIPPolicyClusterSecondaryRangeName  = "ipPolicyClusterSecondaryRangeName"
	GoogleKubernetesEngineConfigFieldIPPolicyCreateSubnetwork           = "ipPolicyCreateSubnetwork"
	GoogleKubernetesEngineConfigFieldIPPolicyNodeIPV4CidrBlock          = "ipPolicyNodeIpv4CidrBlock"
	GoogleKubernetesEngineConfigFieldIPPolicyServicesIPV4CidrBlock      = "ipPolicyServicesIpv4CidrBlock"
	GoogleKubernetesEngineConfigFieldIPPolicyServicesSecondaryRangeName = "ipPolicyServicesSecondaryRangeName"
	GoogleKubernetesEngineConfigFieldIPPolicySubnetworkName             = "ipPolicySubnetworkName"
	GoogleKubernetesEngineConfigFieldImageType                          = "imageType"
	GoogleKubernetesEngineConfigFieldIssueClientCertificate             = "issueClientCertificate"
	GoogleKubernetesEngineConfigFieldLabels                             = "labels"
	GoogleKubernetesEngineConfigFieldLocations                          = "locations"
	GoogleKubernetesEngineConfigFieldMachineType                        = "machineType"
	GoogleKubernetesEngineConfigFieldMaintenanceWindow                  = "maintenanceWindow"
	GoogleKubernetesEngineConfigFieldMasterAuthPassword                 = "masterAuthPassword"
	GoogleKubernetesEngineConfigFieldMasterAuthUsername                 = "masterAuthUsername"
	GoogleKubernetesEngineConfigFieldMasterAuthorizedNetworkCidrBlocks  = "masterAuthorizedNetworkCidrBlocks"
	GoogleKubernetesEngineConfigFieldMasterIPV4CidrBlock                = "masterIpv4CidrBlock"
	GoogleKubernetesEngineConfigFieldMasterVersion                      = "masterVersion"
	GoogleKubernetesEngineConfigFieldMaxNodeCount                       = "maxNodeCount"
	GoogleKubernetesEngineConfigFieldMinNodeCount                       = "minNodeCount"
	GoogleKubernetesEngineConfigFieldNetwork                            = "network"
	GoogleKubernetesEngineConfigFieldNodeCount                          = "nodeCount"
	GoogleKubernetesEngineConfigFieldNodeVersion                        = "nodeVersion"
	GoogleKubernetesEngineConfigFieldProjectID                          = "projectId"
	GoogleKubernetesEngineConfigFieldResourceLabels                     = "resourceLabels"
	GoogleKubernetesEngineConfigFieldSubNetwork                         = "subNetwork"
	GoogleKubernetesEngineConfigFieldUseIPAliases                       = "useIpAliases"
	GoogleKubernetesEngineConfigFieldZone                               = "zone"
)

type GoogleKubernetesEngineConfig struct {
	ClusterIpv4Cidr                    string            `json:"clusterIpv4Cidr,omitempty" yaml:"clusterIpv4Cidr,omitempty"`
	Credential                         string            `json:"credential,omitempty" yaml:"credential,omitempty"`
	Description                        string            `json:"description,omitempty" yaml:"description,omitempty"`
	DiskSizeGb                         int64             `json:"diskSizeGb,omitempty" yaml:"diskSizeGb,omitempty"`
	EnableAlphaFeature                 bool              `json:"enableAlphaFeature,omitempty" yaml:"enableAlphaFeature,omitempty"`
	EnableHTTPLoadBalancing            *bool             `json:"enableHttpLoadBalancing,omitempty" yaml:"enableHttpLoadBalancing,omitempty"`
	EnableHorizontalPodAutoscaling     *bool             `json:"enableHorizontalPodAutoscaling,omitempty" yaml:"enableHorizontalPodAutoscaling,omitempty"`
	EnableKubernetesDashboard          bool              `json:"enableKubernetesDashboard,omitempty" yaml:"enableKubernetesDashboard,omitempty"`
	EnableLegacyAbac                   bool              `json:"enableLegacyAbac,omitempty" yaml:"enableLegacyAbac,omitempty"`
	EnableMasterAuthorizedNetwork      bool              `json:"enableMasterAuthorizedNetwork,omitempty" yaml:"enableMasterAuthorizedNetwork,omitempty"`
	EnableNetworkPolicyConfig          *bool             `json:"enableNetworkPolicyConfig,omitempty" yaml:"enableNetworkPolicyConfig,omitempty"`
	EnableNodepoolAutoscaling          bool              `json:"enableNodepoolAutoscaling,omitempty" yaml:"enableNodepoolAutoscaling,omitempty"`
	EnablePrivateEndpoint              bool              `json:"enablePrivateEndpoint,omitempty" yaml:"enablePrivateEndpoint,omitempty"`
	EnablePrivateNodes                 bool              `json:"enablePrivateNodes,omitempty" yaml:"enablePrivateNodes,omitempty"`
	EnableStackdriverLogging           *bool             `json:"enableStackdriverLogging,omitempty" yaml:"enableStackdriverLogging,omitempty"`
	EnableStackdriverMonitoring        *bool             `json:"enableStackdriverMonitoring,omitempty" yaml:"enableStackdriverMonitoring,omitempty"`
	IPPolicyClusterIPV4CidrBlock       string            `json:"ipPolicyClusterIpv4CidrBlock,omitempty" yaml:"ipPolicyClusterIpv4CidrBlock,omitempty"`
	IPPolicyClusterSecondaryRangeName  string            `json:"ipPolicyClusterSecondaryRangeName,omitempty" yaml:"ipPolicyClusterSecondaryRangeName,omitempty"`
	IPPolicyCreateSubnetwork           bool              `json:"ipPolicyCreateSubnetwork,omitempty" yaml:"ipPolicyCreateSubnetwork,omitempty"`
	IPPolicyNodeIPV4CidrBlock          string            `json:"ipPolicyNodeIpv4CidrBlock,omitempty" yaml:"ipPolicyNodeIpv4CidrBlock,omitempty"`
	IPPolicyServicesIPV4CidrBlock      string            `json:"ipPolicyServicesIpv4CidrBlock,omitempty" yaml:"ipPolicyServicesIpv4CidrBlock,omitempty"`
	IPPolicyServicesSecondaryRangeName string            `json:"ipPolicyServicesSecondaryRangeName,omitempty" yaml:"ipPolicyServicesSecondaryRangeName,omitempty"`
	IPPolicySubnetworkName             string            `json:"ipPolicySubnetworkName,omitempty" yaml:"ipPolicySubnetworkName,omitempty"`
	ImageType                          string            `json:"imageType,omitempty" yaml:"imageType,omitempty"`
	IssueClientCertificate             bool              `json:"issueClientCertificate,omitempty" yaml:"issueClientCertificate,omitempty"`
	Labels                             map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Locations                          []string          `json:"locations,omitempty" yaml:"locations,omitempty"`
	MachineType                        string            `json:"machineType,omitempty" yaml:"machineType,omitempty"`
	MaintenanceWindow                  string            `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`
	MasterAuthPassword                 string            `json:"masterAuthPassword,omitempty" yaml:"masterAuthPassword,omitempty"`
	MasterAuthUsername                 string            `json:"masterAuthUsername,omitempty" yaml:"masterAuthUsername,omitempty"`
	MasterAuthorizedNetworkCidrBlocks  []string          `json:"masterAuthorizedNetworkCidrBlocks,omitempty" yaml:"masterAuthorizedNetworkCidrBlocks,omitempty"`
	MasterIPV4CidrBlock                string            `json:"masterIpv4CidrBlock,omitempty" yaml:"masterIpv4CidrBlock,omitempty"`
	MasterVersion                      string            `json:"masterVersion,omitempty" yaml:"masterVersion,omitempty"`
	MaxNodeCount                       int64             `json:"maxNodeCount,omitempty" yaml:"maxNodeCount,omitempty"`
	MinNodeCount                       int64             `json:"minNodeCount,omitempty" yaml:"minNodeCount,omitempty"`
	Network                            string            `json:"network,omitempty" yaml:"network,omitempty"`
	NodeCount                          int64             `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
	NodeVersion                        string            `json:"nodeVersion,omitempty" yaml:"nodeVersion,omitempty"`
	ProjectID                          string            `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	ResourceLabels                     map[string]string `json:"resourceLabels,omitempty" yaml:"resourceLabels,omitempty"`
	SubNetwork                         string            `json:"subNetwork,omitempty" yaml:"subNetwork,omitempty"`
	UseIPAliases                       bool              `json:"useIpAliases,omitempty" yaml:"useIpAliases,omitempty"`
	Zone                               string            `json:"zone,omitempty" yaml:"zone,omitempty"`
}
