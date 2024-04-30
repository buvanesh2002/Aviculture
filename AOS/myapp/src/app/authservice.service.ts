import { Injectable } from '@angular/core';
import { Router } from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthserviceService {

  constructor(private router:Router){}
    isAuth():boolean{
     if( "fi"){
       return true;
     }
     return false;
    }
    canAccess( val:boolean){
      if(!this.isAuth()){
        this.router.navigate(['login']);
      }
     
    }
    isloggedin:boolean = false
    adminemail:string = ''
    useremail:string = ''
    islogged ():boolean{
      return this.isloggedin 
    }
    dologin(){
      this.isloggedin = true
    }
    adminlogin(test:string){
        this.adminemail=test;
    }

    logout(){
      this.isloggedin = false
      this.router.navigate(['userlogin']);
    }
    adminlogout(){
      this.router.navigate(['login']);
    }

    userlogin(test:string){
      this.useremail = test
    }
}
