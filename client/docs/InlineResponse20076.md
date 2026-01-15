# InlineResponse20076

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | Pointer to **time.Time** | Start time of interval | [optional] 
**EndTime** | Pointer to **time.Time** | End time of interval | [optional] 
**ByInterface** | Pointer to [**[]NetworksNetworkIdApplianceUplinksUsageHistoryByInterface**](NetworksNetworkIdApplianceUplinksUsageHistoryByInterface.md) | List of usage data for each interface | [optional] 

## Methods

### NewInlineResponse20076

`func NewInlineResponse20076() *InlineResponse20076`

NewInlineResponse20076 instantiates a new InlineResponse20076 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20076WithDefaults

`func NewInlineResponse20076WithDefaults() *InlineResponse20076`

NewInlineResponse20076WithDefaults instantiates a new InlineResponse20076 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *InlineResponse20076) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InlineResponse20076) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InlineResponse20076) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InlineResponse20076) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *InlineResponse20076) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *InlineResponse20076) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *InlineResponse20076) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *InlineResponse20076) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetByInterface

`func (o *InlineResponse20076) GetByInterface() []NetworksNetworkIdApplianceUplinksUsageHistoryByInterface`

GetByInterface returns the ByInterface field if non-nil, zero value otherwise.

### GetByInterfaceOk

`func (o *InlineResponse20076) GetByInterfaceOk() (*[]NetworksNetworkIdApplianceUplinksUsageHistoryByInterface, bool)`

GetByInterfaceOk returns a tuple with the ByInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByInterface

`func (o *InlineResponse20076) SetByInterface(v []NetworksNetworkIdApplianceUplinksUsageHistoryByInterface)`

SetByInterface sets ByInterface field to given value.

### HasByInterface

`func (o *InlineResponse20076) HasByInterface() bool`

HasByInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


