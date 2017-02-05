"use strict";
var platform_browser_dynamic_1 = require('@angular/platform-browser-dynamic');
var main_module_1 = require('./src/main.module');
var core_1 = require('@angular/core');
require('./src/css/styles.css');
// Enable production mode unless running locally
if (!/localhost/.test(document.location.host)) {
    core_1.enableProdMode();
}
platform_browser_dynamic_1.platformBrowserDynamic().bootstrapModule(main_module_1.MainModule);
/*
 Copyright 2016 Google Inc. All Rights Reserved.
 Use of this source code is governed by an MIT-style license that
 can be found in the LICENSE file at http://angular.io/license
 */ 
