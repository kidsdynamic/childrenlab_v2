import { Component, OnInit } from '@angular/core';
import {Dashboard} from "../model/dashboard";
import {ServerService} from "../server.service";

@Component({
  selector: 'app-dashboard-main',
  templateUrl: './dashboard-main.component.html',
  styleUrls: ['./dashboard-main.component.scss']
})
export class DashboardMainComponent implements OnInit {

  error: string;
  dashboard: Dashboard;
  constructor(private serverService: ServerService) { }

  ngOnInit() {
    this.dashboard = new Dashboard();

    this.serverService.getDashboardData()
      .then(dashboard => {
        console.error(dashboard);
        this.dashboard = dashboard;
      })
      .catch(err => {
        console.error(err);
      });
  }

}
