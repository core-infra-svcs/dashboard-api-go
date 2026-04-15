# InlineResponse200366Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**History** | Pointer to [**[]InlineResponse200366History**](InlineResponse200366History.md) | Site level connectivity history | [optional] 
**Devices** | Pointer to [**InlineResponse200366Devices**](InlineResponse200366Devices.md) |  | [optional] 

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

### GetSiteId

`func (o *InlineResponse200366Items) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200366Items) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200366Items) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200366Items) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200366Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200366Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200366Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200366Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHistory

`func (o *InlineResponse200366Items) GetHistory() []InlineResponse200366History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *InlineResponse200366Items) GetHistoryOk() (*[]InlineResponse200366History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *InlineResponse200366Items) SetHistory(v []InlineResponse200366History)`

SetHistory sets History field to given value.

### HasHistory

`func (o *InlineResponse200366Items) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200366Items) GetDevices() InlineResponse200366Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200366Items) GetDevicesOk() (*InlineResponse200366Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200366Items) SetDevices(v InlineResponse200366Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200366Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


