import { Injectable } from '@angular/core';
import {Http,Headers,RequestOptions} from '@angular/http';
import {AdminToken} from "./model/admin_login";
import {environment} from "../environments/environment";
import 'rxjs/add/operator/toPromise';
import {LocalStorage} from 'ng2-webstorage';
import {User} from "./model/user";
import {Kid} from "./model/kid";
import {Activity} from "./model/activity";
import {ActivityRaw} from "./model/activity-raw";
import {Dashboard} from "./model/dashboard";

@Injectable()
export class ServerService {

  LOGIN_API = "/admin/login";
  TOKEN_VALIDATION_API = '/v1/user/isTokenValid';
  USER_LIST_API = '/admin/userList';
  KID_LIST_API = '/admin/kidList';
  ACTIVITY_LIST_API = '/admin/activityList';
  ACTIVITY_RAW_LIST_API = '/admin/activityRawList';
  DASHBOARD_API='/admin/dashboard';

  @LocalStorage() public token: AdminToken;

  constructor(private http: Http) { }

  login(userName: string, password: string): Promise<AdminToken> {
    let loginJson = {
      name: userName,
      password: password
    };

    return this.http.post(environment.BaseURL + this.LOGIN_API, JSON.stringify(loginJson))
      .toPromise()
      .then(response => response.json() as AdminToken);

  }

  getUserList(): Promise<User[]> {
    let options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.USER_LIST_API}`, options)
      .toPromise()
      .then(response => response.json() as User[])
      .catch(this.handleError);
  }

  getKidList(): Promise<Kid[]> {
    let options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.KID_LIST_API}`, options)
      .toPromise()
      .then(response => response.json() as Kid[])
      .catch(this.handleError);
  }

  getActivityListByKidId(kidId: number): Promise<Activity[]> {
    let options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.ACTIVITY_LIST_API}/${kidId}`, options)
      .toPromise()
      .then(response => response.json() as Activity[])
      .catch(this.handleError)
  }

  getActivityRawListByKidId(kidId: number): Promise<ActivityRaw[]> {
    let options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.ACTIVITY_RAW_LIST_API}/${kidId}`, options)
      .toPromise()
      .then(response => response.json() as ActivityRaw[])
      .catch(this.handleError)
  }

  getDashboardData(): Promise<Dashboard> {
    let options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.DASHBOARD_API}`, options)
      .toPromise()
      .then(response => response.json() as Dashboard)
      .catch(this.handleError)
  }

  tokenValidation(): Promise<any> {
    if(this.token == null) {
      return Promise.reject("Invalid token");
    }
    return this.http.get(environment.BaseURL + this.TOKEN_VALIDATION_API + "?email=" + this.token.username + "&token=" + this.token.access_token)
      .toPromise()
      .then();

  }

  private addTokenHeader(): RequestOptions{
    let headers = new Headers({
      'x-auth-token': this.token.access_token,
      'Content-Type': 'application/json'
    });
    let options = new RequestOptions({ headers: headers });
    return options;
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error);
  }


}
