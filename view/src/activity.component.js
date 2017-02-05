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
require('rxjs/add/operator/switchMap');
var ActivityComponent = (function () {
    function ActivityComponent(activityService, route, location) {
        this.activityService = activityService;
        this.route = route;
        this.location = location;
        this.error = "";
        this.activities = [];
    }
    ActivityComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.route.params
            .switchMap(function (params) { return _this.activityService.getList(+params['kidId']); })
            .subscribe(function (list) { return _this.activities = list; });
    };
    ActivityComponent.prototype.goBack = function () {
        this.location.back();
    };
    ActivityComponent = __decorate([
        core_1.Component({
            selector: 'activity',
            templateUrl: './template/activity.component.html'
        })
    ], ActivityComponent);
    return ActivityComponent;
}());
exports.ActivityComponent = ActivityComponent;
