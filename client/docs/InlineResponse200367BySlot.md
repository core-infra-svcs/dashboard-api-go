# InlineResponse200367BySlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slot** | Pointer to **string** | Identifier for the logical SIM slot. Accepted values are sim1, sim2, and esim. | [optional] 
**IsActive** | Pointer to **bool** | Whether a SIM is installed in the slot or the eSIM has an active profile. | [optional] 
**Total** | Pointer to **string** | Current data usage on the SIM as a string-encoded byte count | [optional] 
**LastUpdatedAt** | Pointer to **string** | Timestamp of when the usage was last reported from the device | [optional] 
**StartTs** | Pointer to **string** | Start of the current usage period for the device | [optional] 
**EndTs** | Pointer to **string** | End of the current usage period for the device | [optional] 
**Limit** | Pointer to **string** | Data limit for the SIM as a string-encoded byte count, if applicable | [optional] 

## Methods

### NewInlineResponse200367BySlot

`func NewInlineResponse200367BySlot() *InlineResponse200367BySlot`

NewInlineResponse200367BySlot instantiates a new InlineResponse200367BySlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200367BySlotWithDefaults

`func NewInlineResponse200367BySlotWithDefaults() *InlineResponse200367BySlot`

NewInlineResponse200367BySlotWithDefaults instantiates a new InlineResponse200367BySlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlot

`func (o *InlineResponse200367BySlot) GetSlot() string`

GetSlot returns the Slot field if non-nil, zero value otherwise.

### GetSlotOk

`func (o *InlineResponse200367BySlot) GetSlotOk() (*string, bool)`

GetSlotOk returns a tuple with the Slot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlot

`func (o *InlineResponse200367BySlot) SetSlot(v string)`

SetSlot sets Slot field to given value.

### HasSlot

`func (o *InlineResponse200367BySlot) HasSlot() bool`

HasSlot returns a boolean if a field has been set.

### GetIsActive

`func (o *InlineResponse200367BySlot) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *InlineResponse200367BySlot) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *InlineResponse200367BySlot) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *InlineResponse200367BySlot) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetTotal

`func (o *InlineResponse200367BySlot) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InlineResponse200367BySlot) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InlineResponse200367BySlot) SetTotal(v string)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InlineResponse200367BySlot) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200367BySlot) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200367BySlot) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200367BySlot) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200367BySlot) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetStartTs

`func (o *InlineResponse200367BySlot) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200367BySlot) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200367BySlot) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200367BySlot) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200367BySlot) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200367BySlot) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200367BySlot) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200367BySlot) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetLimit

`func (o *InlineResponse200367BySlot) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *InlineResponse200367BySlot) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *InlineResponse200367BySlot) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *InlineResponse200367BySlot) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


