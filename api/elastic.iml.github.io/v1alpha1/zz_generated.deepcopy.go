//go:build !ignore_autogenerated

/*
Copyright 2022.

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
	"github.com/kidlj/dlrover-operator/pkg/common/api/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticJob) DeepCopyInto(out *ElasticJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticJob.
func (in *ElasticJob) DeepCopy() *ElasticJob {
	if in == nil {
		return nil
	}
	out := new(ElasticJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticJobList) DeepCopyInto(out *ElasticJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticJobList.
func (in *ElasticJobList) DeepCopy() *ElasticJobList {
	if in == nil {
		return nil
	}
	out := new(ElasticJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticJobSpec) DeepCopyInto(out *ElasticJobSpec) {
	*out = *in
	if in.ResourceLimits != nil {
		in, out := &in.ResourceLimits, &out.ResourceLimits
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ReplicaSpecs != nil {
		in, out := &in.ReplicaSpecs, &out.ReplicaSpecs
		*out = make(map[v1.ReplicaType]*ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make(map[string]*corev1.EnvVar, len(*in))
		for key, val := range *in {
			var outVal *corev1.EnvVar
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(corev1.EnvVar)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticJobSpec.
func (in *ElasticJobSpec) DeepCopy() *ElasticJobSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticJobStatus) DeepCopyInto(out *ElasticJobStatus) {
	*out = *in
	in.JobStatus.DeepCopyInto(&out.JobStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticJobStatus.
func (in *ElasticJobStatus) DeepCopy() *ElasticJobStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodMeta) DeepCopyInto(out *PodMeta) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodMeta.
func (in *PodMeta) DeepCopy() *PodMeta {
	if in == nil {
		return nil
	}
	out := new(PodMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaResourceSpec) DeepCopyInto(out *ReplicaResourceSpec) {
	*out = *in
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaResourceSpec.
func (in *ReplicaResourceSpec) DeepCopy() *ReplicaResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicaResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaSpec) DeepCopyInto(out *ReplicaSpec) {
	*out = *in
	in.ReplicaSpec.DeepCopyInto(&out.ReplicaSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaSpec.
func (in *ReplicaSpec) DeepCopy() *ReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalePlan) DeepCopyInto(out *ScalePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalePlan.
func (in *ScalePlan) DeepCopy() *ScalePlan {
	if in == nil {
		return nil
	}
	out := new(ScalePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalePlanList) DeepCopyInto(out *ScalePlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalePlanList.
func (in *ScalePlanList) DeepCopy() *ScalePlanList {
	if in == nil {
		return nil
	}
	out := new(ScalePlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalePlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalePlanStatus) DeepCopyInto(out *ScalePlanStatus) {
	*out = *in
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = (*in).DeepCopy()
	}
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalePlanStatus.
func (in *ScalePlanStatus) DeepCopy() *ScalePlanStatus {
	if in == nil {
		return nil
	}
	out := new(ScalePlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleSpec) DeepCopyInto(out *ScaleSpec) {
	*out = *in
	if in.ReplicaResourceSpecs != nil {
		in, out := &in.ReplicaResourceSpecs, &out.ReplicaResourceSpecs
		*out = make(map[v1.ReplicaType]ReplicaResourceSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.CreatePods != nil {
		in, out := &in.CreatePods, &out.CreatePods
		*out = make([]PodMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RemovePods != nil {
		in, out := &in.RemovePods, &out.RemovePods
		*out = make([]PodMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MigratePods != nil {
		in, out := &in.MigratePods, &out.MigratePods
		*out = make([]PodMeta, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PsHosts != nil {
		in, out := &in.PsHosts, &out.PsHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleSpec.
func (in *ScaleSpec) DeepCopy() *ScaleSpec {
	if in == nil {
		return nil
	}
	out := new(ScaleSpec)
	in.DeepCopyInto(out)
	return out
}
