# InlineResponse200126

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product type to rollback (if the network is a combined network) | [optional] 
**Status** | Pointer to **string** | Status of the rollback | [optional] 
**UpgradeBatchId** | Pointer to **string** | Batch ID of the firmware rollback | [optional] 
**Time** | Pointer to **time.Time** | Scheduled time for the rollback | [optional] 
**ToVersion** | Pointer to [**InlineResponse200126ToVersion**](InlineResponse200126ToVersion.md) |  | [optional] 
**Reasons** | Pointer to [**[]InlineResponse200126Reasons**](InlineResponse200126Reasons.md) | Reasons for the rollback | [optional] 
**Predownload** | Pointer to [**NetworksNetworkIdFirmwareUpgradesRollbacksPredownload**](NetworksNetworkIdFirmwareUpgradesRollbacksPredownload.md) |  | [optional] 

## Methods

### NewInlineResponse200126

`func NewInlineResponse200126() *InlineResponse200126`

NewInlineResponse200126 instantiates a new InlineResponse200126 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200126WithDefaults

`func NewInlineResponse200126WithDefaults() *InlineResponse200126`

NewInlineResponse200126WithDefaults instantiates a new InlineResponse200126 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *InlineResponse200126) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *InlineResponse200126) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *InlineResponse200126) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *InlineResponse200126) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200126) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200126) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200126) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200126) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradeBatchId

`func (o *InlineResponse200126) GetUpgradeBatchId() string`

GetUpgradeBatchId returns the UpgradeBatchId field if non-nil, zero value otherwise.

### GetUpgradeBatchIdOk

`func (o *InlineResponse200126) GetUpgradeBatchIdOk() (*string, bool)`

GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBatchId

`func (o *InlineResponse200126) SetUpgradeBatchId(v string)`

SetUpgradeBatchId sets UpgradeBatchId field to given value.

### HasUpgradeBatchId

`func (o *InlineResponse200126) HasUpgradeBatchId() bool`

HasUpgradeBatchId returns a boolean if a field has been set.

### GetTime

`func (o *InlineResponse200126) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse200126) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse200126) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse200126) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse200126) GetToVersion() InlineResponse200126ToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse200126) GetToVersionOk() (*InlineResponse200126ToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse200126) SetToVersion(v InlineResponse200126ToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse200126) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse200126) GetReasons() []InlineResponse200126Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse200126) GetReasonsOk() (*[]InlineResponse200126Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse200126) SetReasons(v []InlineResponse200126Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse200126) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetPredownload

`func (o *InlineResponse200126) GetPredownload() NetworksNetworkIdFirmwareUpgradesRollbacksPredownload`

GetPredownload returns the Predownload field if non-nil, zero value otherwise.

### GetPredownloadOk

`func (o *InlineResponse200126) GetPredownloadOk() (*NetworksNetworkIdFirmwareUpgradesRollbacksPredownload, bool)`

GetPredownloadOk returns a tuple with the Predownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredownload

`func (o *InlineResponse200126) SetPredownload(v NetworksNetworkIdFirmwareUpgradesRollbacksPredownload)`

SetPredownload sets Predownload field to given value.

### HasPredownload

`func (o *InlineResponse200126) HasPredownload() bool`

HasPredownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


