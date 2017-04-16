import { Component, OnInit } from '@angular/core';
import {Kid} from "../model/kid";
import {ServerService} from "../server.service";

@Component({
  selector: 'app-kid-list',
  templateUrl: './kid-list.component.html',
  styleUrls: ['./kid-list.component.scss']
})
export class KidListComponent implements OnInit {

  private error: string;
  private kidList: Kid[];

  constructor(
    private serverService: ServerService
  ) { }

  ngOnInit() {
    this.serverService.getKidList()
      .then(kidList => this.kidList = kidList)
      .catch(error => {
        console.log(error);
        this.error = error;
      })
  }

}
