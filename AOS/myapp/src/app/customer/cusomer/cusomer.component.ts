import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AppService } from 'src/app/app.service';
@Component({
  selector: 'app-cusomer',
  templateUrl: './cusomer.component.html',
  styleUrls: ['./cusomer.component.css']
})
export class CusomerComponent implements OnInit {
  ngOnInit(): void {
      this.fetchCart()
  }
  products: any[] = [];  
  checkoutForm:any ={}
  constructor(public router:Router,private appService: AppService){}
  nav(){
    this.router.navigate(['cart']);
  }
  fetchCart(): void {
    let obj = {};
    this.appService.postRequest("listcart", obj).subscribe((result: any) => {
      console.log("result====", result);
      result.forEach((data : any)=> {
         console.log(data)     
      });
      this.products = result;
      // console.log(this.products) 
    });
  }

  calculateSubtotal(): number {
    let subtotal = 0;
    for (const product of this.products) {
      // Assuming there's a price property in each product
      subtotal += (product.totalamount);
    }
    return subtotal;
  }

  // Calculate the total including taxes, shipping, etc.
  calculateTotal(): number {
    // For simplicity, let's assume total is same as subtotal in this example
    return this.calculateSubtotal();
  }

  placeorder(){
    let obj = {};
    this.appService.postRequest("placeorder", obj).subscribe((result: any) => {
      console.log("result====", result);
     
      // console.log(this.products) 
    });

  }
}
