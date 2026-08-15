# InlineResponse200343Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of device | [optional] 
**Name** | Pointer to **string** | Name of device | [optional] 
**Url** | Pointer to **string** | Dashboard URL for the device | [optional] 
**Model** | Pointer to **string** | Product type of device | [optional] 
**Software** | Pointer to [**InlineResponse200343Software**](InlineResponse200343Software.md) |  | [optional] 
**Modems** | Pointer to [**[]InlineResponse200343Modems**](InlineResponse200343Modems.md) | Information about the available modems and SIMs for the device. Maximum 1 modem. | [optional] 
**Profile** | Pointer to [**InlineResponse200343Profile**](InlineResponse200343Profile.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200343Network**](InlineResponse200343Network.md) |  | [optional] 

## Methods

### NewInlineResponse200343Items

`func NewInlineResponse200343Items() *InlineResponse200343Items`

NewInlineResponse200343Items instantiates a new InlineResponse200343Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200343ItemsWithDefaults

`func NewInlineResponse200343ItemsWithDefaults() *InlineResponse200343Items`

NewInlineResponse200343ItemsWithDefaults instantiates a new InlineResponse200343Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200343Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200343Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200343Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200343Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200343Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200343Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200343Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200343Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200343Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200343Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200343Items) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200343Items) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200343Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200343Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200343Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200343Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSoftware

`func (o *InlineResponse200343Items) GetSoftware() InlineResponse200343Software`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *InlineResponse200343Items) GetSoftwareOk() (*InlineResponse200343Software, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *InlineResponse200343Items) SetSoftware(v InlineResponse200343Software)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *InlineResponse200343Items) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetModems

`func (o *InlineResponse200343Items) GetModems() []InlineResponse200343Modems`

GetModems returns the Modems field if non-nil, zero value otherwise.

### GetModemsOk

`func (o *InlineResponse200343Items) GetModemsOk() (*[]InlineResponse200343Modems, bool)`

GetModemsOk returns a tuple with the Modems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModems

`func (o *InlineResponse200343Items) SetModems(v []InlineResponse200343Modems)`

SetModems sets Modems field to given value.

### HasModems

`func (o *InlineResponse200343Items) HasModems() bool`

HasModems returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse200343Items) GetProfile() InlineResponse200343Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse200343Items) GetProfileOk() (*InlineResponse200343Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse200343Items) SetProfile(v InlineResponse200343Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineResponse200343Items) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200343Items) GetNetwork() InlineResponse200343Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200343Items) GetNetworkOk() (*InlineResponse200343Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200343Items) SetNetwork(v InlineResponse200343Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200343Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


