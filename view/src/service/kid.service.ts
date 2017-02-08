/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Injectable } from '@angular/core'
import { Headers, Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'

import { Kid } from '../model/kid'

@Injectable()
export class KidService {

    private kidListUrl = 'https://www.childrenlab.com/v1/admin/kids/list';

    constructor(private http: Http) {}

    getList(): Promise<Kid[]> {
        return this.http.get(this.kidListUrl)
            .toPromise()
            .then(response => response.json() as Kid[])
            .catch(this.handlerError)
    }

    private handlerError(error: any): Promise<any> {
        console.error('Error: ', error);
        return Promise.reject(error.message || error)
    }
}