import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AppService } from 'src/app/app.service';
import { AuthserviceService } from 'src/app/authservice.service';
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
  constructor(public router:Router,private appService: AppService,public auth: AuthserviceService){}
  nav(){
    this.router.navigate(['cart']);
  }
  fetchCart(): void {
    let obj = {"useremailid":this.auth.useremail};
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
    for (let i = 0; i < this.products.length; i++) {
    let obj = {
      "useremailid":this.auth.useremail,
      "emailid": this.products[i].emailid,
      "address": this.checkoutForm.address,
      "companyname": this.checkoutForm.companyname,
      "country": this.checkoutForm.country,
      "emailaddress": this.checkoutForm.emailaddress,
      "firstname": this.checkoutForm.firstname,
      "lastname": this.checkoutForm.lastname,
      "ordernotes": this.checkoutForm.ordernotes,
      "phone": this.checkoutForm.phone,
      "postalcode": this.checkoutForm.postalcode,
      "state": this.checkoutForm.state,
      "breedname":this.products[i].breedName,
      "birdquantity":this.products[i].birdquantity,
      "eggquanity":this.products[i].eggquanity,
      "totalamount":this.products[i].totalamount
    };
    // console.log(this.checkoutForm)
    this.appService.postRequest("placeorder", obj).subscribe((result: any) => {
      console.log("result====", result);
     
      // console.log(this.products) 
    });

  }
}
}
