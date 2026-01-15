# InlineResponse20074ValueDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **string** | E.g.: \&quot;any\&quot;, \&quot;0\&quot; (also means \&quot;any\&quot;), \&quot;8080\&quot;, \&quot;1-1024\&quot; | [optional] 
**Cidr** | Pointer to **string** | CIDR format address (e.g.\&quot;192.168.10.1\&quot;, which is the same as \&quot;192.168.10.1/32\&quot;), or \&quot;any\&quot; | [optional] 
**Applications** | Pointer to [**[]NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications**](NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications.md) | list of application objects (either majorApplication or nbar) | [optional] 

## Methods

### NewInlineResponse20074ValueDestination

`func NewInlineResponse20074ValueDestination() *InlineResponse20074ValueDestination`

NewInlineResponse20074ValueDestination instantiates a new InlineResponse20074ValueDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20074ValueDestinationWithDefaults

`func NewInlineResponse20074ValueDestinationWithDefaults() *InlineResponse20074ValueDestination`

NewInlineResponse20074ValueDestinationWithDefaults instantiates a new InlineResponse20074ValueDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *InlineResponse20074ValueDestination) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20074ValueDestination) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20074ValueDestination) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20074ValueDestination) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetCidr

`func (o *InlineResponse20074ValueDestination) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *InlineResponse20074ValueDestination) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *InlineResponse20074ValueDestination) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *InlineResponse20074ValueDestination) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetApplications

`func (o *InlineResponse20074ValueDestination) GetApplications() []NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *InlineResponse20074ValueDestination) GetApplicationsOk() (*[]NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *InlineResponse20074ValueDestination) SetApplications(v []NetworksNetworkIdApplianceSdwanInternetPoliciesValueDestinationApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *InlineResponse20074ValueDestination) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplicationsNil

`func (o *InlineResponse20074ValueDestination) SetApplicationsNil(b bool)`

 SetApplicationsNil sets the value for Applications to be an explicit nil

### UnsetApplications
`func (o *InlineResponse20074ValueDestination) UnsetApplications()`

UnsetApplications ensures that no value is present for Applications, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


