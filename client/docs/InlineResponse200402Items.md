# InlineResponse200402Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**History** | Pointer to [**[]InlineResponse200402History**](InlineResponse200402History.md) | Site level connectivity history | [optional] 
**Devices** | Pointer to [**InlineResponse200402Devices**](InlineResponse200402Devices.md) |  | [optional] 

## Methods

### NewInlineResponse200402Items

`func NewInlineResponse200402Items() *InlineResponse200402Items`

NewInlineResponse200402Items instantiates a new InlineResponse200402Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200402ItemsWithDefaults

`func NewInlineResponse200402ItemsWithDefaults() *InlineResponse200402Items`

NewInlineResponse200402ItemsWithDefaults instantiates a new InlineResponse200402Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *InlineResponse200402Items) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200402Items) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200402Items) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200402Items) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200402Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200402Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200402Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200402Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHistory

`func (o *InlineResponse200402Items) GetHistory() []InlineResponse200402History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *InlineResponse200402Items) GetHistoryOk() (*[]InlineResponse200402History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *InlineResponse200402Items) SetHistory(v []InlineResponse200402History)`

SetHistory sets History field to given value.

### HasHistory

`func (o *InlineResponse200402Items) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200402Items) GetDevices() InlineResponse200402Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200402Items) GetDevicesOk() (*InlineResponse200402Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200402Items) SetDevices(v InlineResponse200402Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200402Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


