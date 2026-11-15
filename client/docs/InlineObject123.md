# InlineObject123

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product type to rollback (if the network is a combined network) | [optional] 
**Time** | Pointer to **time.Time** | Scheduled time for the rollback | [optional] 
**Reasons** | [**[]NetworksNetworkIdFirmwareUpgradesRollbacksReasons**](NetworksNetworkIdFirmwareUpgradesRollbacksReasons.md) | Reasons for the rollback | 
**ToVersion** | Pointer to [**NetworksNetworkIdFirmwareUpgradesRollbacksToVersion**](NetworksNetworkIdFirmwareUpgradesRollbacksToVersion.md) |  | [optional] 
**Predownload** | Pointer to [**NetworksNetworkIdFirmwareUpgradesRollbacksPredownload**](NetworksNetworkIdFirmwareUpgradesRollbacksPredownload.md) |  | [optional] 

## Methods

### NewInlineObject123

`func NewInlineObject123(reasons []NetworksNetworkIdFirmwareUpgradesRollbacksReasons, ) *InlineObject123`

NewInlineObject123 instantiates a new InlineObject123 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject123WithDefaults

`func NewInlineObject123WithDefaults() *InlineObject123`

NewInlineObject123WithDefaults instantiates a new InlineObject123 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *InlineObject123) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *InlineObject123) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *InlineObject123) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *InlineObject123) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetTime

`func (o *InlineObject123) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineObject123) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineObject123) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineObject123) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetReasons

`func (o *InlineObject123) GetReasons() []NetworksNetworkIdFirmwareUpgradesRollbacksReasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineObject123) GetReasonsOk() (*[]NetworksNetworkIdFirmwareUpgradesRollbacksReasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineObject123) SetReasons(v []NetworksNetworkIdFirmwareUpgradesRollbacksReasons)`

SetReasons sets Reasons field to given value.


### GetToVersion

`func (o *InlineObject123) GetToVersion() NetworksNetworkIdFirmwareUpgradesRollbacksToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineObject123) GetToVersionOk() (*NetworksNetworkIdFirmwareUpgradesRollbacksToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineObject123) SetToVersion(v NetworksNetworkIdFirmwareUpgradesRollbacksToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineObject123) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetPredownload

`func (o *InlineObject123) GetPredownload() NetworksNetworkIdFirmwareUpgradesRollbacksPredownload`

GetPredownload returns the Predownload field if non-nil, zero value otherwise.

### GetPredownloadOk

`func (o *InlineObject123) GetPredownloadOk() (*NetworksNetworkIdFirmwareUpgradesRollbacksPredownload, bool)`

GetPredownloadOk returns a tuple with the Predownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredownload

`func (o *InlineObject123) SetPredownload(v NetworksNetworkIdFirmwareUpgradesRollbacksPredownload)`

SetPredownload sets Predownload field to given value.

### HasPredownload

`func (o *InlineObject123) HasPredownload() bool`

HasPredownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


