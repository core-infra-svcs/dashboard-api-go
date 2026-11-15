# InlineResponse200372Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Serial of device | [optional] 
**Name** | Pointer to **string** | Name of device | [optional] 
**Url** | Pointer to **string** | Dashboard URL for the device | [optional] 
**Model** | Pointer to **string** | Product type of device | [optional] 
**Software** | Pointer to [**InlineResponse200372Software**](InlineResponse200372Software.md) |  | [optional] 
**Modems** | Pointer to [**[]InlineResponse200372Modems**](InlineResponse200372Modems.md) | Information about the available modems and SIMs for the device. Maximum 1 modem. | [optional] 
**Profile** | Pointer to [**InlineResponse200372Profile**](InlineResponse200372Profile.md) |  | [optional] 
**Network** | Pointer to [**InlineResponse200372Network**](InlineResponse200372Network.md) |  | [optional] 

## Methods

### NewInlineResponse200372Items

`func NewInlineResponse200372Items() *InlineResponse200372Items`

NewInlineResponse200372Items instantiates a new InlineResponse200372Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200372ItemsWithDefaults

`func NewInlineResponse200372ItemsWithDefaults() *InlineResponse200372Items`

NewInlineResponse200372ItemsWithDefaults instantiates a new InlineResponse200372Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200372Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200372Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200372Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200372Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200372Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200372Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200372Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200372Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse200372Items) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse200372Items) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse200372Items) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse200372Items) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200372Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200372Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200372Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200372Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSoftware

`func (o *InlineResponse200372Items) GetSoftware() InlineResponse200372Software`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *InlineResponse200372Items) GetSoftwareOk() (*InlineResponse200372Software, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *InlineResponse200372Items) SetSoftware(v InlineResponse200372Software)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *InlineResponse200372Items) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetModems

`func (o *InlineResponse200372Items) GetModems() []InlineResponse200372Modems`

GetModems returns the Modems field if non-nil, zero value otherwise.

### GetModemsOk

`func (o *InlineResponse200372Items) GetModemsOk() (*[]InlineResponse200372Modems, bool)`

GetModemsOk returns a tuple with the Modems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModems

`func (o *InlineResponse200372Items) SetModems(v []InlineResponse200372Modems)`

SetModems sets Modems field to given value.

### HasModems

`func (o *InlineResponse200372Items) HasModems() bool`

HasModems returns a boolean if a field has been set.

### GetProfile

`func (o *InlineResponse200372Items) GetProfile() InlineResponse200372Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *InlineResponse200372Items) GetProfileOk() (*InlineResponse200372Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *InlineResponse200372Items) SetProfile(v InlineResponse200372Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *InlineResponse200372Items) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200372Items) GetNetwork() InlineResponse200372Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200372Items) GetNetworkOk() (*InlineResponse200372Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200372Items) SetNetwork(v InlineResponse200372Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200372Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


