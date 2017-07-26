import { Component, OnInit } from '@angular/core';
import {BatteryStatus} from "app/model/kid";
import {LocalStorage} from "ng2-webstorage/dist/decorators";
import {AdminToken} from "../model/admin_login";
import {ServerService} from '../server.service';
import {ActivatedRoute, Params} from "@angular/router";

@Component({
  selector: 'app-kid-battery',
  templateUrl: './kid-battery.component.html',
  styleUrls: ['./kid-battery.component.scss']
})
export class KidBatteryComponent implements OnInit {

  error: string;
  batteryStatus: BatteryStatus[];
  macId: string;

  @LocalStorage() public token: AdminToken;
  constructor(private serverService: ServerService, private route: ActivatedRoute) { }

  ngOnInit() {
    this.route.params
      .switchMap((params: Params) => this.serverService.getBatteryStatus(params['macId']))
      .subscribe((batteryStatus) => this.batteryStatus = batteryStatus);
  }

}
