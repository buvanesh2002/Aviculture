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
}
