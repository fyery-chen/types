package v3

import (
	reflect "reflect"

	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	version "k8s.io/apimachinery/pkg/version"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Action).DeepCopyInto(out.(*Action))
			return nil
		}, InType: reflect.TypeOf(&Action{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActiveDirectoryConfig).DeepCopyInto(out.(*ActiveDirectoryConfig))
			return nil
		}, InType: reflect.TypeOf(&ActiveDirectoryConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ActiveDirectoryTestAndApplyInput).DeepCopyInto(out.(*ActiveDirectoryTestAndApplyInput))
			return nil
		}, InType: reflect.TypeOf(&ActiveDirectoryTestAndApplyInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AmazonElasticContainerServiceConfig).DeepCopyInto(out.(*AmazonElasticContainerServiceConfig))
			return nil
		}, InType: reflect.TypeOf(&AmazonElasticContainerServiceConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthConfig).DeepCopyInto(out.(*AuthConfig))
			return nil
		}, InType: reflect.TypeOf(&AuthConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthConfigList).DeepCopyInto(out.(*AuthConfigList))
			return nil
		}, InType: reflect.TypeOf(&AuthConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AzureKubernetesServiceConfig).DeepCopyInto(out.(*AzureKubernetesServiceConfig))
			return nil
		}, InType: reflect.TypeOf(&AzureKubernetesServiceConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ChangePasswordInput).DeepCopyInto(out.(*ChangePasswordInput))
			return nil
		}, InType: reflect.TypeOf(&ChangePasswordInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Cluster).DeepCopyInto(out.(*Cluster))
			return nil
		}, InType: reflect.TypeOf(&Cluster{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterComponentStatus).DeepCopyInto(out.(*ClusterComponentStatus))
			return nil
		}, InType: reflect.TypeOf(&ClusterComponentStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterCondition).DeepCopyInto(out.(*ClusterCondition))
			return nil
		}, InType: reflect.TypeOf(&ClusterCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterEvent).DeepCopyInto(out.(*ClusterEvent))
			return nil
		}, InType: reflect.TypeOf(&ClusterEvent{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterEventList).DeepCopyInto(out.(*ClusterEventList))
			return nil
		}, InType: reflect.TypeOf(&ClusterEventList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterList).DeepCopyInto(out.(*ClusterList))
			return nil
		}, InType: reflect.TypeOf(&ClusterList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRegistrationToken).DeepCopyInto(out.(*ClusterRegistrationToken))
			return nil
		}, InType: reflect.TypeOf(&ClusterRegistrationToken{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRegistrationTokenList).DeepCopyInto(out.(*ClusterRegistrationTokenList))
			return nil
		}, InType: reflect.TypeOf(&ClusterRegistrationTokenList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRegistrationTokenSpec).DeepCopyInto(out.(*ClusterRegistrationTokenSpec))
			return nil
		}, InType: reflect.TypeOf(&ClusterRegistrationTokenSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterRegistrationTokenStatus).DeepCopyInto(out.(*ClusterRegistrationTokenStatus))
			return nil
		}, InType: reflect.TypeOf(&ClusterRegistrationTokenStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterSpec).DeepCopyInto(out.(*ClusterSpec))
			return nil
		}, InType: reflect.TypeOf(&ClusterSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ClusterStatus).DeepCopyInto(out.(*ClusterStatus))
			return nil
		}, InType: reflect.TypeOf(&ClusterStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DynamicSchema).DeepCopyInto(out.(*DynamicSchema))
			return nil
		}, InType: reflect.TypeOf(&DynamicSchema{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DynamicSchemaList).DeepCopyInto(out.(*DynamicSchemaList))
			return nil
		}, InType: reflect.TypeOf(&DynamicSchemaList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DynamicSchemaSpec).DeepCopyInto(out.(*DynamicSchemaSpec))
			return nil
		}, InType: reflect.TypeOf(&DynamicSchemaSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DynamicSchemaStatus).DeepCopyInto(out.(*DynamicSchemaStatus))
			return nil
		}, InType: reflect.TypeOf(&DynamicSchemaStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ExportOutput).DeepCopyInto(out.(*ExportOutput))
			return nil
		}, InType: reflect.TypeOf(&ExportOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Field).DeepCopyInto(out.(*Field))
			return nil
		}, InType: reflect.TypeOf(&Field{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Filter).DeepCopyInto(out.(*Filter))
			return nil
		}, InType: reflect.TypeOf(&Filter{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GenerateKubeConfigOutput).DeepCopyInto(out.(*GenerateKubeConfigOutput))
			return nil
		}, InType: reflect.TypeOf(&GenerateKubeConfigOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubConfig).DeepCopyInto(out.(*GithubConfig))
			return nil
		}, InType: reflect.TypeOf(&GithubConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubConfigApplyInput).DeepCopyInto(out.(*GithubConfigApplyInput))
			return nil
		}, InType: reflect.TypeOf(&GithubConfigApplyInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GithubConfigTestOutput).DeepCopyInto(out.(*GithubConfigTestOutput))
			return nil
		}, InType: reflect.TypeOf(&GithubConfigTestOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GoogleKubernetesEngineConfig).DeepCopyInto(out.(*GoogleKubernetesEngineConfig))
			return nil
		}, InType: reflect.TypeOf(&GoogleKubernetesEngineConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Group).DeepCopyInto(out.(*Group))
			return nil
		}, InType: reflect.TypeOf(&Group{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GroupList).DeepCopyInto(out.(*GroupList))
			return nil
		}, InType: reflect.TypeOf(&GroupList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GroupMember).DeepCopyInto(out.(*GroupMember))
			return nil
		}, InType: reflect.TypeOf(&GroupMember{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*GroupMemberList).DeepCopyInto(out.(*GroupMemberList))
			return nil
		}, InType: reflect.TypeOf(&GroupMemberList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImportClusterYamlInput).DeepCopyInto(out.(*ImportClusterYamlInput))
			return nil
		}, InType: reflect.TypeOf(&ImportClusterYamlInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImportYamlOutput).DeepCopyInto(out.(*ImportYamlOutput))
			return nil
		}, InType: reflect.TypeOf(&ImportYamlOutput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ImportedConfig).DeepCopyInto(out.(*ImportedConfig))
			return nil
		}, InType: reflect.TypeOf(&ImportedConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ListOpts).DeepCopyInto(out.(*ListOpts))
			return nil
		}, InType: reflect.TypeOf(&ListOpts{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ListenConfig).DeepCopyInto(out.(*ListenConfig))
			return nil
		}, InType: reflect.TypeOf(&ListenConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ListenConfigList).DeepCopyInto(out.(*ListenConfigList))
			return nil
		}, InType: reflect.TypeOf(&ListenConfigList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*LocalConfig).DeepCopyInto(out.(*LocalConfig))
			return nil
		}, InType: reflect.TypeOf(&LocalConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Preference).DeepCopyInto(out.(*Preference))
			return nil
		}, InType: reflect.TypeOf(&Preference{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PreferenceList).DeepCopyInto(out.(*PreferenceList))
			return nil
		}, InType: reflect.TypeOf(&PreferenceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Principal).DeepCopyInto(out.(*Principal))
			return nil
		}, InType: reflect.TypeOf(&Principal{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PrincipalList).DeepCopyInto(out.(*PrincipalList))
			return nil
		}, InType: reflect.TypeOf(&PrincipalList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SearchPrincipalsInput).DeepCopyInto(out.(*SearchPrincipalsInput))
			return nil
		}, InType: reflect.TypeOf(&SearchPrincipalsInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SetPasswordInput).DeepCopyInto(out.(*SetPasswordInput))
			return nil
		}, InType: reflect.TypeOf(&SetPasswordInput{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Setting).DeepCopyInto(out.(*Setting))
			return nil
		}, InType: reflect.TypeOf(&Setting{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*SettingList).DeepCopyInto(out.(*SettingList))
			return nil
		}, InType: reflect.TypeOf(&SettingList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Token).DeepCopyInto(out.(*Token))
			return nil
		}, InType: reflect.TypeOf(&Token{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TokenList).DeepCopyInto(out.(*TokenList))
			return nil
		}, InType: reflect.TypeOf(&TokenList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*User).DeepCopyInto(out.(*User))
			return nil
		}, InType: reflect.TypeOf(&User{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*UserList).DeepCopyInto(out.(*UserList))
			return nil
		}, InType: reflect.TypeOf(&UserList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Values).DeepCopyInto(out.(*Values))
			return nil
		}, InType: reflect.TypeOf(&Values{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryConfig) DeepCopyInto(out *ActiveDirectoryConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthConfig.DeepCopyInto(&out.AuthConfig)
	if in.Servers != nil {
		in, out := &in.Servers, &out.Servers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryConfig.
func (in *ActiveDirectoryConfig) DeepCopy() *ActiveDirectoryConfig {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveDirectoryConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveDirectoryTestAndApplyInput) DeepCopyInto(out *ActiveDirectoryTestAndApplyInput) {
	*out = *in
	in.ActiveDirectoryConfig.DeepCopyInto(&out.ActiveDirectoryConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveDirectoryTestAndApplyInput.
func (in *ActiveDirectoryTestAndApplyInput) DeepCopy() *ActiveDirectoryTestAndApplyInput {
	if in == nil {
		return nil
	}
	out := new(ActiveDirectoryTestAndApplyInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmazonElasticContainerServiceConfig) DeepCopyInto(out *AmazonElasticContainerServiceConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmazonElasticContainerServiceConfig.
func (in *AmazonElasticContainerServiceConfig) DeepCopy() *AmazonElasticContainerServiceConfig {
	if in == nil {
		return nil
	}
	out := new(AmazonElasticContainerServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfig) DeepCopyInto(out *AuthConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AllowedPrincipalIDs != nil {
		in, out := &in.AllowedPrincipalIDs, &out.AllowedPrincipalIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfig.
func (in *AuthConfig) DeepCopy() *AuthConfig {
	if in == nil {
		return nil
	}
	out := new(AuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthConfigList) DeepCopyInto(out *AuthConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AuthConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthConfigList.
func (in *AuthConfigList) DeepCopy() *AuthConfigList {
	if in == nil {
		return nil
	}
	out := new(AuthConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AuthConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureKubernetesServiceConfig) DeepCopyInto(out *AzureKubernetesServiceConfig) {
	*out = *in
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureKubernetesServiceConfig.
func (in *AzureKubernetesServiceConfig) DeepCopy() *AzureKubernetesServiceConfig {
	if in == nil {
		return nil
	}
	out := new(AzureKubernetesServiceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChangePasswordInput) DeepCopyInto(out *ChangePasswordInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChangePasswordInput.
func (in *ChangePasswordInput) DeepCopy() *ChangePasswordInput {
	if in == nil {
		return nil
	}
	out := new(ChangePasswordInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComponentStatus) DeepCopyInto(out *ClusterComponentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.ComponentCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComponentStatus.
func (in *ClusterComponentStatus) DeepCopy() *ClusterComponentStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEvent) DeepCopyInto(out *ClusterEvent) {
	*out = *in
	out.Namespaced = in.Namespaced
	in.Event.DeepCopyInto(&out.Event)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEvent.
func (in *ClusterEvent) DeepCopy() *ClusterEvent {
	if in == nil {
		return nil
	}
	out := new(ClusterEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEvent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterEventList) DeepCopyInto(out *ClusterEventList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterEvent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterEventList.
func (in *ClusterEventList) DeepCopy() *ClusterEventList {
	if in == nil {
		return nil
	}
	out := new(ClusterEventList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterEventList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRegistrationToken) DeepCopyInto(out *ClusterRegistrationToken) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRegistrationToken.
func (in *ClusterRegistrationToken) DeepCopy() *ClusterRegistrationToken {
	if in == nil {
		return nil
	}
	out := new(ClusterRegistrationToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRegistrationToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRegistrationTokenList) DeepCopyInto(out *ClusterRegistrationTokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterRegistrationToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRegistrationTokenList.
func (in *ClusterRegistrationTokenList) DeepCopy() *ClusterRegistrationTokenList {
	if in == nil {
		return nil
	}
	out := new(ClusterRegistrationTokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterRegistrationTokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRegistrationTokenSpec) DeepCopyInto(out *ClusterRegistrationTokenSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRegistrationTokenSpec.
func (in *ClusterRegistrationTokenSpec) DeepCopy() *ClusterRegistrationTokenSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterRegistrationTokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterRegistrationTokenStatus) DeepCopyInto(out *ClusterRegistrationTokenStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterRegistrationTokenStatus.
func (in *ClusterRegistrationTokenStatus) DeepCopy() *ClusterRegistrationTokenStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterRegistrationTokenStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.ImportedConfig != nil {
		in, out := &in.ImportedConfig, &out.ImportedConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(ImportedConfig)
			**out = **in
		}
	}
	if in.GoogleKubernetesEngineConfig != nil {
		in, out := &in.GoogleKubernetesEngineConfig, &out.GoogleKubernetesEngineConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(GoogleKubernetesEngineConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AzureKubernetesServiceConfig != nil {
		in, out := &in.AzureKubernetesServiceConfig, &out.AzureKubernetesServiceConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(AzureKubernetesServiceConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.AmazonElasticContainerServiceConfig != nil {
		in, out := &in.AmazonElasticContainerServiceConfig, &out.AmazonElasticContainerServiceConfig
		if *in == nil {
			*out = nil
		} else {
			*out = new(AmazonElasticContainerServiceConfig)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		copy(*out, *in)
	}
	if in.ComponentStatuses != nil {
		in, out := &in.ComponentStatuses, &out.ComponentStatuses
		*out = make([]ClusterComponentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Capacity != nil {
		in, out := &in.Capacity, &out.Capacity
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Allocatable != nil {
		in, out := &in.Allocatable, &out.Allocatable
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	in.AppliedSpec.DeepCopyInto(&out.AppliedSpec)
	if in.FailedSpec != nil {
		in, out := &in.FailedSpec, &out.FailedSpec
		if *in == nil {
			*out = nil
		} else {
			*out = new(ClusterSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Requested != nil {
		in, out := &in.Requested, &out.Requested
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		if *in == nil {
			*out = nil
		} else {
			*out = new(version.Info)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicSchema) DeepCopyInto(out *DynamicSchema) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicSchema.
func (in *DynamicSchema) DeepCopy() *DynamicSchema {
	if in == nil {
		return nil
	}
	out := new(DynamicSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicSchema) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicSchemaList) DeepCopyInto(out *DynamicSchemaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DynamicSchema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicSchemaList.
func (in *DynamicSchemaList) DeepCopy() *DynamicSchemaList {
	if in == nil {
		return nil
	}
	out := new(DynamicSchemaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DynamicSchemaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicSchemaSpec) DeepCopyInto(out *DynamicSchemaSpec) {
	*out = *in
	if in.ResourceMethods != nil {
		in, out := &in.ResourceMethods, &out.ResourceMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceFields != nil {
		in, out := &in.ResourceFields, &out.ResourceFields
		*out = make(map[string]Field, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ResourceActions != nil {
		in, out := &in.ResourceActions, &out.ResourceActions
		*out = make(map[string]Action, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CollectionMethods != nil {
		in, out := &in.CollectionMethods, &out.CollectionMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CollectionFields != nil {
		in, out := &in.CollectionFields, &out.CollectionFields
		*out = make(map[string]Field, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.CollectionActions != nil {
		in, out := &in.CollectionActions, &out.CollectionActions
		*out = make(map[string]Action, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CollectionFilters != nil {
		in, out := &in.CollectionFilters, &out.CollectionFilters
		*out = make(map[string]Filter, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.IncludeableLinks != nil {
		in, out := &in.IncludeableLinks, &out.IncludeableLinks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicSchemaSpec.
func (in *DynamicSchemaSpec) DeepCopy() *DynamicSchemaSpec {
	if in == nil {
		return nil
	}
	out := new(DynamicSchemaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicSchemaStatus) DeepCopyInto(out *DynamicSchemaStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicSchemaStatus.
func (in *DynamicSchemaStatus) DeepCopy() *DynamicSchemaStatus {
	if in == nil {
		return nil
	}
	out := new(DynamicSchemaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportOutput) DeepCopyInto(out *ExportOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportOutput.
func (in *ExportOutput) DeepCopy() *ExportOutput {
	if in == nil {
		return nil
	}
	out := new(ExportOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Field) DeepCopyInto(out *Field) {
	*out = *in
	in.Default.DeepCopyInto(&out.Default)
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Field.
func (in *Field) DeepCopy() *Field {
	if in == nil {
		return nil
	}
	out := new(Field)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
	if in.Modifiers != nil {
		in, out := &in.Modifiers, &out.Modifiers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenerateKubeConfigOutput) DeepCopyInto(out *GenerateKubeConfigOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenerateKubeConfigOutput.
func (in *GenerateKubeConfigOutput) DeepCopy() *GenerateKubeConfigOutput {
	if in == nil {
		return nil
	}
	out := new(GenerateKubeConfigOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubConfig) DeepCopyInto(out *GithubConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthConfig.DeepCopyInto(&out.AuthConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubConfig.
func (in *GithubConfig) DeepCopy() *GithubConfig {
	if in == nil {
		return nil
	}
	out := new(GithubConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GithubConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubConfigApplyInput) DeepCopyInto(out *GithubConfigApplyInput) {
	*out = *in
	in.GithubConfig.DeepCopyInto(&out.GithubConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubConfigApplyInput.
func (in *GithubConfigApplyInput) DeepCopy() *GithubConfigApplyInput {
	if in == nil {
		return nil
	}
	out := new(GithubConfigApplyInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubConfigTestOutput) DeepCopyInto(out *GithubConfigTestOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubConfigTestOutput.
func (in *GithubConfigTestOutput) DeepCopy() *GithubConfigTestOutput {
	if in == nil {
		return nil
	}
	out := new(GithubConfigTestOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleKubernetesEngineConfig) DeepCopyInto(out *GoogleKubernetesEngineConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleKubernetesEngineConfig.
func (in *GoogleKubernetesEngineConfig) DeepCopy() *GoogleKubernetesEngineConfig {
	if in == nil {
		return nil
	}
	out := new(GoogleKubernetesEngineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Group) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupList) DeepCopyInto(out *GroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Group, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupList.
func (in *GroupList) DeepCopy() *GroupList {
	if in == nil {
		return nil
	}
	out := new(GroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMember) DeepCopyInto(out *GroupMember) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMember.
func (in *GroupMember) DeepCopy() *GroupMember {
	if in == nil {
		return nil
	}
	out := new(GroupMember)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupMember) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupMemberList) DeepCopyInto(out *GroupMemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GroupMember, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupMemberList.
func (in *GroupMemberList) DeepCopy() *GroupMemberList {
	if in == nil {
		return nil
	}
	out := new(GroupMemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupMemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportClusterYamlInput) DeepCopyInto(out *ImportClusterYamlInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportClusterYamlInput.
func (in *ImportClusterYamlInput) DeepCopy() *ImportClusterYamlInput {
	if in == nil {
		return nil
	}
	out := new(ImportClusterYamlInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportYamlOutput) DeepCopyInto(out *ImportYamlOutput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportYamlOutput.
func (in *ImportYamlOutput) DeepCopy() *ImportYamlOutput {
	if in == nil {
		return nil
	}
	out := new(ImportYamlOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedConfig) DeepCopyInto(out *ImportedConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedConfig.
func (in *ImportedConfig) DeepCopy() *ImportedConfig {
	if in == nil {
		return nil
	}
	out := new(ImportedConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListOpts) DeepCopyInto(out *ListOpts) {
	*out = *in
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListOpts.
func (in *ListOpts) DeepCopy() *ListOpts {
	if in == nil {
		return nil
	}
	out := new(ListOpts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenConfig) DeepCopyInto(out *ListenConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TOS != nil {
		in, out := &in.TOS, &out.TOS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KnownIPs != nil {
		in, out := &in.KnownIPs, &out.KnownIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GeneratedCerts != nil {
		in, out := &in.GeneratedCerts, &out.GeneratedCerts
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SubjectAlternativeNames != nil {
		in, out := &in.SubjectAlternativeNames, &out.SubjectAlternativeNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenConfig.
func (in *ListenConfig) DeepCopy() *ListenConfig {
	if in == nil {
		return nil
	}
	out := new(ListenConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListenConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenConfigList) DeepCopyInto(out *ListenConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ListenConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenConfigList.
func (in *ListenConfigList) DeepCopy() *ListenConfigList {
	if in == nil {
		return nil
	}
	out := new(ListenConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ListenConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalConfig) DeepCopyInto(out *LocalConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AuthConfig.DeepCopyInto(&out.AuthConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalConfig.
func (in *LocalConfig) DeepCopy() *LocalConfig {
	if in == nil {
		return nil
	}
	out := new(LocalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preference) DeepCopyInto(out *Preference) {
	*out = *in
	out.Namespaced = in.Namespaced
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preference.
func (in *Preference) DeepCopy() *Preference {
	if in == nil {
		return nil
	}
	out := new(Preference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Preference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreferenceList) DeepCopyInto(out *PreferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Preference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreferenceList.
func (in *PreferenceList) DeepCopy() *PreferenceList {
	if in == nil {
		return nil
	}
	out := new(PreferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Principal) DeepCopyInto(out *Principal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.ExtraInfo != nil {
		in, out := &in.ExtraInfo, &out.ExtraInfo
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Principal.
func (in *Principal) DeepCopy() *Principal {
	if in == nil {
		return nil
	}
	out := new(Principal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Principal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalList) DeepCopyInto(out *PrincipalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Principal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalList.
func (in *PrincipalList) DeepCopy() *PrincipalList {
	if in == nil {
		return nil
	}
	out := new(PrincipalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrincipalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SearchPrincipalsInput) DeepCopyInto(out *SearchPrincipalsInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SearchPrincipalsInput.
func (in *SearchPrincipalsInput) DeepCopy() *SearchPrincipalsInput {
	if in == nil {
		return nil
	}
	out := new(SearchPrincipalsInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SetPasswordInput) DeepCopyInto(out *SetPasswordInput) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SetPasswordInput.
func (in *SetPasswordInput) DeepCopy() *SetPasswordInput {
	if in == nil {
		return nil
	}
	out := new(SetPasswordInput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Setting) DeepCopyInto(out *Setting) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Setting.
func (in *Setting) DeepCopy() *Setting {
	if in == nil {
		return nil
	}
	out := new(Setting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Setting) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingList) DeepCopyInto(out *SettingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Setting, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingList.
func (in *SettingList) DeepCopy() *SettingList {
	if in == nil {
		return nil
	}
	out := new(SettingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SettingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Token) DeepCopyInto(out *Token) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.UserPrincipal.DeepCopyInto(&out.UserPrincipal)
	if in.GroupPrincipals != nil {
		in, out := &in.GroupPrincipals, &out.GroupPrincipals
		*out = make([]Principal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ProviderInfo != nil {
		in, out := &in.ProviderInfo, &out.ProviderInfo
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Token.
func (in *Token) DeepCopy() *Token {
	if in == nil {
		return nil
	}
	out := new(Token)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Token) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenList) DeepCopyInto(out *TokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Token, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenList.
func (in *TokenList) DeepCopy() *TokenList {
	if in == nil {
		return nil
	}
	out := new(TokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.PrincipalIDs != nil {
		in, out := &in.PrincipalIDs, &out.PrincipalIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Values) DeepCopyInto(out *Values) {
	*out = *in
	if in.StringSliceValue != nil {
		in, out := &in.StringSliceValue, &out.StringSliceValue
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Values.
func (in *Values) DeepCopy() *Values {
	if in == nil {
		return nil
	}
	out := new(Values)
	in.DeepCopyInto(out)
	return out
}
