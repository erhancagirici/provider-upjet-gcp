// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	defaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/defaultsupportedidpconfig"
	inboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/inboundsamlconfig"
	oauthidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/oauthidpconfig"
	projectdefaultconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/projectdefaultconfig"
	tenant "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenant"
	tenantdefaultsupportedidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantdefaultsupportedidpconfig"
	tenantinboundsamlconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantinboundsamlconfig"
	tenantoauthidpconfig "github.com/upbound/provider-gcp/internal/controller/cluster/identityplatform/tenantoauthidpconfig"
)

// Setup_identityplatform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_identityplatform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		defaultsupportedidpconfig.Setup,
		inboundsamlconfig.Setup,
		oauthidpconfig.Setup,
		projectdefaultconfig.Setup,
		tenant.Setup,
		tenantdefaultsupportedidpconfig.Setup,
		tenantinboundsamlconfig.Setup,
		tenantoauthidpconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
