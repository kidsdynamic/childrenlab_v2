import { Component, OnInit } from '@angular/core';
import {ServerService} from "../server.service";
import {LocalStorageService} from 'ng2-webstorage';
import {LocalStorage} from 'ng2-webstorage';
import { Router } from '@angular/router';
import {AdminToken} from "../model/admin_login";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  username: string;
  password: string;
  showError: boolean;
  showLoading: boolean;
  @LocalStorage() public token: AdminToken;

  constructor(
    private serverService: ServerService,
    private localSt: LocalStorageService,
    private router: Router) {}

  ngOnInit() {
    this.showLoading = true;
    this.showError = false;
    this.tokenValidation();
  }

  login() {
    this.showLoading = true;
    this.serverService.login(this.username, this.password).then(loginToken => {
      this.showError = false;
      this.token = loginToken;
      this.router.navigate(['/dashboard']);
    }).catch(err => {
      this.showLoading = false;
      this.showError = true;
    });
  }

  tokenValidation() {
    this.serverService.tokenValidation().then( () => {
      this.router.navigate(['/dashboard']);
    }).catch(err => {
      this.showLoading = false;
      console.log(err);
    })
  }
}
