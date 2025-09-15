# InlineResponse20073Value

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to **string** | Protocol of &#39;custom&#39; type traffic filter. Must be one of: &#39;tcp&#39;, &#39;udp&#39;, &#39;icmp6&#39; or &#39;any&#39; | [optional] 
**Source** | [**InlineResponse20073ValueSource**](InlineResponse20073ValueSource.md) |  | 
**Destination** | [**NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination**](NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination.md) |  | 

## Methods

### NewInlineResponse20073Value

`func NewInlineResponse20073Value(source InlineResponse20073ValueSource, destination NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination, ) *InlineResponse20073Value`

NewInlineResponse20073Value instantiates a new InlineResponse20073Value object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20073ValueWithDefaults

`func NewInlineResponse20073ValueWithDefaults() *InlineResponse20073Value`

NewInlineResponse20073ValueWithDefaults instantiates a new InlineResponse20073Value object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *InlineResponse20073Value) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *InlineResponse20073Value) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *InlineResponse20073Value) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *InlineResponse20073Value) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSource

`func (o *InlineResponse20073Value) GetSource() InlineResponse20073ValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *InlineResponse20073Value) GetSourceOk() (*InlineResponse20073ValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *InlineResponse20073Value) SetSource(v InlineResponse20073ValueSource)`

SetSource sets Source field to given value.


### GetDestination

`func (o *InlineResponse20073Value) GetDestination() NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InlineResponse20073Value) GetDestinationOk() (*NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InlineResponse20073Value) SetDestination(v NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


