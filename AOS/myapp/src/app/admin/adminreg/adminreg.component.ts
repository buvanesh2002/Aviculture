import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { AppService } from 'src/app/app.service';

@Component({
  selector: 'app-adminreg',
  templateUrl: './adminreg.component.html',
  styleUrls: ['./adminreg.component.css']
})
export class AdminregComponent {
  adminData:any ={};
  constructor(private router: Router, private toastr: ToastrService, private appservice: AppService) {}
  onSubmit() {
    let obj = {
      "name": this.adminData.name,
      "email": this.adminData.email,
      "password": this.adminData.password,
      "phone": this.adminData.phone.toString(),
      "address": this.adminData.address,
      "pincode": this.adminData.pincode.toString()
    };

    console.log(obj);
    
    this.appservice.postRequest("adminreg", this.adminData).subscribe(
      (result: any) => {
        this.showSuccess('admin created successfully!');
        this.resetForm();
        this.router.navigate(['login']);
      },
      (error: any) => {
        if (error.error && error.error.errors) {
         
          for (const field in error.error.errors) {
            this.showError(`Error in ${field}: ${error.error.errors[field]}`);
          }
        } else {
          
          this.showError('Email already exists');
        }
      }
    );
  }

  resetForm() {
    this.adminData = {};
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