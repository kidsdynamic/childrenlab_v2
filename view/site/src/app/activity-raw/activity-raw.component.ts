import { Component, OnInit } from '@angular/core';
import {ServerService} from "../server.service";
import 'rxjs/add/operator/switchMap';
import { ActivatedRoute, Params } from '@angular/router'
import { Location } from '@angular/common'
import {ActivityRaw} from "../model/activity-raw";

@Component({
  selector: 'app-activity-raw',
  templateUrl: './activity-raw.component.html',
  styleUrls: ['./activity-raw.component.scss']
})
export class ActivityRawComponent implements OnInit {

  error: string;
  activityRawList: ActivityRaw[];

  constructor(
    private serverService: ServerService,
    private route: ActivatedRoute,
    private location: Location
  ) { }

  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => this.serverService.getActivityRawListByKidId(params['macId']))
      .subscribe(list => this.activityRawList = list )
  }

  goBack(): void {
    this.location.back();
  }

}
