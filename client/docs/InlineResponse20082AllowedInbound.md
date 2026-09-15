# InlineResponse20082AllowedInbound

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Either of the following: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp-ping&#39; or &#39;any&#39; | [optional] 
**DestinationPorts** | Pointer to **[]string** | An array of ports or port ranges that will be forwarded to the host on the LAN | [optional] 
**AllowedIps** | Pointer to **[]string** | An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or &#39;any&#39; | [optional] 

## Methods

### NewInlineResponse20082AllowedInbound

`func NewInlineResponse20082AllowedInbound() *InlineResponse20082AllowedInbound`

NewInlineResponse20082AllowedInbound instantiates a new InlineResponse20082AllowedInbound object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20082AllowedInboundWithDefaults

`func NewInlineResponse20082AllowedInboundWithDefaults() *InlineResponse20082AllowedInbound`

NewInlineResponse20082AllowedInboundWithDefaults instantiates a new InlineResponse20082AllowedInbound object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20082AllowedInbound) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20082AllowedInbound) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20082AllowedInbound) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20082AllowedInbound) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetDestinationPorts

`func (o *InlineResponse20082AllowedInbound) GetDestinationPorts() []string`

GetDestinationPorts returns the DestinationPorts field if non-nil, zero value otherwise.

### GetDestinationPortsOk

`func (o *InlineResponse20082AllowedInbound) GetDestinationPortsOk() (*[]string, bool)`

GetDestinationPortsOk returns a tuple with the DestinationPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPorts

`func (o *InlineResponse20082AllowedInbound) SetDestinationPorts(v []string)`

SetDestinationPorts sets DestinationPorts field to given value.

### HasDestinationPorts

`func (o *InlineResponse20082AllowedInbound) HasDestinationPorts() bool`

HasDestinationPorts returns a boolean if a field has been set.

### GetAllowedIps

`func (o *InlineResponse20082AllowedInbound) GetAllowedIps() []string`

GetAllowedIps returns the AllowedIps field if non-nil, zero value otherwise.

### GetAllowedIpsOk

`func (o *InlineResponse20082AllowedInbound) GetAllowedIpsOk() (*[]string, bool)`

GetAllowedIpsOk returns a tuple with the AllowedIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIps

`func (o *InlineResponse20082AllowedInbound) SetAllowedIps(v []string)`

SetAllowedIps sets AllowedIps field to given value.

### HasAllowedIps

`func (o *InlineResponse20082AllowedInbound) HasAllowedIps() bool`

HasAllowedIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


