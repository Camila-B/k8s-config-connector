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
func (in *BigQueryDataset) DeepCopyInto(out *BigQueryDataset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDataset.
func (in *BigQueryDataset) DeepCopy() *BigQueryDataset {
	if in == nil {
		return nil
	}
	out := new(BigQueryDataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryDataset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDatasetList) DeepCopyInto(out *BigQueryDatasetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryDataset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDatasetList.
func (in *BigQueryDatasetList) DeepCopy() *BigQueryDatasetList {
	if in == nil {
		return nil
	}
	out := new(BigQueryDatasetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryDatasetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDatasetSpec) DeepCopyInto(out *BigQueryDatasetSpec) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]DatasetAccess, len(*in))
		copy(*out, *in)
	}
	out.DefaultEncryptionConfiguration = in.DefaultEncryptionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDatasetSpec.
func (in *BigQueryDatasetSpec) DeepCopy() *BigQueryDatasetSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryDatasetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDatasetStatus) DeepCopyInto(out *BigQueryDatasetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDatasetStatus.
func (in *BigQueryDatasetStatus) DeepCopy() *BigQueryDatasetStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryDatasetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJob) DeepCopyInto(out *BigQueryJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJob.
func (in *BigQueryJob) DeepCopy() *BigQueryJob {
	if in == nil {
		return nil
	}
	out := new(BigQueryJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJobList) DeepCopyInto(out *BigQueryJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJobList.
func (in *BigQueryJobList) DeepCopy() *BigQueryJobList {
	if in == nil {
		return nil
	}
	out := new(BigQueryJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJobSpec) DeepCopyInto(out *BigQueryJobSpec) {
	*out = *in
	in.Copy.DeepCopyInto(&out.Copy)
	in.Extract.DeepCopyInto(&out.Extract)
	in.Load.DeepCopyInto(&out.Load)
	in.Query.DeepCopyInto(&out.Query)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJobSpec.
func (in *BigQueryJobSpec) DeepCopy() *BigQueryJobSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJobStatus) DeepCopyInto(out *BigQueryJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]JobStatusStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJobStatus.
func (in *BigQueryJobStatus) DeepCopy() *BigQueryJobStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTable) DeepCopyInto(out *BigQueryTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTable.
func (in *BigQueryTable) DeepCopy() *BigQueryTable {
	if in == nil {
		return nil
	}
	out := new(BigQueryTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTableList) DeepCopyInto(out *BigQueryTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTableList.
func (in *BigQueryTableList) DeepCopy() *BigQueryTableList {
	if in == nil {
		return nil
	}
	out := new(BigQueryTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTableSpec) DeepCopyInto(out *BigQueryTableSpec) {
	*out = *in
	if in.Clustering != nil {
		in, out := &in.Clustering, &out.Clustering
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DatasetRef = in.DatasetRef
	out.EncryptionConfiguration = in.EncryptionConfiguration
	in.ExternalDataConfiguration.DeepCopyInto(&out.ExternalDataConfiguration)
	out.MaterializedView = in.MaterializedView
	out.RangePartitioning = in.RangePartitioning
	out.TimePartitioning = in.TimePartitioning
	out.View = in.View
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTableSpec.
func (in *BigQueryTableSpec) DeepCopy() *BigQueryTableSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTableStatus) DeepCopyInto(out *BigQueryTableStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTableStatus.
func (in *BigQueryTableStatus) DeepCopy() *BigQueryTableStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetAccess) DeepCopyInto(out *DatasetAccess) {
	*out = *in
	out.View = in.View
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetAccess.
func (in *DatasetAccess) DeepCopy() *DatasetAccess {
	if in == nil {
		return nil
	}
	out := new(DatasetAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetDefaultEncryptionConfiguration) DeepCopyInto(out *DatasetDefaultEncryptionConfiguration) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetDefaultEncryptionConfiguration.
func (in *DatasetDefaultEncryptionConfiguration) DeepCopy() *DatasetDefaultEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(DatasetDefaultEncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetView) DeepCopyInto(out *DatasetView) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetView.
func (in *DatasetView) DeepCopy() *DatasetView {
	if in == nil {
		return nil
	}
	out := new(DatasetView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobCopy) DeepCopyInto(out *JobCopy) {
	*out = *in
	out.DestinationEncryptionConfiguration = in.DestinationEncryptionConfiguration
	out.DestinationTable = in.DestinationTable
	if in.SourceTables != nil {
		in, out := &in.SourceTables, &out.SourceTables
		*out = make([]JobSourceTables, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobCopy.
func (in *JobCopy) DeepCopy() *JobCopy {
	if in == nil {
		return nil
	}
	out := new(JobCopy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDefaultDataset) DeepCopyInto(out *JobDefaultDataset) {
	*out = *in
	out.DatasetRef = in.DatasetRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDefaultDataset.
func (in *JobDefaultDataset) DeepCopy() *JobDefaultDataset {
	if in == nil {
		return nil
	}
	out := new(JobDefaultDataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDestinationEncryptionConfiguration) DeepCopyInto(out *JobDestinationEncryptionConfiguration) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDestinationEncryptionConfiguration.
func (in *JobDestinationEncryptionConfiguration) DeepCopy() *JobDestinationEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(JobDestinationEncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobDestinationTable) DeepCopyInto(out *JobDestinationTable) {
	*out = *in
	out.TableRef = in.TableRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobDestinationTable.
func (in *JobDestinationTable) DeepCopy() *JobDestinationTable {
	if in == nil {
		return nil
	}
	out := new(JobDestinationTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobErrorResultStatus) DeepCopyInto(out *JobErrorResultStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobErrorResultStatus.
func (in *JobErrorResultStatus) DeepCopy() *JobErrorResultStatus {
	if in == nil {
		return nil
	}
	out := new(JobErrorResultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobErrorsStatus) DeepCopyInto(out *JobErrorsStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobErrorsStatus.
func (in *JobErrorsStatus) DeepCopy() *JobErrorsStatus {
	if in == nil {
		return nil
	}
	out := new(JobErrorsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobExtract) DeepCopyInto(out *JobExtract) {
	*out = *in
	if in.DestinationUris != nil {
		in, out := &in.DestinationUris, &out.DestinationUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SourceTable = in.SourceTable
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobExtract.
func (in *JobExtract) DeepCopy() *JobExtract {
	if in == nil {
		return nil
	}
	out := new(JobExtract)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobLoad) DeepCopyInto(out *JobLoad) {
	*out = *in
	out.DestinationEncryptionConfiguration = in.DestinationEncryptionConfiguration
	out.DestinationTable = in.DestinationTable
	if in.ProjectionFields != nil {
		in, out := &in.ProjectionFields, &out.ProjectionFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SchemaUpdateOptions != nil {
		in, out := &in.SchemaUpdateOptions, &out.SchemaUpdateOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceUris != nil {
		in, out := &in.SourceUris, &out.SourceUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.TimePartitioning = in.TimePartitioning
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobLoad.
func (in *JobLoad) DeepCopy() *JobLoad {
	if in == nil {
		return nil
	}
	out := new(JobLoad)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobQuery) DeepCopyInto(out *JobQuery) {
	*out = *in
	out.DefaultDataset = in.DefaultDataset
	out.DestinationEncryptionConfiguration = in.DestinationEncryptionConfiguration
	out.DestinationTable = in.DestinationTable
	if in.SchemaUpdateOptions != nil {
		in, out := &in.SchemaUpdateOptions, &out.SchemaUpdateOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ScriptOptions = in.ScriptOptions
	if in.UserDefinedFunctionResources != nil {
		in, out := &in.UserDefinedFunctionResources, &out.UserDefinedFunctionResources
		*out = make([]JobUserDefinedFunctionResources, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobQuery.
func (in *JobQuery) DeepCopy() *JobQuery {
	if in == nil {
		return nil
	}
	out := new(JobQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobScriptOptions) DeepCopyInto(out *JobScriptOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobScriptOptions.
func (in *JobScriptOptions) DeepCopy() *JobScriptOptions {
	if in == nil {
		return nil
	}
	out := new(JobScriptOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSourceTable) DeepCopyInto(out *JobSourceTable) {
	*out = *in
	out.TableRef = in.TableRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSourceTable.
func (in *JobSourceTable) DeepCopy() *JobSourceTable {
	if in == nil {
		return nil
	}
	out := new(JobSourceTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSourceTables) DeepCopyInto(out *JobSourceTables) {
	*out = *in
	out.TableRef = in.TableRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSourceTables.
func (in *JobSourceTables) DeepCopy() *JobSourceTables {
	if in == nil {
		return nil
	}
	out := new(JobSourceTables)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatusStatus) DeepCopyInto(out *JobStatusStatus) {
	*out = *in
	if in.ErrorResult != nil {
		in, out := &in.ErrorResult, &out.ErrorResult
		*out = make([]JobErrorResultStatus, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]JobErrorsStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStatusStatus.
func (in *JobStatusStatus) DeepCopy() *JobStatusStatus {
	if in == nil {
		return nil
	}
	out := new(JobStatusStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTimePartitioning) DeepCopyInto(out *JobTimePartitioning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTimePartitioning.
func (in *JobTimePartitioning) DeepCopy() *JobTimePartitioning {
	if in == nil {
		return nil
	}
	out := new(JobTimePartitioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobUserDefinedFunctionResources) DeepCopyInto(out *JobUserDefinedFunctionResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobUserDefinedFunctionResources.
func (in *JobUserDefinedFunctionResources) DeepCopy() *JobUserDefinedFunctionResources {
	if in == nil {
		return nil
	}
	out := new(JobUserDefinedFunctionResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableCsvOptions) DeepCopyInto(out *TableCsvOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableCsvOptions.
func (in *TableCsvOptions) DeepCopy() *TableCsvOptions {
	if in == nil {
		return nil
	}
	out := new(TableCsvOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableEncryptionConfiguration) DeepCopyInto(out *TableEncryptionConfiguration) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableEncryptionConfiguration.
func (in *TableEncryptionConfiguration) DeepCopy() *TableEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(TableEncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableExternalDataConfiguration) DeepCopyInto(out *TableExternalDataConfiguration) {
	*out = *in
	out.CsvOptions = in.CsvOptions
	out.GoogleSheetsOptions = in.GoogleSheetsOptions
	out.HivePartitioningOptions = in.HivePartitioningOptions
	if in.SourceUris != nil {
		in, out := &in.SourceUris, &out.SourceUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableExternalDataConfiguration.
func (in *TableExternalDataConfiguration) DeepCopy() *TableExternalDataConfiguration {
	if in == nil {
		return nil
	}
	out := new(TableExternalDataConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableGoogleSheetsOptions) DeepCopyInto(out *TableGoogleSheetsOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableGoogleSheetsOptions.
func (in *TableGoogleSheetsOptions) DeepCopy() *TableGoogleSheetsOptions {
	if in == nil {
		return nil
	}
	out := new(TableGoogleSheetsOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableHivePartitioningOptions) DeepCopyInto(out *TableHivePartitioningOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableHivePartitioningOptions.
func (in *TableHivePartitioningOptions) DeepCopy() *TableHivePartitioningOptions {
	if in == nil {
		return nil
	}
	out := new(TableHivePartitioningOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableMaterializedView) DeepCopyInto(out *TableMaterializedView) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableMaterializedView.
func (in *TableMaterializedView) DeepCopy() *TableMaterializedView {
	if in == nil {
		return nil
	}
	out := new(TableMaterializedView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableRange) DeepCopyInto(out *TableRange) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableRange.
func (in *TableRange) DeepCopy() *TableRange {
	if in == nil {
		return nil
	}
	out := new(TableRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableRangePartitioning) DeepCopyInto(out *TableRangePartitioning) {
	*out = *in
	out.Range = in.Range
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableRangePartitioning.
func (in *TableRangePartitioning) DeepCopy() *TableRangePartitioning {
	if in == nil {
		return nil
	}
	out := new(TableRangePartitioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableTimePartitioning) DeepCopyInto(out *TableTimePartitioning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableTimePartitioning.
func (in *TableTimePartitioning) DeepCopy() *TableTimePartitioning {
	if in == nil {
		return nil
	}
	out := new(TableTimePartitioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TableView) DeepCopyInto(out *TableView) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TableView.
func (in *TableView) DeepCopy() *TableView {
	if in == nil {
		return nil
	}
	out := new(TableView)
	in.DeepCopyInto(out)
	return out
}