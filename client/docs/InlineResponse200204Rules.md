# InlineResponse200204Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Desctiption of the bonjour forwarding rule | [optional] 
**VlanId** | Pointer to **string** | The ID of the service VLAN. Required | [optional] 
**Services** | Pointer to **[]string** | A list of Bonjour services. At least one service must be specified. Available services are &#39;All Services&#39;, &#39;AFP&#39;, &#39;AirPlay&#39;, &#39;Apple screen share&#39;, &#39;BitTorrent&#39;, &#39;Chromecast&#39;, &#39;FTP&#39;, &#39;iChat&#39;, &#39;iTunes&#39;, &#39;Printers&#39;, &#39;Samba&#39;, &#39;Scanners&#39;, &#39;Spotify&#39; and &#39;SSH&#39; | [optional] 

## Methods

### NewInlineResponse200204Rules

`func NewInlineResponse200204Rules() *InlineResponse200204Rules`

NewInlineResponse200204Rules instantiates a new InlineResponse200204Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200204RulesWithDefaults

`func NewInlineResponse200204RulesWithDefaults() *InlineResponse200204Rules`

NewInlineResponse200204RulesWithDefaults instantiates a new InlineResponse200204Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InlineResponse200204Rules) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200204Rules) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200204Rules) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200204Rules) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlanId

`func (o *InlineResponse200204Rules) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *InlineResponse200204Rules) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *InlineResponse200204Rules) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *InlineResponse200204Rules) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetServices

`func (o *InlineResponse200204Rules) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InlineResponse200204Rules) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InlineResponse200204Rules) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *InlineResponse200204Rules) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


