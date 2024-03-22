import { Component } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-home-cus',
  templateUrl: './home-cus.component.html',
  styleUrls: ['./home-cus.component.css']
})
export class HomeCusComponent {
   
  constructor(private appservice: AppService, public router: Router,private route:ActivatedRoute) {}

  tocart(){
     
  }
}
