# InlineResponse200344Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**InlineResponse200344Device**](InlineResponse200344Device.md) |  | [optional] 
**Active** | Pointer to **bool** | Whether eSIM is currently active SIM on Device | [optional] 
**Eid** | Pointer to **string** | eSIM EID | [optional] 
**LastUpdatedAt** | Pointer to **string** | Last update of eSIM | [optional] 
**Network** | Pointer to [**InlineResponse200344Network**](InlineResponse200344Network.md) |  | [optional] 
**Profiles** | Pointer to [**[]InlineResponse200344Profiles**](InlineResponse200344Profiles.md) | eSIM Profile Information | [optional] 

## Methods

### NewInlineResponse200344Items

`func NewInlineResponse200344Items() *InlineResponse200344Items`

NewInlineResponse200344Items instantiates a new InlineResponse200344Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200344ItemsWithDefaults

`func NewInlineResponse200344ItemsWithDefaults() *InlineResponse200344Items`

NewInlineResponse200344ItemsWithDefaults instantiates a new InlineResponse200344Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *InlineResponse200344Items) GetDevice() InlineResponse200344Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200344Items) GetDeviceOk() (*InlineResponse200344Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200344Items) SetDevice(v InlineResponse200344Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200344Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse200344Items) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200344Items) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200344Items) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200344Items) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEid

`func (o *InlineResponse200344Items) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *InlineResponse200344Items) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *InlineResponse200344Items) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *InlineResponse200344Items) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200344Items) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200344Items) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200344Items) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200344Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200344Items) GetNetwork() InlineResponse200344Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200344Items) GetNetworkOk() (*InlineResponse200344Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200344Items) SetNetwork(v InlineResponse200344Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200344Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProfiles

`func (o *InlineResponse200344Items) GetProfiles() []InlineResponse200344Profiles`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *InlineResponse200344Items) GetProfilesOk() (*[]InlineResponse200344Profiles, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *InlineResponse200344Items) SetProfiles(v []InlineResponse200344Profiles)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *InlineResponse200344Items) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


