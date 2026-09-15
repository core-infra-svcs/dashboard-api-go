# InlineResponse200355Geolocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether geolocation is enabled for the device | [optional] 
**Latitude** | Pointer to **NullableFloat32** | Latitude of the device&#39;s location in decimal degrees. Range: -90 to 90. Null when geolocation is disabled or unavailable. | [optional] 
**Longitude** | Pointer to **NullableFloat32** | Longitude of the device&#39;s location in decimal degrees. Range: -180 to 180. Null when geolocation is disabled or unavailable. | [optional] 
**LastReportedAt** | Pointer to **NullableTime** | UTC ISO 8601 timestamp when the device last reported geolocation telemetry. Null when geolocation is disabled or no location has been reported. | [optional] 

## Methods

### NewInlineResponse200355Geolocation

`func NewInlineResponse200355Geolocation() *InlineResponse200355Geolocation`

NewInlineResponse200355Geolocation instantiates a new InlineResponse200355Geolocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200355GeolocationWithDefaults

`func NewInlineResponse200355GeolocationWithDefaults() *InlineResponse200355Geolocation`

NewInlineResponse200355GeolocationWithDefaults instantiates a new InlineResponse200355Geolocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200355Geolocation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200355Geolocation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200355Geolocation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200355Geolocation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLatitude

`func (o *InlineResponse200355Geolocation) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *InlineResponse200355Geolocation) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *InlineResponse200355Geolocation) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *InlineResponse200355Geolocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *InlineResponse200355Geolocation) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *InlineResponse200355Geolocation) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *InlineResponse200355Geolocation) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *InlineResponse200355Geolocation) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *InlineResponse200355Geolocation) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *InlineResponse200355Geolocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *InlineResponse200355Geolocation) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *InlineResponse200355Geolocation) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetLastReportedAt

`func (o *InlineResponse200355Geolocation) GetLastReportedAt() time.Time`

GetLastReportedAt returns the LastReportedAt field if non-nil, zero value otherwise.

### GetLastReportedAtOk

`func (o *InlineResponse200355Geolocation) GetLastReportedAtOk() (*time.Time, bool)`

GetLastReportedAtOk returns a tuple with the LastReportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportedAt

`func (o *InlineResponse200355Geolocation) SetLastReportedAt(v time.Time)`

SetLastReportedAt sets LastReportedAt field to given value.

### HasLastReportedAt

`func (o *InlineResponse200355Geolocation) HasLastReportedAt() bool`

HasLastReportedAt returns a boolean if a field has been set.

### SetLastReportedAtNil

`func (o *InlineResponse200355Geolocation) SetLastReportedAtNil(b bool)`

 SetLastReportedAtNil sets the value for LastReportedAt to be an explicit nil

### UnsetLastReportedAt
`func (o *InlineResponse200355Geolocation) UnsetLastReportedAt()`

UnsetLastReportedAt ensures that no value is present for LastReportedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


