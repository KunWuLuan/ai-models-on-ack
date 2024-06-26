//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeadGroupSpec) DeepCopyInto(out *HeadGroupSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeadGroupSpec.
func (in *HeadGroupSpec) DeepCopy() *HeadGroupSpec {
	if in == nil {
		return nil
	}
	out := new(HeadGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmCluster) DeepCopyInto(out *SlurmCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmCluster.
func (in *SlurmCluster) DeepCopy() *SlurmCluster {
	if in == nil {
		return nil
	}
	out := new(SlurmCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmClusterList) DeepCopyInto(out *SlurmClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlurmCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmClusterList.
func (in *SlurmClusterList) DeepCopy() *SlurmClusterList {
	if in == nil {
		return nil
	}
	out := new(SlurmClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmClusterSpec) DeepCopyInto(out *SlurmClusterSpec) {
	*out = *in
	in.HeadGroupSpec.DeepCopyInto(&out.HeadGroupSpec)
	if in.WorkerGroupSpecs != nil {
		in, out := &in.WorkerGroupSpecs, &out.WorkerGroupSpecs
		*out = make([]WorkerGroupSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmClusterSpec.
func (in *SlurmClusterSpec) DeepCopy() *SlurmClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SlurmClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmClusterStatus) DeepCopyInto(out *SlurmClusterStatus) {
	*out = *in
	if in.AvailableWorkers != nil {
		in, out := &in.AvailableWorkers, &out.AvailableWorkers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmClusterStatus.
func (in *SlurmClusterStatus) DeepCopy() *SlurmClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SlurmClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJob) DeepCopyInto(out *SlurmJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJob.
func (in *SlurmJob) DeepCopy() *SlurmJob {
	if in == nil {
		return nil
	}
	out := new(SlurmJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobList) DeepCopyInto(out *SlurmJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlurmJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobList.
func (in *SlurmJobList) DeepCopy() *SlurmJobList {
	if in == nil {
		return nil
	}
	out := new(SlurmJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobSpec) DeepCopyInto(out *SlurmJobSpec) {
	*out = *in
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.SlurmClusterSpec != nil {
		in, out := &in.SlurmClusterSpec, &out.SlurmClusterSpec
		*out = new(SlurmClusterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SlurmCluster != nil {
		in, out := &in.SlurmCluster, &out.SlurmCluster
		*out = new(corev1.ObjectReference)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobSpec.
func (in *SlurmJobSpec) DeepCopy() *SlurmJobSpec {
	if in == nil {
		return nil
	}
	out := new(SlurmJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmJobStatus) DeepCopyInto(out *SlurmJobStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.SlurmClusterStatus != nil {
		in, out := &in.SlurmClusterStatus, &out.SlurmClusterStatus
		*out = new(SlurmClusterStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmJobStatus.
func (in *SlurmJobStatus) DeepCopy() *SlurmJobStatus {
	if in == nil {
		return nil
	}
	out := new(SlurmJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubmitterTemplate) DeepCopyInto(out *SubmitterTemplate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubmitterTemplate.
func (in *SubmitterTemplate) DeepCopy() *SubmitterTemplate {
	if in == nil {
		return nil
	}
	out := new(SubmitterTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerGroupSpec) DeepCopyInto(out *WorkerGroupSpec) {
	*out = *in
	if in.PodsToBeDeleted != nil {
		in, out := &in.PodsToBeDeleted, &out.PodsToBeDeleted
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodIndexesToBeCreated != nil {
		in, out := &in.PodIndexesToBeCreated, &out.PodIndexesToBeCreated
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerGroupSpec.
func (in *WorkerGroupSpec) DeepCopy() *WorkerGroupSpec {
	if in == nil {
		return nil
	}
	out := new(WorkerGroupSpec)
	in.DeepCopyInto(out)
	return out
}
