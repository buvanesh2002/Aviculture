import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { DataComponent } from './admin/data/data.component';
import { FlockEntryComponent } from './admin/flock-entry/flock-entry.component';
import { ListFlockComponent } from './admin/list-flock/list-flock.component';
import { UpadteFlockComponent } from './admin/upadte-flock/upadte-flock.component';
import { FlockentrydataComponent } from './admin/flockentrydata/flockentrydata.component';
import {DailyentryComponent} from './admin/dailyentry/dailyentry.component';
import { ListflockentryComponent } from './admin/listflockentry/listflockentry.component';
import { FirstpageComponent } from './admin/firstpage/firstpage.component';
import { AddRemainderComponent } from './admin/add-remainder/add-remainder.component';
import { ListRemainderComponent } from './admin/list-remainder/list-remainder.component';
import { CartComponent } from './customer/furni-1.0.0/cart/cart.component';
import { CusomerComponent } from './customer/cusomer/cusomer.component';
import { AboutComponent } from './customer/about/about.component';
import { HomeCusComponent } from './customer/home-cus/home-cus.component';
import { ContactComponent } from './customer/contact/contact.component';
import { BlogComponent } from './customer/blog/blog.component';
import { ServicesComponent } from './customer/services/services.component';
import { ShopComponent } from './customer/shop/shop.component';
import { ThankyouComponent } from './customer/thankyou/thankyou.component';
import { UserloginComponent } from './userlogin/userlogin.component';
import { UserfirstpageComponent } from './userfirstpage/userfirstpage.component';
import { CustomerregComponent } from './customer/customerreg/customerreg.component';
import { AdminregComponent } from './admin/adminreg/adminreg.component';
import { ViewordersComponent } from './admin/vieworders/vieworders.component';



const routes: Routes = [
  { path: '', redirectTo: '/login', pathMatch: 'full' },
  { path: 'login', component: LoginComponent },
  { path: 'firstpage',component:FirstpageComponent },
  {path: 'data', component: DataComponent},
  { path: 'addflock', component: FlockEntryComponent},
  {path  : 'listflock', component:ListFlockComponent},
  {path  : 'updateflock/:id', component:UpadteFlockComponent},
  {path : 'flockentrydata/:id', component: FlockentrydataComponent},
  {path:'dailyentry', component:DailyentryComponent },
  {path :'listflockentry/:id' , component:ListflockentryComponent},
  {path:'addremainder',component:AddRemainderComponent},
  {path:'listremainder' , component:ListRemainderComponent},
  {path:'cusomer' , component:CusomerComponent},
  { path: 'cart/:id', component: CartComponent },
  { path: 'home-cus' ,component:HomeCusComponent}, 
  { path: 'about' ,component:AboutComponent},
  { path: 'contact', component:ContactComponent},
  { path: 'blog' ,component:BlogComponent},
  { path:'services' ,component:ServicesComponent},
  {path: 'shop' , component:ShopComponent},
  {path:'thankyou', component:ThankyouComponent},
  {path:'userlogin', component:UserloginComponent},
  {path:'userfirstpage', component:UserfirstpageComponent},
  { path: 'cart', component: CartComponent },
  { path: 'customerreg', component: CustomerregComponent },
  { path: 'adminreg', component: AdminregComponent },
  { path: 'vieworders', component: ViewordersComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }

