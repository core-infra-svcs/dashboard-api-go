# InlineResponse20037

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MulticastRoutingId** | Pointer to **string** | ID of the Multicast routing request. Used to check the status of the request. | [optional] 
**Url** | Pointer to **string** | GET this URL to check the status of your Multicast routing request. | [optional] 
**Request** | Pointer to [**InlineResponse2015Request**](InlineResponse2015Request.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the Multicast routing request. | [optional] 
**Interfaces** | Pointer to [**[]InlineResponse20037Interfaces**](InlineResponse20037Interfaces.md) | The interfaces that have PIM enabled | [optional] 
**Routes** | Pointer to [**[]InlineResponse20037Routes**](InlineResponse20037Routes.md) | The multicast routes | [optional] 
**Error** | Pointer to **string** | Description of the error. | [optional] 

## Methods

### NewInlineResponse20037

`func NewInlineResponse20037() *InlineResponse20037`

NewInlineResponse20037 instantiates a new InlineResponse20037 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20037WithDefaults

`func NewInlineResponse20037WithDefaults() *InlineResponse20037`

NewInlineResponse20037WithDefaults instantiates a new InlineResponse20037 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMulticastRoutingId

`func (o *InlineResponse20037) GetMulticastRoutingId() string`

GetMulticastRoutingId returns the MulticastRoutingId field if non-nil, zero value otherwise.

### GetMulticastRoutingIdOk

`func (o *InlineResponse20037) GetMulticastRoutingIdOk() (*string, bool)`

GetMulticastRoutingIdOk returns a tuple with the MulticastRoutingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastRoutingId

`func (o *InlineResponse20037) SetMulticastRoutingId(v string)`

SetMulticastRoutingId sets MulticastRoutingId field to given value.

### HasMulticastRoutingId

`func (o *InlineResponse20037) HasMulticastRoutingId() bool`

HasMulticastRoutingId returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20037) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20037) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20037) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20037) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetRequest

`func (o *InlineResponse20037) GetRequest() InlineResponse2015Request`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *InlineResponse20037) GetRequestOk() (*InlineResponse2015Request, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *InlineResponse20037) SetRequest(v InlineResponse2015Request)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *InlineResponse20037) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20037) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20037) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20037) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20037) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetInterfaces

`func (o *InlineResponse20037) GetInterfaces() []InlineResponse20037Interfaces`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InlineResponse20037) GetInterfacesOk() (*[]InlineResponse20037Interfaces, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InlineResponse20037) SetInterfaces(v []InlineResponse20037Interfaces)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InlineResponse20037) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetRoutes

`func (o *InlineResponse20037) GetRoutes() []InlineResponse20037Routes`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *InlineResponse20037) GetRoutesOk() (*[]InlineResponse20037Routes, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *InlineResponse20037) SetRoutes(v []InlineResponse20037Routes)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *InlineResponse20037) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetError

`func (o *InlineResponse20037) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InlineResponse20037) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InlineResponse20037) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InlineResponse20037) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


