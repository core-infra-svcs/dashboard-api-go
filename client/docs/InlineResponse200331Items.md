# InlineResponse200331Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of device | [optional] 
**Name** | Pointer to **string** | Name of device | [optional] 
**Url** | Pointer to **string** | Dashboard URL for the device | [optional] 
**Model** | Pointer to **string** | Product type of device | [optional] 
**Software** | Pointer to [**InlineResponse200331Software**](InlineResponse200331Software.md) |  | [optional] 
**Modems** | Pointer to [**[]InlineResponse200331Modems**](InlineResponse200331Modems.md) | Information about the available modems and SIMs for the device. Maximum 1 modem. | [optional] 
**Profile** | Pointer to [**InlineResponse200331Profile**](InlineResponse200331Profile.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200331Network**](InlineResponse200331Network.md) |  | [optional] 

## Methods

### NewInlineResponse200331Items

`func NewInlineResponse200331Items() *InlineResponse200331Items`

NewInlineResponse200331Items instantiates a new InlineResponse200331Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200331ItemsWithDefaults

`func NewInlineResponse200331ItemsWithDefaults() *InlineResponse200331Items`

NewInlineResponse200331ItemsWithDefaults instantiates a new InlineResponse200331Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200331Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200331Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200331Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200331Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200331Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200331Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200331Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200331Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200331Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200331Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200331Items) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200331Items) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200331Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200331Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200331Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200331Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSoftware

`func (o *InlineResponse200331Items) GetSoftware() InlineResponse200331Software`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *InlineResponse200331Items) GetSoftwareOk() (*InlineResponse200331Software, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *InlineResponse200331Items) SetSoftware(v InlineResponse200331Software)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *InlineResponse200331Items) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetModems

`func (o *InlineResponse200331Items) GetModems() []InlineResponse200331Modems`

GetModems returns the Modems field if non-nil, zero value otherwise.

### GetModemsOk

`func (o *InlineResponse200331Items) GetModemsOk() (*[]InlineResponse200331Modems, bool)`

GetModemsOk returns a tuple with the Modems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModems

`func (o *InlineResponse200331Items) SetModems(v []InlineResponse200331Modems)`

SetModems sets Modems field to given value.

### HasModems

`func (o *InlineResponse200331Items) HasModems() bool`

HasModems returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse200331Items) GetProfile() InlineResponse200331Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse200331Items) GetProfileOk() (*InlineResponse200331Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse200331Items) SetProfile(v InlineResponse200331Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineResponse200331Items) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200331Items) GetNetwork() InlineResponse200331Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200331Items) GetNetworkOk() (*InlineResponse200331Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200331Items) SetNetwork(v InlineResponse200331Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200331Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


