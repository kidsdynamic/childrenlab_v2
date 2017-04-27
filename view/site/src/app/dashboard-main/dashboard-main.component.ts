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
  options: Object;
  constructor(private serverService: ServerService) { }

  ngOnInit() {
    this.dashboard = new Dashboard();

    this.serverService.getDashboardData()
      .then(dashboard => {
        this.dashboard = dashboard;
      })
      .catch(err => {
        console.error(err);
      });
  }

  setupChart(){
    let data = [];
    for(let s of this.dashboard.signup) {
      data.push(s.signupCount);
    }
    console.error(data);
    this.options = {
      title: { text: 'Sign up'},
      series: [
        data
      ]
    }
  }

}
