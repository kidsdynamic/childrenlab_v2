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
require('rxjs/add/operator/toPromise');
var ActivityService = (function () {
    function ActivityService(http) {
        this.http = http;
        this.activityListUrl = 'https://www.childrenlab.com/v1/activity/list';
    }
    ActivityService.prototype.getList = function (kidId) {
        return this.http.get(this.activityListUrl + "/" + kidId)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handlerError);
    };
    ActivityService.prototype.handlerError = function (error) {
        console.error('Error: ', error);
        return Promise.reject(error.message || error);
    };
    ActivityService = __decorate([
        core_1.Injectable()
    ], ActivityService);
    return ActivityService;
}());
exports.ActivityService = ActivityService;
