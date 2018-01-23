## Overview API

### User API
* [POST   /v1/user/login](#v1userlogin---post)
* [POST   /v1/user/register      ](#v1userregister---post)
* [GET   /v1/user/isTokenValid  ](#v1useristokenvalid---get)
* [GET   /v1/user/isEmailAvailableToRegister  ](#v1userisemailavailabletoregister---get)
* [PUT    /v1/user/updateProfile ](#v1userupdateprofile---put)
* [GET    /v1/user/retrieveUserProfile](#v1userretrieveuserprofile---get)
* [PUT    /v1/user/updateIOSRegistrationId](#v1userupdateiosregistrationid---put)
* [PUT    /v1/user/updateAndroidRegistrationId](#v1userupdateandroidregistrationid---put)
* [GET    /v1/user/findByEmail](#v1userfindbyemail---get)
* [POST    /v1/user/updateLanguage](#v1userupdatelanguage---post)
* [POST    /v1/user/sendResetPasswordEmail](#v1usersendresetpasswordemail---post)
* [GET    /v1/user/getUserByEmail](#v1usergetuserbyemail---get)
* [POST    /v1/user/updatePassword](#v1userupdatepassword---post)


### Kid API
* [POST   /v1/kids/add             ](#v1kidsadd---post)
* [PUT    /v1/kids/update          ](#v1kidsupdate---put)
* [GET    /v1/kids/list          ](#v1kidslist---get)
* [GET    /v1/kids/whoRegisteredMacID          ](#v1kidswhoregisteredmacid---get)
* [DELETE    /v1/kids/delete          ](#v1kidsdelete---delete)
* [POST   /v1/kids/batteryStatus             ](#v1kidsbatterystatus---post)
* [PUT   /v1/kids/updateKidRevertMacID             ](#v1kidsupdatekidrevertmacid---put)

### Avatar
* [POST   /v1/user/avatar/upload   ](#v1useravatarupload---post)
* [POST   /v1/user/avatar/uploadKid](#v1useravataruploadkid---post)

### Activity
* [POST   /v1/activity/uploadRawData](#v1activityuploadrawdata---post)
* [GET   /v1/activity/retrieveData](#v1activityretrievedata---get)
* [GET   /v1/activity/retrieveDataByTime](#v1activityretrievedatabytime---get)
* [GET   /v1/activity/retrieveHourlyDataByTime](#v1activityretrievehourlydatabytime---get)
* [GET   /v1/activity/retrieveMonthlyActivity](#v1activityretrievemonthlyactivity---get)

### Event
* [POST   /v1/event/add](#v1eventadd---post)
* [PUT   /v1/event/update](#v1eventupdate---put)
* [DELETE   /v1/event/delete](#v1eventdelete---delete)
* [GET   /v1/event/retrieveEvents](#v1eventretrieveevents---get)
* [GET   /v1/event/retrieveAllEventsWithTodo](#v1eventretrievealleventswithtodo---get)
* [GET   /v1/event/retrieveAllEventsByKid](#v1eventretrievealleventsbykid---get)

#### TODO
* [PUT   /v1/event/todo/done](#v1eventtododone---put)

### Multi-Host API
* [POST   /v1/subHost/add](#v1subhostadd---post)
* [PUT   /v1/subHost/accept](#v1subhostaccept---put)
* [PUT   /v1/subHost/deny](#v1subhostdeny---put)
* [GET   /v1/subHost/list](#v1subhostlist---get)
* [PUT   /v1/subHost/removeKid](#v1subhostremovekid---put)
* [DELETE   /v1/subHost/delete](#v1subhostdelete---delete)

### Firmware version
* [GET /v1/fw/currentVersion](#v1fwcurrentversionmacid---get)
* [PUT   /v1/fw/firmwareVersion](#v1fwfirmwareversion---put)

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

## /v1/user/isTokenValid - GET

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
| language | No | String | en, ja, es, fr |

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

## /v1/user/isEmailAvailableToRegister - GET
* The API doesn't return any body

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@kidsdynami.com |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | The email is able to register |
| 409     | Conflict. The email is already registered |

### curl
```
curl -X GET "http://localhost:8111/v1/user/isEmailAvailableToRegister?email=jack083001@gmail.com"
```



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
      "name": "Jay",
      "dateCreated": "2016-12-11T22:37:15Z",
      "macId": "13031FCFE5E02",
      "profile": ""
    },
    {
      "id": 19,
      "name": "KIDLLE",
      "dateCreated": "2016-12-18T04:17:35Z",
      "macId": "hgweorahgbkljwhnpi",
      "profile": ""
    },
    {
      "id": 20,
      "name": "KIDLLE",
      "dateCreated": "2016-12-18T21:19:54Z",
      "macId": "hgweorahgbkljwhnpi2",
      "profile": ""
    }
  ],
  "user": {
    "id": 29,
    "email": "lwz1@swing.com",
    "name": "KIDLLE",
    "lastUpdate": "2016-12-18T21:24:57Z",
    "dateCreated": "2016-12-06T00:40:10Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  }
}
```

## /v1/user/updateLanguage - POST
* Update user's language

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| language     | Yes | String |   en |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update the registration id successfully |
| 400     | Bad request. The token is invalid |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl --request POST \
  --url http://localhost:8111/v1/user/updateLanguage \
  --header 'content-type: application/json' \
  --header 'x-auth-token: a87de39359d989f4a3a34d732d9b1529' \
  --data '{\n  "language": "es"\n}'
```
* Success - Return user profile
```
{}
```

## /v1/user/sendResetPasswordEmail - POST
* Send reset password email
* If x-auth-token is provided, no need any parameter

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | No | String |   test@gmail.com |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | send successfully |
| 400     | Bad request. The token is invalid |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST \
  http://localhost:8111/v1/user/sendResetPasswordEmail \
  -H 'x-auth-token: 5dd388750f3a17b7e7942aaa5a612692'
```
* Success - Return user profile
```
{}
```

## /v1/user/getUserByEmail - GET
* Get user information by email

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| email     | Yes | String |   test@gmail.com |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | send successfully |
| 400     | Bad request. The token is invalid |
| 404     | The email is not found in the database |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET \
  'http://localhost:8111/v1/user/getUserByEmail?email=ii14%40ii.com' \
  -H 'x-auth-token: 7802035efda3dfe01fa16fd10f97ec81'
```
* Success - Return user profile
```
{
    "id": 324,
    "email": "ii14@ii.com",
    "firstName": "dufhe",
    "lastName": "fudhf",
    "lastUpdate": "2017-10-03T03:59:30Z",
    "dateCreated": "2017-10-03T03:59:30Z",
    "zipCode": "",
    "phoneNumber": "59596855",
    "profile": "",
    "language": "en",
    "ios_registration_id": "",
    "android_registration_id": "dhIKOCE5U-E:APA91bEi19Y428vQCVaC9xsZlcC5Nt2dmc100KXVu91Y0xR8sWh5sKceTzINyjV7qoUVIZlSYueYOZ5PGU1l15ygezRoi8vcGCa1ZTQj0RvLnZWtXx5oEWCgine2sUMv8gOste3tLHRB",
    "country": "US"
}
```
* Error 404
```
{
    "message": "No record"
}
```

## /v1/user/updatePassword - POST
* This API for the user to update the password after signed in
* The password length has to be longer than 6 characters

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| newPassword     | Yes | String |   aaaaaa |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | send successfully |
| 400     | Bad request. The token is invalid or the password length is not longer than 6 characters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST \
  http://localhost:8111/v1/user/updatePassword \
  -H 'content-type: application/json' \
  -H 'x-auth-token: 7802035efda3dfe01fa16fd10f97ec81' \
  -d '{
	"newPassword": "aaaaaa"
	}'
```
* Success - Return 200 status with empty json
```
{
}
```
* Error 400
```
{
    "message": "The password has to be longer than 6 characters"
}
```

## /v1/user/updateIOSRegistrationId - PUT
* It's for IOS registration ID - push notification

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| registrationId     | Yes | String |   TestRegistrationID |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update the registration id successfully |
| 400     | Bad request. The token is invalid |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: ec83d6e41db5168ddb0b1d28b2e262d6" -d '{
  "registrationId": "TestRegistrationID"
}' "http://localhost:8111/v1/user/updateIOSRegistrationId"
```
* Success - Return user profile
```
{
  "id": 6,
  "email": "jack08300@gmail.com",
  "firstName": "JJJ",
  "lastName": "TTT",
  "lastUpdate": "2017-01-11T03:17:31Z",
  "dateCreated": "2017-01-11T03:17:31Z",
  "zipCode": "11111",
  "phoneNumber": "",
  "profile": "",
  "registrationId": "TestRegistrationID"
}
```

## /v1/user/updateAndroidRegistrationId - PUT
* It's for Android registration ID - push notification

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| registrationId     | Yes | String |   TestRegistrationID |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update the registration id successfully |
| 400     | Bad request. The token is invalid |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT \
  http://dev.childrenlab.com:8080/v1/user/updateAndroidRegistrationId \
  -H 'Content-Type: application/json' \
  -H 'x-auth-token: 0619b2d6e3c065bb2a0ced77b3705c8c' \
  -d '{
	"registrationId": "TestRegistrationID"
}'
```
* Success - Return user profile
```
{
    "id": 156,
    "email": "nn@nn.com",
    "firstName": "jdhcbd",
    "lastName": "fhhchdff",
    "dateCreated": "2017-03-17T04:56:16Z",
    "zipCode": "",
    "phoneNumber": "86866598",
    "profile": "avatar_156.jpg",
    "language": "es",
    "country": "US"
}
```

## /v1/user/findByEmail - GET
* Retrieve user data by the email

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Get user successfully |
| 400     | Bad request. The email parameter is missing |
| 404     | Not found. Can't find user by the email |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: 4b0d53ff1953af73aff2fec2478bd1fe" "http://localhost:8111/v1/user/findByEmail?email=jack0830421@gmail.com"
```
* Success -
```
{
  "id": 2,
  "email": "jack08300@gmail.com",
  "firstName": "Jay",
  "lastName": "Chen",
  "lastUpdate": "0001-01-01T00:00:00Z",
  "dateCreated": "0001-01-01T00:00:00Z",
  "zipCode": "",
  "phoneNumber": "",
  "profile": ""
}
```

## /v1/user/avatar/upload - POST
* Upload user profile image
* Content-Type: multipart/form-data
* The file will be on Amazon S3 - ask me the path
* The file name will be /userProfile/avatar_{userID}

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
* The file will be on Amazon S3 - ask me the path
* The file name will be /userProfile/kid_avatar_{kidID}

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
    "id": 13,
    "name": "kid13",
    "dateCreated": "2017-01-29T23:07:38Z",
    "macId": "Mac_ID3",
    "profile": "kid_avatar_13.jpg",
    "parent": {
      "id": 5,
      "email": "jack08301@gmail.com",
      "name": "Jay",
      "lastUpdate": "2017-01-11T04:16:44Z",
      "dateCreated": "2017-01-11T04:16:44Z",
      "zipCode": "11111",
      "phoneNumber": "11111",
      "profile": "avatar_5.jpg",
      "registrationId": "123test"
    }
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
* It only responses added kid
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| name     | Yes | string | Kiddle   |
| macId      | Yes | string | 8D071FCFE5E0 |

### curl Example
```
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: 8158050a26569ff9c8c5297351c4f56a" -d '{
  "name": "Kids3",
  "macId": "Mac_ID4"
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
  "id": 2,
  "name": "Kids3",
  "dateCreated": "2017-02-02T01:14:16Z",
  "macId": "Mac_ID4",
  "profile": "",
  "parent": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "",
    "phoneNumber": "",
    "profile": ""
  }
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
| name    | No | String |   Jay |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "x-auth-token: pej57nakctvf7gcr7j9m7macdbad3637" -H "Content-Type: application/json" -d '{
	"name": "yes",
	"kidId": 19
}' "http://localhost:8111/v1/kids/update"
```

* Success - 
```
{
  "kid": {
    "id": 13,
    "name": "kid13",
    "dateCreated": "2017-01-29T23:07:38Z",
    "macId": "Mac_ID3",
    "profile": "",
    "parent": {
      "id": 5,
      "email": "jack08301@gmail.com",
      "firstName": "JJJ",
      "lastName": "TTT",
      "lastUpdate": "2017-01-11T04:16:44Z",
      "dateCreated": "2017-01-11T04:16:44Z",
      "zipCode": "11111",
      "phoneNumber": "412412412412",
      "profile": "avatar_5.jpg"
    }
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

## /v1/kids/delete - DELETE
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***
* It can only delete logged in user's kid data

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| kidId    | Yes | Integer |   3 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Delete successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X DELETE -H "Content-Type: application/json" -H "x-auth-token: 1c8306065183d7829fab224be6e43257" "http://localhost:8111/v1/kids/delete?kidId=2"
```

* Success - no response body

## /v1/kids/list - GET
* Retrieve kids belong to the signed in user
* Doesn't need any parameter

#### Request Parameters
None


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 403     | User is not signed in |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: 28e20ffb974205c0747b5aa35d53e53" "http://localhost:8111/v1/kids/list"
```

* Success - 
```
[
  {
    "id": 3,
    "name": "kid13",
    "dateCreated": "2017-02-02T10:25:20Z",
    "macId": "012345678915",
    "profile": "kid_avatar_3.jpg"
  }
]
```

## /v1/kids/whoRegisteredMacID - GET
* Retrieve kid and user information by MAC ID
* Doesn't need x-auth-token

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
    "id": 2,
    "name": "KIDLLE123124",
    "dateCreated": "2017-01-16T22:18:26Z",
    "macId": "hgweorahgbkljwhnpi1",
    "profile": "qwe",
    "parent": {
      "id": 2,
      "email": "jack08300@gmail.com",
      "firstName": "Jay",
      "lastName": "Chen",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "0001-01-01T00:00:00Z",
      "zipCode": "11111",
      "phoneNumber": "11111",
      "profile": "",
      "registrationId": "123test"
    }
  }
}
```

## /v1/kids/batteryStatus - POST
* Upload device battery life
* Content-Type: application/json

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| batteryLife     | Yes | Int | 95   |
| macId      | Yes | string | 8D071FCFE5E0 |
| dateReceived      | Yes | Long | 1487290224 (Unix time) |

### curl Example
```
curl -X POST \
  http://localhost:8111/v1/kids/batteryStatus \
  -H 'x-auth-token: 421d691c595e5b4321f33facc037b956' \
  -d '{
  "macId": "test123",
  "batteryLife": 98,
  "DateReceived": 14442523
}'
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Data added successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - return empty JSON
```
{}
```

## /v1/kids/updateKidRevertMacID - PUT
* Fix kid Mac ID reverse issue
* It's put method, please add the parameter to the URL

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| macId     | Yes | String | 8D071FCFE5E0   |
| kidId      | Yes | Int | 8 |

### curl Example
```
curl -X PUT \
  'http://localhost:8111/v1/kids/updateKidRevertMacID?kidId=8&macId=E0E5CF1ED7C2' \
  -H 'x-auth-token: ab2763ef35f79cc15604eacaa181c824'
```

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Data added successfully |
| 400     | Bad request. Missing some parameters |
| 500     | Internal error. Please send me the error. I will fix it |

* Success - return empty JSON
```
{}
```

## /v1/activity/uploadRawData - POST
* Content-Type: application/json
* Activity data
* Upload with user's timezone offset (e.g. New York Timezone offset is ***-300***)
* When the data duplicate which is when server response ***409*** status, you can ***ignore*** it and process next data

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| -----:|
| indoorActivity      | Yes | String | 1481299119,0,216,2,3,4 |
| outdoorActivity     | Yes | String | 1481299119,1,0,0,0,0 |
| timeZoneOffset | Yes | Integer | -300 |
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
curl -X POST -H "x-auth-token: 28e20ffb974205c0747b5aa35d53e538" -H "Content-Type: application/json" -d '{
  "indoorActivity": "1487125375,0,10,2,3,4",
  "outdoorActivity": "1487125375,1,24,0,0,0",
  "time": 1487125375,
  "timeZoneOffset": -300,
  "macId": "012345678915"
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
| period     | Yes | String | DAILY, WEEKLY MONTHLY, YEARLY  |

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

## /v1/activity/retrieveDataByTime - GET
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| start      | Yes | Long ***Timestamp*** | 1491599032 |
| end        | Yes | Long ***Timestamp*** | 1498089090  |
| kidId      | Yes | Integer | 1  |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receiving Data successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "x-auth-token: 27e217ae4d2907188b8c92cdcf9c85ac" "http://localhost:8111/v1/activity/retrieveDataByTime?start=1491599032&end=1498089090&kidId=1"
```

* Success - 
```
{
  "activities": [
    {
      "id": 5,
      "macId": "hgweorahgbkljwhnpi3",
      "kidId": "1",
      "type": "INDOOR",
      "steps": 10,
      "distance": 0,
      "receivedDate": "2017-05-03T10:58:10Z"
    },
    {
      "id": 6,
      "macId": "hgweorahgbkljwhnpi3",
      "kidId": "1",
      "type": "OUTDOOR",
      "steps": 24,
      "distance": 0,
      "receivedDate": "2017-06-21T23:51:30Z"
    }
  ]
}
```

## /v1/activity/retrieveHourlyDataByTime - GET
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***
* Ordered by receivedDate

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| start      | Yes | Long ***Timestamp*** | 1491599032 |
| end        | Yes | Long ***Timestamp*** | 1498089090  |
| kidId      | Yes | Integer | 1  |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receiving Data successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET \
  'http://dev.childrenlab.com/v1/activity/retrieveHourlyDataByTime?start=1511654400&end=1511740799&kidId=120' \
  -H 'x-auth-token: 3c8b3fde05ca42f04a91ae1a67774a65' 
}'
```

* Success - 
```
{
    "activities": [
        {
            "id": 51,
            "macId": "testtter1",
            "kidId": 120,
            "type": "INDOOR",
            "steps": 10,
            "distance": 0,
            "receivedDate": "2017-11-26T23:59:59Z",
            "DateCreated": "2017-11-26T22:07:11Z",
            "LastUpdated": "2017-11-26T22:07:11Z"
        },
        {
            "id": 52,
            "macId": "testtter1",
            "kidId": 120,
            "type": "OUTDOOR",
            "steps": 24,
            "distance": 0,
            "receivedDate": "2017-11-26T23:59:59Z",
            "DateCreated": "2017-11-26T22:07:11Z",
            "LastUpdated": "2017-11-26T22:07:11Z"
        },
        {
            "id": 47,
            "macId": "testtter1",
            "kidId": 120,
            "type": "INDOOR",
            "steps": 10,
            "distance": 0,
            "receivedDate": "2017-11-26T10:53:20Z",
            "DateCreated": "2017-11-26T22:01:46Z",
            "LastUpdated": "2017-11-26T22:01:46Z"
        },
        {
            "id": 48,
            "macId": "testtter1",
            "kidId": 120,
            "type": "OUTDOOR",
            "steps": 24,
            "distance": 0,
            "receivedDate": "2017-11-26T10:53:20Z",
            "DateCreated": "2017-11-26T22:01:46Z",
            "LastUpdated": "2017-11-26T22:01:46Z"
        },
        {
            "id": 45,
            "macId": "testtter1",
            "kidId": 120,
            "type": "INDOOR",
            "steps": 10,
            "distance": 0,
            "receivedDate": "2017-11-26T08:06:40Z",
            "DateCreated": "2017-11-26T22:01:27Z",
            "LastUpdated": "2017-11-26T22:01:27Z"
        },
        {
            "id": 46,
            "macId": "testtter1",
            "kidId": 120,
            "type": "OUTDOOR",
            "steps": 24,
            "distance": 0,
            "receivedDate": "2017-11-26T08:06:40Z",
            "DateCreated": "2017-11-26T22:01:27Z",
            "LastUpdated": "2017-11-26T22:01:27Z"
        },
        {
            "id": 43,
            "macId": "testtter1",
            "kidId": 120,
            "type": "INDOOR",
            "steps": 20,
            "distance": 0,
            "receivedDate": "2017-11-26T07:50:00Z",
            "DateCreated": "2017-11-26T22:01:07Z",
            "LastUpdated": "2017-11-26T22:01:17Z"
        },
        {
            "id": 44,
            "macId": "testtter1",
            "kidId": 120,
            "type": "OUTDOOR",
            "steps": 48,
            "distance": 0,
            "receivedDate": "2017-11-26T07:50:00Z",
            "DateCreated": "2017-11-26T22:01:07Z",
            "LastUpdated": "2017-11-26T22:01:17Z"
        }
    ]
}
```

## /v1/activity/retrieveMonthlyActivity - GET
* It's ***GET*** method, so no content-type. Add parameters to the ***URL***
* Ordered by month

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| start      | Yes | Long ***Timestamp*** | 1491599032 |
| end        | Yes | Long ***Timestamp*** | 1498089090  |
| kidId      | Yes | Integer | 1  |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receiving Data successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET \
  'http://localhost:8111/v1/activity/retrieveMonthlyActivity?start=1411654400&end=1512237067&kidId=19' \
  -H 'x-auth-token: 3b9b03e26bf40f9e431b5d2b69b89c4c'
```

* Success - 
```
{
    "activities": [
        {
            "macId": "AAAAAABBBB01",
            "type": "INDOOR",
            "month": 2,
            "steps": 346,
            "distance": 0
        },
        {
            "macId": "AAAAAABBBB01",
            "type": "OUTDOOR",
            "month": 2,
            "steps": 346,
            "distance": 0
        },
        {
            "macId": "AAAAAABBBB01",
            "type": "INDOOR",
            "month": 11,
            "steps": 173,
            "distance": 0
        },
        {
            "macId": "AAAAAABBBB01",
            "type": "OUTDOOR",
            "month": 11,
            "steps": 173,
            "distance": 0
        }
    ]
}
```


## /v1/event/add - POST
* Content-Type: application/json
* Only when user has access to the kids, otherwise, the API returns 401 unauthorized
* Date Time format is ***YYYY-MM-ddThh:mm:ssZ***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| kidId             | Yes | Integer Array | [3,4] |
| name               | Yes | String | Test event name |
| startDate           | Yes | String |   2015-08-30T08:20:00 |
| endDate             | Yes | String |   2015-08-31T08:20:00 |
| color               | Yes | String |   #F05D25 |
| description         | No | String |   Test description |
| alert               | No | Integer |   49 |
| repeat               | No | String |   DAILY, MONTHLY |
| timezoneOffset       Yes | Integer |   300 |
| todo               | No | Array |   ["test todo 1", "test todo 2"] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Added successfully |
| 401     | Unauthorized. The user doesn't have permission to add event |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X POST -H "x-auth-token: 28e20ffb974205c0747b5aa35d53e538" -H "Content-Type: application/json" -d '{
  "kidId": [3,4],
  "Name": "Other kids event",
  "startDate": "2017-02-23T08:20:00Z",
  "endDate": "2017-08-31T08:20:00Z",
  "timezoneOffset": 300,
  "color": "#F05D25",
  "description": "ahhahahahahhaa",
  "alert": 25
  
}' "http://localhost:8111/v1/event/add"
```

* Success - Returns added event
```
{
  "event": {
    "id": 5,
    "user": {
      "id": 2,
      "email": "jack08300@gmail.com",
      "firstName": "Jay",
      "lastName": "Chen",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "0001-01-01T00:00:00Z",
      "zipCode": "",
      "phoneNumber": "",
      "profile": "avatar_2.jpg"
    },
    "kid": [
      {
        "id": 3,
        "name": "kid13",
        "dateCreated": "2017-02-02T10:25:20Z",
        "macId": "012345678915",
        "profile": "kid_avatar_3.jpg"
      },
      {
        "id": 4,
        "name": "5",
        "dateCreated": "2017-02-02T10:42:02Z",
        "macId": "012345678916",
        "profile": "kid_avatar_4.jpg"
      }
    ],
    "name": "Other kids event",
    "startDate": "2017-02-23T08:20:00Z",
    "endDate": "2017-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "ahhahahahahhaa",
    "alert": 25,
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-13T00:03:57Z",
    "lastUpdated": "2017-02-13T00:03:57Z"
  }
}
```

## /v1/event/update - PUT
* Content-Type: application/json
* ***User who creates the event has the permission to update the event***
* Date Time format is ***YYYY-MM-ddThh:mm:ssZ***
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
| repeat               | No | String |   DAILY, MONTHLY |
| timezoneOffset       Yes | Integer |   300 |
| todo               | No | Array |   ["test todo 1"] |
| kidId               | Yes | Array |   [118, 120] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Updated successfully |
| 401     | Unauthorized. The user doesn't have permission to update |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT \
  http://localhost:8111/v1/event/update \
  -H 'content-type: application/json' \
  -H 'x-auth-token: 3c8b3fde05ca42f04a91ae1a67774a65' \
  -d '{
"color":"#7572C1",
"description":"123",
"endDate":"2017-02-20T12:00:00Z",
"kidId":[118, 120],
"name":"Storey",
"repeat":"",
"startDate":"2017-02-17T09:30:00Z",
"state":"",
"todo":["todo1"],
"alert":61,
"eventId":402,
"timezoneOffset":0
}'
```

* Success - Returns updated event
```
{
    "event": {
        "id": 402,
        "user": {
            "id": 348,
            "email": "nn17@nn.com",
            "firstName": "jdhx",
            "lastName": "djchsj",
            "lastUpdate": "2017-11-12T20:10:12Z",
            "dateCreated": "2017-11-12T20:10:12Z",
            "zipCode": "",
            "phoneNumber": "5686946894",
            "profile": "",
            "language": "en",
            "ios_registration_id": "",
            "android_registration_id": "dQYymla3XIY:APA91bHjW_PFjbnXDSP-Xjag_49UKgtYTmM3w_PpVD_Z_trPRo0PU-B22XRVJatlW_0Dr94phK8M_AlYEv7DnyCo3_-SOlxYGLRVtugWELmcEzwsg_j-4kL3gc8BErfc61eat55UVC0Q",
            "country": "US"
        },
        "kid": [
            {
                "id": 118,
                "name": "kha hahx",
                "dateCreated": "2017-11-12T20:12:48Z",
                "macId": "606405861CC9",
                "firmwareVersion": "KDV0006-A",
                "profile": ""
            },
            {
                "id": 120,
                "name": "JOHN",
                "dateCreated": "2017-11-19T21:34:45Z",
                "macId": "testtter1",
                "firmwareVersion": "",
                "profile": ""
            }
        ],
        "name": "Storey",
        "startDate": "2017-02-17T09:30:00Z",
        "endDate": "2017-02-20T12:00:00Z",
        "color": "#7572C1",
        "status": "OPEN",
        "description": "123",
        "alert": 61,
        "repeat": "",
        "timezoneOffset": 0,
        "dateCreated": "2017-11-19T21:31:02Z",
        "lastUpdated": "2017-11-19T21:31:02Z",
        "todo": [
            {
                "id": 72,
                "text": "todo1",
                "status": "PENDING",
                "dateCreated": "2017-11-19T21:35:00Z",
                "lastUpdated": "2017-11-19T21:35:00Z"
            }
        ]
    }
}
```

## /v1/event/delete - DELETE
* URL Query
* ***User who creates the event has the permission to delete the event***

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| eventId             | Yes | Integer | 414 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Delete successfully |
| 401     | Unauthorized. The user doesn't have permission to delete |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X DELETE -H "x-auth-token: ec83d6e41db5168ddb0b1d28b2e262d6" "http://localhost:8111/v1/event/delete?eventId=2"
```

* Success - Returns empty json
```
{}
```

## /v1/event/retrieveAllEventsWithTodo - GET
* User can only see kid that has permission to him 
* Date Time format is ***YYYY-MM-ddThh:mm:ssZ***

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "x-auth-token: 42142142421421312312312" "http://localhost:8111/v1/event/retrieveAllEventsWithTodo"
```

* Success - Returns events
```
[
  {
    "id": 4,
    "user": {
      "id": 2,
      "email": "jack08300@gmail.com",
      "firstName": "Jay",
      "lastName": "Chen",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "0001-01-01T00:00:00Z",
      "zipCode": "",
      "phoneNumber": "",
      "profile": "avatar_2.jpg"
    },
    "kid": [
      {
        "id": 5,
        "name": "hello",
        "dateCreated": "2017-02-03T00:40:15Z",
        "macId": "8D071FCFE5E0",
        "profile": "kid_avatar_5.jpg"
      }
    ],
    "name": "Test event name2",
    "startDate": "2017-02-23T08:20:00Z",
    "endDate": "2017-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hafdewadhah",
    "alert": 49,
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-12T22:41:55Z",
    "lastUpdated": "2017-02-12T22:41:55Z",
    "todo": [
      {
        "id": 12,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2017-02-13T04:59:44Z",
        "lastUpdated": "2017-02-13T01:01:51Z"
      }
    ]
  },
  {
    "id": 5,
    "user": {
      "id": 2,
      "email": "jack08300@gmail.com",
      "firstName": "Jay",
      "lastName": "Chen",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "0001-01-01T00:00:00Z",
      "zipCode": "",
      "phoneNumber": "",
      "profile": "avatar_2.jpg"
    },
    "kid": [
      {
        "id": 4,
        "name": "5",
        "dateCreated": "2017-02-02T10:42:02Z",
        "macId": "012345678916",
        "profile": "kid_avatar_4.jpg"
      }
    ],
    "name": "Other kids event",
    "startDate": "2017-02-23T08:20:00Z",
    "endDate": "2017-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "ahhahahahahhaa",
    "alert": 25,
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-13T00:03:57Z",
    "lastUpdated": "2017-02-13T00:03:57Z"
  }
]
```

## /v1/event/retrieveAllEventsByKid - GET

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| kidId        | Yes | Integer | 3 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 401     | Unauthorized, The user doesn't have access to the kid
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "x-auth-token: 42142142421421312312312" "http://localhost:8111/v1/event/retrieveAllEventsByKid?kidId=5"

```

* Success - Returns events
```
[
  {
    "id": 4,
    "user": {
      "id": 2,
      "email": "jack08300@gmail.com",
      "firstName": "Jay",
      "lastName": "Chen",
      "lastUpdate": "0001-01-01T00:00:00Z",
      "dateCreated": "0001-01-01T00:00:00Z",
      "zipCode": "",
      "phoneNumber": "",
      "profile": "avatar_2.jpg"
    },
    "kid": null,
    "name": "Test event name2",
    "startDate": "2017-02-23T08:20:00Z",
    "endDate": "2017-08-31T08:20:00Z",
    "color": "#F05D25",
    "status": "OPEN",
    "description": "Hafdewadhah",
    "alert": 49,
    "repeat": "",
    "timezoneOffset": 300,
    "dateCreated": "2017-02-12T22:41:55Z",
    "lastUpdated": "2017-02-12T22:41:55Z",
    "todo": [
      {
        "id": 12,
        "text": "test todo 2",
        "status": "PENDING",
        "dateCreated": "2017-02-13T04:59:44Z",
        "lastUpdated": "2017-02-13T01:01:51Z"
      }
    ]
  }
]
```


## /v1/event/todo/done - PUT
* Set status = done to todo

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| eventId        | Yes | Integer | 2 |
| todoId          | Yes | Integer | 2 |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | updated successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "x-auth-token: ec83d6e41db5168ddb0b1d28b2e262d6" -H "Content-Type: application/json" -d '{
  "eventId": 2,
  "todoId": 2
}' "http://localhost:8111/v1/event/todo/done"
```

* Success - 
```
{}
```


## /v1/subHost/add - POST
* Send sub host request to the host account
* Content-Type: application/json
* Get ```hostId``` from [/v1/kids/whoRegisteredMacID](#v1kidswhoregisteredmacid---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
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
curl -X POST -H "Content-Type: application/json" -H "x-auth-token: 4094345fecef667df936a632bafabbe6" -d '{
  "hostId": 63
}' "http://localhost:8111/v1/subHost/add"
```

* Success - Returns added request
```
{
  "id": 3,
  "requestFromUser": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "11111",
    "phoneNumber": "11111",
    "profile": "",
    "registrationId": "123test"
  },
  "requestToUser": {
    "id": 63,
    "email": "jack08301@gmail.com",
    "firstName": "JJJ",
    "lastName": "TTT",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "2017-01-17T00:56:06Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  },
  "status": "PENDING",
  "createdDate": "2017-01-18T04:03:22Z",
  "lastUpdated": "2017-01-18T04:03:22Z"
}
```

* Fail (409) -
```
{
  "message": "The request is already exist"
}
```

## /v1/subHost/accept - PUT
* Content-Type: application/json
* Accept the sub host request by allowing sub host to view the kids
* When adding a kid under a subhost, use this API as well
* Get ```subHostId``` from [GET   /v1/subHost/list](#v1subhostlist---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| subHostId             | Yes | Integer | 10 |
| KidId             | Yes | Integer Array | [22, 10] |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Accept successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong. Or the logged in user doesn't have permission |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: 6d312cbc54ce05a279dcc554b984061c" -d '{
  "subHostId": 1,
  "kidId": [
  	9, 10
  ]
}' "http://localhost:8111/v1/subHost/accept"
```

* Success - Returns updated request
```
{
  "id": 3,
  "requestFromUser": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "11111",
    "phoneNumber": "11111",
    "profile": "",
    "registrationId": "123test"
  },
  "requestToUser": {
    "id": 63,
    "email": "jack08301@gmail.com",
    "firstName": "JJJ",
    "lastName": "TTT",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "2017-01-17T00:56:06Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  },
  "status": "ACCEPTED",
  "createdDate": "2017-01-18T04:03:22Z",
  "lastUpdated": "2017-01-18T04:03:22Z",
  "kids": [
    {
      "id": 9,
      "name": "Another KID",
      "dateCreated": "2017-01-17T00:56:58Z",
      "macId": "Test2",
      "profile": ""
    },
    {
      "id": 10,
      "name": "KIDLLE123124",
      "dateCreated": "2017-01-17T02:39:36Z",
      "macId": "Test3",
      "profile": ""
    }
  ]
}
```

## /v1/subHost/deny - PUT
* Content-Type: application/json
* It will delete subhost request
* Get ```subHostId``` from [GET   /v1/subHost/list](#v1subhostlist---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| subHostId             | Yes | Integer | 2 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Accept successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: 58lkp329ejbr4498st59ur2na7e0rmtg" -d '{
  "subHostId": 2
}' "http://localhost:8111/v1/subHost/deny"
```

* Success - Delete completed
```
{}
```

## /v1/subHost/list - GET
* Content-Type: application/json
* If no status parameter, the API returns ALL of sub host belong to the user
* It returns request from and request to Subhost list 

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| status        | No | String | PENDING, ACCEPTED, DENIED |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Receive list successfully |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET -H "Content-Type: application/json" -H "x-auth-token: 58lkp329ejbr4498st59ur2na7e0rmtg" "http://localhost:8111/v1/subHost/list"
```

* Success - Returns updated request
```
{
  "requestFrom": [
    {
      "id": 1,
      "requestFromUser": {
        "id": 3,
        "email": "spencer@kidsdynamic.com",
        "firstName": "JJJ",
        "lastName": "TTT",
        "lastUpdate": "0001-01-01T00:00:00Z",
        "dateCreated": "2017-01-31T05:06:20Z",
        "zipCode": "11111",
        "phoneNumber": "",
        "profile": "avatar_12.jpg"
      },
      "requestToUser": {
        "id": 2,
        "email": "jack08300@gmail.com",
        "firstName": "Jay",
        "lastName": "Chen",
        "lastUpdate": "0001-01-01T00:00:00Z",
        "dateCreated": "0001-01-01T00:00:00Z",
        "zipCode": "",
        "phoneNumber": "",
        "profile": "avatar_2.jpg"
      },
      "status": "ACCEPTED",
      "createdDate": "2017-01-30T02:19:29Z",
      "lastUpdated": "2017-01-30T02:19:29Z",
      "kids": [
        {
          "id": 3,
          "name": "kid13",
          "dateCreated": "2017-02-02T10:25:20Z",
          "macId": "012345678915",
          "profile": "kid_avatar_3.jpg"
        }
      ]
    },
    {
      "id": 2,
      "requestFromUser": {
        "id": 9,
        "email": "test10@swing.com",
        "firstName": "1",
        "lastName": "2",
        "lastUpdate": "0001-01-01T00:00:00Z",
        "dateCreated": "2017-02-02T10:40:20Z",
        "zipCode": "4",
        "phoneNumber": "3",
        "profile": "avatar_12.jpg",
        "registrationId": "B16DDF222DFC1BF75DEF597E9AD14E9AC99935C48F964361A855731DAECD39FE"
      },
      "requestToUser": {
        "id": 2,
        "email": "jack08300@gmail.com",
        "firstName": "Jay",
        "lastName": "Chen",
        "lastUpdate": "0001-01-01T00:00:00Z",
        "dateCreated": "0001-01-01T00:00:00Z",
        "zipCode": "",
        "phoneNumber": "",
        "profile": "avatar_2.jpg"
      },
      "status": "PENDING",
      "createdDate": "2017-02-11T17:28:31Z",
      "lastUpdated": "2017-02-11T17:28:31Z"
    }
  ],
  "requestTo": [
    {
      "id": 3,
      "requestFromUser": {
        "id": 2,
        "email": "jack08300@gmail.com",
        "firstName": "Jay",
        "lastName": "Chen",
        "lastUpdate": "0001-01-01T00:00:00Z",
        "dateCreated": "0001-01-01T00:00:00Z",
        "zipCode": "",
        "phoneNumber": "",
        "profile": "avatar_2.jpg"
      },
      "requestToUser": {
        "id": 3,
        "email": "spencer@kidsdynamic.com",
        "firstName": "JJJ",
        "lastName": "TTT",
        "lastUpdate": "0001-01-01T00:00:00Z",
        "dateCreated": "2017-01-31T05:06:20Z",
        "zipCode": "11111",
        "phoneNumber": "",
        "profile": "avatar_12.jpg"
      },
      "status": "PENDING",
      "createdDate": "2017-02-11T17:28:57Z",
      "lastUpdated": "2017-02-11T17:28:57Z"
    }
  ]
}
```

## /v1/subHost/removeKid - PUT
* Content-Type: application/json
* It will delete a kid under the SubHost
* Get ```subHostId``` from [GET   /v1/subHost/list](#v1subhostlist---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| subHostId     | Yes | Integer | 2 |
| kidId         | Yes | Integer | 1 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Accept successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 401     | Unauthorized. The user doesn't have permission |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X PUT -H "Content-Type: application/json" -H "x-auth-token: 0838289b3d518dd263e2356900a00241" -d '{
  "subHostId": 1,
  "kidId": 1
}' "http://localhost:8111/v1/subHost/removeKid"
```

* Success - Delete completed, and return sub host info
```
{
  "id": 1,
  "requestFromUser": {
    "id": 3,
    "email": "jack083002@gmail.com",
    "firstName": "JJJ",
    "lastName": "TTT",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "2017-02-19T13:48:45Z",
    "zipCode": "11111",
    "phoneNumber": "",
    "profile": ""
  },
  "requestToUser": {
    "id": 2,
    "email": "jack08300@gmail.com",
    "firstName": "Jay",
    "lastName": "Chen",
    "lastUpdate": "0001-01-01T00:00:00Z",
    "dateCreated": "0001-01-01T00:00:00Z",
    "zipCode": "",
    "phoneNumber": "",
    "profile": ""
  },
  "status": "ACCEPTED",
  "createdDate": "2017-02-19T13:51:02Z",
  "lastUpdated": "2017-02-19T13:51:02Z",
  "kids": [
    {
      "id": 3,
      "name": "second kids3",
      "dateCreated": "2017-02-19T15:49:28Z",
      "macId": "Mac_ID3",
      "profile": ""
    }
  ]
}
```

## /v1/subHost/delete - DELETE
* It's ***DELETE*** method, so no content-type. Add parameters to the ***URL***
* It will delete the SubHost
* Get ```subHostId``` from [GET   /v1/subHost/list](#v1subhostlist---get) API

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| subHostId     | Yes | Integer | 25 |


#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Delete successfully |
| 400     | Bad request. Missing some parameters, or the type is wrong |
| 401     | Unauthorized. The user doesn't have permission |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl --request DELETE \
  --url 'http://localhost:8111/v1/subHost/delete?subHostId=25' \
  --header 'x-auth-token: b0d2d81ee32cc8c9ff276f4469b0dd61'
```

* Success - Delete completed, and return sub host info
```
{}
```

## /v1/fw/currentVersion/ - GET
* Retrieve current Firmware version
* The API only return latest firmware version
* the Mac ID is on the path
* The file is on Amazon S3

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| macId     | Yes | String | 6064058630BA |
| fwVersion | Yes | String | KDV0109-CN |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Retrieve successfully |
| 400     | Bad request. The token is invalid or mac id is not present |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET \
  'http://dev.childrenlab.com/v1/fw/currentVersion/?macId=6064058630BA&fwVersion=KDV0109-CN' \
  -H 'x-auth-token: f64ca9d8de4cb508ba6b6e6b4f1462d6'
```
* Success - with kids data
```
{
    "id": 15,
    "version": "KDV0109-CN",
    "fileAUrl": "fw_version/KDV0109-CNA.hex",
    "fileBUrl": "fw_version/KDV0109-CNB.hex",
    "uploadedDate": "2018-01-19T03:27:07Z",
    "active": true
}
```

## /v1/fw/firmwareVersion - PUT
* Send device firmware version to backend
* Send the firmware version every time when user sync

#### Request Parameters
| Parameters    | Required      | Type  | Example  |
| ------------- |:-------------:|:-------------:| :-----|
| macId     | Yes | String | A81B6ABA0749 |
| firmwareVersion | Yes | String | KDV0106-J |

#### Response Status
| Status Code    | Meaning      |
| ------------- |:-------------|
| 200     | Update successfully |
| 400     | Bad request. The token is invalid or mac id is not present |
| 500     | Internal error. Please send me the error. I will fix it |

### curl
```
curl -X GET \
  http://localhost:8111/v1/fw/currentVersion/E01786056460 \
  -H 'x-auth-token: 04b62b70c8464bafb84d2de2464024a2'
```
* Success - No Data
```
{}
```
