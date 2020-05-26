# EdgeCluster

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemOwned** | **bool** | Indicates system owned resource | [optional] [default to null]
**DisplayName** | **string** | Defaults to ID if not set | [optional] [default to null]
**Description** | **string** | Description of this resource | [optional] [default to null]
**Tags** | [**[]Tag**](Tag.md) | Opaque identifiers meaningful to the API user | [optional] [default to null]
**CreateUser** | **string** | ID of the user who created this resource | [optional] [default to null]
**Protection** | **string** | Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite&#x3D;true. UNKNOWN - the _protection field could not be determined for this           entity.  | [optional] [default to null]
**CreateTime** | **int64** | Timestamp of resource creation | [optional] [default to null]
**LastModifiedTime** | **int64** | Timestamp of last modification | [optional] [default to null]
**LastModifiedUser** | **string** | ID of the user who last modified this resource | [optional] [default to null]
**Id** | **string** | Unique identifier of this resource | [optional] [default to null]
**ResourceType** | **string** | The type of this resource. | [optional] [default to null]
**MemberNodeType** | **string** | Edge cluster is homogenous collection of transport nodes. Hence all transport nodes of the cluster must be of same type. This readonly field shows the type of transport nodes.  | [optional] [default to null]
**Members** | [**[]EdgeClusterMember**](EdgeClusterMember.md) | EdgeCluster only supports homogeneous members. These member should be backed by either EdgeNode or PublicCloudGatewayNode. TransportNode type of these nodes should be the same. DeploymentType (VIRTUAL_MACHINE|PHYSICAL_MACHINE) of these EdgeNodes is recommended to be the same. EdgeCluster supports members of different deployment types.  | [optional] [default to null]
**NodeRtepIps** | [**[]NodeRtepIpsConfig**](NodeRtepIpsConfig.md) | List of remote tunnel endpoint ipaddress configured on edge cluster for each transport node. | [optional] [default to null]
**ClusterProfileBindings** | [**[]ClusterProfileTypeIdEntry**](ClusterProfileTypeIdEntry.md) | Edge cluster profile bindings | [optional] [default to null]
**EnableInterSiteForwarding** | **bool** | Flag should be only use in federation for inter site l2 and l3 forwarding. Before enabling this flag, all the edge cluster members must have remote tunnel endpoint configured. TIER0/TIER1 logical routers managed by GM must be associated with edge cluster which has inter-site forwarding enabled.  | [optional] [default to null]
**AllocationRules** | [**[]AllocationRule**](AllocationRule.md) | Set of allocation rules and respected action for auto placement of logical router, DHCP and MDProxy on edge cluster members.  | [optional] [default to null]
**DeploymentType** | **string** | This field is a readonly field which shows the deployment_type of members. It returns UNKNOWN if there are no members, and returns VIRTUAL_MACHINE| PHYSICAL_MACHINE if all edge members are VIRTUAL_MACHINE|PHYSICAL_MACHINE. It returns HYBRID if the cluster contains edge members of both types VIRTUAL_MACHINE and PHYSICAL_MACHINE.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

