import { Component, OnInit } from '@angular/core';
import {ServerService} from "../server.service";
import {Router} from '@angular/router';
import {AdminToken} from "../model/admin_login";
import {LocalStorage} from 'ng2-webstorage';
import {LocalStorageService} from 'ng2-webstorage';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  @LocalStorage() public token: AdminToken;

  constructor(
    private serverService: ServerService,
    private router: Router,
    private storage: LocalStorageService
  ) { }

  ngOnInit() {
    this.serverService.tokenValidation().then().catch(err => {
      this.router.navigate(['/login']);
    })
  }

  logout() {
    this.storage.clear('token');
    this.router.navigate(["/login"]);
  }
}
