import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthserviceService } from '../../authservice.service';

@Component({
  selector: 'app-firstpage',
  templateUrl: './firstpage.component.html',
  styleUrls: ['./firstpage.component.css']
})
export class FirstpageComponent {
  constructor(private router:Router,public auth:AuthserviceService){}
  
  ngOnInit() {
    this.auth.canAccess(true);
  }
  
   
  

}
