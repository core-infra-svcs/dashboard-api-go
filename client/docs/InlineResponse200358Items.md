# InlineResponse200358Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Serial** | Pointer to **string** | Unique serial number for device. | [optional] 
**Name** | Pointer to **string** | Name of device. | [optional] 
**Network** | Pointer to [**InlineResponse200358Network**](InlineResponse200358Network.md) |  | [optional] 
**BasicServiceSets** | Pointer to [**[]InlineResponse200358BasicServiceSets**](InlineResponse200358BasicServiceSets.md) | Status information for wireless access points. | [optional] 

## Methods

### NewInlineResponse200358Items

`func NewInlineResponse200358Items() *InlineResponse200358Items`

NewInlineResponse200358Items instantiates a new InlineResponse200358Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200358ItemsWithDefaults

`func NewInlineResponse200358ItemsWithDefaults() *InlineResponse200358Items`

NewInlineResponse200358ItemsWithDefaults instantiates a new InlineResponse200358Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerial

`func (o *InlineResponse200358Items) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InlineResponse200358Items) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InlineResponse200358Items) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InlineResponse200358Items) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200358Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200358Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200358Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200358Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *InlineResponse200358Items) GetNetwork() InlineResponse200358Network`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *InlineResponse200358Items) GetNetworkOk() (*InlineResponse200358Network, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *InlineResponse200358Items) SetNetwork(v InlineResponse200358Network)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *InlineResponse200358Items) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetBasicServiceSets

`func (o *InlineResponse200358Items) GetBasicServiceSets() []InlineResponse200358BasicServiceSets`

GetBasicServiceSets returns the BasicServiceSets field if non-nil, zero value otherwise.

### GetBasicServiceSetsOk

`func (o *InlineResponse200358Items) GetBasicServiceSetsOk() (*[]InlineResponse200358BasicServiceSets, bool)`

GetBasicServiceSetsOk returns a tuple with the BasicServiceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicServiceSets

`func (o *InlineResponse200358Items) SetBasicServiceSets(v []InlineResponse200358BasicServiceSets)`

SetBasicServiceSets sets BasicServiceSets field to given value.

### HasBasicServiceSets

`func (o *InlineResponse200358Items) HasBasicServiceSets() bool`

HasBasicServiceSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


