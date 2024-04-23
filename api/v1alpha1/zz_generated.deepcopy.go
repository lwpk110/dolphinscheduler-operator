//go:build !ignore_autogenerated

/*
Copyright 2024 zncdata-labs.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlerterImageSpec) DeepCopyInto(out *AlerterImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlerterImageSpec.
func (in *AlerterImageSpec) DeepCopy() *AlerterImageSpec {
	if in == nil {
		return nil
	}
	out := new(AlerterImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlerterSpec) DeepCopyInto(out *AlerterSpec) {
	*out = *in
	out.Image = in.Image
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		**out = **in
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlerterSpec.
func (in *AlerterSpec) DeepCopy() *AlerterSpec {
	if in == nil {
		return nil
	}
	out := new(AlerterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiImageSpec) DeepCopyInto(out *ApiImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiImageSpec.
func (in *ApiImageSpec) DeepCopy() *ApiImageSpec {
	if in == nil {
		return nil
	}
	out := new(ApiImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApiSpec) DeepCopyInto(out *ApiSpec) {
	*out = *in
	out.Image = in.Image
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		**out = **in
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApiSpec.
func (in *ApiSpec) DeepCopy() *ApiSpec {
	if in == nil {
		return nil
	}
	out := new(ApiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUResource) DeepCopyInto(out *CPUResource) {
	*out = *in
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUResource.
func (in *CPUResource) DeepCopy() *CPUResource {
	if in == nil {
		return nil
	}
	out := new(CPUResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOverridesSpec) DeepCopyInto(out *ConfigOverridesSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOverridesSpec.
func (in *ConfigOverridesSpec) DeepCopy() *ConfigOverridesSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigOverridesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.ExtraEnv != nil {
		in, out := &in.ExtraEnv, &out.ExtraEnv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraSecret != nil {
		in, out := &in.ExtraSecret, &out.ExtraSecret
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(ContainerLoggingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (in *ConfigSpec) DeepCopy() *ConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerLoggingSpec) DeepCopyInto(out *ContainerLoggingSpec) {
	*out = *in
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerLoggingSpec.
func (in *ContainerLoggingSpec) DeepCopy() *ContainerLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DolphinschedulerCluster) DeepCopyInto(out *DolphinschedulerCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DolphinschedulerCluster.
func (in *DolphinschedulerCluster) DeepCopy() *DolphinschedulerCluster {
	if in == nil {
		return nil
	}
	out := new(DolphinschedulerCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DolphinschedulerCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DolphinschedulerClusterList) DeepCopyInto(out *DolphinschedulerClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DolphinschedulerCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DolphinschedulerClusterList.
func (in *DolphinschedulerClusterList) DeepCopy() *DolphinschedulerClusterList {
	if in == nil {
		return nil
	}
	out := new(DolphinschedulerClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DolphinschedulerClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DolphinschedulerClusterSpec) DeepCopyInto(out *DolphinschedulerClusterSpec) {
	*out = *in
	out.Image = in.Image
	if in.ClusterConfigSpec != nil {
		in, out := &in.ClusterConfigSpec, &out.ClusterConfigSpec
		*out = new(ClusterConfigSpec)
		**out = **in
	}
	if in.MasterSpec != nil {
		in, out := &in.MasterSpec, &out.MasterSpec
		*out = new(MasterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerSpec != nil {
		in, out := &in.WorkerSpec, &out.WorkerSpec
		*out = new(WorkerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AlerterSpec != nil {
		in, out := &in.AlerterSpec, &out.AlerterSpec
		*out = new(AlerterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Api != nil {
		in, out := &in.Api, &out.Api
		*out = new(ApiSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DolphinschedulerClusterSpec.
func (in *DolphinschedulerClusterSpec) DeepCopy() *DolphinschedulerClusterSpec {
	if in == nil {
		return nil
	}
	out := new(DolphinschedulerClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogLevelSpec) DeepCopyInto(out *LogLevelSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogLevelSpec.
func (in *LogLevelSpec) DeepCopy() *LogLevelSpec {
	if in == nil {
		return nil
	}
	out := new(LogLevelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigSpec) DeepCopyInto(out *LoggingConfigSpec) {
	*out = *in
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make(map[string]*LogLevelSpec, len(*in))
		for key, val := range *in {
			var outVal *LogLevelSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(LogLevelSpec)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Console != nil {
		in, out := &in.Console, &out.Console
		*out = new(LogLevelSpec)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(LogLevelSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigSpec.
func (in *LoggingConfigSpec) DeepCopy() *LoggingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterImageSpec) DeepCopyInto(out *MasterImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterImageSpec.
func (in *MasterImageSpec) DeepCopy() *MasterImageSpec {
	if in == nil {
		return nil
	}
	out := new(MasterImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MasterSpec) DeepCopyInto(out *MasterSpec) {
	*out = *in
	out.ImageSpec = in.ImageSpec
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		**out = **in
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MasterSpec.
func (in *MasterSpec) DeepCopy() *MasterSpec {
	if in == nil {
		return nil
	}
	out := new(MasterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryResource) DeepCopyInto(out *MemoryResource) {
	*out = *in
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryResource.
func (in *MemoryResource) DeepCopy() *MemoryResource {
	if in == nil {
		return nil
	}
	out := new(MemoryResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesSpec) DeepCopyInto(out *ResourcesSpec) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUResource)
		(*in).DeepCopyInto(*out)
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(MemoryResource)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageResource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesSpec.
func (in *ResourcesSpec) DeepCopy() *ResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleGroupSpec) DeepCopyInto(out *RoleGroupSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		**out = **in
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleGroupSpec.
func (in *RoleGroupSpec) DeepCopy() *RoleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(RoleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResource) DeepCopyInto(out *StorageResource) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResource.
func (in *StorageResource) DeepCopy() *StorageResource {
	if in == nil {
		return nil
	}
	out := new(StorageResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResourceSpec) DeepCopyInto(out *StorageResourceSpec) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(StorageResource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResourceSpec.
func (in *StorageResourceSpec) DeepCopy() *StorageResourceSpec {
	if in == nil {
		return nil
	}
	out := new(StorageResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerImageSpec) DeepCopyInto(out *WorkerImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerImageSpec.
func (in *WorkerImageSpec) DeepCopy() *WorkerImageSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerSpec) DeepCopyInto(out *WorkerSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(WorkerImageSpec)
		**out = **in
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*RoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *RoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(RoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		**out = **in
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerSpec.
func (in *WorkerSpec) DeepCopy() *WorkerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerSpec)
	in.DeepCopyInto(out)
	return out
}