# InlineResponse200111Hubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HubId** | Pointer to **string** | The network ID of the hub. | [optional] 
**UseDefaultRoute** | Pointer to **bool** | Indicates whether default route traffic should be sent to this hub. | [optional] 
**Vrfs** | Pointer to [**[]InlineResponse200111Vrfs**](InlineResponse200111Vrfs.md) | The VRFs associated with this default-route hub. | [optional] 

## Methods

### NewInlineResponse200111Hubs

`func NewInlineResponse200111Hubs() *InlineResponse200111Hubs`

NewInlineResponse200111Hubs instantiates a new InlineResponse200111Hubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200111HubsWithDefaults

`func NewInlineResponse200111HubsWithDefaults() *InlineResponse200111Hubs`

NewInlineResponse200111HubsWithDefaults instantiates a new InlineResponse200111Hubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHubId

`func (o *InlineResponse200111Hubs) GetHubId() string`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *InlineResponse200111Hubs) GetHubIdOk() (*string, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubId

`func (o *InlineResponse200111Hubs) SetHubId(v string)`

SetHubId sets HubId field to given value.

### HasHubId

`func (o *InlineResponse200111Hubs) HasHubId() bool`

HasHubId returns a boolean if a field has been set.

### GetUseDefaultRoute

`func (o *InlineResponse200111Hubs) GetUseDefaultRoute() bool`

GetUseDefaultRoute returns the UseDefaultRoute field if non-nil, zero value otherwise.

### GetUseDefaultRouteOk

`func (o *InlineResponse200111Hubs) GetUseDefaultRouteOk() (*bool, bool)`

GetUseDefaultRouteOk returns a tuple with the UseDefaultRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultRoute

`func (o *InlineResponse200111Hubs) SetUseDefaultRoute(v bool)`

SetUseDefaultRoute sets UseDefaultRoute field to given value.

### HasUseDefaultRoute

`func (o *InlineResponse200111Hubs) HasUseDefaultRoute() bool`

HasUseDefaultRoute returns a boolean if a field has been set.

### GetVrfs

`func (o *InlineResponse200111Hubs) GetVrfs() []InlineResponse200111Vrfs`

GetVrfs returns the Vrfs field if non-nil, zero value otherwise.

### GetVrfsOk

`func (o *InlineResponse200111Hubs) GetVrfsOk() (*[]InlineResponse200111Vrfs, bool)`

GetVrfsOk returns a tuple with the Vrfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfs

`func (o *InlineResponse200111Hubs) SetVrfs(v []InlineResponse200111Vrfs)`

SetVrfs sets Vrfs field to given value.

### HasVrfs

`func (o *InlineResponse200111Hubs) HasVrfs() bool`

HasVrfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


