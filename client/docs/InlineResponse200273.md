# InlineResponse200273

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpgradeId** | Pointer to **string** | The upgrade | [optional] 
**UpgradeBatchId** | Pointer to **string** | The upgrade batch | [optional] 
**Network** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesNetwork**](OrganizationsOrganizationIdFirmwareUpgradesNetwork.md) |  | [optional] 
**Status** | Pointer to **string** | Status of upgrade event: [Cancelled, Completed] | [optional] 
**Time** | Pointer to **time.Time** | Scheduled start time | [optional] 
**CompletedAt** | Pointer to **string** | Timestamp when upgrade completed. Null if status pending. | [optional] 
**ProductTypes** | Pointer to **string** | product upgraded [wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor] | [optional] 
**ToVersion** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesToVersion**](OrganizationsOrganizationIdFirmwareUpgradesToVersion.md) |  | [optional] 
**FromVersion** | Pointer to [**OrganizationsOrganizationIdFirmwareUpgradesFromVersion**](OrganizationsOrganizationIdFirmwareUpgradesFromVersion.md) |  | [optional] 

## Methods

### NewInlineResponse200273

`func NewInlineResponse200273() *InlineResponse200273`

NewInlineResponse200273 instantiates a new InlineResponse200273 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200273WithDefaults

`func NewInlineResponse200273WithDefaults() *InlineResponse200273`

NewInlineResponse200273WithDefaults instantiates a new InlineResponse200273 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpgradeId

`func (o *InlineResponse200273) GetUpgradeId() string`

GetUpgradeId returns the UpgradeId field if non-nil, zero value otherwise.

### GetUpgradeIdOk

`func (o *InlineResponse200273) GetUpgradeIdOk() (*string, bool)`

GetUpgradeIdOk returns a tuple with the UpgradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeId

`func (o *InlineResponse200273) SetUpgradeId(v string)`

SetUpgradeId sets UpgradeId field to given value.

### HasUpgradeId

`func (o *InlineResponse200273) HasUpgradeId() bool`

HasUpgradeId returns a boolean if a field has been set.

### GetUpgradeBatchId

`func (o *InlineResponse200273) GetUpgradeBatchId() string`

GetUpgradeBatchId returns the UpgradeBatchId field if non-nil, zero value otherwise.

### GetUpgradeBatchIdOk

`func (o *InlineResponse200273) GetUpgradeBatchIdOk() (*string, bool)`

GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBatchId

`func (o *InlineResponse200273) SetUpgradeBatchId(v string)`

SetUpgradeBatchId sets UpgradeBatchId field to given value.

### HasUpgradeBatchId

`func (o *InlineResponse200273) HasUpgradeBatchId() bool`

HasUpgradeBatchId returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200273) GetNetwork() OrganizationsOrganizationIdFirmwareUpgradesNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200273) GetNetworkOk() (*OrganizationsOrganizationIdFirmwareUpgradesNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200273) SetNetwork(v OrganizationsOrganizationIdFirmwareUpgradesNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200273) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200273) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200273) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200273) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200273) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *InlineResponse200273) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200273) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200273) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200273) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetCompletedAt

`func (o *InlineResponse200273) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *InlineResponse200273) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *InlineResponse200273) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *InlineResponse200273) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetProductTypes

`func (o *InlineResponse200273) GetProductTypes() string`

GetProductTypes returns the ProductTypes field if non-nil, zero value otherwise.

### GetProductTypesOk

`func (o *InlineResponse200273) GetProductTypesOk() (*string, bool)`

GetProductTypesOk returns a tuple with the ProductTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTypes

`func (o *InlineResponse200273) SetProductTypes(v string)`

SetProductTypes sets ProductTypes field to given value.

### HasProductTypes

`func (o *InlineResponse200273) HasProductTypes() bool`

HasProductTypes returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200273) GetToVersion() OrganizationsOrganizationIdFirmwareUpgradesToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200273) GetToVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200273) SetToVersion(v OrganizationsOrganizationIdFirmwareUpgradesToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200273) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetFromVersion

`func (o *InlineResponse200273) GetFromVersion() OrganizationsOrganizationIdFirmwareUpgradesFromVersion`

GetFromVersion returns the FromVersion field if non-nil, zero value otherwise.

### GetFromVersionOk

`func (o *InlineResponse200273) GetFromVersionOk() (*OrganizationsOrganizationIdFirmwareUpgradesFromVersion, bool)`

GetFromVersionOk returns a tuple with the FromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromVersion

`func (o *InlineResponse200273) SetFromVersion(v OrganizationsOrganizationIdFirmwareUpgradesFromVersion)`

SetFromVersion sets FromVersion field to given value.

### HasFromVersion

`func (o *InlineResponse200273) HasFromVersion() bool`

HasFromVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


