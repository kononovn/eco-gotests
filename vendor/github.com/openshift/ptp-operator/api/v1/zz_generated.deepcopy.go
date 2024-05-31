//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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

package v1

import (
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HwConfig) DeepCopyInto(out *HwConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(apiextensionsv1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HwConfig.
func (in *HwConfig) DeepCopy() *HwConfig {
	if in == nil {
		return nil
	}
	out := new(HwConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatchRule) DeepCopyInto(out *MatchRule) {
	*out = *in
	if in.NodeLabel != nil {
		in, out := &in.NodeLabel, &out.NodeLabel
		*out = new(string)
		**out = **in
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatchRule.
func (in *MatchRule) DeepCopy() *MatchRule {
	if in == nil {
		return nil
	}
	out := new(MatchRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMatchList) DeepCopyInto(out *NodeMatchList) {
	*out = *in
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMatchList.
func (in *NodeMatchList) DeepCopy() *NodeMatchList {
	if in == nil {
		return nil
	}
	out := new(NodeMatchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePtpDevice) DeepCopyInto(out *NodePtpDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePtpDevice.
func (in *NodePtpDevice) DeepCopy() *NodePtpDevice {
	if in == nil {
		return nil
	}
	out := new(NodePtpDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePtpDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePtpDeviceList) DeepCopyInto(out *NodePtpDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePtpDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePtpDeviceList.
func (in *NodePtpDeviceList) DeepCopy() *NodePtpDeviceList {
	if in == nil {
		return nil
	}
	out := new(NodePtpDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePtpDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePtpDeviceSpec) DeepCopyInto(out *NodePtpDeviceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePtpDeviceSpec.
func (in *NodePtpDeviceSpec) DeepCopy() *NodePtpDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(NodePtpDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePtpDeviceStatus) DeepCopyInto(out *NodePtpDeviceStatus) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]PtpDevice, len(*in))
		copy(*out, *in)
	}
	if in.Hwconfig != nil {
		in, out := &in.Hwconfig, &out.Hwconfig
		*out = make([]HwConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePtpDeviceStatus.
func (in *NodePtpDeviceStatus) DeepCopy() *NodePtpDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(NodePtpDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpClockThreshold) DeepCopyInto(out *PtpClockThreshold) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpClockThreshold.
func (in *PtpClockThreshold) DeepCopy() *PtpClockThreshold {
	if in == nil {
		return nil
	}
	out := new(PtpClockThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpConfig) DeepCopyInto(out *PtpConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpConfig.
func (in *PtpConfig) DeepCopy() *PtpConfig {
	if in == nil {
		return nil
	}
	out := new(PtpConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PtpConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpConfigList) DeepCopyInto(out *PtpConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PtpConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpConfigList.
func (in *PtpConfigList) DeepCopy() *PtpConfigList {
	if in == nil {
		return nil
	}
	out := new(PtpConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PtpConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpConfigSpec) DeepCopyInto(out *PtpConfigSpec) {
	*out = *in
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = make([]PtpProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Recommend != nil {
		in, out := &in.Recommend, &out.Recommend
		*out = make([]PtpRecommend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpConfigSpec.
func (in *PtpConfigSpec) DeepCopy() *PtpConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PtpConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpConfigStatus) DeepCopyInto(out *PtpConfigStatus) {
	*out = *in
	if in.MatchList != nil {
		in, out := &in.MatchList, &out.MatchList
		*out = make([]NodeMatchList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpConfigStatus.
func (in *PtpConfigStatus) DeepCopy() *PtpConfigStatus {
	if in == nil {
		return nil
	}
	out := new(PtpConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpDevice) DeepCopyInto(out *PtpDevice) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpDevice.
func (in *PtpDevice) DeepCopy() *PtpDevice {
	if in == nil {
		return nil
	}
	out := new(PtpDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpEventConfig) DeepCopyInto(out *PtpEventConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpEventConfig.
func (in *PtpEventConfig) DeepCopy() *PtpEventConfig {
	if in == nil {
		return nil
	}
	out := new(PtpEventConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpOperatorConfig) DeepCopyInto(out *PtpOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpOperatorConfig.
func (in *PtpOperatorConfig) DeepCopy() *PtpOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(PtpOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PtpOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpOperatorConfigList) DeepCopyInto(out *PtpOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PtpOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpOperatorConfigList.
func (in *PtpOperatorConfigList) DeepCopy() *PtpOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(PtpOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PtpOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpOperatorConfigSpec) DeepCopyInto(out *PtpOperatorConfigSpec) {
	*out = *in
	if in.DaemonNodeSelector != nil {
		in, out := &in.DaemonNodeSelector, &out.DaemonNodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EventConfig != nil {
		in, out := &in.EventConfig, &out.EventConfig
		*out = new(PtpEventConfig)
		**out = **in
	}
	if in.EnabledPlugins != nil {
		in, out := &in.EnabledPlugins, &out.EnabledPlugins
		*out = new(map[string]*apiextensionsv1.JSON)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]*apiextensionsv1.JSON, len(*in))
			for key, val := range *in {
				var outVal *apiextensionsv1.JSON
				if val == nil {
					(*out)[key] = nil
				} else {
					in, out := &val, &outVal
					*out = new(apiextensionsv1.JSON)
					(*in).DeepCopyInto(*out)
				}
				(*out)[key] = outVal
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpOperatorConfigSpec.
func (in *PtpOperatorConfigSpec) DeepCopy() *PtpOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PtpOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpOperatorConfigStatus) DeepCopyInto(out *PtpOperatorConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpOperatorConfigStatus.
func (in *PtpOperatorConfigStatus) DeepCopy() *PtpOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(PtpOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpProfile) DeepCopyInto(out *PtpProfile) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Ptp4lOpts != nil {
		in, out := &in.Ptp4lOpts, &out.Ptp4lOpts
		*out = new(string)
		**out = **in
	}
	if in.Phc2sysOpts != nil {
		in, out := &in.Phc2sysOpts, &out.Phc2sysOpts
		*out = new(string)
		**out = **in
	}
	if in.Ts2PhcOpts != nil {
		in, out := &in.Ts2PhcOpts, &out.Ts2PhcOpts
		*out = new(string)
		**out = **in
	}
	if in.Ptp4lConf != nil {
		in, out := &in.Ptp4lConf, &out.Ptp4lConf
		*out = new(string)
		**out = **in
	}
	if in.Phc2sysConf != nil {
		in, out := &in.Phc2sysConf, &out.Phc2sysConf
		*out = new(string)
		**out = **in
	}
	if in.Ts2PhcConf != nil {
		in, out := &in.Ts2PhcConf, &out.Ts2PhcConf
		*out = new(string)
		**out = **in
	}
	if in.PtpSchedulingPolicy != nil {
		in, out := &in.PtpSchedulingPolicy, &out.PtpSchedulingPolicy
		*out = new(string)
		**out = **in
	}
	if in.PtpSchedulingPriority != nil {
		in, out := &in.PtpSchedulingPriority, &out.PtpSchedulingPriority
		*out = new(int64)
		**out = **in
	}
	if in.PtpClockThreshold != nil {
		in, out := &in.PtpClockThreshold, &out.PtpClockThreshold
		*out = new(PtpClockThreshold)
		**out = **in
	}
	if in.PtpSettings != nil {
		in, out := &in.PtpSettings, &out.PtpSettings
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = make(map[string]*apiextensionsv1.JSON, len(*in))
		for key, val := range *in {
			var outVal *apiextensionsv1.JSON
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(apiextensionsv1.JSON)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpProfile.
func (in *PtpProfile) DeepCopy() *PtpProfile {
	if in == nil {
		return nil
	}
	out := new(PtpProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PtpRecommend) DeepCopyInto(out *PtpRecommend) {
	*out = *in
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]MatchRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PtpRecommend.
func (in *PtpRecommend) DeepCopy() *PtpRecommend {
	if in == nil {
		return nil
	}
	out := new(PtpRecommend)
	in.DeepCopyInto(out)
	return out
}
