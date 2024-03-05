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
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this Workflow.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Workflow) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccount),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.ServiceAccountRef,
			Selector:     mg.Spec.ForProvider.ServiceAccountSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccount")
	}
	mg.Spec.ForProvider.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cloudplatform.gcp.upbound.io", "v1beta1", "ServiceAccount", "ServiceAccountList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ServiceAccount),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.ServiceAccountRef,
			Selector:     mg.Spec.InitProvider.ServiceAccountSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ServiceAccount")
	}
	mg.Spec.InitProvider.ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ServiceAccountRef = rsp.ResolvedReference

	return nil
}
