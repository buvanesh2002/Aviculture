import { Component } from '@angular/core';
import { Router } from '@angular/router';
@Component({
  selector: 'app-cusomer',
  templateUrl: './cusomer.component.html',
  styleUrls: ['./cusomer.component.css']
})
export class CusomerComponent {
   
  constructor(public router:Router){}
  nav(){
    this.router.navigate(['cart']);
  }
}
