import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { FormsModule } from '@angular/forms';
import { ToastrModule } from 'ngx-toastr';
import { HttpClientModule } from '@angular/common/http';
import{ BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { DataComponent } from './data/data.component';
import { FlockEntryComponent } from './flock-entry/flock-entry.component';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ListFlockComponent } from './list-flock/list-flock.component';
import { MatTableModule } from '@angular/material/table';
import { UpadteFlockComponent } from './upadte-flock/upadte-flock.component';
import { FlockentrydataComponent } from './flockentrydata/flockentrydata.component';

import { DailyentryComponent } from './dailyentry/dailyentry.component';
import { ListflockentryComponent } from './listflockentry/listflockentry.component';
// import { AppService } from './app.service'; 
import { MatMenuModule } from '@angular/material/menu';
import { MatDividerModule } from '@angular/material/divider';
import { FirstpageComponent } from './firstpage/firstpage.component';
import { AddRemainderComponent } from './add-remainder/add-remainder.component';
import { ListRemainderComponent } from './list-remainder/list-remainder.component';

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
    ListRemainderComponent
    
    
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
