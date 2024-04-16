import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { AppService } from 'src/app/app.service';

@Component({
  selector: 'app-customerreg',
  templateUrl: './customerreg.component.html',
  styleUrls: ['./customerreg.component.css']
})
export class CustomerregComponent {
  userData: any = {}; // Object to store user registration data

  constructor(private router: Router, private toastr: ToastrService, private appservice: AppService) { }

  onSubmit() {
    let obj = {
      "name": this.userData.name,
      "email": this.userData.email,
      "password": this.userData.password,
      "phone": this.userData.phone.toString(),
      "address": this.userData.address,
      "pincode": this.userData.pincode.toString()
    };
    console.log('User registration data:', obj);
      
    this.appservice.postRequest("customerreg", obj).subscribe(
      (result: any) => {
        this.showSuccess('Account created successfully!');
        this.reset();
        this.router.navigate(['login']);
        // this.router.navigate(['listremainder']);
      },
      (error: any) => {
        if (error.error && error.error.errors) {
         
          for (const field in error.error.errors) {
            this.showError(`Error in ${field}: ${error.error.errors[field]}`);
          }
        } else {
          
          this.showError('Email already registered');
        }
      }
    );
    

    // You can send the registration data to a server or perform any other necessary actions
  }
 reset() {
  this.userData ={};
 }
 showSuccess(msg: any) {
  if (msg) {
    this.toastr.success(msg);
  }
}

showError(msg: any) {
  if (msg) {
    this.toastr.error(msg);
  }
}
}