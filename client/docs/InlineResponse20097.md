# InlineResponse20097

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** | Product type to rollback (if the network is a combined network) | [optional] 
**Status** | Pointer to **string** | Status of the rollback | [optional] 
**UpgradeBatchId** | Pointer to **string** | Batch ID of the firmware rollback | [optional] 
**Time** | Pointer to **time.Time** | Scheduled time for the rollback | [optional] 
**ToVersion** | Pointer to [**InlineResponse20097ToVersion**](InlineResponse20097ToVersion.md) |  | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20097Reasons**](InlineResponse20097Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20097

`func NewInlineResponse20097() *InlineResponse20097`

NewInlineResponse20097 instantiates a new InlineResponse20097 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20097WithDefaults

`func NewInlineResponse20097WithDefaults() *InlineResponse20097`

NewInlineResponse20097WithDefaults instantiates a new InlineResponse20097 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *InlineResponse20097) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *InlineResponse20097) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *InlineResponse20097) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *InlineResponse20097) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20097) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20097) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20097) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20097) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpgradeBatchId

`func (o *InlineResponse20097) GetUpgradeBatchId() string`

GetUpgradeBatchId returns the UpgradeBatchId field if non-nil, zero value otherwise.

### GetUpgradeBatchIdOk

`func (o *InlineResponse20097) GetUpgradeBatchIdOk() (*string, bool)`

GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeBatchId

`func (o *InlineResponse20097) SetUpgradeBatchId(v string)`

SetUpgradeBatchId sets UpgradeBatchId field to given value.

### HasUpgradeBatchId

`func (o *InlineResponse20097) HasUpgradeBatchId() bool`

HasUpgradeBatchId returns a boolean if a field has been set.

### GetTime

`func (o *InlineResponse20097) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *InlineResponse20097) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *InlineResponse20097) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *InlineResponse20097) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetToVersion

`func (o *InlineResponse20097) GetToVersion() InlineResponse20097ToVersion`

GetToVersion returns the ToVersion field if non-nil, zero value otherwise.

### GetToVersionOk

`func (o *InlineResponse20097) GetToVersionOk() (*InlineResponse20097ToVersion, bool)`

GetToVersionOk returns a tuple with the ToVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToVersion

`func (o *InlineResponse20097) SetToVersion(v InlineResponse20097ToVersion)`

SetToVersion sets ToVersion field to given value.

### HasToVersion

`func (o *InlineResponse20097) HasToVersion() bool`

HasToVersion returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20097) GetReasons() []InlineResponse20097Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20097) GetReasonsOk() (*[]InlineResponse20097Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20097) SetReasons(v []InlineResponse20097Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20097) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


