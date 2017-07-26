import { NgModule }             from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import {DashboardComponent} from "../dashboard/dashboard.component";
import {LoginComponent} from "../login/login.component";
import {UserListComponent} from "../user-list/user-list.component";
import {KidListComponent} from "../kid-list/kid-list.component";
import { ActivityComponent } from "../activity/activity.component";
import { ActivityRawComponent } from "app/activity-raw/activity-raw.component";
import { FWComponent } from '../fw/fw.component';
import {DashboardMainComponent} from "../dashboard-main/dashboard-main.component";
import {KidBatteryComponent} from "app/kid-battery/kid-battery.component";

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', component: LoginComponent },
  {
    path: 'dashboard',
    component: DashboardComponent,
    children: [
      {
        path: 'userList',
        component: UserListComponent
      },
      {
        path: 'kidList',
        component: KidListComponent
      },
      {
        path: 'kidBattery/:macId',
        component: KidBatteryComponent
      },
      {
        path: 'activity/:kidId',
        component: ActivityComponent
      },
      {
        path: 'activity-raw/:macId',
        component: ActivityRawComponent
      },
      {
        path: 'dashboard-main',
        component: DashboardMainComponent
      },
      {
        path: 'fw-management',
        component: FWComponent
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {useHash: true})],
  exports: [RouterModule]
})
export class RouteModule { }
