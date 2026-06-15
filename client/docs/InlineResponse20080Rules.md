# InlineResponse20080Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIp** | Pointer to **string** | The IP address that will be used to access the internal resource from the WAN | [optional] 
**Uplink** | Pointer to **string** | The physical WAN interface on which the traffic will arrive, formatted as &#39;internetN&#39; where N is an integer representing a valid uplink for the network&#39;s appliance | [optional] 
**PortRules** | Pointer to [**[]InlineResponse20080PortRules**](InlineResponse20080PortRules.md) | An array of associated forwarding rules | [optional] 

## Methods

### NewInlineResponse20080Rules

`func NewInlineResponse20080Rules() *InlineResponse20080Rules`

NewInlineResponse20080Rules instantiates a new InlineResponse20080Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20080RulesWithDefaults

`func NewInlineResponse20080RulesWithDefaults() *InlineResponse20080Rules`

NewInlineResponse20080RulesWithDefaults instantiates a new InlineResponse20080Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIp

`func (o *InlineResponse20080Rules) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *InlineResponse20080Rules) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *InlineResponse20080Rules) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *InlineResponse20080Rules) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetUplink

`func (o *InlineResponse20080Rules) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *InlineResponse20080Rules) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *InlineResponse20080Rules) SetUplink(v string)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *InlineResponse20080Rules) HasUplink() bool`

HasUplink returns a boolean if a field has been set.

### GetPortRules

`func (o *InlineResponse20080Rules) GetPortRules() []InlineResponse20080PortRules`

GetPortRules returns the PortRules field if non-nil, zero value otherwise.

### GetPortRulesOk

`func (o *InlineResponse20080Rules) GetPortRulesOk() (*[]InlineResponse20080PortRules, bool)`

GetPortRulesOk returns a tuple with the PortRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRules

`func (o *InlineResponse20080Rules) SetPortRules(v []InlineResponse20080PortRules)`

SetPortRules sets PortRules field to given value.

### HasPortRules

`func (o *InlineResponse20080Rules) HasPortRules() bool`

HasPortRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


