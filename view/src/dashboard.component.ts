/**
 * Created by yen-chiehchen on 2/4/17.
 */

import {Component} from '@angular/core'

@Component({
    selector: 'app',
    template: `
    <nav class="navbar navbar-default"> 
        <div class="container-fluid"> 
            <a routerLink="/user" class="btn btn-default navbar-btn">User</a>
            <a routerLink="/device" class="btn btn-default navbar-btn">Device-Kid</a>
        </div>
    </nav>
    
    <router-outlet></router-outlet>
    `
})

export class DashboardComponent {
}