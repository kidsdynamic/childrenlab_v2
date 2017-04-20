webpackJsonp([1,4],{

/***/ 115:
/***/ (function(module, exports) {

function webpackEmptyContext(req) {
	throw new Error("Cannot find module '" + req + "'.");
}
webpackEmptyContext.keys = function() { return []; };
webpackEmptyContext.resolve = webpackEmptyContext;
module.exports = webpackEmptyContext;
webpackEmptyContext.id = 115;


/***/ }),

/***/ 116:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_platform_browser_dynamic__ = __webpack_require__(122);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__app_app_module__ = __webpack_require__(125);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__environments_environment__ = __webpack_require__(82);




if (__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].production) {
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["a" /* enableProdMode */])();
}
__webpack_require__.i(__WEBPACK_IMPORTED_MODULE_1__angular_platform_browser_dynamic__["a" /* platformBrowserDynamic */])().bootstrapModule(__WEBPACK_IMPORTED_MODULE_2__app_app_module__["a" /* AppModule */]);
//# sourceMappingURL=main.js.map

/***/ }),

/***/ 12:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_http__ = __webpack_require__(42);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__model_admin_login__ = __webpack_require__(43);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__environments_environment__ = __webpack_require__(82);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4_rxjs_add_operator_toPromise__ = __webpack_require__(225);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4_rxjs_add_operator_toPromise___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_4_rxjs_add_operator_toPromise__);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_5_ng2_webstorage__ = __webpack_require__(22);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ServerService; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};






var ServerService = (function () {
    function ServerService(http) {
        this.http = http;
        this.LOGIN_API = "/admin/login";
        this.TOKEN_VALIDATION_API = '/v1/user/isTokenValid';
        this.USER_LIST_API = '/admin/userList';
        this.KID_LIST_API = '/admin/kidList';
        this.ACTIVITY_LIST_API = '/admin/activityList';
        this.ACTIVITY_RAW_LIST_API = '/admin/activityRawList';
        this.DASHBOARD_API = '/admin/dashboard';
    }
    ServerService.prototype.login = function (userName, password) {
        var loginJson = {
            name: userName,
            password: password
        };
        return this.http.post(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].BaseURL + this.LOGIN_API, JSON.stringify(loginJson))
            .toPromise()
            .then(function (response) { return response.json(); });
    };
    ServerService.prototype.getUserList = function () {
        var options = this.addTokenHeader();
        return this.http.get("" + (__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].BaseURL + this.USER_LIST_API), options)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handleError);
    };
    ServerService.prototype.getKidList = function () {
        var options = this.addTokenHeader();
        return this.http.get("" + (__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].BaseURL + this.KID_LIST_API), options)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handleError);
    };
    ServerService.prototype.getActivityListByKidId = function (kidId) {
        var options = this.addTokenHeader();
        return this.http.get(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].BaseURL + this.ACTIVITY_LIST_API + "/" + kidId, options)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handleError);
    };
    ServerService.prototype.getActivityRawListByKidId = function (kidId) {
        var options = this.addTokenHeader();
        return this.http.get(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].BaseURL + this.ACTIVITY_RAW_LIST_API + "/" + kidId, options)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handleError);
    };
    ServerService.prototype.getDashboardData = function () {
        var options = this.addTokenHeader();
        return this.http.get("" + (__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].BaseURL + this.DASHBOARD_API), options)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handleError);
    };
    ServerService.prototype.tokenValidation = function () {
        if (this.token == null) {
            return Promise.reject("Invalid token");
        }
        return this.http.get(__WEBPACK_IMPORTED_MODULE_3__environments_environment__["a" /* environment */].BaseURL + this.TOKEN_VALIDATION_API + "?email=" + this.token.username + "&token=" + this.token.access_token)
            .toPromise()
            .then();
    };
    ServerService.prototype.addTokenHeader = function () {
        var headers = new __WEBPACK_IMPORTED_MODULE_1__angular_http__["b" /* Headers */]({
            'x-auth-token': this.token.access_token,
            'Content-Type': 'application/json'
        });
        var options = new __WEBPACK_IMPORTED_MODULE_1__angular_http__["c" /* RequestOptions */]({ headers: headers });
        return options;
    };
    ServerService.prototype.handleError = function (error) {
        console.error('Error: ', error);
        return Promise.reject(error.message || error);
    };
    return ServerService;
}());
__decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_5_ng2_webstorage__["b" /* LocalStorage */])(),
    __metadata("design:type", typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_2__model_admin_login__["a" /* AdminToken */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_2__model_admin_login__["a" /* AdminToken */]) === "function" && _a || Object)
], ServerService.prototype, "token", void 0);
ServerService = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["c" /* Injectable */])(),
    __metadata("design:paramtypes", [typeof (_b = typeof __WEBPACK_IMPORTED_MODULE_1__angular_http__["d" /* Http */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__angular_http__["d" /* Http */]) === "function" && _b || Object])
], ServerService);

var _a, _b;
//# sourceMappingURL=server.service.js.map

/***/ }),

/***/ 124:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_router__ = __webpack_require__(21);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AppComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var AppComponent = (function () {
    function AppComponent(router) {
        this.router = router;
    }
    AppComponent.prototype.ngOnInit = function () {
        this.router.navigate(["login"]);
    };
    return AppComponent;
}());
AppComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-root',
        template: __webpack_require__(200),
        styles: [__webpack_require__(184)]
    }),
    __metadata("design:paramtypes", [typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_1__angular_router__["c" /* Router */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__angular_router__["c" /* Router */]) === "function" && _a || Object])
], AppComponent);

var _a;
//# sourceMappingURL=app.component.js.map

/***/ }),

/***/ 125:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_platform_browser__ = __webpack_require__(16);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_forms__ = __webpack_require__(74);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__angular_http__ = __webpack_require__(42);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4__angular_platform_browser_animations__ = __webpack_require__(123);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_5__app_component__ = __webpack_require__(124);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_6__angular_material__ = __webpack_require__(121);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_7__server_service__ = __webpack_require__(12);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_8_ng2_webstorage__ = __webpack_require__(22);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_9__dashboard_dashboard_component__ = __webpack_require__(78);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_10__route_route_module__ = __webpack_require__(127);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_11__login_login_component__ = __webpack_require__(80);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_12__user_list_user_list_component__ = __webpack_require__(81);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_13__kid_list_kid_list_component__ = __webpack_require__(79);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_14__activity_activity_component__ = __webpack_require__(76);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_15__activity_raw_activity_raw_component__ = __webpack_require__(75);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_16__dashboard_main_dashboard_main_component__ = __webpack_require__(77);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AppModule; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};

















var AppModule = (function () {
    function AppModule() {
    }
    return AppModule;
}());
AppModule = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_1__angular_core__["b" /* NgModule */])({
        declarations: [
            __WEBPACK_IMPORTED_MODULE_5__app_component__["a" /* AppComponent */],
            __WEBPACK_IMPORTED_MODULE_9__dashboard_dashboard_component__["a" /* DashboardComponent */],
            __WEBPACK_IMPORTED_MODULE_11__login_login_component__["a" /* LoginComponent */],
            __WEBPACK_IMPORTED_MODULE_12__user_list_user_list_component__["a" /* UserListComponent */],
            __WEBPACK_IMPORTED_MODULE_13__kid_list_kid_list_component__["a" /* KidListComponent */],
            __WEBPACK_IMPORTED_MODULE_14__activity_activity_component__["a" /* ActivityComponent */],
            __WEBPACK_IMPORTED_MODULE_15__activity_raw_activity_raw_component__["a" /* ActivityRawComponent */],
            __WEBPACK_IMPORTED_MODULE_16__dashboard_main_dashboard_main_component__["a" /* DashboardMainComponent */]
        ],
        imports: [
            __WEBPACK_IMPORTED_MODULE_0__angular_platform_browser__["a" /* BrowserModule */],
            __WEBPACK_IMPORTED_MODULE_2__angular_forms__["a" /* FormsModule */],
            __WEBPACK_IMPORTED_MODULE_3__angular_http__["a" /* HttpModule */],
            __WEBPACK_IMPORTED_MODULE_10__route_route_module__["a" /* RouteModule */],
            __WEBPACK_IMPORTED_MODULE_4__angular_platform_browser_animations__["a" /* BrowserAnimationsModule */],
            __WEBPACK_IMPORTED_MODULE_6__angular_material__["a" /* MdButtonModule */],
            __WEBPACK_IMPORTED_MODULE_6__angular_material__["b" /* MdInputModule */],
            __WEBPACK_IMPORTED_MODULE_6__angular_material__["c" /* MdToolbarModule */],
            __WEBPACK_IMPORTED_MODULE_8_ng2_webstorage__["a" /* Ng2Webstorage */],
            __WEBPACK_IMPORTED_MODULE_6__angular_material__["d" /* MdProgressSpinnerModule */],
            __WEBPACK_IMPORTED_MODULE_6__angular_material__["e" /* MdGridListModule */]
        ],
        providers: [
            __WEBPACK_IMPORTED_MODULE_7__server_service__["a" /* ServerService */]
        ],
        bootstrap: [__WEBPACK_IMPORTED_MODULE_5__app_component__["a" /* AppComponent */]]
    })
], AppModule);

//# sourceMappingURL=app.module.js.map

/***/ }),

/***/ 126:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return Dashboard; });
var Dashboard = (function () {
    function Dashboard() {
    }
    return Dashboard;
}());

var DashboardSignup = (function () {
    function DashboardSignup() {
    }
    return DashboardSignup;
}());
var DashboardActivity = (function () {
    function DashboardActivity() {
    }
    return DashboardActivity;
}());
//# sourceMappingURL=dashboard.js.map

/***/ }),

/***/ 127:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_router__ = __webpack_require__(21);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__dashboard_dashboard_component__ = __webpack_require__(78);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__login_login_component__ = __webpack_require__(80);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4__user_list_user_list_component__ = __webpack_require__(81);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_5__kid_list_kid_list_component__ = __webpack_require__(79);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_6__activity_activity_component__ = __webpack_require__(76);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_7_app_activity_raw_activity_raw_component__ = __webpack_require__(75);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_8__dashboard_main_dashboard_main_component__ = __webpack_require__(77);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return RouteModule; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};









var routes = [
    { path: 'login', component: __WEBPACK_IMPORTED_MODULE_3__login_login_component__["a" /* LoginComponent */] },
    { path: '', component: __WEBPACK_IMPORTED_MODULE_3__login_login_component__["a" /* LoginComponent */] },
    {
        path: 'dashboard',
        component: __WEBPACK_IMPORTED_MODULE_2__dashboard_dashboard_component__["a" /* DashboardComponent */],
        children: [
            {
                path: 'userList',
                component: __WEBPACK_IMPORTED_MODULE_4__user_list_user_list_component__["a" /* UserListComponent */]
            },
            {
                path: 'kidList',
                component: __WEBPACK_IMPORTED_MODULE_5__kid_list_kid_list_component__["a" /* KidListComponent */]
            },
            {
                path: 'activity/:kidId',
                component: __WEBPACK_IMPORTED_MODULE_6__activity_activity_component__["a" /* ActivityComponent */]
            },
            {
                path: 'activity-raw/:macId',
                component: __WEBPACK_IMPORTED_MODULE_7_app_activity_raw_activity_raw_component__["a" /* ActivityRawComponent */]
            },
            {
                path: 'dashboard-main',
                component: __WEBPACK_IMPORTED_MODULE_8__dashboard_main_dashboard_main_component__["a" /* DashboardMainComponent */]
            }
        ]
    },
];
var RouteModule = (function () {
    function RouteModule() {
    }
    return RouteModule;
}());
RouteModule = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["b" /* NgModule */])({
        imports: [__WEBPACK_IMPORTED_MODULE_1__angular_router__["a" /* RouterModule */].forRoot(routes, { useHash: true })],
        exports: [__WEBPACK_IMPORTED_MODULE_1__angular_router__["a" /* RouterModule */]]
    })
], RouteModule);

//# sourceMappingURL=route.module.js.map

/***/ }),

/***/ 182:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 183:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 184:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 185:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 186:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "button {\n  font-size: 14px;\n  line-height: 14px;\n  height: 30px;\n  min-width: 75px; }\n\n.logout {\n  float: right; }\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 187:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 188:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "form {\n  text-align: center; }\n  form md-input-container {\n    width: 250px; }\n", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 189:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(7)();
// imports


// module
exports.push([module.i, "", ""]);

// exports


/*** EXPORTS FROM exports-loader ***/
module.exports = module.exports.toString();

/***/ }),

/***/ 198:
/***/ (function(module, exports) {

module.exports = "<h3> Activity - </h3>\n<a (click)=\"goBack()\" class=\"btn btn-link\"><span class=\"glyphicon glyphicon-chevron-left\"></span>Back</a>\n<div class=\"error\">\n    {{ error }}\n</div>\n<table class=\"table table-striped table-hover\">\n    <thead>\n        <tr>\n            <th>\n                ID\n            </th>\n            <th>\n                Indoor\n            </th>\n            <th>\n                Outdoor\n            </th>\n            <th>\n                Time\n            </th>\n            <th>\n                Timezone Offset\n            </th>\n            <th>\n                Date Created\n            </th>\n        </tr>\n    </thead>\n    <tbody>\n        <tr *ngFor=\"let activity of activityRawList\">\n            <td>\n                {{ activity.id }}\n            </td>\n            <td style=\"color: darkcyan;\">\n                {{ activity.indoorActivity }}\n            </td>\n            <td>\n                {{ activity.outdoorActivity }}\n            </td>\n            <td>\n                {{ activity.time }}\n            </td>\n            <td>\n                {{ activity.timeZoneOffset }}\n            </td>\n            <td>\n                {{ activity.dateCreated }}\n            </td>\n\n        </tr>\n    </tbody>\n\n</table>\n"

/***/ }),

/***/ 199:
/***/ (function(module, exports) {

module.exports = "<h3> Activity - </h3>\n<a routerLink=\"/dashboard/kidList\" class=\"btn btn-link\"><span class=\"glyphicon glyphicon-chevron-left\"></span>Back</a>\n<div class=\"error\">\n    {{ error }}\n</div>\n<table class=\"table table-striped table-hover\">\n    <thead>\n        <tr>\n            <th>\n                ID\n            </th>\n            <th>\n                Type\n            </th>\n            <th>\n                Steps\n            </th>\n            <th>\n                Date\n            </th>\n        </tr>\n    </thead>\n    <tbody>\n        <tr *ngFor=\"let activity of activityList\">\n            <td>\n                {{ activity.id }}\n            </td>\n            <td>\n                {{ activity.type }}\n            </td>\n            <td>\n                {{ activity.steps }}\n            </td>\n            <td>\n                {{ activity.receivedDate }}\n            </td>\n\n        </tr>\n    </tbody>\n\n</table>"

/***/ }),

/***/ 200:
/***/ (function(module, exports) {

module.exports = "<header>\n  Kids Dynamic Server\n</header>\n<div class=\"page-container\">\n    <router-outlet></router-outlet>\n\n</div>\n"

/***/ }),

/***/ 201:
/***/ (function(module, exports) {

module.exports = "<h3> Dashboard </h3>\n<div class=\"error\">\n  {{ error }}\n</div>\n\n<div class=\"dashboardTable\">\n  <table class=\"table table-striped table-hover\">\n    <thead>\n    <tr>\n      <th>\n        Sign-up Count\n      </th>\n      <th>\n        Date\n      </th>\n    </tr>\n    </thead>\n    <tbody>\n    <tr *ngFor=\"let d of dashboard.signup\">\n      <td>\n        {{ d.signupCount }}\n      </td>\n      <td>\n        {{ d.date }}\n      </td>\n\n    </tr>\n    </tbody>\n  </table>\n</div>\n<div class=\"dashboardTable\">\n  <table class=\"table table-striped table-hover\">\n    <thead>\n    <tr>\n      <th>\n        Activity Count\n      </th>\n      <th>\n        Date\n      </th>\n    </tr>\n    </thead>\n    <tbody>\n    <tr *ngFor=\"let d of dashboard.activity\">\n      <td>\n        {{ d.activityCount }}\n      </td>\n      <td>\n        {{ d.date }}\n      </td>\n\n    </tr>\n    </tbody>\n  </table>\n</div>\n\n\n"

/***/ }),

/***/ 202:
/***/ (function(module, exports) {

module.exports = "<div class=\"dashboard\">\n  <legend>\n    Menu\n  </legend>\n    <button md-raised-button [routerLink]=\"['/dashboard/dashboard-main']\">Dashboard</button>\n    <button md-raised-button [routerLink]=\"['/dashboard/userList']\">User</button>\n    <button md-raised-button [routerLink]=\"['/dashboard/kidList']\">Kid</button>\n    <button md-raised-button (click)=\"logout()\" class=\"logout\" color=\"accent\">Logout</button>\n  <hr/>\n</div>\n\n<div class=\"container\">\n  <router-outlet></router-outlet>\n</div>\n"

/***/ }),

/***/ 203:
/***/ (function(module, exports) {

module.exports = "<h3> Device-Kid List</h3>\n<div class=\"error\">\n    <label class=\"error-message label-danger label\" *ngIf=\"error\">{{ error }}</label>\n</div>\n<table class=\"table table-striped table-hover\">\n    <thead>\n        <tr>\n            <th>\n                ID\n            </th>\n            <th>\n                Activity\n            </th>\n            <th>\n                Name\n            </th>\n            <th>\n                Mac ID\n            </th>\n            <th>\n                Profile\n            </th>\n            <th>\n                Parent\n            </th>\n            <th>\n                Date Created\n            </th>\n        </tr>\n    </thead>\n    <tbody>\n        <tr *ngFor=\"let kid of kidList\">\n            <td>\n                {{ kid.id }}\n            </td>\n            <td>\n                <a [routerLink]=\"['/dashboard/activity', kid.id]\">Activity</a> |\n                <a [routerLink]=\"['/dashboard/activity-raw', kid.macId]\">Raw Data</a>\n            </td>\n            <td>\n                {{ kid.name }}\n            </td>\n            <td>\n                {{ kid.macId }}\n            </td>\n            <td>\n                {{ kid.profile }}\n            </td>\n            <td>\n                {{ kid.parent.email }}\n            </td>\n            <td>\n                {{ kid.dateCreated }}\n            </td>\n\n        </tr>\n    </tbody>\n\n</table>\n"

/***/ }),

/***/ 204:
/***/ (function(module, exports) {

module.exports = "\n<form>\n  <legend>Admin Server</legend>\n  <label class=\"error-message label-danger label\" *ngIf=\"showError\">Please enter valid credential</label>\n  <div class=\"form-group\">\n    <md-input-container>\n      <input mdInput placeholder=\"Name\" (input)=\"username = $event.target.value\" required />\n      <md-error>This field is required</md-error>\n    </md-input-container>\n  </div>\n  <div class=\"form-group\">\n    <md-input-container>\n      <input mdInput placeholder=\"Password\" type=\"password\" (input)=\"password = $event.target.value\" required />\n      <md-error>This field is required</md-error>\n    </md-input-container>\n  </div>\n  <button md-raised-button (click)=\"login()\">Login</button>\n</form>\n\n<div class=\"overlay loading\" *ngIf=\"showLoading\">\n  <md-spinner></md-spinner>\n</div>\n\n\n"

/***/ }),

/***/ 205:
/***/ (function(module, exports) {

module.exports = "<h3> User List</h3>\n<div class=\"error\">\n  {{ error }}\n</div>\n<table class=\"table table-striped table-hover\">\n  <thead>\n  <tr>\n    <th>\n      ID\n    </th>\n    <th>\n      Email\n    </th>\n    <th>\n      Name(F L)\n    </th>\n    <th>\n      Profile\n    </th>\n    <th>\n      iOS ID\n    </th>\n    <th>\n      Android ID\n    </th>\n    <th>\n      Sign-up Date\n    </th>\n  </tr>\n  </thead>\n  <tbody>\n  <tr *ngFor=\"let user of userList\">\n    <td>\n      {{ user.id }}\n    </td>\n    <td>\n      {{ user.email }}\n    </td>\n    <td>\n      {{ user.firstName }} {{ user.lastName }}\n    </td>\n    <td>\n      {{ user.profile }}\n    </td>\n    <td>\n      {{ user.ios_registration_id.substring(0, 5) }}\n    </td>\n    <td>\n      {{ user.android_registration_id.substring(0, 5) }}\n    </td>\n    <td>\n      {{ user.dateCreated }}\n    </td>\n  </tr>\n  </tbody>\n\n</table>\n"

/***/ }),

/***/ 259:
/***/ (function(module, exports, __webpack_require__) {

module.exports = __webpack_require__(116);


/***/ }),

/***/ 43:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return AdminToken; });
/**
 * Created by yen-chiehchen on 4/15/17.
 */
var AdminToken = (function () {
    function AdminToken() {
    }
    return AdminToken;
}());

//# sourceMappingURL=admin_login.js.map

/***/ }),

/***/ 75:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__server_service__ = __webpack_require__(12);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2_rxjs_add_operator_switchMap__ = __webpack_require__(63);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2_rxjs_add_operator_switchMap___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_2_rxjs_add_operator_switchMap__);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__angular_router__ = __webpack_require__(21);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4__angular_common__ = __webpack_require__(20);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ActivityRawComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};





var ActivityRawComponent = (function () {
    function ActivityRawComponent(serverService, route, location) {
        this.serverService = serverService;
        this.route = route;
        this.location = location;
    }
    ActivityRawComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.route.params
            .switchMap(function (params) { return _this.serverService.getActivityRawListByKidId(params['macId']); })
            .subscribe(function (list) { return _this.activityRawList = list; });
    };
    ActivityRawComponent.prototype.goBack = function () {
        this.location.back();
    };
    return ActivityRawComponent;
}());
ActivityRawComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-activity-raw',
        template: __webpack_require__(198),
        styles: [__webpack_require__(182)]
    }),
    __metadata("design:paramtypes", [typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */]) === "function" && _a || Object, typeof (_b = typeof __WEBPACK_IMPORTED_MODULE_3__angular_router__["b" /* ActivatedRoute */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_3__angular_router__["b" /* ActivatedRoute */]) === "function" && _b || Object, typeof (_c = typeof __WEBPACK_IMPORTED_MODULE_4__angular_common__["e" /* Location */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_4__angular_common__["e" /* Location */]) === "function" && _c || Object])
], ActivityRawComponent);

var _a, _b, _c;
//# sourceMappingURL=activity-raw.component.js.map

/***/ }),

/***/ 76:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__angular_router__ = __webpack_require__(21);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_common__ = __webpack_require__(20);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__server_service__ = __webpack_require__(12);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4_rxjs_add_operator_switchMap__ = __webpack_require__(63);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4_rxjs_add_operator_switchMap___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_4_rxjs_add_operator_switchMap__);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return ActivityComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};





var ActivityComponent = (function () {
    function ActivityComponent(serverService, route, location) {
        this.serverService = serverService;
        this.route = route;
        this.location = location;
    }
    ActivityComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.route.params
            .switchMap(function (params) { return _this.serverService.getActivityListByKidId(+params['kidId']); })
            .subscribe(function (activityList) { return _this.activityList = activityList; });
    };
    return ActivityComponent;
}());
ActivityComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-activity',
        template: __webpack_require__(199),
        styles: [__webpack_require__(183)]
    }),
    __metadata("design:paramtypes", [typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_3__server_service__["a" /* ServerService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_3__server_service__["a" /* ServerService */]) === "function" && _a || Object, typeof (_b = typeof __WEBPACK_IMPORTED_MODULE_1__angular_router__["b" /* ActivatedRoute */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__angular_router__["b" /* ActivatedRoute */]) === "function" && _b || Object, typeof (_c = typeof __WEBPACK_IMPORTED_MODULE_2__angular_common__["e" /* Location */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_2__angular_common__["e" /* Location */]) === "function" && _c || Object])
], ActivityComponent);

var _a, _b, _c;
//# sourceMappingURL=activity.component.js.map

/***/ }),

/***/ 77:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__model_dashboard__ = __webpack_require__(126);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__server_service__ = __webpack_require__(12);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return DashboardMainComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};



var DashboardMainComponent = (function () {
    function DashboardMainComponent(serverService) {
        this.serverService = serverService;
    }
    DashboardMainComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.dashboard = new __WEBPACK_IMPORTED_MODULE_1__model_dashboard__["a" /* Dashboard */]();
        this.serverService.getDashboardData()
            .then(function (dashboard) {
            console.error(dashboard);
            _this.dashboard = dashboard;
        })
            .catch(function (err) {
            console.error(err);
        });
    };
    return DashboardMainComponent;
}());
DashboardMainComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-dashboard-main',
        template: __webpack_require__(201),
        styles: [__webpack_require__(185)]
    }),
    __metadata("design:paramtypes", [typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_2__server_service__["a" /* ServerService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_2__server_service__["a" /* ServerService */]) === "function" && _a || Object])
], DashboardMainComponent);

var _a;
//# sourceMappingURL=dashboard-main.component.js.map

/***/ }),

/***/ 78:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__server_service__ = __webpack_require__(12);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2__angular_router__ = __webpack_require__(21);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__model_admin_login__ = __webpack_require__(43);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4_ng2_webstorage__ = __webpack_require__(22);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return DashboardComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};






var DashboardComponent = (function () {
    function DashboardComponent(serverService, router, storage) {
        this.serverService = serverService;
        this.router = router;
        this.storage = storage;
    }
    DashboardComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.serverService.tokenValidation().then(function () {
            _this.loadDashboard();
        }).catch(function (err) {
            _this.router.navigate(['/login']);
        });
    };
    DashboardComponent.prototype.loadDashboard = function () {
        this.router.navigate(["dashboard/dashboard-main"]);
    };
    DashboardComponent.prototype.logout = function () {
        this.storage.clear('token');
        this.router.navigate(["/login"]);
    };
    return DashboardComponent;
}());
__decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_4_ng2_webstorage__["b" /* LocalStorage */])(),
    __metadata("design:type", typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_3__model_admin_login__["a" /* AdminToken */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_3__model_admin_login__["a" /* AdminToken */]) === "function" && _a || Object)
], DashboardComponent.prototype, "token", void 0);
DashboardComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-dashboard',
        template: __webpack_require__(202),
        styles: [__webpack_require__(186)]
    }),
    __metadata("design:paramtypes", [typeof (_b = typeof __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */]) === "function" && _b || Object, typeof (_c = typeof __WEBPACK_IMPORTED_MODULE_2__angular_router__["c" /* Router */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_2__angular_router__["c" /* Router */]) === "function" && _c || Object, typeof (_d = typeof __WEBPACK_IMPORTED_MODULE_4_ng2_webstorage__["c" /* LocalStorageService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_4_ng2_webstorage__["c" /* LocalStorageService */]) === "function" && _d || Object])
], DashboardComponent);

var _a, _b, _c, _d;
//# sourceMappingURL=dashboard.component.js.map

/***/ }),

/***/ 79:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__server_service__ = __webpack_require__(12);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return KidListComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var KidListComponent = (function () {
    function KidListComponent(serverService) {
        this.serverService = serverService;
    }
    KidListComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.serverService.getKidList()
            .then(function (kidList) { return _this.kidList = kidList; })
            .catch(function (error) {
            console.log(error);
            _this.error = error;
        });
    };
    return KidListComponent;
}());
KidListComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-kid-list',
        template: __webpack_require__(203),
        styles: [__webpack_require__(187)]
    }),
    __metadata("design:paramtypes", [typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */]) === "function" && _a || Object])
], KidListComponent);

var _a;
//# sourceMappingURL=kid-list.component.js.map

/***/ }),

/***/ 80:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__server_service__ = __webpack_require__(12);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_2_ng2_webstorage__ = __webpack_require__(22);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_3__angular_router__ = __webpack_require__(21);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_4__model_admin_login__ = __webpack_require__(43);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return LoginComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};






var LoginComponent = (function () {
    function LoginComponent(serverService, localSt, router) {
        this.serverService = serverService;
        this.localSt = localSt;
        this.router = router;
    }
    LoginComponent.prototype.ngOnInit = function () {
        this.showLoading = true;
        this.showError = false;
        this.tokenValidation();
    };
    LoginComponent.prototype.login = function () {
        var _this = this;
        this.showLoading = true;
        this.serverService.login(this.username, this.password).then(function (loginToken) {
            _this.showError = false;
            _this.token = loginToken;
            _this.router.navigate(['/dashboard']);
        }).catch(function (err) {
            _this.showLoading = false;
            _this.showError = true;
        });
    };
    LoginComponent.prototype.tokenValidation = function () {
        var _this = this;
        this.serverService.tokenValidation().then(function () {
            _this.router.navigate(['/dashboard']);
        }).catch(function (err) {
            _this.showLoading = false;
            console.log(err);
        });
    };
    return LoginComponent;
}());
__decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_2_ng2_webstorage__["b" /* LocalStorage */])(),
    __metadata("design:type", typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_4__model_admin_login__["a" /* AdminToken */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_4__model_admin_login__["a" /* AdminToken */]) === "function" && _a || Object)
], LoginComponent.prototype, "token", void 0);
LoginComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-login',
        template: __webpack_require__(204),
        styles: [__webpack_require__(188)]
    }),
    __metadata("design:paramtypes", [typeof (_b = typeof __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */]) === "function" && _b || Object, typeof (_c = typeof __WEBPACK_IMPORTED_MODULE_2_ng2_webstorage__["c" /* LocalStorageService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_2_ng2_webstorage__["c" /* LocalStorageService */]) === "function" && _c || Object, typeof (_d = typeof __WEBPACK_IMPORTED_MODULE_3__angular_router__["c" /* Router */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_3__angular_router__["c" /* Router */]) === "function" && _d || Object])
], LoginComponent);

var _a, _b, _c, _d;
//# sourceMappingURL=login.component.js.map

/***/ }),

/***/ 81:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__angular_core__ = __webpack_require__(2);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__server_service__ = __webpack_require__(12);
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return UserListComponent; });
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};


var UserListComponent = (function () {
    function UserListComponent(serverService) {
        this.serverService = serverService;
    }
    UserListComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.serverService.getUserList().then(function (userList) {
            _this.userList = userList;
        });
    };
    return UserListComponent;
}());
UserListComponent = __decorate([
    __webpack_require__.i(__WEBPACK_IMPORTED_MODULE_0__angular_core__["_1" /* Component */])({
        selector: 'app-user-list',
        template: __webpack_require__(205),
        styles: [__webpack_require__(189)]
    }),
    __metadata("design:paramtypes", [typeof (_a = typeof __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */] !== "undefined" && __WEBPACK_IMPORTED_MODULE_1__server_service__["a" /* ServerService */]) === "function" && _a || Object])
], UserListComponent);

var _a;
//# sourceMappingURL=user-list.component.js.map

/***/ }),

/***/ 82:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (binding) */ __webpack_require__.d(__webpack_exports__, "a", function() { return environment; });
/**
 * Created by yen-chiehchen on 4/17/17.
 */
/**
 * Created by yen-chiehchen on 4/17/17.
 */ var environment = {
    production: true,
    BaseURL: "http://dev.childrenlab.com"
};
//# sourceMappingURL=environment.js.map

/***/ })

},[259]);
//# sourceMappingURL=main.e2c79758eb49ad864db2.bundle.js.map