# InlineResponse200429Firewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of network enforcement target | 
**Direction** | **string** | Traffic direction that rules will be applied to. Read-only attribute derived from the firewall type | 

## Methods

### NewInlineResponse200429Firewall

`func NewInlineResponse200429Firewall(type_ string, direction string, ) *InlineResponse200429Firewall`

NewInlineResponse200429Firewall instantiates a new InlineResponse200429Firewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200429FirewallWithDefaults

`func NewInlineResponse200429FirewallWithDefaults() *InlineResponse200429Firewall`

NewInlineResponse200429FirewallWithDefaults instantiates a new InlineResponse200429Firewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InlineResponse200429Firewall) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse200429Firewall) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse200429Firewall) SetType(v string)`

SetType sets Type field to given value.


### GetDirection

`func (o *InlineResponse200429Firewall) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InlineResponse200429Firewall) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InlineResponse200429Firewall) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


