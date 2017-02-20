/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Component, OnInit } from '@angular/core'
import { ActivatedRoute, Params } from '@angular/router'
import { Location } from '@angular/common'

import { ActivityRaw } from './model/activity_raw'
import { ActivityService } from './service/activity.service'

import 'rxjs/add/operator/switchMap';

@Component({
    selector: 'activity_raw',
    templateUrl: './template/activity_raw.component.html'
})


export class ActivityRawComponent implements OnInit{
    error: string = "";
    activities: ActivityRaw[] = [];
    constructor(
        private activityService: ActivityService,
        private route: ActivatedRoute,
        private location: Location
    ){}

    ngOnInit(): void {
        this.route.params
            .switchMap((params: Params) => this.activityService.getRawList(params['macId']))
            .subscribe(list => this.activities = list )

    }

    goBack(): void {
        this.location.back();
    }
}