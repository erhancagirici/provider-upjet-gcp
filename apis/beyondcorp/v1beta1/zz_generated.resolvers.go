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
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AppConnection) ResolveReferences( // ResolveReferences of this AppConnection.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Gateway); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("beyondcorp.gcp.upbound.io", "v1beta1", "AppGateway", "AppGatewayList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Gateway[i3].AppGateway),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.Gateway[i3].AppGatewayRef,
				Selector:     mg.Spec.ForProvider.Gateway[i3].AppGatewaySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Gateway[i3].AppGateway")
		}
		mg.Spec.ForProvider.Gateway[i3].AppGateway = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Gateway[i3].AppGatewayRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Gateway); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("beyondcorp.gcp.upbound.io", "v1beta1", "AppGateway", "AppGatewayList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Gateway[i3].AppGateway),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.Gateway[i3].AppGatewayRef,
				Selector:     mg.Spec.InitProvider.Gateway[i3].AppGatewaySelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Gateway[i3].AppGateway")
		}
		mg.Spec.InitProvider.Gateway[i3].AppGateway = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Gateway[i3].AppGatewayRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this AppConnector.
func (mg *AppConnector) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.PrincipalInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.PrincipalInfo[i3].ServiceAccount); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrincipalInfo[i3].ServiceAccount[i4].Email),
					Extract:      resource.ExtractParamPath("email", true),
					Reference:    mg.Spec.ForProvider.PrincipalInfo[i3].ServiceAccount[i4].EmailRef,
					Selector:     mg.Spec.ForProvider.PrincipalInfo[i3].ServiceAccount[i4].EmailSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.PrincipalInfo[i3].ServiceAccount[i4].Email")
			}
			mg.Spec.ForProvider.PrincipalInfo[i3].ServiceAccount[i4].Email = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.PrincipalInfo[i3].ServiceAccount[i4].EmailRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.PrincipalInfo); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.PrincipalInfo[i3].ServiceAccount); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrincipalInfo[i3].ServiceAccount[i4].Email),
					Extract:      resource.ExtractParamPath("email", true),
					Reference:    mg.Spec.InitProvider.PrincipalInfo[i3].ServiceAccount[i4].EmailRef,
					Selector:     mg.Spec.InitProvider.PrincipalInfo[i3].ServiceAccount[i4].EmailSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.PrincipalInfo[i3].ServiceAccount[i4].Email")
			}
			mg.Spec.InitProvider.PrincipalInfo[i3].ServiceAccount[i4].Email = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.PrincipalInfo[i3].ServiceAccount[i4].EmailRef = rsp.ResolvedReference

		}
	}

	return nil
}
