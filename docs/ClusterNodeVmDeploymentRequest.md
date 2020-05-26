# ClusterNodeVmDeploymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentConfig** | [***ClusterNodeVmDeploymentConfig**](ClusterNodeVMDeploymentConfig.md) |  | [default to null]
**VmId** | **string** | ID of the VM maintained internally and used to recognize it. Note: This is automatically generated and cannot be modified.  | [optional] [default to null]
**UserSettings** | [***NodeUserSettings**](NodeUserSettings.md) |  | [optional] [default to null]
**Roles** | **[]string** | List of cluster node role (or roles) which the VM should take on. They specify what type (or types) of cluster node which the new VM should act as. Currently both CONTROLLER and MANAGER must be provided, since this permutation is the only one supported now.  | [default to null]
**FormFactor** | **string** | Specifies the desired \&quot;size\&quot; of the VM  | [optional] [default to FORM_FACTOR.MEDIUM]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

