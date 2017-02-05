import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { MainModule } from './src/main.module';
import { enableProdMode } from '@angular/core';
import './src/css/styles.css'

// Enable production mode unless running locally
if (!/localhost/.test(document.location.host)) {
    enableProdMode();
}

platformBrowserDynamic().bootstrapModule(MainModule);


/*
 Copyright 2016 Google Inc. All Rights Reserved.
 Use of this source code is governed by an MIT-style license that
 can be found in the LICENSE file at http://angular.io/license
 */