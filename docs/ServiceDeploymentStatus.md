# ServiceDeploymentStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentIssues** | [**[]ServiceDeploymentIssue**](ServiceDeploymentIssue.md) | List of issue and detailed description of the issue in case of deployment failure. | [optional] [default to null]
**LastUpdateTimestamp** | **int64** | Timestamp when the data was last updated; unset if data source has never updated the data. | [optional] [default to null]
**DeploymentStatus** | **string** | Deployment status of NXGI Partner Service-VM on a compute collection. It shows the latest status during the process of deployment, redeploy, upgrade, and un-deployment on a compute collection such as VC cluster. | [optional] [default to null]
**SvaCurrentVersion** | **string** | Currently deployed Service Virtual Appliance version. | [optional] [default to null]
**ServiceDeploymentId** | **string** | Id of service deployment. | [optional] [default to null]
**SvaMaxAvailableVersion** | **string** | Max available SVA version for upgrade | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

