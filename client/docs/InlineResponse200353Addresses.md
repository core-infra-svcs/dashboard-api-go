# InlineResponse200353Addresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Type of address for the device uplink. Available options are: ipv4, ipv6. enum &#x3D; [ipv4, ipv6] | [optional] 
**Address** | Pointer to **string** | The address of the wireless LAN controller interface | [optional] 
**Subnet** | Pointer to **string** | The address of the wireless LAN controller interface | [optional] 

## Methods

### NewInlineResponse200353Addresses

`func NewInlineResponse200353Addresses() *InlineResponse200353Addresses`

NewInlineResponse200353Addresses instantiates a new InlineResponse200353Addresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200353AddressesWithDefaults

`func NewInlineResponse200353AddressesWithDefaults() *InlineResponse200353Addresses`

NewInlineResponse200353AddressesWithDefaults instantiates a new InlineResponse200353Addresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse200353Addresses) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse200353Addresses) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse200353Addresses) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse200353Addresses) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetAddress

`func (o *InlineResponse200353Addresses) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InlineResponse200353Addresses) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InlineResponse200353Addresses) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InlineResponse200353Addresses) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSubnet

`func (o *InlineResponse200353Addresses) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *InlineResponse200353Addresses) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *InlineResponse200353Addresses) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *InlineResponse200353Addresses) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


