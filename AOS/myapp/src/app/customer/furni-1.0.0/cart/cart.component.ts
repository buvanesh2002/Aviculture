import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AppService } from 'src/app/app.service';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent implements OnInit {
  products: any[] = [];
  id:any = null;
  count:number=0
  constructor(private appService: AppService, private route: ActivatedRoute, private router: Router) {}

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.id = params['id']; 
     console.log("fetch call=",this.id)// Access the 'id' route parameter
     this.fetchFlockData(this.id);
   });
  }

  fetchFlockData(id: string): void {
    this.appService.postRequest("cartlist", { id: id }).subscribe((result: any) => {
      this.count++
      console.log(this.count)
      console.log("result====", result);
      this.products = result;
    });
  }



  removeProduct(id: any) {
    this.appService.postRequest("removecatrid", { id: id }).subscribe((result: any) => {
      console.log("result====", result);
    });
  }

  calculateSubtotal(): number {
    let subtotal = 0;
    for (const product of this.products) {
      // Assuming there's a price property in each product
      subtotal += (product.price * product.quantity);
    }
    return subtotal;
  }

  // Calculate the total including taxes, shipping, etc.
  calculateTotal(): number {
    // For simplicity, let's assume total is same as subtotal in this example
    return this.calculateSubtotal();
  }
}
