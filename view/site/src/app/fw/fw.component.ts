import { Component, OnInit } from '@angular/core';
import {ServerService} from '../server.service';
import {FWVersion} from '../model/fw-version';
import {environment} from 'environments/environment';

@Component({
  selector: 'app-fw',
  templateUrl: './fw.component.html',
  styleUrls: ['./fw.component.scss']
})
export class FWComponent implements OnInit {

  uploading = false;
  hasFile = false;
  file: File;
  fwList: FWVersion[];
  env = environment;

  constructor(private serverService: ServerService) { }

  ngOnInit() {

    const fileInput = <HTMLInputElement>document.getElementById('file_input_file');
  const fileInputText = <HTMLInputElement>document.getElementById('file_input_text');
    const uploadButton = <HTMLButtonElement>document.getElementById('upload_button');
    const submitButton = <HTMLButtonElement>document.getElementById('submit_button');
    const versionInput = <HTMLButtonElement>document.getElementById('fw_version_name');
    const fileInputContainer = document.getElementsByClassName('file_input_container')[0];

    uploadButton.addEventListener('click', () => {
      fileInput.click();
    });

    submitButton.addEventListener('click', () => {
      this.uploading = true;
      this.file = fileInput.files[0];
      this.serverService.uploadFWFile(versionInput.value, this.file)
        .then(() => {
          this.updateFWList();
        });
    });

    fileInput.addEventListener('change', changeInputText);

    function changeInputText() {
      this.hasFile = true;
      const str = fileInput.value;
      let i;
      if (str.lastIndexOf('\\')) {
        i = str.lastIndexOf('\\') + 1;
      } else if (str.lastIndexOf('/')) {
        i = str.lastIndexOf('/') + 1;
      }
      fileInputText.value = str.slice(i, str.length);
      fileInputText.classList.remove('none');
      fileInputContainer.classList.add('with-text');
      versionInput.focus();
    }
    this.updateFWList();
  }

  updateFWList() {
    this.serverService.getFWList()
      .then(fwList => {
        console.error(fwList);
        this.fwList = fwList;
      })
      .catch(err => {
        console.error(err);
      });
  }

}
