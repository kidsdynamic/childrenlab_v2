/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Component, OnInit } from '@angular/core'
import { ActivatedRoute, Params } from '@angular/router'
import { Location } from '@angular/common'

import { Activity } from './model/activity'
import { ActivityService } from './service/activity.service'

import 'rxjs/add/operator/switchMap';

@Component({
    selector: 'activity',
    templateUrl: './template/activity.component.html'
})


export class ActivityComponent implements OnInit{
    error: string = "";
    activities: Activity[] = [];
    constructor(
        private activityService: ActivityService,
        private route: ActivatedRoute,
        private location: Location
    ){}

    ngOnInit(): void {
        this.route.params
            .switchMap((params: Params) => this.activityService.getList(+params['kidId']))
            .subscribe(list => this.activities = list )

    }

    goBack(): void {
            this.location.back();
    }
}