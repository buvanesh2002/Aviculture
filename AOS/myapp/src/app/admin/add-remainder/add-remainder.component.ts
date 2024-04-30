import { Component } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { AppService } from 'src/app/app.service';
import { AuthserviceService } from 'src/app/authservice.service';

@Component({
  selector: 'app-add-remainder',
  templateUrl: './add-remainder.component.html',
  styleUrls: ['./add-remainder.component.css']
})
export class AddRemainderComponent {
  reminderData: any = {};

  constructor(private router: Router, private toastr: ToastrService, private appservice: AppService,public auth:AuthserviceService) {}
 
  onSubmit() {

    let obj = {
      "emailid" : this.auth.adminemail,
      "reminderId": this.reminderData.reminderId,
      "remindername": this.reminderData.name,
      "beforeDate": this.reminderData.beforeDate.toString(),
      "afterDate": this.reminderData.afterDate.toString(),
      "reminderdate": this.reminderData.date.toString(),
      "remarks": this.reminderData.remarks
    };

   
    console.log(obj);
    
    this.appservice.postRequest("addreminder", obj).subscribe(
      (result: any) => {
        this.showSuccess('Remainder created successfully!');
        this.resetForm();
        this.router.navigate(['listremainder']);
      },
      (error: any) => {
        if (error.error && error.error.errors) {
         
          for (const field in error.error.errors) {
            this.showError(`Error in ${field}: ${error.error.errors[field]}`);
          }
        } else {
          
          this.showError('Error occurred while creating remainder!');
        }
      }
    );
  }

  resetForm() {
    this.reminderData = {};
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
