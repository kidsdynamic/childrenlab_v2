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
var UserComponent = (function () {
    function UserComponent(userService) {
        this.userService = userService;
        this.error = "";
        this.users = [];
    }
    UserComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.userService.getUserList()
            .then(function (userList) { return _this.users = userList; })
            .catch(function (error) {
            console.log(error);
            _this.error = error;
        });
    };
    UserComponent = __decorate([
        core_1.Component({
            selector: 'user',
            templateUrl: './template/user.component.html'
        })
    ], UserComponent);
    return UserComponent;
}());
exports.UserComponent = UserComponent;
