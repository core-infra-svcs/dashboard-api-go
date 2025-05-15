# InlineResponse200352Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200352Network**](InlineResponse200352Network.md) |  | [optional] 
**Serial** | Pointer to **string** | AP cloud ID | [optional] 
**Controller** | Pointer to [**InlineResponse200352Controller**](InlineResponse200352Controller.md) |  | [optional] 
**JoinedAt** | Pointer to **string** | The time when AP joins wireless controller | [optional] 
**Model** | Pointer to **string** | AP model | [optional] 
**Tags** | Pointer to [**[]InlineResponse200352Tags**](InlineResponse200352Tags.md) | The tags of the catalyst access point | [optional] 
**Mode** | Pointer to **string** | AP mode (local, flex, etc.) | [optional] 
**CountryCode** | Pointer to **string** | Country code (2 characters) | [optional] 
**Details** | Pointer to [**[]InlineResponse200352Details**](InlineResponse200352Details.md) | Catalyst access point details | [optional] 

## Methods

### NewInlineResponse200352Items

`func NewInlineResponse200352Items() *InlineResponse200352Items`

NewInlineResponse200352Items instantiates a new InlineResponse200352Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200352ItemsWithDefaults

`func NewInlineResponse200352ItemsWithDefaults() *InlineResponse200352Items`

NewInlineResponse200352ItemsWithDefaults instantiates a new InlineResponse200352Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200352Items) GetNetwork() InlineResponse200352Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200352Items) GetNetworkOk() (*InlineResponse200352Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200352Items) SetNetwork(v InlineResponse200352Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200352Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200352Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200352Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200352Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200352Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetController

`func (o *InlineResponse200352Items) GetController() InlineResponse200352Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *InlineResponse200352Items) GetControllerOk() (*InlineResponse200352Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *InlineResponse200352Items) SetController(v InlineResponse200352Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *InlineResponse200352Items) HasController() bool`

HasController returns a boolean if a field has been set.

### GetJoinedAt

`func (o *InlineResponse200352Items) GetJoinedAt() string`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *InlineResponse200352Items) GetJoinedAtOk() (*string, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *InlineResponse200352Items) SetJoinedAt(v string)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *InlineResponse200352Items) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200352Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200352Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200352Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200352Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200352Items) GetTags() []InlineResponse200352Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200352Items) GetTagsOk() (*[]InlineResponse200352Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200352Items) SetTags(v []InlineResponse200352Tags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200352Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse200352Items) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200352Items) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200352Items) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200352Items) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCountryCode

`func (o *InlineResponse200352Items) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InlineResponse200352Items) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InlineResponse200352Items) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *InlineResponse200352Items) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200352Items) GetDetails() []InlineResponse200352Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200352Items) GetDetailsOk() (*[]InlineResponse200352Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200352Items) SetDetails(v []InlineResponse200352Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200352Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


