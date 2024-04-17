import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AppService } from 'src/app/app.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-cart',
  templateUrl: './cart.component.html',
  styleUrls: ['./cart.component.css']
})
export class CartComponent implements OnInit {
  products: any[] = [];
  id:any = null;
  count:number=0
  birdquant:number=1;
  eggquant:number=1;
  constructor(private appService: AppService, private route: ActivatedRoute, private router: Router,private toastr: ToastrService) {}
  ngOnInit() {
    this.route.params.subscribe(params => {
      this.id = params['id']; 
     console.log("fetch call=",this.id)// Access the 'id' route parameter
     if (this.id != null) {
      this.fetchFlockData(this.id);

     }else{
      this.fetchCart()
     }
   });
  
  }


  fetchFlockData(id: string): void {
    this.appService.postRequest("cartlist", { id: id }).subscribe((res: any) => {
      this.count++
      console.log(this.count)
      this.fetchCart();
    });
  }



  removeProduct(id: any) {
    this.appService.postRequest("removecatrid", { id: id }).subscribe((result: any) => {
      this.fetchCart()
      this.toastr.success(result);
      console.log("result====", result);
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
   

  eggminus(productId: number) {
    const index = this.products.findIndex(product => product.id === productId);
    if (index !== -1) {
      if (this.products[index].eggquantity > 0) {
        this.products[index].eggquantity--;
        this.products[index].totalamount -= (this.products[index].eggprice)
        this.products
        let obj = {
          id: productId, // Replace yourIdValue with the actual value of the id
          eggquantity: this.products[index].eggquantity, // Replace yourEggQuantityValue with the actual value of the eggquantity
          totalamount : this.products[index].totalamount
        };
        
        this.appService.postRequest("eggquantity", obj).subscribe((result: any) => {
          console.log("result====", result);
        });
        
      } else {
        console.error('Quantity cannot be less than 0');
      }
    } else {
      console.error('Product not found');
    }
  }

  eggplus(productId: number) {
    const index = this.products.findIndex(product => product.id === productId);
    if (index !== -1) {
      this.products[index].eggquantity++;
      this.products[index].totalamount += (this.products[index].eggprice)
      if ( this.products[index].eggquantity === this.products[index].noEgg){
        this.toastr.success("This is the Maximum EggQuantity Present For this Breed")
      }
      let obj = {
        id: productId, // Replace yourIdValue with the actual value of the id
        eggquantity: this.products[index].eggquantity,// Replace yourEggQuantityValue with the actual value of the eggquantity
        totalamount : this.products[index].totalamount
      };
      
      this.appService.postRequest("eggquantity", obj).subscribe((result: any) => {
        console.log("result====", result);
      });
    } else {
      console.error('Product not found');
    }
  }

  birdminus(productId: number) {
    const index = this.products.findIndex(product => product.id === productId);
    if (index !== -1) {
      if (this.products[index].birdquantity > 0) {
        this.products[index].birdquantity--;
        this.products[index].totalamount -= (this.products[index].birdprice)
        let obj = {
          id: productId, // Replace yourIdValue with the actual value of the id
          birdquantity: this.products[index].birdquantity, // Replace yourEggQuantityValue with the actual value of the eggquantity
          totalamount : this.products[index].totalamount
        };
        
        this.appService.postRequest("birdquantity", obj).subscribe((result: any) => {
          console.log("result====", result);
        });
      } else {
        console.error('Quantity cannot be less than 0');
      }
    } else {
      console.error('Product not found');
    }
  }

  birdplus(productId: number) {
    const index = this.products.findIndex(product => product.id === productId);
    if (index !== -1) {

      this.products[index].birdquantity++;
      this.products[index].totalamount += (this.products[index].birdprice)
      if ( this.products[index].birdquantity == this.products[index].noBirds){
 
        this.toastr.success("This is the Maximum BirdQuantity Present For this Breed")
      }
   
      let obj = {
        id: productId, // Replace yourIdValue with the actual value of the id
        birdquantity: this.products[index].birdquantity, // Replace yourEggQuantityValue with the actual value of the eggquantity
        totalamount : this.products[index].totalamount
      };
      
      this.appService.postRequest("birdquantity", obj).subscribe((result: any) => {
        console.log("result====", result);
      });
    } else {
      console.error('Product not found');
    }
  }


}


