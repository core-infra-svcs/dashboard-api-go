# InlineResponse200248Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**InlineResponse200248Device**](InlineResponse200248Device.md) |  | [optional] 
**Active** | Pointer to **bool** | Whether eSIM is currently active SIM on Device | [optional] 
**Eid** | Pointer to **string** | eSIM EID | [optional] 
**LastUpdatedAt** | Pointer to **string** | Last update of eSIM | [optional] 
**Network** | Pointer to [**InlineResponse200248Network**](InlineResponse200248Network.md) |  | [optional] 
**Profiles** | Pointer to [**[]InlineResponse200248Profiles**](InlineResponse200248Profiles.md) | eSIM Profile Information | [optional] 

## Methods

### NewInlineResponse200248Items

`func NewInlineResponse200248Items() *InlineResponse200248Items`

NewInlineResponse200248Items instantiates a new InlineResponse200248Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200248ItemsWithDefaults

`func NewInlineResponse200248ItemsWithDefaults() *InlineResponse200248Items`

NewInlineResponse200248ItemsWithDefaults instantiates a new InlineResponse200248Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *InlineResponse200248Items) GetDevice() InlineResponse200248Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200248Items) GetDeviceOk() (*InlineResponse200248Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200248Items) SetDevice(v InlineResponse200248Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200248Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse200248Items) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200248Items) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200248Items) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200248Items) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEid

`func (o *InlineResponse200248Items) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *InlineResponse200248Items) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *InlineResponse200248Items) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *InlineResponse200248Items) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200248Items) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200248Items) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200248Items) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200248Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200248Items) GetNetwork() InlineResponse200248Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200248Items) GetNetworkOk() (*InlineResponse200248Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200248Items) SetNetwork(v InlineResponse200248Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200248Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProfiles

`func (o *InlineResponse200248Items) GetProfiles() []InlineResponse200248Profiles`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *InlineResponse200248Items) GetProfilesOk() (*[]InlineResponse200248Profiles, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *InlineResponse200248Items) SetProfiles(v []InlineResponse200248Profiles)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *InlineResponse200248Items) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


