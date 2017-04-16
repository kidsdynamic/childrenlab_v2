import { Component, OnInit } from '@angular/core';
import {User} from "../model/user";
import {ServerService} from "../server.service";

@Component({
  selector: 'app-user-list',
  templateUrl: './user-list.component.html',
  styleUrls: ['./user-list.component.scss']
})
export class UserListComponent implements OnInit {

  error: string;
  userList: User[];

  constructor(
    private serverService: ServerService
  ) { }

  ngOnInit() {
    this.serverService.getUserList().then(userList => {
      this.userList = userList;
    });

  }

}
