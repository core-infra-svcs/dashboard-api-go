# InlineResponse200314Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**InlineResponse200314Device**](InlineResponse200314Device.md) |  | [optional] 
**Active** | Pointer to **bool** | Whether eSIM is currently active SIM on Device | [optional] 
**Eid** | Pointer to **string** | eSIM EID | [optional] 
**LastUpdatedAt** | Pointer to **string** | Last update of eSIM | [optional] 
**Network** | Pointer to [**InlineResponse200314Network**](InlineResponse200314Network.md) |  | [optional] 
**Profiles** | Pointer to [**[]InlineResponse200314Profiles**](InlineResponse200314Profiles.md) | eSIM Profile Information | [optional] 

## Methods

### NewInlineResponse200314Items

`func NewInlineResponse200314Items() *InlineResponse200314Items`

NewInlineResponse200314Items instantiates a new InlineResponse200314Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200314ItemsWithDefaults

`func NewInlineResponse200314ItemsWithDefaults() *InlineResponse200314Items`

NewInlineResponse200314ItemsWithDefaults instantiates a new InlineResponse200314Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *InlineResponse200314Items) GetDevice() InlineResponse200314Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200314Items) GetDeviceOk() (*InlineResponse200314Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200314Items) SetDevice(v InlineResponse200314Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200314Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse200314Items) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200314Items) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200314Items) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200314Items) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEid

`func (o *InlineResponse200314Items) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *InlineResponse200314Items) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *InlineResponse200314Items) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *InlineResponse200314Items) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200314Items) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200314Items) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200314Items) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200314Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200314Items) GetNetwork() InlineResponse200314Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200314Items) GetNetworkOk() (*InlineResponse200314Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200314Items) SetNetwork(v InlineResponse200314Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200314Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProfiles

`func (o *InlineResponse200314Items) GetProfiles() []InlineResponse200314Profiles`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *InlineResponse200314Items) GetProfilesOk() (*[]InlineResponse200314Profiles, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *InlineResponse200314Items) SetProfiles(v []InlineResponse200314Profiles)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *InlineResponse200314Items) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


