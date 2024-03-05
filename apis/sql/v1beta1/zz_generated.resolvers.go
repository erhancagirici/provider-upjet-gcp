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

	// ResolveReferences of this Database.
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
)

func (mg *Database) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "DatabaseInstance", "DatabaseInstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceRef,
			Selector:     mg.Spec.ForProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this DatabaseInstance.
func (mg *DatabaseInstance) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Settings); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Settings[i3].IPConfiguration); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Settings[i3].IPConfiguration[i4].PrivateNetwork),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.Settings[i3].IPConfiguration[i4].PrivateNetworkRef,
					Selector:     mg.Spec.ForProvider.Settings[i3].IPConfiguration[i4].PrivateNetworkSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Settings[i3].IPConfiguration[i4].PrivateNetwork")
			}
			mg.Spec.ForProvider.Settings[i3].IPConfiguration[i4].PrivateNetwork = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Settings[i3].IPConfiguration[i4].PrivateNetworkRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Settings); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Settings[i3].IPConfiguration); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("compute.gcp.upbound.io", "v1beta1", "Network", "NetworkList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Settings[i3].IPConfiguration[i4].PrivateNetwork),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.Settings[i3].IPConfiguration[i4].PrivateNetworkRef,
					Selector:     mg.Spec.InitProvider.Settings[i3].IPConfiguration[i4].PrivateNetworkSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Settings[i3].IPConfiguration[i4].PrivateNetwork")
			}
			mg.Spec.InitProvider.Settings[i3].IPConfiguration[i4].PrivateNetwork = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Settings[i3].IPConfiguration[i4].PrivateNetworkRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this SSLCert.
func (mg *SSLCert) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "DatabaseInstance", "DatabaseInstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceRef,
			Selector:     mg.Spec.ForProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "DatabaseInstance", "DatabaseInstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.InstanceRef,
			Selector:     mg.Spec.InitProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Instance")
	}
	mg.Spec.InitProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this User.
func (mg *User) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "DatabaseInstance", "DatabaseInstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.InstanceRef,
			Selector:     mg.Spec.ForProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("sql.gcp.upbound.io", "v1beta1", "DatabaseInstance", "DatabaseInstanceList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.InstanceRef,
			Selector:     mg.Spec.InitProvider.InstanceSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Instance")
	}
	mg.Spec.InitProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.InstanceRef = rsp.ResolvedReference

	return nil
}
