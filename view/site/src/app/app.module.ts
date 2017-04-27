import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { AppComponent } from './app.component';
import {MdButtonModule, MdInputModule, MdToolbarModule, MdProgressSpinnerModule, MdGridListModule} from '@angular/material';
import {ServerService} from "./server.service";
import {Ng2Webstorage} from 'ng2-webstorage';
import { DashboardComponent } from './dashboard/dashboard.component';
import {RouteModule} from './route/route.module';
import { LoginComponent } from './login/login.component';
import { UserListComponent } from './user-list/user-list.component';
import { KidListComponent } from './kid-list/kid-list.component';
import { ActivityComponent } from './activity/activity.component';
import { ActivityRawComponent } from './activity-raw/activity-raw.component';
import { DashboardMainComponent } from './dashboard-main/dashboard-main.component';
import { ChartModule } from 'angular2-highcharts';
import { HighchartsStatic } from 'angular2-highcharts/dist/HighchartsService';

declare var require: any;
export function highchartsFactory() {
  return require('highcharts');
}

@NgModule({
  declarations: [
    AppComponent,
    DashboardComponent,
    LoginComponent,
    UserListComponent,
    KidListComponent,
    ActivityComponent,
    ActivityRawComponent,
    DashboardMainComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    RouteModule,
    BrowserAnimationsModule,
    MdButtonModule,
    MdInputModule,
    MdToolbarModule,
    Ng2Webstorage,
    MdProgressSpinnerModule,
    MdGridListModule,
    ChartModule
  ],
  providers: [
    ServerService,
    {
      provide: HighchartsStatic,
      useFactory: highchartsFactory
    }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
