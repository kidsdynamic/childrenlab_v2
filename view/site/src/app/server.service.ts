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
  DELETE_MAC_ID='/admin/deleteMacID';

  @LocalStorage() public token: AdminToken;

  constructor(private http: Http) { }

  login(userName: string, password: string): Promise<AdminToken> {
    const loginJson = {
      name: userName,
      password: password
    };

    return this.http.post(environment.BaseURL + this.LOGIN_API, JSON.stringify(loginJson))
      .toPromise()
      .then(response => response.json() as AdminToken);

  }

  getUserList(): Promise<User[]> {
    const options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.USER_LIST_API}`, options)
      .toPromise()
      .then(response => response.json() as User[])
      .catch(this.handleError);
  }

  getKidList(): Promise<Kid[]> {
    const options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.KID_LIST_API}`, options)
      .toPromise()
      .then(response => response.json() as Kid[])
      .catch(this.handleError);
  }

  getActivityListByKidId(kidId: number): Promise<Activity[]> {
    const options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.ACTIVITY_LIST_API}/${kidId}`, options)
      .toPromise()
      .then(response => response.json() as Activity[])
      .catch(this.handleError)
  }

  getActivityRawListByKidId(kidId: number): Promise<ActivityRaw[]> {
    const options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.ACTIVITY_RAW_LIST_API}/${kidId}`, options)
      .toPromise()
      .then(response => response.json() as ActivityRaw[])
      .catch(this.handleError)
  }

  getDashboardData(): Promise<Dashboard> {
    const options = this.addTokenHeader();
    return this.http.get(`${environment.BaseURL + this.DASHBOARD_API}`, options)
      .toPromise()
      .then(response => response.json() as Dashboard)
      .catch(this.handleError)
  }

  tokenValidation(): Promise<any> {
    if (this.token == null) {
      return Promise.reject('Invalid token');
    }
    return this.http.get(environment.BaseURL + this.TOKEN_VALIDATION_API + '?email=' +
      this.token.username + '&token=' + this.token.access_token)
      .toPromise()
      .then();

  }

  deleteMacId(macId:string): Promise<any> {
    const options = this.addTokenHeader();
    return this.http.delete(`${environment.BaseURL + this.DELETE_MAC_ID}?macId=${macId}`, options)
      .toPromise()
      .then(response => response.json())
      .catch(this.handleError)
  }

  private addTokenHeader(): RequestOptions{
    const headers = new Headers({
      'x-auth-token': this.token.access_token,
      'Content-Type': 'application/json'
    });
    const options = new RequestOptions({ headers: headers });
    return options;
  }

  private handleError(error: any): Promise<any> {
    console.error('Error: ', error);
    return Promise.reject(error.message || error);
  }
  


}
