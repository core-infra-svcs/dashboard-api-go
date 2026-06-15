# InlineResponse20036

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MulticastRoutingId** | Pointer to **string** | ID of the Multicast routing request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this URL to check the status of your Multicast routing request. | [optional] 
**Request** | Pointer to [**InlineResponse2015Request**](InlineResponse2015Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the Multicast routing request. | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse20036Interfaces**](InlineResponse20036Interfaces.md) | The interfaces that have PIM enabled | [optional] 
**Routes** | Pointer to [**[]InlineResponse20036Routes**](InlineResponse20036Routes.md) | The multicast routes | [optional] 
**Error** | Pointer to **string** | Description of the error. | [optional] 

## Methods

### NewInlineResponse20036

`func NewInlineResponse20036() *InlineResponse20036`

NewInlineResponse20036 instantiates a new InlineResponse20036 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20036WithDefaults

`func NewInlineResponse20036WithDefaults() *InlineResponse20036`

NewInlineResponse20036WithDefaults instantiates a new InlineResponse20036 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMulticastRoutingId

`func (o *InlineResponse20036) GetMulticastRoutingId() string`

GetMulticastRoutingId returns the MulticastRoutingId field if non-nil, zero value otherwise.

### GetMulticastRoutingIdOk

`func (o *InlineResponse20036) GetMulticastRoutingIdOk() (*string, bool)`

GetMulticastRoutingIdOk returns a tuple with the MulticastRoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRoutingId

`func (o *InlineResponse20036) SetMulticastRoutingId(v string)`

SetMulticastRoutingId sets MulticastRoutingId field to given value.

### HasMulticastRoutingId

`func (o *InlineResponse20036) HasMulticastRoutingId() bool`

HasMulticastRoutingId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20036) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20036) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20036) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20036) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20036) GetRequest() InlineResponse2015Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20036) GetRequestOk() (*InlineResponse2015Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20036) SetRequest(v InlineResponse2015Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20036) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20036) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20036) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20036) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20036) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse20036) GetInterfaces() []InlineResponse20036Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse20036) GetInterfacesOk() (*[]InlineResponse20036Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse20036) SetInterfaces(v []InlineResponse20036Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse20036) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetRoutes

`func (o *InlineResponse20036) GetRoutes() []InlineResponse20036Routes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *InlineResponse20036) GetRoutesOk() (*[]InlineResponse20036Routes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *InlineResponse20036) SetRoutes(v []InlineResponse20036Routes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *InlineResponse20036) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20036) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20036) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20036) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20036) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


