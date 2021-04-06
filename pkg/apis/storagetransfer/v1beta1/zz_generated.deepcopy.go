// +build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAccessKeyId) DeepCopyInto(out *JobAccessKeyId) {
	*out = *in
	out.ValueFrom = in.ValueFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAccessKeyId.
func (in *JobAccessKeyId) DeepCopy() *JobAccessKeyId {
	if in == nil {
		return nil
	}
	out := new(JobAccessKeyId)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAwsAccessKey) DeepCopyInto(out *JobAwsAccessKey) {
	*out = *in
	out.AccessKeyId = in.AccessKeyId
	out.SecretAccessKey = in.SecretAccessKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAwsAccessKey.
func (in *JobAwsAccessKey) DeepCopy() *JobAwsAccessKey {
	if in == nil {
		return nil
	}
	out := new(JobAwsAccessKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobAwsS3DataSource) DeepCopyInto(out *JobAwsS3DataSource) {
	*out = *in
	out.AwsAccessKey = in.AwsAccessKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobAwsS3DataSource.
func (in *JobAwsS3DataSource) DeepCopy() *JobAwsS3DataSource {
	if in == nil {
		return nil
	}
	out := new(JobAwsS3DataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobGcsDataSink) DeepCopyInto(out *JobGcsDataSink) {
	*out = *in
	out.BucketRef = in.BucketRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobGcsDataSink.
func (in *JobGcsDataSink) DeepCopy() *JobGcsDataSink {
	if in == nil {
		return nil
	}
	out := new(JobGcsDataSink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobGcsDataSource) DeepCopyInto(out *JobGcsDataSource) {
	*out = *in
	out.BucketRef = in.BucketRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobGcsDataSource.
func (in *JobGcsDataSource) DeepCopy() *JobGcsDataSource {
	if in == nil {
		return nil
	}
	out := new(JobGcsDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobHttpDataSource) DeepCopyInto(out *JobHttpDataSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobHttpDataSource.
func (in *JobHttpDataSource) DeepCopy() *JobHttpDataSource {
	if in == nil {
		return nil
	}
	out := new(JobHttpDataSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobObjectConditions) DeepCopyInto(out *JobObjectConditions) {
	*out = *in
	if in.ExcludePrefixes != nil {
		in, out := &in.ExcludePrefixes, &out.ExcludePrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludePrefixes != nil {
		in, out := &in.IncludePrefixes, &out.IncludePrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobObjectConditions.
func (in *JobObjectConditions) DeepCopy() *JobObjectConditions {
	if in == nil {
		return nil
	}
	out := new(JobObjectConditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSchedule) DeepCopyInto(out *JobSchedule) {
	*out = *in
	out.ScheduleEndDate = in.ScheduleEndDate
	out.ScheduleStartDate = in.ScheduleStartDate
	out.StartTimeOfDay = in.StartTimeOfDay
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSchedule.
func (in *JobSchedule) DeepCopy() *JobSchedule {
	if in == nil {
		return nil
	}
	out := new(JobSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobScheduleEndDate) DeepCopyInto(out *JobScheduleEndDate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobScheduleEndDate.
func (in *JobScheduleEndDate) DeepCopy() *JobScheduleEndDate {
	if in == nil {
		return nil
	}
	out := new(JobScheduleEndDate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobScheduleStartDate) DeepCopyInto(out *JobScheduleStartDate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobScheduleStartDate.
func (in *JobScheduleStartDate) DeepCopy() *JobScheduleStartDate {
	if in == nil {
		return nil
	}
	out := new(JobScheduleStartDate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSecretAccessKey) DeepCopyInto(out *JobSecretAccessKey) {
	*out = *in
	out.ValueFrom = in.ValueFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSecretAccessKey.
func (in *JobSecretAccessKey) DeepCopy() *JobSecretAccessKey {
	if in == nil {
		return nil
	}
	out := new(JobSecretAccessKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStartTimeOfDay) DeepCopyInto(out *JobStartTimeOfDay) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStartTimeOfDay.
func (in *JobStartTimeOfDay) DeepCopy() *JobStartTimeOfDay {
	if in == nil {
		return nil
	}
	out := new(JobStartTimeOfDay)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTransferOptions) DeepCopyInto(out *JobTransferOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTransferOptions.
func (in *JobTransferOptions) DeepCopy() *JobTransferOptions {
	if in == nil {
		return nil
	}
	out := new(JobTransferOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTransferSpec) DeepCopyInto(out *JobTransferSpec) {
	*out = *in
	out.AwsS3DataSource = in.AwsS3DataSource
	out.GcsDataSink = in.GcsDataSink
	out.GcsDataSource = in.GcsDataSource
	out.HttpDataSource = in.HttpDataSource
	in.ObjectConditions.DeepCopyInto(&out.ObjectConditions)
	out.TransferOptions = in.TransferOptions
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTransferSpec.
func (in *JobTransferSpec) DeepCopy() *JobTransferSpec {
	if in == nil {
		return nil
	}
	out := new(JobTransferSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobValueFrom) DeepCopyInto(out *JobValueFrom) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobValueFrom.
func (in *JobValueFrom) DeepCopy() *JobValueFrom {
	if in == nil {
		return nil
	}
	out := new(JobValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJob) DeepCopyInto(out *StorageTransferJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJob.
func (in *StorageTransferJob) DeepCopy() *StorageTransferJob {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageTransferJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJobList) DeepCopyInto(out *StorageTransferJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageTransferJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJobList.
func (in *StorageTransferJobList) DeepCopy() *StorageTransferJobList {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageTransferJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJobSpec) DeepCopyInto(out *StorageTransferJobSpec) {
	*out = *in
	out.Schedule = in.Schedule
	in.TransferSpec.DeepCopyInto(&out.TransferSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJobSpec.
func (in *StorageTransferJobSpec) DeepCopy() *StorageTransferJobSpec {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageTransferJobStatus) DeepCopyInto(out *StorageTransferJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageTransferJobStatus.
func (in *StorageTransferJobStatus) DeepCopy() *StorageTransferJobStatus {
	if in == nil {
		return nil
	}
	out := new(StorageTransferJobStatus)
	in.DeepCopyInto(out)
	return out
}