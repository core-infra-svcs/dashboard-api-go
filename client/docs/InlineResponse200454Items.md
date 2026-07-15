# InlineResponse200454Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for the device | [optional] 
**Model** | Pointer to **string** | Model of the device | [optional] 
**Name** | Pointer to **string** | Name of the device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Tags** | Pointer to **[]string** | List of custom tags for the device | [optional] 
**Network** | Pointer to [**InlineResponse200355Network**](InlineResponse200355Network.md) |  | [optional] 
**Events** | Pointer to [**[]InlineResponse200454Events**](InlineResponse200454Events.md) | Events indicating power mode changes for the device | [optional] 

## Methods

### NewInlineResponse200454Items

`func NewInlineResponse200454Items() *InlineResponse200454Items`

NewInlineResponse200454Items instantiates a new InlineResponse200454Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200454ItemsWithDefaults

`func NewInlineResponse200454ItemsWithDefaults() *InlineResponse200454Items`

NewInlineResponse200454ItemsWithDefaults instantiates a new InlineResponse200454Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200454Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200454Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200454Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200454Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetModel

`func (o *InlineResponse200454Items) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InlineResponse200454Items) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InlineResponse200454Items) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InlineResponse200454Items) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200454Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200454Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200454Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200454Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMac

`func (o *InlineResponse200454Items) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *InlineResponse200454Items) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *InlineResponse200454Items) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *InlineResponse200454Items) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse200454Items) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse200454Items) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse200454Items) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse200454Items) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200454Items) GetNetwork() InlineResponse200355Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200454Items) GetNetworkOk() (*InlineResponse200355Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200454Items) SetNetwork(v InlineResponse200355Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200454Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetEvents

`func (o *InlineResponse200454Items) GetEvents() []InlineResponse200454Events`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *InlineResponse200454Items) GetEventsOk() (*[]InlineResponse200454Events, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *InlineResponse200454Items) SetEvents(v []InlineResponse200454Events)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *InlineResponse200454Items) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


