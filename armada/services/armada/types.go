/*
Copyright 2018 The Kubernetes Authors.

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

package armada

type InterfacesDeleteFuture struct{}
type SecurityGroupsCreateOrUpdateFuture struct{}
type SecurityGroup struct{}
type IPAddress struct {
	IPAddress string
}
type PublicIPAddress struct {
	IPAddress string
}
type PublicIPAddressesDeleteFuture struct{}
type VirtualNetworksCreateOrUpdateFuture struct{}
type VirtualNetworksClient struct{}
type InterfacesClient struct{}
type LoadBalancersClient struct{}
type PublicIPAddressesClient struct{}
type SecurityGroupsClient struct{}
type SecurityRule struct{}
type SecurityGroupPropertiesFormat struct{}
type SecurityGroupsDeleteFuture struct{}
type Subnet struct{}
type VirtualNetworkPropertiesFormat struct{}
type VirtualNetwork struct{}
