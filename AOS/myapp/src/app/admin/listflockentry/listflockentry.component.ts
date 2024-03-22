import { Component, OnInit } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';

interface Item {
 
    entrydate: string;
    age: number; // Assuming it should be a number
    openingbirds: number;
    mortality: number;
    birdssold: number;
    closingbirds: number;
    cummortality: number;
    mortalityPercent: number; // Assuming it should be a percentage
    eggsperDay: number;
    eggproducion: number;
    productionPercent: number; // Assuming it should be a percentage
    feed: number;
    feedperBird: number;
    feedperEgg: number;
    cumFPE: number;

}

@Component({
  selector: 'app-listflockentry',
  templateUrl: './listflockentry.component.html',
  styleUrls: ['./listflockentry.component.css']
})
 
export class ListflockentryComponent implements OnInit {

  responseData: Item[] = [];
  id : string ="";

  constructor(private appservice: AppService, public router: Router,private route:ActivatedRoute) { }

 

  ngOnInit() {
    this.route.params.subscribe(params => {
       this.id = params['id']; 
      console.log("fetch call=",this.id)
      this.listData(this.id);
    });
  }


  listData(id:string) {
   
    this.responseData = [];
    this.appservice.postRequest("listparticularflock", { id: id }).subscribe(
      (response) => {
        console.log('List data received:', response);
        if (response && response.length) {
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
