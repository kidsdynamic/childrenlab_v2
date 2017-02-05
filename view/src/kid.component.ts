/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Component, OnInit } from '@angular/core'

import { Kid } from './model/Kid'
import { KidService } from './service/kid.service'

@Component({
    selector: 'device',
    templateUrl: './template/kid.component.html'
})


export class KidComponent implements OnInit{
    error: string = "";
    kids: Kid[] = [];
    constructor(private kidService: KidService){}

    ngOnInit(): void {
        this.kidService.getList()
            .then(list => this.kids = list)
            .catch(error => {
                console.log(error);
                this.error = error;
            })
    }
}