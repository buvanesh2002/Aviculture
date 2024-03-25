import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { FormsModule } from '@angular/forms';
import { ToastrModule } from 'ngx-toastr';
import { HttpClientModule } from '@angular/common/http';
import{ BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DataComponent } from './admin/data/data.component';
import { FlockEntryComponent } from './admin/flock-entry/flock-entry.component';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ListFlockComponent } from './admin/list-flock/list-flock.component';
import { MatTableModule } from '@angular/material/table';
import { UpadteFlockComponent } from './admin/upadte-flock/upadte-flock.component';
import { FlockentrydataComponent } from './admin/flockentrydata/flockentrydata.component';

import { DailyentryComponent } from './admin/dailyentry/dailyentry.component';
import { ListflockentryComponent } from './admin/listflockentry/listflockentry.component';
// import { AppService } from './app.service'; 
import { MatMenuModule } from '@angular/material/menu';
import { MatDividerModule } from '@angular/material/divider';
import { FirstpageComponent } from './admin/firstpage/firstpage.component';
import { AddRemainderComponent } from './admin/add-remainder/add-remainder.component';
import { ListRemainderComponent } from './admin/list-remainder/list-remainder.component';
import { CusomerComponent } from './customer/cusomer/cusomer.component';
import { CartComponent } from './customer/furni-1.0.0/cart/cart.component';
import { ThankyouComponent } from './customer/thankyou/thankyou.component';
import { ShopComponent } from './customer/shop/shop.component';
import { HomeCusComponent } from './customer/home-cus/home-cus.component';
import { AboutComponent } from './customer/about/about.component';
import { ContactComponent } from './customer/contact/contact.component';
import { BlogComponent } from './customer/blog/blog.component';
import { ServicesComponent } from './customer/services/services.component';
import { UserloginComponent } from './userlogin/userlogin.component';
import { UserfirstpageComponent } from './userfirstpage/userfirstpage.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    DataComponent,
    FlockEntryComponent,
    ListFlockComponent,
    UpadteFlockComponent,
    FlockentrydataComponent,
    DailyentryComponent,
    ListflockentryComponent,
    FirstpageComponent,
    AddRemainderComponent,
    ListRemainderComponent,
    CusomerComponent,
    CartComponent,
    ThankyouComponent,
    ShopComponent,
    HomeCusComponent,
    AboutComponent,
    ContactComponent,
    BlogComponent,
    ServicesComponent,
    UserloginComponent,
    UserfirstpageComponent
    
    
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ToastrModule.forRoot(),
    MatInputModule,
    MatFormFieldModule,
    MatTableModule,
    MatMenuModule,
    MatDividerModule
  ],
  providers: [AppComponent],
  bootstrap: [AppComponent]
})
export class AppModule { }
