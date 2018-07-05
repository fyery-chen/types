package client

import (
	"github.com/rancher/norman/types"
)

const (
	ListenConfigBusinessType                         = "listenConfigBusiness"
	ListenConfigBusinessFieldAlgorithm               = "algorithm"
	ListenConfigBusinessFieldAnnotations             = "annotations"
	ListenConfigBusinessFieldCACerts                 = "caCerts"
	ListenConfigBusinessFieldCN                      = "cn"
	ListenConfigBusinessFieldCert                    = "cert"
	ListenConfigBusinessFieldCertFingerprint         = "certFingerprint"
	ListenConfigBusinessFieldCreated                 = "created"
	ListenConfigBusinessFieldCreatorID               = "creatorId"
	ListenConfigBusinessFieldDescription             = "description"
	ListenConfigBusinessFieldDomains                 = "domains"
	ListenConfigBusinessFieldEnabled                 = "enabled"
	ListenConfigBusinessFieldExpiresAt               = "expiresAt"
	ListenConfigBusinessFieldGeneratedCerts          = "generatedCerts"
	ListenConfigBusinessFieldIssuedAt                = "issuedAt"
	ListenConfigBusinessFieldIssuer                  = "issuer"
	ListenConfigBusinessFieldKey                     = "key"
	ListenConfigBusinessFieldKeySize                 = "keySize"
	ListenConfigBusinessFieldKnownIPs                = "knownIps"
	ListenConfigBusinessFieldLabels                  = "labels"
	ListenConfigBusinessFieldMode                    = "mode"
	ListenConfigBusinessFieldName                    = "name"
	ListenConfigBusinessFieldOwnerReferences         = "ownerReferences"
	ListenConfigBusinessFieldRemoved                 = "removed"
	ListenConfigBusinessFieldSerialNumber            = "serialNumber"
	ListenConfigBusinessFieldSubjectAlternativeNames = "subjectAlternativeNames"
	ListenConfigBusinessFieldTOS                     = "tos"
	ListenConfigBusinessFieldUuid                    = "uuid"
	ListenConfigBusinessFieldVersion                 = "version"
)

type ListenConfigBusiness struct {
	types.Resource
	Algorithm               string            `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
	Annotations             map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	CACerts                 string            `json:"caCerts,omitempty" yaml:"caCerts,omitempty"`
	CN                      string            `json:"cn,omitempty" yaml:"cn,omitempty"`
	Cert                    string            `json:"cert,omitempty" yaml:"cert,omitempty"`
	CertFingerprint         string            `json:"certFingerprint,omitempty" yaml:"certFingerprint,omitempty"`
	Created                 string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID               string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Description             string            `json:"description,omitempty" yaml:"description,omitempty"`
	Domains                 []string          `json:"domains,omitempty" yaml:"domains,omitempty"`
	Enabled                 bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ExpiresAt               string            `json:"expiresAt,omitempty" yaml:"expiresAt,omitempty"`
	GeneratedCerts          map[string]string `json:"generatedCerts,omitempty" yaml:"generatedCerts,omitempty"`
	IssuedAt                string            `json:"issuedAt,omitempty" yaml:"issuedAt,omitempty"`
	Issuer                  string            `json:"issuer,omitempty" yaml:"issuer,omitempty"`
	Key                     string            `json:"key,omitempty" yaml:"key,omitempty"`
	KeySize                 int64             `json:"keySize,omitempty" yaml:"keySize,omitempty"`
	KnownIPs                []string          `json:"knownIps,omitempty" yaml:"knownIps,omitempty"`
	Labels                  map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Mode                    string            `json:"mode,omitempty" yaml:"mode,omitempty"`
	Name                    string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences         []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed                 string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	SerialNumber            string            `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty"`
	SubjectAlternativeNames []string          `json:"subjectAlternativeNames,omitempty" yaml:"subjectAlternativeNames,omitempty"`
	TOS                     []string          `json:"tos,omitempty" yaml:"tos,omitempty"`
	Uuid                    string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Version                 int64             `json:"version,omitempty" yaml:"version,omitempty"`
}
type ListenConfigBusinessCollection struct {
	types.Collection
	Data   []ListenConfigBusiness `json:"data,omitempty"`
	client *ListenConfigBusinessClient
}

type ListenConfigBusinessClient struct {
	apiClient *Client
}

type ListenConfigBusinessOperations interface {
	List(opts *types.ListOpts) (*ListenConfigBusinessCollection, error)
	Create(opts *ListenConfigBusiness) (*ListenConfigBusiness, error)
	Update(existing *ListenConfigBusiness, updates interface{}) (*ListenConfigBusiness, error)
	ByID(id string) (*ListenConfigBusiness, error)
	Delete(container *ListenConfigBusiness) error
}

func newListenConfigBusinessClient(apiClient *Client) *ListenConfigBusinessClient {
	return &ListenConfigBusinessClient{
		apiClient: apiClient,
	}
}

func (c *ListenConfigBusinessClient) Create(container *ListenConfigBusiness) (*ListenConfigBusiness, error) {
	resp := &ListenConfigBusiness{}
	err := c.apiClient.Ops.DoCreate(ListenConfigBusinessType, container, resp)
	return resp, err
}

func (c *ListenConfigBusinessClient) Update(existing *ListenConfigBusiness, updates interface{}) (*ListenConfigBusiness, error) {
	resp := &ListenConfigBusiness{}
	err := c.apiClient.Ops.DoUpdate(ListenConfigBusinessType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ListenConfigBusinessClient) List(opts *types.ListOpts) (*ListenConfigBusinessCollection, error) {
	resp := &ListenConfigBusinessCollection{}
	err := c.apiClient.Ops.DoList(ListenConfigBusinessType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ListenConfigBusinessCollection) Next() (*ListenConfigBusinessCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ListenConfigBusinessCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ListenConfigBusinessClient) ByID(id string) (*ListenConfigBusiness, error) {
	resp := &ListenConfigBusiness{}
	err := c.apiClient.Ops.DoByID(ListenConfigBusinessType, id, resp)
	return resp, err
}

func (c *ListenConfigBusinessClient) Delete(container *ListenConfigBusiness) error {
	return c.apiClient.Ops.DoResourceDelete(ListenConfigBusinessType, &container.Resource)
}
