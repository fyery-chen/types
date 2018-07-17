package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapper"
	"github.com/rancher/types/apis/cloud.huawei.com/v3"
	"github.com/rancher/types/factory"
)

var (
	Version = types.APIVersion{
		Version: "v3",
		Group:   "cloud.huawei.com",
		Path:    "/v3",
	}

	Schemas = factory.Schemas(&Version).
		Init(authzTypes).
		Init(schemaTypes).
		Init(businessTypes)
)

func businessTypes(schema *types.Schemas) *types.Schemas {
	return schema.
		AddMapperForType(&Version, v3.Business{},
			m.DisplayName{}).
		MustImport(&Version, v3.Business{}).
		MustImport(&Version, v3.BusinessQuotaCheck{}).
		MustImport(&Version, v3.BusinessQuotaCheckOutput{}).
		MustImportAndCustomize(&Version, v3.Business{}, func(schema *types.Schema) {
			schema.ResourceActions = map[string]types.Action{
				"checkout": {
					Input: "businessQuotaCheck",
					Output: "businessQuotaCheckOutput",
				},
			}
		})
}

func schemaTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.DynamicSchema{})
}

func authzTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.BusinessGlobalRole{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.BusinessRoleTemplate{}, m.DisplayName{}).
		MustImport(&Version, v3.BusinessGlobalRole{}).
		MustImport(&Version, v3.BusinessGlobalRoleBinding{}).
		MustImport(&Version, v3.BusinessRoleTemplate{}).
		MustImport(&Version, v3.BusinessRoleTemplateBinding{}).
		MustImport(&Version, v3.BusinessGlobalRoleBinding{})
}
