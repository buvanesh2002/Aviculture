import { Component } from '@angular/core';
import { AppService } from 'src/app/app.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-shop',
  templateUrl: './shop.component.html',
  styleUrls: ['./shop.component.css']
})
export class ShopComponent {
  responseData: any[] = []; // Modify type of responseData to match your response data structure
  
  constructor(private appservice: AppService, public router: Router, private route: ActivatedRoute) { }

  ngOnInit() {
    this.listData();
  }

  listData() {
    let obj = {};
    this.appservice.postRequest("shoplist", obj).subscribe(
      (response: any[]) => { // Specify type of response as any[]
        console.log('List data received:', response);
        this.responseData = response; // Assign response data to responseData
      },
      (error) => {
        console.error('Failed to fetch list data:', error);
        // Handle error response here
      }
    );
 
  }

  tocart(id:any){
    console.log("in shop=",id);
    this.router.navigate(['cart', JSON.stringify(id)]);

  }
}

   // this.responseData = [
    //   {
    //     BreedName: 'Chair A',
    //     Nobirds: 10,
    //     NoEgg: 20,
    //     Birdprice: 50,
    //     EggPrice: 100
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   {
    //     BreedName: 'Chair B',
    //     Nobirds: 15,
    //     NoEgg: 25,
    //     Birdprice: 60,
    //     EggPrice: 110
    //   },
    //   // Add more dummy data as needed
    // ];