// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dataplex

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgiamresource"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var DataplexAssetIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"lake": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"dataplex_zone": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"asset": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
	},
}

type DataplexAssetIamUpdater struct {
	project      string
	location     string
	lake         string
	dataplexZone string
	asset        string
	d            tpgresource.TerraformResourceData
	Config       *transport_tpg.Config
}

func DataplexAssetIamUpdaterProducer(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgiamresource.ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("lake"); ok {
		values["lake"] = v.(string)
	}

	if v, ok := d.GetOk("dataplex_zone"); ok {
		values["dataplex_zone"] = v.(string)
	}

	if v, ok := d.GetOk("asset"); ok {
		values["asset"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/lakes/(?P<lake>[^/]+)/zones/(?P<dataplex_zone>[^/]+)/assets/(?P<asset>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)/(?P<asset>[^/]+)", "(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)/(?P<asset>[^/]+)", "(?P<asset>[^/]+)"}, d, config, d.Get("asset").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataplexAssetIamUpdater{
		project:      values["project"],
		location:     values["location"],
		lake:         values["lake"],
		dataplexZone: values["dataplex_zone"],
		asset:        values["asset"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("lake", u.lake); err != nil {
		return nil, fmt.Errorf("Error setting lake: %s", err)
	}
	if err := d.Set("dataplex_zone", u.dataplexZone); err != nil {
		return nil, fmt.Errorf("Error setting dataplex_zone: %s", err)
	}
	if err := d.Set("asset", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting asset: %s", err)
	}

	return u, nil
}

func DataplexAssetIdParseFunc(d *schema.ResourceData, config *transport_tpg.Config) error {
	values := make(map[string]string)

	project, _ := tpgresource.GetProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := tpgresource.GetLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := tpgresource.GetImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/lakes/(?P<lake>[^/]+)/zones/(?P<dataplex_zone>[^/]+)/assets/(?P<asset>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)/(?P<asset>[^/]+)", "(?P<location>[^/]+)/(?P<lake>[^/]+)/(?P<dataplex_zone>[^/]+)/(?P<asset>[^/]+)", "(?P<asset>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataplexAssetIamUpdater{
		project:      values["project"],
		location:     values["location"],
		lake:         values["lake"],
		dataplexZone: values["dataplex_zone"],
		asset:        values["asset"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("asset", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting asset: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataplexAssetIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyAssetUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "GET",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
	})
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = tpgresource.Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataplexAssetIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := tpgresource.ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyAssetUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := tpgresource.GetProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := tpgresource.GenerateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    u.Config,
		Method:    "POST",
		Project:   project,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   u.d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataplexAssetIamUpdater) qualifyAssetUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataplexBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", u.project, u.location, u.lake, u.dataplexZone, u.asset), methodIdentifier)
	url, err := tpgresource.ReplaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataplexAssetIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/lakes/%s/zones/%s/assets/%s", u.project, u.location, u.lake, u.dataplexZone, u.asset)
}

func (u *DataplexAssetIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-dataplex-asset-%s", u.GetResourceId())
}

func (u *DataplexAssetIamUpdater) DescribeResource() string {
	return fmt.Sprintf("dataplex asset %q", u.GetResourceId())
}