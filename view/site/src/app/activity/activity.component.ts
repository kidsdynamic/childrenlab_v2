import {Component, OnInit} from '@angular/core';
import {ActivatedRoute,Params} from '@angular/router';
import {Location} from '@angular/common';
import {ServerService} from "../server.service";
import 'rxjs/add/operator/switchMap';
import {Activity} from "../model/activity";

@Component({
  selector: 'app-activity',
  templateUrl: './activity.component.html',
  styleUrls: ['./activity.component.scss']
})
export class ActivityComponent implements OnInit {

  private error: string;
  private activityList: Activity[];

  constructor(private serverService: ServerService,
              private route: ActivatedRoute,
              private location: Location) {
  }

  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => this.serverService.getActivityListByKidId(+params['kidId']))
      .subscribe(activityList => this.activityList = activityList);
  }

}
