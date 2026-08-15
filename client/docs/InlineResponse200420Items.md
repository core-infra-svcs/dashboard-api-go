# InlineResponse200420Items

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Name** | Pointer to **string** | Site name | [optional] 
**History** | Pointer to [**[]InlineResponse200420History**](InlineResponse200420History.md) | Site level connectivity history | [optional] 
**Devices** | Pointer to [**InlineResponse200420Devices**](InlineResponse200420Devices.md) |  | [optional] 

## Methods

### NewInlineResponse200420Items

`func NewInlineResponse200420Items() *InlineResponse200420Items`

NewInlineResponse200420Items instantiates a new InlineResponse200420Items object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200420ItemsWithDefaults

`func NewInlineResponse200420ItemsWithDefaults() *InlineResponse200420Items`

NewInlineResponse200420ItemsWithDefaults instantiates a new InlineResponse200420Items object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *InlineResponse200420Items) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InlineResponse200420Items) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InlineResponse200420Items) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InlineResponse200420Items) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse200420Items) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse200420Items) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse200420Items) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse200420Items) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHistory

`func (o *InlineResponse200420Items) GetHistory() []InlineResponse200420History`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *InlineResponse200420Items) GetHistoryOk() (*[]InlineResponse200420History, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *InlineResponse200420Items) SetHistory(v []InlineResponse200420History)`

SetHistory sets History field to given value.

### HasHistory

`func (o *InlineResponse200420Items) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetDevices

`func (o *InlineResponse200420Items) GetDevices() InlineResponse200420Devices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *InlineResponse200420Items) GetDevicesOk() (*InlineResponse200420Devices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *InlineResponse200420Items) SetDevices(v InlineResponse200420Devices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *InlineResponse200420Items) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


