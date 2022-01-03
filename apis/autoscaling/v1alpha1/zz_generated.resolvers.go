/*
Copyright 2021 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-jet-aws/apis/elasticloadbalancing/v1alpha1"
	common "github.com/crossplane-contrib/provider-jet-aws/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Attachment.
func (mg *Attachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ALBTargetGroupArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.ALBTargetGroupArnRef,
		Selector:     mg.Spec.ForProvider.ALBTargetGroupArnSelector,
		To: reference.To{
			List:    &v1alpha1.TargetGroupList{},
			Managed: &v1alpha1.TargetGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ALBTargetGroupArn")
	}
	mg.Spec.ForProvider.ALBTargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ALBTargetGroupArnRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AutoscalingGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.AutoscalingGroupNameRef,
		Selector:     mg.Spec.ForProvider.AutoscalingGroupNameSelector,
		To: reference.To{
			List:    &AutoscalingGroupList{},
			Managed: &AutoscalingGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AutoscalingGroupName")
	}
	mg.Spec.ForProvider.AutoscalingGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AutoscalingGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this AutoscalingGroup.
func (mg *AutoscalingGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VPCZoneIdentifier),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VPCZoneIdentifierRefs,
		Selector:      mg.Spec.ForProvider.VPCZoneIdentifierSelector,
		To: reference.To{
			List:    &v1alpha11.SubnetList{},
			Managed: &v1alpha11.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCZoneIdentifier")
	}
	mg.Spec.ForProvider.VPCZoneIdentifier = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VPCZoneIdentifierRefs = mrsp.ResolvedReferences

	return nil
}
