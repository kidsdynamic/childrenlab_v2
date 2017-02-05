/**
 * Created by yen-chiehchen on 2/4/17.
 */

import { Component, OnInit } from '@angular/core'

import { User } from './model/user'
import { UserService } from './service/user.service'

@Component({
    selector: 'user',
    templateUrl: './template/user.component.html'
})


export class UserComponent implements OnInit{
    error: string = "";
    users: User[] = [];
    constructor(private userService: UserService){}

    ngOnInit(): void {
        this.userService.getUserList()
            .then(userList => this.users = userList)
            .catch(error => {
                console.log(error);
                this.error = error;
            })
    }
}