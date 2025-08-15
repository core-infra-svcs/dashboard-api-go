# InlineResponse200372Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for device. | [optional] 
**Name** | Pointer to **string** | Name of device. | [optional] 
**Network** | Pointer to [**InlineResponse200372Network**](InlineResponse200372Network.md) |  | [optional] 
**BasicServiceSets** | Pointer to [**[]InlineResponse200372BasicServiceSets**](InlineResponse200372BasicServiceSets.md) | Status information for wireless access points. | [optional] 

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

### GetBasicServiceSets

`func (o *InlineResponse200372Items) GetBasicServiceSets() []InlineResponse200372BasicServiceSets`

GetBasicServiceSets returns the BasicServiceSets field if non-nil, zero value otherwise.

### GetBasicServiceSetsOk

`func (o *InlineResponse200372Items) GetBasicServiceSetsOk() (*[]InlineResponse200372BasicServiceSets, bool)`

GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicServiceSets

`func (o *InlineResponse200372Items) SetBasicServiceSets(v []InlineResponse200372BasicServiceSets)`

SetBasicServiceSets sets BasicServiceSets field to given value.

### HasBasicServiceSets

`func (o *InlineResponse200372Items) HasBasicServiceSets() bool`

HasBasicServiceSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


