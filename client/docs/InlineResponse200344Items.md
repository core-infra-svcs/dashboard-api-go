# InlineResponse200344Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device | [optional] 
**Network** | Pointer to [**InlineResponse200270Network**](InlineResponse200270Network.md) |  | [optional] 
**Events** | Pointer to [**[]InlineResponse200344Events**](InlineResponse200344Events.md) | Events indicating power mode changes for the device | [optional] 

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

### GetSerial

`func (o *InlineResponse200344Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200344Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200344Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200344Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200344Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200344Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200344Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200344Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200344Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200344Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200344Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200344Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200344Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200344Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200344Items) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200344Items) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200344Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200344Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200344Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200344Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200344Items) GetNetwork() InlineResponse200270Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200344Items) GetNetworkOk() (*InlineResponse200270Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200344Items) SetNetwork(v InlineResponse200270Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200344Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetEvents

`func (o *InlineResponse200344Items) GetEvents() []InlineResponse200344Events`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *InlineResponse200344Items) GetEventsOk() (*[]InlineResponse200344Events, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *InlineResponse200344Items) SetEvents(v []InlineResponse200344Events)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *InlineResponse200344Items) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


