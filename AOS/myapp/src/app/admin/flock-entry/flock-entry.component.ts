import { Component } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { AppService } from 'src/app/app.service';
import { AuthserviceService } from 'src/app/authservice.service';

@Component({
  selector: 'app-flock-entry',
  templateUrl: './flock-entry.component.html',
  styleUrls: ['./flock-entry.component.css']
})
export class FlockEntryComponent {
  flockData: any = {};
  
 
  constructor(private router: Router, private toastr: ToastrService, private appservice: AppService,public auth:AuthserviceService) {}

  onSubmit() {
    this.imageToBase64("birdImage")
    .then(base64String => {
        console.log("Base64 image string:", base64String);
        // Proceed with form submission here
        let obj = {
          "emailid" : this.auth.adminemail,
            "flockName": this.flockData.flockName,
            "breedName": this.flockData.breedName,
            "startDate": this.flockData.startDate,
            "startAge": parseInt(this.flockData.startAge),
            "openingBirds": parseInt(this.flockData.openingBirds),
            "shedNumber": this.flockData.shedNumber.toString(),
            "image": base64String // Add base64 image string to the object
        };

        console.log(obj);
        console.log("Returned");

        this.appservice.postRequest("addflock", obj).subscribe(
            (result: any) => {
                this.showSuccess('Flock created successfully!');
                this.resetForm();
                this.router.navigate(['listflock']);
            },
            (error: any) => {
                if (error.error && error.error.errors) {
                    for (const field in error.error.errors) {
                        this.showError(`Error in ${field}: ${error.error.errors[field]}`);
                    }
                } else {
                    this.showError('Error occurred while creating flock!');
                }
            }
        );
    })
    .catch(error => {
        console.error("Error converting image to base64:", error);
    });
}

  resetForm() {
    this.flockData = {};
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

  imageToBase64(elementId: string): Promise<string> {
    console.log("Image converter");
    const fileInputElement = document.getElementById(elementId) as HTMLInputElement;
    if (!fileInputElement || !fileInputElement.files || !fileInputElement.files[0]) {
        console.log("File input or file not found");
        return Promise.reject(new Error(`File input or file not found for element with ID "${elementId}"`));
    }

    const file = fileInputElement.files[0];
    const reader = new FileReader();

    return new Promise((resolve, reject) => {
        reader.onload = () => {
            const base64String = reader.result as string;
            console.log("Base64 image string:", base64String);
            resolve(base64String);
        };

        reader.onerror = (error) => {
            console.error("Error loading image:", error);
            reject(error);
        };

        reader.readAsDataURL(file);
    });
}
}
