# InlineResponse20067

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | ID of the network whose VPN exclusion rules are returned. | 
**NetworkName** | **string** | Name of the network whose VPN exclusion rules are returned. | 
**Custom** | [**[]InlineResponse20067Custom**](InlineResponse20067Custom.md) | Custom VPN exclusion rules. | 
**MajorApplications** | [**[]InlineResponse20067MajorApplications**](InlineResponse20067MajorApplications.md) | Major Application based VPN exclusion rules. | 

## Methods

### NewInlineResponse20067

`func NewInlineResponse20067(networkId string, networkName string, custom []InlineResponse20067Custom, majorApplications []InlineResponse20067MajorApplications, ) *InlineResponse20067`

NewInlineResponse20067 instantiates a new InlineResponse20067 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20067WithDefaults

`func NewInlineResponse20067WithDefaults() *InlineResponse20067`

NewInlineResponse20067WithDefaults instantiates a new InlineResponse20067 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse20067) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20067) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20067) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *InlineResponse20067) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse20067) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse20067) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetCustom

`func (o *InlineResponse20067) GetCustom() []InlineResponse20067Custom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InlineResponse20067) GetCustomOk() (*[]InlineResponse20067Custom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InlineResponse20067) SetCustom(v []InlineResponse20067Custom)`

SetCustom sets Custom field to given value.


### GetMajorApplications

`func (o *InlineResponse20067) GetMajorApplications() []InlineResponse20067MajorApplications`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *InlineResponse20067) GetMajorApplicationsOk() (*[]InlineResponse20067MajorApplications, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *InlineResponse20067) SetMajorApplications(v []InlineResponse20067MajorApplications)`

SetMajorApplications sets MajorApplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


