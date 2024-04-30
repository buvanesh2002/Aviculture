import { Component } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthserviceService } from 'src/app/authservice.service';

interface Item {
  reminderId: string;
  remindername: string;
  beforedate: string;
  afterdate: string;
  reminderdate: string;
  remarks: string;
  status: string;
}


@Component({
  selector: 'app-list-remainder',
  templateUrl: './list-remainder.component.html',
  styleUrls: ['./list-remainder.component.css']
})
export class ListRemainderComponent {
  
  responseData: Item[] = [];


  constructor(private appservice: AppService, private route: ActivatedRoute, private router: Router,public auth:AuthserviceService) {}

  ngOnInit() {
    this.listData()
  }


  listData() {
    let obj = {
      "emailid":this.auth.adminemail
    };
    this.responseData= []
    this.appservice.postRequest("listremainder", obj).subscribe(
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

  id( id:string){
    console.log(id);
    this.router.navigate(['listflockentry',id])
  }



}
