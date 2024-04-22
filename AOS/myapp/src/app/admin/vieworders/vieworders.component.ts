import { Component } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthserviceService } from 'src/app/authservice.service';
import { ToastrService } from 'ngx-toastr';

interface Orders {
  address: string;
  companyname: string;
  country: string;
  emailaddress: string;
  firstname: string;
  lastname: string;
  ordernotes: string;
  phone: string;
  postalcode: string;
  state: string;
  breedname: string;
  birdquantity: string;
  eggquantity: string;
  totalamount: string;
}

@Component({
  selector: 'app-vieworders',
  templateUrl: './vieworders.component.html',
  styleUrls: ['./vieworders.component.css']
})
export class ViewordersComponent {
  responseData: Orders[] = [];


  constructor(private appService: AppService,private toastr: ToastrService, private route: ActivatedRoute, private router: Router,public auth:AuthserviceService) {}

  ngOnInit() {
    this.listData()
  }


  listData() {
    let obj = {};
    this.responseData= []
    this.appService.postRequest("listorder", obj).subscribe(
      (response) => {
        console.log('List data received:', response);
        if (response && response.length) {
          console.log(response)
          console.log(response.id);
          this.responseData.push(...response); 
          
        }
      },
      (error) => {
        console.error('Failed to fetch list data:', error);
        // Handle error response here
      }
    );
  }

  



}

