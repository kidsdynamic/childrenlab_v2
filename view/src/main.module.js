"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
/**
 * Created by yen-chiehchen on 2/4/17.
 */
var core_1 = require('@angular/core');
var platform_browser_1 = require('@angular/platform-browser');
var router_1 = require('@angular/router');
var http_1 = require('@angular/http');
var dashboard_component_1 = require('./dashboard.component');
var user_component_1 = require('./user.component');
var kid_component_1 = require('./kid.component');
var activity_component_1 = require('./activity.component');
var user_service_1 = require('./service/user.service');
var kid_service_1 = require('./service/kid.service');
var activity_service_1 = require('./service/activity.service');
var routing_module_1 = require('./routing.module');
var MainModule = (function () {
    function MainModule() {
    }
    MainModule = __decorate([
        core_1.NgModule({
            imports: [
                platform_browser_1.BrowserModule,
                router_1.RouterModule,
                http_1.HttpModule,
                routing_module_1.Routing
            ],
            declarations: [
                dashboard_component_1.DashboardComponent,
                user_component_1.UserComponent,
                kid_component_1.KidComponent,
                activity_component_1.ActivityComponent
            ],
            providers: [user_service_1.UserService, kid_service_1.KidService, activity_service_1.ActivityService],
            bootstrap: [dashboard_component_1.DashboardComponent]
        })
    ], MainModule);
    return MainModule;
}());
exports.MainModule = MainModule;
