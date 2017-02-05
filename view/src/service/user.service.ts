/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Injectable } from '@angular/core'
import { Headers, Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'

import { User } from '../model/user'

@Injectable()
export class UserService {

    private userListUrl = 'https://www.childrenlab.com/v1/user/userList';

    constructor(private http: Http) {}

    getUserList(): Promise<User[]> {
        return this.http.get(this.userListUrl)
            .toPromise()
            .then(response => response.json() as User[])
            .catch(this.handlerError)
    }

    private handlerError(error: any): Promise<any> {
        console.error('Error: ', error);
        return Promise.reject(error.message || error)
    }
}