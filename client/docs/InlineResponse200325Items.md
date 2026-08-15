# InlineResponse200325Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial number | [optional] 
**Uplink** | Pointer to [**InlineResponse200325Uplink**](InlineResponse200325Uplink.md) |  | [optional] 
**Nameservers** | Pointer to [**InlineResponse20118Nameservers**](InlineResponse20118Nameservers.md) |  | [optional] 
**Sgt** | Pointer to **int32** | Infra Security Group Tag(sgt) value for Trustsec | [optional] 

## Methods

### NewInlineResponse200325Items

`func NewInlineResponse200325Items() *InlineResponse200325Items`

NewInlineResponse200325Items instantiates a new InlineResponse200325Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200325ItemsWithDefaults

`func NewInlineResponse200325ItemsWithDefaults() *InlineResponse200325Items`

NewInlineResponse200325ItemsWithDefaults instantiates a new InlineResponse200325Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200325Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200325Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200325Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200325Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse200325Items) GetUplink() InlineResponse200325Uplink`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse200325Items) GetUplinkOk() (*InlineResponse200325Uplink, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse200325Items) SetUplink(v InlineResponse200325Uplink)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse200325Items) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetNameservers

`func (o *InlineResponse200325Items) GetNameservers() InlineResponse20118Nameservers`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *InlineResponse200325Items) GetNameserversOk() (*InlineResponse20118Nameservers, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *InlineResponse200325Items) SetNameservers(v InlineResponse20118Nameservers)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *InlineResponse200325Items) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetSgt

`func (o *InlineResponse200325Items) GetSgt() int32`

GetSgt returns the Sgt field if non-nil, zero value otherwise.

### GetSgtOk

`func (o *InlineResponse200325Items) GetSgtOk() (*int32, bool)`

GetSgtOk returns a tuple with the Sgt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgt

`func (o *InlineResponse200325Items) SetSgt(v int32)`

SetSgt sets Sgt field to given value.

### HasSgt

`func (o *InlineResponse200325Items) HasSgt() bool`

HasSgt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


