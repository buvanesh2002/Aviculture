import { Component } from '@angular/core';

@Component({
  selector: 'app-customerreg',
  templateUrl: './customerreg.component.html',
  styleUrls: ['./customerreg.component.css']
})
export class CustomerregComponent {
  userData: any = {}; // Object to store user registration data

  constructor() { }

  onSubmit() {
    // Handle user registration submission here
    console.log('User registration data:', this.userData);
    // You can send the registration data to a server or perform any other necessary actions
  }
}