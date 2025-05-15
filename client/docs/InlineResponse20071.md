# InlineResponse20071

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkId** | **string** | ID of the network whose VPN exclusion rules are returned. | 
**NetworkName** | **string** | Name of the network whose VPN exclusion rules are returned. | 
**Custom** | [**[]InlineResponse20071Custom**](InlineResponse20071Custom.md) | Custom VPN exclusion rules. | 
**MajorApplications** | [**[]InlineResponse20071MajorApplications**](InlineResponse20071MajorApplications.md) | Major Application based VPN exclusion rules. | 

## Methods

### NewInlineResponse20071

`func NewInlineResponse20071(networkId string, networkName string, custom []InlineResponse20071Custom, majorApplications []InlineResponse20071MajorApplications, ) *InlineResponse20071`

NewInlineResponse20071 instantiates a new InlineResponse20071 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20071WithDefaults

`func NewInlineResponse20071WithDefaults() *InlineResponse20071`

NewInlineResponse20071WithDefaults instantiates a new InlineResponse20071 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkId

`func (o *InlineResponse20071) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *InlineResponse20071) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *InlineResponse20071) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetNetworkName

`func (o *InlineResponse20071) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *InlineResponse20071) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *InlineResponse20071) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetCustom

`func (o *InlineResponse20071) GetCustom() []InlineResponse20071Custom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *InlineResponse20071) GetCustomOk() (*[]InlineResponse20071Custom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *InlineResponse20071) SetCustom(v []InlineResponse20071Custom)`

SetCustom sets Custom field to given value.


### GetMajorApplications

`func (o *InlineResponse20071) GetMajorApplications() []InlineResponse20071MajorApplications`

GetMajorApplications returns the MajorApplications field if non-nil, zero value otherwise.

### GetMajorApplicationsOk

`func (o *InlineResponse20071) GetMajorApplicationsOk() (*[]InlineResponse20071MajorApplications, bool)`

GetMajorApplicationsOk returns a tuple with the MajorApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorApplications

`func (o *InlineResponse20071) SetMajorApplications(v []InlineResponse20071MajorApplications)`

SetMajorApplications sets MajorApplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


