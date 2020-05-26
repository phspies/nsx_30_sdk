# {{classname}}

All URIs are relative to *https://nsxmanager.your.domain/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddClusterNode**](SystemAdministrationConfigurationApi.md#AddClusterNode) | **Post** /cluster/nodes | Add a controller to the cluster
[**AddClusterNodeVM**](SystemAdministrationConfigurationApi.md#AddClusterNodeVM) | **Post** /cluster/nodes/deployments | Deploy and register a cluster node VM
[**AddComputeManager**](SystemAdministrationConfigurationApi.md#AddComputeManager) | **Post** /fabric/compute-managers | Register compute manager with NSX
[**AddNode**](SystemAdministrationConfigurationApi.md#AddNode) | **Post** /fabric/nodes | Register and Install NSX Components on a Node
[**AddPaceClusterNodeVM**](SystemAdministrationConfigurationApi.md#AddPaceClusterNodeVM) | **Post** /intelligence/nodes/deployments | Deploy and register a Intelligence cluster node VM
[**AddVirtualMachineTagsAddTags**](SystemAdministrationConfigurationApi.md#AddVirtualMachineTagsAddTags) | **Post** /fabric/virtual-machines?action&#x3D;add_tags | Perform action on specified virtual machine e.g. update tags
[**AllocateOrReleaseFromIpBlockSubnet**](SystemAdministrationConfigurationApi.md#AllocateOrReleaseFromIpBlockSubnet) | **Post** /pools/ip-subnets/{subnet-id} | Allocate or Release an IP Address from a Ip Subnet
[**AllocateOrReleaseFromIpPool**](SystemAdministrationConfigurationApi.md#AllocateOrReleaseFromIpPool) | **Post** /pools/ip-pools/{pool-id} | Allocate or Release an IP Address from a Pool
[**CancelApplianceManagementTaskCancel**](SystemAdministrationConfigurationApi.md#CancelApplianceManagementTaskCancel) | **Post** /node/tasks/{task-id}?action&#x3D;cancel | Cancel specified task
[**CancelBundleUploadCancelUpload**](SystemAdministrationConfigurationApi.md#CancelBundleUploadCancelUpload) | **Post** /repository/bundles/{bundle-id}?action&#x3D;cancel_upload | Cancel bundle upload
[**ChangeNodeMode**](SystemAdministrationConfigurationApi.md#ChangeNodeMode) | **Post** /configs/node/mode | NodeMode
[**CheckRabbitMQManagementPort**](SystemAdministrationConfigurationApi.md#CheckRabbitMQManagementPort) | **Get** /node/rabbitmq-management-port | Check if RabbitMQ management port is enabled or not
[**ClearClusterCertificateClearClusterCertificate**](SystemAdministrationConfigurationApi.md#ClearClusterCertificateClearClusterCertificate) | **Post** /cluster/api-certificate?action&#x3D;clear_cluster_certificate | Clear the cluster certificate
[**ClearClusterVirtualIpClearVirtualIp**](SystemAdministrationConfigurationApi.md#ClearClusterVirtualIpClearVirtualIp) | **Post** /cluster/api-virtual-ip?action&#x3D;clear_virtual_ip | Clear cluster virtual IP address
[**CreateApplianceManagementServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateApplianceManagementServiceActionRestart) | **Post** /node/services/node-mgmt?action&#x3D;restart | Restart the node management service
[**CreateAsyncReplicatorServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateAsyncReplicatorServiceActionRestart) | **Post** /node/services/async_replicator?action&#x3D;restart | Restart, start or stop the Async Replicator service
[**CreateAsyncReplicatorServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateAsyncReplicatorServiceActionStart) | **Post** /node/services/async_replicator?action&#x3D;start | Restart, start or stop the Async Replicator service
[**CreateAsyncReplicatorServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateAsyncReplicatorServiceActionStop) | **Post** /node/services/async_replicator?action&#x3D;stop | Restart, start or stop the Async Replicator service
[**CreateClusterBootManagerServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateClusterBootManagerServiceActionRestart) | **Post** /node/services/cluster_manager?action&#x3D;restart | Restart, start or stop the cluster boot manager service
[**CreateClusterBootManagerServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateClusterBootManagerServiceActionStart) | **Post** /node/services/cluster_manager?action&#x3D;start | Restart, start or stop the cluster boot manager service
[**CreateClusterBootManagerServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateClusterBootManagerServiceActionStop) | **Post** /node/services/cluster_manager?action&#x3D;stop | Restart, start or stop the cluster boot manager service
[**CreateClusterProfile**](SystemAdministrationConfigurationApi.md#CreateClusterProfile) | **Post** /cluster-profiles | Create a Cluster Profile
[**CreateCminventoryServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateCminventoryServiceActionRestart) | **Post** /node/services/cm-inventory?action&#x3D;restart | Restart, start or stop the manager service
[**CreateCminventoryServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateCminventoryServiceActionStart) | **Post** /node/services/cm-inventory?action&#x3D;start | Restart, start or stop the manager service
[**CreateCminventoryServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateCminventoryServiceActionStop) | **Post** /node/services/cm-inventory?action&#x3D;stop | Restart, start or stop the manager service
[**CreateComputeCollectionFabricTemplate**](SystemAdministrationConfigurationApi.md#CreateComputeCollectionFabricTemplate) | **Post** /fabric/compute-collection-fabric-templates | Create a compute collection fabric template
[**CreateControllerServerServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateControllerServerServiceActionRestart) | **Post** /node/services/controller?action&#x3D;restart | Restart, start or stop the controller service
[**CreateControllerServerServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateControllerServerServiceActionStart) | **Post** /node/services/controller?action&#x3D;start | Restart, start or stop the controller service
[**CreateControllerServerServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateControllerServerServiceActionStop) | **Post** /node/services/controller?action&#x3D;stop | Restart, start or stop the controller service
[**CreateDirectoryDomain**](SystemAdministrationConfigurationApi.md#CreateDirectoryDomain) | **Post** /directory/domains | Create a directory domain
[**CreateDirectoryLdapServer**](SystemAdministrationConfigurationApi.md#CreateDirectoryLdapServer) | **Post** /directory/domains/{domain-id}/ldap-servers | Create a LDAP server for directory domain
[**CreateEdgeCluster**](SystemAdministrationConfigurationApi.md#CreateEdgeCluster) | **Post** /edge-clusters | Create Edge Cluster
[**CreateFailureDomain**](SystemAdministrationConfigurationApi.md#CreateFailureDomain) | **Post** /failure-domains | Create Failure Domain
[**CreateHostSwitchProfile**](SystemAdministrationConfigurationApi.md#CreateHostSwitchProfile) | **Post** /host-switch-profiles | Create a Hostswitch Profile
[**CreateIdpsReportingServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateIdpsReportingServiceActionRestart) | **Post** /node/services/idps-reporting?action&#x3D;restart | Restart, start or stop the idps-reporting service
[**CreateIdpsReportingServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateIdpsReportingServiceActionStart) | **Post** /node/services/idps-reporting?action&#x3D;start | Restart, start or stop the idps-reporting service
[**CreateIdpsReportingServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateIdpsReportingServiceActionStop) | **Post** /node/services/idps-reporting?action&#x3D;stop | Restart, start or stop the idps-reporting service
[**CreateIntelligenceUpgradeCoordinatorServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateIntelligenceUpgradeCoordinatorServiceActionRestart) | **Post** /node/services/intelligence-upgrade-coordinator?action&#x3D;restart | Restart, start or stop the intelligence upgrade coordinator service
[**CreateIntelligenceUpgradeCoordinatorServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateIntelligenceUpgradeCoordinatorServiceActionStart) | **Post** /node/services/intelligence-upgrade-coordinator?action&#x3D;start | Restart, start or stop the intelligence upgrade coordinator service
[**CreateIntelligenceUpgradeCoordinatorServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateIntelligenceUpgradeCoordinatorServiceActionStop) | **Post** /node/services/intelligence-upgrade-coordinator?action&#x3D;stop | Restart, start or stop the intelligence upgrade coordinator service
[**CreateIpBlock**](SystemAdministrationConfigurationApi.md#CreateIpBlock) | **Post** /pools/ip-blocks | Create a new IP address block.
[**CreateIpBlockSubnet**](SystemAdministrationConfigurationApi.md#CreateIpBlockSubnet) | **Post** /pools/ip-subnets | Create subnet of specified size within an IP block
[**CreateIpPool**](SystemAdministrationConfigurationApi.md#CreateIpPool) | **Post** /pools/ip-pools | Create an IP Pool
[**CreateLiagentServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateLiagentServiceActionRestart) | **Post** /node/services/liagent?action&#x3D;restart | Restart, start or stop the liagent service
[**CreateLiagentServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateLiagentServiceActionStart) | **Post** /node/services/liagent?action&#x3D;start | Restart, start or stop the liagent service
[**CreateLiagentServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateLiagentServiceActionStop) | **Post** /node/services/liagent?action&#x3D;stop | Restart, start or stop the liagent service
[**CreateMigrationCoordinatorServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateMigrationCoordinatorServiceActionRestart) | **Post** /node/services/migration-coordinator?action&#x3D;restart | Restart, start or stop the migration coordinator service
[**CreateMigrationCoordinatorServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateMigrationCoordinatorServiceActionStart) | **Post** /node/services/migration-coordinator?action&#x3D;start | Restart, start or stop the migration coordinator service
[**CreateMigrationCoordinatorServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateMigrationCoordinatorServiceActionStop) | **Post** /node/services/migration-coordinator?action&#x3D;stop | Restart, start or stop the migration coordinator service
[**CreateNSXMessageBusServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateNSXMessageBusServiceActionRestart) | **Post** /node/services/nsx-message-bus?action&#x3D;restart | Restart, start or stop the NSX Message Bus service
[**CreateNSXMessageBusServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateNSXMessageBusServiceActionStart) | **Post** /node/services/nsx-message-bus?action&#x3D;start | Restart, start or stop the NSX Message Bus service
[**CreateNSXMessageBusServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateNSXMessageBusServiceActionStop) | **Post** /node/services/nsx-message-bus?action&#x3D;stop | Restart, start or stop the NSX Message Bus service
[**CreateNTPServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateNTPServiceActionRestart) | **Post** /node/services/ntp?action&#x3D;restart | Restart, start or stop the NTP service
[**CreateNTPServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateNTPServiceActionStart) | **Post** /node/services/ntp?action&#x3D;start | Restart, start or stop the NTP service
[**CreateNTPServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateNTPServiceActionStop) | **Post** /node/services/ntp?action&#x3D;stop | Restart, start or stop the NTP service
[**CreateNodeNetworkRoute**](SystemAdministrationConfigurationApi.md#CreateNodeNetworkRoute) | **Post** /node/network/routes | Create node network route
[**CreateNodeStatsServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateNodeStatsServiceActionRestart) | **Post** /node/services/node-stats?action&#x3D;restart | Restart, start or stop the NSX node-stats service
[**CreateNodeStatsServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateNodeStatsServiceActionStart) | **Post** /node/services/node-stats?action&#x3D;start | Restart, start or stop the NSX node-stats service
[**CreateNodeStatsServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateNodeStatsServiceActionStop) | **Post** /node/services/node-stats?action&#x3D;stop | Restart, start or stop the NSX node-stats service
[**CreateNsxUiServiceServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateNsxUiServiceServiceActionRestart) | **Post** /node/services/ui-service?action&#x3D;restart | Restart, Start and Stop the ui service
[**CreateNsxUiServiceServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateNsxUiServiceServiceActionStart) | **Post** /node/services/ui-service?action&#x3D;start | Restart, Start and Stop the ui service
[**CreateNsxUiServiceServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateNsxUiServiceServiceActionStop) | **Post** /node/services/ui-service?action&#x3D;stop | Restart, Start and Stop the ui service
[**CreateNsxUpgradeAgentServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateNsxUpgradeAgentServiceActionRestart) | **Post** /node/services/nsx-upgrade-agent?action&#x3D;restart | Restart, start or stop the NSX upgrade agent service
[**CreateNsxUpgradeAgentServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateNsxUpgradeAgentServiceActionStart) | **Post** /node/services/nsx-upgrade-agent?action&#x3D;start | Restart, start or stop the NSX upgrade agent service
[**CreateNsxUpgradeAgentServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateNsxUpgradeAgentServiceActionStop) | **Post** /node/services/nsx-upgrade-agent?action&#x3D;stop | Restart, start or stop the NSX upgrade agent service
[**CreatePhonehomeCoordinatorServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreatePhonehomeCoordinatorServiceActionRestart) | **Post** /node/services/telemetry?action&#x3D;restart | Restart, start or stop Telemetry service
[**CreatePhonehomeCoordinatorServiceActionStart**](SystemAdministrationConfigurationApi.md#CreatePhonehomeCoordinatorServiceActionStart) | **Post** /node/services/telemetry?action&#x3D;start | Restart, start or stop Telemetry service
[**CreatePhonehomeCoordinatorServiceActionStop**](SystemAdministrationConfigurationApi.md#CreatePhonehomeCoordinatorServiceActionStop) | **Post** /node/services/telemetry?action&#x3D;stop | Restart, start or stop Telemetry service
[**CreatePlatformClientServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreatePlatformClientServiceActionRestart) | **Post** /node/services/nsx-platform-client?action&#x3D;restart | Restart, start or stop the NSX Platform Client service
[**CreatePlatformClientServiceActionStart**](SystemAdministrationConfigurationApi.md#CreatePlatformClientServiceActionStart) | **Post** /node/services/nsx-platform-client?action&#x3D;start | Restart, start or stop the NSX Platform Client service
[**CreatePlatformClientServiceActionStop**](SystemAdministrationConfigurationApi.md#CreatePlatformClientServiceActionStop) | **Post** /node/services/nsx-platform-client?action&#x3D;stop | Restart, start or stop the NSX Platform Client service
[**CreatePolicyServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreatePolicyServiceActionRestart) | **Post** /node/services/policy?action&#x3D;restart | Restart, start or stop the service
[**CreatePolicyServiceActionStart**](SystemAdministrationConfigurationApi.md#CreatePolicyServiceActionStart) | **Post** /node/services/policy?action&#x3D;start | Restart, start or stop the service
[**CreatePolicyServiceActionStop**](SystemAdministrationConfigurationApi.md#CreatePolicyServiceActionStop) | **Post** /node/services/policy?action&#x3D;stop | Restart, start or stop the service
[**CreateProtonServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateProtonServiceActionRestart) | **Post** /node/services/manager?action&#x3D;restart | Restart, start or stop the service
[**CreateProtonServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateProtonServiceActionStart) | **Post** /node/services/manager?action&#x3D;start | Restart, start or stop the service
[**CreateProtonServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateProtonServiceActionStop) | **Post** /node/services/manager?action&#x3D;stop | Restart, start or stop the service
[**CreateProxyServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateProxyServiceActionRestart) | **Post** /node/services/http?action&#x3D;restart | Restart the http service
[**CreateProxyServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateProxyServiceActionStart) | **Post** /node/services/http?action&#x3D;start | Start the http service
[**CreateProxyServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateProxyServiceActionStop) | **Post** /node/services/http?action&#x3D;stop | Stop the http service
[**CreateProxyServiceApplyCertificateActionApplyCertificate**](SystemAdministrationConfigurationApi.md#CreateProxyServiceApplyCertificateActionApplyCertificate) | **Post** /node/services/http?action&#x3D;apply_certificate | Update http service certificate
[**CreateRabbitMQServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateRabbitMQServiceActionRestart) | **Post** /node/services/mgmt-plane-bus?action&#x3D;restart | Restart, start or stop the Rabbit MQ service
[**CreateRabbitMQServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateRabbitMQServiceActionStart) | **Post** /node/services/mgmt-plane-bus?action&#x3D;start | Restart, start or stop the Rabbit MQ service
[**CreateRabbitMQServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateRabbitMQServiceActionStop) | **Post** /node/services/mgmt-plane-bus?action&#x3D;stop | Restart, start or stop the Rabbit MQ service
[**CreateRepositoryServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateRepositoryServiceActionRestart) | **Post** /node/services/install-upgrade?action&#x3D;restart | Restart, start or stop the NSX install-upgrade service
[**CreateRepositoryServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateRepositoryServiceActionStart) | **Post** /node/services/install-upgrade?action&#x3D;start | Restart, start or stop the NSX install-upgrade service
[**CreateRepositoryServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateRepositoryServiceActionStop) | **Post** /node/services/install-upgrade?action&#x3D;stop | Restart, start or stop the NSX install-upgrade service
[**CreateSNMPServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateSNMPServiceActionRestart) | **Post** /node/services/snmp?action&#x3D;restart | Restart, start or stop the SNMP service
[**CreateSNMPServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateSNMPServiceActionStart) | **Post** /node/services/snmp?action&#x3D;start | Restart, start or stop the SNMP service
[**CreateSNMPServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateSNMPServiceActionStop) | **Post** /node/services/snmp?action&#x3D;stop | Restart, start or stop the SNMP service
[**CreateSSHServiceActionNotifyMpaRestart**](SystemAdministrationConfigurationApi.md#CreateSSHServiceActionNotifyMpaRestart) | **Post** /node/services/ssh/notify_mpa?action&#x3D;restart | Restart, start or stop the ssh service
[**CreateSSHServiceActionNotifyMpaStart**](SystemAdministrationConfigurationApi.md#CreateSSHServiceActionNotifyMpaStart) | **Post** /node/services/ssh/notify_mpa?action&#x3D;start | Restart, start or stop the ssh service
[**CreateSSHServiceActionNotifyMpaStop**](SystemAdministrationConfigurationApi.md#CreateSSHServiceActionNotifyMpaStop) | **Post** /node/services/ssh/notify_mpa?action&#x3D;stop | Restart, start or stop the ssh service
[**CreateSSHServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateSSHServiceActionRestart) | **Post** /node/services/ssh?action&#x3D;restart | Restart, start or stop the ssh service
[**CreateSSHServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateSSHServiceActionStart) | **Post** /node/services/ssh?action&#x3D;start | Restart, start or stop the ssh service
[**CreateSSHServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateSSHServiceActionStop) | **Post** /node/services/ssh?action&#x3D;stop | Restart, start or stop the ssh service
[**CreateSSHServiceRemoveHostFingerprintActionRemoveHostFingerprint**](SystemAdministrationConfigurationApi.md#CreateSSHServiceRemoveHostFingerprintActionRemoveHostFingerprint) | **Post** /node/services/ssh?action&#x3D;remove_host_fingerprint | Remove a host&#x27;s fingerprint from known hosts file
[**CreateSearchServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateSearchServiceActionRestart) | **Post** /node/services/search?action&#x3D;restart | Restart, start or stop the NSX Search service
[**CreateSearchServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateSearchServiceActionStart) | **Post** /node/services/search?action&#x3D;start | Restart, start or stop the NSX Search service
[**CreateSearchServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateSearchServiceActionStop) | **Post** /node/services/search?action&#x3D;stop | Restart, start or stop the NSX Search service
[**CreateSyslogServiceActionRestart**](SystemAdministrationConfigurationApi.md#CreateSyslogServiceActionRestart) | **Post** /node/services/syslog?action&#x3D;restart | Restart, start or stop the syslog service
[**CreateSyslogServiceActionStart**](SystemAdministrationConfigurationApi.md#CreateSyslogServiceActionStart) | **Post** /node/services/syslog?action&#x3D;start | Restart, start or stop the syslog service
[**CreateSyslogServiceActionStop**](SystemAdministrationConfigurationApi.md#CreateSyslogServiceActionStop) | **Post** /node/services/syslog?action&#x3D;stop | Restart, start or stop the syslog service
[**CreateTransportNodeCollection**](SystemAdministrationConfigurationApi.md#CreateTransportNodeCollection) | **Post** /transport-node-collections | Create transport node collection by attaching Transport Node Profile to cluster.
[**CreateTransportNodeForDiscoveredNodeCreateTransportNode**](SystemAdministrationConfigurationApi.md#CreateTransportNodeForDiscoveredNodeCreateTransportNode) | **Post** /fabric/discovered-nodes/{node-ext-id}?action&#x3D;create_transport_node | Created Transport Node for Discovered Node
[**CreateTransportNodeProfile**](SystemAdministrationConfigurationApi.md#CreateTransportNodeProfile) | **Post** /transport-node-profiles | Create a Transport Node Profile
[**CreateTransportNodeWithDeploymentInfo**](SystemAdministrationConfigurationApi.md#CreateTransportNodeWithDeploymentInfo) | **Post** /transport-nodes | Create a Transport Node
[**CreateTransportZone**](SystemAdministrationConfigurationApi.md#CreateTransportZone) | **Post** /transport-zones | Create a Transport Zone
[**CreateTransportZoneProfile**](SystemAdministrationConfigurationApi.md#CreateTransportZoneProfile) | **Post** /transportzone-profiles | Create a transport zone Profile
[**CreateVNIPool**](SystemAdministrationConfigurationApi.md#CreateVNIPool) | **Post** /pools/vni-pools | Create a new VNI Pool.
[**DELETERabbitMQManagementPort**](SystemAdministrationConfigurationApi.md#DELETERabbitMQManagementPort) | **Delete** /node/rabbitmq-management-port | Delete RabbitMQ management port
[**DeleteApplianceManagementTask**](SystemAdministrationConfigurationApi.md#DeleteApplianceManagementTask) | **Delete** /node/tasks/{task-id} | Delete task
[**DeleteAutoDeployedClusterNodeVMDelete**](SystemAdministrationConfigurationApi.md#DeleteAutoDeployedClusterNodeVMDelete) | **Post** /cluster/nodes/deployments/{node-id}?action&#x3D;delete | Attempt to delete an auto-deployed cluster node VM
[**DeleteAutoDeployedPaceClusterNodeVMDelete**](SystemAdministrationConfigurationApi.md#DeleteAutoDeployedPaceClusterNodeVMDelete) | **Post** /intelligence/nodes/deployments/{node-id}?action&#x3D;delete | Attempt to delete an auto-deployed cluster node VM
[**DeleteClusterNodeConfig**](SystemAdministrationConfigurationApi.md#DeleteClusterNodeConfig) | **Delete** /cluster/nodes/{node-id} | Remove a controller from the cluster
[**DeleteClusterProfile**](SystemAdministrationConfigurationApi.md#DeleteClusterProfile) | **Delete** /cluster-profiles/{cluster-profile-id} | Delete a cluster profile
[**DeleteComputeCollectionFabricTemplate**](SystemAdministrationConfigurationApi.md#DeleteComputeCollectionFabricTemplate) | **Delete** /fabric/compute-collection-fabric-templates/{fabric-template-id} | Deletes compute collection fabric template
[**DeleteComputeManager**](SystemAdministrationConfigurationApi.md#DeleteComputeManager) | **Delete** /fabric/compute-managers/{compute-manager-id} | Unregister a compute manager
[**DeleteDirectoryDomain**](SystemAdministrationConfigurationApi.md#DeleteDirectoryDomain) | **Delete** /directory/domains/{domain-id} | Delete a specific domain with given identifier
[**DeleteDirectoryLdapServer**](SystemAdministrationConfigurationApi.md#DeleteDirectoryLdapServer) | **Delete** /directory/domains/{domain-id}/ldap-servers/{server-id} | Delete a LDAP server for directory domain
[**DeleteEdgeCluster**](SystemAdministrationConfigurationApi.md#DeleteEdgeCluster) | **Delete** /edge-clusters/{edge-cluster-id} | Delete Edge Cluster
[**DeleteFailureDomain**](SystemAdministrationConfigurationApi.md#DeleteFailureDomain) | **Delete** /failure-domains/{failure-domain-id} | Delete Failure Domain
[**DeleteHostSwitchProfile**](SystemAdministrationConfigurationApi.md#DeleteHostSwitchProfile) | **Delete** /host-switch-profiles/{host-switch-profile-id} | Delete a Hostswitch Profile
[**DeleteIpBlock**](SystemAdministrationConfigurationApi.md#DeleteIpBlock) | **Delete** /pools/ip-blocks/{block-id} | Delete an IP Address Block
[**DeleteIpBlockSubnet**](SystemAdministrationConfigurationApi.md#DeleteIpBlockSubnet) | **Delete** /pools/ip-subnets/{subnet-id} | Delete subnet within an IP block
[**DeleteIpPool**](SystemAdministrationConfigurationApi.md#DeleteIpPool) | **Delete** /pools/ip-pools/{pool-id} | Delete an IP Pool
[**DeleteMPAConfiguration**](SystemAdministrationConfigurationApi.md#DeleteMPAConfiguration) | **Delete** /node/mpa-config | Delete MPA configuration for this node
[**DeleteManagementPlaneConfiguration**](SystemAdministrationConfigurationApi.md#DeleteManagementPlaneConfiguration) | **Delete** /node/management-plane | Delete management plane configuration for this node
[**DeleteNode**](SystemAdministrationConfigurationApi.md#DeleteNode) | **Delete** /fabric/nodes/{node-id} | Delete a Node
[**DeleteNodeNetworkRoute**](SystemAdministrationConfigurationApi.md#DeleteNodeNetworkRoute) | **Delete** /node/network/routes/{route-id} | Delete node network route
[**DeleteNodeSyslogExporter**](SystemAdministrationConfigurationApi.md#DeleteNodeSyslogExporter) | **Delete** /node/services/syslog/exporters/{exporter-name} | Delete node syslog exporter
[**DeleteTransportNodeCollection**](SystemAdministrationConfigurationApi.md#DeleteTransportNodeCollection) | **Delete** /transport-node-collections/{transport-node-collection-id} | Detach transport node profile from compute collection.
[**DeleteTransportNodeProfile**](SystemAdministrationConfigurationApi.md#DeleteTransportNodeProfile) | **Delete** /transport-node-profiles/{transport-node-profile-id} | Delete a Transport Node Profile
[**DeleteTransportNodeWithDeploymentInfo**](SystemAdministrationConfigurationApi.md#DeleteTransportNodeWithDeploymentInfo) | **Delete** /transport-nodes/{transport-node-id} | Delete a Transport Node
[**DeleteTransportZone**](SystemAdministrationConfigurationApi.md#DeleteTransportZone) | **Delete** /transport-zones/{zone-id} | Delete a Transport Zone
[**DeleteTransportZoneProfile**](SystemAdministrationConfigurationApi.md#DeleteTransportZoneProfile) | **Delete** /transportzone-profiles/{transportzone-profile-id} | Delete a transport zone Profile
[**DeleteVNIPool**](SystemAdministrationConfigurationApi.md#DeleteVNIPool) | **Delete** /pools/vni-pools/{pool-id} | Delete a VNI Pool
[**DetachClusterNodeRemoveNode**](SystemAdministrationConfigurationApi.md#DetachClusterNodeRemoveNode) | **Post** /cluster/{node-id}?action&#x3D;remove_node | Detach a node from the Cluster
[**DisableFlowCacheDisableFlowCache**](SystemAdministrationConfigurationApi.md#DisableFlowCacheDisableFlowCache) | **Post** /transport-nodes/{transport-node-id}?action&#x3D;disable_flow_cache | Disable flow cache for an edge transport node
[**EnableFlowCacheEnableFlowCache**](SystemAdministrationConfigurationApi.md#EnableFlowCacheEnableFlowCache) | **Post** /transport-nodes/{transport-node-id}?action&#x3D;enable_flow_cache | Enable flow cache for an edge transport node
[**GetAllTransportNodesStatus**](SystemAdministrationConfigurationApi.md#GetAllTransportNodesStatus) | **Get** /transport-nodes/status | Get high-level summary of all transport nodes. The service layer does not support source &#x3D; realtime or cached.
[**GetAllTransportZoneStatus**](SystemAdministrationConfigurationApi.md#GetAllTransportZoneStatus) | **Get** /transport-zones/status | Get high-level summary of a transport zone. The service layer does not support source &#x3D; realtime or cached.
[**GetApiServiceConfig**](SystemAdministrationConfigurationApi.md#GetApiServiceConfig) | **Get** /cluster/api-service | Read API service properties
[**GetBackupUiFramesInfo**](SystemAdministrationConfigurationApi.md#GetBackupUiFramesInfo) | **Get** /cluster/backups/ui_frames | Get backup frames for UI
[**GetBundleIds**](SystemAdministrationConfigurationApi.md#GetBundleIds) | **Get** /repository/bundles | Get list of bundle-ids which are available in repository or in-progress 
[**GetBundleUploadPermissions**](SystemAdministrationConfigurationApi.md#GetBundleUploadPermissions) | **Get** /repository/bundles/upload-allowed | Checks bundle upload permissions
[**GetBundleUploadStatus**](SystemAdministrationConfigurationApi.md#GetBundleUploadStatus) | **Get** /repository/bundles/{bundle-id}/upload-status | Get bundle upload status
[**GetCloudNativeServiceInstance**](SystemAdministrationConfigurationApi.md#GetCloudNativeServiceInstance) | **Get** /fabric/cloud-native-service-instances/{external-id} | Returns information about a particular cloud native service instance by external-id. 
[**GetClusterCertificateId**](SystemAdministrationConfigurationApi.md#GetClusterCertificateId) | **Get** /cluster/api-certificate | Read cluster certificate ID
[**GetClusterNodeConfig**](SystemAdministrationConfigurationApi.md#GetClusterNodeConfig) | **Get** /cluster/{node-id} | Read cluster node configuration
[**GetClusterProfile**](SystemAdministrationConfigurationApi.md#GetClusterProfile) | **Get** /cluster-profiles/{cluster-profile-id} | Get cluster profile by Id
[**GetClusterVirtualIp**](SystemAdministrationConfigurationApi.md#GetClusterVirtualIp) | **Get** /cluster/api-virtual-ip | Read cluster virtual IP address
[**GetComputeCollectionFabricTemplate**](SystemAdministrationConfigurationApi.md#GetComputeCollectionFabricTemplate) | **Get** /fabric/compute-collection-fabric-templates/{fabric-template-id} | Get compute collection fabric template by id
[**GetComputeManagerState**](SystemAdministrationConfigurationApi.md#GetComputeManagerState) | **Get** /fabric/compute-managers/{compute-manager-id}/state | Get the realized state of a compute manager
[**GetContainerApplication**](SystemAdministrationConfigurationApi.md#GetContainerApplication) | **Get** /fabric/container-applications/{container-application-id} | Return a Container Application within a container project
[**GetContainerApplicationInstance**](SystemAdministrationConfigurationApi.md#GetContainerApplicationInstance) | **Get** /fabric/container-application-instances/{container-application-instance-id} | Return a container application instance
[**GetContainerCluster**](SystemAdministrationConfigurationApi.md#GetContainerCluster) | **Get** /fabric/container-clusters/{container-cluster-id} | Return a container cluster
[**GetContainerClusterNode**](SystemAdministrationConfigurationApi.md#GetContainerClusterNode) | **Get** /fabric/container-cluster-nodes/{container-cluster-node-id} | Return a container cluster node
[**GetContainerIngressPolicy**](SystemAdministrationConfigurationApi.md#GetContainerIngressPolicy) | **Get** /fabric/container-ingress-policies/{ingress-policy-id} | Returns an ingress policy spec
[**GetContainerNetworkPolicy**](SystemAdministrationConfigurationApi.md#GetContainerNetworkPolicy) | **Get** /fabric/container-network-policies/{network-policy-id} | Return a network policy spec
[**GetContainerProject**](SystemAdministrationConfigurationApi.md#GetContainerProject) | **Get** /fabric/container-projects/{container-project-id} | Return a container project
[**GetControllerProfilerStatus**](SystemAdministrationConfigurationApi.md#GetControllerProfilerStatus) | **Get** /node/services/controller/profiler | Get the status (Enabled/Disabled) of controller profiler
[**GetCurrentBarrier**](SystemAdministrationConfigurationApi.md#GetCurrentBarrier) | **Get** /realization-state-barrier/current | Gets the current barrier number
[**GetDirectoryDomain**](SystemAdministrationConfigurationApi.md#GetDirectoryDomain) | **Get** /directory/domains/{domain-id} | Get a specific domain with given identifier
[**GetDirectoryDomainSyncStats**](SystemAdministrationConfigurationApi.md#GetDirectoryDomainSyncStats) | **Get** /directory/domains/{domain-id}/sync-stats | Get domain sync statistics for the given identifier
[**GetDirectoryLdapServer**](SystemAdministrationConfigurationApi.md#GetDirectoryLdapServer) | **Get** /directory/domains/{domain-id}/ldap-servers/{server-id} | Get a specific LDAP server for a given directory domain
[**GetEdgeClusterAllocationStatus**](SystemAdministrationConfigurationApi.md#GetEdgeClusterAllocationStatus) | **Get** /edge-clusters/{edge-cluster-id}/allocation-status | Get the Allocation details of an edge cluster
[**GetEdgeClusterInterSiteStatus**](SystemAdministrationConfigurationApi.md#GetEdgeClusterInterSiteStatus) | **Get** /edge-clusters/{edge-cluster-id}/inter-site/status | Get inter-site status of the edge cluster
[**GetEdgeClusterState**](SystemAdministrationConfigurationApi.md#GetEdgeClusterState) | **Get** /edge-clusters/{edge-cluster-id}/state | Get the Realized State of a Edge Cluster
[**GetEdgeClusterStatus**](SystemAdministrationConfigurationApi.md#GetEdgeClusterStatus) | **Get** /edge-clusters/{edge-cluster-id}/status | Get the status for the Edge cluster of the given id
[**GetFabricNodeModules**](SystemAdministrationConfigurationApi.md#GetFabricNodeModules) | **Get** /fabric/nodes/{node-id}/modules | Get the module details of a Fabric Node This api is deprecated, use Transport Node API GET /transport-nodes/&amp;lt;transportnode-id&amp;gt;/modules to get fabric node modules. 
[**GetFabricNodeModulesOfTransportNode**](SystemAdministrationConfigurationApi.md#GetFabricNodeModulesOfTransportNode) | **Get** /transport-nodes/{node-id}/modules | Get the module details of a transport node 
[**GetFabricNodeState**](SystemAdministrationConfigurationApi.md#GetFabricNodeState) | **Get** /fabric/nodes/{node-id}/state | Get the Realized State of a Fabric Node.
[**GetFailureDomain**](SystemAdministrationConfigurationApi.md#GetFailureDomain) | **Get** /failure-domains/{failure-domain-id} | Get a Failure Domain
[**GetGlobalConfigs**](SystemAdministrationConfigurationApi.md#GetGlobalConfigs) | **Get** /global-configs/{config-type} | Get global configs for a config type
[**GetHeatmapTransportZoneStatus**](SystemAdministrationConfigurationApi.md#GetHeatmapTransportZoneStatus) | **Get** /transport-zones/{zone-id}/status | Get high-level summary of a transport zone
[**GetHostNodeStatusOnComputeCollection**](SystemAdministrationConfigurationApi.md#GetHostNodeStatusOnComputeCollection) | **Get** /fabric/compute-collections/{cc-ext-id}/member-status | Get status of member host nodes of the compute-collection. Only nsx prepared host nodes in the specified compute-collection are included in the response. cc-ext-id should be of type VC_Cluster.
[**GetHostSwitchProfile**](SystemAdministrationConfigurationApi.md#GetHostSwitchProfile) | **Get** /host-switch-profiles/{host-switch-profile-id} | Get a Hostswitch Profile by ID
[**GetInterSiteEdgeNodeBgpNeighborAdvertisedRoutes**](SystemAdministrationConfigurationApi.md#GetInterSiteEdgeNodeBgpNeighborAdvertisedRoutes) | **Get** /transport-nodes/{edge-node-id}/inter-site/bgp/neighbors/{neighbor-id}/advertised-routes | Get BGP neighbor advertised routes on edge transport node
[**GetInterSiteEdgeNodeBgpNeighborRoutes**](SystemAdministrationConfigurationApi.md#GetInterSiteEdgeNodeBgpNeighborRoutes) | **Get** /transport-nodes/{edge-node-id}/inter-site/bgp/neighbors/{neighbor-id}/routes | Get BGP neighbor learned routes on edge transport node
[**GetInterSiteEdgeNodeBgpSummary**](SystemAdministrationConfigurationApi.md#GetInterSiteEdgeNodeBgpSummary) | **Get** /transport-nodes/{edge-node-id}/inter-site/bgp/summary | Get inter-site BGP summary of edge node
[**GetInterSiteEdgeNodeStatistics**](SystemAdministrationConfigurationApi.md#GetInterSiteEdgeNodeStatistics) | **Get** /transport-nodes/{edge-node-id}/inter-site/statistics | Get inter-site statistics of edge node
[**GetInventoryConfig**](SystemAdministrationConfigurationApi.md#GetInventoryConfig) | **Get** /configs/inventory | Return inventory configuration
[**GetNodeMandatoryAccessControl**](SystemAdministrationConfigurationApi.md#GetNodeMandatoryAccessControl) | **Get** /node/hardening-policy/mandatory-access-control | Gets the enable status for Mandatory Access Control
[**GetNodeMandatoryAccessControlReport**](SystemAdministrationConfigurationApi.md#GetNodeMandatoryAccessControlReport) | **Get** /node/hardening-policy/mandatory-access-control/report | Get the report for Mandatory Access Control
[**GetNodeMode**](SystemAdministrationConfigurationApi.md#GetNodeMode) | **Get** /node/mode | NodeMode
[**GetOvfDeployInfo**](SystemAdministrationConfigurationApi.md#GetOvfDeployInfo) | **Get** /repository/bundles/ovf-deploy-info | Get information of the OVF which will be getting deployed. 
[**GetPaceHostConfiguration**](SystemAdministrationConfigurationApi.md#GetPaceHostConfiguration) | **Get** /intelligence/host-config | Get NSX-Intelligence host configuration
[**GetPhysicalServer**](SystemAdministrationConfigurationApi.md#GetPhysicalServer) | **Get** /fabric/physical-servers/{physical-server-id} | Return a specific physical server
[**GetPnicStatusesForTransportNode**](SystemAdministrationConfigurationApi.md#GetPnicStatusesForTransportNode) | **Get** /transport-nodes/{node-id}/pnic-bond-status | Get high-level summary of a transport node
[**GetRealizationStateBarrierConfig**](SystemAdministrationConfigurationApi.md#GetRealizationStateBarrierConfig) | **Get** /realization-state-barrier/config | Gets the realization state barrier configuration
[**GetRepoSyncStatus**](SystemAdministrationConfigurationApi.md#GetRepoSyncStatus) | **Get** /cluster/nodes/{node-id}/repo_sync/status | Synchronizes the repository data between nsx managers.
[**GetSupportedHostOSTypes**](SystemAdministrationConfigurationApi.md#GetSupportedHostOSTypes) | **Get** /fabric/ostypes | Return list of supported host OS types
[**GetTransportNodeCollection**](SystemAdministrationConfigurationApi.md#GetTransportNodeCollection) | **Get** /transport-node-collections/{transport-node-collection-id} | Get Transport Node collection by id
[**GetTransportNodeCollectionState**](SystemAdministrationConfigurationApi.md#GetTransportNodeCollectionState) | **Get** /transport-node-collections/{transport-node-collection-id}/state | Get Transport Node collection application state
[**GetTransportNodeProfile**](SystemAdministrationConfigurationApi.md#GetTransportNodeProfile) | **Get** /transport-node-profiles/{transport-node-profile-id} | Get a Transport Node
[**GetTransportNodeReport**](SystemAdministrationConfigurationApi.md#GetTransportNodeReport) | **Get** /transport-zones/transport-node-status-report | Creates a status report of transport nodes of all the transport zones
[**GetTransportNodeReportForATransportZone**](SystemAdministrationConfigurationApi.md#GetTransportNodeReportForATransportZone) | **Get** /transport-zones/{zone-id}/transport-node-status-report | Creates a status report of transport nodes in a transport zone
[**GetTransportNodeStateWithDeploymentInfo**](SystemAdministrationConfigurationApi.md#GetTransportNodeStateWithDeploymentInfo) | **Get** /transport-nodes/{transport-node-id}/state | Get a Transport Node&#x27;s State
[**GetTransportNodeStatus**](SystemAdministrationConfigurationApi.md#GetTransportNodeStatus) | **Get** /transport-nodes/{node-id}/status | Read status of a transport node
[**GetTransportNodeWithDeploymentInfo**](SystemAdministrationConfigurationApi.md#GetTransportNodeWithDeploymentInfo) | **Get** /transport-nodes/{transport-node-id} | Get a Transport Node
[**GetTransportZone**](SystemAdministrationConfigurationApi.md#GetTransportZone) | **Get** /transport-zones/{zone-id} | Get a Transport Zone
[**GetTransportZoneProfile**](SystemAdministrationConfigurationApi.md#GetTransportZoneProfile) | **Get** /transportzone-profiles/{transportzone-profile-id} | Get transport zone profile by identifier
[**GetTransportZoneStatus**](SystemAdministrationConfigurationApi.md#GetTransportZoneStatus) | **Get** /transport-zones/{zone-id}/summary | Get a Transport Zone&#x27;s Current Runtime Status Information
[**GetTunnel**](SystemAdministrationConfigurationApi.md#GetTunnel) | **Get** /transport-nodes/{node-id}/tunnels/{tunnel-name} | Tunnel properties
[**HostPrepDiscoveredNodeHostprep**](SystemAdministrationConfigurationApi.md#HostPrepDiscoveredNodeHostprep) | **Post** /fabric/discovered-nodes/{node-ext-id}?action&#x3D;hostprep | (Deprecated) Prepares discovered Node for NSX
[**IncrementRealizationStateBarrierIncrement**](SystemAdministrationConfigurationApi.md#IncrementRealizationStateBarrierIncrement) | **Post** /realization-state-barrier/current?action&#x3D;increment | Increments the barrier count by 1
[**InvokeDeleteClusterCentralAPI**](SystemAdministrationConfigurationApi.md#InvokeDeleteClusterCentralAPI) | **Delete** /cluster/{target-node-id}/{target-uri} | Invoke DELETE request on target cluster node
[**InvokeDeleteFabricCentralAPI**](SystemAdministrationConfigurationApi.md#InvokeDeleteFabricCentralAPI) | **Delete** /fabric/nodes/{target-node-id}/{target-uri} | Invoke DELETE request on target fabric node
[**InvokeDeleteTransportNodeCentralAPI**](SystemAdministrationConfigurationApi.md#InvokeDeleteTransportNodeCentralAPI) | **Delete** /transport-nodes/{target-node-id}/{target-uri} | Invoke DELETE request on target transport node
[**InvokeGetClusterCentralAPI**](SystemAdministrationConfigurationApi.md#InvokeGetClusterCentralAPI) | **Get** /cluster/{target-node-id}/{target-uri} | Invoke GET request on target cluster node
[**InvokeGetFabricCentralAPI**](SystemAdministrationConfigurationApi.md#InvokeGetFabricCentralAPI) | **Get** /fabric/nodes/{target-node-id}/{target-uri} | Invoke GET request on target fabric node
[**InvokeGetTransportNodeCentralAPI**](SystemAdministrationConfigurationApi.md#InvokeGetTransportNodeCentralAPI) | **Get** /transport-nodes/{target-node-id}/{target-uri} | Invoke GET request on target transport node
[**InvokePostClusterCentralAPI**](SystemAdministrationConfigurationApi.md#InvokePostClusterCentralAPI) | **Post** /cluster/{target-node-id}/{target-uri} | Invoke POST request on target cluster node
[**InvokePostFabricCentralAPI**](SystemAdministrationConfigurationApi.md#InvokePostFabricCentralAPI) | **Post** /fabric/nodes/{target-node-id}/{target-uri} | Invoke POST request on target fabric node
[**InvokePostTransportNodeCentralAPI**](SystemAdministrationConfigurationApi.md#InvokePostTransportNodeCentralAPI) | **Post** /transport-nodes/{target-node-id}/{target-uri} | Invoke POST request on target transport node
[**InvokePutClusterCentralAPI**](SystemAdministrationConfigurationApi.md#InvokePutClusterCentralAPI) | **Put** /cluster/{target-node-id}/{target-uri} | Invoke PUT request on target cluster node
[**InvokePutFabricCentralAPI**](SystemAdministrationConfigurationApi.md#InvokePutFabricCentralAPI) | **Put** /fabric/nodes/{target-node-id}/{target-uri} | Invoke PUT request on target fabric node
[**InvokePutTransportNodeCentralAPI**](SystemAdministrationConfigurationApi.md#InvokePutTransportNodeCentralAPI) | **Put** /transport-nodes/{target-node-id}/{target-uri} | Invoke PUT request on target transport node
[**JoinClusterJoinCluster**](SystemAdministrationConfigurationApi.md#JoinClusterJoinCluster) | **Post** /cluster?action&#x3D;join_cluster | Join this node to a NSX Cluster
[**ListAllCloudNativeServiceInstances**](SystemAdministrationConfigurationApi.md#ListAllCloudNativeServiceInstances) | **Get** /fabric/cloud-native-service-instances | Returns the List of cloud native service instances
[**ListApplianceManagementTasks**](SystemAdministrationConfigurationApi.md#ListApplianceManagementTasks) | **Get** /node/tasks | List appliance management tasks
[**ListCentralNodeConfigProfiles**](SystemAdministrationConfigurationApi.md#ListCentralNodeConfigProfiles) | **Get** /configs/central-config/node-config-profiles/ | List all Central Node Config profiles
[**ListClusterNodeConfigs**](SystemAdministrationConfigurationApi.md#ListClusterNodeConfigs) | **Get** /cluster/nodes | List Cluster Node Configurations
[**ListClusterNodeInterfaces**](SystemAdministrationConfigurationApi.md#ListClusterNodeInterfaces) | **Get** /cluster/nodes/{node-id}/network/interfaces | List the specified node&#x27;s Network Interfaces
[**ListClusterNodeVMDeploymentRequests**](SystemAdministrationConfigurationApi.md#ListClusterNodeVMDeploymentRequests) | **Get** /cluster/nodes/deployments | Returns info for all cluster node VM auto-deployment attempts
[**ListClusterProfiles**](SystemAdministrationConfigurationApi.md#ListClusterProfiles) | **Get** /cluster-profiles | List Cluster Profiles
[**ListComputeCollectionFabricTemplates**](SystemAdministrationConfigurationApi.md#ListComputeCollectionFabricTemplates) | **Get** /fabric/compute-collection-fabric-templates | Get compute collection fabric templates
[**ListComputeCollectionPhysicalNetworkInterfaces**](SystemAdministrationConfigurationApi.md#ListComputeCollectionPhysicalNetworkInterfaces) | **Get** /fabric/compute-collections/{cc-ext-id}/network/physical-interfaces | List the Physical Network Interface for all discovered nodes
[**ListComputeCollections**](SystemAdministrationConfigurationApi.md#ListComputeCollections) | **Get** /fabric/compute-collections | Return the List of Compute Collections
[**ListComputeManagers**](SystemAdministrationConfigurationApi.md#ListComputeManagers) | **Get** /fabric/compute-managers | Return the List of Compute managers
[**ListContainerApplicationInstances**](SystemAdministrationConfigurationApi.md#ListContainerApplicationInstances) | **Get** /fabric/container-application-instances | Return the list of container application instance
[**ListContainerApplications**](SystemAdministrationConfigurationApi.md#ListContainerApplications) | **Get** /fabric/container-applications | Return the List of Container Applications
[**ListContainerClusterNodes**](SystemAdministrationConfigurationApi.md#ListContainerClusterNodes) | **Get** /fabric/container-cluster-nodes | Return the list of container cluster nodes
[**ListContainerClusters**](SystemAdministrationConfigurationApi.md#ListContainerClusters) | **Get** /fabric/container-clusters | Return the List of Container Clusters
[**ListContainerIngressPolicies**](SystemAdministrationConfigurationApi.md#ListContainerIngressPolicies) | **Get** /fabric/container-ingress-policies | Return the List of Container Ingress Policies
[**ListContainerNetworkPolicies**](SystemAdministrationConfigurationApi.md#ListContainerNetworkPolicies) | **Get** /fabric/container-network-policies | Return the List of Container Network Policies
[**ListContainerProjects**](SystemAdministrationConfigurationApi.md#ListContainerProjects) | **Get** /fabric/container-projects | Return the list of container projects
[**ListDirectoryDomains**](SystemAdministrationConfigurationApi.md#ListDirectoryDomains) | **Get** /directory/domains | List all configured domains
[**ListDirectoryGroupMemberGroups**](SystemAdministrationConfigurationApi.md#ListDirectoryGroupMemberGroups) | **Get** /directory/domains/{domain-id}/groups/{group-id}/member-groups | List members of a directory group
[**ListDirectoryLdapServers**](SystemAdministrationConfigurationApi.md#ListDirectoryLdapServers) | **Get** /directory/domains/{domain-id}/ldap-servers | List all configured domain LDAP servers
[**ListDiscoveredNodes**](SystemAdministrationConfigurationApi.md#ListDiscoveredNodes) | **Get** /fabric/discovered-nodes | Return the List of Discovered Nodes
[**ListEdgeClusters**](SystemAdministrationConfigurationApi.md#ListEdgeClusters) | **Get** /edge-clusters | List Edge Clusters
[**ListFabricNodeInterfaces**](SystemAdministrationConfigurationApi.md#ListFabricNodeInterfaces) | **Get** /fabric/nodes/{node-id}/network/interfaces | List the specified node&#x27;s Network Interfaces
[**ListFabricNodeNeighborProperties**](SystemAdministrationConfigurationApi.md#ListFabricNodeNeighborProperties) | **Get** /lldp/fabric-nodes/{fabric-node-id}/interfaces | List LLDP Neighbor Properties of Fabric Node
[**ListFailureDomains**](SystemAdministrationConfigurationApi.md#ListFailureDomains) | **Get** /failure-domains | List Failure Domains
[**ListGlobalConfigs**](SystemAdministrationConfigurationApi.md#ListGlobalConfigs) | **Get** /global-configs | List global configurations of a NSX domain
[**ListHostSwitchProfiles**](SystemAdministrationConfigurationApi.md#ListHostSwitchProfiles) | **Get** /host-switch-profiles | List Hostswitch Profiles
[**ListIntelligenceFormFactors**](SystemAdministrationConfigurationApi.md#ListIntelligenceFormFactors) | **Get** /node/intelligence/form-factors | List available NSX Intelligence appliance form factors
[**ListInterSiteEdgeNodeBgpNeighbors**](SystemAdministrationConfigurationApi.md#ListInterSiteEdgeNodeBgpNeighbors) | **Get** /transport-nodes/{edge-node-id}/inter-site/bgp/neighbors | Paginated list of BGP Neighbors on edge transport node
[**ListIpBlockSubnets**](SystemAdministrationConfigurationApi.md#ListIpBlockSubnets) | **Get** /pools/ip-subnets | List subnets within an IP block
[**ListIpBlocks**](SystemAdministrationConfigurationApi.md#ListIpBlocks) | **Get** /pools/ip-blocks | Returns list of configured IP address blocks.
[**ListIpPoolAllocations**](SystemAdministrationConfigurationApi.md#ListIpPoolAllocations) | **Get** /pools/ip-pools/{pool-id}/allocations | List IP Pool Allocations
[**ListIpPools**](SystemAdministrationConfigurationApi.md#ListIpPools) | **Get** /pools/ip-pools | List IP Pools
[**ListMacPools**](SystemAdministrationConfigurationApi.md#ListMacPools) | **Get** /pools/mac-pools | List MAC Pools
[**ListNeighborProperties**](SystemAdministrationConfigurationApi.md#ListNeighborProperties) | **Get** /lldp/transport-nodes/{node-id}/interfaces | List LLDP Neighbor Properties of Transport Node
[**ListNodeCapabilities**](SystemAdministrationConfigurationApi.md#ListNodeCapabilities) | **Get** /fabric/nodes/{node-id}/capabilities | Return the List of Capabilities of a Single Node
[**ListNodeInterfaces**](SystemAdministrationConfigurationApi.md#ListNodeInterfaces) | **Get** /node/network/interfaces | List the NSX Manager&#x27;s Network Interfaces
[**ListNodeNetworkRoutes**](SystemAdministrationConfigurationApi.md#ListNodeNetworkRoutes) | **Get** /node/network/routes | List node network routes
[**ListNodeProcesses**](SystemAdministrationConfigurationApi.md#ListNodeProcesses) | **Get** /node/processes | List node processes
[**ListNodeServices**](SystemAdministrationConfigurationApi.md#ListNodeServices) | **Get** /node/services | List node services
[**ListNodeSyslogExporters**](SystemAdministrationConfigurationApi.md#ListNodeSyslogExporters) | **Get** /node/services/syslog/exporters | List node syslog exporters
[**ListNodes**](SystemAdministrationConfigurationApi.md#ListNodes) | **Get** /fabric/nodes | Return the List of Nodes
[**ListPaceClusterNodeVMDeploymentRequests**](SystemAdministrationConfigurationApi.md#ListPaceClusterNodeVMDeploymentRequests) | **Get** /intelligence/nodes/deployments | Returns info for all cluster node VM auto-deployment attempts
[**ListPhysicalServers**](SystemAdministrationConfigurationApi.md#ListPhysicalServers) | **Get** /fabric/physical-servers | Return the list of physical servers
[**ListRemoteTransportNodeStatus**](SystemAdministrationConfigurationApi.md#ListRemoteTransportNodeStatus) | **Get** /transport-nodes/{node-id}/remote-transport-node-status | Read status of all transport nodes with tunnel connections to transport node 
[**ListTasks**](SystemAdministrationConfigurationApi.md#ListTasks) | **Get** /tasks | Get information about all tasks
[**ListTransportNodeCapabilities**](SystemAdministrationConfigurationApi.md#ListTransportNodeCapabilities) | **Get** /transport-nodes/{transport-node-id}/capabilities | Return the list of capabilities of transport node
[**ListTransportNodeCollections**](SystemAdministrationConfigurationApi.md#ListTransportNodeCollections) | **Get** /transport-node-collections | List Transport Node collections
[**ListTransportNodeInterfaces**](SystemAdministrationConfigurationApi.md#ListTransportNodeInterfaces) | **Get** /transport-nodes/{transport-node-id}/network/interfaces | List the specified transport node&#x27;s network interfaces
[**ListTransportNodeProfiles**](SystemAdministrationConfigurationApi.md#ListTransportNodeProfiles) | **Get** /transport-node-profiles | List Transport Nodes
[**ListTransportNodeStatus**](SystemAdministrationConfigurationApi.md#ListTransportNodeStatus) | **Get** /transport-zones/transport-node-status | Read status of all the transport nodes
[**ListTransportNodeStatusForTransportZone**](SystemAdministrationConfigurationApi.md#ListTransportNodeStatusForTransportZone) | **Get** /transport-zones/{zone-id}/transport-node-status | Read status of transport nodes in a transport zone
[**ListTransportNodesByStateWithDeploymentInfo**](SystemAdministrationConfigurationApi.md#ListTransportNodesByStateWithDeploymentInfo) | **Get** /transport-nodes/state | List transport nodes by realized state
[**ListTransportNodesWithDeploymentInfo**](SystemAdministrationConfigurationApi.md#ListTransportNodesWithDeploymentInfo) | **Get** /transport-nodes | List Transport Nodes
[**ListTransportZoneProfiles**](SystemAdministrationConfigurationApi.md#ListTransportZoneProfiles) | **Get** /transportzone-profiles | List transport zone profiles
[**ListTransportZones**](SystemAdministrationConfigurationApi.md#ListTransportZones) | **Get** /transport-zones | List Transport Zones
[**ListVNIPools**](SystemAdministrationConfigurationApi.md#ListVNIPools) | **Get** /pools/vni-pools | List VNI Pools
[**ListVifs**](SystemAdministrationConfigurationApi.md#ListVifs) | **Get** /fabric/vifs | Return the List of Virtual Network Interfaces (VIFs)
[**ListVirtualMachines**](SystemAdministrationConfigurationApi.md#ListVirtualMachines) | **Get** /fabric/virtual-machines | Return the List of Virtual Machines
[**ListVirtualSwitches**](SystemAdministrationConfigurationApi.md#ListVirtualSwitches) | **Get** /fabric/virtual-switches | Return the List of Virtual Switches
[**ListVmToolsInfo**](SystemAdministrationConfigurationApi.md#ListVmToolsInfo) | **Get** /fabric/virtual-machines/tools-info | Return the list of tools and agents installed in VMs.
[**ListVtepLabelPools**](SystemAdministrationConfigurationApi.md#ListVtepLabelPools) | **Get** /pools/vtep-label-pools | List virtual tunnel endpoint Label Pools
[**PatchPaceHostConfiguration**](SystemAdministrationConfigurationApi.md#PatchPaceHostConfiguration) | **Patch** /intelligence/host-config | Patch NSX-Intelligence host configuration
[**PerformActionOnComputeCollection**](SystemAdministrationConfigurationApi.md#PerformActionOnComputeCollection) | **Post** /fabric/compute-collections/{cc-ext-id} | Perform action specific to NSX on the compute-collection. cc-ext-id should be of type VC_Cluster.
[**PerformHostNodeUpgradeActionUpgradeInfra**](SystemAdministrationConfigurationApi.md#PerformHostNodeUpgradeActionUpgradeInfra) | **Post** /fabric/nodes/{node-id}?action&#x3D;upgrade_infra | Perform a service deployment upgrade on a host node
[**PerformNodeAction**](SystemAdministrationConfigurationApi.md#PerformNodeAction) | **Post** /fabric/nodes/{node-id} | Perform an Action on Fabric Node
[**PerformRepoSyncRepoSync**](SystemAdministrationConfigurationApi.md#PerformRepoSyncRepoSync) | **Post** /cluster/node?action&#x3D;repo_sync | Synchronizes the repository data between nsx managers.
[**PostNodeSyslogExporter**](SystemAdministrationConfigurationApi.md#PostNodeSyslogExporter) | **Post** /node/services/syslog/exporters | Add node syslog exporter
[**QueryTunnels**](SystemAdministrationConfigurationApi.md#QueryTunnels) | **Get** /transport-nodes/{node-id}/tunnels | List of tunnels
[**ReadApplianceManagementService**](SystemAdministrationConfigurationApi.md#ReadApplianceManagementService) | **Get** /node/services/node-mgmt | Read appliance management service properties
[**ReadApplianceManagementServiceStatus**](SystemAdministrationConfigurationApi.md#ReadApplianceManagementServiceStatus) | **Get** /node/services/node-mgmt/status | Read appliance management service status
[**ReadApplianceManagementTaskProperties**](SystemAdministrationConfigurationApi.md#ReadApplianceManagementTaskProperties) | **Get** /node/tasks/{task-id} | Read task properties
[**ReadApplianceNodeStatus**](SystemAdministrationConfigurationApi.md#ReadApplianceNodeStatus) | **Get** /node/status | Read node status
[**ReadAsyncApplianceManagementTaskResponse**](SystemAdministrationConfigurationApi.md#ReadAsyncApplianceManagementTaskResponse) | **Get** /node/tasks/{task-id}/response | Read asynchronous task response
[**ReadAsyncReplicatorService**](SystemAdministrationConfigurationApi.md#ReadAsyncReplicatorService) | **Get** /node/services/async_replicator | Read the Async Replicator service properties
[**ReadAsyncReplicatorServiceStatus**](SystemAdministrationConfigurationApi.md#ReadAsyncReplicatorServiceStatus) | **Get** /node/services/async_replicator/status | Read the Async Replicator service status
[**ReadCentralConfigProperties**](SystemAdministrationConfigurationApi.md#ReadCentralConfigProperties) | **Get** /node/central-config | Read Central Config properties
[**ReadCentralNodeConfigProfile**](SystemAdministrationConfigurationApi.md#ReadCentralNodeConfigProfile) | **Get** /configs/central-config/node-config-profiles/{profile-id} | Read Central Node Config profile
[**ReadClusterBootManagerService**](SystemAdministrationConfigurationApi.md#ReadClusterBootManagerService) | **Get** /node/services/cluster_manager | Read cluster boot manager service properties
[**ReadClusterBootManagerServiceStatus**](SystemAdministrationConfigurationApi.md#ReadClusterBootManagerServiceStatus) | **Get** /node/services/cluster_manager/status | Read cluster boot manager service status
[**ReadClusterConfig**](SystemAdministrationConfigurationApi.md#ReadClusterConfig) | **Get** /cluster | Read Cluster Configuration
[**ReadClusterNodeConfig**](SystemAdministrationConfigurationApi.md#ReadClusterNodeConfig) | **Get** /cluster/nodes/{node-id} | Read Cluster Node Configuration
[**ReadClusterNodeInterface**](SystemAdministrationConfigurationApi.md#ReadClusterNodeInterface) | **Get** /cluster/nodes/{node-id}/network/interfaces/{interface-id} | Read the node&#x27;s Network Interface
[**ReadClusterNodeInterfaceStatistics**](SystemAdministrationConfigurationApi.md#ReadClusterNodeInterfaceStatistics) | **Get** /cluster/nodes/{node-id}/network/interfaces/{interface-id}/stats | Read the NSX Manager/Controller&#x27;s Network Interface Statistics
[**ReadClusterNodeStatus**](SystemAdministrationConfigurationApi.md#ReadClusterNodeStatus) | **Get** /cluster/nodes/{node-id}/status | Read cluster node runtime status
[**ReadClusterNodeVMDeploymentRequest**](SystemAdministrationConfigurationApi.md#ReadClusterNodeVMDeploymentRequest) | **Get** /cluster/nodes/deployments/{node-id} | Returns info for a cluster-node VM auto-deployment attempt
[**ReadClusterNodeVMDeploymentStatus**](SystemAdministrationConfigurationApi.md#ReadClusterNodeVMDeploymentStatus) | **Get** /cluster/nodes/deployments/{node-id}/status | Returns the status of the VM creation/deletion
[**ReadClusterNodesAggregateStatus**](SystemAdministrationConfigurationApi.md#ReadClusterNodesAggregateStatus) | **Get** /cluster/nodes/status | Read cluster runtime status
[**ReadClusterStatus**](SystemAdministrationConfigurationApi.md#ReadClusterStatus) | **Get** /cluster/status | Read Cluster Status
[**ReadCminventoryService**](SystemAdministrationConfigurationApi.md#ReadCminventoryService) | **Get** /node/services/cm-inventory | Read cm inventory service properties
[**ReadCminventoryServiceStatus**](SystemAdministrationConfigurationApi.md#ReadCminventoryServiceStatus) | **Get** /node/services/cm-inventory/status | Read manager service status
[**ReadComputeCollection**](SystemAdministrationConfigurationApi.md#ReadComputeCollection) | **Get** /fabric/compute-collections/{cc-ext-id} | Return Compute Collection Information
[**ReadComputeManager**](SystemAdministrationConfigurationApi.md#ReadComputeManager) | **Get** /fabric/compute-managers/{compute-manager-id} | Return compute manager Information
[**ReadComputeManagerStatus**](SystemAdministrationConfigurationApi.md#ReadComputeManagerStatus) | **Get** /fabric/compute-managers/{compute-manager-id}/status | Return runtime status information for a compute manager
[**ReadControllerServerCertificate**](SystemAdministrationConfigurationApi.md#ReadControllerServerCertificate) | **Get** /node/services/controller/controller-certificate | Read controller server certificate properties
[**ReadControllerServerService**](SystemAdministrationConfigurationApi.md#ReadControllerServerService) | **Get** /node/services/controller | Read controller service properties
[**ReadControllerServerServiceStatus**](SystemAdministrationConfigurationApi.md#ReadControllerServerServiceStatus) | **Get** /node/services/controller/status | Read controller service status
[**ReadDiscoveredNode**](SystemAdministrationConfigurationApi.md#ReadDiscoveredNode) | **Get** /fabric/discovered-nodes/{node-ext-id} | Return Discovered Node Information
[**ReadEdgeCluster**](SystemAdministrationConfigurationApi.md#ReadEdgeCluster) | **Get** /edge-clusters/{edge-cluster-id} | Read Edge Cluster
[**ReadFabricNodeInterface**](SystemAdministrationConfigurationApi.md#ReadFabricNodeInterface) | **Get** /fabric/nodes/{node-id}/network/interfaces/{interface-id} | Read the node&#x27;s Network Interface
[**ReadFabricNodeInterfaceStatistics**](SystemAdministrationConfigurationApi.md#ReadFabricNodeInterfaceStatistics) | **Get** /fabric/nodes/{node-id}/network/interfaces/{interface-id}/stats | Read the NSX Manager&#x27;s Network Interface Statistics
[**ReadFabricNodeNeighborProperties**](SystemAdministrationConfigurationApi.md#ReadFabricNodeNeighborProperties) | **Get** /lldp/fabric-nodes/{fabric-node-id}/interfaces/{interface-name} | Read LLDP Neighbor Properties of Fabric Node by Interface Name
[**ReadIdpsReportingService**](SystemAdministrationConfigurationApi.md#ReadIdpsReportingService) | **Get** /node/services/idps-reporting | Read the idps-reporting service properties
[**ReadIdpsReportingServiceStatus**](SystemAdministrationConfigurationApi.md#ReadIdpsReportingServiceStatus) | **Get** /node/services/idps-reporting/status | Read the idps-reporting service status
[**ReadIntelligenceUpgradeCoordinatorService**](SystemAdministrationConfigurationApi.md#ReadIntelligenceUpgradeCoordinatorService) | **Get** /node/services/intelligence-upgrade-coordinator | Read intelligence upgrade coordinator service properties
[**ReadIntelligenceUpgradeCoordinatorServiceStatus**](SystemAdministrationConfigurationApi.md#ReadIntelligenceUpgradeCoordinatorServiceStatus) | **Get** /node/services/intelligence-upgrade-coordinator/status | Read intelligence upgrade coordinator service status
[**ReadIpBlock**](SystemAdministrationConfigurationApi.md#ReadIpBlock) | **Get** /pools/ip-blocks/{block-id} | Get IP address block information.
[**ReadIpBlockSubnet**](SystemAdministrationConfigurationApi.md#ReadIpBlockSubnet) | **Get** /pools/ip-subnets/{subnet-id} | Get the subnet within an IP block
[**ReadIpPool**](SystemAdministrationConfigurationApi.md#ReadIpPool) | **Get** /pools/ip-pools/{pool-id} | Read IP Pool
[**ReadLiagentService**](SystemAdministrationConfigurationApi.md#ReadLiagentService) | **Get** /node/services/liagent | Read liagent service properties
[**ReadLiagentServiceStatus**](SystemAdministrationConfigurationApi.md#ReadLiagentServiceStatus) | **Get** /node/services/liagent/status | Read liagent service status
[**ReadMPAConfiguration**](SystemAdministrationConfigurationApi.md#ReadMPAConfiguration) | **Get** /node/mpa-config | Get MPA configuration for this node
[**ReadMacPool**](SystemAdministrationConfigurationApi.md#ReadMacPool) | **Get** /pools/mac-pools/{pool-id} | Read MAC Pool
[**ReadManagementConfig**](SystemAdministrationConfigurationApi.md#ReadManagementConfig) | **Get** /configs/management | Read NSX Management nodes global configuration.
[**ReadManagementPlaneConfiguration**](SystemAdministrationConfigurationApi.md#ReadManagementPlaneConfiguration) | **Get** /node/management-plane | Get management plane configuration for this node
[**ReadMigrationCoordinatorService**](SystemAdministrationConfigurationApi.md#ReadMigrationCoordinatorService) | **Get** /node/services/migration-coordinator | Read migration coordinator service properties
[**ReadMigrationCoordinatorServiceStatus**](SystemAdministrationConfigurationApi.md#ReadMigrationCoordinatorServiceStatus) | **Get** /node/services/migration-coordinator/status | Read migration coordinator service status
[**ReadNSXMessageBusService**](SystemAdministrationConfigurationApi.md#ReadNSXMessageBusService) | **Get** /node/services/nsx-message-bus | Read NSX Message Bus service properties
[**ReadNSXMessageBusServiceStatus**](SystemAdministrationConfigurationApi.md#ReadNSXMessageBusServiceStatus) | **Get** /node/services/nsx-message-bus/status | Read NSX Message Bus service status
[**ReadNTPService**](SystemAdministrationConfigurationApi.md#ReadNTPService) | **Get** /node/services/ntp | Read NTP service properties
[**ReadNTPServiceStatus**](SystemAdministrationConfigurationApi.md#ReadNTPServiceStatus) | **Get** /node/services/ntp/status | Read NTP service status
[**ReadNeighborProperties**](SystemAdministrationConfigurationApi.md#ReadNeighborProperties) | **Get** /lldp/transport-nodes/{node-id}/interfaces/{interface-name} | Read LLDP Neighbor Properties of Transport Node by Interface Name 
[**ReadNetworkInterfaceStatistics**](SystemAdministrationConfigurationApi.md#ReadNetworkInterfaceStatistics) | **Get** /node/network/interfaces/{interface-id}/stats | Read the NSX Manager&#x27;s Network Interface Statistics
[**ReadNetworkProperties**](SystemAdministrationConfigurationApi.md#ReadNetworkProperties) | **Get** /node/network | Read network configuration properties
[**ReadNode**](SystemAdministrationConfigurationApi.md#ReadNode) | **Get** /fabric/nodes/{node-id} | Return Node Information
[**ReadNodeInterface**](SystemAdministrationConfigurationApi.md#ReadNodeInterface) | **Get** /node/network/interfaces/{interface-id} | Read the NSX Manager&#x27;s Network Interface
[**ReadNodeNameServers**](SystemAdministrationConfigurationApi.md#ReadNodeNameServers) | **Get** /node/network/name-servers | Read the NSX Manager&#x27;s Name Servers
[**ReadNodeNetworkRoute**](SystemAdministrationConfigurationApi.md#ReadNodeNetworkRoute) | **Get** /node/network/routes/{route-id} | Read node network route
[**ReadNodeProcess**](SystemAdministrationConfigurationApi.md#ReadNodeProcess) | **Get** /node/processes/{process-id} | Read node process
[**ReadNodeProperties**](SystemAdministrationConfigurationApi.md#ReadNodeProperties) | **Get** /node | Read node properties
[**ReadNodeSearchDomains**](SystemAdministrationConfigurationApi.md#ReadNodeSearchDomains) | **Get** /node/network/search-domains | Read the NSX Manager&#x27;s Search Domains
[**ReadNodeStatsService**](SystemAdministrationConfigurationApi.md#ReadNodeStatsService) | **Get** /node/services/node-stats | Read NSX node-stats service properties
[**ReadNodeStatsServiceStatus**](SystemAdministrationConfigurationApi.md#ReadNodeStatsServiceStatus) | **Get** /node/services/node-stats/status | Read NSX node-stats service status
[**ReadNodeStatus**](SystemAdministrationConfigurationApi.md#ReadNodeStatus) | **Get** /fabric/nodes/{node-id}/status | Return Runtime Status Information for a Node
[**ReadNodeSupportBundle**](SystemAdministrationConfigurationApi.md#ReadNodeSupportBundle) | **Get** /node/support-bundle | Read node support bundle
[**ReadNodeSyslogExporter**](SystemAdministrationConfigurationApi.md#ReadNodeSyslogExporter) | **Get** /node/services/syslog/exporters/{exporter-name} | Read node syslog exporter
[**ReadNodeVersion**](SystemAdministrationConfigurationApi.md#ReadNodeVersion) | **Get** /node/version | Read node version
[**ReadNodesStatus**](SystemAdministrationConfigurationApi.md#ReadNodesStatus) | **Get** /fabric/nodes/status | Return Runtime Status Information for given Nodes
[**ReadNsxUiServiceService**](SystemAdministrationConfigurationApi.md#ReadNsxUiServiceService) | **Get** /node/services/ui-service | Read ui service properties
[**ReadNsxUiServiceServiceStatus**](SystemAdministrationConfigurationApi.md#ReadNsxUiServiceServiceStatus) | **Get** /node/services/ui-service/status | Read ui service status
[**ReadNsxUpgradeAgentService**](SystemAdministrationConfigurationApi.md#ReadNsxUpgradeAgentService) | **Get** /node/services/nsx-upgrade-agent | Read NSX upgrade Agent service properties
[**ReadNsxUpgradeAgentServiceStatus**](SystemAdministrationConfigurationApi.md#ReadNsxUpgradeAgentServiceStatus) | **Get** /node/services/nsx-upgrade-agent/status | Read Nsx upgrade agent service status
[**ReadPaceClusterNodeVMDeploymentRequest**](SystemAdministrationConfigurationApi.md#ReadPaceClusterNodeVMDeploymentRequest) | **Get** /intelligence/nodes/deployments/{node-id} | Returns info for a Intelligence cluster node VM auto-deployment attempt
[**ReadPaceClusterNodeVMDeploymentStatus**](SystemAdministrationConfigurationApi.md#ReadPaceClusterNodeVMDeploymentStatus) | **Get** /intelligence/nodes/deployments/{node-id}/status | Returns the status of the VM creation/deletion
[**ReadPhonehomeCoordinatorService**](SystemAdministrationConfigurationApi.md#ReadPhonehomeCoordinatorService) | **Get** /node/services/telemetry | Read Telemetry service properties
[**ReadPhonehomeCoordinatorServiceStatus**](SystemAdministrationConfigurationApi.md#ReadPhonehomeCoordinatorServiceStatus) | **Get** /node/services/telemetry/status | Read Telemetry service status
[**ReadPlatformClientService**](SystemAdministrationConfigurationApi.md#ReadPlatformClientService) | **Get** /node/services/nsx-platform-client | Read NSX Platform Client service properties
[**ReadPlatformClientServiceStatus**](SystemAdministrationConfigurationApi.md#ReadPlatformClientServiceStatus) | **Get** /node/services/nsx-platform-client/status | Read NSX Platform Client service status
[**ReadPolicyService**](SystemAdministrationConfigurationApi.md#ReadPolicyService) | **Get** /node/services/policy | Read service properties
[**ReadPolicyServiceStatus**](SystemAdministrationConfigurationApi.md#ReadPolicyServiceStatus) | **Get** /node/services/policy/status | Read service status
[**ReadProtonService**](SystemAdministrationConfigurationApi.md#ReadProtonService) | **Get** /node/services/manager | Read service properties
[**ReadProtonServiceStatus**](SystemAdministrationConfigurationApi.md#ReadProtonServiceStatus) | **Get** /node/services/manager/status | Read service status
[**ReadProxyService**](SystemAdministrationConfigurationApi.md#ReadProxyService) | **Get** /node/services/http | Read http service properties
[**ReadProxyServiceStatus**](SystemAdministrationConfigurationApi.md#ReadProxyServiceStatus) | **Get** /node/services/http/status | Read http service status
[**ReadRabbitMQService**](SystemAdministrationConfigurationApi.md#ReadRabbitMQService) | **Get** /node/services/mgmt-plane-bus | Read Rabbit MQ service properties
[**ReadRabbitMQServiceStatus**](SystemAdministrationConfigurationApi.md#ReadRabbitMQServiceStatus) | **Get** /node/services/mgmt-plane-bus/status | Read Rabbit MQ service status
[**ReadRepositoryService**](SystemAdministrationConfigurationApi.md#ReadRepositoryService) | **Get** /node/services/install-upgrade | Read NSX install-upgrade service properties
[**ReadRepositoryServiceStatus**](SystemAdministrationConfigurationApi.md#ReadRepositoryServiceStatus) | **Get** /node/services/install-upgrade/status | Read NSX install-upgrade service status
[**ReadSNMPService**](SystemAdministrationConfigurationApi.md#ReadSNMPService) | **Get** /node/services/snmp | Read SNMP service properties
[**ReadSNMPServiceStatus**](SystemAdministrationConfigurationApi.md#ReadSNMPServiceStatus) | **Get** /node/services/snmp/status | Read SNMP service status
[**ReadSNMPV3EngineID**](SystemAdministrationConfigurationApi.md#ReadSNMPV3EngineID) | **Get** /node/services/snmp/v3-engine-id | Read SNMP V3 Engine ID
[**ReadSSHService**](SystemAdministrationConfigurationApi.md#ReadSSHService) | **Get** /node/services/ssh | Read ssh service properties
[**ReadSSHServiceStatus**](SystemAdministrationConfigurationApi.md#ReadSSHServiceStatus) | **Get** /node/services/ssh/status | Read ssh service status
[**ReadSearchService**](SystemAdministrationConfigurationApi.md#ReadSearchService) | **Get** /node/services/search | Read NSX Search service properties
[**ReadSearchServiceStatus**](SystemAdministrationConfigurationApi.md#ReadSearchServiceStatus) | **Get** /node/services/search/status | Read NSX Search service status
[**ReadSyslogService**](SystemAdministrationConfigurationApi.md#ReadSyslogService) | **Get** /node/services/syslog | Read syslog service properties
[**ReadSyslogServiceStatus**](SystemAdministrationConfigurationApi.md#ReadSyslogServiceStatus) | **Get** /node/services/syslog/status | Read syslog service status
[**ReadTaskProperties**](SystemAdministrationConfigurationApi.md#ReadTaskProperties) | **Get** /tasks/{task-id} | Get information about the specified task
[**ReadTaskResult**](SystemAdministrationConfigurationApi.md#ReadTaskResult) | **Get** /tasks/{task-id}/response | Get the response of a task
[**ReadTransportNodeInterface**](SystemAdministrationConfigurationApi.md#ReadTransportNodeInterface) | **Get** /transport-nodes/{transport-node-id}/network/interfaces/{interface-id} | Read the transport node&#x27;s network interface
[**ReadTransportNodeInterfaceStatistics**](SystemAdministrationConfigurationApi.md#ReadTransportNodeInterfaceStatistics) | **Get** /transport-nodes/{transport-node-id}/network/interfaces/{interface-id}/stats | Read the NSX Manager&#x27;s Network Interface Statistics
[**ReadVNIPool**](SystemAdministrationConfigurationApi.md#ReadVNIPool) | **Get** /pools/vni-pools/{pool-id} | Read VNI Pool
[**ReadVtepLabelPool**](SystemAdministrationConfigurationApi.md#ReadVtepLabelPool) | **Get** /pools/vtep-label-pools/{pool-id} | Read a virtual tunnel endpoint label pool
[**ReapplyTNProfileOnDiscoveredNodeReapplyClusterConfig**](SystemAdministrationConfigurationApi.md#ReapplyTNProfileOnDiscoveredNodeReapplyClusterConfig) | **Post** /fabric/discovered-nodes/{node-ext-id}?action&#x3D;reapply_cluster_config | Apply cluster level config on Discovered Node
[**RefreshTransportNode**](SystemAdministrationConfigurationApi.md#RefreshTransportNode) | **Post** /transport-nodes/{transport-node-id}?action&#x3D;refresh_node_configuration&amp;resource_type&#x3D;EdgeNode | Refresh the node configuration for the Edge node.
[**RegisterBatchRequest**](SystemAdministrationConfigurationApi.md#RegisterBatchRequest) | **Post** /batch | Register a Collection of API Calls at a Single End Point
[**RemoveVirtualMachineTagsRemoveTags**](SystemAdministrationConfigurationApi.md#RemoveVirtualMachineTagsRemoveTags) | **Post** /fabric/virtual-machines?action&#x3D;remove_tags | Perform action on specified virtual machine e.g. update tags
[**ReplaceEdgeClusterMemberTransportNodeReplaceTransportNode**](SystemAdministrationConfigurationApi.md#ReplaceEdgeClusterMemberTransportNodeReplaceTransportNode) | **Post** /edge-clusters/{edge-cluster-id}?action&#x3D;replace_transport_node | Replace the transport node in the specified member of the edge-cluster
[**RequestDirectoryDomainSync**](SystemAdministrationConfigurationApi.md#RequestDirectoryDomainSync) | **Post** /directory/domains/{domain-id} | Invoke full sync or delta sync for a specific domain, with additional delay in seconds if needed.  Stop sync will try to stop any pending sync if any to return to idle state.
[**ResetPaceHostConfigurationReset**](SystemAdministrationConfigurationApi.md#ResetPaceHostConfigurationReset) | **Post** /intelligence/host-config?action&#x3D;reset | Reset NSX-Intelligence host configuration
[**ResetPolicyServiceLoggingLevelActionResetManagerLoggingLevels**](SystemAdministrationConfigurationApi.md#ResetPolicyServiceLoggingLevelActionResetManagerLoggingLevels) | **Post** /node/services/policy?action&#x3D;reset-manager-logging-levels | Reset the logging levels to default values
[**ResetProtonServiceLoggingLevelActionResetManagerLoggingLevels**](SystemAdministrationConfigurationApi.md#ResetProtonServiceLoggingLevelActionResetManagerLoggingLevels) | **Post** /node/services/manager?action&#x3D;reset-manager-logging-levels | Reset the logging levels to default values
[**RestartInventorySyncRestartInventorySync**](SystemAdministrationConfigurationApi.md#RestartInventorySyncRestartInventorySync) | **Post** /fabric/nodes/{node-id}?action&#x3D;restart_inventory_sync | Restart the inventory sync for the node if it is paused currently.
[**RestartOrShutdownNodeRestart**](SystemAdministrationConfigurationApi.md#RestartOrShutdownNodeRestart) | **Post** /node?action&#x3D;restart | Restart or shutdown node
[**RestartOrShutdownNodeShutdown**](SystemAdministrationConfigurationApi.md#RestartOrShutdownNodeShutdown) | **Post** /node?action&#x3D;shutdown | Restart or shutdown node
[**RestartTransportNodeInventorySyncRestartInventorySync**](SystemAdministrationConfigurationApi.md#RestartTransportNodeInventorySyncRestartInventorySync) | **Post** /transport-nodes/{transport-node-id}?action&#x3D;restart_inventory_sync | Restart the inventory sync for the node if it is paused currently.
[**RestoreParentClusterConfigurationRestoreClusterConfig**](SystemAdministrationConfigurationApi.md#RestoreParentClusterConfigurationRestoreClusterConfig) | **Post** /transport-nodes/{transport-node-id}?action&#x3D;restore_cluster_config | Apply cluster level Transport Node Profile on overridden host
[**ResyncGlobalConfigsResyncConfig**](SystemAdministrationConfigurationApi.md#ResyncGlobalConfigsResyncConfig) | **Put** /global-configs/{config-type}?action&#x3D;resync_config | Resyncs global configurations of a config-type
[**ResyncTransportNodeResyncHostConfig**](SystemAdministrationConfigurationApi.md#ResyncTransportNodeResyncHostConfig) | **Post** /transport-nodes/{transportnode-id}?action&#x3D;resync_host_config | Resync a Transport Node
[**ScanDirectoryDomainSize**](SystemAdministrationConfigurationApi.md#ScanDirectoryDomainSize) | **Post** /directory/domain-size | Scan  the size of a directory domain
[**SearchDirectoryGroups**](SystemAdministrationConfigurationApi.md#SearchDirectoryGroups) | **Get** /directory/domains/{domain-id}/groups | Search for directory groups within a domain based on the substring of a distinguished name. (e.g. CN&#x3D;User,DC&#x3D;acme,DC&#x3D;com) The search filter pattern can optionally support multiple (up to 100 maximum) search pattern separated by &#x27;|&#x27; (url encoded %7C). In this case, the search results will be returned as the union of all matching criteria. (e.g. CN&#x3D;Ann,CN&#x3D;Users,DC&#x3D;acme,DC&#x3D;com|CN&#x3D;Bob,CN&#x3D;Users,DC&#x3D;acme,DC&#x3D;com)
[**SetClusterCertificateSetClusterCertificate**](SystemAdministrationConfigurationApi.md#SetClusterCertificateSetClusterCertificate) | **Post** /cluster/api-certificate?action&#x3D;set_cluster_certificate | Set the cluster certificate
[**SetClusterVirtualIpSetVirtualIp**](SystemAdministrationConfigurationApi.md#SetClusterVirtualIpSetVirtualIp) | **Post** /cluster/api-virtual-ip?action&#x3D;set_virtual_ip | Set cluster virtual IP address
[**SetControllerProfiler**](SystemAdministrationConfigurationApi.md#SetControllerProfiler) | **Put** /node/services/controller/profiler | Enable or disable controller profiler
[**SetNodeMandatoryAccessControl**](SystemAdministrationConfigurationApi.md#SetNodeMandatoryAccessControl) | **Put** /node/hardening-policy/mandatory-access-control | Enable or disable  Mandatory Access Control
[**SetNodeTimeSetSystemTime**](SystemAdministrationConfigurationApi.md#SetNodeTimeSetSystemTime) | **Post** /node?action&#x3D;set_system_time | Set the node system time
[**SetRabbitMQManagementPort**](SystemAdministrationConfigurationApi.md#SetRabbitMQManagementPort) | **Post** /node/rabbitmq-management-port | Set RabbitMQ management port
[**TestDirectoryLdapServer**](SystemAdministrationConfigurationApi.md#TestDirectoryLdapServer) | **Post** /directory/domains/{domain-id}/ldap-servers/{server-id} | Test a LDAP server connection for directory domain
[**UpdateApiServiceConfig**](SystemAdministrationConfigurationApi.md#UpdateApiServiceConfig) | **Put** /cluster/api-service | Update API service properties
[**UpdateApplianceNodeStatusClearBootupError**](SystemAdministrationConfigurationApi.md#UpdateApplianceNodeStatusClearBootupError) | **Post** /node/status?action&#x3D;clear_bootup_error | Update node status
[**UpdateAsyncReplicatorService**](SystemAdministrationConfigurationApi.md#UpdateAsyncReplicatorService) | **Put** /node/services/async_replicator | Update the async_replicator service properties
[**UpdateCentralConfigProperties**](SystemAdministrationConfigurationApi.md#UpdateCentralConfigProperties) | **Put** /node/central-config | Update Central Config properties
[**UpdateCentralNodeConfigProfile**](SystemAdministrationConfigurationApi.md#UpdateCentralNodeConfigProfile) | **Put** /configs/central-config/node-config-profiles/{node-config-profile-id} | Configure Node Config profile
[**UpdateClusterProfile**](SystemAdministrationConfigurationApi.md#UpdateClusterProfile) | **Put** /cluster-profiles/{cluster-profile-id} | Update a cluster profile
[**UpdateComputeCollectionFabricTemplate**](SystemAdministrationConfigurationApi.md#UpdateComputeCollectionFabricTemplate) | **Put** /fabric/compute-collection-fabric-templates/{fabric-template-id} | Updates compute collection fabric template
[**UpdateComputeManager**](SystemAdministrationConfigurationApi.md#UpdateComputeManager) | **Put** /fabric/compute-managers/{compute-manager-id} | Update compute manager
[**UpdateDirectoryDomain**](SystemAdministrationConfigurationApi.md#UpdateDirectoryDomain) | **Put** /directory/domains/{domain-id} | Update a directory domain
[**UpdateDirectoryLdapServer**](SystemAdministrationConfigurationApi.md#UpdateDirectoryLdapServer) | **Put** /directory/domains/{domain-id}/ldap-servers/{server-id} | Update a LDAP server for directory domain
[**UpdateEdgeCluster**](SystemAdministrationConfigurationApi.md#UpdateEdgeCluster) | **Put** /edge-clusters/{edge-cluster-id} | Update Edge Cluster
[**UpdateFailureDomain**](SystemAdministrationConfigurationApi.md#UpdateFailureDomain) | **Put** /failure-domains/{failure-domain-id} | Update Failure Domain
[**UpdateGlobalConfigs**](SystemAdministrationConfigurationApi.md#UpdateGlobalConfigs) | **Put** /global-configs/{config-type} | Update global configurations of a config type
[**UpdateHostSwitchProfile**](SystemAdministrationConfigurationApi.md#UpdateHostSwitchProfile) | **Put** /host-switch-profiles/{host-switch-profile-id} | Update a Hostswitch Profile
[**UpdateIpBlock**](SystemAdministrationConfigurationApi.md#UpdateIpBlock) | **Put** /pools/ip-blocks/{block-id} | Update an IP Address Block
[**UpdateIpPool**](SystemAdministrationConfigurationApi.md#UpdateIpPool) | **Put** /pools/ip-pools/{pool-id} | Update an IP Pool
[**UpdateMPAConfiguration**](SystemAdministrationConfigurationApi.md#UpdateMPAConfiguration) | **Put** /node/mpa-config | Update MPA configuration for this node
[**UpdateManagementConfig**](SystemAdministrationConfigurationApi.md#UpdateManagementConfig) | **Put** /configs/management | Update NSX Management nodes global configuration
[**UpdateManagementPlaneConfiguration**](SystemAdministrationConfigurationApi.md#UpdateManagementPlaneConfiguration) | **Put** /node/management-plane | Update management plane configuration for this node
[**UpdateNTPService**](SystemAdministrationConfigurationApi.md#UpdateNTPService) | **Put** /node/services/ntp | Update NTP service properties
[**UpdateNode**](SystemAdministrationConfigurationApi.md#UpdateNode) | **Put** /fabric/nodes/{node-id} | Update a Node
[**UpdateNodeInterface**](SystemAdministrationConfigurationApi.md#UpdateNodeInterface) | **Put** /node/network/interfaces/{interface-id} | Update the NSX Manager&#x27;s Network Interface
[**UpdateNodeNameServers**](SystemAdministrationConfigurationApi.md#UpdateNodeNameServers) | **Put** /node/network/name-servers | Update the NSX Manager&#x27;s Name Servers
[**UpdateNodeProperties**](SystemAdministrationConfigurationApi.md#UpdateNodeProperties) | **Put** /node | Update node properties
[**UpdateNodeSearchDomains**](SystemAdministrationConfigurationApi.md#UpdateNodeSearchDomains) | **Put** /node/network/search-domains | Update the NSX Manager&#x27;s Search Domains
[**UpdatePolicyService**](SystemAdministrationConfigurationApi.md#UpdatePolicyService) | **Put** /node/services/policy | Update service properties
[**UpdateProtonService**](SystemAdministrationConfigurationApi.md#UpdateProtonService) | **Put** /node/services/manager | Update service properties
[**UpdateProxyService**](SystemAdministrationConfigurationApi.md#UpdateProxyService) | **Put** /node/services/http | Update http service properties
[**UpdateRealizationStateBarrierConfig**](SystemAdministrationConfigurationApi.md#UpdateRealizationStateBarrierConfig) | **Put** /realization-state-barrier/config | Updates the barrier configuration
[**UpdateRepositoryService**](SystemAdministrationConfigurationApi.md#UpdateRepositoryService) | **Put** /node/services/install-upgrade | Update NSX install-upgrade service properties
[**UpdateSNMPService**](SystemAdministrationConfigurationApi.md#UpdateSNMPService) | **Put** /node/services/snmp | Update SNMP service properties
[**UpdateSNMPV3EngineID**](SystemAdministrationConfigurationApi.md#UpdateSNMPV3EngineID) | **Put** /node/services/snmp/v3-engine-id | Update SNMP V3 Engine ID
[**UpdateSSHService**](SystemAdministrationConfigurationApi.md#UpdateSSHService) | **Put** /node/services/ssh | Update ssh service properties
[**UpdateTransportNodeCollection**](SystemAdministrationConfigurationApi.md#UpdateTransportNodeCollection) | **Put** /transport-node-collections/{transport-node-collection-id} | Update Transport Node collection
[**UpdateTransportNodeMaintenanceMode**](SystemAdministrationConfigurationApi.md#UpdateTransportNodeMaintenanceMode) | **Post** /transport-nodes/{transportnode-id} | Update transport node maintenance mode
[**UpdateTransportNodeProfile**](SystemAdministrationConfigurationApi.md#UpdateTransportNodeProfile) | **Put** /transport-node-profiles/{transport-node-profile-id} | Update a Transport Node Profile
[**UpdateTransportNodeWithDeploymentInfo**](SystemAdministrationConfigurationApi.md#UpdateTransportNodeWithDeploymentInfo) | **Put** /transport-nodes/{transport-node-id} | Update a Transport Node
[**UpdateTransportZone**](SystemAdministrationConfigurationApi.md#UpdateTransportZone) | **Put** /transport-zones/{zone-id} | Update a Transport Zone
[**UpdateTransportZoneProfile**](SystemAdministrationConfigurationApi.md#UpdateTransportZoneProfile) | **Put** /transportzone-profiles/{transportzone-profile-id} | Update a transport zone profile
[**UpdateUcState**](SystemAdministrationConfigurationApi.md#UpdateUcState) | **Put** /node/services/install-upgrade/uc-state | Update UC state properties
[**UpdateVNIPool**](SystemAdministrationConfigurationApi.md#UpdateVNIPool) | **Put** /pools/vni-pools/{pool-id} | Update a VNI Pool
[**UpdateVirtualMachineTagsUpdateTags**](SystemAdministrationConfigurationApi.md#UpdateVirtualMachineTagsUpdateTags) | **Post** /fabric/virtual-machines?action&#x3D;update_tags | Perform action on specified virtual machine e.g. update tags
[**UploadBundleViaLocalFileUpload**](SystemAdministrationConfigurationApi.md#UploadBundleViaLocalFileUpload) | **Post** /repository/bundles?action&#x3D;upload | Upload bundle
[**UploadBundleViaRemoteFile**](SystemAdministrationConfigurationApi.md#UploadBundleViaRemoteFile) | **Post** /repository/bundles | Upload bundle using remote file
[**VerifyDirectoryLdapServer**](SystemAdministrationConfigurationApi.md#VerifyDirectoryLdapServer) | **Post** /directory/ldap-server | Test a directory domain LDAP server connectivity
[**VerifyNodeSyslogExporterVerify**](SystemAdministrationConfigurationApi.md#VerifyNodeSyslogExporterVerify) | **Post** /node/services/syslog/exporters?action&#x3D;verify | Verify node syslog exporter

# **AddClusterNode**
> ClusterNodeConfig AddClusterNode(ctx, body, action)
Add a controller to the cluster

Add a new controller to the NSX cluster. Deprecated. Use POST /cluster?action=join_cluster to join a node to cluster. The controller comes with the new node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddClusterNodeSpec**](AddClusterNodeSpec.md)|  | 
  **action** | **string**|  | 

### Return type

[**ClusterNodeConfig**](ClusterNodeConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddClusterNodeVM**
> ClusterNodeVmDeploymentRequestList AddClusterNodeVM(ctx, body)
Deploy and register a cluster node VM

Deploys a cluster node VM as specified by the deployment config. Once the VM is deployed and powered on, it will automatically join the existing cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddClusterNodeVmInfo**](AddClusterNodeVmInfo.md)|  | 

### Return type

[**ClusterNodeVmDeploymentRequestList**](ClusterNodeVMDeploymentRequestList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddComputeManager**
> ComputeManager AddComputeManager(ctx, body)
Register compute manager with NSX

Registers compute manager with NSX. Inventory service will collect data from the registered compute manager 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeManager**](ComputeManager.md)|  | 

### Return type

[**ComputeManager**](ComputeManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddNode**
> Node AddNode(ctx, body)
Register and Install NSX Components on a Node

Creates a host node (hypervisor) or edge node (router) in the transport network.  When you run this command for a host, NSX Manager attempts to install the NSX kernel modules, which are packaged as VIB, RPM, or DEB files. For the installation to succeed, you must provide the host login credentials and the host thumbprint.  To get the ESXi host thumbprint, SSH to the host and run the <b>openssl x509 -in /etc/vmware/ssl/rui.crt -fingerprint -sha256 -noout</b> command.  To generate host key thumbprint using SHA-256 algorithm please follow the steps below.  Log into the host, making sure that the connection is not vulnerable to a man in the middle attack. Check whether a public key already exists. Host public key is generally located at '/etc/ssh/ssh_host_rsa_key.pub'. If the key is not present then generate a new key by running the following command and follow the instructions.  <b>ssh-keygen -t rsa</b>  Now generate a SHA256 hash of the key using the following command. Please make sure to pass the appropriate file name if the public key is stored with a different file name other than the default 'id_rsa.pub'.  <b>awk '{print $2}' id_rsa.pub | base64 -d | sha256sum -b | sed 's/ .*$//' | xxd -r -p | base64</b> This api is deprecated as part of FN+TN unification. Please use Transport Node API POST /transport-nodes to install NSX components on a node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Node**](Node.md)|  | 

### Return type

[**Node**](Node.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddPaceClusterNodeVM**
> IntelligenceClusterNodeVmDeploymentRequestList AddPaceClusterNodeVM(ctx, body)
Deploy and register a Intelligence cluster node VM

Deploys a Intelligence cluster node VM as specified by the deployment config. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddIntelligenceClusterNodeVmInfo**](AddIntelligenceClusterNodeVmInfo.md)|  | 

### Return type

[**IntelligenceClusterNodeVmDeploymentRequestList**](IntelligenceClusterNodeVMDeploymentRequestList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddVirtualMachineTagsAddTags**
> AddVirtualMachineTagsAddTags(ctx, body)
Perform action on specified virtual machine e.g. update tags

Perform action on a specific virtual machine. External id of the virtual machine needs to be provided in the request body. Some of the actions that can be performed are update tags, add tags, remove tags. To add tags to existing list of tag, use action parameter add_tags. To remove tags from existing list of tag, use action parameter remove_tags. To replace existing tags with new tags, use action parameter update_tags. To clear all tags, provide an empty list and action parameter as update_tags. The vmw-async: True HTTP header cannot be used with this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualMachineTagUpdate**](VirtualMachineTagUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AllocateOrReleaseFromIpBlockSubnet**
> AllocationIpAddress AllocateOrReleaseFromIpBlockSubnet(ctx, body, subnetId, action)
Allocate or Release an IP Address from a Ip Subnet

Allocates or releases an IP address from the specified IP subnet. To allocate an address, include ?action=ALLOCATE in the request and a \"{}\" in the request body. When the request is successful, the response is \"allocation_id\": \"<ip-address>\", where <ip-address> is an IP address from the specified pool. To release an IP address (return it back to the pool), include ?action=RELEASE in the request and \"allocation_id\":<ip-address> in the request body, where <ip-address> is the address to be released. When the request is successful, the response is NULL. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AllocationIpAddress**](AllocationIpAddress.md)|  | 
  **subnetId** | **string**| IP subnet id | 
  **action** | **string**| Specifies allocate or release action | 

### Return type

[**AllocationIpAddress**](AllocationIpAddress.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AllocateOrReleaseFromIpPool**
> AllocationIpAddress AllocateOrReleaseFromIpPool(ctx, body, poolId, action)
Allocate or Release an IP Address from a Pool

Allocates or releases an IP address from the specified IP pool. To allocate an address, include ?action=ALLOCATE in the request and \"allocation_id\":null in the request body. When the request is successful, the response is \"allocation_id\": \"<ip-address>\", where <ip-address> is an IP address from the specified pool. To release an IP address (return it back to the pool), include ?action=RELEASE in the request and \"allocation_id\":<ip-address> in the request body, where <ip-address> is the address to be released. When the request is successful, the response is NULL. Tags, display_name and description attributes are not supported for AllocationIpAddress in this release. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AllocationIpAddress**](AllocationIpAddress.md)|  | 
  **poolId** | **string**| IP pool ID | 
  **action** | **string**| Specifies allocate or release action | 

### Return type

[**AllocationIpAddress**](AllocationIpAddress.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelApplianceManagementTaskCancel**
> CancelApplianceManagementTaskCancel(ctx, taskId)
Cancel specified task

Cancel specified task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelBundleUploadCancelUpload**
> CancelBundleUploadCancelUpload(ctx, bundleId, product)
Cancel bundle upload

Cancel upload of bundle. This API works only when bundle upload is in-progress and will not work during post-processing of bundle. If bundle upload is in-progress, then the API call returns http OK response after cancelling the upload and deleting partially uploaded bundle. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**|  | 
  **product** | **string**| Name of the product | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeNodeMode**
> NodeMode ChangeNodeMode(ctx, body)
NodeMode

Currently only a switch from \"VMC_LOCAL\" to \"VMC\" is supported. Returns a new Node Mode, if the request successfuly changed it. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwitchingToVmcModeParameters**](SwitchingToVmcModeParameters.md)|  | 

### Return type

[**NodeMode**](NodeMode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckRabbitMQManagementPort**
> PortStatus CheckRabbitMQManagementPort(ctx, )
Check if RabbitMQ management port is enabled or not

Returns status as true if RabbitMQ management port is enabled else false

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PortStatus**](PortStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearClusterCertificateClearClusterCertificate**
> ClusterCertificateId ClearClusterCertificateClearClusterCertificate(ctx, certificateId)
Clear the cluster certificate

Clears the certificate used for the MP cluster. This does not affect the certificate itself. This API is deprecated. Instead use the  /api/v1/cluster/api-certificate?action=set_cluster_certificate API to set the cluster certificate to a different one. It just means that from now on, individual certificates will be used on each MP node. This affects all nodes in the cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**| Certificate ID | 

### Return type

[**ClusterCertificateId**](ClusterCertificateId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearClusterVirtualIpClearVirtualIp**
> ClusterVirtualIpProperties ClearClusterVirtualIpClearVirtualIp(ctx, )
Clear cluster virtual IP address

Clears the cluster virtual IP address. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterVirtualIpProperties**](ClusterVirtualIpProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateApplianceManagementServiceActionRestart**
> CreateApplianceManagementServiceActionRestart(ctx, )
Restart the node management service

Restart the node management service

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAsyncReplicatorServiceActionRestart**
> NodeServiceStatusProperties CreateAsyncReplicatorServiceActionRestart(ctx, )
Restart, start or stop the Async Replicator service

Restart, start or stop the Async Replicator service

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

# **CreateAsyncReplicatorServiceActionStart**
> NodeServiceStatusProperties CreateAsyncReplicatorServiceActionStart(ctx, )
Restart, start or stop the Async Replicator service

Restart, start or stop the Async Replicator service

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

# **CreateAsyncReplicatorServiceActionStop**
> NodeServiceStatusProperties CreateAsyncReplicatorServiceActionStop(ctx, )
Restart, start or stop the Async Replicator service

Restart, start or stop the Async Replicator service

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

# **CreateClusterBootManagerServiceActionRestart**
> NodeServiceStatusProperties CreateClusterBootManagerServiceActionRestart(ctx, )
Restart, start or stop the cluster boot manager service

Restart, start or stop the cluster boot manager service

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

# **CreateClusterBootManagerServiceActionStart**
> NodeServiceStatusProperties CreateClusterBootManagerServiceActionStart(ctx, )
Restart, start or stop the cluster boot manager service

Restart, start or stop the cluster boot manager service

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

# **CreateClusterBootManagerServiceActionStop**
> NodeServiceStatusProperties CreateClusterBootManagerServiceActionStop(ctx, )
Restart, start or stop the cluster boot manager service

Restart, start or stop the cluster boot manager service

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

# **CreateClusterProfile**
> ClusterProfile CreateClusterProfile(ctx, body)
Create a Cluster Profile

Create a cluster profile. The resource_type is required. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterProfile**](ClusterProfile.md)|  | 

### Return type

[**ClusterProfile**](ClusterProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCminventoryServiceActionRestart**
> NodeServiceStatusProperties CreateCminventoryServiceActionRestart(ctx, )
Restart, start or stop the manager service

Restart, start or stop the manager service

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

# **CreateCminventoryServiceActionStart**
> NodeServiceStatusProperties CreateCminventoryServiceActionStart(ctx, )
Restart, start or stop the manager service

Restart, start or stop the manager service

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

# **CreateCminventoryServiceActionStop**
> NodeServiceStatusProperties CreateCminventoryServiceActionStop(ctx, )
Restart, start or stop the manager service

Restart, start or stop the manager service

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

# **CreateComputeCollectionFabricTemplate**
> ComputeCollectionFabricTemplate CreateComputeCollectionFabricTemplate(ctx, body)
Create a compute collection fabric template

Fabric templates are fabric configurations applied at the compute collection level. This configurations is used to decide what automated operations should be a run when a host membership changes. This functionality is deprecated. Use Transport Node Profiles instead of this template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeCollectionFabricTemplate**](ComputeCollectionFabricTemplate.md)|  | 

### Return type

[**ComputeCollectionFabricTemplate**](ComputeCollectionFabricTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateControllerServerServiceActionRestart**
> NodeServiceStatusProperties CreateControllerServerServiceActionRestart(ctx, )
Restart, start or stop the controller service

Restart, start or stop the controller service

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

# **CreateControllerServerServiceActionStart**
> NodeServiceStatusProperties CreateControllerServerServiceActionStart(ctx, )
Restart, start or stop the controller service

Restart, start or stop the controller service

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

# **CreateControllerServerServiceActionStop**
> NodeServiceStatusProperties CreateControllerServerServiceActionStop(ctx, )
Restart, start or stop the controller service

Restart, start or stop the controller service

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

# **CreateDirectoryDomain**
> DirectoryDomain CreateDirectoryDomain(ctx, body)
Create a directory domain

Create a directory domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectoryDomain**](DirectoryDomain.md)|  | 

### Return type

[**DirectoryDomain**](DirectoryDomain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDirectoryLdapServer**
> DirectoryLdapServer CreateDirectoryLdapServer(ctx, body, domainId)
Create a LDAP server for directory domain

More than one LDAP server can be created and only one LDAP server is used to synchronize directory objects. If more than one LDAP server is configured, NSX will try all the servers until it is able to successfully connect to one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectoryLdapServer**](DirectoryLdapServer.md)|  | 
  **domainId** | **string**| Directory domain identifier | 

### Return type

[**DirectoryLdapServer**](DirectoryLdapServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEdgeCluster**
> EdgeCluster CreateEdgeCluster(ctx, body)
Create Edge Cluster

Creates a new edge cluster. It only supports homogeneous members. The TransportNodes backed by EdgeNode are only allowed in cluster members. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EdgeCluster**](EdgeCluster.md)|  | 

### Return type

[**EdgeCluster**](EdgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFailureDomain**
> FailureDomain CreateFailureDomain(ctx, body)
Create Failure Domain

Creates a new failure domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FailureDomain**](FailureDomain.md)|  | 

### Return type

[**FailureDomain**](FailureDomain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHostSwitchProfile**
> BaseHostSwitchProfile CreateHostSwitchProfile(ctx, body)
Create a Hostswitch Profile

Creates a hostswitch profile. The resource_type is required. For uplink profiles, the teaming and policy parameters are required. By default, the mtu is 1600 and the transport_vlan is 0. The supported MTU range is 1280 through (uplink_mtu_threshold). (uplink_mtu_threshold) is 9000 by default. Range can be extended by modifying (uplink_mtu_threshold) in SwitchingGlobalConfig to the required upper threshold. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseHostSwitchProfile**](BaseHostSwitchProfile.md)|  | 

### Return type

[**BaseHostSwitchProfile**](BaseHostSwitchProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIdpsReportingServiceActionRestart**
> NodeServiceStatusProperties CreateIdpsReportingServiceActionRestart(ctx, )
Restart, start or stop the idps-reporting service

Restart, start or stop the idps-reporting service

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

# **CreateIdpsReportingServiceActionStart**
> NodeServiceStatusProperties CreateIdpsReportingServiceActionStart(ctx, )
Restart, start or stop the idps-reporting service

Restart, start or stop the idps-reporting service

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

# **CreateIdpsReportingServiceActionStop**
> NodeServiceStatusProperties CreateIdpsReportingServiceActionStop(ctx, )
Restart, start or stop the idps-reporting service

Restart, start or stop the idps-reporting service

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

# **CreateIntelligenceUpgradeCoordinatorServiceActionRestart**
> NodeServiceStatusProperties CreateIntelligenceUpgradeCoordinatorServiceActionRestart(ctx, )
Restart, start or stop the intelligence upgrade coordinator service

Restart, start or stop the intelligence upgrade coordinator service

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

# **CreateIntelligenceUpgradeCoordinatorServiceActionStart**
> NodeServiceStatusProperties CreateIntelligenceUpgradeCoordinatorServiceActionStart(ctx, )
Restart, start or stop the intelligence upgrade coordinator service

Restart, start or stop the intelligence upgrade coordinator service

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

# **CreateIntelligenceUpgradeCoordinatorServiceActionStop**
> NodeServiceStatusProperties CreateIntelligenceUpgradeCoordinatorServiceActionStop(ctx, )
Restart, start or stop the intelligence upgrade coordinator service

Restart, start or stop the intelligence upgrade coordinator service

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

# **CreateIpBlock**
> IpBlock CreateIpBlock(ctx, body)
Create a new IP address block.

Creates a new IPv4 address block using the specified cidr. cidr is a required parameter. display_name & description are optional parameters 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpBlock**](IpBlock.md)|  | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpBlockSubnet**
> IpBlockSubnet CreateIpBlockSubnet(ctx, body)
Create subnet of specified size within an IP block

Carves out a subnet of requested size from the specified IP block. The \"size\" parameter  and the \"block_id \" are the requireds field while invoking this API. If the IP block has sufficient resources/space to allocate a subnet of specified size, the response will contain all the details of the newly created subnet including the display_name, description, cidr & allocation_ranges. Returns a conflict error if the IP block does not have enough resources/space to allocate a subnet of the requested size. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpBlockSubnet**](IpBlockSubnet.md)|  | 

### Return type

[**IpBlockSubnet**](IpBlockSubnet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateIpPool**
> IpPool CreateIpPool(ctx, body)
Create an IP Pool

Creates a new IPv4 or IPv6 address pool. Required parameters are allocation_ranges and cidr. Optional parameters are display_name, description, dns_nameservers, dns_suffix, and gateway_ip. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpPool**](IpPool.md)|  | 

### Return type

[**IpPool**](IpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLiagentServiceActionRestart**
> NodeServiceStatusProperties CreateLiagentServiceActionRestart(ctx, )
Restart, start or stop the liagent service

Restart, start or stop the liagent service

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

# **CreateLiagentServiceActionStart**
> NodeServiceStatusProperties CreateLiagentServiceActionStart(ctx, )
Restart, start or stop the liagent service

Restart, start or stop the liagent service

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

# **CreateLiagentServiceActionStop**
> NodeServiceStatusProperties CreateLiagentServiceActionStop(ctx, )
Restart, start or stop the liagent service

Restart, start or stop the liagent service

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

# **CreateMigrationCoordinatorServiceActionRestart**
> NodeServiceStatusProperties CreateMigrationCoordinatorServiceActionRestart(ctx, )
Restart, start or stop the migration coordinator service

Restart, start or stop the migration coordinator service

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

# **CreateMigrationCoordinatorServiceActionStart**
> NodeServiceStatusProperties CreateMigrationCoordinatorServiceActionStart(ctx, )
Restart, start or stop the migration coordinator service

Restart, start or stop the migration coordinator service

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

# **CreateMigrationCoordinatorServiceActionStop**
> NodeServiceStatusProperties CreateMigrationCoordinatorServiceActionStop(ctx, )
Restart, start or stop the migration coordinator service

Restart, start or stop the migration coordinator service

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

# **CreateNSXMessageBusServiceActionRestart**
> NodeServiceStatusProperties CreateNSXMessageBusServiceActionRestart(ctx, )
Restart, start or stop the NSX Message Bus service

Restart, start or stop the NSX Message Bus service

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

# **CreateNSXMessageBusServiceActionStart**
> NodeServiceStatusProperties CreateNSXMessageBusServiceActionStart(ctx, )
Restart, start or stop the NSX Message Bus service

Restart, start or stop the NSX Message Bus service

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

# **CreateNSXMessageBusServiceActionStop**
> NodeServiceStatusProperties CreateNSXMessageBusServiceActionStop(ctx, )
Restart, start or stop the NSX Message Bus service

Restart, start or stop the NSX Message Bus service

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

# **CreateNTPServiceActionRestart**
> NodeServiceStatusProperties CreateNTPServiceActionRestart(ctx, )
Restart, start or stop the NTP service

Restart, start or stop the NTP service

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

# **CreateNTPServiceActionStart**
> NodeServiceStatusProperties CreateNTPServiceActionStart(ctx, )
Restart, start or stop the NTP service

Restart, start or stop the NTP service

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

# **CreateNTPServiceActionStop**
> NodeServiceStatusProperties CreateNTPServiceActionStop(ctx, )
Restart, start or stop the NTP service

Restart, start or stop the NTP service

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

# **CreateNodeNetworkRoute**
> NodeRouteProperties CreateNodeNetworkRoute(ctx, body)
Create node network route

Add a route to the NSX Manager routing table. For static routes, the route_type, interface_id, netmask, and destination are required parameters. For default routes, the route_type, gateway address, and interface_id are required. For blackhole routes, the route_type and destination are required. All other parameters are optional. When you add a static route, the scope and route_id are created automatically. When you add a default or blackhole route, the route_id is created automatically. The route_id is read-only, meaning that it cannot be modified. All other properties can be modified by deleting and readding the route. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeRouteProperties**](NodeRouteProperties.md)|  | 

### Return type

[**NodeRouteProperties**](NodeRouteProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNodeStatsServiceActionRestart**
> NodeServiceStatusProperties CreateNodeStatsServiceActionRestart(ctx, )
Restart, start or stop the NSX node-stats service

Restart, start or stop the NSX node-stats service

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

# **CreateNodeStatsServiceActionStart**
> NodeServiceStatusProperties CreateNodeStatsServiceActionStart(ctx, )
Restart, start or stop the NSX node-stats service

Restart, start or stop the NSX node-stats service

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

# **CreateNodeStatsServiceActionStop**
> NodeServiceStatusProperties CreateNodeStatsServiceActionStop(ctx, )
Restart, start or stop the NSX node-stats service

Restart, start or stop the NSX node-stats service

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

# **CreateNsxUiServiceServiceActionRestart**
> NodeServiceStatusProperties CreateNsxUiServiceServiceActionRestart(ctx, )
Restart, Start and Stop the ui service

Restart, Start and Stop the ui service

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

# **CreateNsxUiServiceServiceActionStart**
> NodeServiceStatusProperties CreateNsxUiServiceServiceActionStart(ctx, )
Restart, Start and Stop the ui service

Restart, Start and Stop the ui service

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

# **CreateNsxUiServiceServiceActionStop**
> NodeServiceStatusProperties CreateNsxUiServiceServiceActionStop(ctx, )
Restart, Start and Stop the ui service

Restart, Start and Stop the ui service

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

# **CreateNsxUpgradeAgentServiceActionRestart**
> NodeServiceStatusProperties CreateNsxUpgradeAgentServiceActionRestart(ctx, )
Restart, start or stop the NSX upgrade agent service

Restart, start or stop the NSX upgrade agent service

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

# **CreateNsxUpgradeAgentServiceActionStart**
> NodeServiceStatusProperties CreateNsxUpgradeAgentServiceActionStart(ctx, )
Restart, start or stop the NSX upgrade agent service

Restart, start or stop the NSX upgrade agent service

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

# **CreateNsxUpgradeAgentServiceActionStop**
> NodeServiceStatusProperties CreateNsxUpgradeAgentServiceActionStop(ctx, )
Restart, start or stop the NSX upgrade agent service

Restart, start or stop the NSX upgrade agent service

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

# **CreatePhonehomeCoordinatorServiceActionRestart**
> NodeServiceStatusProperties CreatePhonehomeCoordinatorServiceActionRestart(ctx, )
Restart, start or stop Telemetry service

Restart, start or stop Telemetry service

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

# **CreatePhonehomeCoordinatorServiceActionStart**
> NodeServiceStatusProperties CreatePhonehomeCoordinatorServiceActionStart(ctx, )
Restart, start or stop Telemetry service

Restart, start or stop Telemetry service

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

# **CreatePhonehomeCoordinatorServiceActionStop**
> NodeServiceStatusProperties CreatePhonehomeCoordinatorServiceActionStop(ctx, )
Restart, start or stop Telemetry service

Restart, start or stop Telemetry service

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

# **CreatePlatformClientServiceActionRestart**
> NodeServiceStatusProperties CreatePlatformClientServiceActionRestart(ctx, )
Restart, start or stop the NSX Platform Client service

Restart, start or stop the NSX Platform Client service

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

# **CreatePlatformClientServiceActionStart**
> NodeServiceStatusProperties CreatePlatformClientServiceActionStart(ctx, )
Restart, start or stop the NSX Platform Client service

Restart, start or stop the NSX Platform Client service

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

# **CreatePlatformClientServiceActionStop**
> NodeServiceStatusProperties CreatePlatformClientServiceActionStop(ctx, )
Restart, start or stop the NSX Platform Client service

Restart, start or stop the NSX Platform Client service

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

# **CreatePolicyServiceActionRestart**
> NodeServiceStatusProperties CreatePolicyServiceActionRestart(ctx, )
Restart, start or stop the service

Restart, start or stop the service

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

# **CreatePolicyServiceActionStart**
> NodeServiceStatusProperties CreatePolicyServiceActionStart(ctx, )
Restart, start or stop the service

Restart, start or stop the service

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

# **CreatePolicyServiceActionStop**
> NodeServiceStatusProperties CreatePolicyServiceActionStop(ctx, )
Restart, start or stop the service

Restart, start or stop the service

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

# **CreateProtonServiceActionRestart**
> NodeServiceStatusProperties CreateProtonServiceActionRestart(ctx, )
Restart, start or stop the service

Restart, start or stop the service

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

# **CreateProtonServiceActionStart**
> NodeServiceStatusProperties CreateProtonServiceActionStart(ctx, )
Restart, start or stop the service

Restart, start or stop the service

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

# **CreateProtonServiceActionStop**
> NodeServiceStatusProperties CreateProtonServiceActionStop(ctx, )
Restart, start or stop the service

Restart, start or stop the service

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

# **CreateProxyServiceActionRestart**
> CreateProxyServiceActionRestart(ctx, )
Restart the http service

Restart the http service

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProxyServiceActionStart**
> NodeServiceStatusProperties CreateProxyServiceActionStart(ctx, )
Start the http service

Start the http service

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

# **CreateProxyServiceActionStop**
> CreateProxyServiceActionStop(ctx, )
Stop the http service

Stop the http service

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateProxyServiceApplyCertificateActionApplyCertificate**
> CreateProxyServiceApplyCertificateActionApplyCertificate(ctx, certificateId)
Update http service certificate

Applies a security certificate to the http service. In the POST request, the CERTIFICATE_ID references a certificate created with the /api/v1/trust-management APIs. If the certificate used is a CA signed certificate,the request fails if the whole chain(leaf, intermediate, root) is not imported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**| Certificate ID | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRabbitMQServiceActionRestart**
> NodeServiceStatusProperties CreateRabbitMQServiceActionRestart(ctx, )
Restart, start or stop the Rabbit MQ service

Restart, start or stop the Rabbit MQ service

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

# **CreateRabbitMQServiceActionStart**
> NodeServiceStatusProperties CreateRabbitMQServiceActionStart(ctx, )
Restart, start or stop the Rabbit MQ service

Restart, start or stop the Rabbit MQ service

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

# **CreateRabbitMQServiceActionStop**
> NodeServiceStatusProperties CreateRabbitMQServiceActionStop(ctx, )
Restart, start or stop the Rabbit MQ service

Restart, start or stop the Rabbit MQ service

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

# **CreateRepositoryServiceActionRestart**
> NodeServiceStatusProperties CreateRepositoryServiceActionRestart(ctx, )
Restart, start or stop the NSX install-upgrade service

Restart, start or stop the NSX install-upgrade service

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

# **CreateRepositoryServiceActionStart**
> NodeServiceStatusProperties CreateRepositoryServiceActionStart(ctx, )
Restart, start or stop the NSX install-upgrade service

Restart, start or stop the NSX install-upgrade service

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

# **CreateRepositoryServiceActionStop**
> NodeServiceStatusProperties CreateRepositoryServiceActionStop(ctx, )
Restart, start or stop the NSX install-upgrade service

Restart, start or stop the NSX install-upgrade service

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

# **CreateSNMPServiceActionRestart**
> NodeServiceStatusProperties CreateSNMPServiceActionRestart(ctx, )
Restart, start or stop the SNMP service

Restart, start or stop the SNMP service

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

# **CreateSNMPServiceActionStart**
> NodeServiceStatusProperties CreateSNMPServiceActionStart(ctx, )
Restart, start or stop the SNMP service

Restart, start or stop the SNMP service

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

# **CreateSNMPServiceActionStop**
> NodeServiceStatusProperties CreateSNMPServiceActionStop(ctx, )
Restart, start or stop the SNMP service

Restart, start or stop the SNMP service

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

# **CreateSSHServiceActionNotifyMpaRestart**
> NodeServiceStatusProperties CreateSSHServiceActionNotifyMpaRestart(ctx, )
Restart, start or stop the ssh service

Restart, start or stop the ssh service

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

# **CreateSSHServiceActionNotifyMpaStart**
> NodeServiceStatusProperties CreateSSHServiceActionNotifyMpaStart(ctx, )
Restart, start or stop the ssh service

Restart, start or stop the ssh service

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

# **CreateSSHServiceActionNotifyMpaStop**
> NodeServiceStatusProperties CreateSSHServiceActionNotifyMpaStop(ctx, )
Restart, start or stop the ssh service

Restart, start or stop the ssh service

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

# **CreateSSHServiceActionRestart**
> NodeServiceStatusProperties CreateSSHServiceActionRestart(ctx, )
Restart, start or stop the ssh service

Restart, start or stop the ssh service

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

# **CreateSSHServiceActionStart**
> NodeServiceStatusProperties CreateSSHServiceActionStart(ctx, )
Restart, start or stop the ssh service

Restart, start or stop the ssh service

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

# **CreateSSHServiceActionStop**
> NodeServiceStatusProperties CreateSSHServiceActionStop(ctx, )
Restart, start or stop the ssh service

Restart, start or stop the ssh service

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

# **CreateSSHServiceRemoveHostFingerprintActionRemoveHostFingerprint**
> CreateSSHServiceRemoveHostFingerprintActionRemoveHostFingerprint(ctx, body)
Remove a host's fingerprint from known hosts file

Remove a host's fingerprint from known hosts file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**KnownHostParameter**](KnownHostParameter.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSearchServiceActionRestart**
> NodeServiceStatusProperties CreateSearchServiceActionRestart(ctx, )
Restart, start or stop the NSX Search service

Restart, start or stop the NSX Search service

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

# **CreateSearchServiceActionStart**
> NodeServiceStatusProperties CreateSearchServiceActionStart(ctx, )
Restart, start or stop the NSX Search service

Restart, start or stop the NSX Search service

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

# **CreateSearchServiceActionStop**
> NodeServiceStatusProperties CreateSearchServiceActionStop(ctx, )
Restart, start or stop the NSX Search service

Restart, start or stop the NSX Search service

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

# **CreateSyslogServiceActionRestart**
> NodeServiceStatusProperties CreateSyslogServiceActionRestart(ctx, )
Restart, start or stop the syslog service

Restart, start or stop the syslog service

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

# **CreateSyslogServiceActionStart**
> NodeServiceStatusProperties CreateSyslogServiceActionStart(ctx, )
Restart, start or stop the syslog service

Restart, start or stop the syslog service

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

# **CreateSyslogServiceActionStop**
> NodeServiceStatusProperties CreateSyslogServiceActionStop(ctx, )
Restart, start or stop the syslog service

Restart, start or stop the syslog service

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

# **CreateTransportNodeCollection**
> TransportNodeCollection CreateTransportNodeCollection(ctx, body, optional)
Create transport node collection by attaching Transport Node Profile to cluster.

When transport node collection is created the hosts which are part of compute collection will be prepared automatically i.e. NSX Manager attempts to install the NSX components on hosts. Transport nodes for these hosts are created using the configuration specified in transport node profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportNodeCollection**](TransportNodeCollection.md)|  | 
 **optional** | ***SystemAdministrationConfigurationApiCreateTransportNodeCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiCreateTransportNodeCollectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applyProfile** | **optional.**| Indicates if the Transport Node Profile (TNP) configuration should be applied during creation | [default to true]

### Return type

[**TransportNodeCollection**](TransportNodeCollection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransportNodeForDiscoveredNodeCreateTransportNode**
> TransportNode CreateTransportNodeForDiscoveredNodeCreateTransportNode(ctx, body, nodeExtId)
Created Transport Node for Discovered Node

NSX components are installaed on host and transport node is created with given configurations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportNode**](TransportNode.md)|  | 
  **nodeExtId** | **string**|  | 

### Return type

[**TransportNode**](TransportNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransportNodeProfile**
> TransportNodeProfile CreateTransportNodeProfile(ctx, body)
Create a Transport Node Profile

Transport node profile captures the configuration needed to create a transport node. A transport node profile can be attached to compute collections for automatic TN creation of member hosts. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportNodeProfile**](TransportNodeProfile.md)|  | 

### Return type

[**TransportNodeProfile**](TransportNodeProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransportNodeWithDeploymentInfo**
> TransportNode CreateTransportNodeWithDeploymentInfo(ctx, body)
Create a Transport Node

Transport nodes are hypervisor hosts and NSX Edges that will participate in an NSX-T overlay. For a hypervisor host, this means that it hosts VMs that will communicate over NSX-T logical switches. For NSX Edges, this means that it will have logical router uplinks and downlinks.  This API creates transport node for a host node (hypervisor) or edge node (router) in the transport network.  When you run this command for a host, NSX Manager attempts to install the NSX kernel modules, which are packaged as VIB, RPM, or DEB files. For the installation to succeed, you must provide the host login credentials and the host thumbprint.  To get the ESXi host thumbprint, SSH to the host and run the <b>openssl x509 -in /etc/vmware/ssl/rui.crt -fingerprint -sha256 -noout</b> command.  To generate host key thumbprint using SHA-256 algorithm please follow the steps below.  Log into the host, making sure that the connection is not vulnerable to a man in the middle attack. Check whether a public key already exists. Host public key is generally located at '/etc/ssh/ssh_host_rsa_key.pub'. If the key is not present then generate a new key by running the following command and follow the instructions.  <b>ssh-keygen -t rsa</b>  Now generate a SHA256 hash of the key using the following command. Please make sure to pass the appropriate file name if the public key is stored with a different file name other than the default 'id_rsa.pub'.  <b>awk '{print $2}' id_rsa.pub | base64 -d | sha256sum -b | sed 's/ .*$//' | xxd -r -p | base64</b> This api is deprecated as part of FN+TN unification. Please use Transport Node API to install NSX components on a node.  Additional documentation on creating a transport node can be found in the NSX-T Installation Guide.  In order for the transport node to forward packets, the host_switch_spec property must be specified.  Host switches (called bridges in OVS on KVM hypervisors) are the individual switches within the host virtual switch. Virtual machines are connected to the host switches.  When creating a transport node, you need to specify if the host switches are already manually preconfigured on the node, or if NSX should create and manage the host switches. You specify this choice by the type of host switches you pass in the host_switch_spec property of the TransportNode request payload.  For a KVM host, you can preconfigure the host switch, or you can have NSX Manager perform the configuration. For an ESXi host or NSX Edge node, NSX Manager always configures the host switch.  To preconfigure the host switches on a KVM host, pass an array of PreconfiguredHostSwitchSpec objects that describes those host switches. In the current NSX-T release, only one prefonfigured host switch can be specified.  See the PreconfiguredHostSwitchSpec schema definition for documentation on the properties that must be provided. Preconfigured host switches are only supported on KVM hosts, not on ESXi hosts or NSX Edge nodes.  To allow NSX to manage the host switch configuration on KVM hosts, ESXi hosts, or NSX Edge nodes, pass an array of StandardHostSwitchSpec objects in the host_switch_spec property, and NSX will automatically create host switches with the properties you provide. In the current NSX-T release, up to 16 host switches can be automatically managed. See the StandardHostSwitchSpec schema definition for documentation on the properties that must be provided.  Note: Previous versions of NSX-T also used a property named transport_zone_endpoints at TransportNode level. This property is deprecated which creates some combinations of new client along with old client payloads. Examples [1] & [2] show old/existing client request and response by populating transport_zone_endpoints property at TransportNode level. Example [3] shows TransportNode creation request/response by populating transport_zone_endpoints property at StandardHostSwitch level and other new properties.  The request should either provide node_deployement_info or node_id.  If the host node (hypervisor) or edge node (router) is already added in system then it can be converted to transport node by providing node_id in request.  If host node (hypervisor) or edge node (router) is not already present in system then information should be provided under node_deployment_info. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportNode**](TransportNode.md)|  | 

### Return type

[**TransportNode**](TransportNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransportZone**
> TransportZone CreateTransportZone(ctx, body)
Create a Transport Zone

Creates a new transport zone. The required parameters are host_switch_name and transport_type (OVERLAY or VLAN). The optional parameters are description and display_name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportZone**](TransportZone.md)|  | 

### Return type

[**TransportZone**](TransportZone.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTransportZoneProfile**
> TransportZoneProfile CreateTransportZoneProfile(ctx, body)
Create a transport zone Profile

Creates a transport zone profile. The resource_type is required. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportZoneProfile**](TransportZoneProfile.md)|  | 

### Return type

[**TransportZoneProfile**](TransportZoneProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVNIPool**
> VniPool CreateVNIPool(ctx, body)
Create a new VNI Pool.

Creates a new VNI pool using the specified VNI pool range. The range should be non-overlapping with an existing range. If the range in payload is present or overlaps with an existing range, return code 400 with bad request and an error message is returned mentioning that the given range overlaps with an existing range. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VniPool**](VniPool.md)|  | 

### Return type

[**VniPool**](VniPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DELETERabbitMQManagementPort**
> DELETERabbitMQManagementPort(ctx, )
Delete RabbitMQ management port

Delete RabbitMQ management port

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteApplianceManagementTask**
> DeleteApplianceManagementTask(ctx, taskId)
Delete task

Delete task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoDeployedClusterNodeVMDelete**
> DeleteAutoDeployedClusterNodeVMDelete(ctx, nodeId, optional)
Attempt to delete an auto-deployed cluster node VM

Attempts to unregister and undeploy a specified auto-deployed cluster node VM. If it is a member of a cluster, then the VM will be automatically detached from the cluster before being unregistered and undeployed. Alternatively, if the original deployment attempt failed or the VM is not found, cleans up the deployment information associated with the deployment attempt. Note: If a VM has been successfully auto-deployed, then the associated deployment information will not be deleted unless and until the VM is successfully deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiDeleteAutoDeployedClusterNodeVMDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDeleteAutoDeployedClusterNodeVMDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **optional.Bool**| Delete by force | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAutoDeployedPaceClusterNodeVMDelete**
> DeleteAutoDeployedPaceClusterNodeVMDelete(ctx, nodeId, optional)
Attempt to delete an auto-deployed cluster node VM

Attempts to unregister and undeploy a specified auto-deployed cluster node VM. If it is a member of a cluster, then the VM will be automatically detached from the cluster before being unregistered and undeployed. Alternatively, if the original deployment attempt failed or the VM is not found, cleans up the deployment information associated with the deployment attempt. Note: If a VM has been successfully auto-deployed, then the associated deployment information will not be deleted unless and until the VM is successfully deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiDeleteAutoDeployedPaceClusterNodeVMDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDeleteAutoDeployedPaceClusterNodeVMDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forceDelete** | **optional.Bool**| Delete by force | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterNodeConfig**
> DeleteClusterNodeConfig(ctx, nodeId)
Remove a controller from the cluster

Removes the specified controller from the NSX cluster. Before you can remove a controller from the cluster, you must shut down the controller service with the \"stop service controller\" command. Deprecated. Use POST /cluster/<node-id>?action=remove_node to detach a node from cluster. The controller is removed with the node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterProfile**
> DeleteClusterProfile(ctx, clusterProfileId)
Delete a cluster profile

Delete a specified cluster profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComputeCollectionFabricTemplate**
> DeleteComputeCollectionFabricTemplate(ctx, fabricTemplateId)
Deletes compute collection fabric template

Deletes compute collection fabric template for the given id. This functionality is deprecated. Use Transport Node Profiles instead of this template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fabricTemplateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteComputeManager**
> DeleteComputeManager(ctx, computeManagerId)
Unregister a compute manager

Unregisters a specified compute manager 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **computeManagerId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDirectoryDomain**
> DeleteDirectoryDomain(ctx, domainId, optional)
Delete a specific domain with given identifier

Delete a specific domain with given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
 **optional** | ***SystemAdministrationConfigurationApiDeleteDirectoryDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDeleteDirectoryDomainOpts struct
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

# **DeleteDirectoryLdapServer**
> DeleteDirectoryLdapServer(ctx, domainId, serverId)
Delete a LDAP server for directory domain

Delete a LDAP server for directory domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
  **serverId** | **string**| LDAP server identifier | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEdgeCluster**
> DeleteEdgeCluster(ctx, edgeClusterId)
Delete Edge Cluster

Deletes the specified edge cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFailureDomain**
> DeleteFailureDomain(ctx, failureDomainId)
Delete Failure Domain

Deletes an existing failure domain. You can not delete system generated default failure domain. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failureDomainId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHostSwitchProfile**
> DeleteHostSwitchProfile(ctx, hostSwitchProfileId)
Delete a Hostswitch Profile

Deletes a specified hostswitch profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostSwitchProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpBlock**
> DeleteIpBlock(ctx, blockId)
Delete an IP Address Block

Deletes the IP address block with specified id if it exists. IP block cannot be deleted if there are allocated subnets from the block. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockId** | **string**| IP address block id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpBlockSubnet**
> DeleteIpBlockSubnet(ctx, subnetId)
Delete subnet within an IP block

Deletes a subnet with specified id within a given IP address block. Deletion is allowed only when there are no allocated IP addresses from that subnet. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subnetId** | **string**| Subnet id | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteIpPool**
> DeleteIpPool(ctx, poolId, optional)
Delete an IP Pool

Deletes the specified IP address pool. By default, if the IpPool is used in other configurations (such as transport node template), it won't be deleted. In such situations, pass \"force=true\" as query param to force delete the IpPool

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| IP pool ID | 
 **optional** | ***SystemAdministrationConfigurationApiDeleteIpPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDeleteIpPoolOpts struct
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

# **DeleteMPAConfiguration**
> DeleteMPAConfiguration(ctx, )
Delete MPA configuration for this node

Delete the MPA configuration for this node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManagementPlaneConfiguration**
> DeleteManagementPlaneConfiguration(ctx, )
Delete management plane configuration for this node

Delete the management plane configuration for this node.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNode**
> DeleteNode(ctx, nodeId, optional)
Delete a Node

Removes a specified fabric node (host or edge). A fabric node may only be deleted when it is no longer referenced by a Transport Node. If unprepare_host option is set to false, the host will be deleted without uninstalling the NSX components from the host. This api is deprecated, use Transport Node API DELETE /transport-nodes/&lt;transport-node-id&gt; to delete FN. DELETE /transport-nodes/<transport-node-id> to delete FN. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiDeleteNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDeleteNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **unprepareHost** | **optional.Bool**| Delete a host and uninstall NSX components | [default to true]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNodeNetworkRoute**
> DeleteNodeNetworkRoute(ctx, routeId)
Delete node network route

Delete a route from the NSX Manager routing table. You can modify an existing route by deleting it and then posting the modified version of the route. To verify, remove the route ID from the URI, issue a GET request, and note the absense of the deleted route. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeId** | **string**| ID of route to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNodeSyslogExporter**
> DeleteNodeSyslogExporter(ctx, exporterName)
Delete node syslog exporter

Removes a specified rule from the collection of syslog exporter rules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exporterName** | **string**| Name of syslog exporter to delete | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransportNodeCollection**
> DeleteTransportNodeCollection(ctx, transportNodeCollectionId)
Detach transport node profile from compute collection.

By deleting transport node collection, we are detaching the transport node profile(TNP) from the compute collection. It has no effect on existing transport nodes. However, new hosts added to the compute collection will no longer be automatically converted to NSX transport node. Detaching TNP from compute collection does not delete TNP. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeCollectionId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransportNodeProfile**
> DeleteTransportNodeProfile(ctx, transportNodeProfileId)
Delete a Transport Node Profile

Deletes the specified transport node profile. A transport node profile can be deleted only when it is not attached to any compute collection. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransportNodeWithDeploymentInfo**
> DeleteTransportNodeWithDeploymentInfo(ctx, transportNodeId, optional)
Delete a Transport Node

Deletes the specified transport node. Query param force can be used to force delete the host nodes. Force deletion of edge and public cloud gateway nodes is not supported.  It also removes the specified node (host or edge) from system. If unprepare_host option is set to false, then host will be deleted without uninstalling the NSX components from the host. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiDeleteTransportNodeWithDeploymentInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDeleteTransportNodeWithDeploymentInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Force delete the resource even if it is being used somewhere  | [default to false]
 **unprepareHost** | **optional.Bool**| Uninstall NSX components from host while deleting | [default to true]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransportZone**
> DeleteTransportZone(ctx, zoneId)
Delete a Transport Zone

Deletes an existing transport zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTransportZoneProfile**
> DeleteTransportZoneProfile(ctx, transportzoneProfileId)
Delete a transport zone Profile

Deletes a specified transport zone profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportzoneProfileId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVNIPool**
> DeleteVNIPool(ctx, poolId, optional)
Delete a VNI Pool

Deletes the given VNI pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| VNI pool id | 
 **optional** | ***SystemAdministrationConfigurationApiDeleteVNIPoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDeleteVNIPoolOpts struct
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

# **DetachClusterNodeRemoveNode**
> ClusterConfiguration DetachClusterNodeRemoveNode(ctx, nodeId, optional)
Detach a node from the Cluster

Detach a node from the Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| UUID of the node | 
 **optional** | ***SystemAdministrationConfigurationApiDetachClusterNodeRemoveNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiDetachClusterNodeRemoveNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.String**|  | 
 **gracefulShutdown** | **optional.String**|  | [default to false]
 **ignoreRepositoryIpCheck** | **optional.String**|  | [default to false]

### Return type

[**ClusterConfiguration**](ClusterConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableFlowCacheDisableFlowCache**
> DisableFlowCacheDisableFlowCache(ctx, transportNodeId)
Disable flow cache for an edge transport node

Disable flow cache for edge transport node. Caution: This involves restart of the edge dataplane and hence may lead to network disruption. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableFlowCacheEnableFlowCache**
> EnableFlowCacheEnableFlowCache(ctx, transportNodeId)
Enable flow cache for an edge transport node

Enable flow cache for edge transport node. Caution: This involves restart of the edge dataplane and hence may lead to network disruption. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTransportNodesStatus**
> HeatMapTransportZoneStatus GetAllTransportNodesStatus(ctx, optional)
Get high-level summary of all transport nodes. The service layer does not support source = realtime or cached.

Get high-level summary of all transport nodes. The service layer does not support source = realtime or cached.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiGetAllTransportNodesStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetAllTransportNodesStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeType** | **optional.String**| Transport node type | 

### Return type

[**HeatMapTransportZoneStatus**](HeatMapTransportZoneStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTransportZoneStatus**
> HeatMapTransportNodesAggregateStatus GetAllTransportZoneStatus(ctx, )
Get high-level summary of a transport zone. The service layer does not support source = realtime or cached.

Get high-level summary of a transport zone. The service layer does not support source = realtime or cached.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**HeatMapTransportNodesAggregateStatus**](HeatMapTransportNodesAggregateStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetApiServiceConfig**
> ApiServiceConfig GetApiServiceConfig(ctx, )
Read API service properties

Read the configuration of the NSX API service. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ApiServiceConfig**](ApiServiceConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBackupUiFramesInfo**
> BackupUiFramesInfoList GetBackupUiFramesInfo(ctx, optional)
Get backup frames for UI

Returns list of backup frames and some metadata to be used by UI. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiGetBackupUiFramesInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetBackupUiFramesInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **uiTabType** | **optional.String**|  | [default to LOCAL_MANAGER_TAB]

### Return type

[**BackupUiFramesInfoList**](BackupUiFramesInfoList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleIds**
> BundleIds GetBundleIds(ctx, fileType, product)
Get list of bundle-ids which are available in repository or in-progress 

Get list of bundle-ids which are available in repository or in-progress 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fileType** | **string**| Type of file | 
  **product** | **string**| Name of the product | 

### Return type

[**BundleIds**](BundleIds.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleUploadPermissions**
> BundleUploadPermission GetBundleUploadPermissions(ctx, product)
Checks bundle upload permissions

Checks whether bundle upload is allowed on given node for given product. There are different kinds of checks for different products. Some of the checks for Intelligence product are as follows: 1. Is bundle upload-allowed on given node 2. Is bundle upload already in-progress 3. Is Intelliegnce node deployment in-progress 4. Is Intelliegnce node upgrade in-progress 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **product** | **string**| Name of the product | 

### Return type

[**BundleUploadPermission**](BundleUploadPermission.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBundleUploadStatus**
> BundleUploadStatus GetBundleUploadStatus(ctx, bundleId, product)
Get bundle upload status

Get uploaded bundle upload status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bundleId** | **string**|  | 
  **product** | **string**| Name of the product | 

### Return type

[**BundleUploadStatus**](BundleUploadStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCloudNativeServiceInstance**
> CloudNativeServiceInstance GetCloudNativeServiceInstance(ctx, externalId)
Returns information about a particular cloud native service instance by external-id. 

Returns information about a particular cloud native service instance by external-id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **externalId** | **string**|  | 

### Return type

[**CloudNativeServiceInstance**](CloudNativeServiceInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterCertificateId**
> ClusterCertificateId GetClusterCertificateId(ctx, )
Read cluster certificate ID

Returns the ID of the certificate that is used as the cluster certificate for MP 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterCertificateId**](ClusterCertificateId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterNodeConfig**
> ClusterNodeInfo GetClusterNodeConfig(ctx, nodeId)
Read cluster node configuration

Returns information about the specified NSX cluster node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**ClusterNodeInfo**](ClusterNodeInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterProfile**
> ClusterProfile GetClusterProfile(ctx, clusterProfileId)
Get cluster profile by Id

Returns information about a specified cluster profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterProfileId** | **string**|  | 

### Return type

[**ClusterProfile**](ClusterProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterVirtualIp**
> ClusterVirtualIpProperties GetClusterVirtualIp(ctx, )
Read cluster virtual IP address

Returns the configured cluster virtual IP address or null if not configured. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterVirtualIpProperties**](ClusterVirtualIpProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComputeCollectionFabricTemplate**
> ComputeCollectionFabricTemplate GetComputeCollectionFabricTemplate(ctx, fabricTemplateId)
Get compute collection fabric template by id

Get compute collection fabric template for the given id. This functionality is deprecated. Use Transport Node Profiles instead of this template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fabricTemplateId** | **string**|  | 

### Return type

[**ComputeCollectionFabricTemplate**](ComputeCollectionFabricTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetComputeManagerState**
> ConfigurationState GetComputeManagerState(ctx, computeManagerId)
Get the realized state of a compute manager

Get the realized state of a compute manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **computeManagerId** | **string**|  | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainerApplication**
> ContainerApplication GetContainerApplication(ctx, containerApplicationId)
Return a Container Application within a container project

Returns information about a specific Container Application within a project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerApplicationId** | **string**|  | 

### Return type

[**ContainerApplication**](ContainerApplication.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainerApplicationInstance**
> ContainerApplicationInstance GetContainerApplicationInstance(ctx, containerApplicationInstanceId)
Return a container application instance

Returns information about a specific container application instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerApplicationInstanceId** | **string**|  | 

### Return type

[**ContainerApplicationInstance**](ContainerApplicationInstance.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainerCluster**
> ContainerCluster GetContainerCluster(ctx, containerClusterId)
Return a container cluster

Returns information about a specific container cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerClusterId** | **string**|  | 

### Return type

[**ContainerCluster**](ContainerCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainerClusterNode**
> ContainerClusterNode GetContainerClusterNode(ctx, containerClusterNodeId)
Return a container cluster node

Returns information about a specific container cluster node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerClusterNodeId** | **string**|  | 

### Return type

[**ContainerClusterNode**](ContainerClusterNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainerIngressPolicy**
> ContainerIngressPolicy GetContainerIngressPolicy(ctx, ingressPolicyId)
Returns an ingress policy spec

Returns information about a specific ingress policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ingressPolicyId** | **string**|  | 

### Return type

[**ContainerIngressPolicy**](ContainerIngressPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainerNetworkPolicy**
> ContainerNetworkPolicy GetContainerNetworkPolicy(ctx, networkPolicyId)
Return a network policy spec

Returns information about a specific network policy.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkPolicyId** | **string**|  | 

### Return type

[**ContainerNetworkPolicy**](ContainerNetworkPolicy.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContainerProject**
> ContainerProject GetContainerProject(ctx, containerProjectId)
Return a container project

Returns information about a specific project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **containerProjectId** | **string**|  | 

### Return type

[**ContainerProject**](ContainerProject.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetControllerProfilerStatus**
> ControllerProfilerProperties GetControllerProfilerStatus(ctx, )
Get the status (Enabled/Disabled) of controller profiler

Get the status (Enabled/Disabled) of controller profiler

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ControllerProfilerProperties**](ControllerProfilerProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentBarrier**
> CurrentRealizationStateBarrier GetCurrentBarrier(ctx, )
Gets the current barrier number

Returns the current global realization barrier number for NSX. This method has been deprecated. To track realization state, use X-NSX-REQUESTID request header instead. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrentRealizationStateBarrier**](CurrentRealizationStateBarrier.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectoryDomain**
> DirectoryDomain GetDirectoryDomain(ctx, domainId)
Get a specific domain with given identifier

Get a specific domain with given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 

### Return type

[**DirectoryDomain**](DirectoryDomain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectoryDomainSyncStats**
> DirectoryDomainSyncStats GetDirectoryDomainSyncStats(ctx, domainId)
Get domain sync statistics for the given identifier

Get domain sync statistics for the given identifier

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 

### Return type

[**DirectoryDomainSyncStats**](DirectoryDomainSyncStats.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirectoryLdapServer**
> DirectoryLdapServer GetDirectoryLdapServer(ctx, domainId, serverId)
Get a specific LDAP server for a given directory domain

Get a specific LDAP server for a given directory domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
  **serverId** | **string**| LDAP server identifier | 

### Return type

[**DirectoryLdapServer**](DirectoryLdapServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeClusterAllocationStatus**
> EdgeClusterAllocationStatus GetEdgeClusterAllocationStatus(ctx, edgeClusterId)
Get the Allocation details of an edge cluster

Returns the allocation details of cluster and its members. Lists the edge node members, active and standby services of each node, utilization details of configured sub-pools. These allocation details can be monitored by customers to trigger migration of certain service contexts to different edge nodes, to balance the utilization of edge node resources. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 

### Return type

[**EdgeClusterAllocationStatus**](EdgeClusterAllocationStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeClusterInterSiteStatus**
> EdgeClusterInterSiteStatus GetEdgeClusterInterSiteStatus(ctx, edgeClusterId)
Get inter-site status of the edge cluster

Returns the aggregated status for the Edge cluster along with status of all edge nodes in the cluster. It always returns cached response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 

### Return type

[**EdgeClusterInterSiteStatus**](EdgeClusterInterSiteStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeClusterState**
> EdgeClusterState GetEdgeClusterState(ctx, edgeClusterId, optional)
Get the Realized State of a Edge Cluster

Return realized state information of a edge cluster. Any configuration update that affects the edge cluster can use this API to get its realized state by passing a request_id returned by the configuration change operation. e.g. Update configuration of edge cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiGetEdgeClusterStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetEdgeClusterStateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **barrierId** | **optional.Int64**|  | 
 **requestId** | **optional.String**| Realization request ID | 

### Return type

[**EdgeClusterState**](EdgeClusterState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEdgeClusterStatus**
> EdgeClusterStatus GetEdgeClusterStatus(ctx, edgeClusterId, optional)
Get the status for the Edge cluster of the given id

Returns the aggregated status for the Edge cluster along with status of all edge nodes in the cluster. Query parameter \"source=realtime\" is the only supported source. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiGetEdgeClusterStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetEdgeClusterStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**EdgeClusterStatus**](EdgeClusterStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFabricNodeModules**
> SoftwareModuleResult GetFabricNodeModules(ctx, nodeId)
Get the module details of a Fabric Node This api is deprecated, use Transport Node API GET /transport-nodes/&lt;transportnode-id&gt;/modules to get fabric node modules. 

Get the module details of a Fabric Node This api is deprecated, use Transport Node API GET /transport-nodes/&lt;transportnode-id&gt;/modules to get fabric node modules. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**SoftwareModuleResult**](SoftwareModuleResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFabricNodeModulesOfTransportNode**
> SoftwareModuleResult GetFabricNodeModulesOfTransportNode(ctx, nodeId)
Get the module details of a transport node 

Get the module details of a transport node 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**SoftwareModuleResult**](SoftwareModuleResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFabricNodeState**
> ConfigurationState GetFabricNodeState(ctx, nodeId)
Get the Realized State of a Fabric Node.

For edge nodes, returns the current install state when deployment is in progress, NODE_READY when deployment is complete and the failure state when deployment has failed. This api is deprecated. Please use /transport-nodes/&lt;transportnode-id&gt;/state to get realized state of a Fabric Node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**ConfigurationState**](ConfigurationState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFailureDomain**
> FailureDomain GetFailureDomain(ctx, failureDomainId)
Get a Failure Domain

Returns information about a single failure domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **failureDomainId** | **string**|  | 

### Return type

[**FailureDomain**](FailureDomain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalConfigs**
> GlobalConfigs GetGlobalConfigs(ctx, configType)
Get global configs for a config type

Returns global configurations that belong to the config type 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configType** | **string**|  | 

### Return type

[**GlobalConfigs**](GlobalConfigs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHeatmapTransportZoneStatus**
> HeatMapTransportZoneStatus GetHeatmapTransportZoneStatus(ctx, zoneId, optional)
Get high-level summary of a transport zone

Get high-level summary of a transport zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**| ID of transport zone | 
 **optional** | ***SystemAdministrationConfigurationApiGetHeatmapTransportZoneStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetHeatmapTransportZoneStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**HeatMapTransportZoneStatus**](HeatMapTransportZoneStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostNodeStatusOnComputeCollection**
> HostNodeStatusListResult GetHostNodeStatusOnComputeCollection(ctx, ccExtId)
Get status of member host nodes of the compute-collection. Only nsx prepared host nodes in the specified compute-collection are included in the response. cc-ext-id should be of type VC_Cluster.

Get status of member host nodes of the compute-collection. Only nsx prepared host nodes in the specified compute-collection are included in the response. cc-ext-id should be of type VC_Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ccExtId** | **string**|  | 

### Return type

[**HostNodeStatusListResult**](HostNodeStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHostSwitchProfile**
> BaseHostSwitchProfile GetHostSwitchProfile(ctx, hostSwitchProfileId)
Get a Hostswitch Profile by ID

Returns information about a specified hostswitch profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hostSwitchProfileId** | **string**|  | 

### Return type

[**BaseHostSwitchProfile**](BaseHostSwitchProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterSiteEdgeNodeBgpNeighborAdvertisedRoutes**
> BgpNeighborRouteDetails GetInterSiteEdgeNodeBgpNeighborAdvertisedRoutes(ctx, edgeNodeId, neighborId)
Get BGP neighbor advertised routes on edge transport node

Returns routes advertised by BGP neighbor from the given edge transport node. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeNodeId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**BgpNeighborRouteDetails**](BgpNeighborRouteDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterSiteEdgeNodeBgpNeighborRoutes**
> BgpNeighborRouteDetails GetInterSiteEdgeNodeBgpNeighborRoutes(ctx, edgeNodeId, neighborId)
Get BGP neighbor learned routes on edge transport node

Returns routes learned by BGP neighbor from the given edge transport node. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeNodeId** | **string**|  | 
  **neighborId** | **string**|  | 

### Return type

[**BgpNeighborRouteDetails**](BgpNeighborRouteDetails.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterSiteEdgeNodeBgpSummary**
> InterSiteBgpSummary GetInterSiteEdgeNodeBgpSummary(ctx, edgeNodeId)
Get inter-site BGP summary of edge node

Returns BGP summary for all configured neighbors in tunnel VRF on the given egde node. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeNodeId** | **string**|  | 

### Return type

[**InterSiteBgpSummary**](InterSiteBgpSummary.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInterSiteEdgeNodeStatistics**
> NodeInterSiteStatistics GetInterSiteEdgeNodeStatistics(ctx, edgeNodeId)
Get inter-site statistics of edge node

Returns RTEP to RTEP tunnel port statistics of the given edge node. It always returns realtime response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeNodeId** | **string**|  | 

### Return type

[**NodeInterSiteStatistics**](NodeInterSiteStatistics.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInventoryConfig**
> InventoryConfig GetInventoryConfig(ctx, )
Return inventory configuration

Supports retrieving following configuration of inventory module 1. Soft limit on number of compute managers that can be registered. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InventoryConfig**](InventoryConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeMandatoryAccessControl**
> MandatoryAccessControlProperties GetNodeMandatoryAccessControl(ctx, )
Gets the enable status for Mandatory Access Control

Gets the enable status for Mandatory Access Control

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MandatoryAccessControlProperties**](MandatoryAccessControlProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeMandatoryAccessControlReport**
> GetNodeMandatoryAccessControlReport(ctx, )
Get the report for Mandatory Access Control

Get the report for Mandatory Access Control

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNodeMode**
> NodeMode GetNodeMode(ctx, )
NodeMode

Returns current Node Mode. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeMode**](NodeMode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOvfDeployInfo**
> OvfInfo GetOvfDeployInfo(ctx, product)
Get information of the OVF which will be getting deployed. 

Get information of the OVF for specified product which is present in repository and will be used to deploy new VM. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **product** | **string**| Name of the product | 

### Return type

[**OvfInfo**](OvfInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaceHostConfiguration**
> IntelligenceHostConfigurationInfo GetPaceHostConfiguration(ctx, )
Get NSX-Intelligence host configuration

Get the current NSX-Intelligence host configuration. Recommend to keep the value same for flow_data_collection_interval and context_data_collection_interval. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IntelligenceHostConfigurationInfo**](IntelligenceHostConfigurationInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPhysicalServer**
> PhysicalServer GetPhysicalServer(ctx, physicalServerId)
Return a specific physical server

Returns information about physical/bare metal server based on given transport node id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **physicalServerId** | **string**|  | 

### Return type

[**PhysicalServer**](PhysicalServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPnicStatusesForTransportNode**
> PnicBondStatusListResult GetPnicStatusesForTransportNode(ctx, nodeId, optional)
Get high-level summary of a transport node

Get high-level summary of a transport node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| ID of transport node | 
 **optional** | ***SystemAdministrationConfigurationApiGetPnicStatusesForTransportNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetPnicStatusesForTransportNodeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **optional.String**| pNic/bond status | 

### Return type

[**PnicBondStatusListResult**](PnicBondStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRealizationStateBarrierConfig**
> RealizationStateBarrierConfig GetRealizationStateBarrierConfig(ctx, )
Gets the realization state barrier configuration

Returns the current barrier configuration 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RealizationStateBarrierConfig**](RealizationStateBarrierConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoSyncStatus**
> RepoSyncStatusReport GetRepoSyncStatus(ctx, nodeId)
Synchronizes the repository data between nsx managers.

Returns the synchronization status for the manager represented by given <node-id>. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**RepoSyncStatusReport**](RepoSyncStatusReport.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSupportedHostOSTypes**
> SupportedHostOsListResult GetSupportedHostOSTypes(ctx, )
Return list of supported host OS types

Returns names of all supported host OS.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SupportedHostOsListResult**](SupportedHostOSListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeCollection**
> TransportNodeCollection GetTransportNodeCollection(ctx, transportNodeCollectionId)
Get Transport Node collection by id

Returns transport node collection by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeCollectionId** | **string**|  | 

### Return type

[**TransportNodeCollection**](TransportNodeCollection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeCollectionState**
> TransportNodeCollectionState GetTransportNodeCollectionState(ctx, transportNodeCollectionId)
Get Transport Node collection application state

Returns the state of transport node collection based on the states of transport nodes of the hosts which are part of compute collection. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeCollectionId** | **string**|  | 

### Return type

[**TransportNodeCollectionState**](TransportNodeCollectionState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeProfile**
> TransportNodeProfile GetTransportNodeProfile(ctx, transportNodeProfileId)
Get a Transport Node

Returns information about a specified transport node profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeProfileId** | **string**|  | 

### Return type

[**TransportNodeProfile**](TransportNodeProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeReport**
> GetTransportNodeReport(ctx, optional)
Creates a status report of transport nodes of all the transport zones

You must provide the request header \"Accept:application/octet-stream\" when calling this API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiGetTransportNodeReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetTransportNodeReportOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **optional.String**| Data source type. | 
 **status** | **optional.String**| Transport node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeReportForATransportZone**
> GetTransportNodeReportForATransportZone(ctx, zoneId, optional)
Creates a status report of transport nodes in a transport zone

You must provide the request header \"Accept:application/octet-stream\" when calling this API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**| ID of transport zone | 
 **optional** | ***SystemAdministrationConfigurationApiGetTransportNodeReportForATransportZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetTransportNodeReportForATransportZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 
 **status** | **optional.String**| Transport node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeStateWithDeploymentInfo**
> TransportNodeState GetTransportNodeStateWithDeploymentInfo(ctx, transportNodeId)
Get a Transport Node's State

Returns information about the current state of the transport node configuration and information about the associated hostswitch. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

[**TransportNodeState**](TransportNodeState.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeStatus**
> TransportNodeStatus GetTransportNodeStatus(ctx, nodeId, optional)
Read status of a transport node

Read status of a transport node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| ID of transport node | 
 **optional** | ***SystemAdministrationConfigurationApiGetTransportNodeStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetTransportNodeStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**TransportNodeStatus**](TransportNodeStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportNodeWithDeploymentInfo**
> TransportNode GetTransportNodeWithDeploymentInfo(ctx, transportNodeId)
Get a Transport Node

Returns information about a specified transport node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

[**TransportNode**](TransportNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportZone**
> TransportZone GetTransportZone(ctx, zoneId)
Get a Transport Zone

Returns information about a single transport zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**|  | 

### Return type

[**TransportZone**](TransportZone.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportZoneProfile**
> TransportZoneProfile GetTransportZoneProfile(ctx, transportzoneProfileId)
Get transport zone profile by identifier

Returns information about a specified transport zone profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportzoneProfileId** | **string**|  | 

### Return type

[**TransportZoneProfile**](TransportZoneProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransportZoneStatus**
> TransportZoneStatus GetTransportZoneStatus(ctx, zoneId)
Get a Transport Zone's Current Runtime Status Information

Returns information about a specified transport zone, including the number of logical switches in the transport zone, number of logical spitch ports assigned to the transport zone, and number of transport nodes in the transport zone. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**|  | 

### Return type

[**TransportZoneStatus**](TransportZoneStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTunnel**
> TunnelProperties GetTunnel(ctx, nodeId, tunnelName, optional)
Tunnel properties

Tunnel properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| ID of transport node | 
  **tunnelName** | **string**| Tunnel name | 
 **optional** | ***SystemAdministrationConfigurationApiGetTunnelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiGetTunnelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**TunnelProperties**](TunnelProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HostPrepDiscoveredNodeHostprep**
> Node HostPrepDiscoveredNodeHostprep(ctx, nodeExtId)
(Deprecated) Prepares discovered Node for NSX

Prepares(hostprep) discovered node for NSX. NSX LCP bundles are installed on this discovered node. This API is deprecated. Use /fabric/discovered-nodes/<node-ext-id>?action=create_transport_node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeExtId** | **string**|  | 

### Return type

[**Node**](Node.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IncrementRealizationStateBarrierIncrement**
> CurrentRealizationStateBarrier IncrementRealizationStateBarrierIncrement(ctx, )
Increments the barrier count by 1

Increment the current barrier number by 1 for NSX. This method has been deprecated. To track realization state, use X-NSX-REQUESTID request header instead. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CurrentRealizationStateBarrier**](CurrentRealizationStateBarrier.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeDeleteClusterCentralAPI**
> InvokeDeleteClusterCentralAPI(ctx, targetNodeId, targetUri)
Invoke DELETE request on target cluster node

Invoke DELETE request on target cluster node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID or keyword self | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeDeleteFabricCentralAPI**
> InvokeDeleteFabricCentralAPI(ctx, targetNodeId, targetUri)
Invoke DELETE request on target fabric node

Invoke DELETE request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API DELETE /transport-nodes/&lt;transport-node-id&gt;/&lt;target-node-id&gt;/&lt;target-uri&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeDeleteTransportNodeCentralAPI**
> InvokeDeleteTransportNodeCentralAPI(ctx, targetNodeId, targetUri)
Invoke DELETE request on target transport node

Invoke DELETE request on target transport node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeGetClusterCentralAPI**
> InvokeGetClusterCentralAPI(ctx, targetNodeId, targetUri)
Invoke GET request on target cluster node

Invoke GET request on target cluster node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID or keyword self | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeGetFabricCentralAPI**
> InvokeGetFabricCentralAPI(ctx, targetNodeId, targetUri)
Invoke GET request on target fabric node

Invoke GET request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API GET /transport-nodes/&lt;transport-node-id&gt;/&lt;target-node-id&gt;/&lt;target-uri&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokeGetTransportNodeCentralAPI**
> InvokeGetTransportNodeCentralAPI(ctx, targetNodeId, targetUri)
Invoke GET request on target transport node

Invoke GET request on target transport node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePostClusterCentralAPI**
> InvokePostClusterCentralAPI(ctx, targetNodeId, targetUri)
Invoke POST request on target cluster node

Invoke POST request on target cluster node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID or keyword self | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePostFabricCentralAPI**
> InvokePostFabricCentralAPI(ctx, targetNodeId, targetUri)
Invoke POST request on target fabric node

Invoke POST request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API POST /transport-nodes/&lt;transport-node-id&gt;/&lt;target-node-id&gt;/&lt;target-uri&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePostTransportNodeCentralAPI**
> InvokePostTransportNodeCentralAPI(ctx, targetNodeId, targetUri)
Invoke POST request on target transport node

Invoke POST request on target transport node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePutClusterCentralAPI**
> InvokePutClusterCentralAPI(ctx, targetNodeId, targetUri)
Invoke PUT request on target cluster node

Invoke PUT request on target cluster node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID or keyword self | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePutFabricCentralAPI**
> InvokePutFabricCentralAPI(ctx, targetNodeId, targetUri)
Invoke PUT request on target fabric node

Invoke PUT request on target fabric node. This api is deprecated as part of FN+TN unification. Please use Transport Node API PUT /transport-nodes/&lt;transport-node-id&gt;/&lt;target-node-id&gt;/&lt;target-uri&gt; 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InvokePutTransportNodeCentralAPI**
> InvokePutTransportNodeCentralAPI(ctx, targetNodeId, targetUri)
Invoke PUT request on target transport node

Invoke PUT request on target transport node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetNodeId** | **string**| Target node UUID | 
  **targetUri** | **string**| URI of API to invoke on target node | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **JoinClusterJoinCluster**
> ClusterConfiguration JoinClusterJoinCluster(ctx, body)
Join this node to a NSX Cluster

Join this node to a NSX Cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JoinClusterParameters**](JoinClusterParameters.md)|  | 

### Return type

[**ClusterConfiguration**](ClusterConfiguration.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllCloudNativeServiceInstances**
> CloudNativeServiceInstanceListResult ListAllCloudNativeServiceInstances(ctx, optional)
Returns the List of cloud native service instances

Returns information about all cloud native service instances.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListAllCloudNativeServiceInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListAllCloudNativeServiceInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **displayName** | **optional.String**| Display Name of the cloud native service instance | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **serviceType** | **optional.String**| Type of cloud native service; possible values are ELB, RDS | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| NSX node id of the public cloud gateway that reported the service instance | 

### Return type

[**CloudNativeServiceInstanceListResult**](CloudNativeServiceInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListApplianceManagementTasks**
> ApplianceManagementTaskListResult ListApplianceManagementTasks(ctx, optional)
List appliance management tasks

List appliance management tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListApplianceManagementTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListApplianceManagementTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **optional.String**| Fields to include in query results | 
 **requestMethod** | **optional.String**| Request method(s) to include in query result | 
 **requestPath** | **optional.String**| Request URI path(s) to include in query result | 
 **requestUri** | **optional.String**| Request URI(s) to include in query result | 
 **status** | **optional.String**| Status(es) to include in query result | 
 **user** | **optional.String**| Names of users to include in query result | 

### Return type

[**ApplianceManagementTaskListResult**](ApplianceManagementTaskListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCentralNodeConfigProfiles**
> CentralNodeConfigProfileListResult ListCentralNodeConfigProfiles(ctx, )
List all Central Node Config profiles

Returns list of all Central Node Config profiles. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CentralNodeConfigProfileListResult**](CentralNodeConfigProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterNodeConfigs**
> ClusterNodeConfigListResult ListClusterNodeConfigs(ctx, optional)
List Cluster Node Configurations

Returns information about all NSX cluster nodes. Deprecated. Use GET /cluster to get cluster configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListClusterNodeConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListClusterNodeConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ClusterNodeConfigListResult**](ClusterNodeConfigListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterNodeInterfaces**
> NodeInterfacePropertiesListResult ListClusterNodeInterfaces(ctx, nodeId, optional)
List the specified node's Network Interfaces

Returns the number of interfaces on the node and detailed information about each interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network mask, and the IP configuration method (static or DHCP). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiListClusterNodeInterfacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListClusterNodeInterfacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfacePropertiesListResult**](NodeInterfacePropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterNodeVMDeploymentRequests**
> ClusterNodeVmDeploymentRequestList ListClusterNodeVMDeploymentRequests(ctx, )
Returns info for all cluster node VM auto-deployment attempts

Returns request information for every attempted deployment of a cluster node VM. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterNodeVmDeploymentRequestList**](ClusterNodeVMDeploymentRequestList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterProfiles**
> ClusterProfileListResult ListClusterProfiles(ctx, optional)
List Cluster Profiles

Returns paginated list of cluster profiles Cluster profiles define policies for edge cluster and bridge cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListClusterProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListClusterProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeSystemOwned** | **optional.Bool**| Whether the list result contains system resources | [default to true]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **resourceType** | **optional.String**| Supported cluster profiles. | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ClusterProfileListResult**](ClusterProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListComputeCollectionFabricTemplates**
> ComputeCollectionFabricTemplateListResult ListComputeCollectionFabricTemplates(ctx, optional)
Get compute collection fabric templates

Returns compute collection fabric templates. This functionality is deprecated. Use Transport Node Profiles instead of this template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListComputeCollectionFabricTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListComputeCollectionFabricTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **computeCollectionId** | **optional.String**| Compute collection id | 

### Return type

[**ComputeCollectionFabricTemplateListResult**](ComputeCollectionFabricTemplateListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListComputeCollectionPhysicalNetworkInterfaces**
> ComputeCollectionNetworkInterfacesListResult ListComputeCollectionPhysicalNetworkInterfaces(ctx, ccExtId)
List the Physical Network Interface for all discovered nodes

Returns list of physical network interfaces for all discovered nodes in compute collection. Interface information includes PNIC name, hostswitch name it's attached to(if any) and MAC address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ccExtId** | **string**|  | 

### Return type

[**ComputeCollectionNetworkInterfacesListResult**](ComputeCollectionNetworkInterfacesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListComputeCollections**
> ComputeCollectionListResult ListComputeCollections(ctx, optional)
Return the List of Compute Collections

Returns information about all compute collections.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListComputeCollectionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListComputeCollectionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmLocalId** | **optional.String**| Local Id of the compute collection in the Compute Manager | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **discoveredNodeId** | **optional.String**| Id of the discovered node which belongs to this Compute Collection  | 
 **displayName** | **optional.String**| Name of the ComputeCollection in source compute manager | 
 **externalId** | **optional.String**| External ID of the ComputeCollection in the source Compute manager, e.g. mo-ref in VC  | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **nodeId** | **optional.String**| Id of the fabric node created from a discovered node belonging to this Compute Collection  | 
 **originId** | **optional.String**| Id of the compute manager from where this Compute Collection was discovered | 
 **originType** | **optional.String**| ComputeCollection type like VC_Cluster. Here the Compute Manager type prefix would help in differentiating similar named Compute Collection types from different Compute Managers  | 
 **ownerId** | **optional.String**| Id of the owner of compute collection in the Compute Manager | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ComputeCollectionListResult**](ComputeCollectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListComputeManagers**
> ComputeManagerListResult ListComputeManagers(ctx, optional)
Return the List of Compute managers

Returns information about all compute managers.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListComputeManagersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListComputeManagersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **originType** | **optional.String**| Compute manager type like vCenter | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **server** | **optional.String**| IP address or hostname of compute manager | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ComputeManagerListResult**](ComputeManagerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerApplicationInstances**
> ContainerApplicationInstanceListResult ListContainerApplicationInstances(ctx, optional)
Return the list of container application instance

Returns information about all container application instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListContainerApplicationInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListContainerApplicationInstancesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerApplicationId** | **optional.String**| Identifier of the container application | 
 **containerClusterId** | **optional.String**| Identifier of the container cluster | 
 **containerProjectId** | **optional.String**| Identifier of the container project | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerApplicationInstanceListResult**](ContainerApplicationInstanceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerApplications**
> ContainerApplicationListResult ListContainerApplications(ctx, optional)
Return the List of Container Applications

Returns information about all Container Applications.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListContainerApplicationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListContainerApplicationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerClusterId** | **optional.String**| Identifier of the container cluster | 
 **containerProjectId** | **optional.String**| Identifier of the container project | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerApplicationListResult**](ContainerApplicationListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerClusterNodes**
> ContainerClusterNodeListResult ListContainerClusterNodes(ctx, optional)
Return the list of container cluster nodes

Returns information about all container cluster nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListContainerClusterNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListContainerClusterNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerClusterId** | **optional.String**| Identifier of the container cluster | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerClusterNodeListResult**](ContainerClusterNodeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerClusters**
> ContainerClusterListResult ListContainerClusters(ctx, optional)
Return the List of Container Clusters

Returns information about all Container Clusters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListContainerClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListContainerClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterType** | **optional.String**| Type of container cluster | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **infraType** | **optional.String**| Type of infrastructure | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerClusterListResult**](ContainerClusterListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerIngressPolicies**
> ContainerIngressPolicyListResult ListContainerIngressPolicies(ctx, optional)
Return the List of Container Ingress Policies

Returns information about all ingress policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListContainerIngressPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListContainerIngressPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerClusterId** | **optional.String**| Identifier of the container cluster | 
 **containerProjectId** | **optional.String**| Identifier of the container project | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerIngressPolicyListResult**](ContainerIngressPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerNetworkPolicies**
> ContainerNetworkPolicyListResult ListContainerNetworkPolicies(ctx, optional)
Return the List of Container Network Policies

Returns information about all network policies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListContainerNetworkPoliciesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListContainerNetworkPoliciesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerClusterId** | **optional.String**| Identifier of the container cluster | 
 **containerProjectId** | **optional.String**| Identifier of the container project | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerNetworkPolicyListResult**](ContainerNetworkPolicyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListContainerProjects**
> ContainerProjectListResult ListContainerProjects(ctx, optional)
Return the list of container projects

Returns information about all container projects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListContainerProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListContainerProjectsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerClusterId** | **optional.String**| Identifier of the container cluster | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**ContainerProjectListResult**](ContainerProjectListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDirectoryDomains**
> DirectoryDomainListResults ListDirectoryDomains(ctx, optional)
List all configured domains

List all configured domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListDirectoryDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListDirectoryDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DirectoryDomainListResults**](DirectoryDomainListResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDirectoryGroupMemberGroups**
> DirectoryGroupMemberListResults ListDirectoryGroupMemberGroups(ctx, domainId, groupId, optional)
List members of a directory group

A member group could be either direct member of the group specified by group_id or nested member of it. Both direct member groups and nested member groups are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
  **groupId** | **string**| Directory group identifier | 
 **optional** | ***SystemAdministrationConfigurationApiListDirectoryGroupMemberGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListDirectoryGroupMemberGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DirectoryGroupMemberListResults**](DirectoryGroupMemberListResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDirectoryLdapServers**
> DirectoryLdapServerListResults ListDirectoryLdapServers(ctx, domainId, optional)
List all configured domain LDAP servers

List all configured domain LDAP servers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
 **optional** | ***SystemAdministrationConfigurationApiListDirectoryLdapServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListDirectoryLdapServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DirectoryLdapServerListResults**](DirectoryLdapServerListResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiscoveredNodes**
> DiscoveredNodeListResult ListDiscoveredNodes(ctx, optional)
Return the List of Discovered Nodes

Returns information about all discovered nodes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListDiscoveredNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListDiscoveredNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmLocalId** | **optional.String**| Local Id of the discovered node in the Compute Manager | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **displayName** | **optional.String**| Display name of discovered node | 
 **externalId** | **optional.String**| External id of the discovered node, ex. a mo-ref from VC | 
 **hasParent** | **optional.String**| Discovered node has a parent compute collection or is a standalone host | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **ipAddress** | **optional.String**| IP address of the discovered node | 
 **nodeId** | **optional.String**| Id of the fabric node created from the discovered node | 
 **nodeType** | **optional.String**| Discovered Node type like HostNode | 
 **originId** | **optional.String**| Id of the compute manager from where this node was discovered | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **parentComputeCollection** | **optional.String**| External id of the compute collection to which this node belongs | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DiscoveredNodeListResult**](DiscoveredNodeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEdgeClusters**
> EdgeClusterListResult ListEdgeClusters(ctx, optional)
List Edge Clusters

Returns information about the configured edge clusters, which enable you to group together transport nodes of the type EdgeNode and apply fabric profiles to all members of the edge cluster. Each edge node can participate in only one edge cluster. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListEdgeClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListEdgeClustersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**EdgeClusterListResult**](EdgeClusterListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFabricNodeInterfaces**
> NodeInterfacePropertiesListResult ListFabricNodeInterfaces(ctx, nodeId, optional)
List the specified node's Network Interfaces

Returns the number of interfaces on the node and detailed information about each interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network mask, and the IP configuration method (static or DHCP). This api is deprecated. Please use Transport Node API GET /transport-nodes/<transport-node-id>/network/interfaces to list node network interfaces for the corresponding TN. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiListFabricNodeInterfacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListFabricNodeInterfacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminStatus** | **optional.String**| Admin status of the interface | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfacePropertiesListResult**](NodeInterfacePropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFabricNodeNeighborProperties**
> InterfaceNeighborPropertyListResult ListFabricNodeNeighborProperties(ctx, fabricNodeId)
List LLDP Neighbor Properties of Fabric Node

List LLDP Neighbor Properties for all interfaces of Fabric Node 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fabricNodeId** | **string**| ID of fabric node | 

### Return type

[**InterfaceNeighborPropertyListResult**](InterfaceNeighborPropertyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFailureDomains**
> FailureDomainListResult ListFailureDomains(ctx, )
List Failure Domains

Returns information about configured failure domains.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**FailureDomainListResult**](FailureDomainListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGlobalConfigs**
> GlobalConfigsListResult ListGlobalConfigs(ctx, )
List global configurations of a NSX domain

Returns global configurations of a NSX domain grouped by the config types. These global configurations are valid across NSX domain for their respective types unless they are overridden by a more granular configurations. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GlobalConfigsListResult**](GlobalConfigsListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListHostSwitchProfiles**
> HostSwitchProfilesListResult ListHostSwitchProfiles(ctx, optional)
List Hostswitch Profiles

Returns information about the configured hostswitch profiles. Hostswitch profiles define networking policies for hostswitches (sometimes referred to as bridges in OVS). Currently, only uplink teaming is supported. Uplink teaming allows NSX to load balance traffic across different physical NICs (PNICs) on the hypervisor hosts. Multiple teaming policies are supported, including LACP active, LACP passive, load balancing based on source ID, and failover order. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListHostSwitchProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListHostSwitchProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **deploymentType** | **optional.String**| Supported edge deployment type. | 
 **hostswitchProfileType** | **optional.String**| Supported HostSwitch profiles. | 
 **includeSystemOwned** | **optional.Bool**| Whether the list result contains system resources | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **nodeType** | **optional.String**| Fabric node type for which uplink profiles are to be listed | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **uplinkTeamingPolicyName** | **optional.String**| The host switch profile&#x27;s uplink teaming policy name | 

### Return type

[**HostSwitchProfilesListResult**](HostSwitchProfilesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIntelligenceFormFactors**
> IntelligenceFormFactors ListIntelligenceFormFactors(ctx, )
List available NSX Intelligence appliance form factors

Returns information about all form factors available for intelligence nodes 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IntelligenceFormFactors**](IntelligenceFormFactors.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInterSiteEdgeNodeBgpNeighbors**
> BgpNeighborListResult ListInterSiteEdgeNodeBgpNeighbors(ctx, edgeNodeId, optional)
Paginated list of BGP Neighbors on edge transport node

Paginated list of BGP Neighbors on edge transport node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeNodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiListInterSiteEdgeNodeBgpNeighborsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListInterSiteEdgeNodeBgpNeighborsOpts struct
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

# **ListIpBlockSubnets**
> IpBlockSubnetListResult ListIpBlockSubnets(ctx, optional)
List subnets within an IP block

Returns information about all subnets present within an IP address block. Information includes subnet's id, display_name, description, cidr and allocation ranges. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListIpBlockSubnetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListIpBlockSubnetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockId** | **optional.String**|  | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpBlockSubnetListResult**](IpBlockSubnetListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpBlocks**
> IpBlockListResult ListIpBlocks(ctx, optional)
Returns list of configured IP address blocks.

Returns information about configured IP address blocks. Information includes the id, display name, description & CIDR of IP address blocks 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListIpBlocksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListIpBlocksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpBlockListResult**](IpBlockListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpPoolAllocations**
> AllocationIpAddressListResult ListIpPoolAllocations(ctx, poolId)
List IP Pool Allocations

Returns information about which addresses have been allocated from a specified IP address pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| IP pool ID | 

### Return type

[**AllocationIpAddressListResult**](AllocationIpAddressListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListIpPools**
> IpPoolListResult ListIpPools(ctx, optional)
List IP Pools

Returns information about the configured IP address pools. Information includes the display name and description of the pool and the details of each of the subnets in the pool, including the DNS servers, allocation ranges, gateway, and CIDR subnet address. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListIpPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListIpPoolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**IpPoolListResult**](IpPoolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMacPools**
> MacPoolListResult ListMacPools(ctx, optional)
List MAC Pools

Returns a list of all the MAC pools 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListMacPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListMacPoolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**MacPoolListResult**](MacPoolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNeighborProperties**
> InterfaceNeighborPropertyListResult ListNeighborProperties(ctx, nodeId)
List LLDP Neighbor Properties of Transport Node

List LLDP Neighbor Properties for all interfaces of Transport Node 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| ID of transport node | 

### Return type

[**InterfaceNeighborPropertyListResult**](InterfaceNeighborPropertyListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeCapabilities**
> NodeCapabilitiesResult ListNodeCapabilities(ctx, nodeId)
Return the List of Capabilities of a Single Node

Returns information about capabilities of a single fabric host node. Edge nodes do not have capabilities. This api is deprecated, use GET /transport-nodes/&lt;transportnode-id&gt;/capabilities if FN is converted to TN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**NodeCapabilitiesResult**](NodeCapabilitiesResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeInterfaces**
> NodeNetworkInterfacePropertiesListResult ListNodeInterfaces(ctx, )
List the NSX Manager's Network Interfaces

Returns the number of interfaces on the NSX Manager appliance and detailed information about each interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network mask, and the IP configuration method (static or DHCP). 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeNetworkInterfacePropertiesListResult**](NodeNetworkInterfacePropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeNetworkRoutes**
> NodeRoutePropertiesListResult ListNodeNetworkRoutes(ctx, )
List node network routes

Returns detailed information about each route in the NSX Manager routing table. Route information includes the route type (default, static, and so on), a unique route identifier, the route metric, the protocol from which the route was learned, the route source (which is the preferred egress interface), the route destination, and the route scope. The route scope refers to the distance to the destination network: The \"host\" scope leads to a destination address on the NSX Manager, such as a loopback address; the \"link\" scope leads to a destination on the local network; and the \"global\" scope leads to addresses that are more than one hop away. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeRoutePropertiesListResult**](NodeRoutePropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeProcesses**
> NodeProcessPropertiesListResult ListNodeProcesses(ctx, )
List node processes

Returns the number of processes and information about each process. Process information includes 1) mem_resident, which is roughly equivalent to the amount of RAM, in bytes, currently used by the process, 2) parent process ID (ppid), 3) process name, 4) process up time in milliseconds, 5) mem_used, wich is the amount of virtual memory used by the process, in bytes, 6) process start time, in milliseconds since epoch, 7) process ID (pid), 8) CPU time, both user and the system, consumed by the process in milliseconds. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeProcessPropertiesListResult**](NodeProcessPropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeServices**
> NodeServicePropertiesListResult ListNodeServices(ctx, )
List node services

Returns a list of all services available on the NSX Manager applicance. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeServicePropertiesListResult**](NodeServicePropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodeSyslogExporters**
> NodeSyslogExporterPropertiesListResult ListNodeSyslogExporters(ctx, )
List node syslog exporters

Returns the collection of registered syslog exporter rules, if any. The rules specify the collector IP address and port, and the protocol to use. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeSyslogExporterPropertiesListResult**](NodeSyslogExporterPropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNodes**
> NodeListResult ListNodes(ctx, optional)
Return the List of Nodes

Returns information about all fabric nodes (hosts and edges). This api is deprecated as part of FN+TN unification. Please use Transport Node API GET /transport-nodes to list all fabric nodes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListNodesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **discoveredNodeId** | **optional.String**| Id of the discovered node which was converted to create this node | 
 **displayName** | **optional.String**| HostNode display name | 
 **externalId** | **optional.String**| HostNode external id | 
 **hardwareId** | **optional.String**| Hardware Id of the host | 
 **hypervisorOsType** | **optional.String**| HostNode&#x27;s Hypervisor type, for example ESXi, RHEL KVM or UBUNTU KVM. | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **ipAddress** | **optional.String**| Management IP address of the node | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **resourceType** | **optional.String**| Node type from &#x27;HostNode&#x27;, &#x27;EdgeNode&#x27;, &#x27;PublicCloudGatewayNode&#x27; | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**NodeListResult**](NodeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaceClusterNodeVMDeploymentRequests**
> IntelligenceClusterNodeVmDeploymentRequestList ListPaceClusterNodeVMDeploymentRequests(ctx, )
Returns info for all cluster node VM auto-deployment attempts

Returns request information for every attempted deployment of a cluster node VM. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IntelligenceClusterNodeVmDeploymentRequestList**](IntelligenceClusterNodeVMDeploymentRequestList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPhysicalServers**
> PhysicalServerListResult ListPhysicalServers(ctx, optional)
Return the list of physical servers

Returns information of all physical/bare metal servers registered as TN.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListPhysicalServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListPhysicalServersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **displayName** | **optional.String**| Display Name of the physical server | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **osType** | **optional.String**| OS type of the physical server | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**PhysicalServerListResult**](PhysicalServerListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRemoteTransportNodeStatus**
> TransportNodeStatusListResult ListRemoteTransportNodeStatus(ctx, nodeId, optional)
Read status of all transport nodes with tunnel connections to transport node 

Read status of all transport nodes with tunnel connections to transport node 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| ID of transport node | 
 **optional** | ***SystemAdministrationConfigurationApiListRemoteTransportNodeStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListRemoteTransportNodeStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bfdDiagnosticCode** | **optional.String**| BFD diagnostic code of Tunnel | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **tunnelStatus** | **optional.String**| Tunnel Status | 

### Return type

[**TransportNodeStatusListResult**](TransportNodeStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTasks**
> TaskListResult ListTasks(ctx, optional)
Get information about all tasks

Get information about all tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListTasksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTasksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **requestUri** | **optional.String**| Request URI(s) to include in query result | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **status** | **optional.String**| Status(es) to include in query result | 
 **user** | **optional.String**| Names of users to include in query result | 

### Return type

[**TaskListResult**](TaskListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodeCapabilities**
> NodeCapabilitiesResult ListTransportNodeCapabilities(ctx, transportNodeId)
Return the list of capabilities of transport node

Returns information about capabilities of transport host node. Edge nodes do not have capabilities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

[**NodeCapabilitiesResult**](NodeCapabilitiesResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodeCollections**
> TransportNodeCollectionListResult ListTransportNodeCollections(ctx, )
List Transport Node collections

Returns all Transport Node collections

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TransportNodeCollectionListResult**](TransportNodeCollectionListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodeInterfaces**
> NodeInterfacePropertiesListResult ListTransportNodeInterfaces(ctx, transportNodeId, optional)
List the specified transport node's network interfaces

Returns the number of interfaces on the node and detailed information about each interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network mask, and the IP configuration method (static or DHCP). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiListTransportNodeInterfacesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportNodeInterfacesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **adminStatus** | **optional.String**| Admin status of the interface | 
 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfacePropertiesListResult**](NodeInterfacePropertiesListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodeProfiles**
> TransportNodeProfileListResult ListTransportNodeProfiles(ctx, optional)
List Transport Nodes

Returns information about all transport node profiles. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListTransportNodeProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportNodeProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**TransportNodeProfileListResult**](TransportNodeProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodeStatus**
> TransportNodeStatusListResult ListTransportNodeStatus(ctx, optional)
Read status of all the transport nodes

Read status of all the transport nodes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListTransportNodeStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportNodeStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **status** | **optional.String**| Transport node | 

### Return type

[**TransportNodeStatusListResult**](TransportNodeStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodeStatusForTransportZone**
> TransportNodeStatusListResult ListTransportNodeStatusForTransportZone(ctx, zoneId, optional)
Read status of transport nodes in a transport zone

Read status of transport nodes in a transport zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **zoneId** | **string**| ID of transport zone | 
 **optional** | ***SystemAdministrationConfigurationApiListTransportNodeStatusForTransportZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportNodeStatusForTransportZoneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **status** | **optional.String**| Transport node | 

### Return type

[**TransportNodeStatusListResult**](TransportNodeStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodesByStateWithDeploymentInfo**
> TransportNodeStateListResult ListTransportNodesByStateWithDeploymentInfo(ctx, optional)
List transport nodes by realized state

Returns a list of transport node states that have realized state as provided as query parameter 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListTransportNodesByStateWithDeploymentInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportNodesByStateWithDeploymentInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mmState** | **optional.String**| maintenance mode state | 
 **status** | **optional.String**| Realized state of transport nodes | 
 **vtepIp** | **optional.String**| Virtual tunnel endpoint ip address of transport node | 

### Return type

[**TransportNodeStateListResult**](TransportNodeStateListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportNodesWithDeploymentInfo**
> TransportNodeListResult ListTransportNodesWithDeploymentInfo(ctx, optional)
List Transport Nodes

Returns information about all transport nodes along with underlying host or edge details. A transport node is a host or edge that contains hostswitches. A hostswitch can have virtual machines connected to them.  Because each transport node has hostswitches, transport nodes can also have virtual tunnel endpoints, which means that they can be part of the overlay. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListTransportNodesWithDeploymentInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportNodesWithDeploymentInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **inMaintenanceMode** | **optional.Bool**| maintenance mode flag | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **nodeId** | **optional.String**| node identifier | 
 **nodeIp** | **optional.String**| Fabric node IP address | 
 **nodeTypes** | **optional.String**| a list of fabric node types separated by comma or a single type | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **transportZoneId** | **optional.String**| Transport zone identifier | 

### Return type

[**TransportNodeListResult**](TransportNodeListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportZoneProfiles**
> TransportZoneProfileListResult ListTransportZoneProfiles(ctx, optional)
List transport zone profiles

Returns information about the configured transport zone profiles. Transport zone profiles define networking policies for transport zones and transport zone endpoints. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListTransportZoneProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportZoneProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeSystemOwned** | **optional.Bool**| Whether the list result contains system resources | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **resourceType** | **optional.String**| comma-separated list of transport zone profile types, e.g. ?resource_type&#x3D;BfdHealthMonitoringProfile | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**TransportZoneProfileListResult**](TransportZoneProfileListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransportZones**
> TransportZoneListResult ListTransportZones(ctx, optional)
List Transport Zones

Returns information about configured transport zones. NSX requires at least one transport zone. NSX uses transport zones to provide connectivity based on the topology of the underlying network, trust zones, or organizational separations. For example, you might have hypervisors that use one network for management traffic and a different network for VM traffic. This architecture would require two transport zones. The combination of transport zones plus transport connectors enables NSX to form tunnels between hypervisors. Transport zones define which interfaces on the hypervisors can communicate with which other interfaces on other hypervisors to establish overlay tunnels or provide connectivity to a VLAN. A logical switch can be in one (and only one) transport zone. This means that all of a switch's interfaces must be in the same transport zone. However, each hypervisor virtual switch (OVS or VDS) has multiple interfaces (connectors), and each connector can be attached to a different logical switch. For example, on a single hypervisor with two connectors, connector A can be attached to logical switch 1 in transport zone A, while connector B is attached to logical switch 2 in transport zone B. In this way, a single hypervisor can participate in multiple transport zones. The API for creating a transport zone requires that a single host switch be specified for each transport zone, and multiple transport zones can share the same host switch. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListTransportZonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListTransportZonesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includeSystemOwned** | **optional.Bool**| Filter to indicate whether to include system owned Transport Zones. | [default to false]
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **isDefault** | **optional.Bool**| Filter to choose if default transport zones will be returned | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **transportType** | **optional.String**| Filter to choose the type of transport zones to return | 
 **uplinkTeamingPolicyName** | **optional.String**| The transport zone&#x27;s uplink teaming policy name | 

### Return type

[**TransportZoneListResult**](TransportZoneListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVNIPools**
> VniPoolListResult ListVNIPools(ctx, optional)
List VNI Pools

Returns information about the default and configured virtual network identifier (VNI) pools for use when building logical network segments. Each virtual network has a unique ID called a VNI. Instead of creating a new VNI each time you need a new logical switch, you can instead allocate a VNI from a VNI pool. VNI pools are sometimes called segment ID pools. Each VNI pool has a range of usable VNIs. By default, there is one pool with two ranges [5000, 65535] and [65536, 75000]. To create multiple smaller pools, specify a smaller range for each pool such as 75001-75100 and 75101-75200. The VNI range determines the maximum number of logical switches that can be created in each network segment. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListVNIPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListVNIPoolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VniPoolListResult**](VniPoolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVifs**
> VirtualNetworkInterfaceListResult ListVifs(ctx, optional)
Return the List of Virtual Network Interfaces (VIFs)

Returns information about all VIFs. A virtual network interface aggregates network interfaces into a logical interface unit that is indistinuishable from a physical network interface. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListVifsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListVifsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **hostId** | **optional.String**| Id of the host where this vif is located. | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **lportAttachmentId** | **optional.String**| LPort Attachment Id of the virtual network interface. | 
 **ownerVmId** | **optional.String**| External id of the virtual machine. | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **vmId** | **optional.String**| External id of the virtual machine. | 

### Return type

[**VirtualNetworkInterfaceListResult**](VirtualNetworkInterfaceListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVirtualMachines**
> VirtualMachineListResult ListVirtualMachines(ctx, optional)
Return the List of Virtual Machines

Returns information about all virtual machines.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListVirtualMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListVirtualMachinesOpts struct
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

[**VirtualMachineListResult**](VirtualMachineListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVirtualSwitches**
> VirtualSwitchListResult ListVirtualSwitches(ctx, optional)
Return the List of Virtual Switches

Returns information about all virtual switches based on the request parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListVirtualSwitchesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListVirtualSwitchesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cmLocalId** | **optional.String**| Local Id of the virtual switch | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **discoveredNodeId** | **optional.String**| Discovered node ID | 
 **displayName** | **optional.String**| Display name of the virtual switch | 
 **externalId** | **optional.String**| External id of the virtual switch | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **originId** | **optional.String**| ID of the compute manager | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **uuid** | **optional.String**| UUID of the switch | 

### Return type

[**VirtualSwitchListResult**](VirtualSwitchListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVmToolsInfo**
> VmToolsInfoListResult ListVmToolsInfo(ctx, optional)
Return the list of tools and agents installed in VMs.

This API returns the list of tools and agents installed in VMs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListVmToolsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListVmToolsInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VmToolsInfoListResult**](VmToolsInfoListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListVtepLabelPools**
> VtepLabelPoolListResult ListVtepLabelPools(ctx, optional)
List virtual tunnel endpoint Label Pools

Returns a list of all virtual tunnel endpoint label pools 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiListVtepLabelPoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiListVtepLabelPoolsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**VtepLabelPoolListResult**](VtepLabelPoolListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchPaceHostConfiguration**
> IntelligenceHostConfigurationInfo PatchPaceHostConfiguration(ctx, body)
Patch NSX-Intelligence host configuration

Patch the current NSX-Intelligence host configuration. Return error if NSX-Intelligence is not registered with NSX. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IntelligenceHostConfigurationInfo**](IntelligenceHostConfigurationInfo.md)|  | 

### Return type

[**IntelligenceHostConfigurationInfo**](IntelligenceHostConfigurationInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformActionOnComputeCollection**
> PerformActionOnComputeCollection(ctx, ccExtId, optional)
Perform action specific to NSX on the compute-collection. cc-ext-id should be of type VC_Cluster.

Perform action specific to NSX on the compute-collection. cc-ext-id should be of type VC_Cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ccExtId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiPerformActionOnComputeCollectionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiPerformActionOnComputeCollectionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**| Supported actions on compute-collection | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformHostNodeUpgradeActionUpgradeInfra**
> Node PerformHostNodeUpgradeActionUpgradeInfra(ctx, nodeId, optional)
Perform a service deployment upgrade on a host node

Perform a service deployment upgrade on a host node

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiPerformHostNodeUpgradeActionUpgradeInfraOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiPerformHostNodeUpgradeActionUpgradeInfraOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **disableVmMigration** | **optional.Bool**| Should VM migration be disabled during upgrade | [default to false]

### Return type

[**Node**](Node.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformNodeAction**
> Node PerformNodeAction(ctx, nodeId, optional)
Perform an Action on Fabric Node

The supported fabric node actions are enter_maintenance_mode, exit_maintenance_mode for EdgeNode. This API is deprecated, please call TransportNode maintenance mode API to update maintenance mode, refer to \"Update transport node maintenance mode\". 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiPerformNodeActionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiPerformNodeActionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**| Supported fabric node actions | 
 **evacuatePoweredOffVms** | **optional.Bool**| Evacuate powered-off vms | [default to false]
 **vsanMode** | **optional.String**| Vsan decommission mode | [default to ensure_object_accessibility]

### Return type

[**Node**](Node.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PerformRepoSyncRepoSync**
> PerformRepoSyncRepoSync(ctx, )
Synchronizes the repository data between nsx managers.

Attempts to synchronize the repository partition on nsx manager. Repository partition contains packages required for the install and upgrade of nsx components.Normally there is no need to call this API explicitely by the user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostNodeSyslogExporter**
> NodeSyslogExporterProperties PostNodeSyslogExporter(ctx, body)
Add node syslog exporter

Adds a rule for exporting syslog information to a specified server. The required parameters are the rule name (exporter_name); severity level (emerg, alert, crit, and so on); transmission protocol (TCP or UDP); and server IP address or hostname. The optional parameters are the syslog port number, which can be 1 through 65,535 (514, by default); facility level to use when logging messages to syslog (kern, user, mail, and so on); and message IDs (msgids), which identify the types of messages to export. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeSyslogExporterProperties**](NodeSyslogExporterProperties.md)|  | 

### Return type

[**NodeSyslogExporterProperties**](NodeSyslogExporterProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTunnels**
> TunnelList QueryTunnels(ctx, nodeId, optional)
List of tunnels

List of tunnels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| ID of transport node | 
 **optional** | ***SystemAdministrationConfigurationApiQueryTunnelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiQueryTunnelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bfdDiagnosticCode** | **optional.String**| BFD diagnostic code of Tunnel as defined in RFC 5880 | 
 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **remoteNodeId** | **optional.String**|  | 
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 
 **source** | **optional.String**| Data source type. | 
 **status** | **optional.String**| Tunnel status | 

### Return type

[**TunnelList**](TunnelList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadApplianceManagementService**
> NodeServiceProperties ReadApplianceManagementService(ctx, )
Read appliance management service properties

Read appliance management service properties

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

# **ReadApplianceManagementServiceStatus**
> NodeServiceStatusProperties ReadApplianceManagementServiceStatus(ctx, )
Read appliance management service status

Read appliance management service status

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

# **ReadApplianceManagementTaskProperties**
> ApplianceManagementTaskProperties ReadApplianceManagementTaskProperties(ctx, taskId, optional)
Read task properties

Read task properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to read | 
 **optional** | ***SystemAdministrationConfigurationApiReadApplianceManagementTaskPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadApplianceManagementTaskPropertiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **suppressRedirect** | **optional.Bool**| Suppress redirect status if applicable | [default to false]

### Return type

[**ApplianceManagementTaskProperties**](ApplianceManagementTaskProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadApplianceNodeStatus**
> NodeStatusProperties ReadApplianceNodeStatus(ctx, )
Read node status

Returns information about the NSX Manager appliance's file system, CPU, memory, disk usage, and uptime. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeStatusProperties**](NodeStatusProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAsyncApplianceManagementTaskResponse**
> ReadAsyncApplianceManagementTaskResponse(ctx, taskId)
Read asynchronous task response

Read asynchronous task response

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to read | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAsyncReplicatorService**
> NodeAsyncReplicatorServiceProperties ReadAsyncReplicatorService(ctx, )
Read the Async Replicator service properties

Read the Async Replicator service properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeAsyncReplicatorServiceProperties**](NodeAsyncReplicatorServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadAsyncReplicatorServiceStatus**
> NodeServiceStatusProperties ReadAsyncReplicatorServiceStatus(ctx, )
Read the Async Replicator service status

Read the Async Replicator service status

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

# **ReadCentralConfigProperties**
> CentralConfigProperties ReadCentralConfigProperties(ctx, )
Read Central Config properties

Read Central Config properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CentralConfigProperties**](CentralConfigProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCentralNodeConfigProfile**
> CentralNodeConfigProfile ReadCentralNodeConfigProfile(ctx, profileId, optional)
Read Central Node Config profile

Returns properties in specified Central Node Config profile. Sensitive data (like SNMP v2c community strings) are included only if query parameter \"show_sensitive_data\" is true. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**| Central Node Config profile id | 
 **optional** | ***SystemAdministrationConfigurationApiReadCentralNodeConfigProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadCentralNodeConfigProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **showSensitiveData** | **optional.Bool**| Show sensitive data in Central Node Config profile | [default to false]

### Return type

[**CentralNodeConfigProfile**](CentralNodeConfigProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterBootManagerService**
> NodeServiceProperties ReadClusterBootManagerService(ctx, )
Read cluster boot manager service properties

Read cluster boot manager service properties

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

# **ReadClusterBootManagerServiceStatus**
> NodeServiceStatusProperties ReadClusterBootManagerServiceStatus(ctx, )
Read cluster boot manager service status

Read cluster boot manager service status

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

# **ReadClusterConfig**
> ClusterConfig ReadClusterConfig(ctx, )
Read Cluster Configuration

Returns information about the NSX cluster configuration. An NSX cluster has two functions or purposes, commonly referred to as \"roles.\" These two roles are control and management. Each NSX installation has a single cluster. Separate NSX clusters do not share data. In other words, a given data-plane node is attached to only one cluster, not to multiple clusters. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterConfig**](ClusterConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterNodeConfig**
> ClusterNodeConfig ReadClusterNodeConfig(ctx, nodeId)
Read Cluster Node Configuration

Returns information about the specified NSX cluster node. Deprecated. Use GET /cluster/<node-id> to get cluster node configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**ClusterNodeConfig**](ClusterNodeConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterNodeInterface**
> NodeInterfaceProperties ReadClusterNodeInterface(ctx, nodeId, interfaceId, optional)
Read the node's Network Interface

Returns detailed information about the specified interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network  mask, and the IP configuration method (static or DHCP). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
  **interfaceId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadClusterNodeInterfaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadClusterNodeInterfaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfaceProperties**](NodeInterfaceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterNodeInterfaceStatistics**
> NodeInterfaceStatisticsProperties ReadClusterNodeInterfaceStatistics(ctx, nodeId, interfaceId, optional)
Read the NSX Manager/Controller's Network Interface Statistics

On the specified interface, returns the number of received (rx), transmitted (tx), and dropped packets; the number of bytes and errors received and transmitted on the interface; and the number of detected collisions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
  **interfaceId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadClusterNodeInterfaceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadClusterNodeInterfaceStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfaceStatisticsProperties**](NodeInterfaceStatisticsProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterNodeStatus**
> ClusterNodeStatus ReadClusterNodeStatus(ctx, nodeId, optional)
Read cluster node runtime status

Read aggregated runtime status of cluster node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadClusterNodeStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadClusterNodeStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**ClusterNodeStatus**](ClusterNodeStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterNodeVMDeploymentRequest**
> ClusterNodeVmDeploymentRequest ReadClusterNodeVMDeploymentRequest(ctx, nodeId)
Returns info for a cluster-node VM auto-deployment attempt

Returns deployment request information for a specific attempted deployment of a cluster node VM. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**ClusterNodeVmDeploymentRequest**](ClusterNodeVMDeploymentRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterNodeVMDeploymentStatus**
> ClusterNodeVmDeploymentStatusReport ReadClusterNodeVMDeploymentStatus(ctx, nodeId)
Returns the status of the VM creation/deletion

Returns the current deployment or undeployment status for a VM along with any other relevant current information, such as error messages. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**ClusterNodeVmDeploymentStatusReport**](ClusterNodeVMDeploymentStatusReport.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterNodesAggregateStatus**
> ClustersAggregateInfo ReadClusterNodesAggregateStatus(ctx, )
Read cluster runtime status

Read aggregated runtime status of all cluster nodes. Deprecated. Use GET /cluster/status instead. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClustersAggregateInfo**](ClustersAggregateInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadClusterStatus**
> ClusterStatus ReadClusterStatus(ctx, )
Read Cluster Status

Returns status information for the NSX cluster control role and management role. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClusterStatus**](ClusterStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadCminventoryService**
> NodeServiceProperties ReadCminventoryService(ctx, )
Read cm inventory service properties

Read cm inventory service properties

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

# **ReadCminventoryServiceStatus**
> NodeServiceStatusProperties ReadCminventoryServiceStatus(ctx, )
Read manager service status

Read manager service status

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

# **ReadComputeCollection**
> ComputeCollection ReadComputeCollection(ctx, ccExtId)
Return Compute Collection Information

Returns information about a specific compute collection.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ccExtId** | **string**|  | 

### Return type

[**ComputeCollection**](ComputeCollection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadComputeManager**
> ComputeManager ReadComputeManager(ctx, computeManagerId)
Return compute manager Information

Returns information about a specific compute manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **computeManagerId** | **string**|  | 

### Return type

[**ComputeManager**](ComputeManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadComputeManagerStatus**
> ComputeManagerStatus ReadComputeManagerStatus(ctx, computeManagerId)
Return runtime status information for a compute manager

Returns connection and version information about a compute manager 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **computeManagerId** | **string**|  | 

### Return type

[**ComputeManagerStatus**](ComputeManagerStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadControllerServerCertificate**
> CertificateKeyPair ReadControllerServerCertificate(ctx, )
Read controller server certificate properties

Read controller server certificate properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CertificateKeyPair**](CertificateKeyPair.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadControllerServerService**
> NodeServiceProperties ReadControllerServerService(ctx, )
Read controller service properties

Read controller service properties

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

# **ReadControllerServerServiceStatus**
> NodeServiceStatusProperties ReadControllerServerServiceStatus(ctx, )
Read controller service status

Read controller service status

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

# **ReadDiscoveredNode**
> DiscoveredNode ReadDiscoveredNode(ctx, nodeExtId)
Return Discovered Node Information

Returns information about a specific discovered node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeExtId** | **string**|  | 

### Return type

[**DiscoveredNode**](DiscoveredNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadEdgeCluster**
> EdgeCluster ReadEdgeCluster(ctx, edgeClusterId)
Read Edge Cluster

Returns information about the specified edge cluster.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **edgeClusterId** | **string**|  | 

### Return type

[**EdgeCluster**](EdgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFabricNodeInterface**
> NodeInterfaceProperties ReadFabricNodeInterface(ctx, nodeId, interfaceId, optional)
Read the node's Network Interface

Returns detailed information about the specified interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network  mask, and the IP configuration method (static or DHCP). This api is deprecated as part of FN+TN unification. Please use Transport Node API GET /transport-nodes/<transport-node-id>/network/interfaces/<interface-id> to get interface details of a node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
  **interfaceId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadFabricNodeInterfaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadFabricNodeInterfaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfaceProperties**](NodeInterfaceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFabricNodeInterfaceStatistics**
> NodeInterfaceStatisticsProperties ReadFabricNodeInterfaceStatistics(ctx, nodeId, interfaceId, optional)
Read the NSX Manager's Network Interface Statistics

On the specified interface, returns the number of received (rx), transmitted (tx), and dropped packets; the number of bytes and errors received and transmitted on the interface; and the number of detected collisions. This api is deprecated as part of FN+TN unification. Please use /transport-nodes/<transportnode-id>/network/interfaces/<interface-id>/stats to read network interface statistics with contraint FN is converted to TN. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
  **interfaceId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadFabricNodeInterfaceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadFabricNodeInterfaceStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfaceStatisticsProperties**](NodeInterfaceStatisticsProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadFabricNodeNeighborProperties**
> InterfaceNeighborProperties ReadFabricNodeNeighborProperties(ctx, fabricNodeId, interfaceName)
Read LLDP Neighbor Properties of Fabric Node by Interface Name

Read LLDP Neighbor Properties for a specific interface of Fabric Node 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fabricNodeId** | **string**| ID of fabric node | 
  **interfaceName** | **string**| Interface name to read | 

### Return type

[**InterfaceNeighborProperties**](InterfaceNeighborProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIdpsReportingService**
> NodeServiceProperties ReadIdpsReportingService(ctx, )
Read the idps-reporting service properties

Read the idps-reporting service properties

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

# **ReadIdpsReportingServiceStatus**
> NodeServiceStatusProperties ReadIdpsReportingServiceStatus(ctx, )
Read the idps-reporting service status

Read the idps-reporting service status

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

# **ReadIntelligenceUpgradeCoordinatorService**
> NodeServiceProperties ReadIntelligenceUpgradeCoordinatorService(ctx, )
Read intelligence upgrade coordinator service properties

Read intelligence upgrade coordinator service properties

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

# **ReadIntelligenceUpgradeCoordinatorServiceStatus**
> NodeServiceStatusProperties ReadIntelligenceUpgradeCoordinatorServiceStatus(ctx, )
Read intelligence upgrade coordinator service status

Read intelligence upgrade coordinator service status

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

# **ReadIpBlock**
> IpBlock ReadIpBlock(ctx, blockId)
Get IP address block information.

Returns information about the IP address block with specified id. Information includes id, display_name, description & cidr. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockId** | **string**| IP address block id | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIpBlockSubnet**
> IpBlockSubnet ReadIpBlockSubnet(ctx, subnetId)
Get the subnet within an IP block

Returns information about the subnet with specified id within a given IP address block. Information includes display_name, description, cidr and allocation_ranges. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subnetId** | **string**| Subnet id | 

### Return type

[**IpBlockSubnet**](IpBlockSubnet.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadIpPool**
> IpPool ReadIpPool(ctx, poolId)
Read IP Pool

Returns information about the specified IP address pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| IP pool ID | 

### Return type

[**IpPool**](IpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadLiagentService**
> NodeServiceProperties ReadLiagentService(ctx, )
Read liagent service properties

Read liagent service properties

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

# **ReadLiagentServiceStatus**
> NodeServiceStatusProperties ReadLiagentServiceStatus(ctx, )
Read liagent service status

Read liagent service status

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

# **ReadMPAConfiguration**
> MpaConfigProperties ReadMPAConfiguration(ctx, )
Get MPA configuration for this node

Retrieve the MPA configuration for this node to identify the Manager nodes with which this node is communicating.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MpaConfigProperties**](MPAConfigProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMacPool**
> MacPool ReadMacPool(ctx, poolId)
Read MAC Pool

Returns information about the specified MAC pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| MAC pool ID | 

### Return type

[**MacPool**](MacPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadManagementConfig**
> ManagementConfig ReadManagementConfig(ctx, )
Read NSX Management nodes global configuration.

Returns the NSX Management nodes global configuration. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ManagementConfig**](ManagementConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadManagementPlaneConfiguration**
> ManagementPlaneProperties ReadManagementPlaneConfiguration(ctx, )
Get management plane configuration for this node

Retrieve the management plane configuration for this node to identify the Manager node with which the controller service is communicating.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ManagementPlaneProperties**](ManagementPlaneProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadMigrationCoordinatorService**
> NodeServiceProperties ReadMigrationCoordinatorService(ctx, )
Read migration coordinator service properties

Read migration coordinator service properties

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

# **ReadMigrationCoordinatorServiceStatus**
> NodeServiceStatusProperties ReadMigrationCoordinatorServiceStatus(ctx, )
Read migration coordinator service status

Read migration coordinator service status

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

# **ReadNSXMessageBusService**
> NodeServiceProperties ReadNSXMessageBusService(ctx, )
Read NSX Message Bus service properties

Read NSX Message Bus service properties

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

# **ReadNSXMessageBusServiceStatus**
> NodeServiceStatusProperties ReadNSXMessageBusServiceStatus(ctx, )
Read NSX Message Bus service status

Read NSX Message Bus service status

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

# **ReadNTPService**
> NodeNtpServiceProperties ReadNTPService(ctx, )
Read NTP service properties

Read NTP service properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeNtpServiceProperties**](NodeNtpServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNTPServiceStatus**
> NodeServiceStatusProperties ReadNTPServiceStatus(ctx, )
Read NTP service status

Read NTP service status

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

# **ReadNeighborProperties**
> InterfaceNeighborProperties ReadNeighborProperties(ctx, nodeId, interfaceName)
Read LLDP Neighbor Properties of Transport Node by Interface Name 

Read LLDP Neighbor Properties for a specific interface of Transport Node 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**| ID of transport node | 
  **interfaceName** | **string**| Interface name to read | 

### Return type

[**InterfaceNeighborProperties**](InterfaceNeighborProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNetworkInterfaceStatistics**
> NodeInterfaceStatisticsProperties ReadNetworkInterfaceStatistics(ctx, interfaceId)
Read the NSX Manager's Network Interface Statistics

On the specified interface, returns the number of received (rx), transmitted (tx), and dropped packets; the number of bytes and errors received and transmitted on the interface; and the number of detected collisions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| ID of interface to read | 

### Return type

[**NodeInterfaceStatisticsProperties**](NodeInterfaceStatisticsProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNetworkProperties**
> NodeNetworkProperties ReadNetworkProperties(ctx, )
Read network configuration properties

Read network configuration properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeNetworkProperties**](NodeNetworkProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNode**
> Node ReadNode(ctx, nodeId)
Return Node Information

Returns information about a specific fabric node (host or edge). This api is deprecated, use Transport Node API GET /transport-nodes/&lt;transport-node-id&gt; to get fabric node information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**Node**](Node.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeInterface**
> NodeNetworkInterfaceProperties ReadNodeInterface(ctx, interfaceId)
Read the NSX Manager's Network Interface

Returns detailed information about the specified interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network  mask, and the IP configuration method. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **interfaceId** | **string**| ID of interface to read | 

### Return type

[**NodeNetworkInterfaceProperties**](NodeNetworkInterfaceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeNameServers**
> NodeNameServersProperties ReadNodeNameServers(ctx, )
Read the NSX Manager's Name Servers

Returns the list of servers that the NSX Manager node uses to look up IP addresses associated with given domain names. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeNameServersProperties**](NodeNameServersProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeNetworkRoute**
> NodeRouteProperties ReadNodeNetworkRoute(ctx, routeId)
Read node network route

Returns detailed information about a specified route in the NSX Manager routing table. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeId** | **string**| ID of route to read | 

### Return type

[**NodeRouteProperties**](NodeRouteProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeProcess**
> NodeProcessProperties ReadNodeProcess(ctx, processId)
Read node process

Returns information for a specified process ID (pid).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **processId** | **string**| ID of process to read | 

### Return type

[**NodeProcessProperties**](NodeProcessProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeProperties**
> NodeProperties ReadNodeProperties(ctx, )
Read node properties

Returns information about the NSX appliance. Information includes release number, time zone, system time, kernel version, message of the day (motd), and host name. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeProperties**](NodeProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeSearchDomains**
> NodeSearchDomainsProperties ReadNodeSearchDomains(ctx, )
Read the NSX Manager's Search Domains

Returns the domain list that the NSX Manager node uses to complete unqualified host names. When a host name does not include a fully qualified domain name (FQDN), the NSX Management node appends the first-listed domain name to the host name before the host name is looked up. The NSX Management node continues this for each entry in the domain list until it finds a match. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeSearchDomainsProperties**](NodeSearchDomainsProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeStatsService**
> NodeServiceProperties ReadNodeStatsService(ctx, )
Read NSX node-stats service properties

Read NSX node-stats service properties

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

# **ReadNodeStatsServiceStatus**
> NodeServiceStatusProperties ReadNodeStatsServiceStatus(ctx, )
Read NSX node-stats service status

Read NSX node-stats service status

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

# **ReadNodeStatus**
> NodeStatus ReadNodeStatus(ctx, nodeId, optional)
Return Runtime Status Information for a Node

Returns connectivity, heartbeat, and version information about a fabric node (host or edge). Note that the LCP connectivity status remains down until after the fabric node has been added as a transpot node and the NSX host switch has been successfully installed. See POST /api/v1/transport-nodes. This api is deprecated, use GET /api/v1/transport-nodes/&lt;node-id&gt;/status to get status information of a node with constraint FN is converted to TN. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadNodeStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadNodeStatusOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeStatus**](NodeStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeSupportBundle**
> ReadNodeSupportBundle(ctx, optional)
Read node support bundle

Read node support bundle

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiReadNodeSupportBundleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadNodeSupportBundleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.Bool**| Include all files | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeSyslogExporter**
> NodeSyslogExporterProperties ReadNodeSyslogExporter(ctx, exporterName)
Read node syslog exporter

Returns information about a specific syslog collection point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exporterName** | **string**| Name of syslog exporter | 

### Return type

[**NodeSyslogExporterProperties**](NodeSyslogExporterProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodeVersion**
> NodeVersion ReadNodeVersion(ctx, )
Read node version

Read node version

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeVersion**](NodeVersion.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNodesStatus**
> NodeStatusListResult ReadNodesStatus(ctx, nodeIds)
Return Runtime Status Information for given Nodes

Returns connectivity, heartbeat, and version information about all fabric nodes (host or edge). This api is deprecated as part of FN+TN unification. Please use Transport Node Status API /transport-nodes/&lt;node-id&gt;/status to get status information of a node and to get all transport nodes ids use GET /transport-nodes. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeIds** | **string**| List of requested Nodes. | 

### Return type

[**NodeStatusListResult**](NodeStatusListResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadNsxUiServiceService**
> NodeServiceProperties ReadNsxUiServiceService(ctx, )
Read ui service properties

Read ui service properties

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

# **ReadNsxUiServiceServiceStatus**
> NodeServiceStatusProperties ReadNsxUiServiceServiceStatus(ctx, )
Read ui service status

Read ui service status

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

# **ReadNsxUpgradeAgentService**
> NodeServiceProperties ReadNsxUpgradeAgentService(ctx, )
Read NSX upgrade Agent service properties

Read NSX upgrade Agent service properties

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

# **ReadNsxUpgradeAgentServiceStatus**
> NodeServiceStatusProperties ReadNsxUpgradeAgentServiceStatus(ctx, )
Read Nsx upgrade agent service status

Read Nsx upgrade agent service status

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

# **ReadPaceClusterNodeVMDeploymentRequest**
> IntelligenceClusterNodeVmDeploymentRequest ReadPaceClusterNodeVMDeploymentRequest(ctx, nodeId)
Returns info for a Intelligence cluster node VM auto-deployment attempt

Returns deployment request information for a specific attempted deployment of a cluster node VM. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**IntelligenceClusterNodeVmDeploymentRequest**](IntelligenceClusterNodeVMDeploymentRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPaceClusterNodeVMDeploymentStatus**
> IntelligenceClusterNodeVmDeploymentStatusReport ReadPaceClusterNodeVMDeploymentStatus(ctx, nodeId)
Returns the status of the VM creation/deletion

Returns the current deployment or undeployment status for a VM along with any other relevant current information, such as error messages. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

[**IntelligenceClusterNodeVmDeploymentStatusReport**](IntelligenceClusterNodeVMDeploymentStatusReport.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPhonehomeCoordinatorService**
> NodeServiceProperties ReadPhonehomeCoordinatorService(ctx, )
Read Telemetry service properties

Read Telemetry service properties

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

# **ReadPhonehomeCoordinatorServiceStatus**
> NodeServiceStatusProperties ReadPhonehomeCoordinatorServiceStatus(ctx, )
Read Telemetry service status

Read Telemetry service status

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

# **ReadPlatformClientService**
> NodeServiceProperties ReadPlatformClientService(ctx, )
Read NSX Platform Client service properties

Read NSX Platform Client service properties

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

# **ReadPlatformClientServiceStatus**
> NodeServiceStatusProperties ReadPlatformClientServiceStatus(ctx, )
Read NSX Platform Client service status

Read NSX Platform Client service status

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

# **ReadPolicyService**
> NodePolicyServiceProperties ReadPolicyService(ctx, )
Read service properties

Read service properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodePolicyServiceProperties**](NodePolicyServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadPolicyServiceStatus**
> NodeServiceStatusProperties ReadPolicyServiceStatus(ctx, )
Read service status

Read service status

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

# **ReadProtonService**
> NodeProtonServiceProperties ReadProtonService(ctx, )
Read service properties

Read service properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeProtonServiceProperties**](NodeProtonServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadProtonServiceStatus**
> NodeServiceStatusProperties ReadProtonServiceStatus(ctx, )
Read service status

Read service status

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

# **ReadProxyService**
> NodeHttpServiceProperties ReadProxyService(ctx, )
Read http service properties

This API is deprecated.  Read the configuration of the http service by calling the GET /api/v1/cluster/api-service API. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeHttpServiceProperties**](NodeHttpServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadProxyServiceStatus**
> NodeServiceStatusProperties ReadProxyServiceStatus(ctx, )
Read http service status

Read http service status

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

# **ReadRabbitMQService**
> NodeServiceProperties ReadRabbitMQService(ctx, )
Read Rabbit MQ service properties

Read Rabbit MQ service properties

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

# **ReadRabbitMQServiceStatus**
> NodeServiceStatusProperties ReadRabbitMQServiceStatus(ctx, )
Read Rabbit MQ service status

Read Rabbit MQ service status

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

# **ReadRepositoryService**
> NodeInstallUpgradeServiceProperties ReadRepositoryService(ctx, )
Read NSX install-upgrade service properties

Read NSX install-upgrade service properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeInstallUpgradeServiceProperties**](NodeInstallUpgradeServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadRepositoryServiceStatus**
> NodeServiceStatusProperties ReadRepositoryServiceStatus(ctx, )
Read NSX install-upgrade service status

Read NSX install-upgrade service status

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

# **ReadSNMPService**
> NodeSnmpServiceProperties ReadSNMPService(ctx, optional)
Read SNMP service properties

Read SNMP service properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemAdministrationConfigurationApiReadSNMPServiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadSNMPServiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showSensitiveData** | **optional.Bool**| Show SNMP sensitive data or not | [default to false]

### Return type

[**NodeSnmpServiceProperties**](NodeSnmpServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSNMPServiceStatus**
> NodeServiceStatusProperties ReadSNMPServiceStatus(ctx, )
Read SNMP service status

Read SNMP service status

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

# **ReadSNMPV3EngineID**
> NodeSnmpV3EngineId ReadSNMPV3EngineID(ctx, )
Read SNMP V3 Engine ID

Read SNMP V3 Engine ID

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeSnmpV3EngineId**](NodeSnmpV3EngineID.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSSHService**
> NodeSshServiceProperties ReadSSHService(ctx, )
Read ssh service properties

Read ssh service properties

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeSshServiceProperties**](NodeSshServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadSSHServiceStatus**
> NodeServiceStatusProperties ReadSSHServiceStatus(ctx, )
Read ssh service status

Read ssh service status

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

# **ReadSearchService**
> NodeServiceProperties ReadSearchService(ctx, )
Read NSX Search service properties

Read NSX Search service properties

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

# **ReadSearchServiceStatus**
> NodeServiceStatusProperties ReadSearchServiceStatus(ctx, )
Read NSX Search service status

Read NSX Search service status

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

# **ReadSyslogService**
> NodeServiceProperties ReadSyslogService(ctx, )
Read syslog service properties

Read syslog service properties

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

# **ReadSyslogServiceStatus**
> NodeServiceStatusProperties ReadSyslogServiceStatus(ctx, )
Read syslog service status

Read syslog service status

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

# **ReadTaskProperties**
> TaskProperties ReadTaskProperties(ctx, taskId)
Get information about the specified task

Get information about the specified task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to read | 

### Return type

[**TaskProperties**](TaskProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTaskResult**
> interface{} ReadTaskResult(ctx, taskId)
Get the response of a task

Get the response of a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**| ID of task to read | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTransportNodeInterface**
> NodeInterfaceProperties ReadTransportNodeInterface(ctx, transportNodeId, interfaceId, optional)
Read the transport node's network interface

Returns detailed information about the specified interface. Interface information includes MTU, broadcast and host IP addresses, link and admin status, MAC address, network  mask, and the IP configuration method (static or DHCP). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 
  **interfaceId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadTransportNodeInterfaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadTransportNodeInterfaceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfaceProperties**](NodeInterfaceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTransportNodeInterfaceStatistics**
> NodeInterfaceStatisticsProperties ReadTransportNodeInterfaceStatistics(ctx, transportNodeId, interfaceId, optional)
Read the NSX Manager's Network Interface Statistics

On the specified interface, returns the number of received (rx), transmitted (tx), and dropped packets; the number of bytes and errors received and transmitted on the interface; and the number of detected collisions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 
  **interfaceId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiReadTransportNodeInterfaceStatisticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiReadTransportNodeInterfaceStatisticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Data source type. | 

### Return type

[**NodeInterfaceStatisticsProperties**](NodeInterfaceStatisticsProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVNIPool**
> VniPool ReadVNIPool(ctx, poolId)
Read VNI Pool

Returns information about the specified virtual network identifier (VNI) pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| VNI pool ID | 

### Return type

[**VniPool**](VniPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadVtepLabelPool**
> VtepLabelPool ReadVtepLabelPool(ctx, poolId)
Read a virtual tunnel endpoint label pool

Returns information about the specified virtual tunnel endpoint label pool. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **poolId** | **string**| Virtual tunnel endpoint label pool ID | 

### Return type

[**VtepLabelPool**](VtepLabelPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReapplyTNProfileOnDiscoveredNodeReapplyClusterConfig**
> TransportNode ReapplyTNProfileOnDiscoveredNodeReapplyClusterConfig(ctx, nodeExtId)
Apply cluster level config on Discovered Node

When transport node profile (TNP) is applied to a cluster, if any validation fails (e.g. VMs running on host) then transport node (TN) is not created. In that case after the required action is taken (e.g. VMs powered off), you can call this API to try to create TN for that discovered node. Do not call this API if Transport Node already exists for the discovered node. In that case use API on transport node. /transport-nodes/<transport-node-id>?action=restore_cluster_config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeExtId** | **string**|  | 

### Return type

[**TransportNode**](TransportNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefreshTransportNode**
> RefreshTransportNode(ctx, transportNodeId)
Refresh the node configuration for the Edge node.

The API is applicable for Edge transport nodes. If you update the VM configuration and find a discrepancy in VM configuration at NSX Manager, then use this API to refresh configuration at NSX Manager. It refreshes the VM configuration from sources external to MP. Sources include vSphere Server and the edge node. After this action, the API GET api/v1/transport-nodes will show refreshed data. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterBatchRequest**
> BatchResponse RegisterBatchRequest(ctx, body, optional)
Register a Collection of API Calls at a Single End Point

Enables you to make multiple API requests using a single request. The batch API takes in an array of logical HTTP requests represented as JSON arrays. Each request has a method (GET, PUT, POST, or DELETE), a relative_url (the portion of the URL after https://&lt;nsx-mgr&gt;/api/), optional headers array (corresponding to HTTP headers) and an optional body (for POST and PUT requests). The batch API returns an array of logical HTTP responses represented as JSON arrays. Each response has a status code, an optional headers array and an optional body (which is a JSON-encoded string). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRequest**](BatchRequest.md)|  | 
 **optional** | ***SystemAdministrationConfigurationApiRegisterBatchRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiRegisterBatchRequestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **atomic** | **optional.**| transactional atomicity for the batch of requests embedded in the batch list | [default to false]

### Return type

[**BatchResponse**](BatchResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveVirtualMachineTagsRemoveTags**
> RemoveVirtualMachineTagsRemoveTags(ctx, body)
Perform action on specified virtual machine e.g. update tags

Perform action on a specific virtual machine. External id of the virtual machine needs to be provided in the request body. Some of the actions that can be performed are update tags, add tags, remove tags. To add tags to existing list of tag, use action parameter add_tags. To remove tags from existing list of tag, use action parameter remove_tags. To replace existing tags with new tags, use action parameter update_tags. To clear all tags, provide an empty list and action parameter as update_tags. The vmw-async: True HTTP header cannot be used with this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualMachineTagUpdate**](VirtualMachineTagUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplaceEdgeClusterMemberTransportNodeReplaceTransportNode**
> EdgeCluster ReplaceEdgeClusterMemberTransportNodeReplaceTransportNode(ctx, body, edgeClusterId)
Replace the transport node in the specified member of the edge-cluster

Replace the transport node in the specified member of the edge-cluster. This is a disruptive action. This will move all the LogicalRouterPorts(uplink and routerLink) host on the old transport_node to the new transport_node. The transportNode cannot be present in another member of any edgeClusters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EdgeClusterMemberTransportNode**](EdgeClusterMemberTransportNode.md)|  | 
  **edgeClusterId** | **string**|  | 

### Return type

[**EdgeCluster**](EdgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestDirectoryDomainSync**
> RequestDirectoryDomainSync(ctx, domainId, action, optional)
Invoke full sync or delta sync for a specific domain, with additional delay in seconds if needed.  Stop sync will try to stop any pending sync if any to return to idle state.

Invoke full sync or delta sync for a specific domain, with additional delay in seconds if needed.  Stop sync will try to stop any pending sync if any to return to idle state.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
  **action** | **string**| Sync type requested | 
 **optional** | ***SystemAdministrationConfigurationApiRequestDirectoryDomainSyncOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiRequestDirectoryDomainSyncOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **delay** | **optional.Int64**| Request to execute the sync with some delay in seconds | [default to 0]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPaceHostConfigurationReset**
> IntelligenceHostConfigurationInfo ResetPaceHostConfigurationReset(ctx, )
Reset NSX-Intelligence host configuration

Reset NSX-Intelligence host configuration to the default setting. Clear NSX-Intelligence host configuration if NSX-Intelligence is not registered with NSX. Return the NSX-Intelligence host configuration after reset operation. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IntelligenceHostConfigurationInfo**](IntelligenceHostConfigurationInfo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetPolicyServiceLoggingLevelActionResetManagerLoggingLevels**
> ResetPolicyServiceLoggingLevelActionResetManagerLoggingLevels(ctx, )
Reset the logging levels to default values

Reset the logging levels to default values

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResetProtonServiceLoggingLevelActionResetManagerLoggingLevels**
> ResetProtonServiceLoggingLevelActionResetManagerLoggingLevels(ctx, )
Reset the logging levels to default values

Reset the logging levels to default values

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartInventorySyncRestartInventorySync**
> RestartInventorySyncRestartInventorySync(ctx, nodeId)
Restart the inventory sync for the node if it is paused currently.

Restart the inventory sync for the node if it is currently internally paused. After this action the next inventory sync coming from the node is processed. This api is deprecated as part of FN+TN unification. Please use Transport Node API POST /transport-nodes/&lt;transport-node-id&gt;?action=restart_inventory_sync to restart inventory sync of node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartOrShutdownNodeRestart**
> RestartOrShutdownNodeRestart(ctx, )
Restart or shutdown node

Restarts or shuts down the NSX appliance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartOrShutdownNodeShutdown**
> RestartOrShutdownNodeShutdown(ctx, )
Restart or shutdown node

Restarts or shuts down the NSX appliance.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestartTransportNodeInventorySyncRestartInventorySync**
> RestartTransportNodeInventorySyncRestartInventorySync(ctx, transportNodeId)
Restart the inventory sync for the node if it is paused currently.

Restart the inventory sync for the node if it is currently internally paused. After this action the next inventory sync coming from the node is processed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RestoreParentClusterConfigurationRestoreClusterConfig**
> RestoreParentClusterConfigurationRestoreClusterConfig(ctx, transportNodeId)
Apply cluster level Transport Node Profile on overridden host

A host can be overridden to have different configuration than Transport Node Profile(TNP) on cluster. This action will restore such overridden host back to cluster level TNP.  This API can be used in other case. When TNP is applied to a cluster, if any validation fails (e.g. VMs running on host) then existing transport node (TN) is not updated. In that case after the issue is resolved manually (e.g. VMs powered off), you can call this API to update TN as per cluster level TNP. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportNodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResyncGlobalConfigsResyncConfig**
> GlobalConfigs ResyncGlobalConfigsResyncConfig(ctx, body, configType)
Resyncs global configurations of a config-type

It is similar to update global configurations but this request would trigger update even if the configs are unmodified. However, the realization of the new configurations is config-type specific. Refer to config-type specific documentation for details about the configuration push state. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalConfigs**](GlobalConfigs.md)|  | 
  **configType** | **string**|  | 

### Return type

[**GlobalConfigs**](GlobalConfigs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResyncTransportNodeResyncHostConfig**
> ResyncTransportNodeResyncHostConfig(ctx, transportnodeId)
Resync a Transport Node

Resync the TransportNode configuration on a host. It is similar to updating the TransportNode with existing configuration, but force synce these configurations to the host (no backend optimizations). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportnodeId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ScanDirectoryDomainSize**
> DirectoryDomainSize ScanDirectoryDomainSize(ctx, body)
Scan  the size of a directory domain

This call scans the size of a directory domain. It may be very | expensive to run this call in some AD domain deployments. Please | use it with caution.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectoryDomain**](DirectoryDomain.md)|  | 

### Return type

[**DirectoryDomainSize**](DirectoryDomainSize.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchDirectoryGroups**
> DirectoryGroupListResults SearchDirectoryGroups(ctx, domainId, filterValue, optional)
Search for directory groups within a domain based on the substring of a distinguished name. (e.g. CN=User,DC=acme,DC=com) The search filter pattern can optionally support multiple (up to 100 maximum) search pattern separated by '|' (url encoded %7C). In this case, the search results will be returned as the union of all matching criteria. (e.g. CN=Ann,CN=Users,DC=acme,DC=com|CN=Bob,CN=Users,DC=acme,DC=com)

Search for directory groups within a domain based on the substring of a distinguished name. (e.g. CN=User,DC=acme,DC=com) The search filter pattern can optionally support multiple (up to 100 maximum) search pattern separated by '|' (url encoded %7C). In this case, the search results will be returned as the union of all matching criteria. (e.g. CN=Ann,CN=Users,DC=acme,DC=com|CN=Bob,CN=Users,DC=acme,DC=com)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
  **filterValue** | **string**| Name search filter value | 
 **optional** | ***SystemAdministrationConfigurationApiSearchDirectoryGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiSearchDirectoryGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Opaque cursor to be used for getting next page of records (supplied by current result page) | 
 **includedFields** | **optional.String**| Comma separated list of fields that should be included in query result | 
 **pageSize** | **optional.Int64**| Maximum number of results to return in this page (server may return fewer) | [default to 1000]
 **sortAscending** | **optional.Bool**|  | 
 **sortBy** | **optional.String**| Field by which records are sorted | 

### Return type

[**DirectoryGroupListResults**](DirectoryGroupListResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetClusterCertificateSetClusterCertificate**
> ClusterCertificateId SetClusterCertificateSetClusterCertificate(ctx, certificateId)
Set the cluster certificate

Sets the certificate used for the MP cluster. This affects all nodes in the cluster. If the certificate used is a CA signed certificate,the request fails if the whole chain(leaf, intermediate, root) is not imported. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**| Certificate ID | 

### Return type

[**ClusterCertificateId**](ClusterCertificateId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetClusterVirtualIpSetVirtualIp**
> ClusterVirtualIpProperties SetClusterVirtualIpSetVirtualIp(ctx, ipAddress)
Set cluster virtual IP address

Sets the cluster virtual IP address. Note, all nodes in the management cluster must be in the same subnet. If not, a 409 CONFLICT status is returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ipAddress** | **string**| Virtual IP address, 0.0.0.0 if not configured | 

### Return type

[**ClusterVirtualIpProperties**](ClusterVirtualIpProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetControllerProfiler**
> SetControllerProfiler(ctx, body)
Enable or disable controller profiler

Enable or disable controller profiler

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ControllerProfilerProperties**](ControllerProfilerProperties.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNodeMandatoryAccessControl**
> MandatoryAccessControlProperties SetNodeMandatoryAccessControl(ctx, body)
Enable or disable  Mandatory Access Control

Enable or disable  Mandatory Access Control

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MandatoryAccessControlProperties**](MandatoryAccessControlProperties.md)|  | 

### Return type

[**MandatoryAccessControlProperties**](MandatoryAccessControlProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetNodeTimeSetSystemTime**
> SetNodeTimeSetSystemTime(ctx, body)
Set the node system time

Set the node system time to the given time in UTC in the RFC3339 format 'yyyy-mm-ddThh:mm:ssZ'. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeTime**](NodeTime.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRabbitMQManagementPort**
> SetRabbitMQManagementPort(ctx, )
Set RabbitMQ management port

Set RabbitMQ management port

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestDirectoryLdapServer**
> TestDirectoryLdapServer(ctx, domainId, serverId, action)
Test a LDAP server connection for directory domain

The API tests a LDAP server connection for an already configured domain. If the connection is successful, the response will be HTTP status 200. Otherwise the response will be HTTP status 500 and corresponding error message will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainId** | **string**| Directory domain identifier | 
  **serverId** | **string**| LDAP server identifier | 
  **action** | **string**| LDAP server test requested | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApiServiceConfig**
> ApiServiceConfig UpdateApiServiceConfig(ctx, body)
Update API service properties

Read the configuration of the NSX API service. Changes are applied to all nodes in the cluster. The API service on each node will restart after it is updated using this API. There may be a delay of up to a minute or so between the time this API call completes and when the new configuration goes into effect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiServiceConfig**](ApiServiceConfig.md)|  | 

### Return type

[**ApiServiceConfig**](ApiServiceConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateApplianceNodeStatusClearBootupError**
> NodeStatusProperties UpdateApplianceNodeStatusClearBootupError(ctx, )
Update node status

Clear node bootup status 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NodeStatusProperties**](NodeStatusProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAsyncReplicatorService**
> NodeAsyncReplicatorServiceProperties UpdateAsyncReplicatorService(ctx, body)
Update the async_replicator service properties

Update the async_replicator service properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeAsyncReplicatorServiceProperties**](NodeAsyncReplicatorServiceProperties.md)|  | 

### Return type

[**NodeAsyncReplicatorServiceProperties**](NodeAsyncReplicatorServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCentralConfigProperties**
> CentralConfigProperties UpdateCentralConfigProperties(ctx, body)
Update Central Config properties

Update Central Config properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CentralConfigProperties**](CentralConfigProperties.md)|  | 

### Return type

[**CentralConfigProperties**](CentralConfigProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCentralNodeConfigProfile**
> CentralNodeConfigProfile UpdateCentralNodeConfigProfile(ctx, body, nodeConfigProfileId)
Configure Node Config profile

Updates properties in the specified Central Node Config profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CentralNodeConfigProfile**](CentralNodeConfigProfile.md)|  | 
  **nodeConfigProfileId** | **string**|  | 

### Return type

[**CentralNodeConfigProfile**](CentralNodeConfigProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateClusterProfile**
> ClusterProfile UpdateClusterProfile(ctx, body, clusterProfileId)
Update a cluster profile

Modifie a specified cluster profile. The body of the PUT request must include the resource_type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ClusterProfile**](ClusterProfile.md)|  | 
  **clusterProfileId** | **string**|  | 

### Return type

[**ClusterProfile**](ClusterProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateComputeCollectionFabricTemplate**
> ComputeCollectionFabricTemplate UpdateComputeCollectionFabricTemplate(ctx, body, fabricTemplateId)
Updates compute collection fabric template

Updates compute collection fabric template for the given id. This functionality is deprecated. Use Transport Node Profiles instead of this template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeCollectionFabricTemplate**](ComputeCollectionFabricTemplate.md)|  | 
  **fabricTemplateId** | **string**|  | 

### Return type

[**ComputeCollectionFabricTemplate**](ComputeCollectionFabricTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateComputeManager**
> ComputeManager UpdateComputeManager(ctx, body, computeManagerId)
Update compute manager

Updates a specified compute manager 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComputeManager**](ComputeManager.md)|  | 
  **computeManagerId** | **string**|  | 

### Return type

[**ComputeManager**](ComputeManager.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDirectoryDomain**
> DirectoryDomain UpdateDirectoryDomain(ctx, body, domainId)
Update a directory domain

Update to any field in the directory domain will trigger a full sync

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectoryDomain**](DirectoryDomain.md)|  | 
  **domainId** | **string**| Directory domain identifier | 

### Return type

[**DirectoryDomain**](DirectoryDomain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDirectoryLdapServer**
> DirectoryLdapServer UpdateDirectoryLdapServer(ctx, body, domainId, serverId)
Update a LDAP server for directory domain

Update a LDAP server for directory domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectoryLdapServer**](DirectoryLdapServer.md)|  | 
  **domainId** | **string**| Directory domain identifier | 
  **serverId** | **string**| LDAP server identifier | 

### Return type

[**DirectoryLdapServer**](DirectoryLdapServer.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEdgeCluster**
> EdgeCluster UpdateEdgeCluster(ctx, body, edgeClusterId)
Update Edge Cluster

Modifies the specified edge cluster. Modifiable parameters include the description, display_name, transport-node-id. If the optional fabric_profile_binding is included, resource_type and profile_id are required. User should do a GET on the edge-cluster and obtain the payload and retain the member_index of the existing members as returning in the GET output. For new member additions, the member_index cannot be defined by the user, user can read the system allocated index to the new member in the output of this API call or by doing a GET call. User cannot use this PUT api to replace the transport_node of an existing member because this is a disruption action, we have exposed a explicit API for doing so, refer to \"ReplaceEdgeClusterMemberTransportNode\" EdgeCluster only supports homogeneous members. The TransportNodes backed by EdgeNode are only allowed in cluster members. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EdgeCluster**](EdgeCluster.md)|  | 
  **edgeClusterId** | **string**|  | 

### Return type

[**EdgeCluster**](EdgeCluster.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFailureDomain**
> FailureDomain UpdateFailureDomain(ctx, body, failureDomainId)
Update Failure Domain

Updates an existing failure domain. Modifiable parameters are display_name, preferred_active_edge_services flag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FailureDomain**](FailureDomain.md)|  | 
  **failureDomainId** | **string**|  | 

### Return type

[**FailureDomain**](FailureDomain.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateGlobalConfigs**
> GlobalConfigs UpdateGlobalConfigs(ctx, body, configType)
Update global configurations of a config type

Updates global configurations that belong to a config type. The request must include the updated values along with the unmodified values. The values that are updated(different) would trigger update to config-type specific state. However, the realization of the new configurations is config-type specific. Refer to config-type specific documentation for details about the config- uration push state. Policy api will overwrite the fipsGlobalConfig set using MP api. Always use https://<policyIp>/policy/api/v1/infra/global-config to update fips config- uration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GlobalConfigs**](GlobalConfigs.md)|  | 
  **configType** | **string**|  | 

### Return type

[**GlobalConfigs**](GlobalConfigs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateHostSwitchProfile**
> BaseHostSwitchProfile UpdateHostSwitchProfile(ctx, body, hostSwitchProfileId)
Update a Hostswitch Profile

Modifies a specified hostswitch profile. The body of the PUT request must include the resource_type. For uplink profiles, the put request must also include teaming parameters. Modifiable attributes include display_name, mtu, and transport_vlan. For uplink teaming policies, uplink_name and policy are also modifiable. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BaseHostSwitchProfile**](BaseHostSwitchProfile.md)|  | 
  **hostSwitchProfileId** | **string**|  | 

### Return type

[**BaseHostSwitchProfile**](BaseHostSwitchProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpBlock**
> IpBlock UpdateIpBlock(ctx, body, blockId)
Update an IP Address Block

Modifies the IP address block with specifed id. display_name, description and cidr are parameters that can be modified. If a new cidr is specified, it should contain all existing subnets in the IP block. Returns a conflict error if the IP address block cidr can not be modified due to the presence of subnets that it contains. Eg: If the IP block contains a subnet 192.168.0.1/24 and we try to change the IP block cidr to 10.1.0.1/16, it results in a conflict. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpBlock**](IpBlock.md)|  | 
  **blockId** | **string**| IP address block id | 

### Return type

[**IpBlock**](IpBlock.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateIpPool**
> IpPool UpdateIpPool(ctx, body, poolId)
Update an IP Pool

Modifies the specified IP address pool. Modifiable parameters include the description, display_name, and all subnet information. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IpPool**](IpPool.md)|  | 
  **poolId** | **string**| IP pool ID | 

### Return type

[**IpPool**](IpPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMPAConfiguration**
> MpaConfigProperties UpdateMPAConfiguration(ctx, body)
Update MPA configuration for this node

Update the MPA configuration for this node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MpaConfigProperties**](MpaConfigProperties.md)|  | 

### Return type

[**MpaConfigProperties**](MPAConfigProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateManagementConfig**
> ManagementConfig UpdateManagementConfig(ctx, body)
Update NSX Management nodes global configuration

Modifies the NSX Management nodes global configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ManagementConfig**](ManagementConfig.md)|  | 

### Return type

[**ManagementConfig**](ManagementConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateManagementPlaneConfiguration**
> ManagementPlaneProperties UpdateManagementPlaneConfiguration(ctx, body)
Update management plane configuration for this node

Update the management plane configuration for this node.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ManagementPlaneProperties**](ManagementPlaneProperties.md)|  | 

### Return type

[**ManagementPlaneProperties**](ManagementPlaneProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNTPService**
> NodeNtpServiceProperties UpdateNTPService(ctx, body)
Update NTP service properties

Update NTP service properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeNtpServiceProperties**](NodeNtpServiceProperties.md)|  | 

### Return type

[**NodeNtpServiceProperties**](NodeNtpServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNode**
> Node UpdateNode(ctx, body, nodeId)
Update a Node

Modifies attributes of a fabric node (host or edge). This api is deprecated as part of FN+TN unification. Please use Transport Node API PUT /transport-nodes/&lt;transport-node-id&gt; to update fabric node details. API PUT /transport-nodes/<transport-node-id> to update fabric node details. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Node**](Node.md)|  | 
  **nodeId** | **string**|  | 

### Return type

[**Node**](Node.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeInterface**
> NodeNetworkInterfaceProperties UpdateNodeInterface(ctx, body, interfaceId)
Update the NSX Manager's Network Interface

Updates the specified interface properties. You cannot change the properties <code>ip_configuration</code>, <code>ip_addresses</code>, or <code>plane</code>. NSX Manager must have a static IP address. You must use NSX CLI to configure a controller or an edge node. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeNetworkInterfaceProperties**](NodeNetworkInterfaceProperties.md)|  | 
  **interfaceId** | **string**| ID of interface to update | 

### Return type

[**NodeNetworkInterfaceProperties**](NodeNetworkInterfaceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeNameServers**
> NodeNameServersProperties UpdateNodeNameServers(ctx, body)
Update the NSX Manager's Name Servers

Modifies the list of servers that the NSX Manager node uses to look up IP addresses associated with given domain names. If DHCP is configured, this method returns a 409 CONFLICT error, because DHCP manages the list of name servers. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeNameServersProperties**](NodeNameServersProperties.md)|  | 

### Return type

[**NodeNameServersProperties**](NodeNameServersProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeProperties**
> NodeProperties UpdateNodeProperties(ctx, body)
Update node properties

Modifies NSX appliance properties. Modifiable properties include the timezone, message of the day (motd), and hostname. The NSX appliance node_version, system_time, and kernel_version are read only and cannot be modified with this method. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeProperties**](NodeProperties.md)|  | 

### Return type

[**NodeProperties**](NodeProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNodeSearchDomains**
> NodeSearchDomainsProperties UpdateNodeSearchDomains(ctx, body)
Update the NSX Manager's Search Domains

Modifies the list of domain names that the NSX Manager node uses to complete unqualified host names. If DHCP is configured, this method returns a 409 CONFLICT error, because DHCP manages the list of name servers. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeSearchDomainsProperties**](NodeSearchDomainsProperties.md)|  | 

### Return type

[**NodeSearchDomainsProperties**](NodeSearchDomainsProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePolicyService**
> NodePolicyServiceProperties UpdatePolicyService(ctx, body)
Update service properties

Update service properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodePolicyServiceProperties**](NodePolicyServiceProperties.md)|  | 

### Return type

[**NodePolicyServiceProperties**](NodePolicyServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProtonService**
> NodeProtonServiceProperties UpdateProtonService(ctx, body)
Update service properties

Update service properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeProtonServiceProperties**](NodeProtonServiceProperties.md)|  | 

### Return type

[**NodeProtonServiceProperties**](NodeProtonServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProxyService**
> NodeHttpServiceProperties UpdateProxyService(ctx, body)
Update http service properties

This API is deprecated.  Make changes to the http service configuration by calling the PUT /api/v1/cluster/api-service API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeHttpServiceProperties**](NodeHttpServiceProperties.md)|  | 

### Return type

[**NodeHttpServiceProperties**](NodeHttpServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRealizationStateBarrierConfig**
> RealizationStateBarrierConfig UpdateRealizationStateBarrierConfig(ctx, body)
Updates the barrier configuration

Updates the barrier configuration having interval set in milliseconds The new interval that automatically increments the global realization number 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RealizationStateBarrierConfig**](RealizationStateBarrierConfig.md)|  | 

### Return type

[**RealizationStateBarrierConfig**](RealizationStateBarrierConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepositoryService**
> NodeInstallUpgradeServiceProperties UpdateRepositoryService(ctx, body)
Update NSX install-upgrade service properties

Update NSX install-upgrade service properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeInstallUpgradeServiceProperties**](NodeInstallUpgradeServiceProperties.md)|  | 

### Return type

[**NodeInstallUpgradeServiceProperties**](NodeInstallUpgradeServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSNMPService**
> NodeSnmpServiceProperties UpdateSNMPService(ctx, body)
Update SNMP service properties

Update SNMP service properties.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeSnmpServiceProperties**](NodeSnmpServiceProperties.md)|  | 

### Return type

[**NodeSnmpServiceProperties**](NodeSnmpServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSNMPV3EngineID**
> NodeSnmpV3EngineId UpdateSNMPV3EngineID(ctx, body)
Update SNMP V3 Engine ID

Update SNMP V3 Engine ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeSnmpV3EngineId**](NodeSnmpV3EngineId.md)|  | 

### Return type

[**NodeSnmpV3EngineId**](NodeSnmpV3EngineID.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSSHService**
> NodeSshServiceProperties UpdateSSHService(ctx, body)
Update ssh service properties

Update ssh service properties. If the start_on_boot property is updated to true, existing ssh sessions if any are stopped and the ssh service is restarted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NodeSshServiceProperties**](NodeSshServiceProperties.md)|  | 

### Return type

[**NodeSshServiceProperties**](NodeSshServiceProperties.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransportNodeCollection**
> TransportNodeCollection UpdateTransportNodeCollection(ctx, body, transportNodeCollectionId)
Update Transport Node collection

Attach different transport node profile to compute collection by updating transport node collection. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportNodeCollection**](TransportNodeCollection.md)|  | 
  **transportNodeCollectionId** | **string**|  | 

### Return type

[**TransportNodeCollection**](TransportNodeCollection.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransportNodeMaintenanceMode**
> UpdateTransportNodeMaintenanceMode(ctx, transportnodeId, optional)
Update transport node maintenance mode

Put transport node into maintenance mode or exit from maintenance mode.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transportnodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiUpdateTransportNodeMaintenanceModeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiUpdateTransportNodeMaintenanceModeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransportNodeProfile**
> TransportNodeProfile UpdateTransportNodeProfile(ctx, body, transportNodeProfileId, optional)
Update a Transport Node Profile

When configurations of a transport node profile(TNP) is updated, all the transport nodes in all the compute collections to which this TNP is attached are updated to reflect the updated configuration. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportNodeProfile**](TransportNodeProfile.md)|  | 
  **transportNodeProfileId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiUpdateTransportNodeProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiUpdateTransportNodeProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **esxMgmtIfMigrationDest** | **optional.**| The network ids to which the ESX vmk interfaces will be migrated | 
 **ifId** | **optional.**| The ESX vmk interfaces to migrate | 
 **pingIp** | **optional.**| IP Addresses to ping right after ESX vmk interfaces were migrated. | 
 **skipValidation** | **optional.**| Whether to skip front-end validation for vmk/vnic/pnic migration | [default to false]
 **vnic** | **optional.**| The ESX vmk interfaces and/or VM NIC to migrate | 
 **vnicMigrationDest** | **optional.**| The migration destinations of ESX vmk interfaces and/or VM NIC | 

### Return type

[**TransportNodeProfile**](TransportNodeProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransportNodeWithDeploymentInfo**
> TransportNode UpdateTransportNodeWithDeploymentInfo(ctx, body, transportNodeId, optional)
Update a Transport Node

Modifies the transport node information. The host_switch_name field must match the host_switch_name value specified in the transport zone (API: transport-zones). You must create the associated uplink profile (API: host-switch-profiles) before you can specify an uplink_name here. If the host is an ESX and has only one physical NIC being used by a vSphere standard switch, TransportNodeUpdateParameters should be used to migrate the management interface and the physical NIC into a logical switch that is in a transport zone this transport node will join or has already joined. If the migration is already done, TransportNodeUpdateParameters can also be used to migrate the management interface and the physical NIC back to a vSphere standard switch. In other cases, the TransportNodeUpdateParameters should NOT be used. When updating transport node you should follow pattern where you should fetch the existing transport node and then only modify the required properties keeping other properties as is.  It also modifies attributes of node (host or edge).  Note: Previous versions of NSX-T also used a property named transport_zone_endpoints at TransportNode level. This property is deprecated which creates some combinations of new client along with old client payloads. Examples [1] shows old/existing client request and response by populating transport_zone_endpoints property at TransportNode level. Example [2] shows TransportNode updating TransportNode from exmaple [1] request/response by adding a new StandardHostSwitch by populating transport_zone_endpoints at StandardHostSwitch level. TransportNode level transport_zone_endpoints will ONLY have TransportZoneEndpoints that were originally specified here during create/update operation and does not include TransportZoneEndpoints that were directly specified at StandardHostSwitch level. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportNode**](TransportNode.md)|  | 
  **transportNodeId** | **string**|  | 
 **optional** | ***SystemAdministrationConfigurationApiUpdateTransportNodeWithDeploymentInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemAdministrationConfigurationApiUpdateTransportNodeWithDeploymentInfoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **esxMgmtIfMigrationDest** | **optional.**| The network ids to which the ESX vmk interfaces will be migrated | 
 **ifId** | **optional.**| The ESX vmk interfaces to migrate | 
 **pingIp** | **optional.**| IP Addresses to ping right after ESX vmk interfaces were migrated. | 
 **skipValidation** | **optional.**| Whether to skip front-end validation for vmk/vnic/pnic migration | [default to false]
 **vnic** | **optional.**| The ESX vmk interfaces and/or VM NIC to migrate | 
 **vnicMigrationDest** | **optional.**| The migration destinations of ESX vmk interfaces and/or VM NIC | 

### Return type

[**TransportNode**](TransportNode.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransportZone**
> TransportZone UpdateTransportZone(ctx, body, zoneId)
Update a Transport Zone

Updates an existing transport zone. Modifiable parameters are is_default, description, and display_name. The request must include the existing host_switch_name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportZone**](TransportZone.md)|  | 
  **zoneId** | **string**|  | 

### Return type

[**TransportZone**](TransportZone.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTransportZoneProfile**
> TransportZoneProfile UpdateTransportZoneProfile(ctx, body, transportzoneProfileId)
Update a transport zone profile

Modifies a specified transport zone profile. The body of the PUT request must include the resource_type. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TransportZoneProfile**](TransportZoneProfile.md)|  | 
  **transportzoneProfileId** | **string**|  | 

### Return type

[**TransportZoneProfile**](TransportZoneProfile.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUcState**
> UpdateUcState(ctx, body)
Update UC state properties

Update UC state properties

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UcStateProperties**](UcStateProperties.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVNIPool**
> VniPool UpdateVNIPool(ctx, body, poolId)
Update a VNI Pool

Updates the specified VNI pool. Modifiable parameters include description, display_name and ranges. Ranges can be added, modified or deleted. Overlapping ranges are not allowed. Only range end can be modified for any existing range. Range shrinking or deletion is not allowed if there are any allocated VNIs. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VniPool**](VniPool.md)|  | 
  **poolId** | **string**| VNI pool ID | 

### Return type

[**VniPool**](VniPool.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVirtualMachineTagsUpdateTags**
> UpdateVirtualMachineTagsUpdateTags(ctx, body)
Perform action on specified virtual machine e.g. update tags

Perform action on a specific virtual machine. External id of the virtual machine needs to be provided in the request body. Some of the actions that can be performed are update tags, add tags, remove tags. To add tags to existing list of tag, use action parameter add_tags. To remove tags from existing list of tag, use action parameter remove_tags. To replace existing tags with new tags, use action parameter update_tags. To clear all tags, provide an empty list and action parameter as update_tags. The vmw-async: True HTTP header cannot be used with this API. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VirtualMachineTagUpdate**](VirtualMachineTagUpdate.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBundleViaLocalFileUpload**
> BundleId UploadBundleViaLocalFileUpload(ctx, file, fileType, product)
Upload bundle

Upload the bundle. This call returns after upload is completed. You can check bundle processing status periodically by retrieving bundle upload-status to find out if the upload and processing is completed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | **string****string**|  | 
  **fileType** | **string**| Type of file | 
  **product** | **string**| Name of the product | 

### Return type

[**BundleId**](BundleId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadBundleViaRemoteFile**
> BundleId UploadBundleViaRemoteFile(ctx, body, fileType, product)
Upload bundle using remote file

Upload the bundle from remote bundle URL. The call returns after fetch is initiated. Check status by periodically retrieving bundle upload status using GET /repository/bundles/<bundle-id>/upload-status. The upload is complete when the status is SUCCESS. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoteBundleUrl**](RemoteBundleUrl.md)|  | 
  **fileType** | **string**| Type of file | 
  **product** | **string**| Name of the product | 

### Return type

[**BundleId**](BundleId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyDirectoryLdapServer**
> DirectoryLdapServerStatus VerifyDirectoryLdapServer(ctx, body, action)
Test a directory domain LDAP server connectivity

This API tests a LDAP server connectivity before the actual domain or LDAP server is configured. If the connectivity is good, the response will be HTTP status 200. Otherwise the response will be HTTP status 500 and corresponding error message will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DirectoryLdapServer**](DirectoryLdapServer.md)|  | 
  **action** | **string**| LDAP server test requested | 

### Return type

[**DirectoryLdapServerStatus**](DirectoryLdapServerStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyNodeSyslogExporterVerify**
> VerifyNodeSyslogExporterVerify(ctx, )
Verify node syslog exporter

Collect iptables rules needed for all existing syslog exporters and verify if the existing iptables rules are the same. If not, remove the stale rules and add the new rules to make sure all exporters work properly. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

