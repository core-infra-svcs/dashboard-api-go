# InlineResponse200300Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number | [optional] 
**Uplink** | Pointer to [**InlineResponse200300Uplink**](InlineResponse200300Uplink.md) |  | [optional] 
**Nameservers** | Pointer to [**InlineResponse20113Nameservers**](InlineResponse20113Nameservers.md) |  | [optional] 
**Sgt** | Pointer to **int32** | Infra Security Group Tag(sgt) value for Trustsec | [optional] 

## Methods

### NewInlineResponse200300Items

`func NewInlineResponse200300Items() *InlineResponse200300Items`

NewInlineResponse200300Items instantiates a new InlineResponse200300Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200300ItemsWithDefaults

`func NewInlineResponse200300ItemsWithDefaults() *InlineResponse200300Items`

NewInlineResponse200300ItemsWithDefaults instantiates a new InlineResponse200300Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200300Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200300Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200300Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200300Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200300Items) GetUplink() InlineResponse200300Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200300Items) GetUplinkOk() (*InlineResponse200300Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200300Items) SetUplink(v InlineResponse200300Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200300Items) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse200300Items) GetNameservers() InlineResponse20113Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse200300Items) GetNameserversOk() (*InlineResponse20113Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse200300Items) SetNameservers(v InlineResponse20113Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse200300Items) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetSgt

`func (o *InlineResponse200300Items) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineResponse200300Items) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineResponse200300Items) SetSgt(v int32)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineResponse200300Items) HasSgt() bool`

HasSgt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


