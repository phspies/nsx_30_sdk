# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBgpNeighbor**](ManagementPlaneApiApi.md#AddBgpNeighbor) | **Post** /logical-routers/{logical-router-id}/routing/bgp/neighbors | Add a new BGP Neighbor on a Logical Router
[**AddIPPrefixList**](ManagementPlaneApiApi.md#AddIPPrefixList) | **Post** /logical-routers/{logical-router-id}/routing/ip-prefix-lists | Add IPPrefixList on a Logical Router
[**AddInstanceEndpoint**](ManagementPlaneApiApi.md#AddInstanceEndpoint) | **Post** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-endpoints | Add an InstanceEndpoint for a Service Instance
[**AddMACAddress**](ManagementPlaneApiApi.md#AddMACAddress) | **Post** /mac-sets/{mac-set-id}/members | Add a MAC address to a MACSet
[**AddMemberAddMember**](ManagementPlaneApiApi.md#AddMemberAddMember) | **Post** /firewall/excludelist?action&#x3D;add_member | Add a new object in the exclude list
[**AddNatRule**](ManagementPlaneApiApi.md#AddNatRule) | **Post** /logical-routers/{logical-router-id}/nat/rules | Add a NAT rule in a specific logical router
[**AddNatRulesCreateMultiple**](ManagementPlaneApiApi.md#AddNatRulesCreateMultiple) | **Post** /logical-routers/{logical-router-id}/nat/rules?action&#x3D;create_multiple | Add multiple NAT rules in a specific logical router
[**AddOrRemoveNSGroupExpression**](ManagementPlaneApiApi.md#AddOrRemoveNSGroupExpression) | **Post** /ns-groups/{ns-group-id} | Add NSGroup expression
[**AddPBRRuleInSection**](ManagementPlaneApiApi.md#AddPBRRuleInSection) | **Post** /pbr/sections/{section-id}/rules | Add a Single Rule in a Section
[**AddPBRRulesInSectionCreateMultiple**](ManagementPlaneApiApi.md#AddPBRRulesInSectionCreateMultiple) | **Post** /pbr/sections/{section-id}/rules?action&#x3D;create_multiple | Add Multiple Rules in a Section
[**AddPBRSection**](ManagementPlaneApiApi.md#AddPBRSection) | **Post** /pbr/sections | Create a New Empty Section
[**AddPBRSectionWithRulesCreateWithRules**](ManagementPlaneApiApi.md#AddPBRSectionWithRulesCreateWithRules) | **Post** /pbr/sections?action&#x3D;create_with_rules | Create a Section with Rules
[**AddRemoveIPAddress**](ManagementPlaneApiApi.md#AddRemoveIPAddress) | **Post** /ip-sets/{ip-set-id} | Add a IP address to a IPSet
[**AddRouteMap**](ManagementPlaneApiApi.md#AddRouteMap) | **Post** /logical-routers/{logical-router-id}/routing/route-maps | Add RouteMap on a Logical Router
[**AddRuleInSection**](ManagementPlaneApiApi.md#AddRuleInSection) | **Post** /firewall/sections/{section-id}/rules | Add a Single Rule in a Section
[**AddRulesInSectionCreateMultiple**](ManagementPlaneApiApi.md#AddRulesInSectionCreateMultiple) | **Post** /firewall/sections/{section-id}/rules?action&#x3D;create_multiple | Add Multiple Rules in a Section
[**AddSIServiceProfile**](ManagementPlaneApiApi.md#AddSIServiceProfile) | **Post** /serviceinsertion/services/{service-id}/service-profiles | Add ServiceProfile for a given Service.
[**AddSection**](ManagementPlaneApiApi.md#AddSection) | **Post** /firewall/sections | Create a New Empty Section
[**AddSectionWithRulesCreateWithRules**](ManagementPlaneApiApi.md#AddSectionWithRulesCreateWithRules) | **Post** /firewall/sections?action&#x3D;create_with_rules | Create a Section with Rules
[**AddServiceAttachment**](ManagementPlaneApiApi.md#AddServiceAttachment) | **Post** /serviceinsertion/service-attachments | Add a Service Attachment.
[**AddServiceChain**](ManagementPlaneApiApi.md#AddServiceChain) | **Post** /serviceinsertion/service-chains | Add Service Chain
[**AddServiceInsertionExcludeListMemberAddMember**](ManagementPlaneApiApi.md#AddServiceInsertionExcludeListMemberAddMember) | **Post** /serviceinsertion/excludelist?action&#x3D;add_member | Add a new member in the exclude list
[**AddServiceInsertionRuleInSection**](ManagementPlaneApiApi.md#AddServiceInsertionRuleInSection) | **Post** /serviceinsertion/sections/{section-id}/rules | Add a Single Rule in a Section
[**AddServiceInsertionRulesInSectionCreateMultiple**](ManagementPlaneApiApi.md#AddServiceInsertionRulesInSectionCreateMultiple) | **Post** /serviceinsertion/sections/{section-id}/rules?action&#x3D;create_multiple | Add Multiple Rules in a Section
[**AddServiceInsertionSection**](ManagementPlaneApiApi.md#AddServiceInsertionSection) | **Post** /serviceinsertion/sections | Create a New Empty Section
[**AddServiceInsertionSectionWithRulesCreateWithRules**](ManagementPlaneApiApi.md#AddServiceInsertionSectionWithRulesCreateWithRules) | **Post** /serviceinsertion/sections?action&#x3D;create_with_rules | Create a Section with Rules
[**AddServiceInsertionService**](ManagementPlaneApiApi.md#AddServiceInsertionService) | **Post** /serviceinsertion/services | Create a Service-Insertion Service
[**AddServiceInstance**](ManagementPlaneApiApi.md#AddServiceInstance) | **Post** /serviceinsertion/services/{service-id}/service-instances | Add a Service Instance for a specified Service.
[**AddStaticRoute**](ManagementPlaneApiApi.md#AddStaticRoute) | **Post** /logical-routers/{logical-router-id}/routing/static-routes | Add Static Routes on a Logical Router
[**AddVendorTemplate**](ManagementPlaneApiApi.md#AddVendorTemplate) | **Post** /serviceinsertion/services/{service-id}/vendor-templates | Add Vendor Template for a given Service
[**CheckMemberIfExistsCheckIfExists**](ManagementPlaneApiApi.md#CheckMemberIfExistsCheckIfExists) | **Post** /firewall/excludelist?action&#x3D;check_if_exists | Check if the object a member of the exclude list
[**ClearDnsForwarderCacheClearCache**](ManagementPlaneApiApi.md#ClearDnsForwarderCacheClearCache) | **Post** /dns/forwarders/{forwarder-id}?action&#x3D;clear_cache | Clear the current cache of the DNS forwarder.
[**CreateApplProxyActionRestart**](ManagementPlaneApiApi.md#CreateApplProxyActionRestart) | **Post** /node/services/applianceproxy?action&#x3D;restart | Restart, start or stop the Appliance Proxy Service
[**CreateApplProxyActionStart**](ManagementPlaneApiApi.md#CreateApplProxyActionStart) | **Post** /node/services/applianceproxy?action&#x3D;start | Restart, start or stop the Appliance Proxy Service
[**CreateApplProxyActionStop**](ManagementPlaneApiApi.md#CreateApplProxyActionStop) | **Post** /node/services/applianceproxy?action&#x3D;stop | Restart, start or stop the Appliance Proxy Service
[**CreateBGPCommunityList**](ManagementPlaneApiApi.md#CreateBGPCommunityList) | **Post** /logical-routers/{logical-router-id}/routing/bgp/community-lists | Create a new BGP community list on a logical router
[**CreateBridgeCluster**](ManagementPlaneApiApi.md#CreateBridgeCluster) | **Post** /bridge-clusters | Create a Bridge Cluster
[**CreateBridgeEndpoint**](ManagementPlaneApiApi.md#CreateBridgeEndpoint) | **Post** /bridge-endpoints | Create a Bridge Endpoint
[**CreateBridgeEndpointProfile**](ManagementPlaneApiApi.md#CreateBridgeEndpointProfile) | **Post** /bridge-endpoint-profiles | Create a Bridge Endpoint Profile
[**CreateDADProfile**](ManagementPlaneApiApi.md#CreateDADProfile) | **Post** /ipv6/dad-profiles | Create a new DADProfile
[**CreateDhcpIpPool**](ManagementPlaneApiApi.md#CreateDhcpIpPool) | **Post** /dhcp/servers/{server-id}/ip-pools | Create an ip pool for a DHCP server
[**CreateDhcpProfile**](ManagementPlaneApiApi.md#CreateDhcpProfile) | **Post** /dhcp/server-profiles | Create a DHCP server profile
[**CreateDhcpRelay**](ManagementPlaneApiApi.md#CreateDhcpRelay) | **Post** /dhcp/relays | Create a DHCP Relay Service
[**CreateDhcpRelayProfile**](ManagementPlaneApiApi.md#CreateDhcpRelayProfile) | **Post** /dhcp/relay-profiles | Create a DHCP Relay Profile
[**CreateDhcpServer**](ManagementPlaneApiApi.md#CreateDhcpServer) | **Post** /dhcp/servers | Create a DHCP server
[**CreateDhcpStaticBinding**](ManagementPlaneApiApi.md#CreateDhcpStaticBinding) | **Post** /dhcp/servers/{server-id}/static-bindings | Create a static binding for a DHCP server
[**CreateDhcpV6IpPool**](ManagementPlaneApiApi.md#CreateDhcpV6IpPool) | **Post** /dhcp/servers/{server-id}/ipv6-ip-pools | Create an ip pool for a DHCP IPv6 server
[**CreateDhcpV6StaticBinding**](ManagementPlaneApiApi.md#CreateDhcpV6StaticBinding) | **Post** /dhcp/servers/{server-id}/ipv6-static-bindings | Create a static binding for a DHCP IPv6 server
[**CreateDnsForwader**](ManagementPlaneApiApi.md#CreateDnsForwader) | **Post** /dns/forwarders | Create a DNS forwader
[**CreateFirewallProfile**](ManagementPlaneApiApi.md#CreateFirewallProfile) | **Post** /firewall/profiles | Create a firewall profile.
[**CreateIPSecVPNDPDProfile**](ManagementPlaneApiApi.md#CreateIPSecVPNDPDProfile) | **Post** /vpn/ipsec/dpd-profiles | Create dead peer detection (DPD) profile
[**CreateIPSecVPNIKEProfile**](ManagementPlaneApiApi.md#CreateIPSecVPNIKEProfile) | **Post** /vpn/ipsec/ike-profiles | Create custom internet key exchange (IKE) Profile
[**CreateIPSecVPNLocalEndpoint**](ManagementPlaneApiApi.md#CreateIPSecVPNLocalEndpoint) | **Post** /vpn/ipsec/local-endpoints | Create custom local endpoint
[**CreateIPSecVPNPeerEndPoint**](ManagementPlaneApiApi.md#CreateIPSecVPNPeerEndPoint) | **Post** /vpn/ipsec/peer-endpoints | Create custom peer endpoint
[**CreateIPSecVPNService**](ManagementPlaneApiApi.md#CreateIPSecVPNService) | **Post** /vpn/ipsec/services | Create VPN service
[**CreateIPSecVPNSession**](ManagementPlaneApiApi.md#CreateIPSecVPNSession) | **Post** /vpn/ipsec/sessions | Create new VPN session
[**CreateIPSecVPNTunnelProfile**](ManagementPlaneApiApi.md#CreateIPSecVPNTunnelProfile) | **Post** /vpn/ipsec/tunnel-profiles | Create custom IPSec tunnel profile
[**CreateIPSet**](ManagementPlaneApiApi.md#CreateIPSet) | **Post** /ip-sets | Create IPSet
[**CreateIpfixCollectorConfig**](ManagementPlaneApiApi.md#CreateIpfixCollectorConfig) | **Post** /ipfix/collectorconfigs | Create a new IPFIX collector configuration
[**CreateIpfixCollectorUpmProfile**](ManagementPlaneApiApi.md#CreateIpfixCollectorUpmProfile) | **Post** /ipfix-collector-profiles | Create a new IPFIX collector profile
[**CreateIpfixConfig**](ManagementPlaneApiApi.md#CreateIpfixConfig) | **Post** /ipfix/configs | Create a new IPFIX configuration
[**CreateIpfixUpmProfile**](ManagementPlaneApiApi.md#CreateIpfixUpmProfile) | **Post** /ipfix-profiles | Create a new IPFIX profile
[**CreateL2VpnService**](ManagementPlaneApiApi.md#CreateL2VpnService) | **Post** /vpn/l2vpn/services | Create L2VPN service
[**CreateL2VpnSession**](ManagementPlaneApiApi.md#CreateL2VpnSession) | **Post** /vpn/l2vpn/sessions | Create L2VPN session
[**CreateLoadBalancerApplicationProfile**](ManagementPlaneApiApi.md#CreateLoadBalancerApplicationProfile) | **Post** /loadbalancer/application-profiles | Create a load balancer application profile
[**CreateLoadBalancerClientSslProfile**](ManagementPlaneApiApi.md#CreateLoadBalancerClientSslProfile) | **Post** /loadbalancer/client-ssl-profiles | Create a load balancer client-ssl profile
[**CreateLoadBalancerMonitor**](ManagementPlaneApiApi.md#CreateLoadBalancerMonitor) | **Post** /loadbalancer/monitors | Create a load balancer monitor
[**CreateLoadBalancerPersistenceProfile**](ManagementPlaneApiApi.md#CreateLoadBalancerPersistenceProfile) | **Post** /loadbalancer/persistence-profiles | Create a load balancer persistence profile
[**CreateLoadBalancerPool**](ManagementPlaneApiApi.md#CreateLoadBalancerPool) | **Post** /loadbalancer/pools | Create a load balancer pool
[**CreateLoadBalancerRule**](ManagementPlaneApiApi.md#CreateLoadBalancerRule) | **Post** /loadbalancer/rules | Create a load balancer rule
[**CreateLoadBalancerServerSslProfile**](ManagementPlaneApiApi.md#CreateLoadBalancerServerSslProfile) | **Post** /loadbalancer/server-ssl-profiles | Create a load balancer server-ssl profile
[**CreateLoadBalancerService**](ManagementPlaneApiApi.md#CreateLoadBalancerService) | **Post** /loadbalancer/services | Create a load balancer service
[**CreateLoadBalancerTcpProfile**](ManagementPlaneApiApi.md#CreateLoadBalancerTcpProfile) | **Post** /loadbalancer/tcp-profiles | Create a load balancer TCP profile
[**CreateLoadBalancerVirtualServer**](ManagementPlaneApiApi.md#CreateLoadBalancerVirtualServer) | **Post** /loadbalancer/virtual-servers | Create a load balancer virtual server
[**CreateLoadBalancerVirtualServerWithRulesCreateWithRules**](ManagementPlaneApiApi.md#CreateLoadBalancerVirtualServerWithRulesCreateWithRules) | **Post** /loadbalancer/virtual-servers?action&#x3D;create_with_rules | Create a load balancer virtual server with rules
[**CreateLogicalPort**](ManagementPlaneApiApi.md#CreateLogicalPort) | **Post** /logical-ports | Create a Logical Port
[**CreateLogicalRouter**](ManagementPlaneApiApi.md#CreateLogicalRouter) | **Post** /logical-routers | Create a Logical Router
[**CreateLogicalRouterPort**](ManagementPlaneApiApi.md#CreateLogicalRouterPort) | **Post** /logical-router-ports | Create a Logical Router Port
[**CreateLogicalSwitch**](ManagementPlaneApiApi.md#CreateLogicalSwitch) | **Post** /logical-switches | Create a Logical Switch
[**CreateMACSet**](ManagementPlaneApiApi.md#CreateMACSet) | **Post** /mac-sets | Create MACSet
[**CreateManualHealthCheck**](ManagementPlaneApiApi.md#CreateManualHealthCheck) | **Post** /manual-health-checks | Create a new manual health check request
[**CreateMetadataProxy**](ManagementPlaneApiApi.md#CreateMetadataProxy) | **Post** /md-proxies | Create a metadata proxy
[**CreateNDRAProfile**](ManagementPlaneApiApi.md#CreateNDRAProfile) | **Post** /ipv6/nd-ra-profiles | Create a new NDRA Profile
[**CreateNSGroup**](ManagementPlaneApiApi.md#CreateNSGroup) | **Post** /ns-groups | Create NSGroup
[**CreateNSProfile**](ManagementPlaneApiApi.md#CreateNSProfile) | **Post** /ns-profiles | Create NSProfile
[**CreateNSService**](ManagementPlaneApiApi.md#CreateNSService) | **Post** /ns-services | Create NSService
[**CreateNSServiceGroup**](ManagementPlaneApiApi.md#CreateNSServiceGroup) | **Post** /ns-service-groups | Create NSServiceGroup
[**CreatePacketCaptureSession**](ManagementPlaneApiApi.md#CreatePacketCaptureSession) | **Post** /pktcap/session | Create an new packet capture session
[**CreatePortMirroringSessions**](ManagementPlaneApiApi.md#CreatePortMirroringSessions) | **Post** /mirror-sessions | Create a mirror session
[**CreateServiceConfig**](ManagementPlaneApiApi.md#CreateServiceConfig) | **Post** /service-configs | Create service config
[**CreateSolutionConfig**](ManagementPlaneApiApi.md#CreateSolutionConfig) | **Post** /serviceinsertion/services/{service-id}/solution-configs | Add Solution Config for a given Service
[**CreateStaticHopBfdPeer**](ManagementPlaneApiApi.md#CreateStaticHopBfdPeer) | **Post** /logical-routers/{logical-router-id}/routing/static-routes/bfd-peers | Create a static hop BFD peer
[**CreateSwitchingProfile**](ManagementPlaneApiApi.md#CreateSwitchingProfile) | **Post** /switching-profiles | Create a Switching Profile
[**CreateTraceflow**](ManagementPlaneApiApi.md#CreateTraceflow) | **Post** /traceflows | Initiate a Traceflow Operation on the Specified Port
[**DeleteADhcpLease**](ManagementPlaneApiApi.md#DeleteADhcpLease) | **Delete** /dhcp/servers/{server-id}/leases | Delete a single DHCP lease entry specified by ip and mac.
[**DeleteAllCaptureSessionsDelete**](ManagementPlaneApiApi.md#DeleteAllCaptureSessionsDelete) | **Post** /pktcap/sessions?action&#x3D;delete | Delete all the packet capture sessions
[**DeleteBGPCommunityList**](ManagementPlaneApiApi.md#DeleteBGPCommunityList) | **Delete** /logical-routers/{logical-router-id}/routing/bgp/community-lists/{community-list-id} | Delete a specific BGP community list from a logical router
[**DeleteBgpNeighbor**](ManagementPlaneApiApi.md#DeleteBgpNeighbor) | **Delete** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{id} | Delete a specific BGP Neighbor on a Logical Router
[**DeleteBridgeCluster**](ManagementPlaneApiApi.md#DeleteBridgeCluster) | **Delete** /bridge-clusters/{bridgecluster-id} | Delete a Bridge Cluster
[**DeleteBridgeEndpoint**](ManagementPlaneApiApi.md#DeleteBridgeEndpoint) | **Delete** /bridge-endpoints/{bridgeendpoint-id} | Delete a Bridge Endpoint
[**DeleteBridgeEndpointProfile**](ManagementPlaneApiApi.md#DeleteBridgeEndpointProfile) | **Delete** /bridge-endpoint-profiles/{bridgeendpointprofile-id} | Delete a Bridge Endpoint Profile
[**DeleteDADProfile**](ManagementPlaneApiApi.md#DeleteDADProfile) | **Delete** /ipv6/dad-profiles/{dad-profile-id} | Delete DAD Profile
[**DeleteDhcpIpPool**](ManagementPlaneApiApi.md#DeleteDhcpIpPool) | **Delete** /dhcp/servers/{server-id}/ip-pools/{pool-id} | Delete a DHCP server&#x27;s IP pool
[**DeleteDhcpProfile**](ManagementPlaneApiApi.md#DeleteDhcpProfile) | **Delete** /dhcp/server-profiles/{profile-id} | Delete a DHCP server profile
[**DeleteDhcpRelay**](ManagementPlaneApiApi.md#DeleteDhcpRelay) | **Delete** /dhcp/relays/{relay-id} | Delete a DHCP Relay Service
[**DeleteDhcpRelayProfile**](ManagementPlaneApiApi.md#DeleteDhcpRelayProfile) | **Delete** /dhcp/relay-profiles/{relay-profile-id} | Delete a DHCP Relay Profile
[**DeleteDhcpServer**](ManagementPlaneApiApi.md#DeleteDhcpServer) | **Delete** /dhcp/servers/{server-id} | Delete a DHCP server
[**DeleteDhcpStaticBinding**](ManagementPlaneApiApi.md#DeleteDhcpStaticBinding) | **Delete** /dhcp/servers/{server-id}/static-bindings/{binding-id} | Delete a static binding
[**DeleteDhcpV6IpPool**](ManagementPlaneApiApi.md#DeleteDhcpV6IpPool) | **Delete** /dhcp/servers/{server-id}/ipv6-ip-pools/{pool-id} | Delete a DHCP IPv6 server&#x27;s IP pool
[**DeleteDhcpV6StaticBinding**](ManagementPlaneApiApi.md#DeleteDhcpV6StaticBinding) | **Delete** /dhcp/servers/{server-id}/ipv6-static-bindings/{binding-id} | Delete a static binding for DHCP IPv6 server
[**DeleteDnsForwarder**](ManagementPlaneApiApi.md#DeleteDnsForwarder) | **Delete** /dns/forwarders/{forwarder-id} | Delete a specific DNS forwarder
[**DeleteFirewallProfile**](ManagementPlaneApiApi.md#DeleteFirewallProfile) | **Delete** /firewall/profiles/{profile-id} | Delete a firewall profile.
[**DeleteIPPrefixList**](ManagementPlaneApiApi.md#DeleteIPPrefixList) | **Delete** /logical-routers/{logical-router-id}/routing/ip-prefix-lists/{id} | Delete a specific IPPrefixList on a Logical Router
[**DeleteIPSecVPNDPDProfile**](ManagementPlaneApiApi.md#DeleteIPSecVPNDPDProfile) | **Delete** /vpn/ipsec/dpd-profiles/{ipsec-vpn-dpd-profile-id} | Delete dead peer detection (DPD) profile
[**DeleteIPSecVPNIKEProfile**](ManagementPlaneApiApi.md#DeleteIPSecVPNIKEProfile) | **Delete** /vpn/ipsec/ike-profiles/{ipsec-vpn-ike-profile-id} | Delete custom IKE Profile
[**DeleteIPSecVPNLocalEndpoint**](ManagementPlaneApiApi.md#DeleteIPSecVPNLocalEndpoint) | **Delete** /vpn/ipsec/local-endpoints/{ipsec-vpn-local-endpoint-id} | Delete custom IPSec local endpoint
[**DeleteIPSecVPNPeerEndpoint**](ManagementPlaneApiApi.md#DeleteIPSecVPNPeerEndpoint) | **Delete** /vpn/ipsec/peer-endpoints/{ipsec-vpn-peer-endpoint-id} | Delete custom IPSec VPN peer endpoint
[**DeleteIPSecVPNService**](ManagementPlaneApiApi.md#DeleteIPSecVPNService) | **Delete** /vpn/ipsec/services/{ipsec-vpn-service-id} | Delete IPSec VPN service
[**DeleteIPSecVPNSession**](ManagementPlaneApiApi.md#DeleteIPSecVPNSession) | **Delete** /vpn/ipsec/sessions/{ipsec-vpn-session-id} | Delete IPSec VPN session
[**DeleteIPSecVPNTunnelProfile**](ManagementPlaneApiApi.md#DeleteIPSecVPNTunnelProfile) | **Delete** /vpn/ipsec/tunnel-profiles/{ipsec-vpn-tunnel-profile-id} | Delete custom IPSecTunnelProfile
[**DeleteIPSet**](ManagementPlaneApiApi.md#DeleteIPSet) | **Delete** /ip-sets/{ip-set-id} | Delete IPSet
[**DeleteInstanceEndpoint**](ManagementPlaneApiApi.md#DeleteInstanceEndpoint) | **Delete** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-endpoints/{instance-endpoint-id} | Delete a particular InstanceEndpoint.
[**DeleteIpfixCollectorConfig**](ManagementPlaneApiApi.md#DeleteIpfixCollectorConfig) | **Delete** /ipfix/collectorconfigs/{collector-config-id} | Delete an existing IPFIX collector configuration
[**DeleteIpfixCollectorUpmProfile**](ManagementPlaneApiApi.md#DeleteIpfixCollectorUpmProfile) | **Delete** /ipfix-collector-profiles/{ipfix-collector-profile-id} | Delete an existing IPFIX collector profile
[**DeleteIpfixConfig**](ManagementPlaneApiApi.md#DeleteIpfixConfig) | **Delete** /ipfix/configs/{config-id} | Delete an existing IPFIX configuration
[**DeleteIpfixUpmProfile**](ManagementPlaneApiApi.md#DeleteIpfixUpmProfile) | **Delete** /ipfix-profiles/{ipfix-profile-id} | Delete an existing IPFIX profile
[**DeleteL2VpnService**](ManagementPlaneApiApi.md#DeleteL2VpnService) | **Delete** /vpn/l2vpn/services/{l2vpn-service-id} | Delete a L2VPN service
[**DeleteL2VpnSession**](ManagementPlaneApiApi.md#DeleteL2VpnSession) | **Delete** /vpn/l2vpn/sessions/{l2vpn-session-id} | Delete a L2VPN session
[**DeleteLoadBalancerApplicationProfile**](ManagementPlaneApiApi.md#DeleteLoadBalancerApplicationProfile) | **Delete** /loadbalancer/application-profiles/{application-profile-id} | Delete a load balancer application profile
[**DeleteLoadBalancerClientSslProfile**](ManagementPlaneApiApi.md#DeleteLoadBalancerClientSslProfile) | **Delete** /loadbalancer/client-ssl-profiles/{client-ssl-profile-id} | Delete a load balancer client-ssl profile
[**DeleteLoadBalancerMonitor**](ManagementPlaneApiApi.md#DeleteLoadBalancerMonitor) | **Delete** /loadbalancer/monitors/{monitor-id} | Delete a load balancer monitor
[**DeleteLoadBalancerPersistenceProfile**](ManagementPlaneApiApi.md#DeleteLoadBalancerPersistenceProfile) | **Delete** /loadbalancer/persistence-profiles/{persistence-profile-id} | Delete a load balancer persistence profile
[**DeleteLoadBalancerPool**](ManagementPlaneApiApi.md#DeleteLoadBalancerPool) | **Delete** /loadbalancer/pools/{pool-id} | Delete a load balancer pool
[**DeleteLoadBalancerRule**](ManagementPlaneApiApi.md#DeleteLoadBalancerRule) | **Delete** /loadbalancer/rules/{rule-id} | Delete a load balancer rule
[**DeleteLoadBalancerServerSslProfile**](ManagementPlaneApiApi.md#DeleteLoadBalancerServerSslProfile) | **Delete** /loadbalancer/server-ssl-profiles/{server-ssl-profile-id} | Delete a load balancer server-ssl profile
[**DeleteLoadBalancerService**](ManagementPlaneApiApi.md#DeleteLoadBalancerService) | **Delete** /loadbalancer/services/{service-id} | Delete a load balancer service
[**DeleteLoadBalancerTcpProfile**](ManagementPlaneApiApi.md#DeleteLoadBalancerTcpProfile) | **Delete** /loadbalancer/tcp-profiles/{tcp-profile-id} | Delete a load balancer TCP profile
[**DeleteLoadBalancerVirtualServer**](ManagementPlaneApiApi.md#DeleteLoadBalancerVirtualServer) | **Delete** /loadbalancer/virtual-servers/{virtual-server-id} | Delete a load balancer virtual server
[**DeleteLogicalPort**](ManagementPlaneApiApi.md#DeleteLogicalPort) | **Delete** /logical-ports/{lport-id} | Delete a Logical Port
[**DeleteLogicalRouter**](ManagementPlaneApiApi.md#DeleteLogicalRouter) | **Delete** /logical-routers/{logical-router-id} | Delete a Logical Router
[**DeleteLogicalRouterPort**](ManagementPlaneApiApi.md#DeleteLogicalRouterPort) | **Delete** /logical-router-ports/{logical-router-port-id} | Delete a Logical Router Port
[**DeleteLogicalSwitch**](ManagementPlaneApiApi.md#DeleteLogicalSwitch) | **Delete** /logical-switches/{lswitch-id} | Delete a Logical Switch
[**DeleteMACSet**](ManagementPlaneApiApi.md#DeleteMACSet) | **Delete** /mac-sets/{mac-set-id} | Delete MACSet
[**DeleteManualHealthCheck**](ManagementPlaneApiApi.md#DeleteManualHealthCheck) | **Delete** /manual-health-checks/{manual-health-check-id} | Delete an existing manual health check
[**DeleteMetadataProxy**](ManagementPlaneApiApi.md#DeleteMetadataProxy) | **Delete** /md-proxies/{proxy-id} | Delete a metadata proxy
[**DeleteNDRAProfile**](ManagementPlaneApiApi.md#DeleteNDRAProfile) | **Delete** /ipv6/nd-ra-profiles/{nd-ra-profile-id} | Delete NDRA Profile
[**DeleteNSGroup**](ManagementPlaneApiApi.md#DeleteNSGroup) | **Delete** /ns-groups/{ns-group-id} | Delete NSGroup
[**DeleteNSProfile**](ManagementPlaneApiApi.md#DeleteNSProfile) | **Delete** /ns-profiles/{ns-profile-id} | Delete NSProfile
[**DeleteNSService**](ManagementPlaneApiApi.md#DeleteNSService) | **Delete** /ns-services/{ns-service-id} | Delete NSService
[**DeleteNSServiceGroup**](ManagementPlaneApiApi.md#DeleteNSServiceGroup) | **Delete** /ns-service-groups/{ns-service-group-id} | Delete NSServiceGroup
[**DeleteNatRule**](ManagementPlaneApiApi.md#DeleteNatRule) | **Delete** /logical-routers/{logical-router-id}/nat/rules/{rule-id} | Delete a specific NAT rule from a logical router
[**DeletePBRRule**](ManagementPlaneApiApi.md#DeletePBRRule) | **Delete** /pbr/sections/{section-id}/rules/{rule-id} | Delete an Existing Rule
[**DeletePBRSection**](ManagementPlaneApiApi.md#DeletePBRSection) | **Delete** /pbr/sections/{section-id} | Delete an Existing Section and Its Associated Rules
[**DeletePacketCaptureSessionDelete**](ManagementPlaneApiApi.md#DeletePacketCaptureSessionDelete) | **Post** /pktcap/session/{session-id}?action&#x3D;delete | Delete the packet capture session by session id.
[**DeletePortMirroringSession**](ManagementPlaneApiApi.md#DeletePortMirroringSession) | **Delete** /mirror-sessions/{mirror-session-id} | Delete the mirror session
[**DeleteRouteMap**](ManagementPlaneApiApi.md#DeleteRouteMap) | **Delete** /logical-routers/{logical-router-id}/routing/route-maps/{id} | Delete a specific RouteMap on a Logical Router
[**DeleteRule**](ManagementPlaneApiApi.md#DeleteRule) | **Delete** /firewall/sections/{section-id}/rules/{rule-id} | Delete an Existing Rule
[**DeleteSIServiceProfile**](ManagementPlaneApiApi.md#DeleteSIServiceProfile) | **Delete** /serviceinsertion/services/{service-id}/service-profiles/{service-profile-id} | Delete a particular ServiceProfile.
[**DeleteSection**](ManagementPlaneApiApi.md#DeleteSection) | **Delete** /firewall/sections/{section-id} | Delete an Existing Section and Its Associated Rules
[**DeleteServiceAttachment**](ManagementPlaneApiApi.md#DeleteServiceAttachment) | **Delete** /serviceinsertion/service-attachments/{service-attachment-id} | Delete an existing service attachment
[**DeleteServiceChain**](ManagementPlaneApiApi.md#DeleteServiceChain) | **Delete** /serviceinsertion/service-chains/{service-chain-id} | Delete a Service Chain.
[**DeleteServiceConfig**](ManagementPlaneApiApi.md#DeleteServiceConfig) | **Delete** /service-configs/{config-set-id} | Delete Service Config
[**DeleteServiceDeployment**](ManagementPlaneApiApi.md#DeleteServiceDeployment) | **Delete** /serviceinsertion/services/{service-id}/service-deployments/{service-deployment-id} | Remove service deployment
[**DeleteServiceInsertionRule**](ManagementPlaneApiApi.md#DeleteServiceInsertionRule) | **Delete** /serviceinsertion/sections/{section-id}/rules/{rule-id} | Delete an Existing Rule
[**DeleteServiceInsertionSection**](ManagementPlaneApiApi.md#DeleteServiceInsertionSection) | **Delete** /serviceinsertion/sections/{section-id} | Delete an Existing Section and Its Associated Rules
[**DeleteServiceInsertionService**](ManagementPlaneApiApi.md#DeleteServiceInsertionService) | **Delete** /serviceinsertion/services/{service-id} | Delete an existing Service and the Service-Instance associated with it.
[**DeleteServiceInstance**](ManagementPlaneApiApi.md#DeleteServiceInstance) | **Delete** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id} | Delete an existing Service-Instance
[**DeleteServiceManager**](ManagementPlaneApiApi.md#DeleteServiceManager) | **Delete** /serviceinsertion/service-managers/{service-manager-id} | Delete service manager
[**DeleteServiceVMsDelete**](ManagementPlaneApiApi.md#DeleteServiceVMsDelete) | **Post** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-runtimes?action&#x3D;delete | Remove service VMs either as standalone or HA
[**DeleteSolutionConfig**](ManagementPlaneApiApi.md#DeleteSolutionConfig) | **Delete** /serviceinsertion/services/{service-id}/solution-configs/{solution-config-id} | Deletes solution config information.
[**DeleteStaticHopBfdPeer**](ManagementPlaneApiApi.md#DeleteStaticHopBfdPeer) | **Delete** /logical-routers/{logical-router-id}/routing/static-routes/bfd-peers/{bfd-peer-id} | Delete a specified static route BFD peer cofigured on a specified logical router
[**DeleteStaticRoute**](ManagementPlaneApiApi.md#DeleteStaticRoute) | **Delete** /logical-routers/{logical-router-id}/routing/static-routes/{id} | Delete a specific Static Route on a Logical Router
[**DeleteSwitchingProfile**](ManagementPlaneApiApi.md#DeleteSwitchingProfile) | **Delete** /switching-profiles/{switching-profile-id} | Delete a Switching Profile
[**DeleteTraceflow**](ManagementPlaneApiApi.md#DeleteTraceflow) | **Delete** /traceflows/{traceflow-id} | Delete the Traceflow round
[**DeleteVendorTemplate**](ManagementPlaneApiApi.md#DeleteVendorTemplate) | **Delete** /serviceinsertion/services/{service-id}/vendor-templates/{vendor-template-id} | Delete a particular vendor tempalte.
[**DeployService**](ManagementPlaneApiApi.md#DeployService) | **Post** /serviceinsertion/services/{service-id}/service-deployments | Deploys a particular service
[**DeployServiceVMsDeploy**](ManagementPlaneApiApi.md#DeployServiceVMsDeploy) | **Post** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-runtimes?action&#x3D;deploy | Deploy and set up service VMs either as standalone or HA
[**DisableDnsForwarderDisable**](ManagementPlaneApiApi.md#DisableDnsForwarderDisable) | **Post** /dns/forwarders/{forwarder-id}?action&#x3D;disable | Disable the DNS forwarder.
[**DisableFirewallOnTargetResourceDisableFirewall**](ManagementPlaneApiApi.md#DisableFirewallOnTargetResourceDisableFirewall) | **Post** /firewall/status/{context-type}/{id}?action&#x3D;disable_firewall | Disable firewall on target resource in dfw context
[**EffectiveProfiles**](ManagementPlaneApiApi.md#EffectiveProfiles) | **Get** /service-configs/effective-profiles | Get Effective Profiles for an Entity
[**EnableDnsForwarderEnable**](ManagementPlaneApiApi.md#EnableDnsForwarderEnable) | **Post** /dns/forwarders/{forwarder-id}?action&#x3D;enable | Enable the DNS forwarder.
[**EnableFirewallOnTargetResourceEnableFirewall**](ManagementPlaneApiApi.md#EnableFirewallOnTargetResourceEnableFirewall) | **Post** /firewall/status/{context-type}/{id}?action&#x3D;enable_firewall | Enable firewall on target resource in dfw context
[**GetAffectedVms**](ManagementPlaneApiApi.md#GetAffectedVms) | **Post** /intrusion-services/affected-vms | Get the list of the VMs affected for that signature
[**GetAllIdsEvents**](ManagementPlaneApiApi.md#GetAllIdsEvents) | **Post** /intrusion-services/ids-events | Get the list of the IDS events that are detected, grouped by signature id.
[**GetAssociations**](ManagementPlaneApiApi.md#GetAssociations) | **Get** /associations | Get ResourceReference objects to which the given resource belongs to 
[**GetBgpNeighborAdvertisedRoutes**](ManagementPlaneApiApi.md#GetBgpNeighborAdvertisedRoutes) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{neighbor-id}/advertised-routes | Get BGP neighbor advertised routes
[**GetBgpNeighborAdvertisedRoutesInCsvFormatCsv**](ManagementPlaneApiApi.md#GetBgpNeighborAdvertisedRoutesInCsvFormatCsv) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{neighbor-id}/advertised-routes?format&#x3D;csv | Get BGP neighbor advertised routes in CSV format 
[**GetBgpNeighborRoutes**](ManagementPlaneApiApi.md#GetBgpNeighborRoutes) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{neighbor-id}/routes | Get BGP neighbor learned routes
[**GetBgpNeighborRoutesInCsvFormatCsv**](ManagementPlaneApiApi.md#GetBgpNeighborRoutesInCsvFormatCsv) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{neighbor-id}/routes?format&#x3D;csv | Get BGP neighbor learned routes in CSV format 
[**GetBgpNeighborsStatus**](ManagementPlaneApiApi.md#GetBgpNeighborsStatus) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors/status | Get the status of all the BGP neighbors for the Logical Router of the given id
[**GetBridgeCluster**](ManagementPlaneApiApi.md#GetBridgeCluster) | **Get** /bridge-clusters/{bridgecluster-id} | Get Information about a bridge cluster
[**GetBridgeClusterStatus**](ManagementPlaneApiApi.md#GetBridgeClusterStatus) | **Get** /bridge-clusters/{cluster-id}/status | Returns status of a specified Bridge Cluster
[**GetBridgeEndpoint**](ManagementPlaneApiApi.md#GetBridgeEndpoint) | **Get** /bridge-endpoints/{bridgeendpoint-id} | Get Information about a bridge endpoint
[**GetBridgeEndpointProfile**](ManagementPlaneApiApi.md#GetBridgeEndpointProfile) | **Get** /bridge-endpoint-profiles/{bridgeendpointprofile-id} | Get Information about a bridge endpoint Profile
[**GetBridgeEndpointStatistics**](ManagementPlaneApiApi.md#GetBridgeEndpointStatistics) | **Get** /bridge-endpoints/{endpoint-id}/statistics | Returns statistics of a specified Bridge Endpoint
[**GetBridgeEndpointStatus**](ManagementPlaneApiApi.md#GetBridgeEndpointStatus) | **Get** /bridge-endpoints/{endpoint-id}/status | Returns status of a specified Bridge Endpoint
[**GetCaptureFile**](ManagementPlaneApiApi.md#GetCaptureFile) | **Get** /pktcap/session/{session-id}/capturefile | Get packet capture file
[**GetComputeCollectionStatusById**](ManagementPlaneApiApi.md#GetComputeCollectionStatusById) | **Get** /idfw/compute-collections/{compute-collection-ext-id}/status | Get list of compute collections and status.
[**GetDhcpIpPoolState**](ManagementPlaneApiApi.md#GetDhcpIpPoolState) | **Get** /dhcp/servers/{server-id}/ip-pools/{pool-id}/state | Get the realized state of a dhcp ip pool
[**GetDhcpLeaseInfo**](ManagementPlaneApiApi.md#GetDhcpLeaseInfo) | **Get** /dhcp/servers/{server-id}/leases | Get specific leases of a given dhcp server
[**GetDhcpServerState**](ManagementPlaneApiApi.md#GetDhcpServerState) | **Get** /dhcp/servers/{server-id}/state | Get the realized state of a dhcp server
[**GetDhcpStaticBindingState**](ManagementPlaneApiApi.md#GetDhcpStaticBindingState) | **Get** /dhcp/servers/{server-id}/static-bindings/{binding-id}/state | Get the realized state of a dhcp static binding
[**GetDhcpStatistics**](ManagementPlaneApiApi.md#GetDhcpStatistics) | **Get** /dhcp/servers/{server-id}/statistics | Get DHCP statistics with given dhcp server id
[**GetDhcpStatus**](ManagementPlaneApiApi.md#GetDhcpStatus) | **Get** /dhcp/servers/{server-id}/status | Get DHCP service status with given dhcp server id
[**GetDnsForwarderState**](ManagementPlaneApiApi.md#GetDnsForwarderState) | **Get** /dns/forwarders/{forwarder-id}/state | Get the realized state of a DNS forwarder
[**GetDnsForwarderStatistics**](ManagementPlaneApiApi.md#GetDnsForwarderStatistics) | **Get** /dns/forwarders/{forwarder-id}/statistics | Get statistics of given dns forwarder
[**GetDnsForwarderStatus**](ManagementPlaneApiApi.md#GetDnsForwarderStatus) | **Get** /dns/forwarders/{forwarder-id}/status | Get current status of the given DNS forwarder
[**GetEffectiveActiveDirectoryGroups**](ManagementPlaneApiApi.md#GetEffectiveActiveDirectoryGroups) | **Get** /ns-groups/{ns-group-id}/effective-directory-group-members | Get Effective Directory Groups of the specified NSGroup.
[**GetEffectiveCloudNativeServiceInstances**](ManagementPlaneApiApi.md#GetEffectiveCloudNativeServiceInstances) | **Get** /ns-groups/{ns-group-id}/effective-cloud-native-service-instance-members | Get Effective Cloud Native Service Instances of the specified NSGroup.
[**GetEffectiveIPAddressMembers**](ManagementPlaneApiApi.md#GetEffectiveIPAddressMembers) | **Get** /ns-groups/{ns-group-id}/effective-ip-address-members | Get Effective IPAddress translated from the NSGroup
[**GetEffectiveIPSetMembers**](ManagementPlaneApiApi.md#GetEffectiveIPSetMembers) | **Get** /ns-groups/{ns-group-id}/effective-ipset-members | Get Effective IPSets of the specified NSGroup.
[**GetEffectiveLogicalPortMembers**](ManagementPlaneApiApi.md#GetEffectiveLogicalPortMembers) | **Get** /ns-groups/{ns-group-id}/effective-logical-port-members | Get Effective Logical Ports translated from the NSgroup
[**GetEffectiveLogicalSwitchMembers**](ManagementPlaneApiApi.md#GetEffectiveLogicalSwitchMembers) | **Get** /ns-groups/{ns-group-id}/effective-logical-switch-members | Get Effective switch members translated from the NSGroup
[**GetEffectivePhysicalServerMembers**](ManagementPlaneApiApi.md#GetEffectivePhysicalServerMembers) | **Get** /ns-groups/{ns-group-id}/effective-physical-server-members | Get Effective Physical Server Memebers of the specified NSGroup.
[**GetEffectiveTransportNodeMembers**](ManagementPlaneApiApi.md#GetEffectiveTransportNodeMembers) | **Get** /ns-groups/{ns-group-id}/effective-transport-node-members | Get effective transport node members translated from the NSGroup
[**GetEffectiveVIFMembers**](ManagementPlaneApiApi.md#GetEffectiveVIFMembers) | **Get** /ns-groups/{ns-group-id}/effective-vif-members | Get effective VIF members translated from the NSGroup
[**GetEffectiveVirtualMachineMembers**](ManagementPlaneApiApi.md#GetEffectiveVirtualMachineMembers) | **Get** /ns-groups/{ns-group-id}/effective-virtual-machine-members | Get Effective Virtual Machine members of the specified NSGroup.
[**GetEnabledComputeCollection**](ManagementPlaneApiApi.md#GetEnabledComputeCollection) | **Get** /idfw/idfw-compute-collections/{cc-ext-id} | Get IDFW compute collection.
[**GetExcludeList**](ManagementPlaneApiApi.md#GetExcludeList) | **Get** /firewall/excludelist | Get list of entities in exclude list
[**GetFailedDnsQueries**](ManagementPlaneApiApi.md#GetFailedDnsQueries) | **Get** /dns/forwarders/{forwarder-id}/failed-queries | Get the recent failed DNS queries
[**GetFirewallProfile**](ManagementPlaneApiApi.md#GetFirewallProfile) | **Get** /firewall/profiles/{profile-id} | Get all firewall session timer profiles.
[**GetFirewallSectionStats**](ManagementPlaneApiApi.md#GetFirewallSectionStats) | **Get** /firewall/sections/{section-id}/rules/stats | Get Firewall section level statistics section
[**GetFirewallStats**](ManagementPlaneApiApi.md#GetFirewallStats) | **Get** /firewall/sections/{section-id}/rules/{rule-id}/stats | Get Firewall rule level statistics
[**GetFirewallStatus**](ManagementPlaneApiApi.md#GetFirewallStatus) | **Get** /firewall/status/{context-type} | Get firewall global status for dfw context
[**GetFirewallStatusOnTargetResource**](ManagementPlaneApiApi.md#GetFirewallStatusOnTargetResource) | **Get** /firewall/status/{context-type}/{id} | Get firewall status for target resource in dfw context
[**GetForwardingPath**](ManagementPlaneApiApi.md#GetForwardingPath) | **Get** /logical-ports/{lport-id}/forwarding-path | Get networking entities between two logical ports with VIF attachment
[**GetIDSProfile**](ManagementPlaneApiApi.md#GetIDSProfile) | **Get** /intrusion-services/profiles/{ids-profile-id} | Get IDSProfile
[**GetIPAddresses**](ManagementPlaneApiApi.md#GetIPAddresses) | **Get** /ip-sets/{ip-set-id}/members | Get all IPAddresses in a IPSet
[**GetIPSecVPNDPDProfile**](ManagementPlaneApiApi.md#GetIPSecVPNDPDProfile) | **Get** /vpn/ipsec/dpd-profiles/{ipsec-vpn-dpd-profile-id} | Get IPSec dead peer detection (DPD) profile
[**GetIPSecVPNIKEProfile**](ManagementPlaneApiApi.md#GetIPSecVPNIKEProfile) | **Get** /vpn/ipsec/ike-profiles/{ipsec-vpn-ike-profile-id} | Get IKE Profile
[**GetIPSecVPNIKEService**](ManagementPlaneApiApi.md#GetIPSecVPNIKEService) | **Get** /vpn/services/{service-id}/summary | Cumulative statistics for one IKE service instance
[**GetIPSecVPNIKESessionStatus**](ManagementPlaneApiApi.md#GetIPSecVPNIKESessionStatus) | **Get** /vpn/ipsec/sessions/{session-id}/status | Get IPSec VPN IKE session status
[**GetIPSecVPNLocalEndpoint**](ManagementPlaneApiApi.md#GetIPSecVPNLocalEndpoint) | **Get** /vpn/ipsec/local-endpoints/{ipsec-vpn-local-endpoint-id} | Get custom IPSec local endpoint
[**GetIPSecVPNPeerEndpoint**](ManagementPlaneApiApi.md#GetIPSecVPNPeerEndpoint) | **Get** /vpn/ipsec/peer-endpoints/{ipsec-vpn-peer-endpoint-id} | Get IPSec VPN peer endpoint
[**GetIPSecVPNPeerEndpointWithPSKShowSensitiveData**](ManagementPlaneApiApi.md#GetIPSecVPNPeerEndpointWithPSKShowSensitiveData) | **Get** /vpn/ipsec/peer-endpoints/{ipsec-vpn-peer-endpoint-id}?action&#x3D;show-sensitive-data | Get IPSec VPN peer endpoint with PSK
[**GetIPSecVPNService**](ManagementPlaneApiApi.md#GetIPSecVPNService) | **Get** /vpn/ipsec/services/{ipsec-vpn-service-id} | Get IPSec VPN service
[**GetIPSecVPNSession**](ManagementPlaneApiApi.md#GetIPSecVPNSession) | **Get** /vpn/ipsec/sessions/{ipsec-vpn-session-id} | Fetch IPSec VPN session
[**GetIPSecVPNSessionState**](ManagementPlaneApiApi.md#GetIPSecVPNSessionState) | **Get** /vpn/ipsec/sessions/{ipsec-vpn-session-id}/state | Get the Realized State of a IPSec VPN Session
[**GetIPSecVPNSessionStatistics**](ManagementPlaneApiApi.md#GetIPSecVPNSessionStatistics) | **Get** /vpn/ipsec/sessions/{session-id}/statistics | Get IPSec VPN session statistics
[**GetIPSecVPNSessionSummary**](ManagementPlaneApiApi.md#GetIPSecVPNSessionSummary) | **Get** /vpn/ipsec/sessions/summary | VPN session summary
[**GetIPSecVPNTunnelProfile**](ManagementPlaneApiApi.md#GetIPSecVPNTunnelProfile) | **Get** /vpn/ipsec/tunnel-profiles/{ipsec-vpn-tunnel-profile-id} | Get IPSec tunnel profile
[**GetIdsDashboardSummary**](ManagementPlaneApiApi.md#GetIdsDashboardSummary) | **Post** /intrusion-services/ids-summary | Get the summary of the intrusions that were detected.
[**GetInstanceEndpoint**](ManagementPlaneApiApi.md#GetInstanceEndpoint) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-endpoints/{instance-endpoint-id} | Get a particular instance endpoint for a service instance.
[**GetIpfixCollectorConfig**](ManagementPlaneApiApi.md#GetIpfixCollectorConfig) | **Get** /ipfix/collectorconfigs/{collector-config-id} | Get an existing IPFIX collector configuration
[**GetIpfixCollectorUpmProfile**](ManagementPlaneApiApi.md#GetIpfixCollectorUpmProfile) | **Get** /ipfix-collector-profiles/{ipfix-collector-profile-id} | Get an existing IPFIX collector profile
[**GetIpfixConfig**](ManagementPlaneApiApi.md#GetIpfixConfig) | **Get** /ipfix/configs/{config-id} | Get an existing IPFIX configuration
[**GetIpfixObsPoints**](ManagementPlaneApiApi.md#GetIpfixObsPoints) | **Get** /ipfix-obs-points | Get the list of IPFIX observation points
[**GetIpfixUpmProfile**](ManagementPlaneApiApi.md#GetIpfixUpmProfile) | **Get** /ipfix-profiles/{ipfix-profile-id} | Get an existing IPFIX profile
[**GetL2ForwarderRemoteMacs**](ManagementPlaneApiApi.md#GetL2ForwarderRemoteMacs) | **Get** /logical-switches/{logical-switch-id}/inter-site-forwarder/site-span-info | Get L2 forwarder remote mac addresses
[**GetL2ForwarderStatistics**](ManagementPlaneApiApi.md#GetL2ForwarderStatistics) | **Get** /logical-switches/{logical-switch-id}/inter-site-forwarder/statistics | Get L2 forwarder statistics
[**GetL2ForwarderStatus**](ManagementPlaneApiApi.md#GetL2ForwarderStatus) | **Get** /logical-switches/{logical-switch-id}/inter-site-forwarder/status | Get L2 forwarder status
[**GetL2VPNSessionRemoteMacsForLS**](ManagementPlaneApiApi.md#GetL2VPNSessionRemoteMacsForLS) | **Get** /vpn/l2vpn/sessions/{session-id}/remote-mac | Get L2VPN session remote mac for logical switch
[**GetL2VPNSessionStatistics**](ManagementPlaneApiApi.md#GetL2VPNSessionStatistics) | **Get** /vpn/l2vpn/sessions/{session-id}/statistics | Get L2VPN session statistics
[**GetL2VPNSessionStatus**](ManagementPlaneApiApi.md#GetL2VPNSessionStatus) | **Get** /vpn/l2vpn/sessions/{session-id}/status | Get L2VPN session status
[**GetL2VPNSessionSummary**](ManagementPlaneApiApi.md#GetL2VPNSessionSummary) | **Get** /vpn/l2vpn/sessions/summary | Get status summary of all existing L2VPN sessions.
[**GetL2VpnService**](ManagementPlaneApiApi.md#GetL2VpnService) | **Get** /vpn/l2vpn/services/{l2vpn-service-id} | Get L2VPN service
[**GetL2VpnSession**](ManagementPlaneApiApi.md#GetL2VpnSession) | **Get** /vpn/l2vpn/sessions/{l2vpn-session-id} | Get a L2VPN session
[**GetL2VpnSessionPeerCodes**](ManagementPlaneApiApi.md#GetL2VpnSessionPeerCodes) | **Get** /vpn/l2vpn/sessions/{l2vpn-session-id}/peer-codes | Get peer codes for the L2VpnSession
[**GetLoadBalancerPoolStatistics**](ManagementPlaneApiApi.md#GetLoadBalancerPoolStatistics) | **Get** /loadbalancer/services/{service-id}/pools/{pool-id}/statistics | Get the statistics of load balancer pool
[**GetLoadBalancerPoolStatus**](ManagementPlaneApiApi.md#GetLoadBalancerPoolStatus) | **Get** /loadbalancer/services/{service-id}/pools/{pool-id}/status | Get the status of load balancer pool
[**GetLoadBalancerServiceStatistics**](ManagementPlaneApiApi.md#GetLoadBalancerServiceStatistics) | **Get** /loadbalancer/services/{service-id}/statistics | Get the statistics of load balancer service
[**GetLoadBalancerServiceStatus**](ManagementPlaneApiApi.md#GetLoadBalancerServiceStatus) | **Get** /loadbalancer/services/{service-id}/status | Get the status of the given load balancer service
[**GetLoadBalancerVirtualServerStatistics**](ManagementPlaneApiApi.md#GetLoadBalancerVirtualServerStatistics) | **Get** /loadbalancer/services/{service-id}/virtual-servers/{virtual-server-id}/statistics | Get the statistics of the given load balancer virtual server
[**GetLoadBalancerVirtualServerStatus**](ManagementPlaneApiApi.md#GetLoadBalancerVirtualServerStatus) | **Get** /loadbalancer/services/{service-id}/virtual-servers/{virtual-server-id}/status | Get the status of the load balancer virtual server
[**GetLogicalPort**](ManagementPlaneApiApi.md#GetLogicalPort) | **Get** /logical-ports/{lport-id} | Get Information About a Logical Port
[**GetLogicalPortMacTable**](ManagementPlaneApiApi.md#GetLogicalPortMacTable) | **Get** /logical-ports/{lport-id}/mac-table | Get MAC table of a logical port with a given port id (lport-id)
[**GetLogicalPortMacTableInCsvFormatCsv**](ManagementPlaneApiApi.md#GetLogicalPortMacTableInCsvFormatCsv) | **Get** /logical-ports/{lport-id}/mac-table?format&#x3D;csv | Get MAC table of a logical port with a given port id (lport-id)
[**GetLogicalPortOperationalStatus**](ManagementPlaneApiApi.md#GetLogicalPortOperationalStatus) | **Get** /logical-ports/{lport-id}/status | Get Operational Status for Logical Port of a Given Port ID (lport-id)
[**GetLogicalPortState**](ManagementPlaneApiApi.md#GetLogicalPortState) | **Get** /logical-ports/{lport-id}/state | Get realized state &amp; location of a logical port
[**GetLogicalPortStatistics**](ManagementPlaneApiApi.md#GetLogicalPortStatistics) | **Get** /logical-ports/{lport-id}/statistics | Get Statistics for Logical Port of a Given Port ID (lport-id)
[**GetLogicalPortStatusSummary**](ManagementPlaneApiApi.md#GetLogicalPortStatusSummary) | **Get** /logical-ports/status | Get Operational Status Summary of All Logical Ports in the System
[**GetLogicalRouterForwardingTable**](ManagementPlaneApiApi.md#GetLogicalRouterForwardingTable) | **Get** /logical-routers/{logical-router-id}/routing/forwarding-table | Get FIB table on a specified node for a logical router
[**GetLogicalRouterForwardingTableInCsvFormatCsv**](ManagementPlaneApiApi.md#GetLogicalRouterForwardingTableInCsvFormatCsv) | **Get** /logical-routers/{logical-router-id}/routing/forwarding-table?format&#x3D;csv | Get FIB table on a specified node for a logical router
[**GetLogicalRouterPortArpTable**](ManagementPlaneApiApi.md#GetLogicalRouterPortArpTable) | **Get** /logical-router-ports/{logical-router-port-id}/arp-table | Get the ARP table (IPv4) or Neighbor Discovery table (IPv6) for the Logical Router Port of the given id 
[**GetLogicalRouterPortArpTableInCsvFormatCsv**](ManagementPlaneApiApi.md#GetLogicalRouterPortArpTableInCsvFormatCsv) | **Get** /logical-router-ports/{logical-router-port-id}/arp-table?format&#x3D;csv | Get the ARP table (IPv4) or Neighbor Discovery table (IPv6) for the Logical Router Port of the given id 
[**GetLogicalRouterPortState**](ManagementPlaneApiApi.md#GetLogicalRouterPortState) | **Get** /logical-router-ports/{logical-router-port-id}/state | Get the Realized State of a Logical Router Port
[**GetLogicalRouterPortStatistics**](ManagementPlaneApiApi.md#GetLogicalRouterPortStatistics) | **Get** /logical-router-ports/{logical-router-port-id}/statistics | Get the statistics of a specified logical router port on all or a specified node
[**GetLogicalRouterPortStatisticsSummary**](ManagementPlaneApiApi.md#GetLogicalRouterPortStatisticsSummary) | **Get** /logical-router-ports/{logical-router-port-id}/statistics/summary | Get the statistics summary of a specified logical router port
[**GetLogicalRouterRouteTable**](ManagementPlaneApiApi.md#GetLogicalRouterRouteTable) | **Get** /logical-routers/{logical-router-id}/routing/route-table | Get route table on a given node for a logical router
[**GetLogicalRouterRouteTableInCsvFormatCsv**](ManagementPlaneApiApi.md#GetLogicalRouterRouteTableInCsvFormatCsv) | **Get** /logical-routers/{logical-router-id}/routing/route-table?format&#x3D;csv | Get route table on a node for a logical router
[**GetLogicalRouterRoutingTable**](ManagementPlaneApiApi.md#GetLogicalRouterRoutingTable) | **Get** /logical-routers/{logical-router-id}/routing/routing-table | Get RIB table on a specified node for a logical router
[**GetLogicalRouterRoutingTableInCsvFormatCsv**](ManagementPlaneApiApi.md#GetLogicalRouterRoutingTableInCsvFormatCsv) | **Get** /logical-routers/{logical-router-id}/routing/routing-table?format&#x3D;csv | Get RIB table on a specified node for a logical router
[**GetLogicalRouterState**](ManagementPlaneApiApi.md#GetLogicalRouterState) | **Get** /logical-routers/{logical-router-id}/state | Get the Realized State of a Logical Router
[**GetLogicalRouterStatus**](ManagementPlaneApiApi.md#GetLogicalRouterStatus) | **Get** /logical-routers/{logical-router-id}/status | Get the status for the Logical Router of the given id
[**GetLogicalServiceRouterClusterState**](ManagementPlaneApiApi.md#GetLogicalServiceRouterClusterState) | **Get** /logical-routers/{logical-router-id}/service-cluster/state | Get the Realized State of a Logical Service Router Cluster
[**GetLogicalSwitch**](ManagementPlaneApiApi.md#GetLogicalSwitch) | **Get** /logical-switches/{lswitch-id} | Get Logical Switch associated with the provided id (lswitch-id)
[**GetLogicalSwitchMacTable**](ManagementPlaneApiApi.md#GetLogicalSwitchMacTable) | **Get** /logical-switches/{lswitch-id}/mac-table | Get MAC Table for Logical Switch of the Given ID (lswitch-id)
[**GetLogicalSwitchMacTableInCsvFormatCsv**](ManagementPlaneApiApi.md#GetLogicalSwitchMacTableInCsvFormatCsv) | **Get** /logical-switches/{lswitch-id}/mac-table?format&#x3D;csv | Get MAC Table for Logical Switch of the Given ID (lswitch-id)
[**GetLogicalSwitchState**](ManagementPlaneApiApi.md#GetLogicalSwitchState) | **Get** /logical-switches/{lswitch-id}/state | Get the realized state associated with provided logical switch id
[**GetLogicalSwitchStatistics**](ManagementPlaneApiApi.md#GetLogicalSwitchStatistics) | **Get** /logical-switches/{lswitch-id}/statistics | Get Statistics for Logical Switch of the Given ID (lswitch-id)
[**GetLogicalSwitchStatus**](ManagementPlaneApiApi.md#GetLogicalSwitchStatus) | **Get** /logical-switches/{lswitch-id}/summary | Get Logical Switch runtime status info for a given logical switch
[**GetLogicalSwitchStatusSummary**](ManagementPlaneApiApi.md#GetLogicalSwitchStatusSummary) | **Get** /logical-switches/status | Get Status Summary of All Logical Switches in the System
[**GetLogicalSwitchVtepTable**](ManagementPlaneApiApi.md#GetLogicalSwitchVtepTable) | **Get** /logical-switches/{lswitch-id}/vtep-table | Get virtual tunnel endpoint table for logical switch of the given ID (lswitch-id) 
[**GetLogicalSwitchVtepTableInCsvFormatCsv**](ManagementPlaneApiApi.md#GetLogicalSwitchVtepTableInCsvFormatCsv) | **Get** /logical-switches/{lswitch-id}/vtep-table?format&#x3D;csv | Get virtual tunnel endpoint table for logical switch of the given ID (lswitch-id) 
[**GetMACAddresses**](ManagementPlaneApiApi.md#GetMACAddresses) | **Get** /mac-sets/{mac-set-id}/members | Get all MACAddresses in a MACSet
[**GetManualHealthCheck**](ManagementPlaneApiApi.md#GetManualHealthCheck) | **Get** /manual-health-checks/{manual-health-check-id} | Get an existing manual health check
[**GetMasterSwitchSetting**](ManagementPlaneApiApi.md#GetMasterSwitchSetting) | **Get** /idfw/master-switch-setting | Get Identity Firewall master switch enabled/disabled
[**GetMemberTypes**](ManagementPlaneApiApi.md#GetMemberTypes) | **Get** /ns-groups/{ns-group-id}/member-types | Get member types from NSGroup
[**GetMetadataProxyStatistics**](ManagementPlaneApiApi.md#GetMetadataProxyStatistics) | **Get** /md-proxies/{proxy-id}/statistics | Get Metadata Proxy statistics with given proxy id
[**GetMetadataProxyStatus**](ManagementPlaneApiApi.md#GetMetadataProxyStatus) | **Get** /md-proxies/{proxy-id}/{logical-switch-id}/status | Get Metadata Proxy status with given proxy id and attached logical switch.
[**GetNatRule**](ManagementPlaneApiApi.md#GetNatRule) | **Get** /logical-routers/{logical-router-id}/nat/rules/{rule-id} | Get a specific NAT rule from a given logical router
[**GetNatStatisticsPerLogicalRouter**](ManagementPlaneApiApi.md#GetNatStatisticsPerLogicalRouter) | **Get** /logical-routers/{logical-router-id}/nat/rules/statistics | Get the statistics of all rules of the logical router
[**GetNatStatisticsPerRule**](ManagementPlaneApiApi.md#GetNatStatisticsPerRule) | **Get** /logical-routers/{logical-router-id}/nat/rules/{rule-id}/statistics | Get the statistics of a specified logical router NAT Rule
[**GetNatStatisticsPerTransportNode**](ManagementPlaneApiApi.md#GetNatStatisticsPerTransportNode) | **Get** /transport-nodes/{node-id}/statistics/nat-rules | Get statistics for all logical router NAT rules on a transport node
[**GetNormalizations**](ManagementPlaneApiApi.md#GetNormalizations) | **Get** /normalizations | Get normalizations based on the query parameters
[**GetNsgroupVmDetails**](ManagementPlaneApiApi.md#GetNsgroupVmDetails) | **Get** /idfw/nsgroup-vm-details/{group-id} | Get all IDFW NSGroup VM details for a given NSGroup
[**GetPBRRule**](ManagementPlaneApiApi.md#GetPBRRule) | **Get** /pbr/sections/{section-id}/rules/{rule-id} | Read an Existing Rule
[**GetPBRRuleStats**](ManagementPlaneApiApi.md#GetPBRRuleStats) | **Get** /pbr/sections/{section-id}/rules/{rule-id}/stats | Get PBR rule level statistics.
[**GetPBRRules**](ManagementPlaneApiApi.md#GetPBRRules) | **Get** /pbr/sections/{section-id}/rules | Get All the Rules for a Section
[**GetPBRSection**](ManagementPlaneApiApi.md#GetPBRSection) | **Get** /pbr/sections/{section-id} | Get an Existing Section
[**GetPBRSectionStats**](ManagementPlaneApiApi.md#GetPBRSectionStats) | **Get** /pbr/sections/{section-id}/rules/stats | Get PBR section level statistics.
[**GetPBRSectionWithRulesListWithRules**](ManagementPlaneApiApi.md#GetPBRSectionWithRulesListWithRules) | **Post** /pbr/sections/{section-id}?action&#x3D;list_with_rules | Get an Existing Section, Including Rules
[**GetPeerConfig**](ManagementPlaneApiApi.md#GetPeerConfig) | **Get** /vpn/ipsec/sessions/{ipsec-vpn-session-id}/peer-config | Get VPN configuration for the peer site
[**GetPortMirroringSession**](ManagementPlaneApiApi.md#GetPortMirroringSession) | **Get** /mirror-sessions/{mirror-session-id} | Get the mirror session
[**GetRule**](ManagementPlaneApiApi.md#GetRule) | **Get** /firewall/sections/{section-id}/rules/{rule-id} | Read an Existing Rule
[**GetRuleState**](ManagementPlaneApiApi.md#GetRuleState) | **Get** /firewall/rules/{rule-id}/state | Get the Realized State of a Firewall Rule
[**GetRules**](ManagementPlaneApiApi.md#GetRules) | **Get** /firewall/sections/{section-id}/rules | Get All the Rules for a Section
[**GetRuntimeInterfaceOperationalStatus**](ManagementPlaneApiApi.md#GetRuntimeInterfaceOperationalStatus) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-runtimes/{instance-runtime-id}/interfaces/{interface_index}/status | Get operational status for an interface
[**GetRuntimeInterfaceStatistics**](ManagementPlaneApiApi.md#GetRuntimeInterfaceStatistics) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-runtimes/{instance-runtime-id}/interfaces/{interface_index}/statistics | Get statistics for a given interface identified by the interface index
[**GetSIServiceProfile**](ManagementPlaneApiApi.md#GetSIServiceProfile) | **Get** /serviceinsertion/services/{service-id}/service-profiles/{service-profile-id} | Get a particular ServiceProfile for a Service.
[**GetSection**](ManagementPlaneApiApi.md#GetSection) | **Get** /firewall/sections/{section-id} | Get an Existing Section
[**GetSectionState**](ManagementPlaneApiApi.md#GetSectionState) | **Get** /firewall/sections/{section-id}/state | Get the Realized State of a Firewall Section
[**GetSectionWithRulesListWithRules**](ManagementPlaneApiApi.md#GetSectionWithRulesListWithRules) | **Post** /firewall/sections/{section-id}?action&#x3D;list_with_rules | Get an Existing Section, Including Rules
[**GetSectionsSummary**](ManagementPlaneApiApi.md#GetSectionsSummary) | **Get** /firewall/sections/summary | Get the summary of sections in the firewall configuration.
[**GetServiceAssociations**](ManagementPlaneApiApi.md#GetServiceAssociations) | **Get** /ns-groups/{nsgroup-id}/service-associations | Get services to which the given nsgroup belongs to 
[**GetServiceAttachment**](ManagementPlaneApiApi.md#GetServiceAttachment) | **Get** /serviceinsertion/service-attachments/{service-attachment-id} | Get a particular service attachment.
[**GetServiceChain**](ManagementPlaneApiApi.md#GetServiceChain) | **Get** /serviceinsertion/service-chains/{service-chain-id} | Get a particular service chain.
[**GetServiceDeployment**](ManagementPlaneApiApi.md#GetServiceDeployment) | **Get** /serviceinsertion/services/{service-id}/service-deployments/{service-deployment-id} | Get a particular service deployment.
[**GetServiceDeploymentState**](ManagementPlaneApiApi.md#GetServiceDeploymentState) | **Get** /serviceinsertion/services/{service-id}/service-deployments/{service-deployment-id}/state | Get Service-Deployment state for Service.
[**GetServiceDeploymentStatus**](ManagementPlaneApiApi.md#GetServiceDeploymentStatus) | **Get** /serviceinsertion/services/{service-id}/service-deployments/{service-deployment-id}/status | Get a particular service deployment status.
[**GetServiceDeployments**](ManagementPlaneApiApi.md#GetServiceDeployments) | **Get** /serviceinsertion/services/{service-id}/service-deployments | Get all service deployments for the given service id
[**GetServiceInsertionExcludeList**](ManagementPlaneApiApi.md#GetServiceInsertionExcludeList) | **Get** /serviceinsertion/excludelist | Get list of members in exclude list
[**GetServiceInsertionRule**](ManagementPlaneApiApi.md#GetServiceInsertionRule) | **Get** /serviceinsertion/sections/{section-id}/rules/{rule-id} | Read an Existing Rule
[**GetServiceInsertionRules**](ManagementPlaneApiApi.md#GetServiceInsertionRules) | **Get** /serviceinsertion/sections/{section-id}/rules | Get All the Rules for a Section
[**GetServiceInsertionSection**](ManagementPlaneApiApi.md#GetServiceInsertionSection) | **Get** /serviceinsertion/sections/{section-id} | Get an Existing Section
[**GetServiceInsertionSectionWithRulesListWithRules**](ManagementPlaneApiApi.md#GetServiceInsertionSectionWithRulesListWithRules) | **Post** /serviceinsertion/sections/{section-id}?action&#x3D;list_with_rules | Get an Existing Section, Including Rules
[**GetServiceInsertionService**](ManagementPlaneApiApi.md#GetServiceInsertionService) | **Get** /serviceinsertion/services/{service-id} | Get an existing Service
[**GetServiceInsertionStatus**](ManagementPlaneApiApi.md#GetServiceInsertionStatus) | **Get** /serviceinsertion/status/{context-type} | Get ServiceInsertion global status for a context
[**GetServiceInstance**](ManagementPlaneApiApi.md#GetServiceInstance) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id} | Get Service-Instance for Service.
[**GetServiceInstanceNSGroups**](ManagementPlaneApiApi.md#GetServiceInstanceNSGroups) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/group-associations | Get NSgroups for a given ServiceInstance.
[**GetServiceInstanceState**](ManagementPlaneApiApi.md#GetServiceInstanceState) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/state | Get Service-Instance state for Service.
[**GetServiceInstanceStatus**](ManagementPlaneApiApi.md#GetServiceInstanceStatus) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/status | Get Service-Instance status for Service.
[**GetServiceManager**](ManagementPlaneApiApi.md#GetServiceManager) | **Get** /serviceinsertion/service-managers/{service-manager-id} | Get service manager
[**GetServiceProfileNSGroups**](ManagementPlaneApiApi.md#GetServiceProfileNSGroups) | **Get** /serviceinsertion/services/{service-id}/service-profiles/{service-profile-id}/nsgroups | Get NSgroups for a given ServiceProfile.
[**GetSolutionConfig**](ManagementPlaneApiApi.md#GetSolutionConfig) | **Get** /serviceinsertion/services/{service-id}/solution-configs/{solution-config-id} | Get Solution Config Information for a given solution config id.
[**GetStandaloneHostsSwitchSetting**](ManagementPlaneApiApi.md#GetStandaloneHostsSwitchSetting) | **Get** /idfw/standalone-host-switch-setting | Get Standalone hosts switch enabled/disabled
[**GetSwitchIpfixConfig**](ManagementPlaneApiApi.md#GetSwitchIpfixConfig) | **Get** /ipfix-obs-points/switch-global | Read global switch IPFIX export configuration
[**GetSwitchingProfile**](ManagementPlaneApiApi.md#GetSwitchingProfile) | **Get** /switching-profiles/{switching-profile-id} | Get Switching Profile by ID
[**GetSwitchingProfileStatus**](ManagementPlaneApiApi.md#GetSwitchingProfileStatus) | **Get** /switching-profiles/{switching-profile-id}/summary | Get Counts of Ports and Switches Using This Switching Profile
[**GetSystemStats**](ManagementPlaneApiApi.md#GetSystemStats) | **Get** /idfw/system-stats | Get IDFW system statistics data
[**GetTraceflow**](ManagementPlaneApiApi.md#GetTraceflow) | **Get** /traceflows/{traceflow-id} | Get the Traceflow round status and result summary
[**GetTraceflowObservations**](ManagementPlaneApiApi.md#GetTraceflowObservations) | **Get** /traceflows/{traceflow-id}/observations | Get observations for the Traceflow round
[**GetUnassociatedVirtualMachines**](ManagementPlaneApiApi.md#GetUnassociatedVirtualMachines) | **Get** /ns-groups/unassociated-virtual-machines | Get the list of all the virtual machines that are not a part of any existing NSGroup.
[**GetUserStats**](ManagementPlaneApiApi.md#GetUserStats) | **Get** /idfw/user-stats/{user-id} | Get IDFW user login events for a given user
[**GetVendorTemplate**](ManagementPlaneApiApi.md#GetVendorTemplate) | **Get** /serviceinsertion/services/{service-id}/vendor-templates/{vendor-template-id} | Get a particular vendor template for a given service.
[**GetVmStats**](ManagementPlaneApiApi.md#GetVmStats) | **Get** /idfw/vm-stats/{vm-ext-id} | Get IDFW user login events for a given VM
[**ListBGPCommunityLists**](ManagementPlaneApiApi.md#ListBGPCommunityLists) | **Get** /logical-routers/{logical-router-id}/routing/bgp/community-lists | Paginated list of BGP community lists on a logical router
[**ListBgpNeighbors**](ManagementPlaneApiApi.md#ListBgpNeighbors) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors | Paginated list of BGP Neighbors on a Logical Router
[**ListBridgeClusters**](ManagementPlaneApiApi.md#ListBridgeClusters) | **Get** /bridge-clusters | List All Bridge Clusters
[**ListBridgeEndpointProfiles**](ManagementPlaneApiApi.md#ListBridgeEndpointProfiles) | **Get** /bridge-endpoint-profiles | List All Bridge Endpoint Profiles
[**ListBridgeEndpoints**](ManagementPlaneApiApi.md#ListBridgeEndpoints) | **Get** /bridge-endpoints | List All Bridge Endpoints
[**ListComputeCollectionStatuses**](ManagementPlaneApiApi.md#ListComputeCollectionStatuses) | **Get** /idfw/compute-collections/status | List all IDFW enabled ComputeCollection statuses.
[**ListDADProfiles**](ManagementPlaneApiApi.md#ListDADProfiles) | **Get** /ipv6/dad-profiles | Read All IPV6 DADProfiles
[**ListDhcpIpPools**](ManagementPlaneApiApi.md#ListDhcpIpPools) | **Get** /dhcp/servers/{server-id}/ip-pools | Get a paginated list of a DHCP server&#x27;s IP pools
[**ListDhcpProfiles**](ManagementPlaneApiApi.md#ListDhcpProfiles) | **Get** /dhcp/server-profiles | Get a paginated list of DHCP server profiles
[**ListDhcpRelayProfiles**](ManagementPlaneApiApi.md#ListDhcpRelayProfiles) | **Get** /dhcp/relay-profiles | List All DHCP Relay Profiles
[**ListDhcpRelays**](ManagementPlaneApiApi.md#ListDhcpRelays) | **Get** /dhcp/relays | List all DHCP Relay Services
[**ListDhcpServers**](ManagementPlaneApiApi.md#ListDhcpServers) | **Get** /dhcp/servers | Get a paginated list of DHCP servers
[**ListDhcpStaticBindings**](ManagementPlaneApiApi.md#ListDhcpStaticBindings) | **Get** /dhcp/servers/{server-id}/static-bindings | Get a paginated list of a DHCP server&#x27;s static bindings
[**ListDhcpV6IpPools**](ManagementPlaneApiApi.md#ListDhcpV6IpPools) | **Get** /dhcp/servers/{server-id}/ipv6-ip-pools | Get a paginated list of a DHCP IPv6 server&#x27;s IP pools
[**ListDhcpV6StaticBindings**](ManagementPlaneApiApi.md#ListDhcpV6StaticBindings) | **Get** /dhcp/servers/{server-id}/ipv6-static-bindings | Get a paginated list of a DHCP IPv6 server&#x27;s static bindings
[**ListDnsForwaders**](ManagementPlaneApiApi.md#ListDnsForwaders) | **Get** /dns/forwarders | Get a paginated list of DNS forwarders
[**ListEnabledComputeCollections**](ManagementPlaneApiApi.md#ListEnabledComputeCollections) | **Get** /idfw/idfw-compute-collections | List all Identity firewall compute collections
[**ListFirewallProfiles**](ManagementPlaneApiApi.md#ListFirewallProfiles) | **Get** /firewall/profiles | Get firewall profiles available.
[**ListFirewallStatus**](ManagementPlaneApiApi.md#ListFirewallStatus) | **Get** /firewall/status | List all firewall status for supported contexts
[**ListIPPrefixLists**](ManagementPlaneApiApi.md#ListIPPrefixLists) | **Get** /logical-routers/{logical-router-id}/routing/ip-prefix-lists | Paginated List of IPPrefixLists
[**ListIPSecVPNDPDProfiles**](ManagementPlaneApiApi.md#ListIPSecVPNDPDProfiles) | **Get** /vpn/ipsec/dpd-profiles | Get IPSec dead peer detection (DPD)  profile list result
[**ListIPSecVPNIKEProfiles**](ManagementPlaneApiApi.md#ListIPSecVPNIKEProfiles) | **Get** /vpn/ipsec/ike-profiles | List IKE profiles
[**ListIPSecVPNLocalEndpoints**](ManagementPlaneApiApi.md#ListIPSecVPNLocalEndpoints) | **Get** /vpn/ipsec/local-endpoints | Get IPSec local endpoint list result
[**ListIPSecVPNPeerEndpoints**](ManagementPlaneApiApi.md#ListIPSecVPNPeerEndpoints) | **Get** /vpn/ipsec/peer-endpoints | Get IPSecVPNPeerEndpoint List Result
[**ListIPSecVPNServices**](ManagementPlaneApiApi.md#ListIPSecVPNServices) | **Get** /vpn/ipsec/services | Get IPSec VPN service list result
[**ListIPSecVPNSessions**](ManagementPlaneApiApi.md#ListIPSecVPNSessions) | **Get** /vpn/ipsec/sessions | Get IPSec VPN session list result
[**ListIPSecVPNTunnelProfiles**](ManagementPlaneApiApi.md#ListIPSecVPNTunnelProfiles) | **Get** /vpn/ipsec/tunnel-profiles | Get IPSecTunnelProfile List Result
[**ListIPSets**](ManagementPlaneApiApi.md#ListIPSets) | **Get** /ip-sets | List IPSets
[**ListInstanceEndpoints**](ManagementPlaneApiApi.md#ListInstanceEndpoints) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-endpoints | List all InstanceEndpoints of a Service Instance.
[**ListInstanceRuntimes**](ManagementPlaneApiApi.md#ListInstanceRuntimes) | **Get** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-runtimes | Returns list of instance runtimes of service VM being deployed
[**ListIpfixCollectorConfig**](ManagementPlaneApiApi.md#ListIpfixCollectorConfig) | **Get** /ipfix/collectorconfigs | List IPFIX collector configurations
[**ListIpfixCollectorUpmProfiles**](ManagementPlaneApiApi.md#ListIpfixCollectorUpmProfiles) | **Get** /ipfix-collector-profiles | List IPFIX Collector Profies
[**ListIpfixConfig**](ManagementPlaneApiApi.md#ListIpfixConfig) | **Get** /ipfix/configs | List IPFIX configuration
[**ListIpfixUpmProfiles**](ManagementPlaneApiApi.md#ListIpfixUpmProfiles) | **Get** /ipfix-profiles | List IPFIX Profies
[**ListL2VpnServices**](ManagementPlaneApiApi.md#ListL2VpnServices) | **Get** /vpn/l2vpn/services | Get all L2VPN services
[**ListL2VpnSessions**](ManagementPlaneApiApi.md#ListL2VpnSessions) | **Get** /vpn/l2vpn/sessions | Get all L2VPN sessions
[**ListLoadBalancerApplicationProfiles**](ManagementPlaneApiApi.md#ListLoadBalancerApplicationProfiles) | **Get** /loadbalancer/application-profiles | Retrieve a paginated list of load balancer application profiles
[**ListLoadBalancerClientSslProfiles**](ManagementPlaneApiApi.md#ListLoadBalancerClientSslProfiles) | **Get** /loadbalancer/client-ssl-profiles | Retrieve a paginated list of load balancer client-ssl profiles
[**ListLoadBalancerMonitors**](ManagementPlaneApiApi.md#ListLoadBalancerMonitors) | **Get** /loadbalancer/monitors | Retrieve a paginated list of load balancer monitors
[**ListLoadBalancerPersistenceProfiles**](ManagementPlaneApiApi.md#ListLoadBalancerPersistenceProfiles) | **Get** /loadbalancer/persistence-profiles | Retrieve a paginated list of load balancer persistence profiles
[**ListLoadBalancerPoolStatistics**](ManagementPlaneApiApi.md#ListLoadBalancerPoolStatistics) | **Get** /loadbalancer/services/{service-id}/pools/statistics | Get the statistics list of load balancer pools
[**ListLoadBalancerPoolStatuses**](ManagementPlaneApiApi.md#ListLoadBalancerPoolStatuses) | **Get** /loadbalancer/services/{service-id}/pools/status | Get the status list of load balancer pools
[**ListLoadBalancerPools**](ManagementPlaneApiApi.md#ListLoadBalancerPools) | **Get** /loadbalancer/pools | Retrieve a paginated list of load balancer pools
[**ListLoadBalancerRules**](ManagementPlaneApiApi.md#ListLoadBalancerRules) | **Get** /loadbalancer/rules | Retrieve a paginated list of load balancer rules
[**ListLoadBalancerServerSslProfiles**](ManagementPlaneApiApi.md#ListLoadBalancerServerSslProfiles) | **Get** /loadbalancer/server-ssl-profiles | Retrieve a paginated list of load balancer server-ssl profiles
[**ListLoadBalancerServices**](ManagementPlaneApiApi.md#ListLoadBalancerServices) | **Get** /loadbalancer/services | Retrieve a paginated list of load balancer services
[**ListLoadBalancerSslCiphersAndProtocols**](ManagementPlaneApiApi.md#ListLoadBalancerSslCiphersAndProtocols) | **Get** /loadbalancer/ssl/ciphers-and-protocols | Retrieve a list of supported SSL ciphers and protocols
[**ListLoadBalancerTcpProfiles**](ManagementPlaneApiApi.md#ListLoadBalancerTcpProfiles) | **Get** /loadbalancer/tcp-profiles | Retrieve a paginated list of load balancer TCP profiles
[**ListLoadBalancerVirtualServerStatuses**](ManagementPlaneApiApi.md#ListLoadBalancerVirtualServerStatuses) | **Get** /loadbalancer/services/{service-id}/virtual-servers/status | Get the status list of virtual servers in given load balancer service
[**ListLoadBalancerVirtualServers**](ManagementPlaneApiApi.md#ListLoadBalancerVirtualServers) | **Get** /loadbalancer/virtual-servers | Retrieve a paginated list of load balancer virtual servers
[**ListLoadBalancerVirtualServersStatistics**](ManagementPlaneApiApi.md#ListLoadBalancerVirtualServersStatistics) | **Get** /loadbalancer/services/{service-id}/virtual-servers/statistics | Get the statistics list of virtual servers
[**ListLogicalPorts**](ManagementPlaneApiApi.md#ListLogicalPorts) | **Get** /logical-ports | List All Logical Ports
[**ListLogicalRouterPorts**](ManagementPlaneApiApi.md#ListLogicalRouterPorts) | **Get** /logical-router-ports | List Logical Router Ports
[**ListLogicalRouters**](ManagementPlaneApiApi.md#ListLogicalRouters) | **Get** /logical-routers | List Logical Routers
[**ListLogicalSwitches**](ManagementPlaneApiApi.md#ListLogicalSwitches) | **Get** /logical-switches | List all Logical Switches
[**ListLogicalSwitchesByState**](ManagementPlaneApiApi.md#ListLogicalSwitchesByState) | **Get** /logical-switches/state | List logical switches by realized state
[**ListMACSets**](ManagementPlaneApiApi.md#ListMACSets) | **Get** /mac-sets | List MACSets
[**ListManualHealthChecks**](ManagementPlaneApiApi.md#ListManualHealthChecks) | **Get** /manual-health-checks | List manual health checks
[**ListMetadataProxy**](ManagementPlaneApiApi.md#ListMetadataProxy) | **Get** /md-proxies | Get a paginated list of metadata proxies
[**ListNDRAProfiles**](ManagementPlaneApiApi.md#ListNDRAProfiles) | **Get** /ipv6/nd-ra-profiles | Read All IPV6 NDRA Profiles
[**ListNSGroups**](ManagementPlaneApiApi.md#ListNSGroups) | **Get** /ns-groups | List NSGroups
[**ListNSProfiles**](ManagementPlaneApiApi.md#ListNSProfiles) | **Get** /ns-profiles | List NSProfiles
[**ListNSServiceGroups**](ManagementPlaneApiApi.md#ListNSServiceGroups) | **Get** /ns-service-groups | List all NSServiceGroups
[**ListNSServices**](ManagementPlaneApiApi.md#ListNSServices) | **Get** /ns-services | List all NSServices
[**ListNSSupportedAttributes**](ManagementPlaneApiApi.md#ListNSSupportedAttributes) | **Get** /ns-profiles/attributes | List NSProfile supported attribute and sub-attributes
[**ListNSSupportedAttributesTypes**](ManagementPlaneApiApi.md#ListNSSupportedAttributesTypes) | **Get** /ns-profiles/attribute-types | List NSProfile supported attribute types
[**ListNatRules**](ManagementPlaneApiApi.md#ListNatRules) | **Get** /logical-routers/{logical-router-id}/nat/rules | List NAT rules of the logical router
[**ListPBRSections**](ManagementPlaneApiApi.md#ListPBRSections) | **Get** /pbr/sections | List All PBR Sections
[**ListPacketCaptureSessions**](ManagementPlaneApiApi.md#ListPacketCaptureSessions) | **Get** /pktcap/sessions | Get the information of all packet capture sessions
[**ListPortMirroringSession**](ManagementPlaneApiApi.md#ListPortMirroringSession) | **Get** /mirror-sessions | List all mirror sessions
[**ListRouteMaps**](ManagementPlaneApiApi.md#ListRouteMaps) | **Get** /logical-routers/{logical-router-id}/routing/route-maps | Paginated List of RouteMaps
[**ListSIServiceProfiles**](ManagementPlaneApiApi.md#ListSIServiceProfiles) | **Get** /serviceinsertion/services/{service-id}/service-profiles | List all Service Profiles of a Service.
[**ListSections**](ManagementPlaneApiApi.md#ListSections) | **Get** /firewall/sections | List All Firewall Sections
[**ListServiceAttachments**](ManagementPlaneApiApi.md#ListServiceAttachments) | **Get** /serviceinsertion/service-attachments | Get all service attachments.
[**ListServiceChainMappings**](ManagementPlaneApiApi.md#ListServiceChainMappings) | **Get** /serviceinsertion/services/{service-id}/service-profiles/{service-profile-id}/service-chain-mappings | List all ServiceChainMappings.
[**ListServiceChains**](ManagementPlaneApiApi.md#ListServiceChains) | **Get** /serviceinsertion/service-chains | List all ServiceChains.
[**ListServiceConfigs**](ManagementPlaneApiApi.md#ListServiceConfigs) | **Get** /service-configs | List service configs
[**ListServiceInsertionSections**](ManagementPlaneApiApi.md#ListServiceInsertionSections) | **Get** /serviceinsertion/sections | List All Service Insertion Sections
[**ListServiceInsertionServices**](ManagementPlaneApiApi.md#ListServiceInsertionServices) | **Get** /serviceinsertion/services | List all Service-Insertion Services.
[**ListServiceInsertionStatus**](ManagementPlaneApiApi.md#ListServiceInsertionStatus) | **Get** /serviceinsertion/status | List all service insertion status for supported contexts
[**ListServiceInstances**](ManagementPlaneApiApi.md#ListServiceInstances) | **Get** /serviceinsertion/service-instances | Get all Service-Instances present in system
[**ListServiceInstancesForService**](ManagementPlaneApiApi.md#ListServiceInstancesForService) | **Get** /serviceinsertion/services/{service-id}/service-instances | Get all Service-Instances for Service.
[**ListServiceManagers**](ManagementPlaneApiApi.md#ListServiceManagers) | **Get** /serviceinsertion/service-managers | List service managers
[**ListServicePaths**](ManagementPlaneApiApi.md#ListServicePaths) | **Get** /serviceinsertion/service-chains/{service-chain-id}/service-paths | List all service paths
[**ListSolutionConfigs**](ManagementPlaneApiApi.md#ListSolutionConfigs) | **Get** /serviceinsertion/services/{service-id}/solution-configs | Get Solution Config Information associated with a given service.
[**ListStaticHopBfdPeers**](ManagementPlaneApiApi.md#ListStaticHopBfdPeers) | **Get** /logical-routers/{logical-router-id}/routing/static-routes/bfd-peers | List static routes BFD Peers
[**ListStaticRoutes**](ManagementPlaneApiApi.md#ListStaticRoutes) | **Get** /logical-routers/{logical-router-id}/routing/static-routes | Paginated List of Static Routes
[**ListSwitchingProfiles**](ManagementPlaneApiApi.md#ListSwitchingProfiles) | **Get** /switching-profiles | List Switching Profiles
[**ListTraceflows**](ManagementPlaneApiApi.md#ListTraceflows) | **Get** /traceflows | List all Traceflow rounds
[**ListTransportNodeStatusesByComputeCollectionId**](ManagementPlaneApiApi.md#ListTransportNodeStatusesByComputeCollectionId) | **Get** /idfw/compute-collections/{cc-ext-id}/transport-nodes/status | List all transport node and statuses based on idfw enabled ComputeCollection ID.
[**ListUserSessions**](ManagementPlaneApiApi.md#ListUserSessions) | **Get** /idfw/user-session-data | Get user session data
[**ListVendorTemplates**](ManagementPlaneApiApi.md#ListVendorTemplates) | **Get** /serviceinsertion/services/{service-id}/vendor-templates | List all VendorTemplates of a Service.
[**ListVirtualMachineStatusesByTransportNodeId**](ManagementPlaneApiApi.md#ListVirtualMachineStatusesByTransportNodeId) | **Get** /idfw/transport-nodes/{transport-node-id}/vms/status | List all VM and statuses based on transport node ID of idfw enabled compute collection.
[**LockSectionLock**](ManagementPlaneApiApi.md#LockSectionLock) | **Post** /firewall/sections/{section-id}?action&#x3D;lock | Lock a section
[**LookupAddress**](ManagementPlaneApiApi.md#LookupAddress) | **Get** /dns/forwarders/{forwarder-id}/nslookup | Resolve a given address via the DNS forwarder
[**PerformPoolMemberAction**](ManagementPlaneApiApi.md#PerformPoolMemberAction) | **Post** /loadbalancer/pools/{pool-id} | Add, remove, or modify load balancer pool members
[**ReAllocateServiceRoutersReallocate**](ManagementPlaneApiApi.md#ReAllocateServiceRoutersReallocate) | **Post** /logical-routers/{logical-router-id}?action&#x3D;reallocate | Re allocate edge node placement of TIER1 service routers
[**ReProcessLogicalRouterReprocess**](ManagementPlaneApiApi.md#ReProcessLogicalRouterReprocess) | **Post** /logical-routers/{logical-router-id}?action&#x3D;reprocess | Reprocesses a logical router configuration and publish updates to controller
[**ReadAdvertiseRuleList**](ManagementPlaneApiApi.md#ReadAdvertiseRuleList) | **Get** /logical-routers/{logical-router-id}/routing/advertisement/rules | Read the Advertisement Rules on a Logical Router
[**ReadAdvertisementConfig**](ManagementPlaneApiApi.md#ReadAdvertisementConfig) | **Get** /logical-routers/{logical-router-id}/routing/advertisement | Read the Advertisement Configuration on a Logical Router
[**ReadApplProxy**](ManagementPlaneApiApi.md#ReadApplProxy) | **Get** /node/services/applianceproxy | Read the Appliance Proxy service properties
[**ReadApplProxyStatus**](ManagementPlaneApiApi.md#ReadApplProxyStatus) | **Get** /node/services/applianceproxy/status | Read the Appliance Proxy service status
[**ReadBGPCommunityList**](ManagementPlaneApiApi.md#ReadBGPCommunityList) | **Get** /logical-routers/{logical-router-id}/routing/bgp/community-lists/{community-list-id} | Read a specific BGP community list from a logical router
[**ReadBgpConfig**](ManagementPlaneApiApi.md#ReadBgpConfig) | **Get** /logical-routers/{logical-router-id}/routing/bgp | Read the BGP Configuration on a Logical Router
[**ReadBgpNeighbor**](ManagementPlaneApiApi.md#ReadBgpNeighbor) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{id} | Read a specific BGP Neighbor on a Logical Router
[**ReadBgpNeighborWithPasswordShowSensitiveData**](ManagementPlaneApiApi.md#ReadBgpNeighborWithPasswordShowSensitiveData) | **Get** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{id}?action&#x3D;show-sensitive-data | Read a specific BGP Neighbor with password on a Logical Router
[**ReadDADProfile**](ManagementPlaneApiApi.md#ReadDADProfile) | **Get** /ipv6/dad-profiles/{dad-profile-id} | Read specified IPV6 DADProfile
[**ReadDebugInfoText**](ManagementPlaneApiApi.md#ReadDebugInfoText) | **Get** /logical-routers/{logical-router-id}/debug-info?format&#x3D;text | Read the debug information for the logical router
[**ReadDhcpIpPool**](ManagementPlaneApiApi.md#ReadDhcpIpPool) | **Get** /dhcp/servers/{server-id}/ip-pools/{pool-id} | Get a DHCP server&#x27;s IP pool with the specified pool ID
[**ReadDhcpProfile**](ManagementPlaneApiApi.md#ReadDhcpProfile) | **Get** /dhcp/server-profiles/{profile-id} | Get a DHCP server profile
[**ReadDhcpRelay**](ManagementPlaneApiApi.md#ReadDhcpRelay) | **Get** /dhcp/relays/{relay-id} | Read a DHCP Relay Service
[**ReadDhcpRelayProfile**](ManagementPlaneApiApi.md#ReadDhcpRelayProfile) | **Get** /dhcp/relay-profiles/{relay-profile-id} | Read a DHCP Relay Profile
[**ReadDhcpServer**](ManagementPlaneApiApi.md#ReadDhcpServer) | **Get** /dhcp/servers/{server-id} | Get a DHCP server with v4 and/or v6 servers
[**ReadDhcpStaticBinding**](ManagementPlaneApiApi.md#ReadDhcpStaticBinding) | **Get** /dhcp/servers/{server-id}/static-bindings/{binding-id} | Get a DHCP server&#x27;s static binding with the specified binding ID
[**ReadDhcpV6IpPool**](ManagementPlaneApiApi.md#ReadDhcpV6IpPool) | **Get** /dhcp/servers/{server-id}/ipv6-ip-pools/{pool-id} | Get a DHCP IPv6 server&#x27;s IP pool with the specified pool ID
[**ReadDhcpV6StaticBinding**](ManagementPlaneApiApi.md#ReadDhcpV6StaticBinding) | **Get** /dhcp/servers/{server-id}/ipv6-static-bindings/{binding-id} | Get a DHCP IPv6 server&#x27;s static binding with the specified binding ID
[**ReadDnsForwader**](ManagementPlaneApiApi.md#ReadDnsForwader) | **Get** /dns/forwarders/{forwarder-id} | Retrieve a DNS forwarder
[**ReadFirewallRule**](ManagementPlaneApiApi.md#ReadFirewallRule) | **Get** /firewall/rules/{rule-id} | Read an Existing Rule
[**ReadIPPrefixList**](ManagementPlaneApiApi.md#ReadIPPrefixList) | **Get** /logical-routers/{logical-router-id}/routing/ip-prefix-lists/{id} | Get a specific IPPrefixList on a Logical Router
[**ReadIPSet**](ManagementPlaneApiApi.md#ReadIPSet) | **Get** /ip-sets/{ip-set-id} | Read IPSet
[**ReadLoadBalancerApplicationProfile**](ManagementPlaneApiApi.md#ReadLoadBalancerApplicationProfile) | **Get** /loadbalancer/application-profiles/{application-profile-id} | Retrieve a load balancer application profile
[**ReadLoadBalancerClientSslProfile**](ManagementPlaneApiApi.md#ReadLoadBalancerClientSslProfile) | **Get** /loadbalancer/client-ssl-profiles/{client-ssl-profile-id} | Retrieve a load balancer client-ssl profile
[**ReadLoadBalancerMonitor**](ManagementPlaneApiApi.md#ReadLoadBalancerMonitor) | **Get** /loadbalancer/monitors/{monitor-id} | Retrieve a load balancer monitor
[**ReadLoadBalancerNodeUsage**](ManagementPlaneApiApi.md#ReadLoadBalancerNodeUsage) | **Get** /loadbalancer/usage-per-node/{node-id} | Read load balancer usage for the given node
[**ReadLoadBalancerNodeUsageSummary**](ManagementPlaneApiApi.md#ReadLoadBalancerNodeUsageSummary) | **Get** /loadbalancer/node-usage-summary | Read load balancer node usage summary
[**ReadLoadBalancerPersistenceProfile**](ManagementPlaneApiApi.md#ReadLoadBalancerPersistenceProfile) | **Get** /loadbalancer/persistence-profiles/{persistence-profile-id} | Retrieve a load balancer persistence profile
[**ReadLoadBalancerPool**](ManagementPlaneApiApi.md#ReadLoadBalancerPool) | **Get** /loadbalancer/pools/{pool-id} | Retrieve a load balancer pool
[**ReadLoadBalancerRule**](ManagementPlaneApiApi.md#ReadLoadBalancerRule) | **Get** /loadbalancer/rules/{rule-id} | Retrieve a load balancer rule
[**ReadLoadBalancerServerSslProfile**](ManagementPlaneApiApi.md#ReadLoadBalancerServerSslProfile) | **Get** /loadbalancer/server-ssl-profiles/{server-ssl-profile-id} | Retrieve a load balancer server-ssl profile
[**ReadLoadBalancerService**](ManagementPlaneApiApi.md#ReadLoadBalancerService) | **Get** /loadbalancer/services/{service-id} | Retrieve a load balancer service
[**ReadLoadBalancerServiceDebugInfo**](ManagementPlaneApiApi.md#ReadLoadBalancerServiceDebugInfo) | **Get** /loadbalancer/services/{service-id}/debug-info | Read the debug information of the load balancer service
[**ReadLoadBalancerServiceUsage**](ManagementPlaneApiApi.md#ReadLoadBalancerServiceUsage) | **Get** /loadbalancer/services/{service-id}/usage | Read the usage information of the given load balancer service
[**ReadLoadBalancerTcpProfile**](ManagementPlaneApiApi.md#ReadLoadBalancerTcpProfile) | **Get** /loadbalancer/tcp-profiles/{tcp-profile-id} | Retrieve a load balancer TCP profile
[**ReadLoadBalancerVirtualServer**](ManagementPlaneApiApi.md#ReadLoadBalancerVirtualServer) | **Get** /loadbalancer/virtual-servers/{virtual-server-id} | Retrieve a load balancer virtual server
[**ReadLogicalRouter**](ManagementPlaneApiApi.md#ReadLogicalRouter) | **Get** /logical-routers/{logical-router-id} | Read Logical Router
[**ReadLogicalRouterPort**](ManagementPlaneApiApi.md#ReadLogicalRouterPort) | **Get** /logical-router-ports/{logical-router-port-id} | Read Logical Router Port
[**ReadMACSet**](ManagementPlaneApiApi.md#ReadMACSet) | **Get** /mac-sets/{mac-set-id} | Read MACSet
[**ReadMetadataProxy**](ManagementPlaneApiApi.md#ReadMetadataProxy) | **Get** /md-proxies/{proxy-id} | Get a metadata proxy
[**ReadNDRAProfile**](ManagementPlaneApiApi.md#ReadNDRAProfile) | **Get** /ipv6/nd-ra-profiles/{nd-ra-profile-id} | Read specified IPV6 NDRA Profile
[**ReadNSGroup**](ManagementPlaneApiApi.md#ReadNSGroup) | **Get** /ns-groups/{ns-group-id} | Read NSGroup
[**ReadNSProfile**](ManagementPlaneApiApi.md#ReadNSProfile) | **Get** /ns-profiles/{ns-profile-id} | Read NSProfile
[**ReadNSService**](ManagementPlaneApiApi.md#ReadNSService) | **Get** /ns-services/{ns-service-id} | Read NSService
[**ReadNSServiceGroup**](ManagementPlaneApiApi.md#ReadNSServiceGroup) | **Get** /ns-service-groups/{ns-service-group-id} | Read NSServiceGroup
[**ReadPacketCaptureSession**](ManagementPlaneApiApi.md#ReadPacketCaptureSession) | **Get** /pktcap/session/{session-id} | Get the status of packet capture session
[**ReadRedistributionConfig**](ManagementPlaneApiApi.md#ReadRedistributionConfig) | **Get** /logical-routers/{logical-router-id}/routing/redistribution | Read the Redistribution Configuration on a Logical Router
[**ReadRedistributionRuleList**](ManagementPlaneApiApi.md#ReadRedistributionRuleList) | **Get** /logical-routers/{logical-router-id}/routing/redistribution/rules | Read All the Redistribution Rules on a Logical Router
[**ReadRouteMap**](ManagementPlaneApiApi.md#ReadRouteMap) | **Get** /logical-routers/{logical-router-id}/routing/route-maps/{id} | Get a specific RouteMap on a Logical Router
[**ReadRoutingBfdConfig**](ManagementPlaneApiApi.md#ReadRoutingBfdConfig) | **Get** /logical-routers/{logical-router-id}/routing/bfd-config | Read the Routing BFD Configuration
[**ReadRoutingConfig**](ManagementPlaneApiApi.md#ReadRoutingConfig) | **Get** /logical-routers/{logical-router-id}/routing | Read the Routing Configuration
[**ReadServiceConfig**](ManagementPlaneApiApi.md#ReadServiceConfig) | **Get** /service-configs/{config-set-id} | Read Service Config
[**ReadStaticHopBfdPeer**](ManagementPlaneApiApi.md#ReadStaticHopBfdPeer) | **Get** /logical-routers/{logical-router-id}/routing/static-routes/bfd-peers/{bfd-peer-id} | Read a static route BFD peer
[**ReadStaticRoute**](ManagementPlaneApiApi.md#ReadStaticRoute) | **Get** /logical-routers/{logical-router-id}/routing/static-routes/{id} | Get a specific Static Route on a Logical Router
[**ReallocateDhcpProfileEdgeClusterReallocate**](ManagementPlaneApiApi.md#ReallocateDhcpProfileEdgeClusterReallocate) | **Post** /dhcp/server-profiles/{server-profile-id}?action&#x3D;reallocate | Reallocate edge cluster and members of given DHCP profile.
[**RegisterServiceManager**](ManagementPlaneApiApi.md#RegisterServiceManager) | **Post** /serviceinsertion/service-managers | Register service manager
[**RemoveMACAddress**](ManagementPlaneApiApi.md#RemoveMACAddress) | **Delete** /mac-sets/{mac-set-id}/members/{mac-address} | Remove a MAC address from given MACSet
[**RemoveMemberRemoveMember**](ManagementPlaneApiApi.md#RemoveMemberRemoveMember) | **Post** /firewall/excludelist?action&#x3D;remove_member | Remove an existing object from the exclude list
[**RemoveServiceInsertionExcludeListMemberRemoveMember**](ManagementPlaneApiApi.md#RemoveServiceInsertionExcludeListMemberRemoveMember) | **Post** /serviceinsertion/excludelist?action&#x3D;remove_member | Remove an existing object from the exclude list
[**ResetFirewallRuleStatsReset**](ManagementPlaneApiApi.md#ResetFirewallRuleStatsReset) | **Post** /firewall/stats?action&#x3D;reset | Reset firewall rule statistics
[**ResetIPSecVPNSessionStatisticsReset**](ManagementPlaneApiApi.md#ResetIPSecVPNSessionStatisticsReset) | **Post** /vpn/ipsec/sessions/{session-id}/statistics?action&#x3D;reset | Reset the statistics of the given VPN session
[**ResolveSourceEntities**](ManagementPlaneApiApi.md#ResolveSourceEntities) | **Get** /serviceinsertion/source-entities | Resolve &#x27;source node id&#x27; value to source entities.
[**RestartPacketCaptureSessionRestart**](ManagementPlaneApiApi.md#RestartPacketCaptureSessionRestart) | **Post** /pktcap/session/{session-id}?action&#x3D;restart | Restart the packet capture session
[**RevisePBRRuleRevise**](ManagementPlaneApiApi.md#RevisePBRRuleRevise) | **Post** /pbr/sections/{section-id}/rules/{rule-id}?action&#x3D;revise | Update an Existing Rule and Reorder the Rule
[**RevisePBRSectionRevise**](ManagementPlaneApiApi.md#RevisePBRSectionRevise) | **Post** /pbr/sections/{section-id}?action&#x3D;revise | Update an Existing Section, including Its Position
[**RevisePBRSectionWithRulesReviseWithRules**](ManagementPlaneApiApi.md#RevisePBRSectionWithRulesReviseWithRules) | **Post** /pbr/sections/{section-id}?action&#x3D;revise_with_rules | Update an Existing Section with Rules
[**ReviseRuleRevise**](ManagementPlaneApiApi.md#ReviseRuleRevise) | **Post** /firewall/sections/{section-id}/rules/{rule-id}?action&#x3D;revise | Update an Existing Rule and Reorder the Rule
[**ReviseSectionRevise**](ManagementPlaneApiApi.md#ReviseSectionRevise) | **Post** /firewall/sections/{section-id}?action&#x3D;revise | Update an Existing Section, Including Its Position
[**ReviseSectionWithRulesReviseWithRules**](ManagementPlaneApiApi.md#ReviseSectionWithRulesReviseWithRules) | **Post** /firewall/sections/{section-id}?action&#x3D;revise_with_rules | Update an Existing Section with Rules
[**ReviseServiceInsertionRuleRevise**](ManagementPlaneApiApi.md#ReviseServiceInsertionRuleRevise) | **Post** /serviceinsertion/sections/{section-id}/rules/{rule-id}?action&#x3D;revise | Update an Existing Rule and Reorder the Rule
[**ReviseServiceInsertionSectionRevise**](ManagementPlaneApiApi.md#ReviseServiceInsertionSectionRevise) | **Post** /serviceinsertion/sections/{section-id}?action&#x3D;revise | Update an Existing Section, Including Its Position
[**ReviseServiceInsertionSectionWithRulesReviseWithRules**](ManagementPlaneApiApi.md#ReviseServiceInsertionSectionWithRulesReviseWithRules) | **Post** /serviceinsertion/sections/{section-id}?action&#x3D;revise_with_rules | Update an Existing Section with Rules
[**ServiceConfigBatchOperation**](ManagementPlaneApiApi.md#ServiceConfigBatchOperation) | **Post** /service-configs/batch | Creates/Updates service configs sent in batch request
[**TerminatePacketCaptureSessionTerminate**](ManagementPlaneApiApi.md#TerminatePacketCaptureSessionTerminate) | **Post** /pktcap/session/{session-id}?action&#x3D;terminate | Terminate the packet capture session by session id
[**UnSetPasswordOnBgpNeighbor**](ManagementPlaneApiApi.md#UnSetPasswordOnBgpNeighbor) | **Post** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{id} | Unset/Delete password property on specific BGP Neighbor on Logical Router
[**UnlockSectionUnlock**](ManagementPlaneApiApi.md#UnlockSectionUnlock) | **Post** /firewall/sections/{section-id}?action&#x3D;unlock | Unlock a section
[**UpdateAdvertiseRuleList**](ManagementPlaneApiApi.md#UpdateAdvertiseRuleList) | **Put** /logical-routers/{logical-router-id}/routing/advertisement/rules | Update the Advertisement Rules on a Logical Router
[**UpdateAdvertisementConfig**](ManagementPlaneApiApi.md#UpdateAdvertisementConfig) | **Put** /logical-routers/{logical-router-id}/routing/advertisement | Update the Advertisement Configuration on a Logical Router
[**UpdateBGPCommunityList**](ManagementPlaneApiApi.md#UpdateBGPCommunityList) | **Put** /logical-routers/{logical-router-id}/routing/bgp/community-lists/{community-list-id} | Update a specific BGP community list from a logical router
[**UpdateBGPCommunityListOld**](ManagementPlaneApiApi.md#UpdateBGPCommunityListOld) | **Put** /logical-routers/{logical-router-id}/routing/bgp/communty-lists/{community-list-id} | Update a specific BGP community list from a logical router
[**UpdateBgpConfig**](ManagementPlaneApiApi.md#UpdateBgpConfig) | **Put** /logical-routers/{logical-router-id}/routing/bgp | Update the BGP Configuration on a Logical Router
[**UpdateBgpNeighbor**](ManagementPlaneApiApi.md#UpdateBgpNeighbor) | **Put** /logical-routers/{logical-router-id}/routing/bgp/neighbors/{id} | Update a specific BGP Neighbor on a Logical Router
[**UpdateBridgeCluster**](ManagementPlaneApiApi.md#UpdateBridgeCluster) | **Put** /bridge-clusters/{bridgecluster-id} | Update a Bridge Cluster
[**UpdateBridgeEndpoint**](ManagementPlaneApiApi.md#UpdateBridgeEndpoint) | **Put** /bridge-endpoints/{bridgeendpoint-id} | Update a Bridge Endpoint
[**UpdateBridgeEndpointProfile**](ManagementPlaneApiApi.md#UpdateBridgeEndpointProfile) | **Put** /bridge-endpoint-profiles/{bridgeendpointprofile-id} | Update a Bridge Endpoint Profile
[**UpdateDADProfile**](ManagementPlaneApiApi.md#UpdateDADProfile) | **Put** /ipv6/dad-profiles/{dad-profile-id} | Update DADProfile
[**UpdateDhcpIpPool**](ManagementPlaneApiApi.md#UpdateDhcpIpPool) | **Put** /dhcp/servers/{server-id}/ip-pools/{pool-id} | Update a DHCP server&#x27;s IP pool
[**UpdateDhcpProfile**](ManagementPlaneApiApi.md#UpdateDhcpProfile) | **Put** /dhcp/server-profiles/{profile-id} | Update a DHCP server profile
[**UpdateDhcpRelay**](ManagementPlaneApiApi.md#UpdateDhcpRelay) | **Put** /dhcp/relays/{relay-id} | Update a DHCP Relay Service
[**UpdateDhcpRelayProfile**](ManagementPlaneApiApi.md#UpdateDhcpRelayProfile) | **Put** /dhcp/relay-profiles/{relay-profile-id} | Update a DHCP Relay Profile
[**UpdateDhcpServer**](ManagementPlaneApiApi.md#UpdateDhcpServer) | **Put** /dhcp/servers/{server-id} | Update a DHCP server with v4 and/or v6 servers
[**UpdateDhcpStaticBinding**](ManagementPlaneApiApi.md#UpdateDhcpStaticBinding) | **Put** /dhcp/servers/{server-id}/static-bindings/{binding-id} | Update a DHCP server&#x27;s static binding
[**UpdateDhcpV6IpPool**](ManagementPlaneApiApi.md#UpdateDhcpV6IpPool) | **Put** /dhcp/servers/{server-id}/ipv6-ip-pools/{pool-id} | Update a DHCP IPv6 server&#x27;s IP pool
[**UpdateDhcpV6StaticBinding**](ManagementPlaneApiApi.md#UpdateDhcpV6StaticBinding) | **Put** /dhcp/servers/{server-id}/ipv6-static-bindings/{binding-id} | Update a DHCP IPv6 server&#x27;s static binding
[**UpdateDnsForwarder**](ManagementPlaneApiApi.md#UpdateDnsForwarder) | **Put** /dns/forwarders/{forwarder-id} | Update a specific DNS forwarder
[**UpdateEnabledComputeCollection**](ManagementPlaneApiApi.md#UpdateEnabledComputeCollection) | **Put** /idfw/idfw-compute-collections/{cc-ext-id} | Update IDFW compute collection
[**UpdateExcludeList**](ManagementPlaneApiApi.md#UpdateExcludeList) | **Put** /firewall/excludelist | Modify exclude list
[**UpdateFirewallProfile**](ManagementPlaneApiApi.md#UpdateFirewallProfile) | **Put** /firewall/profiles/{profile-id} | Update a firewall profile.
[**UpdateFirewallStatus**](ManagementPlaneApiApi.md#UpdateFirewallStatus) | **Put** /firewall/status/{context-type} | Update global firewall status for dfw context
[**UpdateIPPrefixList**](ManagementPlaneApiApi.md#UpdateIPPrefixList) | **Put** /logical-routers/{logical-router-id}/routing/ip-prefix-lists/{id} | Update a specific IPPrefixList on a Logical Router
[**UpdateIPSecVPNDPDProfile**](ManagementPlaneApiApi.md#UpdateIPSecVPNDPDProfile) | **Put** /vpn/ipsec/dpd-profiles/{ipsec-vpn-dpd-profile-id} | Edit IPSec dead peer detection (DPD) profile
[**UpdateIPSecVPNIKEProfile**](ManagementPlaneApiApi.md#UpdateIPSecVPNIKEProfile) | **Put** /vpn/ipsec/ike-profiles/{ipsec-vpn-ike-profile-id} | Edit custom IKE Profile
[**UpdateIPSecVPNLocalEndpoint**](ManagementPlaneApiApi.md#UpdateIPSecVPNLocalEndpoint) | **Put** /vpn/ipsec/local-endpoints/{ipsec-vpn-local-endpoint-id} | Edit custom IPSec local endpoint
[**UpdateIPSecVPNPeerEndpoint**](ManagementPlaneApiApi.md#UpdateIPSecVPNPeerEndpoint) | **Put** /vpn/ipsec/peer-endpoints/{ipsec-vpn-peer-endpoint-id} | Edit custom IPSecPeerEndpoint
[**UpdateIPSecVPNService**](ManagementPlaneApiApi.md#UpdateIPSecVPNService) | **Put** /vpn/ipsec/services/{ipsec-vpn-service-id} | Edit IPSec VPN service
[**UpdateIPSecVPNSession**](ManagementPlaneApiApi.md#UpdateIPSecVPNSession) | **Put** /vpn/ipsec/sessions/{ipsec-vpn-session-id} | Edit IPSec VPN session
[**UpdateIPSecVPNTunnelProfile**](ManagementPlaneApiApi.md#UpdateIPSecVPNTunnelProfile) | **Put** /vpn/ipsec/tunnel-profiles/{ipsec-vpn-tunnel-profile-id} | Edit custom IPSecTunnelProfile
[**UpdateIPSet**](ManagementPlaneApiApi.md#UpdateIPSet) | **Put** /ip-sets/{ip-set-id} | Update IPSet
[**UpdateIpfixCollectorConfig**](ManagementPlaneApiApi.md#UpdateIpfixCollectorConfig) | **Put** /ipfix/collectorconfigs/{collector-config-id} | Update an existing IPFIX collector configuration
[**UpdateIpfixCollectorUpmProfile**](ManagementPlaneApiApi.md#UpdateIpfixCollectorUpmProfile) | **Put** /ipfix-collector-profiles/{ipfix-collector-profile-id} | Update an existing IPFIX collector profile
[**UpdateIpfixConfig**](ManagementPlaneApiApi.md#UpdateIpfixConfig) | **Put** /ipfix/configs/{config-id} | Update an existing IPFIX configuration
[**UpdateIpfixUpmProfile**](ManagementPlaneApiApi.md#UpdateIpfixUpmProfile) | **Put** /ipfix-profiles/{ipfix-profile-id} | Update an existing IPFIX profile
[**UpdateL2VpnService**](ManagementPlaneApiApi.md#UpdateL2VpnService) | **Put** /vpn/l2vpn/services/{l2vpn-service-id} | Edit a L2VPN service
[**UpdateL2VpnSession**](ManagementPlaneApiApi.md#UpdateL2VpnSession) | **Put** /vpn/l2vpn/sessions/{l2vpn-session-id} | Edit a L2VPN session
[**UpdateLoadBalancerApplicationProfile**](ManagementPlaneApiApi.md#UpdateLoadBalancerApplicationProfile) | **Put** /loadbalancer/application-profiles/{application-profile-id} | Update a load balancer application profile
[**UpdateLoadBalancerClientSslProfile**](ManagementPlaneApiApi.md#UpdateLoadBalancerClientSslProfile) | **Put** /loadbalancer/client-ssl-profiles/{client-ssl-profile-id} | Update a load balancer client-ssl profile
[**UpdateLoadBalancerMonitor**](ManagementPlaneApiApi.md#UpdateLoadBalancerMonitor) | **Put** /loadbalancer/monitors/{monitor-id} | Update a load balancer monitor
[**UpdateLoadBalancerPersistenceProfile**](ManagementPlaneApiApi.md#UpdateLoadBalancerPersistenceProfile) | **Put** /loadbalancer/persistence-profiles/{persistence-profile-id} | Update a load balancer persistence profile
[**UpdateLoadBalancerPool**](ManagementPlaneApiApi.md#UpdateLoadBalancerPool) | **Put** /loadbalancer/pools/{pool-id} | Update a load balancer pool
[**UpdateLoadBalancerRule**](ManagementPlaneApiApi.md#UpdateLoadBalancerRule) | **Put** /loadbalancer/rules/{rule-id} | Update a load balancer rule
[**UpdateLoadBalancerServerSslProfile**](ManagementPlaneApiApi.md#UpdateLoadBalancerServerSslProfile) | **Put** /loadbalancer/server-ssl-profiles/{server-ssl-profile-id} | Update a load balancer server-ssl profile
[**UpdateLoadBalancerService**](ManagementPlaneApiApi.md#UpdateLoadBalancerService) | **Put** /loadbalancer/services/{service-id} | Update a load balancer service
[**UpdateLoadBalancerTcpProfile**](ManagementPlaneApiApi.md#UpdateLoadBalancerTcpProfile) | **Put** /loadbalancer/tcp-profiles/{tcp-profile-id} | Update a load balancer TCP profile
[**UpdateLoadBalancerVirtualServer**](ManagementPlaneApiApi.md#UpdateLoadBalancerVirtualServer) | **Put** /loadbalancer/virtual-servers/{virtual-server-id} | Update a load balancer virtual server
[**UpdateLoadBalancerVirtualServerWithRulesUpdateWithRules**](ManagementPlaneApiApi.md#UpdateLoadBalancerVirtualServerWithRulesUpdateWithRules) | **Put** /loadbalancer/virtual-servers/{virtual-server-id}?action&#x3D;update_with_rules | Update a load balancer virtual server with rules
[**UpdateLogicalPort**](ManagementPlaneApiApi.md#UpdateLogicalPort) | **Put** /logical-ports/{lport-id} | Update a Logical Port
[**UpdateLogicalRouter**](ManagementPlaneApiApi.md#UpdateLogicalRouter) | **Put** /logical-routers/{logical-router-id} | Update a Logical Router
[**UpdateLogicalRouterPort**](ManagementPlaneApiApi.md#UpdateLogicalRouterPort) | **Put** /logical-router-ports/{logical-router-port-id} | Update a Logical Router Port
[**UpdateLogicalSwitch**](ManagementPlaneApiApi.md#UpdateLogicalSwitch) | **Put** /logical-switches/{lswitch-id} | Update a Logical Switch
[**UpdateMACSet**](ManagementPlaneApiApi.md#UpdateMACSet) | **Put** /mac-sets/{mac-set-id} | Update MACSet
[**UpdateMasterSwitchSetting**](ManagementPlaneApiApi.md#UpdateMasterSwitchSetting) | **Put** /idfw/master-switch-setting | Update IDFW master switch setting enabled/disabled
[**UpdateMetadataProxy**](ManagementPlaneApiApi.md#UpdateMetadataProxy) | **Put** /md-proxies/{proxy-id} | Update a metadata proxy
[**UpdateNDRAProfile**](ManagementPlaneApiApi.md#UpdateNDRAProfile) | **Put** /ipv6/nd-ra-profiles/{nd-ra-profile-id} | Update NDRA Profile
[**UpdateNSGroup**](ManagementPlaneApiApi.md#UpdateNSGroup) | **Put** /ns-groups/{ns-group-id} | Update NSGroup
[**UpdateNSProfile**](ManagementPlaneApiApi.md#UpdateNSProfile) | **Put** /ns-profiles/{ns-profile-id} | Update NSProfile
[**UpdateNSService**](ManagementPlaneApiApi.md#UpdateNSService) | **Put** /ns-services/{ns-service-id} | Update NSService
[**UpdateNSServiceGroup**](ManagementPlaneApiApi.md#UpdateNSServiceGroup) | **Put** /ns-service-groups/{ns-service-group-id} | Update NSServiceGroup
[**UpdateNatRule**](ManagementPlaneApiApi.md#UpdateNatRule) | **Put** /logical-routers/{logical-router-id}/nat/rules/{rule-id} | Update a specific NAT rule from a given logical router
[**UpdatePBRRule**](ManagementPlaneApiApi.md#UpdatePBRRule) | **Put** /pbr/sections/{section-id}/rules/{rule-id} | Update an Existing Rule
[**UpdatePBRSection**](ManagementPlaneApiApi.md#UpdatePBRSection) | **Put** /pbr/sections/{section-id} | Update an Existing Section
[**UpdatePBRSectionWithRulesUpdateWithRules**](ManagementPlaneApiApi.md#UpdatePBRSectionWithRulesUpdateWithRules) | **Post** /pbr/sections/{section-id}?action&#x3D;update_with_rules | Update an Existing Section, Including Its Rules
[**UpdatePortMirroringSession**](ManagementPlaneApiApi.md#UpdatePortMirroringSession) | **Put** /mirror-sessions/{mirror-session-id} | Update the mirror session
[**UpdateRedistributionConfig**](ManagementPlaneApiApi.md#UpdateRedistributionConfig) | **Put** /logical-routers/{logical-router-id}/routing/redistribution | Update the Redistribution Configuration on a Logical Router
[**UpdateRedistributionRuleList**](ManagementPlaneApiApi.md#UpdateRedistributionRuleList) | **Put** /logical-routers/{logical-router-id}/routing/redistribution/rules | Update All the Redistribution Rules on a Logical Router
[**UpdateRouteMap**](ManagementPlaneApiApi.md#UpdateRouteMap) | **Put** /logical-routers/{logical-router-id}/routing/route-maps/{id} | Update a specific RouteMap on a Logical Router
[**UpdateRoutingBfdConfig**](ManagementPlaneApiApi.md#UpdateRoutingBfdConfig) | **Put** /logical-routers/{logical-router-id}/routing/bfd-config | Update the BFD Configuration for BFD peers for routing
[**UpdateRoutingConfig**](ManagementPlaneApiApi.md#UpdateRoutingConfig) | **Put** /logical-routers/{logical-router-id}/routing | Update the Routing Configuration
[**UpdateRule**](ManagementPlaneApiApi.md#UpdateRule) | **Put** /firewall/sections/{section-id}/rules/{rule-id} | Update an Existing Rule
[**UpdateSection**](ManagementPlaneApiApi.md#UpdateSection) | **Put** /firewall/sections/{section-id} | Update an Existing Section
[**UpdateSectionWithRulesUpdateWithRules**](ManagementPlaneApiApi.md#UpdateSectionWithRulesUpdateWithRules) | **Post** /firewall/sections/{section-id}?action&#x3D;update_with_rules | Update an Existing Section, Including Its Rules
[**UpdateServiceAttachment**](ManagementPlaneApiApi.md#UpdateServiceAttachment) | **Put** /serviceinsertion/service-attachments/{service-attachment-id} | Update an existing service attachment.
[**UpdateServiceConfig**](ManagementPlaneApiApi.md#UpdateServiceConfig) | **Put** /service-configs/{config-set-id} | Update service config
[**UpdateServiceDeployment**](ManagementPlaneApiApi.md#UpdateServiceDeployment) | **Put** /serviceinsertion/services/{service-id}/service-deployments/{service-deployment-id} | Update an existing Service Deployment.
[**UpdateServiceInsertionExcludeList**](ManagementPlaneApiApi.md#UpdateServiceInsertionExcludeList) | **Put** /serviceinsertion/excludelist | Modify exclude list
[**UpdateServiceInsertionRule**](ManagementPlaneApiApi.md#UpdateServiceInsertionRule) | **Put** /serviceinsertion/sections/{section-id}/rules/{rule-id} | Update an Existing Rule
[**UpdateServiceInsertionSection**](ManagementPlaneApiApi.md#UpdateServiceInsertionSection) | **Put** /serviceinsertion/sections/{section-id} | Update an Existing Section
[**UpdateServiceInsertionSectionWithRulesUpdateWithRules**](ManagementPlaneApiApi.md#UpdateServiceInsertionSectionWithRulesUpdateWithRules) | **Post** /serviceinsertion/sections/{section-id}?action&#x3D;update_with_rules | Update an Existing Section, Including Its Rules
[**UpdateServiceInsertionService**](ManagementPlaneApiApi.md#UpdateServiceInsertionService) | **Put** /serviceinsertion/services/{service-id} | Update an existing Service
[**UpdateServiceInsertionStatus**](ManagementPlaneApiApi.md#UpdateServiceInsertionStatus) | **Put** /serviceinsertion/status/{context-type} | Update global ServiceInsertion status for a context
[**UpdateServiceInstance**](ManagementPlaneApiApi.md#UpdateServiceInstance) | **Put** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id} | Update an existing Service-Instance.
[**UpdateServiceManager**](ManagementPlaneApiApi.md#UpdateServiceManager) | **Put** /serviceinsertion/service-managers/{service-manager-id} | Update service manager
[**UpdateServiceVmState**](ManagementPlaneApiApi.md#UpdateServiceVmState) | **Post** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-runtimes/{instance-runtime-id} | Update maintenance mode or runtime state of a service VM
[**UpdateSolutionConfig**](ManagementPlaneApiApi.md#UpdateSolutionConfig) | **Put** /serviceinsertion/services/{service-id}/solution-configs/{solution-config-id} | Updates Solution Config for a given Service
[**UpdateStandaloneHostsSwitchSetting**](ManagementPlaneApiApi.md#UpdateStandaloneHostsSwitchSetting) | **Put** /idfw/standalone-host-switch-setting | Update IDFW master switch setting enabled/disabled
[**UpdateStaticHopBfdPeer**](ManagementPlaneApiApi.md#UpdateStaticHopBfdPeer) | **Put** /logical-routers/{logical-router-id}/routing/static-routes/bfd-peers/{bfd-peer-id} | Update a static route BFD peer
[**UpdateStaticRoute**](ManagementPlaneApiApi.md#UpdateStaticRoute) | **Put** /logical-routers/{logical-router-id}/routing/static-routes/{id} | Update a specific Static Route Rule on a Logical Router
[**UpdateSwitchIpfixConfig**](ManagementPlaneApiApi.md#UpdateSwitchIpfixConfig) | **Put** /ipfix-obs-points/switch-global | Update global switch IPFIX export configuration
[**UpdateSwitchingProfile**](ManagementPlaneApiApi.md#UpdateSwitchingProfile) | **Put** /switching-profiles/{switching-profile-id} | Update a Switching Profile
[**UpgradeServiceDeploymentUpgrade**](ManagementPlaneApiApi.md#UpgradeServiceDeploymentUpgrade) | **Post** /serviceinsertion/services/{service-id}/service-deployments/{service-deployment-id}?action&#x3D;upgrade | Upgrade all VMs part of this service deployment to new Spec OVF.
[**UpgradeServiceVMsUpgrade**](ManagementPlaneApiApi.md#UpgradeServiceVMsUpgrade) | **Post** /serviceinsertion/services/{service-id}/service-instances/{service-instance-id}/instance-runtimes?action&#x3D;upgrade | Upgrade service VMs using newer version of OVF
[**VerifyPortMirroringSessionVerify**](ManagementPlaneApiApi.md#VerifyPortMirroringSessionVerify) | **Post** /mirror-sessions/{mirror-session-id}?action&#x3D;verify | Verify whether the mirror session is still valid

# **AddBgpNeighbor**
> BgpNeighbor AddBgpNeighbor(ctx, body, logicalRouterId)
Add a new BGP Neighbor on a Logical Router

Add a new BGP Neighbor on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BgpNeighbor**](BgpNeighbor.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**BgpNeighbor**](BgpNeighbor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddIPPrefixList**
> IpPrefixList AddIPPrefixList(ctx, body, logicalRouterId)
Add IPPrefixList on a Logical Router

Adds a new IPPrefixList on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpPrefixList**](IpPrefixList.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**IpPrefixList**](IPPrefixList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddInstanceEndpoint**
> InstanceEndpoint AddInstanceEndpoint(ctx, body, serviceId, serviceInstanceId)
Add an InstanceEndpoint for a Service Instance

Adds a new instance endpoint. It belongs to one service instance and is attached to one service attachment. It represents a redirection target for a Rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**InstanceEndpoint**](InstanceEndpoint.md)|  | 
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

[**InstanceEndpoint**](InstanceEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMACAddress**
> MacAddressElement AddMACAddress(ctx, body, macSetId)
Add a MAC address to a MACSet

Add an individual MAC address to a MACSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacAddressElement**](MacAddressElement.md)|  | 
  **macSetId** | **string**| MAC Set Id | 

### Return type

[**MacAddressElement**](MACAddressElement.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddMemberAddMember**
> ResourceReference AddMemberAddMember(ctx, body)
Add a new object in the exclude list

Add a new object in the exclude list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceReference**](ResourceReference.md)|  | 

### Return type

[**ResourceReference**](ResourceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNatRule**
> NatRule AddNatRule(ctx, body, logicalRouterId)
Add a NAT rule in a specific logical router

Add a NAT rule in a specific logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NatRule**](NatRule.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**NatRule**](NatRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNatRulesCreateMultiple**
> NatRuleList AddNatRulesCreateMultiple(ctx, body, logicalRouterId)
Add multiple NAT rules in a specific logical router

Create multiple NAT rules in a specific logical router. The API succeeds only when all rules are accepted and created successfully. Any one validation voilation will fail the API, no rule will be created. The ruleIds of each rules can be found from the responsed message. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NatRuleList**](NatRuleList.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**NatRuleList**](NatRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddOrRemoveNSGroupExpression**
> NsGroup AddOrRemoveNSGroupExpression(ctx, body, nsGroupId, action)
Add NSGroup expression

Add/remove the expressions passed in the request body to/from the NSGroup 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsGroupExpressionList**](NsGroupExpressionList.md)|  | 
  **nsGroupId** | **string**| NSGroup Id | 
  **action** | **string**| Specifies addition or removal action | 

### Return type

[**NsGroup**](NSGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPBRRuleInSection**
> PbrRule AddPBRRuleInSection(ctx, body, sectionId, optional)
Add a Single Rule in a Section

Adds a new PBR rule in existing PBR section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrRule**](PbrRule.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiAddPBRRuleInSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddPBRRuleInSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**PbrRule**](PBRRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPBRRulesInSectionCreateMultiple**
> PbrRuleList AddPBRRulesInSectionCreateMultiple(ctx, body, sectionId, optional)
Add Multiple Rules in a Section

Create multiple PBR rules in existing PBR section bounded by limit of 1000 PBR rules per section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrRuleList**](PbrRuleList.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiAddPBRRulesInSectionCreateMultipleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddPBRRulesInSectionCreateMultipleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**PbrRuleList**](PBRRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPBRSection**
> PbrSection AddPBRSection(ctx, body, optional)
Create a New Empty Section

Creates new empty PBR section in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrSection**](PbrSection.md)|  | 
 **optional** | ***ManagementPlaneApiApiAddPBRSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddPBRSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**PbrSection**](PBRSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPBRSectionWithRulesCreateWithRules**
> PbrSectionRuleList AddPBRSectionWithRulesCreateWithRules(ctx, body, optional)
Create a Section with Rules

Creates a new PBR section with rules. The limit on the number of rules is defined by maxItems in collection types for PBRRule (PBRRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number of rule references is not supported.  Instead, to create sections, use: POST /api/v1/pbr/sections  To create rules, use: POST /api/v1/pbr/sections/&lt;section-id&gt;/rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrSectionRuleList**](PbrSectionRuleList.md)|  | 
 **optional** | ***ManagementPlaneApiApiAddPBRSectionWithRulesCreateWithRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddPBRSectionWithRulesCreateWithRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**PbrSectionRuleList**](PBRSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRemoveIPAddress**
> IpAddressElement AddRemoveIPAddress(ctx, body, ipSetId, action)
Add a IP address to a IPSet

Add/Remove an individual IP address to an IPSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpAddressElement**](IpAddressElement.md)|  | 
  **ipSetId** | **string**| IP Set Id | 
  **action** | **string**| Specifies addition or removal action | 

### Return type

[**IpAddressElement**](IPAddressElement.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRouteMap**
> RouteMap AddRouteMap(ctx, body, logicalRouterId)
Add RouteMap on a Logical Router

Adds a new RouteMap on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteMap**](RouteMap.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**RouteMap**](RouteMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRuleInSection**
> FirewallRule AddRuleInSection(ctx, body, sectionId, optional)
Add a Single Rule in a Section

Adds a new firewall rule in existing firewall section. Adding firewall rule to a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale content to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallRule**](FirewallRule.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiAddRuleInSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddRuleInSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddRulesInSectionCreateMultiple**
> FirewallRuleList AddRulesInSectionCreateMultiple(ctx, body, sectionId, optional)
Add Multiple Rules in a Section

Create multiple firewall rules in existing firewall section bounded by limit of 1000 firewall rules per section. Adding multiple firewall rules in a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallRuleList**](FirewallRuleList.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiAddRulesInSectionCreateMultipleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddRulesInSectionCreateMultipleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**FirewallRuleList**](FirewallRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSIServiceProfile**
> BaseServiceProfile AddSIServiceProfile(ctx, body, serviceId)
Add ServiceProfile for a given Service.

Adds a new service profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseServiceProfile**](BaseServiceProfile.md)|  | 
  **serviceId** | **string**|  | 

### Return type

[**BaseServiceProfile**](BaseServiceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSection**
> FirewallSection AddSection(ctx, body, optional)
Create a New Empty Section

Creates new empty firewall section in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSection**](FirewallSection.md)|  | 
 **optional** | ***ManagementPlaneApiApiAddSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**FirewallSection**](FirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddSectionWithRulesCreateWithRules**
> FirewallSectionRuleList AddSectionWithRulesCreateWithRules(ctx, body, optional)
Create a Section with Rules

Creates a new firewall section with rules. The limit on the number of rules is defined by maxItems in collection types for FirewallRule (FirewallRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number of rule references is not supported.  Instead, to create sections, use: POST /api/v1/firewall/sections  To create rules, use: POST /api/v1/firewall/sections/&lt;section-id&gt;/rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSectionRuleList**](FirewallSectionRuleList.md)|  | 
 **optional** | ***ManagementPlaneApiApiAddSectionWithRulesCreateWithRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddSectionWithRulesCreateWithRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**FirewallSectionRuleList**](FirewallSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceAttachment**
> ServiceAttachment AddServiceAttachment(ctx, body)
Add a Service Attachment.

Adds a new Service attachment. A service attachment represents a point on NSX entity (Example: Logical Router) to which service instance can be connected through an InstanceEndpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceAttachment**](ServiceAttachment.md)|  | 

### Return type

[**ServiceAttachment**](ServiceAttachment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceChain**
> ServiceChain AddServiceChain(ctx, body)
Add Service Chain

Adds a new service chain. Service Chains is can contain profile belonging to same or different Service(s). It represents a redirection target for a Rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceChain**](ServiceChain.md)|  | 

### Return type

[**ServiceChain**](ServiceChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceInsertionExcludeListMemberAddMember**
> ResourceReference AddServiceInsertionExcludeListMemberAddMember(ctx, body)
Add a new member in the exclude list

Add a new member in the exclude list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourceReference**](ResourceReference.md)|  | 

### Return type

[**ResourceReference**](ResourceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceInsertionRuleInSection**
> ServiceInsertionRule AddServiceInsertionRuleInSection(ctx, body, sectionId, optional)
Add a Single Rule in a Section

Adds a new serviceinsertion rule in existing serviceinsertion section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionRule**](ServiceInsertionRule.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiAddServiceInsertionRuleInSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddServiceInsertionRuleInSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**ServiceInsertionRule**](ServiceInsertionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceInsertionRulesInSectionCreateMultiple**
> ServiceInsertionRuleList AddServiceInsertionRulesInSectionCreateMultiple(ctx, body, sectionId, optional)
Add Multiple Rules in a Section

Create multiple serviceinsertion rules in existing serviceinsertion section bounded by limit of 1000 serviceinsertion rules per section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionRuleList**](ServiceInsertionRuleList.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiAddServiceInsertionRulesInSectionCreateMultipleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddServiceInsertionRulesInSectionCreateMultipleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**ServiceInsertionRuleList**](ServiceInsertionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceInsertionSection**
> ServiceInsertionSection AddServiceInsertionSection(ctx, body, optional)
Create a New Empty Section

Creates new empty Service Insertion section in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionSection**](ServiceInsertionSection.md)|  | 
 **optional** | ***ManagementPlaneApiApiAddServiceInsertionSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddServiceInsertionSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**ServiceInsertionSection**](ServiceInsertionSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceInsertionSectionWithRulesCreateWithRules**
> ServiceInsertionSectionRuleList AddServiceInsertionSectionWithRulesCreateWithRules(ctx, body, optional)
Create a Section with Rules

Creates a new serviceinsertion section with rules. The limit on the number of rules is defined by maxItems in collection types for ServiceInsertionRule (ServiceInsertionRuleXXXList types). When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported.  Instead, to create sections, use: POST /api/v1/serviceinsertion/sections  To create rules, use: POST /api/v1/serviceinsertion/sections/&lt;section-id&gt;/rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionSectionRuleList**](ServiceInsertionSectionRuleList.md)|  | 
 **optional** | ***ManagementPlaneApiApiAddServiceInsertionSectionWithRulesCreateWithRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiAddServiceInsertionSectionWithRulesCreateWithRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**ServiceInsertionSectionRuleList**](ServiceInsertionSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceInsertionService**
> ServiceDefinition AddServiceInsertionService(ctx, body)
Create a Service-Insertion Service

Creates new Service-Insertion Service in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceDefinition**](ServiceDefinition.md)|  | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddServiceInstance**
> BaseServiceInstance AddServiceInstance(ctx, body, serviceId)
Add a Service Instance for a specified Service.

Adds a new Service-Instance under the specified Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseServiceInstance**](BaseServiceInstance.md)|  | 
  **serviceId** | **string**|  | 

### Return type

[**BaseServiceInstance**](BaseServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddStaticRoute**
> StaticRoute AddStaticRoute(ctx, body, logicalRouterId)
Add Static Routes on a Logical Router

Adds a new static route on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StaticRoute**](StaticRoute.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**StaticRoute**](StaticRoute.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVendorTemplate**
> VendorTemplate AddVendorTemplate(ctx, body, serviceId)
Add Vendor Template for a given Service

Adds a new vendor template. Vendor templates are service level objects, registered to be used in Service Profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VendorTemplate**](VendorTemplate.md)|  | 
  **serviceId** | **string**|  | 

### Return type

[**VendorTemplate**](VendorTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckMemberIfExistsCheckIfExists**
> ResourceReference CheckMemberIfExistsCheckIfExists(ctx, objectId, optional)
Check if the object a member of the exclude list

Check if the object a member of the exclude list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| identifier of the object | 
 **optional** | ***ManagementPlaneApiApiCheckMemberIfExistsCheckIfExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiCheckMemberIfExistsCheckIfExistsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deepCheck** | **optional.Bool**| Check all parents | [default to false]
 **objectType** | **optional.String**| Object type of an entity | 

### Return type

[**ResourceReference**](ResourceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearDnsForwarderCacheClearCache**
> ClearDnsForwarderCacheClearCache(ctx, forwarderId)
Clear the current cache of the DNS forwarder.

Clear the current cache of the DNS forwarder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplProxyActionRestart**
> NodeServiceStatusProperties CreateApplProxyActionRestart(ctx, )
Restart, start or stop the Appliance Proxy Service

Restart, start or stop the Appliance Proxy Service

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeServiceStatusProperties**](NodeServiceStatusProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplProxyActionStart**
> NodeServiceStatusProperties CreateApplProxyActionStart(ctx, )
Restart, start or stop the Appliance Proxy Service

Restart, start or stop the Appliance Proxy Service

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeServiceStatusProperties**](NodeServiceStatusProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplProxyActionStop**
> NodeServiceStatusProperties CreateApplProxyActionStop(ctx, )
Restart, start or stop the Appliance Proxy Service

Restart, start or stop the Appliance Proxy Service

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeServiceStatusProperties**](NodeServiceStatusProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBGPCommunityList**
> BgpCommunityList CreateBGPCommunityList(ctx, body, logicalRouterId)
Create a new BGP community list on a logical router

Add a new BGP Community List on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BgpCommunityList**](BgpCommunityList.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**BgpCommunityList**](BGPCommunityList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBridgeCluster**
> BridgeCluster CreateBridgeCluster(ctx, body)
Create a Bridge Cluster

Creates a bridge cluster. It is collection of transport nodes that will do the bridging for overlay network to vlan networks. Bridge cluster may have one or more transport nodes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BridgeCluster**](BridgeCluster.md)|  | 

### Return type

[**BridgeCluster**](BridgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBridgeEndpoint**
> BridgeEndpoint CreateBridgeEndpoint(ctx, body)
Create a Bridge Endpoint

Creates a Bridge Endpoint. It describes the physical attributes of the bridge like vlan. A logical port can be attached to a vif providing bridging functionality from the logical overlay network to the physical vlan network 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BridgeEndpoint**](BridgeEndpoint.md)|  | 

### Return type

[**BridgeEndpoint**](BridgeEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBridgeEndpointProfile**
> BridgeEndpointProfile CreateBridgeEndpointProfile(ctx, body)
Create a Bridge Endpoint Profile

Creates a Bridge Endpoint Profile. Profile contains edge cluster id, indexes of the member nodes, fialover mode and high availability mode for a Bridge EndPoint 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BridgeEndpointProfile**](BridgeEndpointProfile.md)|  | 

### Return type

[**BridgeEndpointProfile**](BridgeEndpointProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDADProfile**
> DadProfile CreateDADProfile(ctx, body)
Create a new DADProfile

Adds a new DADProfile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DadProfile**](DadProfile.md)|  | 

### Return type

[**DadProfile**](DADProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpIpPool**
> DhcpIpPool CreateDhcpIpPool(ctx, body, serverId)
Create an ip pool for a DHCP server

Create an ip pool for a local DHCP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpIpPool**](DhcpIpPool.md)|  | 
  **serverId** | **string**|  | 

### Return type

[**DhcpIpPool**](DhcpIpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpProfile**
> DhcpProfile CreateDhcpProfile(ctx, body)
Create a DHCP server profile

Create a DHCP server profile. If no edge member is specified, edge members to run the dhcp servers will be auto-allocated from the edge cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpProfile**](DhcpProfile.md)|  | 

### Return type

[**DhcpProfile**](DhcpProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpRelay**
> DhcpRelayService CreateDhcpRelay(ctx, body)
Create a DHCP Relay Service

Creates a dhcp relay service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpRelayService**](DhcpRelayService.md)|  | 

### Return type

[**DhcpRelayService**](DhcpRelayService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpRelayProfile**
> DhcpRelayProfile CreateDhcpRelayProfile(ctx, body)
Create a DHCP Relay Profile

Creates a dhcp relay profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpRelayProfile**](DhcpRelayProfile.md)|  | 

### Return type

[**DhcpRelayProfile**](DhcpRelayProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpServer**
> LogicalDhcpServer CreateDhcpServer(ctx, body)
Create a DHCP server

Create a logical DHCP server with v4 and/or v6 servers. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalDhcpServer**](LogicalDhcpServer.md)|  | 

### Return type

[**LogicalDhcpServer**](LogicalDhcpServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpStaticBinding**
> DhcpStaticBinding CreateDhcpStaticBinding(ctx, body, serverId)
Create a static binding for a DHCP server

Create a static binding for a logical DHCP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpStaticBinding**](DhcpStaticBinding.md)|  | 
  **serverId** | **string**|  | 

### Return type

[**DhcpStaticBinding**](DhcpStaticBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpV6IpPool**
> DhcpV6IpPool CreateDhcpV6IpPool(ctx, body, serverId)
Create an ip pool for a DHCP IPv6 server

Create an ip pool for a local DHCP IPv6 server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpV6IpPool**](DhcpV6IpPool.md)|  | 
  **serverId** | **string**|  | 

### Return type

[**DhcpV6IpPool**](DhcpV6IpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDhcpV6StaticBinding**
> DhcpV6StaticBinding CreateDhcpV6StaticBinding(ctx, body, serverId)
Create a static binding for a DHCP IPv6 server

Create a static binding for a logical DHCP IPv6 server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpV6StaticBinding**](DhcpV6StaticBinding.md)|  | 
  **serverId** | **string**|  | 

### Return type

[**DhcpV6StaticBinding**](DhcpV6StaticBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDnsForwader**
> DnsForwarder CreateDnsForwader(ctx, body)
Create a DNS forwader

Create a DNS forwader upon a logical router. There is only one DNS forwarder can be created upon a given logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsForwarder**](DnsForwarder.md)|  | 

### Return type

[**DnsForwarder**](DnsForwarder.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFirewallProfile**
> BaseFirewallProfile CreateFirewallProfile(ctx, body)
Create a firewall profile.

Create a firewall profile with values provided. It creates profile based resource_type in the payload. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseFirewallProfile**](BaseFirewallProfile.md)|  | 

### Return type

[**BaseFirewallProfile**](BaseFirewallProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSecVPNDPDProfile**
> IpSecVpndpdProfile CreateIPSecVPNDPDProfile(ctx, body)
Create dead peer detection (DPD) profile

Create dead peer detection (DPD) profile. Any change in profile affects all sessions consuming this profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpndpdProfile**](IpSecVpndpdProfile.md)|  | 

### Return type

[**IpSecVpndpdProfile**](IPSecVPNDPDProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSecVPNIKEProfile**
> IpSecVpnikeProfile CreateIPSecVPNIKEProfile(ctx, body)
Create custom internet key exchange (IKE) Profile

Create custom internet key exchange (IKE) Profile. IKE Profile is a reusable profile that captures IKE and phase one negotiation parameters. System will be pre provisioned with system owned non editable default IKE profile and suggested set of profiles that can be used for peering with popular remote peers like AWS VPN. User can create custom profiles as needed. Any change in profile affects all sessions consuming this profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnikeProfile**](IpSecVpnikeProfile.md)|  | 

### Return type

[**IpSecVpnikeProfile**](IPSecVPNIKEProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSecVPNLocalEndpoint**
> IpSecVpnLocalEndpoint CreateIPSecVPNLocalEndpoint(ctx, body)
Create custom local endpoint

Create custom IPSec local endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnLocalEndpoint**](IpSecVpnLocalEndpoint.md)|  | 

### Return type

[**IpSecVpnLocalEndpoint**](IPSecVPNLocalEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSecVPNPeerEndPoint**
> IpSecVpnPeerEndpoint CreateIPSecVPNPeerEndPoint(ctx, body)
Create custom peer endpoint

Create custom IPSec peer endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnPeerEndpoint**](IpSecVpnPeerEndpoint.md)|  | 

### Return type

[**IpSecVpnPeerEndpoint**](IPSecVPNPeerEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSecVPNService**
> IpSecVpnService CreateIPSecVPNService(ctx, body)
Create VPN service

Create VPN service for given logical router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnService**](IpSecVpnService.md)|  | 

### Return type

[**IpSecVpnService**](IPSecVPNService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSecVPNSession**
> IpSecVpnSession CreateIPSecVPNSession(ctx, body)
Create new VPN session

Create new VPN session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnSession**](IpSecVpnSession.md)|  | 

### Return type

[**IpSecVpnSession**](IPSecVPNSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSecVPNTunnelProfile**
> IpSecVpnTunnelProfile CreateIPSecVPNTunnelProfile(ctx, body)
Create custom IPSec tunnel profile

Create custom IPSec tunnel profile. IPSec tunnel profile is a reusable profile that captures phase two negotiation parameters and tunnel properties. System will be provisioned with system owned non editable default IPSec tunnel profile. Any change in profile affects all sessions consuming this profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnTunnelProfile**](IpSecVpnTunnelProfile.md)|  | 

### Return type

[**IpSecVpnTunnelProfile**](IPSecVPNTunnelProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIPSet**
> IpSet CreateIPSet(ctx, body)
Create IPSet

Creates a new IPSet that can group either IPv4 or IPv6 individual ip addresses, ranges or subnets. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSet**](IpSet.md)|  | 

### Return type

[**IpSet**](IPSet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpfixCollectorConfig**
> IpfixCollectorConfig CreateIpfixCollectorConfig(ctx, body)
Create a new IPFIX collector configuration

Create a new IPFIX collector configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixCollectorConfig**](IpfixCollectorConfig.md)|  | 

### Return type

[**IpfixCollectorConfig**](IpfixCollectorConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpfixCollectorUpmProfile**
> IpfixCollectorUpmProfile CreateIpfixCollectorUpmProfile(ctx, body)
Create a new IPFIX collector profile

Create a new IPFIX collector profile with essential properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixCollectorUpmProfile**](IpfixCollectorUpmProfile.md)|  | 

### Return type

[**IpfixCollectorUpmProfile**](IpfixCollectorUpmProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpfixConfig**
> IpfixConfig CreateIpfixConfig(ctx, body)
Create a new IPFIX configuration

Create a new IPFIX configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixConfig**](IpfixConfig.md)|  | 

### Return type

[**IpfixConfig**](IpfixConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpfixUpmProfile**
> IpfixUpmProfile CreateIpfixUpmProfile(ctx, body)
Create a new IPFIX profile

Create a new IPFIX profile with essential properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixUpmProfile**](IpfixUpmProfile.md)|  | 

### Return type

[**IpfixUpmProfile**](IpfixUpmProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateL2VpnService**
> L2VpnService CreateL2VpnService(ctx, body)
Create L2VPN service

Create L2VPN service for a given logical router

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**L2VpnService**](L2VpnService.md)|  | 

### Return type

[**L2VpnService**](L2VpnService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateL2VpnSession**
> L2VpnSession CreateL2VpnSession(ctx, body)
Create L2VPN session

Create L2VPN session and bind to a L2VPNService

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**L2VpnSession**](L2VpnSession.md)|  | 

### Return type

[**L2VpnSession**](L2VpnSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerApplicationProfile**
> LbAppProfile CreateLoadBalancerApplicationProfile(ctx, body)
Create a load balancer application profile

Create a load balancer application profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbAppProfile**](LbAppProfile.md)|  | 

### Return type

[**LbAppProfile**](LbAppProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerClientSslProfile**
> LbClientSslProfile CreateLoadBalancerClientSslProfile(ctx, body)
Create a load balancer client-ssl profile

Create a load balancer client-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbClientSslProfile**](LbClientSslProfile.md)|  | 

### Return type

[**LbClientSslProfile**](LbClientSslProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerMonitor**
> LbMonitor CreateLoadBalancerMonitor(ctx, body)
Create a load balancer monitor

Create a load balancer monitor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbMonitor**](LbMonitor.md)|  | 

### Return type

[**LbMonitor**](LbMonitor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerPersistenceProfile**
> LbPersistenceProfile CreateLoadBalancerPersistenceProfile(ctx, body)
Create a load balancer persistence profile

Create a load balancer persistence profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbPersistenceProfile**](LbPersistenceProfile.md)|  | 

### Return type

[**LbPersistenceProfile**](LbPersistenceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerPool**
> LbPool CreateLoadBalancerPool(ctx, body)
Create a load balancer pool

Create a load balancer pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbPool**](LbPool.md)|  | 

### Return type

[**LbPool**](LbPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerRule**
> LbRule CreateLoadBalancerRule(ctx, body)
Create a load balancer rule

Create a load balancer rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbRule**](LbRule.md)|  | 

### Return type

[**LbRule**](LbRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerServerSslProfile**
> LbServerSslProfile CreateLoadBalancerServerSslProfile(ctx, body)
Create a load balancer server-ssl profile

Create a load balancer server-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbServerSslProfile**](LbServerSslProfile.md)|  | 

### Return type

[**LbServerSslProfile**](LbServerSslProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerService**
> LbService CreateLoadBalancerService(ctx, body)
Create a load balancer service

Create a load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbService**](LbService.md)|  | 

### Return type

[**LbService**](LbService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerTcpProfile**
> LbTcpProfile CreateLoadBalancerTcpProfile(ctx, body)
Create a load balancer TCP profile

Create a load balancer TCP profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbTcpProfile**](LbTcpProfile.md)|  | 

### Return type

[**LbTcpProfile**](LbTcpProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerVirtualServer**
> LbVirtualServer CreateLoadBalancerVirtualServer(ctx, body)
Create a load balancer virtual server

Create a load balancer virtual server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbVirtualServer**](LbVirtualServer.md)|  | 

### Return type

[**LbVirtualServer**](LbVirtualServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoadBalancerVirtualServerWithRulesCreateWithRules**
> LbVirtualServerWithRule CreateLoadBalancerVirtualServerWithRulesCreateWithRules(ctx, body)
Create a load balancer virtual server with rules

It is used to create virtual servers, the associated rules and bind the rules to the virtual server. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. If the virtual server need to consume some existed rules without change, those rules should not be specified in this array, otherwise, the rules are updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbVirtualServerWithRule**](LbVirtualServerWithRule.md)|  | 

### Return type

[**LbVirtualServerWithRule**](LbVirtualServerWithRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLogicalPort**
> LogicalPort CreateLogicalPort(ctx, body)
Create a Logical Port

Creates a new logical switch port. The required parameters are the associated logical_switch_id and admin_state (UP or DOWN). Optional parameters are the attachment and switching_profile_ids. If you don't specify switching_profile_ids, default switching profiles are assigned to the port. If you don't specify an attachment, the switch port remains empty. To configure an attachment, you must specify an id, and optionally you can specify an attachment_type (VIF or LOGICALROUTER). The attachment_type is VIF by default. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalPort**](LogicalPort.md)|  | 

### Return type

[**LogicalPort**](LogicalPort.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLogicalRouter**
> LogicalRouter CreateLogicalRouter(ctx, body)
Create a Logical Router

Creates a logical router. The required parameters are router_type (TIER0 or TIER1) and edge_cluster_id (TIER0 only). Optional parameters include internal and external transit network addresses. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalRouter**](LogicalRouter.md)|  | 

### Return type

[**LogicalRouter**](LogicalRouter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLogicalRouterPort**
> LogicalRouterPort CreateLogicalRouterPort(ctx, body)
Create a Logical Router Port

Creates a logical router port. The required parameters include resource_type (LogicalRouterUpLinkPort, LogicalRouterDownLinkPort, LogicalRouterLinkPort, LogicalRouterLoopbackPort, LogicalRouterCentralizedServicePort); and logical_router_id (the router to which each logical router port is assigned). The service_bindings parameter is optional. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalRouterPort**](LogicalRouterPort.md)|  | 

### Return type

[**LogicalRouterPort**](LogicalRouterPort.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLogicalSwitch**
> LogicalSwitch CreateLogicalSwitch(ctx, body)
Create a Logical Switch

Creates a new logical switch. The request must include the transport_zone_id, display_name, and admin_state (UP or DOWN). The replication_mode (MTEP or SOURCE) is required for overlay logical switches, but not for VLAN-based logical switches. A vlan needs to be provided for VLAN-based logical switches 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalSwitch**](LogicalSwitch.md)|  | 

### Return type

[**LogicalSwitch**](LogicalSwitch.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMACSet**
> MacSet CreateMACSet(ctx, body)
Create MACSet

Creates a new MACSet that can group individual MAC addresses. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacSet**](MacSet.md)|  | 

### Return type

[**MacSet**](MACSet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateManualHealthCheck**
> ManualHealthCheck CreateManualHealthCheck(ctx, body)
Create a new manual health check request

Create a new manual health check request with essential properties. It's disallowed to create new one until the count of in-progress manual health check is less than 50. A manual health check will be deleted automatically after finished for 24 hours. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ManualHealthCheck**](ManualHealthCheck.md)|  | 

### Return type

[**ManualHealthCheck**](ManualHealthCheck.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMetadataProxy**
> MetadataProxy CreateMetadataProxy(ctx, body)
Create a metadata proxy

Create a metadata proxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetadataProxy**](MetadataProxy.md)|  | 

### Return type

[**MetadataProxy**](MetadataProxy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNDRAProfile**
> NdraProfile CreateNDRAProfile(ctx, body)
Create a new NDRA Profile

Adds a new NDRAProfile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NdraProfile**](NdraProfile.md)|  | 

### Return type

[**NdraProfile**](NDRAProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNSGroup**
> NsGroup CreateNSGroup(ctx, body)
Create NSGroup

Creates a new NSGroup that can group NSX resources - VIFs, Lports and LSwitches as well as the grouping objects - IPSet, MACSet and other NSGroups. For NSGroups containing VM criteria(both static and dynamic), system VMs will not be included as members. This filter applies at VM level only. Exceptions are as follows: 1. LogicalPorts and VNI of System VMs will be included in NSGroup if the criteria  is based on LogicalPort, LogicalSwitch or VNI directly. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsGroup**](NsGroup.md)|  | 

### Return type

[**NsGroup**](NSGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNSProfile**
> NsProfile CreateNSProfile(ctx, body)
Create NSProfile

Creates a new NSProfile which allows users to encapsulate attribute and sub-attributes of network services. Rules for using attributes and sub-attributes in single NSProfile 1. One type of attribute can't have multiple occurrences. ( Example -  Attribute type APP_ID can be used only once per NSProfile.) 2. Values for an attribute are mentioned as array of strings.  ( Example - For type APP_ID , values can be mentioned as [\"SSL\",\"FTP\"].) 3. If sub-attribtes are mentioned for an attribute, then only single  value is allowed for that attribute. 4. To get a list of supported  attributes and sub-attributes fire the following REST API  GET https://&lt;nsx-mgr&gt;/api/v1/ns-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsProfile**](NsProfile.md)|  | 

### Return type

[**NsProfile**](NSProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNSService**
> NsService CreateNSService(ctx, body)
Create NSService

Creates a new NSService which allows users to specify characteristics to use for matching network traffic. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsService**](NsService.md)|  | 

### Return type

[**NsService**](NSService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNSServiceGroup**
> NsServiceGroup CreateNSServiceGroup(ctx, body)
Create NSServiceGroup

Creates a new NSServiceGroup which can contain NSServices. A given NSServiceGroup can contain either only ether type of NSServices or only non-ether type of NSServices, i.e. an NSServiceGroup cannot contain a mix of both ether and non-ether types of NSServices. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsServiceGroup**](NsServiceGroup.md)|  | 

### Return type

[**NsServiceGroup**](NSServiceGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePacketCaptureSession**
> PacketCaptureSession CreatePacketCaptureSession(ctx, body)
Create an new packet capture session

Create an new packet capture session on given node with specified options 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PacketCaptureRequest**](PacketCaptureRequest.md)|  | 

### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePortMirroringSessions**
> PortMirroringSession CreatePortMirroringSessions(ctx, body)
Create a mirror session

Create a mirror session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortMirroringSession**](PortMirroringSession.md)|  | 

### Return type

[**PortMirroringSession**](PortMirroringSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceConfig**
> ServiceConfig CreateServiceConfig(ctx, body)
Create service config

Creates a new service config that can group profiles and configs 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceConfig**](ServiceConfig.md)|  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSolutionConfig**
> SolutionConfig CreateSolutionConfig(ctx, body, serviceId)
Add Solution Config for a given Service

Adds a solution config. Solution Config are service level objects, required for configuring the NXGI partner Service after deployment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SolutionConfig**](SolutionConfig.md)|  | 
  **serviceId** | **string**|  | 

### Return type

[**SolutionConfig**](SolutionConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateStaticHopBfdPeer**
> StaticHopBfdPeer CreateStaticHopBfdPeer(ctx, body, logicalRouterId)
Create a static hop BFD peer

Creates a BFD peer for static route. The required parameters includes peer IP address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StaticHopBfdPeer**](StaticHopBfdPeer.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**StaticHopBfdPeer**](StaticHopBfdPeer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSwitchingProfile**
> BaseSwitchingProfile CreateSwitchingProfile(ctx, body)
Create a Switching Profile

Creates a new, custom qos, port-mirroring, spoof-guard or port-security switching profile. You can override their default switching profile assignments by creating a new switching profile and assigning it to one or more logical switches. You cannot override the default ipfix or ip_discovery switching profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseSwitchingProfile**](BaseSwitchingProfile.md)|  | 

### Return type

[**BaseSwitchingProfile**](BaseSwitchingProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTraceflow**
> Traceflow CreateTraceflow(ctx, body)
Initiate a Traceflow Operation on the Specified Port

Initiate a Traceflow Operation on the Specified Port

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TraceflowRequest**](TraceflowRequest.md)|  | 

### Return type

[**Traceflow**](Traceflow.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteADhcpLease**
> DeleteADhcpLease(ctx, serverId, ip, mac)
Delete a single DHCP lease entry specified by ip and mac.

Delete a single DHCP lease entry specified by ip and mac.  The DHCP server matches the DHCP lease with the given ip address and the mac address. The matched lease entry will be deleted. If no lease matches, the request is ignored.  The DHCP lease to be deleted will be removed by the system from both active and standby node. The system will report error if the DHCP lease could not be removed from both nodes. If the DHCP lease could not be removed on either node, please check the DHCP server status. Once the DHCP server status is UP, please invoke the deletion API again to ensure the lease gets deleted from both nodes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **ip** | **string**| IPv4 or IPv6 address | 
  **mac** | **string**| MAC Address | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllCaptureSessionsDelete**
> PacketCaptureSessionList DeleteAllCaptureSessionsDelete(ctx, )
Delete all the packet capture sessions

Delete all the packet capture sessions. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PacketCaptureSessionList**](PacketCaptureSessionList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBGPCommunityList**
> DeleteBGPCommunityList(ctx, logicalRouterId, communityListId)
Delete a specific BGP community list from a logical router

Delete a specific BGP community list from a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **communityListId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBgpNeighbor**
> DeleteBgpNeighbor(ctx, logicalRouterId, id)
Delete a specific BGP Neighbor on a Logical Router

Delete a specific BGP Neighbor on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBridgeCluster**
> DeleteBridgeCluster(ctx, bridgeclusterId)
Delete a Bridge Cluster

Removes the specified Bridge Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeclusterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBridgeEndpoint**
> DeleteBridgeEndpoint(ctx, bridgeendpointId)
Delete a Bridge Endpoint

Deletes the specified Bridge Endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeendpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBridgeEndpointProfile**
> DeleteBridgeEndpointProfile(ctx, bridgeendpointprofileId)
Delete a Bridge Endpoint Profile

Deletes the specified Bridge Endpoint Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeendpointprofileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDADProfile**
> DeleteDADProfile(ctx, dadProfileId)
Delete DAD Profile

Delete DADProfile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dadProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpIpPool**
> DeleteDhcpIpPool(ctx, serverId, poolId)
Delete a DHCP server's IP pool

Delete a specific ip pool of a given logical DHCP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpProfile**
> DeleteDhcpProfile(ctx, profileId)
Delete a DHCP server profile

Delete a DHCP server profile specified by the profile id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpRelay**
> DeleteDhcpRelay(ctx, relayId)
Delete a DHCP Relay Service

Deletes the specified dhcp relay service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **relayId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpRelayProfile**
> DeleteDhcpRelayProfile(ctx, relayProfileId)
Delete a DHCP Relay Profile

Deletes the specified dhcp relay profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **relayProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpServer**
> DeleteDhcpServer(ctx, serverId)
Delete a DHCP server

Delete a logical DHCP server specified by server id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpStaticBinding**
> DeleteDhcpStaticBinding(ctx, serverId, bindingId)
Delete a static binding

Delete a specific static binding of a given logical DHCP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpV6IpPool**
> DeleteDhcpV6IpPool(ctx, serverId, poolId)
Delete a DHCP IPv6 server's IP pool

Delete a specific ip pool of a given logical DHCP IPv6 server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDhcpV6StaticBinding**
> DeleteDhcpV6StaticBinding(ctx, serverId, bindingId)
Delete a static binding for DHCP IPv6 server

Delete a specific static binding of a given logical DHCP IPv6 server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDnsForwarder**
> DeleteDnsForwarder(ctx, forwarderId)
Delete a specific DNS forwarder

Delete a specific DNS forwarder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFirewallProfile**
> DeleteFirewallProfile(ctx, profileId)
Delete a firewall profile.

Deletes a firewall profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPPrefixList**
> DeleteIPPrefixList(ctx, logicalRouterId, id)
Delete a specific IPPrefixList on a Logical Router

Deletes a specific IPPrefixList on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSecVPNDPDProfile**
> DeleteIPSecVPNDPDProfile(ctx, ipsecVpnDpdProfileId, optional)
Delete dead peer detection (DPD) profile

Delete dead peer detection (DPD) profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnDpdProfileId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSecVPNDPDProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSecVPNDPDProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSecVPNIKEProfile**
> DeleteIPSecVPNIKEProfile(ctx, ipsecVpnIkeProfileId, optional)
Delete custom IKE Profile

Delete custom IKE Profile. Profile can not be deleted if profile has references to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnIkeProfileId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSecVPNIKEProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSecVPNIKEProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSecVPNLocalEndpoint**
> DeleteIPSecVPNLocalEndpoint(ctx, ipsecVpnLocalEndpointId, optional)
Delete custom IPSec local endpoint

Delete custom IPSec local endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnLocalEndpointId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSecVPNLocalEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSecVPNLocalEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSecVPNPeerEndpoint**
> DeleteIPSecVPNPeerEndpoint(ctx, ipsecVpnPeerEndpointId, optional)
Delete custom IPSec VPN peer endpoint

Delete custom IPSec VPN peer endpoint. All references are strong references and dependent peer endpoints can not be deleted if being referenced.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnPeerEndpointId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSecVPNPeerEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSecVPNPeerEndpointOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSecVPNService**
> DeleteIPSecVPNService(ctx, ipsecVpnServiceId, optional)
Delete IPSec VPN service

Delete IPSec VPN service for given router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnServiceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSecVPNServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSecVPNServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSecVPNSession**
> DeleteIPSecVPNSession(ctx, ipsecVpnSessionId, optional)
Delete IPSec VPN session

Delete IPSec VPN session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnSessionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSecVPNSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSecVPNSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSecVPNTunnelProfile**
> DeleteIPSecVPNTunnelProfile(ctx, ipsecVpnTunnelProfileId, optional)
Delete custom IPSecTunnelProfile

Delete custom IPSec Tunnel Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnTunnelProfileId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSecVPNTunnelProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSecVPNTunnelProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIPSet**
> DeleteIPSet(ctx, ipSetId, optional)
Delete IPSet

Deletes the specified IPSet.  By default, if the IPSet is added to an NSGroup, it won't be deleted. In such situations, pass \"force=true\" as query param to force delete the IPSet. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSetId** | **string**| IPSet Id | 
 **optional** | ***ManagementPlaneApiApiDeleteIPSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteIPSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInstanceEndpoint**
> DeleteInstanceEndpoint(ctx, serviceId, serviceInstanceId, instanceEndpointId)
Delete a particular InstanceEndpoint.

Delete instance endpoint information for a given instace endpoint. Please make sure to delete all the Service Insertion Rules, which refer to this Endpoint as 'redirect_tos' target. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 
  **instanceEndpointId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpfixCollectorConfig**
> DeleteIpfixCollectorConfig(ctx, collectorConfigId)
Delete an existing IPFIX collector configuration

Delete an existing IPFIX collector configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectorConfigId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpfixCollectorUpmProfile**
> DeleteIpfixCollectorUpmProfile(ctx, ipfixCollectorProfileId)
Delete an existing IPFIX collector profile

Delete an existing IPFIX collector profile by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipfixCollectorProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpfixConfig**
> DeleteIpfixConfig(ctx, configId)
Delete an existing IPFIX configuration

Delete an existing IPFIX configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpfixUpmProfile**
> DeleteIpfixUpmProfile(ctx, ipfixProfileId)
Delete an existing IPFIX profile

Delete an existing IPFIX profile by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipfixProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteL2VpnService**
> DeleteL2VpnService(ctx, l2vpnServiceId, optional)
Delete a L2VPN service

Delete a specific L2VPN service. If there are any L2VpnSessions on this L2VpnService, those needs to be deleted first.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **l2vpnServiceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteL2VpnServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteL2VpnServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteL2VpnSession**
> DeleteL2VpnSession(ctx, l2vpnSessionId)
Delete a L2VPN session

Delete a specific L2VPN session. If there are any logical switch ports attached to it, those needs to be deleted first.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **l2vpnSessionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerApplicationProfile**
> DeleteLoadBalancerApplicationProfile(ctx, applicationProfileId)
Delete a load balancer application profile

Delete a load balancer application profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerClientSslProfile**
> DeleteLoadBalancerClientSslProfile(ctx, clientSslProfileId)
Delete a load balancer client-ssl profile

Delete a load balancer client-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientSslProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerMonitor**
> DeleteLoadBalancerMonitor(ctx, monitorId)
Delete a load balancer monitor

Delete a load balancer monitor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerPersistenceProfile**
> DeleteLoadBalancerPersistenceProfile(ctx, persistenceProfileId)
Delete a load balancer persistence profile

Delete a load balancer persistence profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **persistenceProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerPool**
> DeleteLoadBalancerPool(ctx, poolId)
Delete a load balancer pool

Delete a load balancer pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerRule**
> DeleteLoadBalancerRule(ctx, ruleId)
Delete a load balancer rule

Delete a load balancer rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerServerSslProfile**
> DeleteLoadBalancerServerSslProfile(ctx, serverSslProfileId)
Delete a load balancer server-ssl profile

Delete a load balancer server-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverSslProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerService**
> DeleteLoadBalancerService(ctx, serviceId)
Delete a load balancer service

Delete a load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerTcpProfile**
> DeleteLoadBalancerTcpProfile(ctx, tcpProfileId)
Delete a load balancer TCP profile

Delete a load balancer TCP profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tcpProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoadBalancerVirtualServer**
> DeleteLoadBalancerVirtualServer(ctx, virtualServerId, optional)
Delete a load balancer virtual server

Delete a load balancer virtual server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualServerId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteLoadBalancerVirtualServerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteLoadBalancerVirtualServerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteAssociatedRules** | **optional.Bool**| Delete associated rules | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogicalPort**
> DeleteLogicalPort(ctx, lportId, optional)
Delete a Logical Port

Deletes the specified logical switch port. By default, if logical port has attachments, or it is added to any NSGroup, the deletion will be failed. Option detach could be used for deleting logical port forcibly. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteLogicalPortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteLogicalPortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detach** | **optional.Bool**| force delete even if attached or referenced by a group | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogicalRouter**
> DeleteLogicalRouter(ctx, logicalRouterId, optional)
Delete a Logical Router

Deletes the specified logical router. You must delete associated logical router ports before you can delete a logical router. Otherwise use force delete which will delete all related ports and other entities associated with that LR. To force delete logical router pass force=true in query param. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteLogicalRouterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteLogicalRouterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascadeDeleteLinkedPorts** | **optional.Bool**| Flag to specify whether to delete related logical switch ports | [default to false]
 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogicalRouterPort**
> DeleteLogicalRouterPort(ctx, logicalRouterPortId, optional)
Delete a Logical Router Port

Deletes the specified logical router port. You must delete logical router ports before you can delete the associated logical router. To Delete Tier0 router link port you must have to delete attached tier1 router link port, otherwise pass \"force=true\" as query param to force delete the Tier0 router link port. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterPortId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteLogicalRouterPortOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteLogicalRouterPortOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascadeDeleteLinkedPorts** | **optional.Bool**| Flag to specify whether to delete related logical switch ports | [default to false]
 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLogicalSwitch**
> DeleteLogicalSwitch(ctx, lswitchId, optional)
Delete a Logical Switch

Removes a logical switch from the associated overlay or VLAN transport zone. By default, a logical switch cannot be deleted if there are logical ports on the switch, or it is added to a NSGroup. Cascade option can be used to delete all ports and the logical switch. Detach option can be used to delete the logical switch forcibly. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteLogicalSwitchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteLogicalSwitchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**| Delete a Logical Switch and all the logical ports in it, if none of the logical ports have any attachment.  | [default to false]
 **detach** | **optional.Bool**| Force delete a logical switch | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMACSet**
> DeleteMACSet(ctx, macSetId, optional)
Delete MACSet

Deletes the specified MACSet. By default, if the MACSet is added to an NSGroup, it won't be deleted. In such situations, pass \"force=true\" as query param to force delete the MACSet. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macSetId** | **string**| MACSet Id | 
 **optional** | ***ManagementPlaneApiApiDeleteMACSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteMACSetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManualHealthCheck**
> DeleteManualHealthCheck(ctx, manualHealthCheckId)
Delete an existing manual health check

Delete an existing manual health check by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **manualHealthCheckId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMetadataProxy**
> DeleteMetadataProxy(ctx, proxyId)
Delete a metadata proxy

Delete a metadata proxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNDRAProfile**
> DeleteNDRAProfile(ctx, ndRaProfileId)
Delete NDRA Profile

Delete NDRAProfile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ndRaProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNSGroup**
> DeleteNSGroup(ctx, nsGroupId, optional)
Delete NSGroup

Deletes the specified NSGroup. By default, if the NSGroup is added to another NSGroup, it won't be deleted. In such situations, pass \"force=true\" as query param to force delete the NSGroup. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiDeleteNSGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteNSGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNSProfile**
> DeleteNSProfile(ctx, nsProfileId, optional)
Delete NSProfile

Deletes the specified NSProfile. By default, if the NSProfile is consumed in a Firewall rule, it won't get deleted. In such situations, pass \"force=true\" as query param to force delete the NSProfile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsProfileId** | **string**| NSProfile Id | 
 **optional** | ***ManagementPlaneApiApiDeleteNSProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteNSProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNSService**
> DeleteNSService(ctx, nsServiceId, optional)
Delete NSService

Deletes the specified NSService. By default, if the NSService is being referred in an NSServiceGroup, it can't be deleted. In such situations, pass \"force=true\" as a parameter to force delete the NSService. System defined NSServices can't be deleted using \"force\" flag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsServiceId** | **string**| NSService Id | 
 **optional** | ***ManagementPlaneApiApiDeleteNSServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteNSServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNSServiceGroup**
> DeleteNSServiceGroup(ctx, nsServiceGroupId, optional)
Delete NSServiceGroup

Deletes the specified NSServiceGroup. By default, if the NSServiceGroup is consumed in a Firewall rule, it won't get deleted. In such situations, pass \"force=true\" as query param to force delete the NSServiceGroup. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsServiceGroupId** | **string**| NSServiceGroup Id | 
 **optional** | ***ManagementPlaneApiApiDeleteNSServiceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteNSServiceGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNatRule**
> DeleteNatRule(ctx, logicalRouterId, ruleId)
Delete a specific NAT rule from a logical router

Delete a specific NAT rule from a logical router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePBRRule**
> DeletePBRRule(ctx, sectionId, ruleId)
Delete an Existing Rule

Delete existing PBR rule in a PBR section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePBRSection**
> DeletePBRSection(ctx, sectionId, optional)
Delete an Existing Section and Its Associated Rules

Removes PBR section from the system. PBR section with rules can only be deleted by passing \"cascade=true\" parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeletePBRSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeletePBRSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**| Flag to cascade delete of this object to all it&#x27;s child objects. | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePacketCaptureSessionDelete**
> PacketCaptureSession DeletePacketCaptureSessionDelete(ctx, sessionId)
Delete the packet capture session by session id.

Before calling this method, terminate any running capture session. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**| Packet capture session id | 

### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePortMirroringSession**
> DeletePortMirroringSession(ctx, mirrorSessionId)
Delete the mirror session

Delete the mirror session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mirrorSessionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRouteMap**
> DeleteRouteMap(ctx, logicalRouterId, id)
Delete a specific RouteMap on a Logical Router

Deletes a specific RouteMap on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRule**
> DeleteRule(ctx, sectionId, ruleId)
Delete an Existing Rule

Delete existing firewall rule in a firewall section. Deleting firewall rule in a section modifies parent section and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSIServiceProfile**
> DeleteSIServiceProfile(ctx, serviceId, serviceProfileId)
Delete a particular ServiceProfile.

Delete service profile for a given service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSection**
> DeleteSection(ctx, sectionId, optional)
Delete an Existing Section and Its Associated Rules

Removes firewall section from the system. Firewall section with rules can only be deleted by passing \"cascade=true\" parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**| Flag to cascade delete of this object to all it&#x27;s child objects. | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceAttachment**
> DeleteServiceAttachment(ctx, serviceAttachmentId)
Delete an existing service attachment

Delete existing service attachment from system. Before deletion, please make sure that, no instance endpoints are connected to this attachment. In turn no appliance should be connected to this attachment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceAttachmentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceChain**
> DeleteServiceChain(ctx, serviceChainId)
Delete a Service Chain.

Delete a particular service chain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceChainId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceConfig**
> DeleteServiceConfig(ctx, configSetId)
Delete Service Config

Deletes the specified service config 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configSetId** | **string**| service Ccnfig Id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceDeployment**
> DeleteServiceDeployment(ctx, serviceId, serviceDeploymentId, optional)
Remove service deployment

Remove the service deployment. Will remove all the Service VMs that were created as part of this deployment. User can send optional force delete option which will force remove the deployment, but should be used only when the regular delete is not working. Regular delete will ensure proper cleanup of Service VMs and related objects. Directly calling this API without trying regular undeploy will result in unexpected results, and orphan objects. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceDeploymentId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteServiceDeploymentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteServiceDeploymentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceInsertionRule**
> DeleteServiceInsertionRule(ctx, sectionId, ruleId)
Delete an Existing Rule

Delete existing serviceinsertion rule in a serviceinsertion section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceInsertionSection**
> DeleteServiceInsertionSection(ctx, sectionId, optional)
Delete an Existing Section and Its Associated Rules

Removes serviceinsertion section from the system. ServiceInsertion section with rules can only be deleted by passing \"cascade=true\" parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteServiceInsertionSectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteServiceInsertionSectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**| Flag to cascade delete of this object to all it&#x27;s child objects. | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceInsertionService**
> DeleteServiceInsertionService(ctx, serviceId, optional)
Delete an existing Service and the Service-Instance associated with it.

Removes Service-Insertion Service from the system. A Service with Service-Instances can only be deleted by passing \"cascade=true\" parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteServiceInsertionServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteServiceInsertionServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cascade** | **optional.Bool**| Flag to cascade delete all the child objects, associated with it. | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceInstance**
> DeleteServiceInstance(ctx, serviceId, serviceInstanceId)
Delete an existing Service-Instance

Delete existing Service-Instance for a given Service-Insertion Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceManager**
> DeleteServiceManager(ctx, serviceManagerId)
Delete service manager

Delete service-manager which is registered with NSX with basic details like name, username, password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceManagerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceVMsDelete**
> DeleteServiceVMsDelete(ctx, serviceId, serviceInstanceId)
Remove service VMs either as standalone or HA

Undeploy one service VM as standalone or two service VMs as HA. Associated deployment information and instance runtime will also be deleted once service VMs have been un-deployed successfully. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSolutionConfig**
> DeleteSolutionConfig(ctx, serviceId, solutionConfigId)
Deletes solution config information.

Deletes solution config information for a given service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **solutionConfigId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStaticHopBfdPeer**
> DeleteStaticHopBfdPeer(ctx, logicalRouterId, bfdPeerId, optional)
Delete a specified static route BFD peer cofigured on a specified logical router

Deletes the specified BFD peer present on specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **bfdPeerId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteStaticHopBfdPeerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteStaticHopBfdPeerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStaticRoute**
> DeleteStaticRoute(ctx, logicalRouterId, id)
Delete a specific Static Route on a Logical Router

Deletes a specific static route on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSwitchingProfile**
> DeleteSwitchingProfile(ctx, switchingProfileId, optional)
Delete a Switching Profile

Deletes the specified switching profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **switchingProfileId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiDeleteSwitchingProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiDeleteSwitchingProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unbind** | **optional.Bool**| force unbinding of logical switches and ports from a switching profile | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTraceflow**
> DeleteTraceflow(ctx, traceflowId)
Delete the Traceflow round

Delete the Traceflow round

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **traceflowId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVendorTemplate**
> DeleteVendorTemplate(ctx, serviceId, vendorTemplateId)
Delete a particular vendor tempalte.

Delete vendor template information for a given service. Please make sure to delete all the Service Profile(s), which refer to this vendor tempalte before deleting the template itself. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **vendorTemplateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployService**
> ServiceDeployment DeployService(ctx, body, serviceId)
Deploys a particular service

This will deploy a particular service on a given cluster / host. Internally multiple service instance can be created during the deployment. If there are no issues in the parameters, the call returns immediately, and the service VMs will be deployed asynchronously. To get the overall status of the deployment or to get the status of individual service vm, please use the deployment status APIs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceDeployment**](ServiceDeployment.md)|  | 
  **serviceId** | **string**|  | 

### Return type

[**ServiceDeployment**](ServiceDeployment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeployServiceVMsDeploy**
> DeployServiceVMsDeploy(ctx, serviceId, serviceInstanceId)
Deploy and set up service VMs either as standalone or HA

Deploys one service VM as standalone, or two service VMs as HA where one VM is active and another one is standby.  During the deployment of service VMs, service will be set up based on deployment events using callbacks. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableDnsForwarderDisable**
> DisableDnsForwarderDisable(ctx, forwarderId)
Disable the DNS forwarder.

Disable the DNS forwarder if the forwarder is currently enbled. If the DNS forwarder is already disabled, the forwarder will not be re-disabled.  Please note, once a DNS forwarder is disabled then enabled, the previous DNS forwarder statistics counters will be reset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableFirewallOnTargetResourceDisableFirewall**
> TargetResourceStatus DisableFirewallOnTargetResourceDisableFirewall(ctx, contextType, id)
Disable firewall on target resource in dfw context

Disable firewall on target resource in dfw context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextType** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**TargetResourceStatus**](TargetResourceStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EffectiveProfiles**
> EffectiveProfileListResult EffectiveProfiles(ctx, resourceId, resourceType, optional)
Get Effective Profiles for an Entity

Returns the effective profiles applied to the specified Resource. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceId** | **string**| The resource for which the effective profiles are to be fetched | 
  **resourceType** | **string**| Valid Resource type in effective profiles API | 
 **optional** | ***ManagementPlaneApiApiEffectiveProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiEffectiveProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveProfileListResult**](EffectiveProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableDnsForwarderEnable**
> EnableDnsForwarderEnable(ctx, forwarderId)
Enable the DNS forwarder.

Enable the DNS forwarder if the forwarder is currently disabled. If the DNS forwarder is already enabled, the forwarder will not be re-enabled.  Please note, once a DNS forwarder is disabled then enabled, the previous DNS forwarder statistics counters will be reset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableFirewallOnTargetResourceEnableFirewall**
> TargetResourceStatus EnableFirewallOnTargetResourceEnableFirewall(ctx, contextType, id)
Enable firewall on target resource in dfw context

Enable firewall on target resource in dfw context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextType** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**TargetResourceStatus**](TargetResourceStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAffectedVms**
> IdsVmList GetAffectedVms(ctx, body, optional)
Get the list of the VMs affected for that signature

Get the list of the VMs affected pertaining to a specific signature. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsEventDataRequest**](IdsEventDataRequest.md)|  | 
 **optional** | ***ManagementPlaneApiApiGetAffectedVmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetAffectedVmsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.**|  | 
 **sortBy** | **optional.**| Field by which records are sorted | 

### Return type

[**IdsVmList**](IdsVmList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllIdsEvents**
> IdsEventsBySignatureResult GetAllIdsEvents(ctx, body)
Get the list of the IDS events that are detected, grouped by signature id.

Get the list of the IDS events that are detected with the total number of intrusions detected, their severity and the time they occurred, grouped by signature id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsEventDataRequest**](IdsEventDataRequest.md)|  | 

### Return type

[**IdsEventsBySignatureResult**](IDSEventsBySignatureResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssociations**
> AssociationListResult GetAssociations(ctx, associatedResourceType, resourceId, resourceType, optional)
Get ResourceReference objects to which the given resource belongs to 

Returns information about resources that are associated with the given resource. Id and type of the resource for which associated resources are to be fetched are to be specified as query parameter in the URI. Resource type of the associated resources must be specified as query parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **associatedResourceType** | **string**| Resource type valid for use as target in association API | 
  **resourceId** | **string**| The resource for which associated resources are to be fetched | 
  **resourceType** | **string**| Resource type valid for use as source in association API | 
 **optional** | ***ManagementPlaneApiApiGetAssociationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetAssociationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **fetchAncestors** | **optional.Bool**| Fetch complete list of associated resources considering containment and nesting  | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**AssociationListResult**](AssociationListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighborAdvertisedRoutes**
> BgpNeighborRouteDetails GetBgpNeighborAdvertisedRoutes(ctx, logicalRouterId, neighborId)
Get BGP neighbor advertised routes

Returns routes advertised by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**BgpNeighborRouteDetails**](BgpNeighborRouteDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighborAdvertisedRoutesInCsvFormatCsv**
> BgpNeighborRouteDetailsInCsvFormat GetBgpNeighborAdvertisedRoutesInCsvFormatCsv(ctx, logicalRouterId, neighborId)
Get BGP neighbor advertised routes in CSV format 

Returns routes advertised by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled in CSV format. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**BgpNeighborRouteDetailsInCsvFormat**](BgpNeighborRouteDetailsInCsvFormat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighborRoutes**
> BgpNeighborRouteDetails GetBgpNeighborRoutes(ctx, logicalRouterId, neighborId)
Get BGP neighbor learned routes

Returns routes learned by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**BgpNeighborRouteDetails**](BgpNeighborRouteDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighborRoutesInCsvFormatCsv**
> BgpNeighborRouteDetailsInCsvFormat GetBgpNeighborRoutesInCsvFormatCsv(ctx, logicalRouterId, neighborId)
Get BGP neighbor learned routes in CSV format 

Returns routes learned by BGP neighbor from all edge transport nodes on which this neighbor is currently enabled in CSV format. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**BgpNeighborRouteDetailsInCsvFormat**](BgpNeighborRouteDetailsInCsvFormat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBgpNeighborsStatus**
> BgpNeighborsStatusListResult GetBgpNeighborsStatus(ctx, logicalRouterId, optional)
Get the status of all the BGP neighbors for the Logical Router of the given id

Returns the status of all the BGP neighbors for the Logical Router of the given id. To get BGP neighbors status for the logical router from particular node, parameter \"transport_node_id=<transportnode_id>\" needs to be specified. Query parameter \"source=realtime\" is the only supported source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetBgpNeighborsStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetBgpNeighborsStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| Transport node id | 

### Return type

[**BgpNeighborsStatusListResult**](BgpNeighborsStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBridgeCluster**
> BridgeCluster GetBridgeCluster(ctx, bridgeclusterId)
Get Information about a bridge cluster

Returns information about a specified bridge cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeclusterId** | **string**|  | 

### Return type

[**BridgeCluster**](BridgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBridgeClusterStatus**
> BridgeClusterStatus GetBridgeClusterStatus(ctx, clusterId, optional)
Returns status of a specified Bridge Cluster

Get the status for the Bridge Cluster of the given cluster id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetBridgeClusterStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetBridgeClusterStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**BridgeClusterStatus**](BridgeClusterStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBridgeEndpoint**
> BridgeEndpoint GetBridgeEndpoint(ctx, bridgeendpointId)
Get Information about a bridge endpoint

Returns information about a specified bridge endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeendpointId** | **string**|  | 

### Return type

[**BridgeEndpoint**](BridgeEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBridgeEndpointProfile**
> BridgeEndpointProfile GetBridgeEndpointProfile(ctx, bridgeendpointprofileId)
Get Information about a bridge endpoint Profile

Returns information about a specified bridge endpoint profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeendpointprofileId** | **string**|  | 

### Return type

[**BridgeEndpointProfile**](BridgeEndpointProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBridgeEndpointStatistics**
> BridgeEndpointStatistics GetBridgeEndpointStatistics(ctx, endpointId, optional)
Returns statistics of a specified Bridge Endpoint

Get the statistics for the Bridge Endpoint of the given Endpoint id (endpoint-id)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetBridgeEndpointStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetBridgeEndpointStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**BridgeEndpointStatistics**](BridgeEndpointStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBridgeEndpointStatus**
> BridgeEndpointStatus GetBridgeEndpointStatus(ctx, endpointId, optional)
Returns status of a specified Bridge Endpoint

Get the status for the Bridge Endpoint of the given Endpoint id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetBridgeEndpointStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetBridgeEndpointStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**BridgeEndpointStatus**](BridgeEndpointStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCaptureFile**
> GetCaptureFile(ctx, sessionId)
Get packet capture file

You must provide the request header \"Accept:application/octet-stream\" when calling this API. The capture file can only be found in MP which receives the capture request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**| Packet capture session id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComputeCollectionStatusById**
> IdfwComputeCollectionStatus GetComputeCollectionStatusById(ctx, computeCollectionExtId)
Get list of compute collections and status.

Retrieve the compute collection status by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **computeCollectionExtId** | **string**|  | 

### Return type

[**IdfwComputeCollectionStatus**](IdfwComputeCollectionStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpIpPoolState**
> ConfigurationState GetDhcpIpPoolState(ctx, serverId, poolId, optional)
Get the realized state of a dhcp ip pool

Return realized state information of a dhcp ip pool. After a dhcp ip pool is created or updated, you can invoke this API to get the realization information of the ip pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **poolId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetDhcpIpPoolStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetDhcpIpPoolStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpLeaseInfo**
> DhcpLeases GetDhcpLeaseInfo(ctx, serverId, optional)
Get specific leases of a given dhcp server

Get specific leases of a given dhcp server. As a dhcp server could manage millions of leases, the API has to limit the number of the returned leases via two mutually-excluded request parameters, i.e. \"pool_id\" and \"address\". Either a \"pool_id\" or an \"address\" can be provided, but not both in a same call.  If a \"pool_id\" is specified, the leases of the specific pool are returned. If an \"address\" is specified, only the lease(s) represented y this address is(are) returned. The \"address\" can be a single IP, an ip-range, or a mac address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetDhcpLeaseInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetDhcpLeaseInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **optional.String**| can be an ip address, or an ip range, or a mac address | 
 **poolId** | **optional.String**| The uuid of dhcp ip pool | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**DhcpLeases**](DhcpLeases.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpServerState**
> ConfigurationState GetDhcpServerState(ctx, serverId, optional)
Get the realized state of a dhcp server

Return realized state information of a dhcp server. After a dhcp server is created or updated, you can invoke this API to get the realization information of the server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetDhcpServerStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetDhcpServerStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpStaticBindingState**
> ConfigurationState GetDhcpStaticBindingState(ctx, serverId, bindingId, optional)
Get the realized state of a dhcp static binding

Return realized state information of a dhcp static binding. After a dhcp static binding is created or updated, you can invoke this API to get the realization information of the static binding. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **bindingId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetDhcpStaticBindingStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetDhcpStaticBindingStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpStatistics**
> DhcpStatistics GetDhcpStatistics(ctx, serverId)
Get DHCP statistics with given dhcp server id

Returns the statistics of the given dhcp server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 

### Return type

[**DhcpStatistics**](DhcpStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDhcpStatus**
> DhcpServerStatus GetDhcpStatus(ctx, serverId)
Get DHCP service status with given dhcp server id

Returns the service status of the given dhcp server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 

### Return type

[**DhcpServerStatus**](DhcpServerStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDnsForwarderState**
> ConfigurationState GetDnsForwarderState(ctx, forwarderId, optional)
Get the realized state of a DNS forwarder

Return the realized state information of a DNS forwarder. After a DNS forwarder was created or updated, you can invoke this API to check the realization state of the forwarder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetDnsForwarderStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetDnsForwarderStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDnsForwarderStatistics**
> DnsForwarderStatistics GetDnsForwarderStatistics(ctx, forwarderId)
Get statistics of given dns forwarder

Returns the statistics of the given dns forwarder specified by forwarder id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 

### Return type

[**DnsForwarderStatistics**](DnsForwarderStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDnsForwarderStatus**
> DnsForwarderStatus GetDnsForwarderStatus(ctx, forwarderId)
Get current status of the given DNS forwarder

Returns the current status of the given DNS forwarder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 

### Return type

[**DnsForwarderStatus**](DnsForwarderStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveActiveDirectoryGroups**
> EffectiveMemberResourceListResult GetEffectiveActiveDirectoryGroups(ctx, nsGroupId, optional)
Get Effective Directory Groups of the specified NSGroup.

Returns effective directory groups which are members of the specified NSGroup. This API is applicable only for NSGroups containing DirectoryGroup member type. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveActiveDirectoryGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveActiveDirectoryGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberResourceListResult**](EffectiveMemberResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveCloudNativeServiceInstances**
> EffectiveMemberResourceListResult GetEffectiveCloudNativeServiceInstances(ctx, nsGroupId, optional)
Get Effective Cloud Native Service Instances of the specified NSGroup.

Returns effective cloud native service instances of the specified NSGroup. This API is applicable only for NSGroups containing CloudNativeServiceInstance member type. For NSGroups containing other member types,it returns an empty list. target_id in response is external_id of CloudNativeServiceInstance 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveCloudNativeServiceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveCloudNativeServiceInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberResourceListResult**](EffectiveMemberResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveIPAddressMembers**
> EffectiveIpAddressMemberListResult GetEffectiveIPAddressMembers(ctx, nsGroupId, optional)
Get Effective IPAddress translated from the NSGroup

Returns effective ip address members of the specified NSGroup. This API is applicable only for NSGroups containing either VirtualMachine, VIF, LogicalSwitch, LogicalPort or IPSet member type. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveIPAddressMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveIPAddressMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveIpAddressMemberListResult**](EffectiveIPAddressMemberListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveIPSetMembers**
> EffectiveMemberResourceListResult GetEffectiveIPSetMembers(ctx, nsGroupId, optional)
Get Effective IPSets of the specified NSGroup.

Returns effective IPSets which are members of the specified NSGroup. This API is applicable only for NSGroups containing IPSet member type. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveIPSetMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveIPSetMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberResourceListResult**](EffectiveMemberResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveLogicalPortMembers**
> EffectiveMemberResourceListResult GetEffectiveLogicalPortMembers(ctx, nsGroupId, optional)
Get Effective Logical Ports translated from the NSgroup

Returns effective logical port members of the specified NSGroup. This API is applicable only for NSGroups containing either VirtualMachines, LogicalSwitch or LogicalPort member types.For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveLogicalPortMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveLogicalPortMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberResourceListResult**](EffectiveMemberResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveLogicalSwitchMembers**
> EffectiveMemberResourceListResult GetEffectiveLogicalSwitchMembers(ctx, nsGroupId, optional)
Get Effective switch members translated from the NSGroup

Returns effective logical switch members of the specified NSGroup. This API is applicable for NSGroups containing LogicalSwitch members. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveLogicalSwitchMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveLogicalSwitchMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberResourceListResult**](EffectiveMemberResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectivePhysicalServerMembers**
> EffectiveMemberResourceListResult GetEffectivePhysicalServerMembers(ctx, nsGroupId, optional)
Get Effective Physical Server Memebers of the specified NSGroup.

Returns effective physical server members of the specified NSGroup. This API is applicable only for NSGroups containing Physical Server member type. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectivePhysicalServerMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectivePhysicalServerMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberResourceListResult**](EffectiveMemberResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveTransportNodeMembers**
> EffectiveMemberResourceListResult GetEffectiveTransportNodeMembers(ctx, nsGroupId, optional)
Get effective transport node members translated from the NSGroup

Returns effective transport node members of the specified NSGroup. This API is applicable only for NSGroups containing TransportNode member type. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveTransportNodeMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveTransportNodeMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberResourceListResult**](EffectiveMemberResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveVIFMembers**
> VirtualNetworkInterfaceListResult GetEffectiveVIFMembers(ctx, nsGroupId, optional)
Get effective VIF members translated from the NSGroup

Returns effective VIF members of the specified NSGroup. This API is applicable only for NSGroups containing either VirtualMachines or VIF member type. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveVIFMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveVIFMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VirtualNetworkInterfaceListResult**](VirtualNetworkInterfaceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEffectiveVirtualMachineMembers**
> VirtualMachineListResult GetEffectiveVirtualMachineMembers(ctx, nsGroupId, optional)
Get Effective Virtual Machine members of the specified NSGroup.

Returns effective virtual machine members of the specified NSGroup. This API is applicable only for NSGroups containing VirtualMachine member type. For NSGroups containing other member types,it returns an empty list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetEffectiveVirtualMachineMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetEffectiveVirtualMachineMembersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VirtualMachineListResult**](VirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnabledComputeCollection**
> IdfwEnabledComputeCollection GetEnabledComputeCollection(ctx, ccExtId)
Get IDFW compute collection.

Get enable/disable status of individual compute collections for IDFW. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ccExtId** | **string**|  | 

### Return type

[**IdfwEnabledComputeCollection**](IdfwEnabledComputeCollection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExcludeList**
> ExcludeList GetExcludeList(ctx, )
Get list of entities in exclude list

Get list of entities in exclude list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ExcludeList**](ExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFailedDnsQueries**
> DnsFailedQueries GetFailedDnsQueries(ctx, forwarderId, optional)
Get the recent failed DNS queries

Return the given count of recent failed DNS queries from DNS forwarder. Since the DNS forwarder is running in Acitve/Standby HA mode on transport nodes, the given count of queries will be returned from each nodes. Hence the total queries returned could be doubled. If no count is specified, 100 recent failed queries are returned. If the recent failures is less than the given count, all the failures will be returned. The maximum count is 1,000. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetFailedDnsQueriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetFailedDnsQueriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **optional.Int64**| The count of the failed DNS queries | [default to 100]

### Return type

[**DnsFailedQueries**](DnsFailedQueries.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallProfile**
> BaseFirewallProfile GetFirewallProfile(ctx, profileId)
Get all firewall session timer profiles.

Return firewall session timer profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**BaseFirewallProfile**](BaseFirewallProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallSectionStats**
> FirewallStatsList GetFirewallSectionStats(ctx, sectionId, optional)
Get Firewall section level statistics section

Get aggregated statistics for all rules for a given firewall section. The API only supports access to cached (source=cached) statistical data collected offline in the system. Data includes total number of packets, bytes, sessions counters and popularity index for a firewall rule and overall session count, max session count and max popularity index for all firewall rules on transport nodes or edge nodes. Aggregated statistics like maximum popularity index, maximum session count and total session count are computed with lower frequency compared to individual generic rule statistics, hence they may have a computation delay up to 15 minutes to reflect in response to this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetFirewallSectionStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetFirewallSectionStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**FirewallStatsList**](FirewallStatsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallStats**
> FirewallStats GetFirewallStats(ctx, sectionId, ruleId, optional)
Get Firewall rule level statistics

Get aggregated statistics for a rule for given firewall section. The API only supports access to cached (source=cached) statistical data collected offline in the system. Data includes total number of packets, bytes, sessions counters and popularity index for a firewall rule and overall session count, max session count and max popularity index for all firewall rules on transport nodes or edge nodes. Aggregated statistics like maximum popularity index, maximum session count and total session count are computed with lower frequency compared to individual generic rule statistics, hence they may have a computation delay up to 15 minutes to reflect in response to this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetFirewallStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetFirewallStatsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**FirewallStats**](FirewallStats.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallStatus**
> FirewallStatus GetFirewallStatus(ctx, contextType)
Get firewall global status for dfw context

Get firewall global status for dfw context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextType** | **string**|  | 

### Return type

[**FirewallStatus**](FirewallStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirewallStatusOnTargetResource**
> TargetResourceStatus GetFirewallStatusOnTargetResource(ctx, contextType, id)
Get firewall status for target resource in dfw context

Get firewall status for target resource in dfw context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextType** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**TargetResourceStatus**](TargetResourceStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForwardingPath**
> PortConnectionEntities GetForwardingPath(ctx, lportId, peerPortId)
Get networking entities between two logical ports with VIF attachment

Get networking entities between two logical ports with VIF attachment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**| ID of source port | 
  **peerPortId** | **string**| ID of peer port | 

### Return type

[**PortConnectionEntities**](PortConnectionEntities.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIDSProfile**
> IdsProfile GetIDSProfile(ctx, idsProfileId)
Get IDSProfile

Returns information about the specified IDSProfile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idsProfileId** | **string**| IDSProfile Id | 

### Return type

[**IdsProfile**](IDSProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPAddresses**
> IpAddressElementListResult GetIPAddresses(ctx, ipSetId)
Get all IPAddresses in a IPSet

List all IP addresses in a IPSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSetId** | **string**| IP Set Id | 

### Return type

[**IpAddressElementListResult**](IPAddressElementListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNDPDProfile**
> IpSecVpndpdProfile GetIPSecVPNDPDProfile(ctx, ipsecVpnDpdProfileId)
Get IPSec dead peer detection (DPD) profile

Get IPSec dead peer detection (DPD) profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnDpdProfileId** | **string**|  | 

### Return type

[**IpSecVpndpdProfile**](IPSecVPNDPDProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNIKEProfile**
> IpSecVpnikeProfile GetIPSecVPNIKEProfile(ctx, ipsecVpnIkeProfileId)
Get IKE Profile

Get custom IKE Profile, given the particular id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnIkeProfileId** | **string**|  | 

### Return type

[**IpSecVpnikeProfile**](IPSecVPNIKEProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNIKEService**
> IpSecVpnikeServiceSummary GetIPSecVPNIKEService(ctx, serviceId, optional)
Cumulative statistics for one IKE service instance

Cumulative statistics for one IKE service instance. Query parameter source supports only cached mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetIPSecVPNIKEServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetIPSecVPNIKEServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**IpSecVpnikeServiceSummary**](IPSecVPNIKEServiceSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNIKESessionStatus**
> IpSecVpnSessionStatus GetIPSecVPNIKESessionStatus(ctx, sessionId, optional)
Get IPSec VPN IKE session status

List status of IPSec session. Query parameter source supports both realtime and cached mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetIPSecVPNIKESessionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetIPSecVPNIKESessionStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**IpSecVpnSessionStatus**](IPSecVPNSessionStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNLocalEndpoint**
> IpSecVpnLocalEndpoint GetIPSecVPNLocalEndpoint(ctx, ipsecVpnLocalEndpointId)
Get custom IPSec local endpoint

Get custom IPSec local endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnLocalEndpointId** | **string**|  | 

### Return type

[**IpSecVpnLocalEndpoint**](IPSecVPNLocalEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNPeerEndpoint**
> IpSecVpnPeerEndpoint GetIPSecVPNPeerEndpoint(ctx, ipsecVpnPeerEndpointId)
Get IPSec VPN peer endpoint

Get custom IPSec VPN peer endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnPeerEndpointId** | **string**|  | 

### Return type

[**IpSecVpnPeerEndpoint**](IPSecVPNPeerEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNPeerEndpointWithPSKShowSensitiveData**
> IpSecVpnPeerEndpoint GetIPSecVPNPeerEndpointWithPSKShowSensitiveData(ctx, ipsecVpnPeerEndpointId)
Get IPSec VPN peer endpoint with PSK

Get custom IPSec VPN peer endpoint with PSK.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnPeerEndpointId** | **string**|  | 

### Return type

[**IpSecVpnPeerEndpoint**](IPSecVPNPeerEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNService**
> IpSecVpnService GetIPSecVPNService(ctx, ipsecVpnServiceId)
Get IPSec VPN service

Get IPSec VPN service for given logical router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnServiceId** | **string**|  | 

### Return type

[**IpSecVpnService**](IPSecVPNService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNSession**
> IpSecVpnSession GetIPSecVPNSession(ctx, ipsecVpnSessionId)
Fetch IPSec VPN session

Fetch IPSec VPN session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnSessionId** | **string**|  | 

### Return type

[**IpSecVpnSession**](IPSecVPNSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNSessionState**
> IpSecVpnSessionState GetIPSecVPNSessionState(ctx, ipsecVpnSessionId, optional)
Get the Realized State of a IPSec VPN Session

Return realized state information of a ipsec vpn session. Any configuration update that affects the ipsec vpn session can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of ipsec vpn session, service, endpoints, profiles, etc. It will return a service disabled error, if the ipsec vpn service associated with the session is disabled. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnSessionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetIPSecVPNSessionStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetIPSecVPNSessionStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**IpSecVpnSessionState**](IPSecVPNSessionState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNSessionStatistics**
> IpSecVpnSessionStatistics GetIPSecVPNSessionStatistics(ctx, sessionId, optional)
Get IPSec VPN session statistics

Get statistics of a vpn session across all tunnels and IKE session. Query parameter \"source=realtime\" is the only supported source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetIPSecVPNSessionStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetIPSecVPNSessionStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**IpSecVpnSessionStatistics**](IPSecVPNSessionStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNSessionSummary**
> IpSecVpnSessionSummary GetIPSecVPNSessionSummary(ctx, optional)
VPN session summary

VPN session summary gets summary per vpn sessions and IKE session. Query parameter source supports only cached mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiGetIPSecVPNSessionSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetIPSecVPNSessionSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **siteId** | **optional.String**| Peer site id | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**IpSecVpnSessionSummary**](IPSecVPNSessionSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIPSecVPNTunnelProfile**
> IpSecVpnTunnelProfile GetIPSecVPNTunnelProfile(ctx, ipsecVpnTunnelProfileId)
Get IPSec tunnel profile

Get custom IPSec Tunnel Profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnTunnelProfileId** | **string**|  | 

### Return type

[**IpSecVpnTunnelProfile**](IPSecVPNTunnelProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIdsDashboardSummary**
> IdsSummaryListResult GetIdsDashboardSummary(ctx, body, optional)
Get the summary of the intrusions that were detected.

Get the summary of all the intrusions that are detected grouped by signature with details including signature name, id, severity, attack type, protocol, first and recent occurence, and affected users and VMs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdsEventDataRequest**](IdsEventDataRequest.md)|  | 
 **optional** | ***ManagementPlaneApiApiGetIdsDashboardSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetIdsDashboardSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.**|  | 
 **sortBy** | **optional.**| Field by which records are sorted | 

### Return type

[**IdsSummaryListResult**](IDSSummaryListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInstanceEndpoint**
> InstanceEndpoint GetInstanceEndpoint(ctx, serviceId, serviceInstanceId, instanceEndpointId)
Get a particular instance endpoint for a service instance.

Returns detailed Endpoint information for a given InstanceEndpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 
  **instanceEndpointId** | **string**|  | 

### Return type

[**InstanceEndpoint**](InstanceEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpfixCollectorConfig**
> IpfixCollectorConfig GetIpfixCollectorConfig(ctx, collectorConfigId)
Get an existing IPFIX collector configuration

Get an existing IPFIX collector configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectorConfigId** | **string**|  | 

### Return type

[**IpfixCollectorConfig**](IpfixCollectorConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpfixCollectorUpmProfile**
> IpfixCollectorUpmProfile GetIpfixCollectorUpmProfile(ctx, ipfixCollectorProfileId)
Get an existing IPFIX collector profile

Get an existing IPFIX collector profile by profile ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipfixCollectorProfileId** | **string**|  | 

### Return type

[**IpfixCollectorUpmProfile**](IpfixCollectorUpmProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpfixConfig**
> IpfixConfig GetIpfixConfig(ctx, configId)
Get an existing IPFIX configuration

Get an existing IPFIX configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configId** | **string**|  | 

### Return type

[**IpfixConfig**](IpfixConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpfixObsPoints**
> IpfixObsPointsListResult GetIpfixObsPoints(ctx, )
Get the list of IPFIX observation points

Deprecated - Please use /ipfix-profiles for switch IPFIX profile and /ipfix-collector-profiles for IPFIX collector profile. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IpfixObsPointsListResult**](IpfixObsPointsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetIpfixUpmProfile**
> IpfixUpmProfile GetIpfixUpmProfile(ctx, ipfixProfileId)
Get an existing IPFIX profile

Get an existing IPFIX profile by profile ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipfixProfileId** | **string**|  | 

### Return type

[**IpfixUpmProfile**](IpfixUpmProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2ForwarderRemoteMacs**
> L2ForwarderRemoteMacs GetL2ForwarderRemoteMacs(ctx, logicalSwitchId)
Get L2 forwarder remote mac addresses

Returns remote mac addresses of the l2 forwarder on logical switch. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalSwitchId** | **string**|  | 

### Return type

[**L2ForwarderRemoteMacs**](L2ForwarderRemoteMacs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2ForwarderStatistics**
> L2ForwarderStatistics GetL2ForwarderStatistics(ctx, logicalSwitchId)
Get L2 forwarder statistics

Returns statistics of the l2 forwarder on logical switch. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalSwitchId** | **string**|  | 

### Return type

[**L2ForwarderStatistics**](L2ForwarderStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2ForwarderStatus**
> L2ForwarderStatus GetL2ForwarderStatus(ctx, logicalSwitchId, optional)
Get L2 forwarder status

Returns status per transport node of the l2 forwarder on logical switch. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalSwitchId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetL2ForwarderStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetL2ForwarderStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**L2ForwarderStatus**](L2ForwarderStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VPNSessionRemoteMacsForLS**
> L2VpnSessionRemoteMacs GetL2VPNSessionRemoteMacsForLS(ctx, sessionId, optional)
Get L2VPN session remote mac for logical switch

Get L2VPN session remote mac for logical switch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetL2VPNSessionRemoteMacsForLSOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetL2VPNSessionRemoteMacsForLSOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logicalSwitchId** | **optional.String**| logical switch identifier | 

### Return type

[**L2VpnSessionRemoteMacs**](L2VPNSessionRemoteMacs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VPNSessionStatistics**
> L2VpnSessionStatistics GetL2VPNSessionStatistics(ctx, sessionId, optional)
Get L2VPN session statistics

Get statistics of a L2VPN session. Query parameter source=realtime is the only supported source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetL2VPNSessionStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetL2VPNSessionStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**L2VpnSessionStatistics**](L2VPNSessionStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VPNSessionStatus**
> L2VpnSessionStatus GetL2VPNSessionStatus(ctx, sessionId, optional)
Get L2VPN session status

Aggregated status of L2VPN session. Query parameter source=realtime|cached is supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetL2VPNSessionStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetL2VPNSessionStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**L2VpnSessionStatus**](L2VPNSessionStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VPNSessionSummary**
> L2VpnSessionSummary GetL2VPNSessionSummary(ctx, optional)
Get status summary of all existing L2VPN sessions.

Load all the existing L2VPN sessions and return the status summary of all L2VPN sessions. Query parameter source supports only cached mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiGetL2VPNSessionSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetL2VPNSessionSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **optional.String**| Data source type. | 

### Return type

[**L2VpnSessionSummary**](L2VPNSessionSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VpnService**
> L2VpnService GetL2VpnService(ctx, l2vpnServiceId)
Get L2VPN service

Get a specific L2VPN service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **l2vpnServiceId** | **string**|  | 

### Return type

[**L2VpnService**](L2VpnService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VpnSession**
> L2VpnSession GetL2VpnSession(ctx, l2vpnSessionId)
Get a L2VPN session

Get a specific L2VPN session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **l2vpnSessionId** | **string**|  | 

### Return type

[**L2VpnSession**](L2VpnSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetL2VpnSessionPeerCodes**
> L2VpnSessionPeerCodes GetL2VpnSessionPeerCodes(ctx, l2vpnSessionId)
Get peer codes for the L2VpnSession

Get peer codes for the L2VPN session to program the remote side of the tunnel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **l2vpnSessionId** | **string**|  | 

### Return type

[**L2VpnSessionPeerCodes**](L2VpnSessionPeerCodes.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerPoolStatistics**
> LbPoolStatistics GetLoadBalancerPoolStatistics(ctx, serviceId, poolId, optional)
Get the statistics of load balancer pool

Returns the statistics of the given load balancer pool by given load balancer serives id and load balancer pool id. Currently, only realtime mode is supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **poolId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLoadBalancerPoolStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLoadBalancerPoolStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**LbPoolStatistics**](LbPoolStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerPoolStatus**
> LbPoolStatus GetLoadBalancerPoolStatus(ctx, serviceId, poolId, optional)
Get the status of load balancer pool

Returns the status of the given load balancer pool by given load balancer serives id and load balancer pool id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **poolId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLoadBalancerPoolStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLoadBalancerPoolStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**LbPoolStatus**](LbPoolStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerServiceStatistics**
> LbServiceStatistics GetLoadBalancerServiceStatistics(ctx, serviceId, optional)
Get the statistics of load balancer service

Returns the statistics of the given load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLoadBalancerServiceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLoadBalancerServiceStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LbServiceStatistics**](LbServiceStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerServiceStatus**
> LbServiceStatus GetLoadBalancerServiceStatus(ctx, serviceId, optional)
Get the status of the given load balancer service

Returns the status of the given load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLoadBalancerServiceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLoadBalancerServiceStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LbServiceStatus**](LbServiceStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerVirtualServerStatistics**
> LbVirtualServerStatistics GetLoadBalancerVirtualServerStatistics(ctx, serviceId, virtualServerId, optional)
Get the statistics of the given load balancer virtual server

Returns the statistics of the load balancer virtual server by given load  balancer serives id and load balancer virtual server id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **virtualServerId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLoadBalancerVirtualServerStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLoadBalancerVirtualServerStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**LbVirtualServerStatistics**](LbVirtualServerStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLoadBalancerVirtualServerStatus**
> LbVirtualServerStatus GetLoadBalancerVirtualServerStatus(ctx, serviceId, virtualServerId, optional)
Get the status of the load balancer virtual server

Returns the status of the virtual server by given load balancer serives id and load balancer virtual server id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **virtualServerId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLoadBalancerVirtualServerStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLoadBalancerVirtualServerStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**LbVirtualServerStatus**](LbVirtualServerStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalPort**
> LogicalPort GetLogicalPort(ctx, lportId)
Get Information About a Logical Port

Returns information about a specified logical port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**|  | 

### Return type

[**LogicalPort**](LogicalPort.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalPortMacTable**
> LogicalPortMacAddressListResult GetLogicalPortMacTable(ctx, lportId, optional)
Get MAC table of a logical port with a given port id (lport-id)

Returns MAC table of a specified logical port. If the target transport node id is not provided, the NSX manager will ask the controller for the transport node where the logical port is located. The query parameter \"source=cached\" is not supported. MAC table retrieval is not supported on logical ports that are attached to a logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalPortMacTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalPortMacTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**LogicalPortMacAddressListResult**](LogicalPortMacAddressListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalPortMacTableInCsvFormatCsv**
> LogicalPortMacAddressCsvListResult GetLogicalPortMacTableInCsvFormatCsv(ctx, lportId, optional)
Get MAC table of a logical port with a given port id (lport-id)

Returns MAC table in CSV format of a specified logical port. If the target transport node id is not provided, the NSX manager will ask the controller for the transport node where the logical port is located. The query parameter \"source=cached\" is not supported. MAC table retrieval is not supported on logical ports that are attached to a logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalPortMacTableInCsvFormatCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalPortMacTableInCsvFormatCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**LogicalPortMacAddressCsvListResult**](LogicalPortMacAddressCsvListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalPortOperationalStatus**
> LogicalPortOperationalStatus GetLogicalPortOperationalStatus(ctx, lportId, optional)
Get Operational Status for Logical Port of a Given Port ID (lport-id)

Returns operational status of a specified logical port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalPortOperationalStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalPortOperationalStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalPortOperationalStatus**](LogicalPortOperationalStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalPortState**
> LogicalPortState GetLogicalPortState(ctx, lportId)
Get realized state & location of a logical port

Returns transport node id for a specified logical port. Also returns information about all address bindings of the specified logical port. This includes address bindings discovered via various snooping methods like ARP snooping, DHCP snooping etc. and addressing bindings that are realized based on user configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**|  | 

### Return type

[**LogicalPortState**](LogicalPortState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalPortStatistics**
> LogicalPortStatistics GetLogicalPortStatistics(ctx, lportId, optional)
Get Statistics for Logical Port of a Given Port ID (lport-id)

Returns statistics of a specified logical port. If the logical port is attached to a logical router port, query parameter \"source=realtime\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lportId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalPortStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalPortStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalPortStatistics**](LogicalPortStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalPortStatusSummary**
> LogicalPortStatusSummary GetLogicalPortStatusSummary(ctx, optional)
Get Operational Status Summary of All Logical Ports in the System

Returns operational status of all logical ports. The query parameter \"source=realtime\" is not supported. Pagination is not supported for this API. The query parameters \"cursor\", \"sort_ascending\", \"sort_by\", \"page_size\" and \"included_fields\" will be ignored. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiGetLogicalPortStatusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalPortStatusSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentId** | **optional.String**| Logical Port attachment Id | 
 **attachmentType** | **optional.String**| Type of attachment for logical port; for query only. | 
 **bridgeClusterId** | **optional.String**| Bridge Cluster identifier | 
 **containerPortsOnly** | **optional.Bool**| Only container VIF logical ports will be returned if true | [default to false]
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **diagnostic** | **optional.Bool**| Flag to enable showing of transit logical port. | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **logicalSwitchId** | **optional.String**| Logical Switch identifier | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **parentVifId** | **optional.String**| ID of the VIF of type PARENT | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **switchingProfileId** | **optional.String**| Network Profile identifier | 
 **transportNodeId** | **optional.String**| Transport node identifier | 
 **transportZoneId** | **optional.String**| Transport zone identifier | 

### Return type

[**LogicalPortStatusSummary**](LogicalPortStatusSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterForwardingTable**
> LogicalRouterRouteTable GetLogicalRouterForwardingTable(ctx, logicalRouterId, transportNodeId, optional)
Get FIB table on a specified node for a logical router

Returns the FIB for the logical router on a node of the given transport-node-id. Query parameter \"transport_node_id=<transport-node-id>\" is required. To filter the result by network address, paramter \"network_prefix=<a.b.c.d/mask>\" needs to be specified. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **transportNodeId** | **string**| TransportNode Id | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterForwardingTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterForwardingTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **networkPrefix** | **optional.String**| IPv4 or IPv6 CIDR Block | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalRouterRouteTable**](LogicalRouterRouteTable.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterForwardingTableInCsvFormatCsv**
> LogicalRouterRouteTableInCsvFormat GetLogicalRouterForwardingTableInCsvFormatCsv(ctx, logicalRouterId, transportNodeId, optional)
Get FIB table on a specified node for a logical router

Returns the FIB table in CSV format for the logical router on a node of the given transport-node-id. Query parameter \"transport_node_id=<transport-node-id>\" is required. To filter the result by network address, paramter \"network_prefix=<a.b.c.d/mask>\" needs to be specified. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **transportNodeId** | **string**| TransportNode Id | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterForwardingTableInCsvFormatCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterForwardingTableInCsvFormatCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkPrefix** | **optional.String**| IPv4 or IPv6 CIDR Block | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalRouterRouteTableInCsvFormat**](LogicalRouterRouteTableInCsvFormat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterPortArpTable**
> LogicalRouterPortArpTable GetLogicalRouterPortArpTable(ctx, logicalRouterPortId, optional)
Get the ARP table (IPv4) or Neighbor Discovery table (IPv6) for the Logical Router Port of the given id 

Returns ARP table (IPv4) or Neighbor Discovery table (IPv6) for the Logical Router Port of the given id, on a node if a query parameter \"transport_node_id=<transport-node-id>\" is given. The transport_node_id parameter is mandatory if the router port is not uplink type. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterPortId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterPortArpTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterPortArpTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**LogicalRouterPortArpTable**](LogicalRouterPortArpTable.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterPortArpTableInCsvFormatCsv**
> LogicalRouterPortArpTableInCsvFormat GetLogicalRouterPortArpTableInCsvFormatCsv(ctx, logicalRouterPortId, optional)
Get the ARP table (IPv4) or Neighbor Discovery table (IPv6) for the Logical Router Port of the given id 

Returns ARP table (IPv4) or Neighbor Discovery table (IPv6) in CSV format for the Logical Router Port of the given id, on a node if a query parameter \"transport_node_id=<transport-node-id>\" is given. The transport_node_id parameter is mandatory if the router port is not uplink type. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterPortId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterPortArpTableInCsvFormatCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterPortArpTableInCsvFormatCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**LogicalRouterPortArpTableInCsvFormat**](LogicalRouterPortArpTableInCsvFormat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterPortState**
> LogicalRouterPortState GetLogicalRouterPortState(ctx, logicalRouterPortId, optional)
Get the Realized State of a Logical Router Port

Return realized state information of a logical router port. Any configuration update that affects the logical router port can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of logical router ports, dhcp relays, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterPortId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterPortStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterPortStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**LogicalRouterPortState**](LogicalRouterPortState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterPortStatistics**
> LogicalRouterPortStatistics GetLogicalRouterPortStatistics(ctx, logicalRouterPortId, optional)
Get the statistics of a specified logical router port on all or a specified node

Returns the statistics for the Logical Router Port. If query parameter \"transport_node_id=<transport-node-id>\" is given,  only the statistics from the given node for the logical router port will be returned. Otherwise the statistics from each node for the same logical router port will be returned. The transport_node_id is mandatory if the router port is not uplink type. The query parameter \"source=cached\" will be ignored and it will always return realtime statistics of the logical router port. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterPortId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterPortStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterPortStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**LogicalRouterPortStatistics**](LogicalRouterPortStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterPortStatisticsSummary**
> LogicalRouterPortStatisticsSummary GetLogicalRouterPortStatisticsSummary(ctx, logicalRouterPortId, optional)
Get the statistics summary of a specified logical router port

Returns the summation of statistics from all nodes for the Specified Logical Router Port. The query parameter \"source=realtime\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterPortId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterPortStatisticsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterPortStatisticsSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalRouterPortStatisticsSummary**](LogicalRouterPortStatisticsSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterRouteTable**
> LogicalRouterRouteTable GetLogicalRouterRouteTable(ctx, logicalRouterId, transportNodeId, optional)
Get route table on a given node for a logical router

Deprecated - Please use /logical-routers/<logical-router-id>/routing/routing-table for RIB and /logical-routers/<logical-router-id>/routing/forwarding-table for FIB. Returns the route table for the logical router on a node of the given transport-node-id. Query parameter \"transport_node_id=<transport-node-id>\" is required. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **transportNodeId** | **string**| TransportNode Id | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterRouteTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterRouteTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalRouterRouteTable**](LogicalRouterRouteTable.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterRouteTableInCsvFormatCsv**
> LogicalRouterRouteTableInCsvFormat GetLogicalRouterRouteTableInCsvFormatCsv(ctx, logicalRouterId, transportNodeId, optional)
Get route table on a node for a logical router

Deprecated - Please use /logical-routers/<logical-router-id>/routing/routing-table for RIB and /logical-routers/<logical-router-id>/routing/forwarding-table for FIB. Returns the route table in CSV format for the logical router on a node of the given transport-node-id. Query parameter \"transport_node_id=<transport-node-id>\" is required. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **transportNodeId** | **string**| TransportNode Id | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterRouteTableInCsvFormatCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterRouteTableInCsvFormatCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalRouterRouteTableInCsvFormat**](LogicalRouterRouteTableInCsvFormat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterRoutingTable**
> LogicalRouterRouteTable GetLogicalRouterRoutingTable(ctx, logicalRouterId, transportNodeId, optional)
Get RIB table on a specified node for a logical router

Returns the route table(RIB) for the logical router on a node of the given transport-node-id. Query parameter \"transport_node_id=<transport-node-id>\" is required. To filter the result by network address, parameter \"network_prefix=<a.b.c.d/mask>\" needs to be specified. To filter the result by route source, parameter \"route_source=<source_type>\" needs to be specified where source_type can be BGP, STATIC, CONNECTED, NSX_STATIC, TIER1_NAT or TIER0_NAT. It is also possible to filter the RIB table using both network address and route source filter together. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **transportNodeId** | **string**| TransportNode Id | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterRoutingTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterRoutingTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **networkPrefix** | **optional.String**| IPv4 or IPv6 CIDR Block | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **routeSource** | **optional.String**| Route source filter parameter | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **vrfTable** | **optional.String**| VRF filter parameter | 

### Return type

[**LogicalRouterRouteTable**](LogicalRouterRouteTable.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterRoutingTableInCsvFormatCsv**
> LogicalRouterRouteTableInCsvFormat GetLogicalRouterRoutingTableInCsvFormatCsv(ctx, logicalRouterId, transportNodeId, optional)
Get RIB table on a specified node for a logical router

Returns the route table in CSV format for the logical router on a node of the given transport-node-id. Query parameter \"transport_node_id=<transport-node-id>\" is required. To filter the result by network address, paramter \"network_prefix=<a.b.c.d/mask>\" needs to be specified. To filter the result by route source, parameter \"route_source=<source_type>\" needs to be specified where source_type can be BGP, STATIC, CONNECTED, NSX_STATIC, TIER1_NAT or TIER0_NAT. It is also possible to filter the RIB table using both network address and route source filter together. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **transportNodeId** | **string**| TransportNode Id | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterRoutingTableInCsvFormatCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterRoutingTableInCsvFormatCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkPrefix** | **optional.String**| IPv4 or IPv6 CIDR Block | 
 **routeSource** | **optional.String**| Route source filter parameter | 
 **source** | **optional.String**| Data source type. | 
 **vrfTable** | **optional.String**| VRF filter parameter | 

### Return type

[**LogicalRouterRouteTableInCsvFormat**](LogicalRouterRouteTableInCsvFormat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterState**
> LogicalRouterState GetLogicalRouterState(ctx, logicalRouterId, optional)
Get the Realized State of a Logical Router

Return realized state information of a logical router. Any configuration update that affects the logical router can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of logical router, static routes, etc. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**LogicalRouterState**](LogicalRouterState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalRouterStatus**
> LogicalRouterStatus GetLogicalRouterStatus(ctx, logicalRouterId, optional)
Get the status for the Logical Router of the given id

Returns status for the Logical Router of the given id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalRouterStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalRouterStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalRouterStatus**](LogicalRouterStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalServiceRouterClusterState**
> LogicalServiceRouterClusterState GetLogicalServiceRouterClusterState(ctx, logicalRouterId, optional)
Get the Realized State of a Logical Service Router Cluster

Return realized state information of a logical service router cluster. Any configuration update that affects the logical service router cluster can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of nat, bgp, bfd, etc.  What is a Service Router? When a service cannot be distributed is enabled on a Logical Router, a Service Router (SR) is instantiated. Some examples of services that are not distributed are NAT, DHCP server, Metadata Proxy, Edge Firewall, Load Balancer and so on. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalServiceRouterClusterStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalServiceRouterClusterStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**LogicalServiceRouterClusterState**](LogicalServiceRouterClusterState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitch**
> LogicalSwitch GetLogicalSwitch(ctx, lswitchId)
Get Logical Switch associated with the provided id (lswitch-id)

Returns information about the specified logical switch Id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 

### Return type

[**LogicalSwitch**](LogicalSwitch.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchMacTable**
> MacAddressListResult GetLogicalSwitchMacTable(ctx, lswitchId, optional)
Get MAC Table for Logical Switch of the Given ID (lswitch-id)

Returns MAC table of a specified logical switch from the given transport node if a transport node id is given in the query parameter from the Central Controller Plane. The query parameter \"source=cached\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalSwitchMacTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalSwitchMacTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**MacAddressListResult**](MacAddressListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchMacTableInCsvFormatCsv**
> MacAddressCsvListResult GetLogicalSwitchMacTableInCsvFormatCsv(ctx, lswitchId, optional)
Get MAC Table for Logical Switch of the Given ID (lswitch-id)

Returns MAC table of a specified logical switch in CSV format from the given transport node if a transport node id is given in the query parameter from the Central Controller Plane. The query parameter \"source=cached\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalSwitchMacTableInCsvFormatCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalSwitchMacTableInCsvFormatCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**MacAddressCsvListResult**](MacAddressCsvListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchState**
> LogicalSwitchState GetLogicalSwitchState(ctx, lswitchId)
Get the realized state associated with provided logical switch id

Returns current state of the logical switch configuration and details of only out-of-sync transport nodes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 

### Return type

[**LogicalSwitchState**](LogicalSwitchState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchStatistics**
> LogicalSwitchStatistics GetLogicalSwitchStatistics(ctx, lswitchId, optional)
Get Statistics for Logical Switch of the Given ID (lswitch-id)

Returns statistics  of a specified logical switch. The query parameter \"source=realtime\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalSwitchStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalSwitchStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LogicalSwitchStatistics**](LogicalSwitchStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchStatus**
> LogicalSwitchStatus GetLogicalSwitchStatus(ctx, lswitchId)
Get Logical Switch runtime status info for a given logical switch

Returns the number of ports assigned to a logical switch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 

### Return type

[**LogicalSwitchStatus**](LogicalSwitchStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchStatusSummary**
> LogicalSwitchStatusSummary GetLogicalSwitchStatusSummary(ctx, optional)
Get Status Summary of All Logical Switches in the System

Returns Operational status of all logical switches. The query parameter \"source=realtime\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiGetLogicalSwitchStatusSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalSwitchStatusSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **diagnostic** | **optional.Bool**| Flag to enable showing of transit logical switch. | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **switchingProfileId** | **optional.String**| Switching Profile identifier | 
 **transportType** | **optional.String**| Mode of transport supported in the transport zone for this logical switch | 
 **transportZoneId** | **optional.String**| Transport zone identifier | 
 **uplinkTeamingPolicyName** | **optional.String**| The logical switch&#x27;s uplink teaming policy name | 
 **vlan** | **optional.Int64**| Virtual Local Area Network Identifier | 
 **vni** | **optional.Int32**| VNI of the OVERLAY LogicalSwitch(es) to return. | 

### Return type

[**LogicalSwitchStatusSummary**](LogicalSwitchStatusSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchVtepTable**
> VtepListResult GetLogicalSwitchVtepTable(ctx, lswitchId, optional)
Get virtual tunnel endpoint table for logical switch of the given ID (lswitch-id) 

Returns the virtual tunnel endpoint table of a specified logical switch from the given transport node if a transport node id is given in the query parameter, from the Central Controller Plane. The query parameter \"source=cached\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalSwitchVtepTableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalSwitchVtepTableOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**VtepListResult**](VtepListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogicalSwitchVtepTableInCsvFormatCsv**
> VtepCsvListResult GetLogicalSwitchVtepTableInCsvFormatCsv(ctx, lswitchId, optional)
Get virtual tunnel endpoint table for logical switch of the given ID (lswitch-id) 

Returns virtual tunnel endpoint table of a specified logical switch in CSV format from the given transport node if a transport node id is given in the query parameter from the Central Controller Plane. The query parameter \"source=cached\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **lswitchId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetLogicalSwitchVtepTableInCsvFormatCsvOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetLogicalSwitchVtepTableInCsvFormatCsvOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 
 **transportNodeId** | **optional.String**| TransportNode Id | 

### Return type

[**VtepCsvListResult**](VtepCsvListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMACAddresses**
> MacAddressElementListResult GetMACAddresses(ctx, macSetId)
Get all MACAddresses in a MACSet

List all MAC addresses in a MACSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macSetId** | **string**| MAC Set Id | 

### Return type

[**MacAddressElementListResult**](MACAddressElementListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetManualHealthCheck**
> ManualHealthCheck GetManualHealthCheck(ctx, manualHealthCheckId)
Get an existing manual health check

Get an existing manual health check by health check ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **manualHealthCheckId** | **string**|  | 

### Return type

[**ManualHealthCheck**](ManualHealthCheck.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMasterSwitchSetting**
> IdfwMasterSwitchSetting GetMasterSwitchSetting(ctx, )
Get Identity Firewall master switch enabled/disabled

Fetches IDFW master switch setting to check whether master switch is enabled or disabled 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdfwMasterSwitchSetting**](IdfwMasterSwitchSetting.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMemberTypes**
> EffectiveMemberTypeListResult GetMemberTypes(ctx, nsGroupId, optional)
Get member types from NSGroup

Returns member types for a specified NSGroup including child NSGroups. This considers static members and members added via membership criteria only 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiGetMemberTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetMemberTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EffectiveMemberTypeListResult**](EffectiveMemberTypeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadataProxyStatistics**
> MetadataProxyStatistics GetMetadataProxyStatistics(ctx, proxyId, optional)
Get Metadata Proxy statistics with given proxy id

Returns the statistics of the given metatada proxy. If no logical switch is provided, all staticstics of all the logical switches the proxy was attached will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetMetadataProxyStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetMetadataProxyStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logicalSwitchId** | **optional.String**| The uuid of logical switch | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**MetadataProxyStatistics**](MetadataProxyStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadataProxyStatus**
> MetadataProxyStatus GetMetadataProxyStatus(ctx, proxyId, logicalSwitchId)
Get Metadata Proxy status with given proxy id and attached logical switch.

Returns the status of the given metadata proxy and attached logical switch. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyId** | **string**|  | 
  **logicalSwitchId** | **string**|  | 

### Return type

[**MetadataProxyStatus**](MetadataProxyStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatRule**
> NatRule GetNatRule(ctx, logicalRouterId, ruleId)
Get a specific NAT rule from a given logical router

Get a specific NAT rule from a given logical router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**NatRule**](NatRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatStatisticsPerLogicalRouter**
> NatStatisticsPerLogicalRouter GetNatStatisticsPerLogicalRouter(ctx, logicalRouterId, optional)
Get the statistics of all rules of the logical router

Returns the summation of statistics for all rules from all nodes for the Specified Logical Router. Also gives the per transport node statistics for provided logical router. The query parameter \"source=realtime\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetNatStatisticsPerLogicalRouterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetNatStatisticsPerLogicalRouterOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**NatStatisticsPerLogicalRouter**](NatStatisticsPerLogicalRouter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatStatisticsPerRule**
> NatStatisticsPerRule GetNatStatisticsPerRule(ctx, logicalRouterId, ruleId, optional)
Get the statistics of a specified logical router NAT Rule

Returns the summation of statistics from all nodes for the Specified Logical Router NAT Rule. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetNatStatisticsPerRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetNatStatisticsPerRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**NatStatisticsPerRule**](NatStatisticsPerRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNatStatisticsPerTransportNode**
> NatStatisticsPerTransportNode GetNatStatisticsPerTransportNode(ctx, nodeId, optional)
Get statistics for all logical router NAT rules on a transport node

Returns the summation of statistics for all rules from all logical routers which are present on given transport node. Only cached statistics are supported. The query parameter \"source=realtime\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetNatStatisticsPerTransportNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetNatStatisticsPerTransportNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**NatStatisticsPerTransportNode**](NatStatisticsPerTransportNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNormalizations**
> NormalizedResourceListResult GetNormalizations(ctx, preferredNormalizationType, resourceId, resourceType, optional)
Get normalizations based on the query parameters

Returns the list of normalized resources based on the query parameters. Id and Type of the resource on which the normalizations is to be performed, are to be specified as query parameters in the URI. The target resource types to which normalization is to be done should also be specified as query parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **preferredNormalizationType** | **string**| Resource type valid for use as target in normalization API. | 
  **resourceId** | **string**| Identifier of the resource on which normalization is to be performed | 
  **resourceType** | **string**| Resource type valid for use as source in normalization API. | 
 **optional** | ***ManagementPlaneApiApiGetNormalizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetNormalizationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NormalizedResourceListResult**](NormalizedResourceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNsgroupVmDetails**
> IdfwNsgroupVmDetailListResult GetNsgroupVmDetails(ctx, groupId)
Get all IDFW NSGroup VM details for a given NSGroup

Get all Identity Firewall NSGroup VM details for a given NSGroup.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**|  | 

### Return type

[**IdfwNsgroupVmDetailListResult**](IdfwNsgroupVmDetailListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPBRRule**
> PbrRule GetPBRRule(ctx, sectionId, ruleId)
Read an Existing Rule

Return existing PBR rule information in a PBR section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**PbrRule**](PBRRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPBRRuleStats**
> PbrStats GetPBRRuleStats(ctx, sectionId, ruleId)
Get PBR rule level statistics.

Get aggregated statistics for a rule for given PBR rule. Stats include total number of packets and total number of bytes for the PBR rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**PbrStats**](PBRStats.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPBRRules**
> PbrRuleListResult GetPBRRules(ctx, sectionId, optional)
Get All the Rules for a Section

Return all PBR rule(s) information for a given PBR section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetPBRRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetPBRRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appliedTos** | **optional.String**| AppliedTo&#x27;s referenced by this section or section&#x27;s Distributed Service Rules . | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **destinations** | **optional.String**| Destinations referenced by this section&#x27;s Distributed Service Rules . | 
 **filterType** | **optional.String**| Filter type | [default to FILTER]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **services** | **optional.String**| NSService referenced by this section&#x27;s Distributed Service Rules . | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sources** | **optional.String**| Sources referenced by this section&#x27;s Distributed Service Rules . | 

### Return type

[**PbrRuleListResult**](PBRRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPBRSection**
> PbrSection GetPBRSection(ctx, sectionId)
Get an Existing Section

Returns information about PBR section for the identifier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

[**PbrSection**](PBRSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPBRSectionStats**
> PbrStatsList GetPBRSectionStats(ctx, sectionId)
Get PBR section level statistics.

Get aggregated statistics for all rules for a given pbr section. Data includes total number of packets, and total number of bytes for all PBR rules in the given section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

[**PbrStatsList**](PBRStatsList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPBRSectionWithRulesListWithRules**
> PbrSectionRuleList GetPBRSectionWithRulesListWithRules(ctx, sectionId)
Get an Existing Section, Including Rules

Returns PBR section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number rule references is not supported.  Instead, to read PBR rules, use: GET /api/v1/pbr/sections/&lt;section-id&gt;/rules with the appropriate page_size. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

[**PbrSectionRuleList**](PBRSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPeerConfig**
> string GetPeerConfig(ctx, ipsecVpnSessionId)
Get VPN configuration for the peer site

API to download VPN configuration for the peer site. The configuration contains pre-shared key and secret; be careful when sharing or storing it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipsecVpnSessionId** | **string**|  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPortMirroringSession**
> PortMirroringSession GetPortMirroringSession(ctx, mirrorSessionId)
Get the mirror session

Get the mirror session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mirrorSessionId** | **string**|  | 

### Return type

[**PortMirroringSession**](PortMirroringSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRule**
> FirewallRule GetRule(ctx, sectionId, ruleId)
Read an Existing Rule

Return existing firewall rule information in a firewall section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuleState**
> RuleState GetRuleState(ctx, ruleId, optional)
Get the Realized State of a Firewall Rule

Return realized state information of a firewall rule. Returned response is same as rule's section realization state response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetRuleStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetRuleStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**RuleState**](RuleState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRules**
> FirewallRuleListResult GetRules(ctx, sectionId, optional)
Get All the Rules for a Section

Return all firewall rule(s) information for a given firewall section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appliedTos** | **optional.String**| AppliedTo&#x27;s referenced by this section or section&#x27;s Distributed Service Rules . | 
 **contextProfiles** | **optional.String**| Limits results to sections having rules with specific Context Profiles. | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **deepSearch** | **optional.Bool**| Toggle to search with direct or indirect references. | [default to false]
 **destinations** | **optional.String**| Destinations referenced by this section&#x27;s Distributed Service Rules . | 
 **extendedSources** | **optional.String**| Limits results to sections having rules with specific Extended Sources. | 
 **filterType** | **optional.String**| Filter type | [default to FILTER]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **searchInvalidReferences** | **optional.Bool**| Return invalid references in results. | [default to false]
 **services** | **optional.String**| NSService referenced by this section&#x27;s Distributed Service Rules . | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sources** | **optional.String**| Sources referenced by this section&#x27;s Distributed Service Rules . | 

### Return type

[**FirewallRuleListResult**](FirewallRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeInterfaceOperationalStatus**
> RuntimeInterfaceOperationalStatus GetRuntimeInterfaceOperationalStatus(ctx, serviceId, serviceInstanceId, instanceRuntimeId, interfaceIndex, optional)
Get operational status for an interface

Returns operational status of a specified interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 
  **instanceRuntimeId** | **string**|  | 
  **interfaceIndex** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetRuntimeInterfaceOperationalStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetRuntimeInterfaceOperationalStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **source** | **optional.String**| Data source type. | 

### Return type

[**RuntimeInterfaceOperationalStatus**](RuntimeInterfaceOperationalStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRuntimeInterfaceStatistics**
> RuntimeInterfaceStatistics GetRuntimeInterfaceStatistics(ctx, serviceId, serviceInstanceId, instanceRuntimeId, interfaceIndex, optional)
Get statistics for a given interface identified by the interface index

Returns statistics of a specified interface via associated logical port. If the logical port is attached to a logical router port, query parameter \"source=realtime\" is not supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 
  **instanceRuntimeId** | **string**|  | 
  **interfaceIndex** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetRuntimeInterfaceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetRuntimeInterfaceStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **source** | **optional.String**| Data source type. | 

### Return type

[**RuntimeInterfaceStatistics**](RuntimeInterfaceStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSIServiceProfile**
> BaseServiceProfile GetSIServiceProfile(ctx, serviceId, serviceProfileId)
Get a particular ServiceProfile for a Service.

Returns detailed service profile information for a given Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceProfileId** | **string**|  | 

### Return type

[**BaseServiceProfile**](BaseServiceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSection**
> FirewallSection GetSection(ctx, sectionId)
Get an Existing Section

Returns information about firewall section for the identifier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

[**FirewallSection**](FirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSectionState**
> FirewallSectionState GetSectionState(ctx, sectionId, optional)
Get the Realized State of a Firewall Section

Return realized state information of a firewall section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetSectionStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetSectionStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**FirewallSectionState**](FirewallSectionState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSectionWithRulesListWithRules**
> FirewallSectionRuleList GetSectionWithRulesListWithRules(ctx, sectionId)
Get an Existing Section, Including Rules

Returns firewall section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules with a large number rule references is not supported.  Instead, to read firewall rules, use: GET /api/v1/firewall/sections/&lt;section-id&gt;/rules with the appropriate page_size. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

[**FirewallSectionRuleList**](FirewallSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSectionsSummary**
> FirewallSectionsSummaryList GetSectionsSummary(ctx, optional)
Get the summary of sections in the firewall configuration.

List the summary of number of sections and number of rules for each firewall category (L2DFW, L3DFW). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiGetSectionsSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetSectionsSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **optional.String**| Data source type. | 

### Return type

[**FirewallSectionsSummaryList**](FirewallSectionsSummaryList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceAssociations**
> ServiceAssociationListResult GetServiceAssociations(ctx, nsgroupId, serviceType, optional)
Get services to which the given nsgroup belongs to 

Returns information about services that are associated with the given NSGroup. The service name is passed by service_type parameter 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsgroupId** | **string**|  | 
  **serviceType** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetServiceAssociationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetServiceAssociationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **fetchParentgroupAssociations** | **optional.Bool**| Fetch complete list of associated resources considering nesting  | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceAssociationListResult**](ServiceAssociationListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceAttachment**
> ServiceAttachment GetServiceAttachment(ctx, serviceAttachmentId)
Get a particular service attachment.

Returns detailed Attachment information for a given service attachment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceAttachmentId** | **string**|  | 

### Return type

[**ServiceAttachment**](ServiceAttachment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceChain**
> ServiceChain GetServiceChain(ctx, serviceChainId)
Get a particular service chain.

Returns detailed service chain information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceChainId** | **string**|  | 

### Return type

[**ServiceChain**](ServiceChain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceDeployment**
> ServiceDeployment GetServiceDeployment(ctx, serviceId, serviceDeploymentId)
Get a particular service deployment.

Returns detail of service deployment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceDeploymentId** | **string**|  | 

### Return type

[**ServiceDeployment**](ServiceDeployment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceDeploymentState**
> ConfigurationState GetServiceDeploymentState(ctx, serviceId, serviceDeploymentId)
Get Service-Deployment state for Service.

Returns configuration state of deployed partner service using service insertion framework. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceDeploymentId** | **string**|  | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceDeploymentStatus**
> ServiceDeploymentStatus GetServiceDeploymentStatus(ctx, serviceId, serviceDeploymentId, optional)
Get a particular service deployment status.

Returns current status of the deployment of partner service. Available only for EPP Services. By default this API would return cached status. Caching happens every 3 minutes. For realtime status, query parameter \"source=realtime\" needs to be passed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceDeploymentId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetServiceDeploymentStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetServiceDeploymentStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**ServiceDeploymentStatus**](ServiceDeploymentStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceDeployments**
> ServiceDeploymentListResult GetServiceDeployments(ctx, serviceId)
Get all service deployments for the given service id

Returns the list of deployments for the given service 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**ServiceDeploymentListResult**](ServiceDeploymentListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInsertionExcludeList**
> SiExcludeList GetServiceInsertionExcludeList(ctx, )
Get list of members in exclude list

Get list of members in exclude list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SiExcludeList**](SIExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInsertionRule**
> ServiceInsertionRule GetServiceInsertionRule(ctx, sectionId, ruleId)
Read an Existing Rule

Return existing serviceinsertion rule information in a serviceinsertion section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**ServiceInsertionRule**](ServiceInsertionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInsertionRules**
> ServiceInsertionRuleListResult GetServiceInsertionRules(ctx, sectionId, optional)
Get All the Rules for a Section

Return all serviceinsertion rule(s) information for a given serviceinsertion section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetServiceInsertionRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetServiceInsertionRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appliedTos** | **optional.String**| AppliedTo&#x27;s referenced by this section or section&#x27;s Distributed Service Rules . | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **destinations** | **optional.String**| Destinations referenced by this section&#x27;s Distributed Service Rules . | 
 **filterType** | **optional.String**| Filter type | [default to FILTER]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **services** | **optional.String**| NSService referenced by this section&#x27;s Distributed Service Rules . | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sources** | **optional.String**| Sources referenced by this section&#x27;s Distributed Service Rules . | 

### Return type

[**ServiceInsertionRuleListResult**](ServiceInsertionRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInsertionSection**
> ServiceInsertionSection GetServiceInsertionSection(ctx, sectionId)
Get an Existing Section

Returns information about serviceinsertion section for the identifier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

[**ServiceInsertionSection**](ServiceInsertionSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInsertionSectionWithRulesListWithRules**
> ServiceInsertionSectionRuleList GetServiceInsertionSectionWithRulesListWithRules(ctx, sectionId)
Get an Existing Section, Including Rules

Returns serviceinsertion section information with rules for a section identifier. When invoked on a section with a large number of rules, this API is supported only at low rates of invocation (not more than 4-5 times per minute). The typical latency of this API with about 1024 rules is about 4-5 seconds. This API should not be invoked with large payloads at automation speeds. More than 50 rules are not supported.  Instead, to read serviceinsertion rules, use: GET /api/v1/serviceinsertion/sections/&lt;section-id&gt;/rules with the appropriate page_size. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sectionId** | **string**|  | 

### Return type

[**ServiceInsertionSectionRuleList**](ServiceInsertionSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInsertionService**
> ServiceDefinition GetServiceInsertionService(ctx, serviceId)
Get an existing Service

Returns information about Service-Insertion Service with the given identifier. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInsertionStatus**
> ServiceInsertionStatus GetServiceInsertionStatus(ctx, contextType)
Get ServiceInsertion global status for a context

Get ServiceInsertion global status for a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **contextType** | **string**|  | 

### Return type

[**ServiceInsertionStatus**](ServiceInsertionStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstance**
> BaseServiceInstance GetServiceInstance(ctx, serviceId, serviceInstanceId)
Get Service-Instance for Service.

Returns Service-Instance information for a given Service-Insertion Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

[**BaseServiceInstance**](BaseServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstanceNSGroups**
> ServiceInstanceNsGroups GetServiceInstanceNSGroups(ctx, serviceId, serviceInstanceId)
Get NSgroups for a given ServiceInstance.

Returns list of NSGroups used in Service Insertion North-South rules for a given Service Instance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

[**ServiceInstanceNsGroups**](ServiceInstanceNSGroups.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstanceState**
> ConfigurationState GetServiceInstanceState(ctx, serviceId, serviceInstanceId)
Get Service-Instance state for Service.

Returns configuration state of one instance of a deployed partner service using service insertion framework. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceInstanceStatus**
> ServiceInstanceStatus GetServiceInstanceStatus(ctx, serviceId, serviceInstanceId, optional)
Get Service-Instance status for Service.

Returns status of one instance of a deployed partner service using service insertion framework. By default this API would return cached status. Caching happens every 3 minutes. For realtime status, query parameter \"source=realtime\" needs to be passed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetServiceInstanceStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetServiceInstanceStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**ServiceInstanceStatus**](ServiceInstanceStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceManager**
> ServiceManager GetServiceManager(ctx, serviceManagerId)
Get service manager

Retrieve service-manager details like name, username, password, vendor ID, thumbprint for a given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceManagerId** | **string**|  | 

### Return type

[**ServiceManager**](ServiceManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceProfileNSGroups**
> ServiceProfileNsGroups GetServiceProfileNSGroups(ctx, serviceId, serviceProfileId)
Get NSgroups for a given ServiceProfile.

Returns list of NSGroups used in Service Insertion rules for a given Service Profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceProfileId** | **string**|  | 

### Return type

[**ServiceProfileNsGroups**](ServiceProfileNSGroups.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSolutionConfig**
> SolutionConfig GetSolutionConfig(ctx, serviceId, solutionConfigId)
Get Solution Config Information for a given solution config id.

Returns Solution Config information for a given solution config id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **solutionConfigId** | **string**|  | 

### Return type

[**SolutionConfig**](SolutionConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStandaloneHostsSwitchSetting**
> IdfwStandaloneHostsSwitchSetting GetStandaloneHostsSwitchSetting(ctx, )
Get Standalone hosts switch enabled/disabled

Fetches IDFW standalone hosts switch setting to check whether standalone hosts is enabled or disabled 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdfwStandaloneHostsSwitchSetting**](IdfwStandaloneHostsSwitchSetting.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSwitchIpfixConfig**
> IpfixObsPointConfig GetSwitchIpfixConfig(ctx, )
Read global switch IPFIX export configuration

Deprecated - Please use /ipfix-profiles/<ipfix-profile-id> for switch IPFIX profile and /ipfix-collector-profiles/<ipfix-collector-profile-id> for IPFIX collector profile. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IpfixObsPointConfig**](IpfixObsPointConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSwitchingProfile**
> BaseSwitchingProfile GetSwitchingProfile(ctx, switchingProfileId)
Get Switching Profile by ID

Returns information about a specified switching profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **switchingProfileId** | **string**|  | 

### Return type

[**BaseSwitchingProfile**](BaseSwitchingProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSwitchingProfileStatus**
> SwitchingProfileStatus GetSwitchingProfileStatus(ctx, switchingProfileId)
Get Counts of Ports and Switches Using This Switching Profile

Get Counts of Ports and Switches Using This Switching Profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **switchingProfileId** | **string**|  | 

### Return type

[**SwitchingProfileStatus**](SwitchingProfileStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSystemStats**
> IdfwSystemStats GetSystemStats(ctx, )
Get IDFW system statistics data

Get IDFW system statistics data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdfwSystemStats**](IdfwSystemStats.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTraceflow**
> Traceflow GetTraceflow(ctx, traceflowId)
Get the Traceflow round status and result summary

Get the Traceflow round status and result summary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **traceflowId** | **string**|  | 

### Return type

[**Traceflow**](Traceflow.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTraceflowObservations**
> TraceflowObservationListResult GetTraceflowObservations(ctx, traceflowId, optional)
Get observations for the Traceflow round

Get observations for the Traceflow round

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **traceflowId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiGetTraceflowObservationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetTraceflowObservationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentName** | **optional.String**| Observations having the given component name will be listed. | 
 **componentType** | **optional.String**| Observations having the given component type will be listed. | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **resourceType** | **optional.String**| The type of observations that will be listed. | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **transportNodeName** | **optional.String**| Observations having the given transport node name will be listed. | 

### Return type

[**TraceflowObservationListResult**](TraceflowObservationListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnassociatedVirtualMachines**
> UnassociatedVmListResult GetUnassociatedVirtualMachines(ctx, optional)
Get the list of all the virtual machines that are not a part of any existing NSGroup.

Get the list of all the virtual machines that are not a part of any existing NSGroup. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiGetUnassociatedVirtualMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiGetUnassociatedVirtualMachinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **displayName** | **optional.String**| Display Name of the virtual machine | 
 **externalId** | **optional.String**| External id of the virtual machine | 
 **hostId** | **optional.String**| Id of the host where this vif is located | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**UnassociatedVmListResult**](UnassociatedVMListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserStats**
> IdfwUserStats GetUserStats(ctx, userId)
Get IDFW user login events for a given user

Get IDFW user login events for a given user (all active plus up to 5 most recent archived entries). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

[**IdfwUserStats**](IdfwUserStats.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVendorTemplate**
> VendorTemplate GetVendorTemplate(ctx, serviceId, vendorTemplateId)
Get a particular vendor template for a given service.

Returns detailed vendor template information for a given service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **vendorTemplateId** | **string**|  | 

### Return type

[**VendorTemplate**](VendorTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVmStats**
> IdfwVmStats GetVmStats(ctx, vmExtId)
Get IDFW user login events for a given VM

Get IDFW user login events for a given VM (all active plus up to 5 most recent archived entries). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vmExtId** | **string**|  | 

### Return type

[**IdfwVmStats**](IdfwVmStats.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBGPCommunityLists**
> BgpCommunityListListResult ListBGPCommunityLists(ctx, logicalRouterId, optional)
Paginated list of BGP community lists on a logical router

Paginated list of BGP Community Lists on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListBGPCommunityListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListBGPCommunityListsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**BgpCommunityListListResult**](BGPCommunityListListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBgpNeighbors**
> BgpNeighborListResult ListBgpNeighbors(ctx, logicalRouterId, optional)
Paginated list of BGP Neighbors on a Logical Router

Paginated list of BGP Neighbors on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListBgpNeighborsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListBgpNeighborsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**BgpNeighborListResult**](BgpNeighborListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBridgeClusters**
> BridgeClusterListResult ListBridgeClusters(ctx, optional)
List All Bridge Clusters

Returns information about all configured bridge clusters 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListBridgeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListBridgeClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**BridgeClusterListResult**](BridgeClusterListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBridgeEndpointProfiles**
> BridgeEndpointProfileListResult ListBridgeEndpointProfiles(ctx, optional)
List All Bridge Endpoint Profiles

Returns information about all configured bridge endoint profiles 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListBridgeEndpointProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListBridgeEndpointProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **edgeClusterId** | **optional.String**| Edge Cluster Identifier | 
 **failoverMode** | **optional.String**|  | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**BridgeEndpointProfileListResult**](BridgeEndpointProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBridgeEndpoints**
> BridgeEndpointListResult ListBridgeEndpoints(ctx, optional)
List All Bridge Endpoints

Returns information about all configured bridge endoints 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListBridgeEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListBridgeEndpointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bridgeClusterId** | **optional.String**| Bridge Cluster Identifier | 
 **bridgeEndpointProfileId** | **optional.String**| Bridge endpoint profile used by the edge cluster | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **logicalSwitchId** | **optional.String**| Logical Switch Identifier | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **vlanTransportZoneId** | **optional.String**| VLAN transport zone id used by the edge cluster | 

### Return type

[**BridgeEndpointListResult**](BridgeEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListComputeCollectionStatuses**
> IdfwComputeCollectionListResult ListComputeCollectionStatuses(ctx, )
List all IDFW enabled ComputeCollection statuses.

Retrieve all the Compute collection status. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdfwComputeCollectionListResult**](IdfwComputeCollectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDADProfiles**
> DadProfileListResult ListDADProfiles(ctx, optional)
Read All IPV6 DADProfiles

Returns all IPv6 DADProfiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListDADProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDADProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DadProfileListResult**](DADProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpIpPools**
> DhcpIpPoolListResult ListDhcpIpPools(ctx, serverId, optional)
Get a paginated list of a DHCP server's IP pools

List the ip pools of a logical DHCP server with pagination support. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListDhcpIpPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpIpPoolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DhcpIpPoolListResult**](DhcpIpPoolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpProfiles**
> DhcpProfileListResult ListDhcpProfiles(ctx, optional)
Get a paginated list of DHCP server profiles

Get a paginated list of DHCP server profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListDhcpProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DhcpProfileListResult**](DhcpProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpRelayProfiles**
> DhcpRelayProfileListResult ListDhcpRelayProfiles(ctx, optional)
List All DHCP Relay Profiles

Returns information about all dhcp relay profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListDhcpRelayProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpRelayProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DhcpRelayProfileListResult**](DhcpRelayProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpRelays**
> DhcpRelayServiceListResult ListDhcpRelays(ctx, optional)
List all DHCP Relay Services

Returns information about all configured dhcp relay services. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListDhcpRelaysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpRelaysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DhcpRelayServiceListResult**](DhcpRelayServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpServers**
> LogicalDhcpServerListResult ListDhcpServers(ctx, optional)
Get a paginated list of DHCP servers

List logical DHCP servers with pagination support.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListDhcpServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LogicalDhcpServerListResult**](LogicalDhcpServerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpStaticBindings**
> DhcpStaticBindingListResult ListDhcpStaticBindings(ctx, serverId, optional)
Get a paginated list of a DHCP server's static bindings

Return a paginated list of a static bindings of a given logical DHCP server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListDhcpStaticBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpStaticBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DhcpStaticBindingListResult**](DhcpStaticBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpV6IpPools**
> DhcpV6IpPoolListResult ListDhcpV6IpPools(ctx, serverId, optional)
Get a paginated list of a DHCP IPv6 server's IP pools

List the ip pools of a logical DHCP IPv6 server with pagination support. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListDhcpV6IpPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpV6IpPoolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DhcpV6IpPoolListResult**](DhcpV6IpPoolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDhcpV6StaticBindings**
> DhcpV6StaticBindingListResult ListDhcpV6StaticBindings(ctx, serverId, optional)
Get a paginated list of a DHCP IPv6 server's static bindings

Return a paginated list of a static bindings of a given logical DHCP IPv6 server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListDhcpV6StaticBindingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDhcpV6StaticBindingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DhcpV6StaticBindingListResult**](DhcpV6StaticBindingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDnsForwaders**
> DnsForwarderListResult ListDnsForwaders(ctx, optional)
Get a paginated list of DNS forwarders

Get a paginated list of DNS forwarders. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListDnsForwadersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListDnsForwadersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DnsForwarderListResult**](DnsForwarderListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnabledComputeCollections**
> IdfwEnabledComputeCollectionListResult ListEnabledComputeCollections(ctx, optional)
List all Identity firewall compute collections

List all Identity firewall compute collections. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListEnabledComputeCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListEnabledComputeCollectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IdfwEnabledComputeCollectionListResult**](IdfwEnabledComputeCollectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallProfiles**
> FirewallProfileListResult ListFirewallProfiles(ctx, resourceType, optional)
Get firewall profiles available.

List all the firewall profiles available by requested resource_type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **resourceType** | **string**| Profile resource type | 
 **optional** | ***ManagementPlaneApiApiListFirewallProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListFirewallProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**FirewallProfileListResult**](FirewallProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFirewallStatus**
> FirewallStatusListResult ListFirewallStatus(ctx, )
List all firewall status for supported contexts

List all firewall status for supported contexts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FirewallStatusListResult**](FirewallStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPPrefixLists**
> IpPrefixListListResult ListIPPrefixLists(ctx, logicalRouterId, optional)
Paginated List of IPPrefixLists

Paginated List of IPPrefixLists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListIPPrefixListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPPrefixListsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpPrefixListListResult**](IPPrefixListListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSecVPNDPDProfiles**
> IpSecVpndpdProfileListResult ListIPSecVPNDPDProfiles(ctx, optional)
Get IPSec dead peer detection (DPD)  profile list result

Get paginated list of all dead peer detection (DPD) profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSecVPNDPDProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSecVPNDPDProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSecVpndpdProfileListResult**](IPSecVPNDPDProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSecVPNIKEProfiles**
> IpSecVpnikeProfileListResult ListIPSecVPNIKEProfiles(ctx, optional)
List IKE profiles

Get paginated list of all IKE Profiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSecVPNIKEProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSecVPNIKEProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSecVpnikeProfileListResult**](IPSecVPNIKEProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSecVPNLocalEndpoints**
> IpSecVpnLocalEndpointListResult ListIPSecVPNLocalEndpoints(ctx, optional)
Get IPSec local endpoint list result

Get paginated list of all local endpoints.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSecVPNLocalEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSecVPNLocalEndpointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **ipsecVpnServiceId** | **optional.String**| Id of the IPSec VPN service | 
 **logicalRouterId** | **optional.String**| Id of logical router | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSecVpnLocalEndpointListResult**](IPSecVPNLocalEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSecVPNPeerEndpoints**
> IpSecVpnPeerEndpointListResult ListIPSecVPNPeerEndpoints(ctx, optional)
Get IPSecVPNPeerEndpoint List Result

Get paginated list of all peer endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSecVPNPeerEndpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSecVPNPeerEndpointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSecVpnPeerEndpointListResult**](IPSecVPNPeerEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSecVPNServices**
> IpSecVpnServiceListResult ListIPSecVPNServices(ctx, optional)
Get IPSec VPN service list result

Get paginated list of all IPSec VPN services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSecVPNServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSecVPNServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSecVpnServiceListResult**](IPSecVPNServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSecVPNSessions**
> IpSecVpnSessionListResult ListIPSecVPNSessions(ctx, optional)
Get IPSec VPN session list result

Get paginated list of all IPSec VPN sessions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSecVPNSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSecVPNSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **ipsecVpnServiceId** | **optional.String**| Id of the IPSec VPN service | 
 **logicalRouterId** | **optional.String**| Id of logical router | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sessionType** | **optional.String**| Resource types of IPsec VPN session | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSecVpnSessionListResult**](IPSecVPNSessionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSecVPNTunnelProfiles**
> IpSecVpnTunnelProfileListResult ListIPSecVPNTunnelProfiles(ctx, optional)
Get IPSecTunnelProfile List Result

Get paginated list of all IPSecTunnelProfiles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSecVPNTunnelProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSecVPNTunnelProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSecVpnTunnelProfileListResult**](IPSecVPNTunnelProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIPSets**
> IpSetListResult ListIPSets(ctx, optional)
List IPSets

Returns paginated list of IPSets 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIPSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIPSetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpSetListResult**](IPSetListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstanceEndpoints**
> InstanceEndpointListResult ListInstanceEndpoints(ctx, serviceId, serviceInstanceId)
List all InstanceEndpoints of a Service Instance.

List all InstanceEndpoints of a service instance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

[**InstanceEndpointListResult**](InstanceEndpointListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInstanceRuntimes**
> InstanceRuntimeListResult ListInstanceRuntimes(ctx, serviceId, serviceInstanceId)
Returns list of instance runtimes of service VM being deployed

Returns list of instance runtimes of service VMs being deployed for a given service instance id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

[**InstanceRuntimeListResult**](InstanceRuntimeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpfixCollectorConfig**
> IpfixCollectorConfigListResult ListIpfixCollectorConfig(ctx, optional)
List IPFIX collector configurations

List IPFIX collector configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIpfixCollectorConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIpfixCollectorConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpfixCollectorConfigListResult**](IpfixCollectorConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpfixCollectorUpmProfiles**
> IpfixCollectorUpmProfileListResult ListIpfixCollectorUpmProfiles(ctx, optional)
List IPFIX Collector Profies

Query IPFIX collector profiles with list parameters. List result can be filtered by profile type defined by IpfixCollectorUpmProfileType. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIpfixCollectorUpmProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIpfixCollectorUpmProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **profileTypes** | **optional.String**| IPFIX Collector Profile Type List | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpfixCollectorUpmProfileListResult**](IpfixCollectorUpmProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpfixConfig**
> IpfixConfigListResult ListIpfixConfig(ctx, optional)
List IPFIX configuration

List IPFIX configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIpfixConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIpfixConfigOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appliedTo** | **optional.String**| Applied To | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **ipfixConfigType** | **optional.String**| Supported IPFIX Config Types. | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpfixConfigListResult**](IpfixConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpfixUpmProfiles**
> IpfixUpmProfileListResult ListIpfixUpmProfiles(ctx, optional)
List IPFIX Profies

Query IPFIX profiles with list parameters. List result can be filtered by profile type defined by IpfixUpmProfileType. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListIpfixUpmProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListIpfixUpmProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appliedToEntityId** | **optional.String**| ID of Entity Applied with Profile | 
 **appliedToEntityType** | **optional.String**| Supported Entity Types | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **profileTypes** | **optional.String**| IPFIX Profile Type List | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpfixUpmProfileListResult**](IpfixUpmProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListL2VpnServices**
> L2VpnServiceListResult ListL2VpnServices(ctx, optional)
Get all L2VPN services

Get paginated list of all L2VPN services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListL2VpnServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListL2VpnServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**L2VpnServiceListResult**](L2VpnServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListL2VpnSessions**
> L2VpnSessionListResult ListL2VpnSessions(ctx, optional)
Get all L2VPN sessions

Get paginated list of all L2VPN sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListL2VpnSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListL2VpnSessionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **l2vpnServiceId** | **optional.String**| Id of the L2Vpn Service | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**L2VpnSessionListResult**](L2VpnSessionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerApplicationProfiles**
> LbAppProfileListResult ListLoadBalancerApplicationProfiles(ctx, optional)
Retrieve a paginated list of load balancer application profiles

Retrieve a paginated list of load balancer application profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerApplicationProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerApplicationProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| application profile type | 

### Return type

[**LbAppProfileListResult**](LbAppProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerClientSslProfiles**
> LbClientSslProfileListResult ListLoadBalancerClientSslProfiles(ctx, optional)
Retrieve a paginated list of load balancer client-ssl profiles

Retrieve a paginated list of load balancer client-ssl profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerClientSslProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerClientSslProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbClientSslProfileListResult**](LbClientSslProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerMonitors**
> LbMonitorListResult ListLoadBalancerMonitors(ctx, optional)
Retrieve a paginated list of load balancer monitors

Retrieve a paginated list of load balancer monitors. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerMonitorsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerMonitorsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| monitor query type | 

### Return type

[**LbMonitorListResult**](LbMonitorListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerPersistenceProfiles**
> LbPersistenceProfileListResult ListLoadBalancerPersistenceProfiles(ctx, optional)
Retrieve a paginated list of load balancer persistence profiles

Retrieve a paginated list of load balancer persistence profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerPersistenceProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerPersistenceProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **type_** | **optional.String**| persistence profile type | 

### Return type

[**LbPersistenceProfileListResult**](LbPersistenceProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerPoolStatistics**
> LbPoolStatisticsListResult ListLoadBalancerPoolStatistics(ctx, serviceId, optional)
Get the statistics list of load balancer pools

Returns the statistics list of load balancer pools in given load balancer service. Currently, only realtime mode is supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListLoadBalancerPoolStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerPoolStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LbPoolStatisticsListResult**](LbPoolStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerPoolStatuses**
> LbPoolStatusListResult ListLoadBalancerPoolStatuses(ctx, serviceId, optional)
Get the status list of load balancer pools

Returns the status list of load balancer pools in given load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListLoadBalancerPoolStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerPoolStatusesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LbPoolStatusListResult**](LbPoolStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerPools**
> LbPoolListResult ListLoadBalancerPools(ctx, optional)
Retrieve a paginated list of load balancer pools

Retrieve a paginated list of load balancer pools. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerPoolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbPoolListResult**](LbPoolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerRules**
> LbRuleListResult ListLoadBalancerRules(ctx, optional)
Retrieve a paginated list of load balancer rules

Retrieve a paginated list of load balancer rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbRuleListResult**](LbRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerServerSslProfiles**
> LbServerSslProfileListResult ListLoadBalancerServerSslProfiles(ctx, optional)
Retrieve a paginated list of load balancer server-ssl profiles

Retrieve a paginated list of load balancer server-ssl profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerServerSslProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerServerSslProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbServerSslProfileListResult**](LbServerSslProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerServices**
> LbServiceListResult ListLoadBalancerServices(ctx, optional)
Retrieve a paginated list of load balancer services

Retrieve a paginated list of load balancer services. When logical_router_id is specified in request parameters, the associated load balancer services which are related to the given logical router returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **logicalRouterId** | **optional.String**| Logical router identifier | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbServiceListResult**](LbServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerSslCiphersAndProtocols**
> LbSslCipherAndProtocolListResult ListLoadBalancerSslCiphersAndProtocols(ctx, optional)
Retrieve a list of supported SSL ciphers and protocols

Retrieve a list of supported SSL ciphers and protocols. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerSslCiphersAndProtocolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerSslCiphersAndProtocolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbSslCipherAndProtocolListResult**](LbSslCipherAndProtocolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerTcpProfiles**
> LbTcpProfileListResult ListLoadBalancerTcpProfiles(ctx, optional)
Retrieve a paginated list of load balancer TCP profiles

Retrieve a paginated list of load balancer TCP profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerTcpProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerTcpProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbTcpProfileListResult**](LbTcpProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerVirtualServerStatuses**
> LbVirtualServerStatusListResult ListLoadBalancerVirtualServerStatuses(ctx, serviceId, optional)
Get the status list of virtual servers in given load balancer service

Returns the status list of virtual servers in given load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListLoadBalancerVirtualServerStatusesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerVirtualServerStatusesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LbVirtualServerStatusListResult**](LbVirtualServerStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerVirtualServers**
> LbVirtualServerListResult ListLoadBalancerVirtualServers(ctx, optional)
Retrieve a paginated list of load balancer virtual servers

Retrieve a paginated list of load balancer virtual servers. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLoadBalancerVirtualServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerVirtualServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LbVirtualServerListResult**](LbVirtualServerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoadBalancerVirtualServersStatistics**
> LbVirtualServerStatisticsListResult ListLoadBalancerVirtualServersStatistics(ctx, serviceId, optional)
Get the statistics list of virtual servers

Returns the statistics list of virtual servers in given load balancer service. Currently, only realtime mode is supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListLoadBalancerVirtualServersStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLoadBalancerVirtualServersStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**LbVirtualServerStatisticsListResult**](LbVirtualServerStatisticsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogicalPorts**
> LogicalPortListResult ListLogicalPorts(ctx, optional)
List All Logical Ports

Returns information about all configured logical switch ports. Logical switch ports connect to VM virtual network interface cards (NICs). Each logical port is associated with one logical switch. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLogicalPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLogicalPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachmentId** | **optional.String**| Logical Port attachment Id | 
 **attachmentType** | **optional.String**| Type of attachment for logical port; for query only. | 
 **bridgeClusterId** | **optional.String**| Bridge Cluster identifier | 
 **containerPortsOnly** | **optional.Bool**| Only container VIF logical ports will be returned if true | [default to false]
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **diagnostic** | **optional.Bool**| Flag to enable showing of transit logical port. | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **logicalSwitchId** | **optional.String**| Logical Switch identifier | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **parentVifId** | **optional.String**| ID of the VIF of type PARENT | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **switchingProfileId** | **optional.String**| Network Profile identifier | 
 **transportNodeId** | **optional.String**| Transport node identifier | 
 **transportZoneId** | **optional.String**| Transport zone identifier | 

### Return type

[**LogicalPortListResult**](LogicalPortListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogicalRouterPorts**
> LogicalRouterPortListResult ListLogicalRouterPorts(ctx, optional)
List Logical Router Ports

Returns information about all logical router ports. Information includes the resource_type (LogicalRouterUpLinkPort, LogicalRouterDownLinkPort, LogicalRouterLinkPort, LogicalRouterLoopbackPort, LogicalRouterCentralizedServicePort); logical_router_id (the router to which each logical router port is assigned); and any service_bindings (such as DHCP relay service). The GET request can include a query parameter (logical_router_id or logical_switch_id). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLogicalRouterPortsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLogicalRouterPortsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **logicalRouterId** | **optional.String**| Logical Router identifier | 
 **logicalSwitchId** | **optional.String**| Logical Switch identifier | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **resourceType** | **optional.String**| Resource types of logical router port | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**LogicalRouterPortListResult**](LogicalRouterPortListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogicalRouters**
> LogicalRouterListResult ListLogicalRouters(ctx, optional)
List Logical Routers

Returns information about all logical routers, including the UUID, internal and external transit network addresses, and the router type (TIER0 or TIER1). You can get information for only TIER0 routers or only the TIER1 routers by including the router_type query parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLogicalRoutersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLogicalRoutersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **routerType** | **optional.String**| Type of Logical Router | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **vrfsOnLogicalRouterId** | **optional.String**| List all VRFs on the specified logical router. | 

### Return type

[**LogicalRouterListResult**](LogicalRouterListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogicalSwitches**
> LogicalSwitchListResult ListLogicalSwitches(ctx, optional)
List all Logical Switches

Returns information about all configured logical switches.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLogicalSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLogicalSwitchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **diagnostic** | **optional.Bool**| Flag to enable showing of transit logical switch. | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **switchingProfileId** | **optional.String**| Switching Profile identifier | 
 **transportType** | **optional.String**| Mode of transport supported in the transport zone for this logical switch | 
 **transportZoneId** | **optional.String**| Transport zone identifier | 
 **uplinkTeamingPolicyName** | **optional.String**| The logical switch&#x27;s uplink teaming policy name | 
 **vlan** | **optional.Int64**| Virtual Local Area Network Identifier | 
 **vni** | **optional.Int32**| VNI of the OVERLAY LogicalSwitch(es) to return. | 

### Return type

[**LogicalSwitchListResult**](LogicalSwitchListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogicalSwitchesByState**
> LogicalSwitchStateListResult ListLogicalSwitchesByState(ctx, optional)
List logical switches by realized state

Returns a list of logical switches states that have realized state as provided as query parameter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListLogicalSwitchesByStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListLogicalSwitchesByStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Realized state of logical switches | 

### Return type

[**LogicalSwitchStateListResult**](LogicalSwitchStateListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMACSets**
> MacSetListResult ListMACSets(ctx, optional)
List MACSets

Returns paginated list of MACSets 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListMACSetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListMACSetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MacSetListResult**](MACSetListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListManualHealthChecks**
> ManualHealthCheckListResult ListManualHealthChecks(ctx, optional)
List manual health checks

Query manual health checks with list parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListManualHealthChecksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListManualHealthChecksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ManualHealthCheckListResult**](ManualHealthCheckListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMetadataProxy**
> MetadataProxyListResult ListMetadataProxy(ctx, optional)
Get a paginated list of metadata proxies

Get a paginated list of metadata proxies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListMetadataProxyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListMetadataProxyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MetadataProxyListResult**](MetadataProxyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNDRAProfiles**
> NdraProfileListResult ListNDRAProfiles(ctx, optional)
Read All IPV6 NDRA Profiles

Returns all IPv6 NDRA Profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListNDRAProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListNDRAProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NdraProfileListResult**](NDRAProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSGroups**
> NsGroupListResult ListNSGroups(ctx, optional)
List NSGroups

List the NSGroups in a paginated format. The page size is restricted to 50 NSGroups so that the size of the response remains small even in the worst case. Optionally, specify valid member types as request parameter to filter NSGroups. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListNSGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListNSGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **memberTypes** | **optional.String**| Specify member types to filter corresponding NSGroups  | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **populateReferences** | **optional.Bool**| Populate metadata of resource referenced by NSGroupExpressions  | [default to false]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NsGroupListResult**](NSGroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSProfiles**
> NsProfileListResult ListNSProfiles(ctx, optional)
List NSProfiles

List the NSProfiles created in a paginated format.The page size is restricted to 50 NSProfiles, so that the size of the response remains small even when there are high number of NSProfiles with multiple attributes and multiple attribute values for each attribute. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListNSProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListNSProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeType** | **optional.String**| Fetch NSProfiles for the given attribute type | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NsProfileListResult**](NSProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSServiceGroups**
> NsServiceGroupListResult ListNSServiceGroups(ctx, optional)
List all NSServiceGroups

Returns paginated list of NSServiceGroups 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListNSServiceGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListNSServiceGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **defaultService** | **optional.Bool**| Fetch all default NSServiceGroups | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NsServiceGroupListResult**](NSServiceGroupListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSServices**
> NsServiceListResult ListNSServices(ctx, optional)
List all NSServices

Returns paginated list of NSServices 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListNSServicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListNSServicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **defaultService** | **optional.Bool**| Fetch all default NSServices | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NsServiceListResult**](NSServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSSupportedAttributes**
> NsSupportedAttributesListResult ListNSSupportedAttributes(ctx, optional)
List NSProfile supported attribute and sub-attributes

Returns supported attribute and sub-attributes for specified attribute type with their supported values, if provided in query/request parameter, else will fetch all supported attribute and sub-attributes for all supported attribute types. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListNSSupportedAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListNSSupportedAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributeSource** | **optional.String**| Fetch attributes source | 
 **attributeType** | **optional.String**| Fetch attributes and sub-attributes for the given attribute type | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NsSupportedAttributesListResult**](NSSupportedAttributesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNSSupportedAttributesTypes**
> NsSupportedAttributeTypesResult ListNSSupportedAttributesTypes(ctx, )
List NSProfile supported attribute types

Returns supported attribute type strings for NSProfile. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NsSupportedAttributeTypesResult**](NSSupportedAttributeTypesResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNatRules**
> NatRuleListResult ListNatRules(ctx, logicalRouterId, optional)
List NAT rules of the logical router

Returns paginated list of all user defined NAT rules of the specific logical router. If a rule_type is provided, only the given type of rules will be returned. If no rule_type is specified, the rule_type will be defaulted to NATv4, i.e. only the NATv4 rules will be listed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListNatRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListNatRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **ruleType** | **optional.String**| Action type for getting NAT rules | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NatRuleListResult**](NatRuleListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPBRSections**
> PbrSectionListResult ListPBRSections(ctx, optional)
List All PBR Sections

List all PBR section in paginated form. A default page size is limited to 1000 PBR sections. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListPBRSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListPBRSectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appliedTos** | **optional.String**| AppliedTo&#x27;s referenced by this section or section&#x27;s Distributed Service Rules . | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **destinations** | **optional.String**| Destinations referenced by this section&#x27;s Distributed Service Rules . | 
 **excludeAppliedToType** | **optional.String**| Resource type valid for use as AppliedTo filter in section API | 
 **filterType** | **optional.String**| Filter type | [default to FILTER]
 **includeAppliedToType** | **optional.String**| Resource type valid for use as AppliedTo filter in section API | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **services** | **optional.String**| NSService referenced by this section&#x27;s Distributed Service Rules . | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sources** | **optional.String**| Sources referenced by this section&#x27;s Distributed Service Rules . | 

### Return type

[**PbrSectionListResult**](PBRSectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPacketCaptureSessions**
> PacketCaptureSessionList ListPacketCaptureSessions(ctx, )
Get the information of all packet capture sessions

Get the information of all packet capture sessions. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PacketCaptureSessionList**](PacketCaptureSessionList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPortMirroringSession**
> PortMirroringSessionListResult ListPortMirroringSession(ctx, optional)
List all mirror sessions

List all mirror sessions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListPortMirroringSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListPortMirroringSessionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PortMirroringSessionListResult**](PortMirroringSessionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRouteMaps**
> RouteMapListResult ListRouteMaps(ctx, logicalRouterId, optional)
Paginated List of RouteMaps

Paginated List of RouteMaps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListRouteMapsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListRouteMapsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**RouteMapListResult**](RouteMapListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSIServiceProfiles**
> SiServiceProfileListResult ListSIServiceProfiles(ctx, serviceId)
List all Service Profiles of a Service.

List all service profiles of a service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**SiServiceProfileListResult**](SIServiceProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSections**
> FirewallSectionListResult ListSections(ctx, optional)
List All Firewall Sections

List all firewall section in paginated form. A default page size is limited to 1000 firewall sections. By default list of section is filtered by LAYER3 type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListSectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appliedTos** | **optional.String**| AppliedTo&#x27;s referenced by this section or section&#x27;s Distributed Service Rules . | 
 **contextProfiles** | **optional.String**| Limits results to sections having rules with specific Context Profiles. | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **deepSearch** | **optional.Bool**| Toggle to search with direct or indirect references. | [default to false]
 **destinations** | **optional.String**| Destinations referenced by this section&#x27;s Distributed Service Rules . | 
 **enforcedOn** | **optional.String**| Type of attachment for logical port; for query only. | 
 **excludeAppliedToType** | **optional.String**| Resource type valid for use as AppliedTo filter in section API | 
 **extendedSources** | **optional.String**| Limits results to sections having rules with specific Extended Sources. | 
 **filterType** | **optional.String**| Filter type | [default to FILTER]
 **includeAppliedToType** | **optional.String**| Resource type valid for use as AppliedTo filter in section API | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **locked** | **optional.Bool**| Limit results to sections which are locked/unlocked | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **searchInvalidReferences** | **optional.Bool**| Return invalid references in results. | [default to false]
 **searchScope** | **optional.String**| Limit result to sections of a specific enforcement point | 
 **services** | **optional.String**| NSService referenced by this section&#x27;s Distributed Service Rules . | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sources** | **optional.String**| Sources referenced by this section&#x27;s Distributed Service Rules . | 
 **type_** | **optional.String**| Section Type | [default to LAYER3]

### Return type

[**FirewallSectionListResult**](FirewallSectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceAttachments**
> ServiceAttachmentListResult ListServiceAttachments(ctx, )
Get all service attachments.

Returns all Service-Attachement(s) present in the system. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceAttachmentListResult**](ServiceAttachmentListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceChainMappings**
> ServiceChainMappingListResult ListServiceChainMappings(ctx, serviceId, serviceProfileId)
List all ServiceChainMappings.

List all service chain mappings in the system for the given service profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceProfileId** | **string**|  | 

### Return type

[**ServiceChainMappingListResult**](ServiceChainMappingListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceChains**
> ServiceChainListResult ListServiceChains(ctx, )
List all ServiceChains.

List all service chains in the system. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceChainListResult**](ServiceChainListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceConfigs**
> ServiceConfigListResult ListServiceConfigs(ctx, optional)
List service configs

List of all service configs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListServiceConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListServiceConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **profileType** | **optional.String**| Fetch ServiceConfig for the given attribute profile_type | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServiceConfigListResult**](ServiceConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceInsertionSections**
> ServiceInsertionSectionListResult ListServiceInsertionSections(ctx, optional)
List All Service Insertion Sections

List all Service Insertion section in paginated form. A default page size is limited to 1000 sections. By default, the list of section is filtered by L3REDIRECT type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListServiceInsertionSectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListServiceInsertionSectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appliedTos** | **optional.String**| AppliedTo&#x27;s referenced by this section or section&#x27;s Distributed Service Rules . | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **destinations** | **optional.String**| Destinations referenced by this section&#x27;s Distributed Service Rules . | 
 **excludeAppliedToType** | **optional.String**| Resource type valid for use as AppliedTo filter in section API | 
 **filterType** | **optional.String**| Filter type | [default to FILTER]
 **includeAppliedToType** | **optional.String**| Resource type valid for use as AppliedTo filter in section API | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **services** | **optional.String**| NSService referenced by this section&#x27;s Distributed Service Rules . | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **sources** | **optional.String**| Sources referenced by this section&#x27;s Distributed Service Rules . | 
 **type_** | **optional.String**| Section Type | [default to L3REDIRECT]

### Return type

[**ServiceInsertionSectionListResult**](ServiceInsertionSectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceInsertionServices**
> ServiceInsertionServiceListResult ListServiceInsertionServices(ctx, )
List all Service-Insertion Services.

List all Service-Insertion Service Definitions. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceInsertionServiceListResult**](ServiceInsertionServiceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceInsertionStatus**
> ServiceInsertionStatusListResult ListServiceInsertionStatus(ctx, )
List all service insertion status for supported contexts

List all service insertion status for supported contexts

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceInsertionStatusListResult**](ServiceInsertionStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceInstances**
> ServiceInstanceListResult ListServiceInstances(ctx, optional)
Get all Service-Instances present in system

Returns all Service-Instance(s) of all Services present in system. When request parameter (deployed_to or service_deployment_id) is provided as a part of request, it will filter out Service-Instances accordingly. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListServiceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListServiceInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployedTo** | **optional.String**| Deployed_to referenced by service instances present in system | 
 **serviceDeploymentId** | **optional.String**| Service Deployment Id using which the instances were deployed | 

### Return type

[**ServiceInstanceListResult**](ServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceInstancesForService**
> ServiceInstanceListResult ListServiceInstancesForService(ctx, serviceId)
Get all Service-Instances for Service.

Returns all Service-Instance(s) for a given Service-Insertion Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**ServiceInstanceListResult**](ServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceManagers**
> ServiceManagerListResult ListServiceManagers(ctx, )
List service managers

List all service managers. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ServiceManagerListResult**](ServiceManagerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServicePaths**
> ServicePathListResult ListServicePaths(ctx, serviceChainId, optional)
List all service paths

List all service paths for the given service chain for the given service chain id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceChainId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListServicePathsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListServicePathsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ServicePathListResult**](ServicePathListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSolutionConfigs**
> SolutionConfigListResult ListSolutionConfigs(ctx, serviceId)
Get Solution Config Information associated with a given service.

Returns Solution Config information for a given service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**SolutionConfigListResult**](SolutionConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStaticHopBfdPeers**
> StaticHopBfdPeerListResult ListStaticHopBfdPeers(ctx, logicalRouterId, optional)
List static routes BFD Peers

Returns information about all BFD peers created on specified logical router for static routes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListStaticHopBfdPeersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListStaticHopBfdPeersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**StaticHopBfdPeerListResult**](StaticHopBfdPeerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStaticRoutes**
> StaticRouteListResult ListStaticRoutes(ctx, logicalRouterId, optional)
Paginated List of Static Routes

Returns information about configured static routes, including the network address and next hops for each static route. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListStaticRoutesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListStaticRoutesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**StaticRouteListResult**](StaticRouteListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSwitchingProfiles**
> SwitchingProfilesListResult ListSwitchingProfiles(ctx, optional)
List Switching Profiles

Returns information about the system-default and user-configured switching profiles. Each switching profile has a unique ID, a display name, and various other read-only and configurable properties. The default switching profiles are assigned automatically to each switch. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListSwitchingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListSwitchingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeSystemOwned** | **optional.Bool**| Whether the list result contains system resources | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **switchingProfileType** | **optional.String**| comma-separated list of switching profile types, e.g. ?switching_profile_type&#x3D;QosSwitchingProfile,IpDiscoverySwitchingProfile | 

### Return type

[**SwitchingProfilesListResult**](SwitchingProfilesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTraceflows**
> TraceflowListResult ListTraceflows(ctx, optional)
List all Traceflow rounds

List all Traceflow rounds; if a logical port id is given as a query parameter, only those originated from the logical port are returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiListTraceflowsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListTraceflowsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **lportId** | **optional.String**| id of the source logical port where the trace flows originated | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**TraceflowListResult**](TraceflowListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodeStatusesByComputeCollectionId**
> IdfwTransportNodeStatusListResult ListTransportNodeStatusesByComputeCollectionId(ctx, ccExtId)
List all transport node and statuses based on idfw enabled ComputeCollection ID.

Retrieve all the transport node and status by idfw enabled   ComputeCollection ID in the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ccExtId** | **string**|  | 

### Return type

[**IdfwTransportNodeStatusListResult**](IdfwTransportNodeStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserSessions**
> IdfwUserSessionDataAndMappings ListUserSessions(ctx, )
Get user session data

Get user session data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdfwUserSessionDataAndMappings**](IdfwUserSessionDataAndMappings.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVendorTemplates**
> VendorTemplateListResult ListVendorTemplates(ctx, serviceId, optional)
List all VendorTemplates of a Service.

List all vendor templates of a service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiListVendorTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiListVendorTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vendorTemplateName** | **optional.String**| Name of vendor template | 

### Return type

[**VendorTemplateListResult**](VendorTemplateListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVirtualMachineStatusesByTransportNodeId**
> IdfwVirtualMachineStatusListResult ListVirtualMachineStatusesByTransportNodeId(ctx, transportNodeId)
List all VM and statuses based on transport node ID of idfw enabled compute collection.

Retrieve all the VM and status by transport node ID of idfw enabled compute collection in the request. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

[**IdfwVirtualMachineStatusListResult**](IdfwVirtualMachineStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LockSectionLock**
> FirewallSection LockSectionLock(ctx, body, sectionId)
Lock a section

Lock a section 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSectionLock**](FirewallSectionLock.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**FirewallSection**](FirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupAddress**
> DnsAnswer LookupAddress(ctx, forwarderId, optional)
Resolve a given address via the DNS forwarder

Query the nameserver for an ip-address or a FQDN of the given an address optionally using an specified DNS server. If the address is a fqdn, nslookup will resolve ip-address with it. If the address is an ip-address, do a reverse lookup and answer fqdn(s). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiLookupAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiLookupAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **address** | **optional.String**| IP address or FQDN for nslookup | 
 **serverIp** | **optional.String**| IPv4 address | 
 **sourceIp** | **optional.String**| IPv4 address | 

### Return type

[**DnsAnswer**](DnsAnswer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformPoolMemberAction**
> LbPool PerformPoolMemberAction(ctx, body, poolId, action)
Add, remove, or modify load balancer pool members

For ADD_MEMBERS, pool members will be created and added to load balancer pool. This action is only valid for static pool members. For REMOVE_MEMBERS, pool members will be removed from load balancer pool via IP and port in pool member settings. This action is only valid for static pool members. For UPDATE_MEMBERS, pool members admin state will be updated. This action is valid for both static pool members and dynamic pool members. For dynamic pool members, this update will be stored in customized_members field in load balancer pool member group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PoolMemberSettingList**](PoolMemberSettingList.md)|  | 
  **poolId** | **string**|  | 
  **action** | **string**| Specifies addition, removal and modification action | 

### Return type

[**LbPool**](LbPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReAllocateServiceRoutersReallocate**
> LogicalRouter ReAllocateServiceRoutersReallocate(ctx, body, logicalRouterId)
Re allocate edge node placement of TIER1 service routers

API to re allocate edge node placement for TIER1 logical router. You can re-allocate service routers of TIER1 in same edge cluster or different edge cluster. You can also place edge nodes manually and provide maximum two indices for HA mode ACTIVE_STANDBY. To re-allocate on new edge cluster you must have existing edge cluster for TIER1 logical router. This will be disruptive operation and all existing statistics of logical router will be remove. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceRouterAllocationConfig**](ServiceRouterAllocationConfig.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**LogicalRouter**](LogicalRouter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReProcessLogicalRouterReprocess**
> ReProcessLogicalRouterReprocess(ctx, logicalRouterId)
Reprocesses a logical router configuration and publish updates to controller

Reprocess logical router configuration and configuration of related entities like logical router ports, static routing, etc. Any missing Updates are published to controller. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAdvertiseRuleList**
> AdvertiseRuleList ReadAdvertiseRuleList(ctx, logicalRouterId)
Read the Advertisement Rules on a Logical Router

Returns the advertisement rule list for the specified TIER1 logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**AdvertiseRuleList**](AdvertiseRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAdvertisementConfig**
> AdvertisementConfig ReadAdvertisementConfig(ctx, logicalRouterId)
Read the Advertisement Configuration on a Logical Router

Returns information about the routes to be advertised by the specified TIER1 logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**AdvertisementConfig**](AdvertisementConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadApplProxy**
> NodeServiceProperties ReadApplProxy(ctx, )
Read the Appliance Proxy service properties

Read the Appliance Proxy service properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeServiceProperties**](NodeServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadApplProxyStatus**
> NodeServiceStatusProperties ReadApplProxyStatus(ctx, )
Read the Appliance Proxy service status

Read the Appliance Proxy service status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeServiceStatusProperties**](NodeServiceStatusProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadBGPCommunityList**
> BgpCommunityList ReadBGPCommunityList(ctx, logicalRouterId, communityListId)
Read a specific BGP community list from a logical router

Read a specific BGP community list from a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **communityListId** | **string**|  | 

### Return type

[**BgpCommunityList**](BGPCommunityList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadBgpConfig**
> BgpConfig ReadBgpConfig(ctx, logicalRouterId)
Read the BGP Configuration on a Logical Router

Returns information about the BGP configuration on a specified logical router. Information includes whether or not the BGP configuration is enabled, the AS number, and whether or not graceful restart is enabled. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**BgpConfig**](BgpConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadBgpNeighbor**
> BgpNeighbor ReadBgpNeighbor(ctx, logicalRouterId, id)
Read a specific BGP Neighbor on a Logical Router

Read a specific BGP Neighbor on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**BgpNeighbor**](BgpNeighbor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadBgpNeighborWithPasswordShowSensitiveData**
> BgpNeighbor ReadBgpNeighborWithPasswordShowSensitiveData(ctx, logicalRouterId, id)
Read a specific BGP Neighbor with password on a Logical Router

Read a specific BGP Neighbor details with password on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**BgpNeighbor**](BgpNeighbor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDADProfile**
> DadProfile ReadDADProfile(ctx, dadProfileId)
Read specified IPV6 DADProfile

Returns information about specified IPv6 DADProfile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dadProfileId** | **string**|  | 

### Return type

[**DadProfile**](DADProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDebugInfoText**
> string ReadDebugInfoText(ctx, logicalRouterId)
Read the debug information for the logical router

API to download below information as text which will be used for debugging and troubleshooting. 1) Logical router sub-components and ports. 2) Routing configuration as sent to central control plane. 3) TIER1 advertised network information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpIpPool**
> DhcpIpPool ReadDhcpIpPool(ctx, serverId, poolId)
Get a DHCP server's IP pool with the specified pool ID

Return a specific ip pool of a given logical DHCP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

[**DhcpIpPool**](DhcpIpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpProfile**
> DhcpProfile ReadDhcpProfile(ctx, profileId)
Get a DHCP server profile

Return the DHCP profile specified by the profile id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 

### Return type

[**DhcpProfile**](DhcpProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpRelay**
> DhcpRelayService ReadDhcpRelay(ctx, relayId)
Read a DHCP Relay Service

Returns the dhcp relay service information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **relayId** | **string**|  | 

### Return type

[**DhcpRelayService**](DhcpRelayService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpRelayProfile**
> DhcpRelayProfile ReadDhcpRelayProfile(ctx, relayProfileId)
Read a DHCP Relay Profile

Returns information about the specified dhcp relay profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **relayProfileId** | **string**|  | 

### Return type

[**DhcpRelayProfile**](DhcpRelayProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpServer**
> LogicalDhcpServer ReadDhcpServer(ctx, serverId)
Get a DHCP server with v4 and/or v6 servers

Retrieve a logical DHCP server specified by server id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 

### Return type

[**LogicalDhcpServer**](LogicalDhcpServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpStaticBinding**
> DhcpStaticBinding ReadDhcpStaticBinding(ctx, serverId, bindingId)
Get a DHCP server's static binding with the specified binding ID

Return a specific static binding of a given logical DHCP server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

[**DhcpStaticBinding**](DhcpStaticBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpV6IpPool**
> DhcpV6IpPool ReadDhcpV6IpPool(ctx, serverId, poolId)
Get a DHCP IPv6 server's IP pool with the specified pool ID

Return a specific ip pool of a given logical DHCP IPv6 server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

[**DhcpV6IpPool**](DhcpV6IpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDhcpV6StaticBinding**
> DhcpV6StaticBinding ReadDhcpV6StaticBinding(ctx, serverId, bindingId)
Get a DHCP IPv6 server's static binding with the specified binding ID

Return a specific static binding of a given logical DHCP IPv6 server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

[**DhcpV6StaticBinding**](DhcpV6StaticBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadDnsForwader**
> DnsForwarder ReadDnsForwader(ctx, forwarderId)
Retrieve a DNS forwarder

Retrieve a DNS forwarder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **forwarderId** | **string**|  | 

### Return type

[**DnsForwarder**](DnsForwarder.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFirewallRule**
> FirewallRule ReadFirewallRule(ctx, ruleId)
Read an Existing Rule

Return existing firewall rule information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIPPrefixList**
> IpPrefixList ReadIPPrefixList(ctx, logicalRouterId, id)
Get a specific IPPrefixList on a Logical Router

Read a specific IPPrefixList on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**IpPrefixList**](IPPrefixList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIPSet**
> IpSet ReadIPSet(ctx, ipSetId)
Read IPSet

Returns information about the specified IPSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipSetId** | **string**| IPSet Id | 

### Return type

[**IpSet**](IPSet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerApplicationProfile**
> LbAppProfile ReadLoadBalancerApplicationProfile(ctx, applicationProfileId)
Retrieve a load balancer application profile

Retrieve a load balancer application profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **applicationProfileId** | **string**|  | 

### Return type

[**LbAppProfile**](LbAppProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerClientSslProfile**
> LbClientSslProfile ReadLoadBalancerClientSslProfile(ctx, clientSslProfileId)
Retrieve a load balancer client-ssl profile

Retrieve a load balancer client-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientSslProfileId** | **string**|  | 

### Return type

[**LbClientSslProfile**](LbClientSslProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerMonitor**
> LbMonitor ReadLoadBalancerMonitor(ctx, monitorId)
Retrieve a load balancer monitor

Retrieve a load balancer monitor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **monitorId** | **string**|  | 

### Return type

[**LbMonitor**](LbMonitor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerNodeUsage**
> LbNodeUsage ReadLoadBalancerNodeUsage(ctx, nodeId)
Read load balancer usage for the given node

API is used to retrieve the usage of load balancer entities which include current number and remaining number of credits, virtual Servers, pools, pool Members and different size of LB services from the given node. Currently only Edge node is supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**LbNodeUsage**](LbNodeUsage.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerNodeUsageSummary**
> LbNodeUsageSummary ReadLoadBalancerNodeUsageSummary(ctx, optional)
Read load balancer node usage summary

API is used to retrieve the load balancer node usage summary for all nodes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ManagementPlaneApiApiReadLoadBalancerNodeUsageSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReadLoadBalancerNodeUsageSummaryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUsages** | **optional.Bool**| Whether to include node usages | 

### Return type

[**LbNodeUsageSummary**](LbNodeUsageSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerPersistenceProfile**
> LbPersistenceProfile ReadLoadBalancerPersistenceProfile(ctx, persistenceProfileId)
Retrieve a load balancer persistence profile

Retrieve a load balancer persistence profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **persistenceProfileId** | **string**|  | 

### Return type

[**LbPersistenceProfile**](LbPersistenceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerPool**
> LbPool ReadLoadBalancerPool(ctx, poolId)
Retrieve a load balancer pool

Retrieve a load balancer pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**|  | 

### Return type

[**LbPool**](LbPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerRule**
> LbRule ReadLoadBalancerRule(ctx, ruleId)
Retrieve a load balancer rule

Retrieve a load balancer rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **string**|  | 

### Return type

[**LbRule**](LbRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerServerSslProfile**
> LbServerSslProfile ReadLoadBalancerServerSslProfile(ctx, serverSslProfileId)
Retrieve a load balancer server-ssl profile

Retrieve a load balancer server-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serverSslProfileId** | **string**|  | 

### Return type

[**LbServerSslProfile**](LbServerSslProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerService**
> LbService ReadLoadBalancerService(ctx, serviceId)
Retrieve a load balancer service

Retrieve a load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**LbService**](LbService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerServiceDebugInfo**
> LbServiceDebugInfo ReadLoadBalancerServiceDebugInfo(ctx, serviceId)
Read the debug information of the load balancer service

API to download below information which will be used for debugging and troubleshooting. 1) Load balancer service 2) Load balancer associated virtual servers 3) Load balancer associated pools 4) Load balancer associated profiles such as persistence, SSL, application. 5) Load balancer associated monitors 6) Load balancer associated rules 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**LbServiceDebugInfo**](LbServiceDebugInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerServiceUsage**
> LbServiceUsage ReadLoadBalancerServiceUsage(ctx, serviceId)
Read the usage information of the given load balancer service

API to fetch the capacity and current usage of the given load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 

### Return type

[**LbServiceUsage**](LbServiceUsage.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerTcpProfile**
> LbTcpProfile ReadLoadBalancerTcpProfile(ctx, tcpProfileId)
Retrieve a load balancer TCP profile

Retrieve a load balancer TCP profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tcpProfileId** | **string**|  | 

### Return type

[**LbTcpProfile**](LbTcpProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLoadBalancerVirtualServer**
> LbVirtualServer ReadLoadBalancerVirtualServer(ctx, virtualServerId)
Retrieve a load balancer virtual server

Retrieve a load balancer virtual server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **virtualServerId** | **string**|  | 

### Return type

[**LbVirtualServer**](LbVirtualServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLogicalRouter**
> LogicalRouter ReadLogicalRouter(ctx, logicalRouterId)
Read Logical Router

Returns information about the specified logical router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**LogicalRouter**](LogicalRouter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLogicalRouterPort**
> LogicalRouterPort ReadLogicalRouterPort(ctx, logicalRouterPortId)
Read Logical Router Port

Returns information about the specified logical router port.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterPortId** | **string**|  | 

### Return type

[**LogicalRouterPort**](LogicalRouterPort.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMACSet**
> MacSet ReadMACSet(ctx, macSetId)
Read MACSet

Returns information about the specified MACSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macSetId** | **string**| MACSet Id | 

### Return type

[**MacSet**](MACSet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMetadataProxy**
> MetadataProxy ReadMetadataProxy(ctx, proxyId)
Get a metadata proxy

Get a metadata proxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **proxyId** | **string**|  | 

### Return type

[**MetadataProxy**](MetadataProxy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNDRAProfile**
> NdraProfile ReadNDRAProfile(ctx, ndRaProfileId)
Read specified IPV6 NDRA Profile

Returns information about specified IPv6 NDRA Profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ndRaProfileId** | **string**|  | 

### Return type

[**NdraProfile**](NDRAProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSGroup**
> NsGroup ReadNSGroup(ctx, nsGroupId, optional)
Read NSGroup

Returns information about the specified NSGroup. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsGroupId** | **string**| NSGroup Id | 
 **optional** | ***ManagementPlaneApiApiReadNSGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReadNSGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **populateReferences** | **optional.Bool**| Populate metadata of resource referenced by NSGroupExpressions  | [default to false]

### Return type

[**NsGroup**](NSGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSProfile**
> NsProfile ReadNSProfile(ctx, nsProfileId)
Read NSProfile

Returns information about the specified NSProfile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsProfileId** | **string**| NSProfile Id | 

### Return type

[**NsProfile**](NSProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSService**
> NsService ReadNSService(ctx, nsServiceId)
Read NSService

Returns information about the specified NSService 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsServiceId** | **string**| NSService Id | 

### Return type

[**NsService**](NSService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNSServiceGroup**
> NsServiceGroup ReadNSServiceGroup(ctx, nsServiceGroupId)
Read NSServiceGroup

Returns information about the specified NSServiceGroup 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nsServiceGroupId** | **string**| NSServiceGroup Id | 

### Return type

[**NsServiceGroup**](NSServiceGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPacketCaptureSession**
> PacketCaptureSession ReadPacketCaptureSession(ctx, sessionId)
Get the status of packet capture session

Get the packet capture status information by session id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**| Packet capture session id | 

### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRedistributionConfig**
> RedistributionConfig ReadRedistributionConfig(ctx, logicalRouterId)
Read the Redistribution Configuration on a Logical Router

Returns information about configured route redistribution for the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**RedistributionConfig**](RedistributionConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRedistributionRuleList**
> RedistributionRuleList ReadRedistributionRuleList(ctx, logicalRouterId)
Read All the Redistribution Rules on a Logical Router

Returns all the route redistribution rules for the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**RedistributionRuleList**](RedistributionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRouteMap**
> RouteMap ReadRouteMap(ctx, logicalRouterId, id)
Get a specific RouteMap on a Logical Router

Read a specific RouteMap on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**RouteMap**](RouteMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoutingBfdConfig**
> BfdConfig ReadRoutingBfdConfig(ctx, logicalRouterId)
Read the Routing BFD Configuration

Returns the BFD configuration for all routing BFD peers. This will be inherited |   by all BFD peers for LogicalRouter unless overriden while configuring the Peer. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**BfdConfig**](BfdConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRoutingConfig**
> RoutingConfig ReadRoutingConfig(ctx, logicalRouterId)
Read the Routing Configuration

Returns the routing configuration for a specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 

### Return type

[**RoutingConfig**](RoutingConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadServiceConfig**
> ServiceConfig ReadServiceConfig(ctx, configSetId)
Read Service Config

Returns information about the specified Service Config. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configSetId** | **string**| Service Config Id | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadStaticHopBfdPeer**
> StaticHopBfdPeer ReadStaticHopBfdPeer(ctx, logicalRouterId, bfdPeerId)
Read a static route BFD peer

Read the BFD peer having specified ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **bfdPeerId** | **string**|  | 

### Return type

[**StaticHopBfdPeer**](StaticHopBfdPeer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadStaticRoute**
> StaticRoute ReadStaticRoute(ctx, logicalRouterId, id)
Get a specific Static Route on a Logical Router

Read a specific static routes on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**StaticRoute**](StaticRoute.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReallocateDhcpProfileEdgeClusterReallocate**
> DhcpProfile ReallocateDhcpProfileEdgeClusterReallocate(ctx, body, serverProfileId)
Reallocate edge cluster and members of given DHCP profile.

As changing edge-cluster-id of a DhcpProfile by a PUT is disallowed, this re-allocate API is used to modify the edge-cluster-id and members of a given DhcpProfile.  Only the edge-cluster-id and the edge-cluster-member-indexes fields will be picked up by this re-allication API. The othere fields in the payload will be ignored.  If the edge-cluster-id in the payload DhcpProfile is different from the current edge-cluster-id of the profile, the referencing DHCP server(s) will be re-allocated to the new edge cluster. If the edge-cluster-id is not changed, the referencing DHCP server(s) will be re-allocated to the given edge members in the edge cluster. In this case, this REST API will act same as that of updating a DhcpProfile.  If the edge cluster member indexes are provided, they should exist in the given edge cluster. If the indexes are not specified in the DhcpProfile, edge members will be auto-allocated from the given edge cluster.  Please note that re-allocating edge-cluster will cause lose of all exisitng DHCP lease information. This API is used only when loosing DHCP leases is not a real problem, e.g. cross-site migration or failover and all client hosts will be reboot and get new IP addresses. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpProfile**](DhcpProfile.md)|  | 
  **serverProfileId** | **string**|  | 

### Return type

[**DhcpProfile**](DhcpProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterServiceManager**
> ServiceManager RegisterServiceManager(ctx, body)
Register service manager

Register service-manager with NSX with basic details like name, username, password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceManager**](ServiceManager.md)|  | 

### Return type

[**ServiceManager**](ServiceManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMACAddress**
> RemoveMACAddress(ctx, macSetId, macAddress)
Remove a MAC address from given MACSet

Remove an individual MAC address from a MACSet 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **macSetId** | **string**| MACSet Id | 
  **macAddress** | **string**| MAC address to be removed | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMemberRemoveMember**
> ResourceReference RemoveMemberRemoveMember(ctx, objectId, optional)
Remove an existing object from the exclude list

Remove an existing object from the exclude list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| identifier of the object | 
 **optional** | ***ManagementPlaneApiApiRemoveMemberRemoveMemberOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiRemoveMemberRemoveMemberOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deepCheck** | **optional.Bool**| Check all parents | [default to false]
 **objectType** | **optional.String**| Object type of an entity | 

### Return type

[**ResourceReference**](ResourceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveServiceInsertionExcludeListMemberRemoveMember**
> ResourceReference RemoveServiceInsertionExcludeListMemberRemoveMember(ctx, objectId)
Remove an existing object from the exclude list

Remove an existing object from the exclude list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| Identifier of the object | 

### Return type

[**ResourceReference**](ResourceReference.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetFirewallRuleStatsReset**
> ResetFirewallRuleStatsReset(ctx, category)
Reset firewall rule statistics

Sets firewall rule statistics counter to zero. This operation is supported for given category, for example: L3DFW i.e. for all layer3 firewall (transport nodes only) rules or L3EDGE i.e. for all layer3 edge firewall (edge nodes only) rules or L3BRIDGEPORT i.e. for all layer3 bridge port firewall (bridge ports only) rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **category** | **string**| Aggregation statistic category | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetIPSecVPNSessionStatisticsReset**
> ResetIPSecVPNSessionStatisticsReset(ctx, sessionId)
Reset the statistics of the given VPN session

Reset the statistics of the given VPN session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResolveSourceEntities**
> SourceEntityResult ResolveSourceEntities(ctx, sourceNodeValue)
Resolve 'source node id' value to source entities.

Service insertion data path inserts unique 'source node id' value into each packet. This API can be used to identify the source of the packet using this value. It can be resolved to multiple source entities. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sourceNodeValue** | **string**| value | 

### Return type

[**SourceEntityResult**](SourceEntityResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartPacketCaptureSessionRestart**
> PacketCaptureSession RestartPacketCaptureSessionRestart(ctx, sessionId)
Restart the packet capture session

Restart the packet capture session 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**| Packet capture session id | 

### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevisePBRRuleRevise**
> PbrRule RevisePBRRuleRevise(ctx, body, sectionId, ruleId, optional)
Update an Existing Rule and Reorder the Rule

Modifies existing PBR rule along with relative position among other PBR rules inside a PBR section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrRule**](PbrRule.md)|  | 
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiRevisePBRRuleReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiRevisePBRRuleReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**PbrRule**](PBRRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevisePBRSectionRevise**
> PbrSection RevisePBRSectionRevise(ctx, body, sectionId, optional)
Update an Existing Section, including Its Position

Modifies an existing PBR section along with its relative position among other PBR sections in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrSection**](PbrSection.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiRevisePBRSectionReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiRevisePBRSectionReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**PbrSection**](PBRSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevisePBRSectionWithRulesReviseWithRules**
> PbrSectionRuleList RevisePBRSectionWithRulesReviseWithRules(ctx, body, sectionId, optional)
Update an Existing Section with Rules

Modifies an existing PBR section along with its relative position among other PBR sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds.  Instead, to move a section above or below another section, use: POST /api/v1/pbr/sections/&lt;section-id&gt;?action=revise  To modify rules, use: PUT /api/v1/pbr/sections/&lt;section-id&gt;/rules/&lt;rule-id&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrSectionRuleList**](PbrSectionRuleList.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiRevisePBRSectionWithRulesReviseWithRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiRevisePBRSectionWithRulesReviseWithRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**PbrSectionRuleList**](PBRSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseRuleRevise**
> FirewallRule ReviseRuleRevise(ctx, body, sectionId, ruleId, optional)
Update an Existing Rule and Reorder the Rule

Modifies existing firewall rule along with relative position among other firewall rules inside a firewall section. Revising firewall rule in a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallRule**](FirewallRule.md)|  | 
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiReviseRuleReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReviseRuleReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseSectionRevise**
> FirewallSection ReviseSectionRevise(ctx, body, sectionId, optional)
Update an Existing Section, Including Its Position

Modifies an existing firewall section along with its relative position among other firewall sections in the system. Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSection**](FirewallSection.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiReviseSectionReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReviseSectionReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**FirewallSection**](FirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseSectionWithRulesReviseWithRules**
> FirewallSectionRuleList ReviseSectionWithRulesReviseWithRules(ctx, body, sectionId, optional)
Update an Existing Section with Rules

Modifies an existing firewall section along with its relative position among other firewall sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds.  Instead, to move a section above or below another section, use: POST /api/v1/firewall/sections/&lt;section-id&gt;?action=revise  To modify rules, use: PUT /api/v1/firewall/sections/&lt;section-id&gt;/rules/&lt;rule-id&gt;  Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSectionRuleList**](FirewallSectionRuleList.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiReviseSectionWithRulesReviseWithRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReviseSectionWithRulesReviseWithRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**FirewallSectionRuleList**](FirewallSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseServiceInsertionRuleRevise**
> ServiceInsertionRule ReviseServiceInsertionRuleRevise(ctx, body, sectionId, ruleId, optional)
Update an Existing Rule and Reorder the Rule

Modifies existing serviceinsertion rule along with relative position among other serviceinsertion rules inside a serviceinsertion section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionRule**](ServiceInsertionRule.md)|  | 
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiReviseServiceInsertionRuleReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReviseServiceInsertionRuleReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**ServiceInsertionRule**](ServiceInsertionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseServiceInsertionSectionRevise**
> ServiceInsertionSection ReviseServiceInsertionSectionRevise(ctx, body, sectionId, optional)
Update an Existing Section, Including Its Position

Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections in the system. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionSection**](ServiceInsertionSection.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiReviseServiceInsertionSectionReviseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReviseServiceInsertionSectionReviseOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**ServiceInsertionSection**](ServiceInsertionSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviseServiceInsertionSectionWithRulesReviseWithRules**
> ServiceInsertionSectionRuleList ReviseServiceInsertionSectionWithRulesReviseWithRules(ctx, body, sectionId, optional)
Update an Existing Section with Rules

Modifies an existing serviceinsertion section along with its relative position among other serviceinsertion sections with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds.  Instead, to move a section above or below another section, use: POST /api/v1/serviceinsertion/sections/&lt;section-id&gt;?action=revise  To modify rules, use: PUT /api/v1/serviceinsertion/sections/&lt;section-id&gt;/rules/&lt;rule-id&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionSectionRuleList**](ServiceInsertionSectionRuleList.md)|  | 
  **sectionId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiReviseServiceInsertionSectionWithRulesReviseWithRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiReviseServiceInsertionSectionWithRulesReviseWithRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.**| Identifier of the anchor rule or section. This is a required field in case operation like &#x27;insert_before&#x27; and &#x27;insert_after&#x27;. | 
 **operation** | **optional.**| Operation | [default to insert_top]

### Return type

[**ServiceInsertionSectionRuleList**](ServiceInsertionSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceConfigBatchOperation**
> ServiceConfigListResult ServiceConfigBatchOperation(ctx, body)
Creates/Updates service configs sent in batch request

Creates/Updates new service configs sent in batch request. This API returns ALL the service configs that are created/updated. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceConfigList**](ServiceConfigList.md)|  | 

### Return type

[**ServiceConfigListResult**](ServiceConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TerminatePacketCaptureSessionTerminate**
> PacketCaptureSession TerminatePacketCaptureSessionTerminate(ctx, sessionId)
Terminate the packet capture session by session id

Terminate the packet capture session by session id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionId** | **string**| Packet capture session id | 

### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnSetPasswordOnBgpNeighbor**
> BgpNeighbor UnSetPasswordOnBgpNeighbor(ctx, logicalRouterId, id, optional)
Unset/Delete password property on specific BGP Neighbor on Logical Router

Unset/Delete the password property on the specific BGP Neighbor. No other property of the BgpNeighbor can be updated using this API 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiUnSetPasswordOnBgpNeighborOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiUnSetPasswordOnBgpNeighborOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **action** | **optional.String**|  | 

### Return type

[**BgpNeighbor**](BgpNeighbor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlockSectionUnlock**
> FirewallSection UnlockSectionUnlock(ctx, body, sectionId)
Unlock a section

Unlock a section 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSectionLock**](FirewallSectionLock.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**FirewallSection**](FirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdvertiseRuleList**
> AdvertiseRuleList UpdateAdvertiseRuleList(ctx, body, logicalRouterId)
Update the Advertisement Rules on a Logical Router

Modifies the advertisement rules on the specified logical router. The PUT request must include all the rules with the networks parameter. Modifiable parameters are networks, display_name, and description. Set the rules list to empty to delete/clear all rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvertiseRuleList**](AdvertiseRuleList.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**AdvertiseRuleList**](AdvertiseRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdvertisementConfig**
> AdvertisementConfig UpdateAdvertisementConfig(ctx, body, logicalRouterId)
Update the Advertisement Configuration on a Logical Router

Modifies the route advertisement configuration on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdvertisementConfig**](AdvertisementConfig.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**AdvertisementConfig**](AdvertisementConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBGPCommunityList**
> BgpCommunityList UpdateBGPCommunityList(ctx, body, logicalRouterId, communityListId)
Update a specific BGP community list from a logical router

Update a specific BGP community list from a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BgpCommunityList**](BgpCommunityList.md)|  | 
  **logicalRouterId** | **string**|  | 
  **communityListId** | **string**|  | 

### Return type

[**BgpCommunityList**](BGPCommunityList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBGPCommunityListOld**
> BgpCommunityList UpdateBGPCommunityListOld(ctx, body, logicalRouterId, communityListId)
Update a specific BGP community list from a logical router

Update a specific BGP community list from a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BgpCommunityList**](BgpCommunityList.md)|  | 
  **logicalRouterId** | **string**|  | 
  **communityListId** | **string**|  | 

### Return type

[**BgpCommunityList**](BGPCommunityList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBgpConfig**
> BgpConfig UpdateBgpConfig(ctx, body, logicalRouterId)
Update the BGP Configuration on a Logical Router

Modifies the BGP configuration on a specified TIER0 logical router. Modifiable parameters include enabled, graceful_restart, as_number. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BgpConfig**](BgpConfig.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**BgpConfig**](BgpConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBgpNeighbor**
> BgpNeighbor UpdateBgpNeighbor(ctx, body, logicalRouterId, id)
Update a specific BGP Neighbor on a Logical Router

Update a specific BGP Neighbor on a Logical Router 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BgpNeighbor**](BgpNeighbor.md)|  | 
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**BgpNeighbor**](BgpNeighbor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBridgeCluster**
> BridgeCluster UpdateBridgeCluster(ctx, body, bridgeclusterId)
Update a Bridge Cluster

Modifies a existing bridge cluster. One of more transport nodes can be added or removed from the bridge cluster using this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BridgeCluster**](BridgeCluster.md)|  | 
  **bridgeclusterId** | **string**|  | 

### Return type

[**BridgeCluster**](BridgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBridgeEndpoint**
> BridgeEndpoint UpdateBridgeEndpoint(ctx, body, bridgeendpointId)
Update a Bridge Endpoint

Modifies a existing bridge endpoint. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BridgeEndpoint**](BridgeEndpoint.md)|  | 
  **bridgeendpointId** | **string**|  | 

### Return type

[**BridgeEndpoint**](BridgeEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBridgeEndpointProfile**
> BridgeEndpointProfile UpdateBridgeEndpointProfile(ctx, body, bridgeendpointprofileId)
Update a Bridge Endpoint Profile

Modifies a existing bridge endpoint profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BridgeEndpointProfile**](BridgeEndpointProfile.md)|  | 
  **bridgeendpointprofileId** | **string**|  | 

### Return type

[**BridgeEndpointProfile**](BridgeEndpointProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDADProfile**
> DadProfile UpdateDADProfile(ctx, body, dadProfileId)
Update DADProfile

Update DADProfile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DadProfile**](DadProfile.md)|  | 
  **dadProfileId** | **string**|  | 

### Return type

[**DadProfile**](DADProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpIpPool**
> DhcpIpPool UpdateDhcpIpPool(ctx, body, serverId, poolId)
Update a DHCP server's IP pool

Update a specific ip pool of a given logical DHCP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpIpPool**](DhcpIpPool.md)|  | 
  **serverId** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

[**DhcpIpPool**](DhcpIpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpProfile**
> DhcpProfile UpdateDhcpProfile(ctx, body, profileId)
Update a DHCP server profile

If both the edge_cluster_member_indexes in the DhcpProfile are changed in a same PUT API, e.g. change from [a,b] to [x,y], the current DHCP server leases will be lost, which could cause the network crash due to ip conflicts. Hence the suggestion is to change only one member index in one single update, e.g. from [a, b] to [a,y].  Please note, the edge_cluster_id in DhcpProfile can NOT be changed by this PUT operation because all existing DHCP leases will lost. If losing leases is not a problem, a dedicated re-allocation API is suggested to modify the edge-cluster-id, i.e. \"POST /api/v1/dhcp/dhcp-profiles/<profileId>?action=reallocate\".  Meanwhile, if the edge_cluster_member_indexes was specified currently but now is changed to none (not specified) via a PUT operation, the edge nodes will not be auto-selected from edge cluster. Instead, the previously-allocated edge nodes will continue to be used by the DHCP server. This is because changing both edge nodes of a DHCP server will lose all existing leases. In case re-allocation is required and leases lost is not a problem (or can be recovered), please invoke the reallocate API mentioned above with new DhcpProfile to accomplish the intent. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpProfile**](DhcpProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

[**DhcpProfile**](DhcpProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpRelay**
> DhcpRelayService UpdateDhcpRelay(ctx, body, relayId)
Update a DHCP Relay Service

Modifies the specified dhcp relay service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpRelayService**](DhcpRelayService.md)|  | 
  **relayId** | **string**|  | 

### Return type

[**DhcpRelayService**](DhcpRelayService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpRelayProfile**
> DhcpRelayProfile UpdateDhcpRelayProfile(ctx, body, relayProfileId)
Update a DHCP Relay Profile

Modifies the specified dhcp relay profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpRelayProfile**](DhcpRelayProfile.md)|  | 
  **relayProfileId** | **string**|  | 

### Return type

[**DhcpRelayProfile**](DhcpRelayProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpServer**
> LogicalDhcpServer UpdateDhcpServer(ctx, body, serverId)
Update a DHCP server with v4 and/or v6 servers

Update a logical DHCP server with new configurations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalDhcpServer**](LogicalDhcpServer.md)|  | 
  **serverId** | **string**|  | 

### Return type

[**LogicalDhcpServer**](LogicalDhcpServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpStaticBinding**
> DhcpStaticBinding UpdateDhcpStaticBinding(ctx, body, serverId, bindingId)
Update a DHCP server's static binding

Update a specific static binding of a given local DHCP server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpStaticBinding**](DhcpStaticBinding.md)|  | 
  **serverId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

[**DhcpStaticBinding**](DhcpStaticBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpV6IpPool**
> DhcpV6IpPool UpdateDhcpV6IpPool(ctx, body, serverId, poolId)
Update a DHCP IPv6 server's IP pool

Update a specific ip pool of a given logical DHCP IPv6 server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpV6IpPool**](DhcpV6IpPool.md)|  | 
  **serverId** | **string**|  | 
  **poolId** | **string**|  | 

### Return type

[**DhcpV6IpPool**](DhcpV6IpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDhcpV6StaticBinding**
> DhcpV6StaticBinding UpdateDhcpV6StaticBinding(ctx, body, serverId, bindingId)
Update a DHCP IPv6 server's static binding

Update a specific static binding of a given local DHCP IPv6 server.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DhcpV6StaticBinding**](DhcpV6StaticBinding.md)|  | 
  **serverId** | **string**|  | 
  **bindingId** | **string**|  | 

### Return type

[**DhcpV6StaticBinding**](DhcpV6StaticBinding.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDnsForwarder**
> DnsForwarder UpdateDnsForwarder(ctx, body, forwarderId)
Update a specific DNS forwarder

Update a specific DNS forwarder. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DnsForwarder**](DnsForwarder.md)|  | 
  **forwarderId** | **string**|  | 

### Return type

[**DnsForwarder**](DnsForwarder.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnabledComputeCollection**
> IdfwEnabledComputeCollection UpdateEnabledComputeCollection(ctx, body, ccExtId)
Update IDFW compute collection

Enable/disable individual compute collections for IDFW. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdfwEnabledComputeCollection**](IdfwEnabledComputeCollection.md)|  | 
  **ccExtId** | **string**|  | 

### Return type

[**IdfwEnabledComputeCollection**](IdfwEnabledComputeCollection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExcludeList**
> ExcludeList UpdateExcludeList(ctx, body)
Modify exclude list

Modify exclude list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExcludeList**](ExcludeList.md)|  | 

### Return type

[**ExcludeList**](ExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFirewallProfile**
> BaseFirewallProfile UpdateFirewallProfile(ctx, body, profileId)
Update a firewall profile.

Update user configurable properties of firewall profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseFirewallProfile**](BaseFirewallProfile.md)|  | 
  **profileId** | **string**|  | 

### Return type

[**BaseFirewallProfile**](BaseFirewallProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFirewallStatus**
> FirewallStatus UpdateFirewallStatus(ctx, body, contextType)
Update global firewall status for dfw context

Update global firewall status for dfw context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallStatus**](FirewallStatus.md)|  | 
  **contextType** | **string**|  | 

### Return type

[**FirewallStatus**](FirewallStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPPrefixList**
> IpPrefixList UpdateIPPrefixList(ctx, body, logicalRouterId, id)
Update a specific IPPrefixList on a Logical Router

Update a specific IPPrefixList on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpPrefixList**](IpPrefixList.md)|  | 
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**IpPrefixList**](IPPrefixList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSecVPNDPDProfile**
> IpSecVpndpdProfile UpdateIPSecVPNDPDProfile(ctx, body, ipsecVpnDpdProfileId)
Edit IPSec dead peer detection (DPD) profile

Edit IPSec dead peer detection (DPD) profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpndpdProfile**](IpSecVpndpdProfile.md)|  | 
  **ipsecVpnDpdProfileId** | **string**|  | 

### Return type

[**IpSecVpndpdProfile**](IPSecVPNDPDProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSecVPNIKEProfile**
> IpSecVpnikeProfile UpdateIPSecVPNIKEProfile(ctx, body, ipsecVpnIkeProfileId)
Edit custom IKE Profile

Edit custom IKE Profile. System owned profiles are non editable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnikeProfile**](IpSecVpnikeProfile.md)|  | 
  **ipsecVpnIkeProfileId** | **string**|  | 

### Return type

[**IpSecVpnikeProfile**](IPSecVPNIKEProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSecVPNLocalEndpoint**
> IpSecVpnLocalEndpoint UpdateIPSecVPNLocalEndpoint(ctx, body, ipsecVpnLocalEndpointId)
Edit custom IPSec local endpoint

Edit custom IPSec local endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnLocalEndpoint**](IpSecVpnLocalEndpoint.md)|  | 
  **ipsecVpnLocalEndpointId** | **string**|  | 

### Return type

[**IpSecVpnLocalEndpoint**](IPSecVPNLocalEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSecVPNPeerEndpoint**
> IpSecVpnPeerEndpoint UpdateIPSecVPNPeerEndpoint(ctx, body, ipsecVpnPeerEndpointId)
Edit custom IPSecPeerEndpoint

Edit custom IPSec peer endpoint. System owned endpoints are non editable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnPeerEndpoint**](IpSecVpnPeerEndpoint.md)|  | 
  **ipsecVpnPeerEndpointId** | **string**|  | 

### Return type

[**IpSecVpnPeerEndpoint**](IPSecVPNPeerEndpoint.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSecVPNService**
> IpSecVpnService UpdateIPSecVPNService(ctx, body, ipsecVpnServiceId)
Edit IPSec VPN service

Edit IPSec VPN service for given logical router.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnService**](IpSecVpnService.md)|  | 
  **ipsecVpnServiceId** | **string**|  | 

### Return type

[**IpSecVpnService**](IPSecVPNService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSecVPNSession**
> IpSecVpnSession UpdateIPSecVPNSession(ctx, body, ipsecVpnSessionId)
Edit IPSec VPN session

Edit IPSec VPN session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnSession**](IpSecVpnSession.md)|  | 
  **ipsecVpnSessionId** | **string**|  | 

### Return type

[**IpSecVpnSession**](IPSecVPNSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSecVPNTunnelProfile**
> IpSecVpnTunnelProfile UpdateIPSecVPNTunnelProfile(ctx, body, ipsecVpnTunnelProfileId)
Edit custom IPSecTunnelProfile

Edit custom IPSec Tunnel Profile. System owned profiles are non editable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSecVpnTunnelProfile**](IpSecVpnTunnelProfile.md)|  | 
  **ipsecVpnTunnelProfileId** | **string**|  | 

### Return type

[**IpSecVpnTunnelProfile**](IPSecVPNTunnelProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIPSet**
> IpSet UpdateIPSet(ctx, body, ipSetId)
Update IPSet

Updates the specified IPSet. Modifiable parameters include description, display_name and ip_addresses. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpSet**](IpSet.md)|  | 
  **ipSetId** | **string**| IPSet Id | 

### Return type

[**IpSet**](IPSet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpfixCollectorConfig**
> IpfixCollectorConfig UpdateIpfixCollectorConfig(ctx, body, collectorConfigId)
Update an existing IPFIX collector configuration

Update an existing IPFIX collector configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixCollectorConfig**](IpfixCollectorConfig.md)|  | 
  **collectorConfigId** | **string**|  | 

### Return type

[**IpfixCollectorConfig**](IpfixCollectorConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpfixCollectorUpmProfile**
> IpfixCollectorUpmProfile UpdateIpfixCollectorUpmProfile(ctx, body, ipfixCollectorProfileId)
Update an existing IPFIX collector profile

Update an existing IPFIX collector profile with profile ID and modified properties. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixCollectorUpmProfile**](IpfixCollectorUpmProfile.md)|  | 
  **ipfixCollectorProfileId** | **string**|  | 

### Return type

[**IpfixCollectorUpmProfile**](IpfixCollectorUpmProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpfixConfig**
> IpfixConfig UpdateIpfixConfig(ctx, body, configId)
Update an existing IPFIX configuration

Update an existing IPFIX configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixConfig**](IpfixConfig.md)|  | 
  **configId** | **string**|  | 

### Return type

[**IpfixConfig**](IpfixConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpfixUpmProfile**
> IpfixUpmProfile UpdateIpfixUpmProfile(ctx, body, ipfixProfileId)
Update an existing IPFIX profile

Update an existing IPFIX profile with profile ID and modified properties. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixUpmProfile**](IpfixUpmProfile.md)|  | 
  **ipfixProfileId** | **string**|  | 

### Return type

[**IpfixUpmProfile**](IpfixUpmProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateL2VpnService**
> L2VpnService UpdateL2VpnService(ctx, body, l2vpnServiceId)
Edit a L2VPN service

Edit a specific L2VPN service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**L2VpnService**](L2VpnService.md)|  | 
  **l2vpnServiceId** | **string**|  | 

### Return type

[**L2VpnService**](L2VpnService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateL2VpnSession**
> L2VpnSession UpdateL2VpnSession(ctx, body, l2vpnSessionId)
Edit a L2VPN session

Edit a specific L2VPN session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**L2VpnSession**](L2VpnSession.md)|  | 
  **l2vpnSessionId** | **string**|  | 

### Return type

[**L2VpnSession**](L2VpnSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerApplicationProfile**
> LbAppProfile UpdateLoadBalancerApplicationProfile(ctx, body, applicationProfileId)
Update a load balancer application profile

Update a load balancer application profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbAppProfile**](LbAppProfile.md)|  | 
  **applicationProfileId** | **string**|  | 

### Return type

[**LbAppProfile**](LbAppProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerClientSslProfile**
> LbClientSslProfile UpdateLoadBalancerClientSslProfile(ctx, body, clientSslProfileId)
Update a load balancer client-ssl profile

Update a load balancer client-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbClientSslProfile**](LbClientSslProfile.md)|  | 
  **clientSslProfileId** | **string**|  | 

### Return type

[**LbClientSslProfile**](LbClientSslProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerMonitor**
> LbMonitor UpdateLoadBalancerMonitor(ctx, body, monitorId)
Update a load balancer monitor

Update a load balancer monitor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbMonitor**](LbMonitor.md)|  | 
  **monitorId** | **string**|  | 

### Return type

[**LbMonitor**](LbMonitor.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerPersistenceProfile**
> LbPersistenceProfile UpdateLoadBalancerPersistenceProfile(ctx, body, persistenceProfileId)
Update a load balancer persistence profile

Update a load balancer persistence profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbPersistenceProfile**](LbPersistenceProfile.md)|  | 
  **persistenceProfileId** | **string**|  | 

### Return type

[**LbPersistenceProfile**](LbPersistenceProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerPool**
> LbPool UpdateLoadBalancerPool(ctx, body, poolId)
Update a load balancer pool

Update a load balancer pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbPool**](LbPool.md)|  | 
  **poolId** | **string**|  | 

### Return type

[**LbPool**](LbPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerRule**
> LbRule UpdateLoadBalancerRule(ctx, body, ruleId)
Update a load balancer rule

Update a load balancer rule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbRule**](LbRule.md)|  | 
  **ruleId** | **string**|  | 

### Return type

[**LbRule**](LbRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerServerSslProfile**
> LbServerSslProfile UpdateLoadBalancerServerSslProfile(ctx, body, serverSslProfileId)
Update a load balancer server-ssl profile

Update a load balancer server-ssl profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbServerSslProfile**](LbServerSslProfile.md)|  | 
  **serverSslProfileId** | **string**|  | 

### Return type

[**LbServerSslProfile**](LbServerSslProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerService**
> LbService UpdateLoadBalancerService(ctx, body, serviceId)
Update a load balancer service

Update a load balancer service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbService**](LbService.md)|  | 
  **serviceId** | **string**|  | 

### Return type

[**LbService**](LbService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerTcpProfile**
> LbTcpProfile UpdateLoadBalancerTcpProfile(ctx, body, tcpProfileId)
Update a load balancer TCP profile

Update a load balancer TCP profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbTcpProfile**](LbTcpProfile.md)|  | 
  **tcpProfileId** | **string**|  | 

### Return type

[**LbTcpProfile**](LbTcpProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerVirtualServer**
> LbVirtualServer UpdateLoadBalancerVirtualServer(ctx, body, virtualServerId)
Update a load balancer virtual server

Update a load balancer virtual server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbVirtualServer**](LbVirtualServer.md)|  | 
  **virtualServerId** | **string**|  | 

### Return type

[**LbVirtualServer**](LbVirtualServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLoadBalancerVirtualServerWithRulesUpdateWithRules**
> LbVirtualServerWithRule UpdateLoadBalancerVirtualServerWithRulesUpdateWithRules(ctx, body, virtualServerId)
Update a load balancer virtual server with rules

It is used to update virtual servers, the associated rules and update the binding of virtual server and rules. To add new rules, make sure the rules which have no identifier specified, the new rules are automatically generated and associated to the virtual server. To delete old rules, the rules should not be configured in new action, the UUID of deleted rules should be also removed from rule_ids. To update rules, the rules should be specified with new change and configured with identifier. If there are some rules which are not modified, those rule should not be specified in the rules list, the UUID list of rules should be specified in rule_ids of LbVirtualServer. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LbVirtualServerWithRule**](LbVirtualServerWithRule.md)|  | 
  **virtualServerId** | **string**|  | 

### Return type

[**LbVirtualServerWithRule**](LbVirtualServerWithRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLogicalPort**
> LogicalPort UpdateLogicalPort(ctx, body, lportId)
Update a Logical Port

Modifies an existing logical switch port. Parameters that can be modified include attachment_type (LOGICALROUTER, VIF), admin_state (UP or DOWN), attachment id and switching_profile_ids. You cannot modify the logical_switch_id. In other words, you cannot move an existing port from one switch to another switch. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalPort**](LogicalPort.md)|  | 
  **lportId** | **string**|  | 

### Return type

[**LogicalPort**](LogicalPort.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLogicalRouter**
> LogicalRouter UpdateLogicalRouter(ctx, body, logicalRouterId)
Update a Logical Router

Modifies the specified logical router. Modifiable attributes include the internal_transit_network, external_transit_networks, and edge_cluster_id (for TIER0 routers). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalRouter**](LogicalRouter.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**LogicalRouter**](LogicalRouter.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLogicalRouterPort**
> LogicalRouterPort UpdateLogicalRouterPort(ctx, body, logicalRouterPortId)
Update a Logical Router Port

Modifies the specified logical router port. Required parameters include the resource_type and logical_router_id. Modifiable parameters include the resource_type (LogicalRouterUpLinkPort, LogicalRouterDownLinkPort, LogicalRouterLinkPort, LogicalRouterLoopbackPort, LogicalRouterCentralizedServicePort), logical_router_id (to reassign the port to a different router), and service_bindings. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalRouterPort**](LogicalRouterPort.md)|  | 
  **logicalRouterPortId** | **string**|  | 

### Return type

[**LogicalRouterPort**](LogicalRouterPort.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLogicalSwitch**
> LogicalSwitch UpdateLogicalSwitch(ctx, body, lswitchId)
Update a Logical Switch

Modifies attributes of an existing logical switch. Modifiable attributes include admin_state, replication_mode, switching_profile_ids and VLAN spec. You cannot modify the original transport_zone_id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LogicalSwitch**](LogicalSwitch.md)|  | 
  **lswitchId** | **string**|  | 

### Return type

[**LogicalSwitch**](LogicalSwitch.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMACSet**
> MacSet UpdateMACSet(ctx, body, macSetId)
Update MACSet

Updates the specified MACSet. Modifiable parameters include the description, display_name and mac_addresses. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MacSet**](MacSet.md)|  | 
  **macSetId** | **string**| MACSet Id | 

### Return type

[**MacSet**](MACSet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMasterSwitchSetting**
> IdfwMasterSwitchSetting UpdateMasterSwitchSetting(ctx, body)
Update IDFW master switch setting enabled/disabled

Update Identity Firewall master switch setting (true=enabled / false=disabled). Identity Firewall master switch setting enables or disables Identity Firewall feature across the system.  It affects compute collections, hypervisor and virtual machines.  This operation is expensive and also has big impact and implication on system perforamce. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdfwMasterSwitchSetting**](IdfwMasterSwitchSetting.md)|  | 

### Return type

[**IdfwMasterSwitchSetting**](IdfwMasterSwitchSetting.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMetadataProxy**
> MetadataProxy UpdateMetadataProxy(ctx, body, proxyId)
Update a metadata proxy

Update a metadata proxy

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MetadataProxy**](MetadataProxy.md)|  | 
  **proxyId** | **string**|  | 

### Return type

[**MetadataProxy**](MetadataProxy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNDRAProfile**
> NdraProfile UpdateNDRAProfile(ctx, body, ndRaProfileId)
Update NDRA Profile

Update NDRAProfile 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NdraProfile**](NdraProfile.md)|  | 
  **ndRaProfileId** | **string**|  | 

### Return type

[**NdraProfile**](NDRAProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNSGroup**
> NsGroup UpdateNSGroup(ctx, body, nsGroupId)
Update NSGroup

Updates the specified NSGroup. Modifiable parameters include the description, display_name and members. For NSGroups containing VM criteria(both static and dynamic), system VMs will not be included as members. This filter applies at VM level only. Exceptions are as follows. 1. LogicalPorts and VNI of system VMs will be included in NSGroup if the criteria  is based on LogicalPort, LogicalSwitch or VNI directly. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsGroup**](NsGroup.md)|  | 
  **nsGroupId** | **string**| NSGroup Id | 

### Return type

[**NsGroup**](NSGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNSProfile**
> NsProfile UpdateNSProfile(ctx, body, nsProfileId)
Update NSProfile

Updates the specified NSProfile. Rules for using attributes and sub-attributes in single NSProfile 1. One type of attribute can't have multiple occurrences. ( Example -  Attribute type APP_ID can be used only once per NSProfile.) 2. Values for an attribute are mentioned as array of strings.  ( Example - For type APP_ID , values can be mentioned as [\"SSL\",\"FTP\"].) 3. If sub-attribtes are mentioned for an attribute, then only single  value is allowed for that attribute. 4. To get a list of supported  attributes and sub-attributes fire the following REST API  GET https://&lt;nsx-mgr&gt;/api/v1/ns-profiles/attributes 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsProfile**](NsProfile.md)|  | 
  **nsProfileId** | **string**| NSProfile Id | 

### Return type

[**NsProfile**](NSProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNSService**
> NsService UpdateNSService(ctx, body, nsServiceId)
Update NSService

Updates the specified NSService. Modifiable parameters include the description, display_name and the NSService element. The system defined NSServices can't be modified 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsService**](NsService.md)|  | 
  **nsServiceId** | **string**| NSService Id | 

### Return type

[**NsService**](NSService.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNSServiceGroup**
> NsServiceGroup UpdateNSServiceGroup(ctx, body, nsServiceGroupId)
Update NSServiceGroup

Updates the specified NSService. Modifiable parameters include the description, display_name and members. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NsServiceGroup**](NsServiceGroup.md)|  | 
  **nsServiceGroupId** | **string**| NSServiceGroup Id | 

### Return type

[**NsServiceGroup**](NSServiceGroup.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNatRule**
> NatRule UpdateNatRule(ctx, body, logicalRouterId, ruleId)
Update a specific NAT rule from a given logical router

Update a specific NAT rule from a given logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NatRule**](NatRule.md)|  | 
  **logicalRouterId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**NatRule**](NatRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePBRRule**
> PbrRule UpdatePBRRule(ctx, body, sectionId, ruleId)
Update an Existing Rule

Modifies existing rule in a PBR section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrRule**](PbrRule.md)|  | 
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**PbrRule**](PBRRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePBRSection**
> PbrSection UpdatePBRSection(ctx, body, sectionId)
Update an Existing Section

Modifies the specified section, but does not modify the section's associated rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrSection**](PbrSection.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**PbrSection**](PBRSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePBRSectionWithRulesUpdateWithRules**
> PbrSectionRuleList UpdatePBRSectionWithRulesUpdateWithRules(ctx, body, sectionId)
Update an Existing Section, Including Its Rules

Modifies existing PBR section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds.  Instead, to update rule content, use: PUT /api/v1/pbr/sections/&lt;section-id&gt;/rules/&lt;rule-id&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PbrSectionRuleList**](PbrSectionRuleList.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**PbrSectionRuleList**](PBRSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePortMirroringSession**
> PortMirroringSession UpdatePortMirroringSession(ctx, body, mirrorSessionId)
Update the mirror session

Update the mirror session

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PortMirroringSession**](PortMirroringSession.md)|  | 
  **mirrorSessionId** | **string**|  | 

### Return type

[**PortMirroringSession**](PortMirroringSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRedistributionConfig**
> RedistributionConfig UpdateRedistributionConfig(ctx, body, logicalRouterId)
Update the Redistribution Configuration on a Logical Router

Modifies existing route redistribution rules for the specified TIER0 logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedistributionConfig**](RedistributionConfig.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**RedistributionConfig**](RedistributionConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRedistributionRuleList**
> RedistributionRuleList UpdateRedistributionRuleList(ctx, body, logicalRouterId)
Update All the Redistribution Rules on a Logical Router

Modifies all route redistribution rules for the specified TIER0 logical router. Set the rules list to empty to delete/clear all rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedistributionRuleList**](RedistributionRuleList.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**RedistributionRuleList**](RedistributionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRouteMap**
> RouteMap UpdateRouteMap(ctx, body, logicalRouterId, id)
Update a specific RouteMap on a Logical Router

Update a specific RouteMap on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RouteMap**](RouteMap.md)|  | 
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**RouteMap**](RouteMap.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoutingBfdConfig**
> BfdConfig UpdateRoutingBfdConfig(ctx, body, logicalRouterId)
Update the BFD Configuration for BFD peers for routing

Modifies the BFD configuration for routing BFD peers. Note - the configuration |   changes apply only to those routing BFD peers for which the BFD configuration has |   not been overridden at Peer level. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BfdConfig**](BfdConfig.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**BfdConfig**](BfdConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRoutingConfig**
> RoutingConfig UpdateRoutingConfig(ctx, body, logicalRouterId)
Update the Routing Configuration

Modifies the routing configuration for a specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RoutingConfig**](RoutingConfig.md)|  | 
  **logicalRouterId** | **string**|  | 

### Return type

[**RoutingConfig**](RoutingConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRule**
> FirewallRule UpdateRule(ctx, body, sectionId, ruleId)
Update an Existing Rule

Modifies existing firewall rule in a firewall section. Updating firewall rule in a section modifies parent section entity and simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallRule**](FirewallRule.md)|  | 
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**FirewallRule**](FirewallRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSection**
> FirewallSection UpdateSection(ctx, body, sectionId)
Update an Existing Section

Modifies the specified section, but does not modify the section's associated rules. Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSection**](FirewallSection.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**FirewallSection**](FirewallSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSectionWithRulesUpdateWithRules**
> FirewallSectionRuleList UpdateSectionWithRulesUpdateWithRules(ctx, body, sectionId)
Update an Existing Section, Including Its Rules

Modifies existing firewall section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds.  Instead, to update rule content, use: PUT /api/v1/firewall/sections/&lt;section-id&gt;/rules/&lt;rule-id&gt;  Simultaneous update (modify) operations on same section are not allowed to prevent overwriting stale contents to firewall section. If a concurrent update is performed, HTTP response code 409 will be returned to the client operating on stale data. That client should retrieve the firewall section again and re-apply its update. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FirewallSectionRuleList**](FirewallSectionRuleList.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**FirewallSectionRuleList**](FirewallSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceAttachment**
> ServiceAttachment UpdateServiceAttachment(ctx, body, serviceAttachmentId)
Update an existing service attachment.

Modifies an existing service attachment. Updates to name, description and Logical Router list only supported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceAttachment**](ServiceAttachment.md)|  | 
  **serviceAttachmentId** | **string**|  | 

### Return type

[**ServiceAttachment**](ServiceAttachment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceConfig**
> ServiceConfig UpdateServiceConfig(ctx, body, configSetId)
Update service config

Updates the specified ServiceConfig. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceConfig**](ServiceConfig.md)|  | 
  **configSetId** | **string**| service config Id | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceDeployment**
> ServiceDeployment UpdateServiceDeployment(ctx, body, serviceId, serviceDeploymentId)
Update an existing Service Deployment.

This API is deprecated since only property we can change on service deployment is display name, which is used for the SVM name. Changing the name will cause the name of the deployment to go out of sync with the deployed VM. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceDeployment**](ServiceDeployment.md)|  | 
  **serviceId** | **string**|  | 
  **serviceDeploymentId** | **string**|  | 

### Return type

[**ServiceDeployment**](ServiceDeployment.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInsertionExcludeList**
> SiExcludeList UpdateServiceInsertionExcludeList(ctx, body)
Modify exclude list

Modify exclude list. This includes adding/removing members in the list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SiExcludeList**](SiExcludeList.md)|  | 

### Return type

[**SiExcludeList**](SIExcludeList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInsertionRule**
> ServiceInsertionRule UpdateServiceInsertionRule(ctx, body, sectionId, ruleId)
Update an Existing Rule

Modifies existing serviceinsertion rule in a serviceinsertion section. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionRule**](ServiceInsertionRule.md)|  | 
  **sectionId** | **string**|  | 
  **ruleId** | **string**|  | 

### Return type

[**ServiceInsertionRule**](ServiceInsertionRule.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInsertionSection**
> ServiceInsertionSection UpdateServiceInsertionSection(ctx, body, sectionId)
Update an Existing Section

Modifies the specified section, but does not modify the section's associated rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionSection**](ServiceInsertionSection.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**ServiceInsertionSection**](ServiceInsertionSection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInsertionSectionWithRulesUpdateWithRules**
> ServiceInsertionSectionRuleList UpdateServiceInsertionSectionWithRulesUpdateWithRules(ctx, body, sectionId)
Update an Existing Section, Including Its Rules

Modifies existing serviceinsertion section along with its association with rules. When invoked on a large number of rules, this API is supported only at low rates of invocation (not more than 2 times per minute). The typical latency of this API with about 1024 rules is about 15 seconds in a cluster setup. This API should not be invoked with large payloads at automation speeds.  Instead, to update rule content, use: PUT /api/v1/serviceinsertion/sections/&lt;section-id&gt;/rules/&lt;rule-id&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionSectionRuleList**](ServiceInsertionSectionRuleList.md)|  | 
  **sectionId** | **string**|  | 

### Return type

[**ServiceInsertionSectionRuleList**](ServiceInsertionSectionRuleList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInsertionService**
> ServiceDefinition UpdateServiceInsertionService(ctx, body, serviceId)
Update an existing Service

Modifies the specified Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceDefinition**](ServiceDefinition.md)|  | 
  **serviceId** | **string**|  | 

### Return type

[**ServiceDefinition**](ServiceDefinition.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInsertionStatus**
> ServiceInsertionStatus UpdateServiceInsertionStatus(ctx, body, contextType)
Update global ServiceInsertion status for a context

Update global ServiceInsertion status for a context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInsertionStatus**](ServiceInsertionStatus.md)|  | 
  **contextType** | **string**|  | 

### Return type

[**ServiceInsertionStatus**](ServiceInsertionStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceInstance**
> BaseServiceInstance UpdateServiceInstance(ctx, body, serviceId, serviceInstanceId)
Update an existing Service-Instance.

Modifies an existing Service-Instance for a given Service-Insertion Service. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseServiceInstance**](BaseServiceInstance.md)|  | 
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

[**BaseServiceInstance**](BaseServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceManager**
> ServiceManager UpdateServiceManager(ctx, body, serviceManagerId)
Update service manager

Update service-manager which is registered with NSX with basic details like name, username, password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceManager**](ServiceManager.md)|  | 
  **serviceManagerId** | **string**|  | 

### Return type

[**ServiceManager**](ServiceManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceVmState**
> UpdateServiceVmState(ctx, serviceId, serviceInstanceId, instanceRuntimeId, optional)
Update maintenance mode or runtime state of a service VM

Set service VM either in or out of maintenance mode for maintenance mode, or in service or out of service for runtime state. Only one value can be set at one time. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 
  **instanceRuntimeId** | **string**|  | 
 **optional** | ***ManagementPlaneApiApiUpdateServiceVmStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManagementPlaneApiApiUpdateServiceVmStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **action** | **optional.String**|  | 
 **unhealthyReason** | **optional.String**| Reason for the unhealthy state | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSolutionConfig**
> SolutionConfig UpdateSolutionConfig(ctx, body, serviceId, solutionConfigId)
Updates Solution Config for a given Service

Updates a solution config. Solution Config are service level objects, required for configuring the NXGI partner Service after deployment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SolutionConfig**](SolutionConfig.md)|  | 
  **serviceId** | **string**|  | 
  **solutionConfigId** | **string**|  | 

### Return type

[**SolutionConfig**](SolutionConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStandaloneHostsSwitchSetting**
> IdfwStandaloneHostsSwitchSetting UpdateStandaloneHostsSwitchSetting(ctx, body)
Update IDFW master switch setting enabled/disabled

Update Identity Firewall standalone hosts switch setting (true=enabled / false=disabled). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IdfwStandaloneHostsSwitchSetting**](IdfwStandaloneHostsSwitchSetting.md)|  | 

### Return type

[**IdfwStandaloneHostsSwitchSetting**](IdfwStandaloneHostsSwitchSetting.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStaticHopBfdPeer**
> StaticHopBfdPeer UpdateStaticHopBfdPeer(ctx, body, logicalRouterId, bfdPeerId)
Update a static route BFD peer

Modifies the static route BFD peer. Modifiable parameters includes peer IP, enable flag and configuration of the BFD peer. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StaticHopBfdPeer**](StaticHopBfdPeer.md)|  | 
  **logicalRouterId** | **string**|  | 
  **bfdPeerId** | **string**|  | 

### Return type

[**StaticHopBfdPeer**](StaticHopBfdPeer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateStaticRoute**
> StaticRoute UpdateStaticRoute(ctx, body, logicalRouterId, id)
Update a specific Static Route Rule on a Logical Router

Update a specific static route on the specified logical router. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StaticRoute**](StaticRoute.md)|  | 
  **logicalRouterId** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**StaticRoute**](StaticRoute.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSwitchIpfixConfig**
> IpfixObsPointConfig UpdateSwitchIpfixConfig(ctx, body)
Update global switch IPFIX export configuration

Deprecated - Please use /ipfix-profiles/<ipfix-profile-id> for switch IPFIX profile and /ipfix-collector-profiles/<ipfix-collector-profile-id> for IPFIX collector profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpfixObsPointConfig**](IpfixObsPointConfig.md)|  | 

### Return type

[**IpfixObsPointConfig**](IpfixObsPointConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSwitchingProfile**
> BaseSwitchingProfile UpdateSwitchingProfile(ctx, body, switchingProfileId)
Update a Switching Profile

Updates the user-configurable parameters of a switching profile. Only the qos, port-mirroring, spoof-guard and port-security switching profiles can be modified. You cannot modify the ipfix or ip-discovery switching profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseSwitchingProfile**](BaseSwitchingProfile.md)|  | 
  **switchingProfileId** | **string**|  | 

### Return type

[**BaseSwitchingProfile**](BaseSwitchingProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeServiceDeploymentUpgrade**
> UpgradeServiceDeploymentUpgrade(ctx, body, serviceId, serviceDeploymentId)
Upgrade all VMs part of this service deployment to new Spec OVF.

If new deployment spec is provided, the deployment will be moved to the provided spec provided that current deployment state is either UPGRADE_FAILED or DEPLOYMENT_SUCCESSFUL If same deployment spec is provided, upgrade will be done only if current deployment state is UPGRADE_FAILED  

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeploymentSpecName**](DeploymentSpecName.md)|  | 
  **serviceId** | **string**|  | 
  **serviceDeploymentId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeServiceVMsUpgrade**
> UpgradeServiceVMsUpgrade(ctx, serviceId, serviceInstanceId)
Upgrade service VMs using newer version of OVF

Upgrade service VMs using newer version of OVF. Upgrade is a 2 step process. Update the 'deployment_spec_name' in the ServiceInstance to the new DeploymentSpec to which the service VMs will be upgraded, folowed by this 'upgrade' api. In case of HA, the stand-by service VM will be upgrade first. Once it has been upgraded, it switches to be the Active one and then the other VM will be upgrade. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceId** | **string**|  | 
  **serviceInstanceId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyPortMirroringSessionVerify**
> VerifyPortMirroringSessionVerify(ctx, mirrorSessionId)
Verify whether the mirror session is still valid

Verify whether all participants are on the same transport node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mirrorSessionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

