/*
	Copyright 2019 The Crossplane Authors.

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

package v1alpha1

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*ComputeClusterVmHostRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeComputeClusterVmHostRule(r, ctyValue)
}

func DecodeComputeClusterVmHostRule(prev *ComputeClusterVmHostRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeComputeClusterVmHostRule_AntiAffinityHostGroupName(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmHostRule_ComputeClusterId(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmHostRule_Enabled(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmHostRule_Mandatory(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmHostRule_Name(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmHostRule_VmGroupName(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmHostRule_AffinityHostGroupName(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmHostRule_AntiAffinityHostGroupName(p *ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	p.AntiAffinityHostGroupName = ctwhy.ValueAsString(vals["anti_affinity_host_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmHostRule_ComputeClusterId(p *ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	p.ComputeClusterId = ctwhy.ValueAsString(vals["compute_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmHostRule_Enabled(p *ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmHostRule_Mandatory(p *ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	p.Mandatory = ctwhy.ValueAsBool(vals["mandatory"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmHostRule_Name(p *ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmHostRule_VmGroupName(p *ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	p.VmGroupName = ctwhy.ValueAsString(vals["vm_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmHostRule_AffinityHostGroupName(p *ComputeClusterVmHostRuleParameters, vals map[string]cty.Value) {
	p.AffinityHostGroupName = ctwhy.ValueAsString(vals["affinity_host_group_name"])
}