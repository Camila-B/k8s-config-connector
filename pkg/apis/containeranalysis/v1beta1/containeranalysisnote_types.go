// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NoteAffectedVersionEnd struct {
	/* Used to correct mistakes in the version numbering scheme. */
	Epoch int `json:"epoch,omitempty"`
	/* Human readable version string. This string is of the form :- and is only set when kind is NORMAL. */
	FullName string `json:"fullName,omitempty"`
	/* Required. Distinguishes between sentinel MIN/MAX versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE */
	Kind string `json:"kind,omitempty"`
	/* Required only when version kind is NORMAL. The main part of the version name. */
	Name string `json:"name,omitempty"`
	/* The iteration of the package build from the above version. */
	Revision string `json:"revision,omitempty"`
}

type NoteAffectedVersionStart struct {
	/* Used to correct mistakes in the version numbering scheme. */
	Epoch int `json:"epoch,omitempty"`
	/* Human readable version string. This string is of the form :- and is only set when kind is NORMAL. */
	FullName string `json:"fullName,omitempty"`
	/* Required. Distinguishes between sentinel MIN/MAX versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE */
	Kind string `json:"kind,omitempty"`
	/* Required only when version kind is NORMAL. The main part of the version name. */
	Name string `json:"name,omitempty"`
	/* The iteration of the package build from the above version. */
	Revision string `json:"revision,omitempty"`
}

type NoteAttestation struct {
	/* Hint hints at the purpose of the attestation authority. */
	Hint NoteHint `json:"hint,omitempty"`
}

type NoteBuild struct {
	/* Required. Immutable. Version of the builder which produced this build. */
	BuilderVersion string `json:"builderVersion,omitempty"`
}

type NoteCvssV3 struct {
	/*  Possible values: ATTACK_COMPLEXITY_UNSPECIFIED, ATTACK_COMPLEXITY_LOW, ATTACK_COMPLEXITY_HIGH */
	AttackComplexity string `json:"attackComplexity,omitempty"`
	/* Base Metrics Represents the intrinsic characteristics of a vulnerability that are constant over time and across user environments. Possible values: ATTACK_VECTOR_UNSPECIFIED, ATTACK_VECTOR_NETWORK, ATTACK_VECTOR_ADJACENT, ATTACK_VECTOR_LOCAL, ATTACK_VECTOR_PHYSICAL */
	AttackVector string `json:"attackVector,omitempty"`
	/*  Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW, IMPACT_NONE */
	AvailabilityImpact string `json:"availabilityImpact,omitempty"`
	/* The base score is a function of the base metric scores. */
	BaseScore float64 `json:"baseScore,omitempty"`
	/*  Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW, IMPACT_NONE */
	ConfidentialityImpact string `json:"confidentialityImpact,omitempty"`
	/*  */
	ExploitabilityScore float64 `json:"exploitabilityScore,omitempty"`
	/*  */
	ImpactScore float64 `json:"impactScore,omitempty"`
	/*  Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW, IMPACT_NONE */
	IntegrityImpact string `json:"integrityImpact,omitempty"`
	/*  Possible values: PRIVILEGES_REQUIRED_UNSPECIFIED, PRIVILEGES_REQUIRED_NONE, PRIVILEGES_REQUIRED_LOW, PRIVILEGES_REQUIRED_HIGH */
	PrivilegesRequired string `json:"privilegesRequired,omitempty"`
	/*  Possible values: SCOPE_UNSPECIFIED, SCOPE_UNCHANGED, SCOPE_CHANGED */
	Scope string `json:"scope,omitempty"`
	/*  Possible values: USER_INTERACTION_UNSPECIFIED, USER_INTERACTION_NONE, USER_INTERACTION_REQUIRED */
	UserInteraction string `json:"userInteraction,omitempty"`
}

type NoteDeployment struct {
	/* Required. Resource URI for the artifact being deployed. */
	ResourceUri []string `json:"resourceUri,omitempty"`
}

type NoteDetails struct {
	/* Required. The (https://cpe.mitre.org/specification/) this vulnerability affects. */
	AffectedCpeUri string `json:"affectedCpeUri,omitempty"`
	/* Required. The package this vulnerability affects. */
	AffectedPackage string `json:"affectedPackage,omitempty"`
	/* The version number at the end of an interval in which this vulnerability exists. A vulnerability can affect a package between version numbers that are disjoint sets of intervals (example: ) each of which will be represented in its own Detail. If a specific affected version is provided by a vulnerability database, affected_version_start and affected_version_end will be the same in that Detail. */
	AffectedVersionEnd NoteAffectedVersionEnd `json:"affectedVersionEnd,omitempty"`
	/* The version number at the start of an interval in which this vulnerability exists. A vulnerability can affect a package between version numbers that are disjoint sets of intervals (example: ) each of which will be represented in its own Detail. If a specific affected version is provided by a vulnerability database, affected_version_start and affected_version_end will be the same in that Detail. */
	AffectedVersionStart NoteAffectedVersionStart `json:"affectedVersionStart,omitempty"`
	/* A vendor-specific description of this vulnerability. */
	Description string `json:"description,omitempty"`
	/* The distro recommended (https://cpe.mitre.org/specification/) to update to that contains a fix for this vulnerability. It is possible for this to be different from the affected_cpe_uri. */
	FixedCpeUri string `json:"fixedCpeUri,omitempty"`
	/* The distro recommended package to update to that contains a fix for this vulnerability. It is possible for this to be different from the affected_package. */
	FixedPackage string `json:"fixedPackage,omitempty"`
	/* The distro recommended version to update to that contains a fix for this vulnerability. Setting this to VersionKind.MAXIMUM means no such version is yet available. */
	FixedVersion NoteFixedVersion `json:"fixedVersion,omitempty"`
	/* Whether this detail is obsolete. Occurrences are expected not to point to obsolete details. */
	IsObsolete bool `json:"isObsolete,omitempty"`
	/* The type of package; whether native or non native (e.g., ruby gems, node.js packages, etc.). */
	PackageType string `json:"packageType,omitempty"`
	/* The distro assigned severity of this vulnerability. */
	SeverityName string `json:"severityName,omitempty"`
	/* The time this information was last changed at the source. This is an upstream timestamp from the underlying information source - e.g. Ubuntu security tracker. */
	SourceUpdateTime string `json:"sourceUpdateTime,omitempty"`
}

type NoteDiscovery struct {
	/* Required. Immutable. The kind of analysis that is handled by this discovery. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE */
	AnalysisKind string `json:"analysisKind,omitempty"`
}

type NoteDistribution struct {
	/* The CPU architecture for which packages in this distribution channel were built. Possible values: ARCHITECTURE_UNSPECIFIED, X86, X64 */
	Architecture string `json:"architecture,omitempty"`
	/* Required. The cpe_uri in (https://cpe.mitre.org/specification/) denoting the package manager version distributing a package. */
	CpeUri string `json:"cpeUri,omitempty"`
	/* The distribution channel-specific description of this package. */
	Description string `json:"description,omitempty"`
	/* The latest available version of this package in this distribution channel. */
	LatestVersion NoteLatestVersion `json:"latestVersion,omitempty"`
	/* A freeform string denoting the maintainer of this package. */
	Maintainer string `json:"maintainer,omitempty"`
	/* The distribution channel-specific homepage for this package. */
	Url string `json:"url,omitempty"`
}

type NoteFingerprint struct {
	/* Required. The layer ID of the final layer in the Docker image's v1 representation. */
	V1Name string `json:"v1Name,omitempty"`
	/* Required. The ordered list of v2 blobs that represent a given image. */
	V2Blob []string `json:"v2Blob,omitempty"`
}

type NoteFixedVersion struct {
	/* Used to correct mistakes in the version numbering scheme. */
	Epoch int `json:"epoch,omitempty"`
	/* Human readable version string. This string is of the form :- and is only set when kind is NORMAL. */
	FullName string `json:"fullName,omitempty"`
	/* Required. Distinguishes between sentinel MIN/MAX versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE */
	Kind string `json:"kind,omitempty"`
	/* Required only when version kind is NORMAL. The main part of the version name. */
	Name string `json:"name,omitempty"`
	/* The iteration of the package build from the above version. */
	Revision string `json:"revision,omitempty"`
}

type NoteFixingKbs struct {
	/* The KB name (generally of the form KB+ (e.g., KB123456)). */
	Name string `json:"name,omitempty"`
	/* A link to the KB in the (https://www.catalog.update.microsoft.com/). */
	Url string `json:"url,omitempty"`
}

type NoteHint struct {
	/* Required. The human readable name of this attestation authority, for example "qa". */
	HumanReadableName string `json:"humanReadableName,omitempty"`
}

type NoteImage struct {
	/* Required. Immutable. The fingerprint of the base image. */
	Fingerprint NoteFingerprint `json:"fingerprint,omitempty"`
	/* Required. Immutable. The resource_url for the resource representing the basis of associated occurrence images. */
	ResourceUrl string `json:"resourceUrl,omitempty"`
}

type NoteLatestVersion struct {
	/* Used to correct mistakes in the version numbering scheme. */
	Epoch int `json:"epoch,omitempty"`
	/* Human readable version string. This string is of the form :- and is only set when kind is NORMAL. */
	FullName string `json:"fullName,omitempty"`
	/* Required. Distinguishes between sentinel MIN/MAX versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE */
	Kind string `json:"kind,omitempty"`
	/* Required only when version kind is NORMAL. The main part of the version name. */
	Name string `json:"name,omitempty"`
	/* The iteration of the package build from the above version. */
	Revision string `json:"revision,omitempty"`
}

type NotePackage struct {
	/* The various channels by which a package is distributed. */
	Distribution []NoteDistribution `json:"distribution,omitempty"`
	/* Required. Immutable. The name of the package. */
	Name string `json:"name,omitempty"`
}

type NoteRelatedUrl struct {
	/* Label to describe usage of the URL. */
	Label string `json:"label,omitempty"`
	/* Specific URL associated with the resource. */
	Url string `json:"url,omitempty"`
}

type NoteVulnerability struct {
	/* The CVSS score of this vulnerability. CVSS score is on a scale of 0 - 10 where 0 indicates low severity and 10 indicates high severity. */
	CvssScore float64 `json:"cvssScore,omitempty"`
	/* The full description of the CVSSv3 for this vulnerability. */
	CvssV3 NoteCvssV3 `json:"cvssV3,omitempty"`
	/* Details of all known distros and packages affected by this vulnerability. */
	Details []NoteDetails `json:"details,omitempty"`
	/* The note provider assigned severity of this vulnerability. Possible values: SEVERITY_UNSPECIFIED, MINIMAL, LOW, MEDIUM, HIGH, CRITICAL */
	Severity string `json:"severity,omitempty"`
	/* The time this information was last changed at the source. This is an upstream timestamp from the underlying information source - e.g. Ubuntu security tracker. */
	SourceUpdateTime string `json:"sourceUpdateTime,omitempty"`
	/* Windows details get their own format because the information format and model don't match a normal detail. Specifically Windows updates are done as patches, thus Windows vulnerabilities really are a missing package, rather than a package being at an incorrect version. */
	WindowsDetails []NoteWindowsDetails `json:"windowsDetails,omitempty"`
}

type NoteWindowsDetails struct {
	/* Required. The (https://cpe.mitre.org/specification/) this vulnerability affects. */
	CpeUri string `json:"cpeUri,omitempty"`
	/* The description of this vulnerability. */
	Description string `json:"description,omitempty"`
	/* Required. The names of the KBs which have hotfixes to mitigate this vulnerability. Note that there may be multiple hotfixes (and thus multiple KBs) that mitigate a given vulnerability. Currently any listed KBs presence is considered a fix. */
	FixingKbs []NoteFixingKbs `json:"fixingKbs,omitempty"`
	/* Required. The name of this vulnerability. */
	Name string `json:"name,omitempty"`
}

type ContainerAnalysisNoteSpec struct {
	/* A note describing an attestation role. */
	Attestation NoteAttestation `json:"attestation,omitempty"`
	/* A note describing build provenance for a verifiable build. */
	Build NoteBuild `json:"build,omitempty"`
	/* A note describing something that can be deployed. */
	Deployment NoteDeployment `json:"deployment,omitempty"`
	/* A note describing the initial analysis of a resource. */
	Discovery NoteDiscovery `json:"discovery,omitempty"`
	/* Time of expiration for this note. Empty if note does not expire. */
	ExpirationTime string `json:"expirationTime,omitempty"`
	/* A note describing a base image. */
	Image NoteImage `json:"image,omitempty"`
	/* A detailed description of this note. */
	LongDescription string `json:"longDescription,omitempty"`
	/* Required for non-Windows OS. The package this Upgrade is for. */
	Package NotePackage `json:"package,omitempty"`
	/*  */
	RelatedNoteNames []v1alpha1.ResourceRef `json:"relatedNoteNames,omitempty"`
	/* URLs associated with this note. */
	RelatedUrl []NoteRelatedUrl `json:"relatedUrl,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* A one sentence description of this note. */
	ShortDescription string `json:"shortDescription,omitempty"`
	/* A note describing a package vulnerability. */
	Vulnerability NoteVulnerability `json:"vulnerability,omitempty"`
}

type NoteFingerprintStatus struct {
	/* Output only. The name of the image's v2 blobs computed via: ) Only the name of the final blob is kept. */
	V2Name string `json:"v2Name,omitempty"`
}

type NoteImageStatus struct {
	/*  */
	Fingerprint NoteFingerprintStatus `json:"fingerprint,omitempty"`
}

type ContainerAnalysisNoteStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerAnalysisNote's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. The time this note was created. This field can be used as a filter in list requests. */
	CreateTime string `json:"createTime,omitempty"`
	/*  */
	Image NoteImageStatus `json:"image,omitempty"`
	/* Output only. The time this note was last updated. This field can be used as a filter in list requests. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerAnalysisNote is the Schema for the containeranalysis API
// +k8s:openapi-gen=true
type ContainerAnalysisNote struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerAnalysisNoteSpec   `json:"spec,omitempty"`
	Status ContainerAnalysisNoteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerAnalysisNoteList contains a list of ContainerAnalysisNote
type ContainerAnalysisNoteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerAnalysisNote `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerAnalysisNote{}, &ContainerAnalysisNoteList{})
}