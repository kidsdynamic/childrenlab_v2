/**
 * Created by yen-chiehchen on 2/4/17.
 */

import {Injectable, Inject} from '@angular/core'
import { Headers, Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'

import { Activity } from '../model/activity'
import { ActivityRaw } from '../model/activity_raw'
import { APP_CONFIG, IAppConfig } from '../constant/app.config'

@Injectable()
export class ActivityService {

    private activityListUrl = 'activity/list';
    private activityRawListUrl = 'admin/activity/raw';

    constructor(private http: Http, @Inject(APP_CONFIG) private config: IAppConfig ) {}

    getList(kidId: number): Promise<Activity[]> {
        return this.http.get(`${this.config.apiEndpoint + this.activityListUrl}/${kidId}`)
            .toPromise()
            .then(response => response.json() as Activity[])
            .catch(this.handlerError)
    }

    getRawList(macId: string): Promise<ActivityRaw[]> {
        return this.http.get(`${this.config.apiEndpoint + this.activityRawListUrl}/${macId}`)
            .toPromise()
            .then(response => response.json() as ActivityRaw[])
            .catch(this.handlerError)
    }

    private handlerError(error: any): Promise<any> {
        console.error('Error: ', error);
        return Promise.reject(error.message || error)
    }
}