# JoinClusterParameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Username on the cluster node | [optional] [default to null]
**CertficateSha256Thumbprint** | **string** | SHA256 Thumbprint of the API certificate of the cluster node | [default to null]
**Token** | **string** | Limited time OAuth token instead of the username/password | [optional] [default to null]
**ClusterId** | **string** | UUID of the cluster to join | [default to null]
**Password** | **string** | Password of the user on the cluster node | [optional] [default to null]
**IpAddress** | **string** | IP address of a node already part of the cluster to join | [default to null]
**Port** | **int64** | API port on the cluster node | [optional] [default to 443]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

