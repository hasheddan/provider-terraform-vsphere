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
	r, ok := mr.(*ComputeClusterVmAntiAffinityRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeComputeClusterVmAntiAffinityRule(r, ctyValue)
}

func DecodeComputeClusterVmAntiAffinityRule(prev *ComputeClusterVmAntiAffinityRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeComputeClusterVmAntiAffinityRule_ComputeClusterId(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmAntiAffinityRule_Enabled(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmAntiAffinityRule_Mandatory(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmAntiAffinityRule_Name(&new.Spec.ForProvider, valMap)
	DecodeComputeClusterVmAntiAffinityRule_VirtualMachineIds(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmAntiAffinityRule_ComputeClusterId(p *ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.ComputeClusterId = ctwhy.ValueAsString(vals["compute_cluster_id"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmAntiAffinityRule_Enabled(p *ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmAntiAffinityRule_Mandatory(p *ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.Mandatory = ctwhy.ValueAsBool(vals["mandatory"])
}

//primitiveTypeDecodeTemplate
func DecodeComputeClusterVmAntiAffinityRule_Name(p *ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeComputeClusterVmAntiAffinityRule_VirtualMachineIds(p *ComputeClusterVmAntiAffinityRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["virtual_machine_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.VirtualMachineIds = goVals
}