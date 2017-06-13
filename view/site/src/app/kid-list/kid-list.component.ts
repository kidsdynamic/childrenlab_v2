import { Component, OnInit } from '@angular/core';
import {Kid} from '../model/kid';
import {ServerService} from '../server.service';
import { LocalStorage } from 'ng2-webstorage';
import { AdminToken } from 'app/model/admin_login';
import * as swal from 'sweetalert';

@Component({
  selector: 'app-kid-list',
  templateUrl: './kid-list.component.html',
  styleUrls: ['./kid-list.component.scss']
})
export class KidListComponent implements OnInit {

  error: string;
  kidList: Kid[];

  @LocalStorage() public token: AdminToken;

  constructor(
    private serverService: ServerService
  ) { }

  ngOnInit() {
    this.serverService.getKidList()
      .then(kidList => this.kidList = kidList)
      .catch(error => {
        console.log(error);
        this.error = error;
      });
  }

  delete(macId:string, id:number) {
    swal({
        title: "Are you sure?",
        text: `The Mac ID( ${macId} ) will be deleted forever!`,
        type: "warning",
        showCancelButton: true,
        confirmButtonColor: "#AEDEF4",
        confirmButtonText: "Yes, delete it!",
        cancelButtonText: "No!",
        showLoaderOnConfirm: true,
        closeOnConfirm: false,
        closeOnCancel: true
      },
      (isConfirm) => {
        if (isConfirm) {
          this.deleteKidFromServer(macId, id);
          
        }
      });
  }

  deleteKidFromServer(macId:string, id:number) {
    this.serverService.deleteMacId(macId)
      .then(() => {
        this.kidList.splice(id, 1);
        swal("Deleted!", "The MAC ID has been deleted.", "success");
      })
      .catch(error => {
        swal("Deleted!", "Can't delete the Mac ID, please contact admin user.", "error");
      })
  }

}
