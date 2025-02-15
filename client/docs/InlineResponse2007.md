# InlineResponse2007

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnet** | Pointer to **string** | Subnet | [optional] 
**VlanId** | Pointer to **int32** | VLAN ID | [optional] 
**UsedCount** | Pointer to **int32** | Count of used IP addresses in subnet | [optional] 
**FreeCount** | Pointer to **int32** | Count of free IP addresses in subnet | [optional] 

## Methods

### NewInlineResponse2007

`func NewInlineResponse2007() *InlineResponse2007`

NewInlineResponse2007 instantiates a new InlineResponse2007 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2007WithDefaults

`func NewInlineResponse2007WithDefaults() *InlineResponse2007`

NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnet

`func (o *InlineResponse2007) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse2007) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse2007) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse2007) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse2007) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse2007) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse2007) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse2007) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetUsedCount

`func (o *InlineResponse2007) GetUsedCount() int32`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *InlineResponse2007) GetUsedCountOk() (*int32, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *InlineResponse2007) SetUsedCount(v int32)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *InlineResponse2007) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetFreeCount

`func (o *InlineResponse2007) GetFreeCount() int32`

GetFreeCount returns the FreeCount field if non-nil, zero value otherwise.

### GetFreeCountOk

`func (o *InlineResponse2007) GetFreeCountOk() (*int32, bool)`

GetFreeCountOk returns a tuple with the FreeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeCount

`func (o *InlineResponse2007) SetFreeCount(v int32)`

SetFreeCount sets FreeCount field to given value.

### HasFreeCount

`func (o *InlineResponse2007) HasFreeCount() bool`

HasFreeCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


