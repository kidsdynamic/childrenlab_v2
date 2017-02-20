/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Injectable, Inject } from '@angular/core'
import { Headers, Http} from '@angular/http'
import 'rxjs/add/operator/toPromise'
import { APP_CONFIG, IAppConfig } from '../constant/app.config'

import { Kid } from '../model/kid'

@Injectable()
export class KidService {

    private kidListUrl = 'admin/kids/list';

    constructor(private http: Http, @Inject(APP_CONFIG) private config: IAppConfig ) {}

    getList(): Promise<Kid[]> {
        return this.http.get(`${this.config.apiEndpoint + this.kidListUrl}`)
            .toPromise()
            .then(response => response.json() as Kid[])
            .catch(this.handlerError)
    }

    private handlerError(error: any): Promise<any> {
        console.error('Error: ', error);
        return Promise.reject(error.message || error)
    }
}