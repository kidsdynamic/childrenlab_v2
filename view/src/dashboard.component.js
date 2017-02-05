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
var DashboardComponent = (function () {
    function DashboardComponent() {
    }
    DashboardComponent = __decorate([
        core_1.Component({
            selector: 'app',
            template: "\n    <nav class=\"navbar navbar-default\"> \n        <div class=\"container-fluid\"> \n            <a routerLink=\"/user\" class=\"btn btn-default navbar-btn\">User</a>\n            <a routerLink=\"/device\" class=\"btn btn-default navbar-btn\">Device-Kid</a>\n        </div>\n    </nav>\n    \n    <router-outlet></router-outlet>\n    "
        })
    ], DashboardComponent);
    return DashboardComponent;
}());
exports.DashboardComponent = DashboardComponent;
