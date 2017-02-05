/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Injectable } from '@angular/core'
import { Headers, Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'

import { Activity } from '../model/activity'

@Injectable()
export class ActivityService {

    private activityListUrl = 'https://www.childrenlab.com/v1/activity/list';

    constructor(private http: Http) {}

    getList(kidId: number): Promise<Activity[]> {
        return this.http.get(`${this.activityListUrl}/${kidId}`)
            .toPromise()
            .then(response => response.json() as Activity[])
            .catch(this.handlerError)
    }

    private handlerError(error: any): Promise<any> {
        console.error('Error: ', error);
        return Promise.reject(error.message || error)
    }
}