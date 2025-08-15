# InlineResponse200256Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to [**InlineResponse200256Device**](InlineResponse200256Device.md) |  | [optional] 
**Active** | Pointer to **bool** | Whether eSIM is currently active SIM on Device | [optional] 
**Eid** | Pointer to **string** | eSIM EID | [optional] 
**LastUpdatedAt** | Pointer to **string** | Last update of eSIM | [optional] 
**Network** | Pointer to [**InlineResponse200256Network**](InlineResponse200256Network.md) |  | [optional] 
**Profiles** | Pointer to [**[]InlineResponse200256Profiles**](InlineResponse200256Profiles.md) | eSIM Profile Information | [optional] 

## Methods

### NewInlineResponse200256Items

`func NewInlineResponse200256Items() *InlineResponse200256Items`

NewInlineResponse200256Items instantiates a new InlineResponse200256Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200256ItemsWithDefaults

`func NewInlineResponse200256ItemsWithDefaults() *InlineResponse200256Items`

NewInlineResponse200256ItemsWithDefaults instantiates a new InlineResponse200256Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *InlineResponse200256Items) GetDevice() InlineResponse200256Device`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *InlineResponse200256Items) GetDeviceOk() (*InlineResponse200256Device, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *InlineResponse200256Items) SetDevice(v InlineResponse200256Device)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *InlineResponse200256Items) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetActive

`func (o *InlineResponse200256Items) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InlineResponse200256Items) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InlineResponse200256Items) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InlineResponse200256Items) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEid

`func (o *InlineResponse200256Items) GetEid() string`

GetEid returns the Eid field if non-nil, zero value otherwise.

### GetEidOk

`func (o *InlineResponse200256Items) GetEidOk() (*string, bool)`

GetEidOk returns a tuple with the Eid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEid

`func (o *InlineResponse200256Items) SetEid(v string)`

SetEid sets Eid field to given value.

### HasEid

`func (o *InlineResponse200256Items) HasEid() bool`

HasEid returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *InlineResponse200256Items) GetLastUpdatedAt() string`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *InlineResponse200256Items) GetLastUpdatedAtOk() (*string, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *InlineResponse200256Items) SetLastUpdatedAt(v string)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *InlineResponse200256Items) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200256Items) GetNetwork() InlineResponse200256Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200256Items) GetNetworkOk() (*InlineResponse200256Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200256Items) SetNetwork(v InlineResponse200256Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200256Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProfiles

`func (o *InlineResponse200256Items) GetProfiles() []InlineResponse200256Profiles`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *InlineResponse200256Items) GetProfilesOk() (*[]InlineResponse200256Profiles, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *InlineResponse200256Items) SetProfiles(v []InlineResponse200256Profiles)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *InlineResponse200256Items) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


