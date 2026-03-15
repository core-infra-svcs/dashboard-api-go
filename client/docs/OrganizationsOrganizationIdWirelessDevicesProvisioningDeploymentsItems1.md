# OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentId** | Pointer to **string** | zero touch deployment request identifier | [optional] 
**Devices** | [**OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1**](OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1.md) |  | 
**Status** | **string** | Status of the zero touch deployment request. enum &#x3D; [ready, in progress, completed, failed] | 
**Type** | **string** | Type of the zero touch deployment request. enum &#x3D; [deploy, replace] | 
**Network** | Pointer to [**OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsNetwork**](OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsNetwork.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp of when the zero touch deployment request was created | [optional] 
**RequestedAt** | Pointer to **time.Time** | Timestamp of when the zero touch deployment request was created | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** | Timestamp of when the zero touch deployment request was last updated | [optional] 
**CompletedAt** | Pointer to **time.Time** | Timestamp of when the zero touch deployment request was completed | [optional] 
**Errors** | Pointer to **[]string** | Array of error message(s) if any | [optional] 

## Methods

### NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1

`func NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1(devices OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1, status string, type_ string, ) *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1`

NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1 instantiates a new OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1WithDefaults

`func NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1WithDefaults() *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1`

NewOrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1WithDefaults instantiates a new OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentId

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetDevices

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetDevices() OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetDevicesOk() (*OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetDevices(v OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsDevices1)`

SetDevices sets Devices field to given value.


### GetStatus

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetType(v string)`

SetType sets Type field to given value.


### GetNetwork

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetNetwork() OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetNetworkOk() (*OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetNetwork(v OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRequestedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetErrors

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *OrganizationsOrganizationIdWirelessDevicesProvisioningDeploymentsItems1) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


