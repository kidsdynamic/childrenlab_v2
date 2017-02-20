/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Injectable, Inject } from '@angular/core'
import { Headers, Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'
import { APP_CONFIG, IAppConfig } from '../constant/app.config'

import { User } from '../model/user'

@Injectable()
export class UserService {

    private userListUrl = 'user/userList';

    constructor(private http: Http, @Inject(APP_CONFIG) private config: IAppConfig ) {}

    getUserList(): Promise<User[]> {
        return this.http.get(`${this.config.apiEndpoint + this.userListUrl}`)
            .toPromise()
            .then(response => response.json() as User[])
            .catch(this.handlerError)
    }

    private handlerError(error: any): Promise<any> {
        console.error('Error: ', error);
        return Promise.reject(error.message || error)
    }
}