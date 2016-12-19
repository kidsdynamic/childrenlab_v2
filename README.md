## Overview API
* [POST   /v1/user/login](#v1userlogin--post)
* [POST   /v1/user/register      ](#v1userregister---post)
* [POST   /v1/user/isTokenValid  ](#v1useristokenvalid---post)
* [PUT    /v1/user/updateProfile ](#v1userupdateprofile---put)
* [GET    /v1/user/retrieveUserProfile](#v1userretrieveuserprofile---get)
* [POST   /v1/kids/add             ](#v1kidsadd---post)
* [PUT    /v1/kids/update          ](#v1kidsupdate---put)
* [POST   /v1/user/avatar/upload   ](#v1useravatarupload---post)
* [POST   /v1/user/avatar/uploadKid](#v1useravataruploadkid---post)
* [POST   /v1/activity/uploadRawData](#v1activityuploadrawdata---post)



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


