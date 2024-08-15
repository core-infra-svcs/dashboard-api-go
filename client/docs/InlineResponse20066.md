# InlineResponse20066

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | ID of the network whose VPN exclusion rules are returned. | 
**NetworkName** | **string** | Name of the network whose VPN exclusion rules are returned. | 
**Custom** | [**[]InlineResponse20066Custom**](InlineResponse20066Custom.md) | Custom VPN exclusion rules. | 
**MajorApplications** | [**[]InlineResponse20066MajorApplications**](InlineResponse20066MajorApplications.md) | Major Application based VPN exclusion rules. | 

## Methods

### NewInlineResponse20066

`func NewInlineResponse20066(networkId string, networkName string, custom []InlineResponse20066Custom, majorApplications []InlineResponse20066MajorApplications, ) *InlineResponse20066`

NewInlineResponse20066 instantiates a new InlineResponse20066 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20066WithDefaults

`func NewInlineResponse20066WithDefaults() *InlineResponse20066`

NewInlineResponse20066WithDefaults instantiates a new InlineResponse20066 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse20066) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20066) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20066) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *InlineResponse20066) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse20066) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse20066) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetCustom

`func (o *InlineResponse20066) GetCustom() []InlineResponse20066Custom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InlineResponse20066) GetCustomOk() (*[]InlineResponse20066Custom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InlineResponse20066) SetCustom(v []InlineResponse20066Custom)`

SetCustom sets Custom field to given value.


### GetMajorApplications

`func (o *InlineResponse20066) GetMajorApplications() []InlineResponse20066MajorApplications`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *InlineResponse20066) GetMajorApplicationsOk() (*[]InlineResponse20066MajorApplications, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *InlineResponse20066) SetMajorApplications(v []InlineResponse20066MajorApplications)`

SetMajorApplications sets MajorApplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


