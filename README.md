## Overview API

### User API
* [POST   /v1/user/login](#v1userlogin---post)
* [POST   /v1/user/register      ](#v1userregister---post)
* [POST   /v1/user/isTokenValid  ](#v1useristokenvalid---post)
* [PUT    /v1/user/updateProfile ](#v1userupdateprofile---put)
* [GET    /v1/user/retrieveUserProfile](#v1userretrieveuserprofile---get)


### Kid API
* [POST   /v1/kids/add             ](#v1kidsadd---post)
* [PUT    /v1/kids/update          ](#v1kidsupdate---put)
* [GET    /v1/kids/whoRegisteredMacID          ](#v1kidswhoregisteredmacid---get)

### Avatar
* [POST   /v1/user/avatar/upload   ](#v1useravatarupload---post)
* [POST   /v1/user/avatar/uploadKid](#v1useravataruploadkid---post)

### Activity
* [POST   /v1/activity/uploadRawData](#v1activityuploadrawdata---post)
* [GET   /v1/activity/retrieveData](#v1activityretrievedata---get)

### Event
* [POST   /v1/event/add](#v1eventadd---post)
* [PUT   /v1/event/update](#v1eventupdate---put)
* [DELETE   /v1/event/delete](#v1eventdelete---delete)
* [GET   /v1/event/retrieveEvents](#v1eventretrieveevents---get)

### Multi-Host API
* [POST   /v1/subHost/add](#v1subhostadd---post)

## /v1/user/login - POST
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| :------------- |:-------------:|:-------------:| :-----|
| email     | Yes | String |   test@kidsdynami.com |
| password     | Yes | String |   aaaaaa |

#### Response Status
| Status Code    | Meaning      |
| :------------- |:-------------|
| 200     | Success Login |
| 400     | Bad Request. Login failed |

* Success response body:
```
{
  "access_token": "2f69e1d858a04b4f35dd477c76921b34",
  "username": "teste"
}
```

## /v1/user/isTokenValid - POST
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@kidsdynami.com |
| token     | Yes | String |   2f69e1d858a04b4f35dd477c76921b34 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | The token is valid |
| 403     | Forbidden. The token is not valid |
* The API doesn't return any JSON

## /v1/user/register - POST
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@kidsdynami.com |
| password     | Yes | String |   aaaaaa |
| firstName     | Yes | String |   Jay |
| lastName     | Yes | String |   Chen |
| phoneNumber     | No | String |   3442314231 |
| zipCode     | No | String |   11101 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Register success |
| 400     | Bad request. Missing some parameters |
| 409     | Conflict. The email is already registered |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it doesn't return JSON
* Fail - response body:
```
{
  "message": "The email is already registered"
}
```
* Internal error - response body:
```
{
  "message": "Error when insert data",
  "error": "Test error"
}
```

## /v1/user/updateProfile - PUT
* Content-Type: application/json
* The API doesn't support email and password update

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| firstName    | No | String |   Jay |
| lastName     | No | String |   Chen |
| phoneNumber  | No | String |   3442314231 |
| zipCode      | No | String |   11101 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |
* Success - 
```
{
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "firstName": "KIDLLE",
    "lastName": "YES",
    "lastUpdate": "2016-12-18T21:24:57Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "11111",
    "phoneNumber": "3444943214",
    "profile": ""
  }
}
```
* Internal error - response body:
```
{
  "message": "Error when insert data",
  "error": "Test error"
}
```

## /v1/user/retrieveUserProfile - GET
* Content-Type: application/json
* Retrieve user data by the header token
* No any parameter required, just need x-auth-token on the header
* It retrieves with kids data that belong to the user

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 400     | Bad request. The token is invalid |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '' "http://localhost:8111/v1/user/retrieveUserProfile"
```
* Success - with kids data
```
{
  "kids": [
    {
      "id": 18,
      "firstName": "Jay",
      "lastName": "Chen",
      "dateCreated": "2016-12-11T22:37:15Z",
      "macId": "13031FCFE5E02",
      "profile": ""
    },
    {
      "id": 19,
      "firstName": "KIDLLE",
      "lastName": "YES",
      "dateCreated": "2016-12-18T04:17:35Z",
      "macId": "hgweorahgbkljwhnpi",
      "profile": ""
    },
    {
      "id": 20,
      "firstName": "KIDLLE",
      "lastName": "YES",
      "dateCreated": "2016-12-18T21:19:54Z",
      "macId": "hgweorahgbkljwhnpi2",
      "profile": ""
    }
  ],
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "firstName": "KIDLLE",
    "lastName": "YES",
    "lastUpdate": "2016-12-18T21:24:57Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  }
}
```

## /v1/user/avatar/upload - POST
* Upload user profile image
* Content-Type: multipart/form-data

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| upload     | Yes | file |    |

### curl Example
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637"  -H "Content-Type: multipart/form-data; -F "upload=@CwEiSDRVMAE4IO4.jpg" "http://localhost:8111/v1/user/avatar/upload"
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | File upload successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it updates profile property
```
{
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "firstName": "KIDLLE",
    "lastName": "YES",
    "lastUpdate": "2016-12-18T04:11:02Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": "avatar_29.jpg"
  }
}
```
* Bad request error - response body:
```
{
  "error": {
    "ErrorString": "request Content-Type isn't multipart/form-data"
  },
  "message": "upload parameter is required"
}
```

## /v1/user/avatar/uploadKid - POST
* Upload kid profile image
* Content-Type: multipart/form-data

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| upload     | Yes | File |    |
| kidId      | Yes | Integer | |

### curl Example
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -H "Content-Type: multipart/form-data; -F "kidId=19" -F "upload=@CwEiSDRVMAE4IO4.jpg" "http://localhost:8111/v1/user/avatar/uploadKid"
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | File upload successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it updates profile property
```
{
  "kid": {
    "id": 19,
    "firstName": "KIDLLE",
    "lastName": "YES",
    "dateCreated": "2016-12-18T04:17:35Z",
    "macId": "",
    "profile": "kid_avatar_19.jpg"
  }
}
```
* Bad request error - response body:
```
{
  "error": {
    "ErrorString": "request Content-Type isn't multipart/form-data"
  },
  "message": "upload parameter is required"
}
```

## /v1/kids/add - POST
* Add kid
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| firstName     | Yes | string | Kiddle   |
| lastName      | Yes | string | JY |
| macId      | Yes | string | 8D071FCFE5E0 |

### curl Example
```
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '{
	"firstName": "Kiddle",
	"lastName": "JY",
	"macId": "8D071FCFE5E0"
}' "http://localhost:8111/v1/kids/add"
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Data added successfully |
| 400     | Bad request. Missing some parameters |
| 409     | Conflict. The mac ID is already registered |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - it returns all of kids belong to the user
```
{
  "kids": [
    {
      "id": 18,
      "firstName": "Jay",
      "lastName": "Chen",
      "dateCreated": "2016-12-11T22:37:15Z",
      "macId": "13031FCFE5E02",
      "profile": ""
    },
    {
      "id": 19,
      "firstName": "KIDLLE",
      "lastName": "YES",
      "dateCreated": "2016-12-18T04:17:35Z",
      "macId": "hgweorahgbkljwhnpi",
      "profile": ""
    },
    {
      "id": 20,
      "firstName": "KIDLLE",
      "lastName": "YES",
      "dateCreated": "2016-12-18T21:19:54Z",
      "macId": "hgweorahgbkljwhnpi2",
      "profile": ""
    }
  ]
}
```
* Conflict error - response body:
```
{
  "message": "The device is already registered"
}
```

## /v1/kids/update - PUT
* Content-Type: application/json
* Update kid info

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| kidId    | Yes | String |   19 |
| firstName    | No | String |   Jay |
| lastName     | No | String |   Chen |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -H "Content-Type: application/json" -d '{
	"lastName": "yes",
	"kidId": 19
}' "http://localhost:8111/v1/kids/update"
```

* Success - 
```
{
  "kid": {
    "id": 19,
    "firstName": "KIDLLE",
    "lastName": "yes",
    "dateCreated": "2016-12-18T04:17:35Z",
    "macId": "",
    "profile": ""
  }
}
```
* Internal error - response body:
```
{
  "message": "Error when insert data",
  "error": "Test error"
}
```

## /v1/kids/whoRegisteredMacID - GET
* Retrieve kid and user information by MAC ID

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| macId    | Yes | String |   hgweorahgbkljwhnpi2 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 404     | The user not found which means the MAC ID is not registered by anyone |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" "http://localhost:8111/v1/kids/whoRegisteredMacID?macId=hgweorahgbkljwhnpi2"
```

* Success - 
```
{
  "kid": {
    "id": 20,
    "firstName": "KIDLLE",
    "lastName": "YES",
    "dateCreated": "2016-12-21T03:24:57Z",
    "macId": "",
    "profile": "",
    "ParentID": 29
  },
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "firstName": "q",
    "lastName": "w",
    "lastUpdate": "2016-12-19T22:28:07Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "",
    "phoneNumber": "555",
    "profile": ""
  }
}
```

## /v1/activity/uploadRawData - POST
* Content-Type: application/json
* Activity data
* When the data duplicate which is when server response ***409*** status, you can ***ignore*** it and process next data

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| indoorActivity      | Yes | String | 1481299119,0,216,2,3,4 |
| outdoorActivity     | Yes | String | 1481299119,1,0,0,0,0 |
| time                | Yes | Long |   1470885849 |
| macId               | Yes | String |   hgweorahgbkljwhnpi2 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 409     | Conflict. The data is already exist |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -H "Content-Type: application/json" -d '{
	"indoorActivity": "1481299119,0,216,2,3,4",
	"outdoorActivity": "1481299119,1,0,0,0,0",
	"time": 1470885849,
	"macId": "hgweorahgbkljwhnpi2"
}' "http://localhost:8111/v1/activity/uploadRawData"
```

* Success - No any JSON response

* 409 status - duplicate data:
```
{
  "message": "This is a duplicate data"
}
```


## /v1/activity/retrieveData - GET
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| kidId      | Yes | Integer | 20 |
| period     | Yes | String | DAILY, MONTHLY, YEARLY  |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receiving Data successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" "http://localhost:8111/v1/activity/dailyData?kidId=20&period=DAILY"
```

* Success - 
```
{
  "activities": [
    {
      "id": 121,
      "macId": "hgweorahgbkljwhnpi2",
      "kidId": "20",
      "type": "INDOOR",
      "steps": 3298,
      "distance": 0,
      "receivedDate": "2016-12-13T19:55:02Z"
    },
    {
      "id": 122,
      "macId": "hgweorahgbkljwhnpi2",
      "kidId": "20",
      "type": "OUTDOOR",
      "steps": 42,
      "distance": 0,
      "receivedDate": "2016-12-13T19:55:02Z"
    }
  ]
}
```

* 400 Bad Request - 
```
{
  "error": {
    "Func": "ParseInt",
    "Num": "",
    "Err": {}
  },
  "message": "kidId should be int type."
}
```

## /v1/event/add - POST
* Content-Type: application/json
* Date Time format is ***YYYY-MM-ddThh:mm:ss***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| kidId             | Yes | Integer | 20 |
| name               | Yes | String | Test event name |
| startDate           | Yes | String |   2015-08-30T08:20:00 |
| endDate             | Yes | String |   2015-08-31T08:20:00 |
| color               | Yes | String |   #F05D25 |
| description         | No | String |   Test description |
| alert               | No | Integer |   49 |
| city               | No | String |   New York |
| state               | No | String |   New York |
| repeat               | No | String |   DAILY, MONTHLY |
| timezoneOffset       Yes | Integer |   300 |
| todo               | No | Array |   ["test todo 1", "test todo 2"] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Added successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '{
  "kidId": 20,
  "Name": "Test event name",
  "startDate": "2015-08-30T08:20:00",
  "endDate": "2015-08-31T08:20:00",
  "timezoneOffset": 300,
  "color": "#F05D25",
  "description": "Hahah",
  "alert": 49,
  "city": "New York",
  "state": "New York",
  "todo": [
  	"test todo 1", "test todo 2"	
  ]
  
}' "http://localhost:8111/v1/event/add"
```

* Success - Returns added event
```
{
  "event": {
    "id": 398,
    "userId": 29,
    "kidId": 20,
    "name": "Test event name",
    "startDate": "2015-08-30T08:20:00Z",
    "endDate": "2015-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2016-12-26T06:25:19Z",
    "lastUpdated": "2016-12-26T06:25:19Z",
    "todo": [
      {
        "id": 38,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2016-12-26T06:25:19Z",
        "lastUpdated": "2016-12-26T06:25:19Z"
      },
      {
        "id": 39,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2016-12-26T06:25:19Z",
        "lastUpdated": "2016-12-26T06:25:19Z"
      }
    ]
  }
}
```

## /v1/event/update - PUT
* Content-Type: application/json
* Date Time format is ***YYYY-MM-ddThh:mm:ss***
* Send the parameter even user does not change it

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| eventId             | Yes | Integer | 414 |
| name               | Yes | String | Test event name |
| startDate           | Yes | String |   2015-08-30T08:20:00 |
| endDate             | Yes | String |   2015-08-31T08:20:00 |
| color               | Yes | String |   #F05D25 |
| description         | No | String |   Test description |
| alert               | No | Integer |   49 |
| city               | No | String |   New York |
| state               | No | String |   New York |
| repeat               | No | String |   DAILY, MONTHLY |
| timezoneOffset       Yes | Integer |   300 |
| todo               | No | Array |   ["test todo 1"] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Updated successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '{
  "eventId": 414,
  "Name": "Test event name2",
  "startDate": "2015-08-30T08:20:00",
  "endDate": "2015-08-31T08:20:00",
  "timezoneOffset": 300,
  "color": "#F05D25",
  "description": "Hahah",
  "alert": 49,
  "city": "New York",
  "state": "New York",
  "todo": [
  	"test todo 1"	
  ]
  
}' "http://localhost:8111/v1/event/update"
```

* Success - Returns updated event
```
{
  "event": {
    "id": 414,
    "userId": 29,
    "kidId": 20,
    "name": "Test event name2",
    "startDate": "2015-08-30T08:20:00Z",
    "endDate": "2015-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hahah",
    "alert": 49,
    "city": "New York",
    "state": "New York",
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2016-12-26T21:21:53Z",
    "lastUpdated": "2016-12-27T00:59:28Z",
    "todo": [
      {
        "id": 83,
        "text": "test todo 1",
        "status": "PENDING",
        "dateCreated": "2016-12-27T00:59:28Z",
        "lastUpdated": "2016-12-27T00:59:28Z"
      }
    ]
  }
}
```

## /v1/event/delete - DELETE
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| eventId             | Yes | Integer | 414 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Delete successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X DELETE -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '{
  "eventId": 414
}' "http://localhost:8111/v1/event/delete"
```

* Success - Returns empty json
```
{}
```

## /v1/event/retrieveEvents - GET
* Date Time format is ***YYYY-MM-ddThh:mm:ss***
* If trying to retrieve month events, do not use '00' as month value. Example: 2016-12-01T00:00:00 to retrieve 2016-12 events.


#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| period        | Yes | String | MONTH, DAY |
| date          | Yes | String | 2016-12-26T01:00:00 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" "http://localhost:8111/v1/event/retrieveEvents?period=MONTH&date=2016-12-01T00:00:00"
```

* Success - Returns updated event
```
{
  "events": [
    {
      "id": 410,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-27T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "New York",
      "state": "New York",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:17:30Z",
      "lastUpdated": "2016-12-26T21:17:30Z",
      "todo": [
        {
          "id": 66,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:30Z",
          "lastUpdated": "2016-12-26T21:17:30Z"
        },
        {
          "id": 67,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:30Z",
          "lastUpdated": "2016-12-26T21:17:30Z"
        }
      ]
    },
    {
      "id": 411,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-26T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "New York",
      "state": "New York",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:17:35Z",
      "lastUpdated": "2016-12-26T21:17:35Z",
      "todo": [
        {
          "id": 68,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:35Z",
          "lastUpdated": "2016-12-26T21:17:35Z"
        },
        {
          "id": 69,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:35Z",
          "lastUpdated": "2016-12-26T21:17:35Z"
        }
      ]
    },
    {
      "id": 412,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-26T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "",
      "state": "",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:17:45Z",
      "lastUpdated": "2016-12-26T21:17:45Z",
      "todo": [
        {
          "id": 70,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:45Z",
          "lastUpdated": "2016-12-26T21:17:45Z"
        },
        {
          "id": 71,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:17:45Z",
          "lastUpdated": "2016-12-26T21:17:45Z"
        }
      ]
    },
    {
      "id": 413,
      "userId": 29,
      "kidId": 20,
      "name": "Test event name",
      "startDate": "2016-12-26T21:21:53Z",
      "endDate": "2015-08-31T08:20:00Z",
      "color": "#F05D25",
      "status": "OPEN",
      "description": "Hahah",
      "alert": 49,
      "city": "New York",
      "state": "New York",
      "repeat": "",
      "timezoneOffset": 300,
      "dateCreated": "2016-12-26T21:19:16Z",
      "lastUpdated": "2016-12-26T21:19:16Z",
      "todo": [
        {
          "id": 72,
          "text": "test todo 1",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:19:16Z",
          "lastUpdated": "2016-12-26T21:19:16Z"
        },
        {
          "id": 73,
          "text": "test todo 2",
          "status": "PENDING",
          "dateCreated": "2016-12-26T21:19:16Z",
          "lastUpdated": "2016-12-26T21:19:16Z"
        }
      ]
    }
  ]
}
```

## /v1/subHost/add - POST
* Content-Type: application/json
* Get ```hostId``` from [/v1/kids/whoRegisteredMacID](#v1kidswhoregisteredmacid---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| macId             | Yes | String | 13031FCFE5E0 |
| hostId            | Yes | Integer | 35 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Added successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 409     | Conflict. The request is already exists |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -d '{
  "kidId": 20,
  "Name": "Test event name",
  "startDate": "2015-08-30T08:20:00",
  "endDate": "2015-08-31T08:20:00",
  "timezoneOffset": 300,
  "color": "#F05D25",
  "description": "Hahah",
  "alert": 49,
  "city": "New York",
  "state": "New York",
  "todo": [
  	"test todo 1", "test todo 2"	
  ]
  
}' "http://localhost:8111/v1/event/add"
```

* Success - Returns added request
```
{
  "SubHostRequest": {
    "id": 10,
    "macId": "13031FCFE5E0",
    "requestFromID": 29,
    "requestToID": 35,
    "status": "PENDING",
    "createdDate": "2017-01-01T03:00:06Z",
    "lastUpdated": "2017-01-01T03:00:06Z"
  }
}
```

* Fail (409) -
```
{
  "message": "The request is already exist"
}
```