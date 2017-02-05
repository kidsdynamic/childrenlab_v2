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
var KidComponent = (function () {
    function KidComponent(kidService) {
        this.kidService = kidService;
        this.error = "";
        this.kids = [];
    }
    KidComponent.prototype.ngOnInit = function () {
        var _this = this;
        this.kidService.getList()
            .then(function (list) { return _this.kids = list; })
            .catch(function (error) {
            console.log(error);
            _this.error = error;
        });
    };
    KidComponent = __decorate([
        core_1.Component({
            moduleId: module.id,
            selector: 'device',
            templateUrl: './template/kid.component.html'
        })
    ], KidComponent);
    return KidComponent;
}());
exports.KidComponent = KidComponent;
