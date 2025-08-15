# InlineResponse200366Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**InlineResponse200366Network**](InlineResponse200366Network.md) |  | [optional] 
**Serial** | Pointer to **string** | AP cloud ID | [optional] 
**Controller** | Pointer to [**InlineResponse200366Controller**](InlineResponse200366Controller.md) |  | [optional] 
**JoinedAt** | Pointer to **string** | The time when AP joins wireless controller | [optional] 
**Model** | Pointer to **string** | AP model | [optional] 
**Tags** | Pointer to [**[]InlineResponse200366Tags**](InlineResponse200366Tags.md) | The tags of the catalyst access point | [optional] 
**Mode** | Pointer to **string** | AP mode (local, flex, etc.) | [optional] 
**CountryCode** | Pointer to **string** | Country code (2 characters) | [optional] 
**Details** | Pointer to [**[]InlineResponse200366Details**](InlineResponse200366Details.md) | Catalyst access point details | [optional] 

## Methods

### NewInlineResponse200366Items

`func NewInlineResponse200366Items() *InlineResponse200366Items`

NewInlineResponse200366Items instantiates a new InlineResponse200366Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200366ItemsWithDefaults

`func NewInlineResponse200366ItemsWithDefaults() *InlineResponse200366Items`

NewInlineResponse200366ItemsWithDefaults instantiates a new InlineResponse200366Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *InlineResponse200366Items) GetNetwork() InlineResponse200366Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200366Items) GetNetworkOk() (*InlineResponse200366Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200366Items) SetNetwork(v InlineResponse200366Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200366Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSerial

`func (o *InlineResponse200366Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200366Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200366Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200366Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetController

`func (o *InlineResponse200366Items) GetController() InlineResponse200366Controller`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *InlineResponse200366Items) GetControllerOk() (*InlineResponse200366Controller, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *InlineResponse200366Items) SetController(v InlineResponse200366Controller)`

SetController sets Controller field to given value.

### HasController

`func (o *InlineResponse200366Items) HasController() bool`

HasController returns a boolean if a field has been set.

### GetJoinedAt

`func (o *InlineResponse200366Items) GetJoinedAt() string`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *InlineResponse200366Items) GetJoinedAtOk() (*string, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *InlineResponse200366Items) SetJoinedAt(v string)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *InlineResponse200366Items) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200366Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200366Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200366Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200366Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200366Items) GetTags() []InlineResponse200366Tags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200366Items) GetTagsOk() (*[]InlineResponse200366Tags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200366Items) SetTags(v []InlineResponse200366Tags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200366Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMode

`func (o *InlineResponse200366Items) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse200366Items) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse200366Items) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse200366Items) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetCountryCode

`func (o *InlineResponse200366Items) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *InlineResponse200366Items) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *InlineResponse200366Items) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *InlineResponse200366Items) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDetails

`func (o *InlineResponse200366Items) GetDetails() []InlineResponse200366Details`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InlineResponse200366Items) GetDetailsOk() (*[]InlineResponse200366Details, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InlineResponse200366Items) SetDetails(v []InlineResponse200366Details)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InlineResponse200366Items) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


