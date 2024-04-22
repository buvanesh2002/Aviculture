import { Component, OnInit } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthserviceService } from 'src/app/authservice.service';


interface Item {
  id: string;
  flockName:string;
  breedName: string;
  startDate: string;
  age: string;
  openingBirds: string;
  shedNumber: string;
  active: string;
}

@Component({
  selector: 'app-list-flock',
  templateUrl: './list-flock.component.html',
  styleUrls: ['./list-flock.component.css']
})

export class ListFlockComponent implements OnInit {
  
 responseData: Item[] = [];


 constructor(private appservice: AppService, private route: ActivatedRoute, private router: Router,public auth:AuthserviceService) {}
  ngOnInit() {
    this.listData()
  }


  listData() {
    let obj = {};
    this.responseData= []
    this.appservice.postRequest("listflock", obj).subscribe(
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
  getname( id:string){
    console.log(id);
    this.router.navigate(['updateflock',id])
}
id( id:string){
  console.log(id);
  this.router.navigate(['listflockentry',id])
}
}
