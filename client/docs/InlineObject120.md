# InlineObject120

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product type to rollback (if the network is a combined network) | [optional] 
**Time** | Pointer to **time.Time** | Scheduled time for the rollback | [optional] 
**Reasons** | [**[]NetworksNetworkIdFirmwareUpgradesRollbacksReasons**](NetworksNetworkIdFirmwareUpgradesRollbacksReasons.md) | Reasons for the rollback | 
**ToVersion** | Pointer to [**NetworksNetworkIdFirmwareUpgradesRollbacksToVersion**](NetworksNetworkIdFirmwareUpgradesRollbacksToVersion.md) |  | [optional] 
**Predownload** | Pointer to [**NetworksNetworkIdFirmwareUpgradesRollbacksPredownload**](NetworksNetworkIdFirmwareUpgradesRollbacksPredownload.md) |  | [optional] 

## Methods

### NewInlineObject120

`func NewInlineObject120(reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons, ) *InlineObject120`

NewInlineObject120 instantiates a new InlineObject120 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject120WithDefaults

`func NewInlineObject120WithDefaults() *InlineObject120`

NewInlineObject120WithDefaults instantiates a new InlineObject120 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *InlineObject120) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *InlineObject120) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *InlineObject120) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *InlineObject120) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTime

`func (o *InlineObject120) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineObject120) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineObject120) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineObject120) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetReasons

`func (o *InlineObject120) GetReasons() []NetworksNetworkIdFirmwareUpgradesRollbacksReasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineObject120) GetReasonsOk() (*[]NetworksNetworkIdFirmwareUpgradesRollbacksReasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineObject120) SetReasons(v []NetworksNetworkIdFirmwareUpgradesRollbacksReasons)`

SetReasons sets Reasons field to given value.


### GetToVersion

`func (o *InlineObject120) GetToVersion() NetworksNetworkIdFirmwareUpgradesRollbacksToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineObject120) GetToVersionOk() (*NetworksNetworkIdFirmwareUpgradesRollbacksToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineObject120) SetToVersion(v NetworksNetworkIdFirmwareUpgradesRollbacksToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineObject120) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetPredownload

`func (o *InlineObject120) GetPredownload() NetworksNetworkIdFirmwareUpgradesRollbacksPredownload`

GetPredownload returns the Predownload field if non-nil, zero value otherwise.

### GetPredownloadOk

`func (o *InlineObject120) GetPredownloadOk() (*NetworksNetworkIdFirmwareUpgradesRollbacksPredownload, bool)`

GetPredownloadOk returns a tuple with the Predownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredownload

`func (o *InlineObject120) SetPredownload(v NetworksNetworkIdFirmwareUpgradesRollbacksPredownload)`

SetPredownload sets Predownload field to given value.

### HasPredownload

`func (o *InlineObject120) HasPredownload() bool`

HasPredownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


