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
var KidService = (function () {
    function KidService(http) {
        this.http = http;
        this.kidListUrl = 'https://www.childrenlab.com/v1/kids/list';
    }
    KidService.prototype.getList = function () {
        return this.http.get(this.kidListUrl)
            .toPromise()
            .then(function (response) { return response.json(); })
            .catch(this.handlerError);
    };
    KidService.prototype.handlerError = function (error) {
        console.error('Error: ', error);
        return Promise.reject(error.message || error);
    };
    KidService = __decorate([
        core_1.Injectable()
    ], KidService);
    return KidService;
}());
exports.KidService = KidService;
