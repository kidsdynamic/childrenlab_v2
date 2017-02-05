/**
 * Created by yen-chiehchen on 2/4/17.
 */
"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var core_1 = require('@angular/core');
var router_1 = require('@angular/router');
var user_component_1 = require('./user.component');
var dashboard_component_1 = require('./dashboard.component');
var kid_component_1 = require('./kid.component');
var activity_component_1 = require('./activity.component');
var routes = [
    { path: 'dashboard', component: dashboard_component_1.DashboardComponent },
    { path: 'user', component: user_component_1.UserComponent },
    { path: 'device', component: kid_component_1.KidComponent },
    { path: 'activity/:kidId', component: activity_component_1.ActivityComponent }
];
var Routing = (function () {
    function Routing() {
    }
    Routing = __decorate([
        core_1.NgModule({
            imports: [router_1.RouterModule.forRoot(routes, { useHash: true })],
            exports: [router_1.RouterModule]
        })
    ], Routing);
    return Routing;
}());
exports.Routing = Routing;
