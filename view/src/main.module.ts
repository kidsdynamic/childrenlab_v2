/**
 * Created by yen-chiehchen on 2/4/17.
 */
import { NgModule } from '@angular/core'
import { BrowserModule } from '@angular/platform-browser'
import { RouterModule } from '@angular/router'
import { HttpModule } from '@angular/http'

import { DashboardComponent } from './dashboard.component'
import { UserComponent } from './user.component'
import { KidComponent } from './kid.component'
import { ActivityComponent } from './activity.component'
import { ActivityRawComponent } from './activity_raw.component'

import { UserService } from './service/user.service'
import { KidService } from './service/kid.service'
import { ActivityService } from './service/activity.service'
import { APP_CONFIG, AppConfig } from './constant/app.config'

import { Routing } from './routing.module'



@NgModule({
    imports: [
        BrowserModule,
        RouterModule,
        HttpModule,
        Routing
    ],
    declarations: [
        DashboardComponent,
        UserComponent,
        KidComponent,
        ActivityComponent,
        ActivityRawComponent
    ],
    providers: [
        UserService,
        KidService,
        ActivityService,
        {
            provide: APP_CONFIG,
            useValue: AppConfig
        }

    ],
    bootstrap: [ DashboardComponent ]
})

export class MainModule{}