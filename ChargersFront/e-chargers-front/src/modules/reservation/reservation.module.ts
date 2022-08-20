import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateReservationComponent } from './components/create-reservation/create-reservation.component';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { SharedModule } from '../shared/shared.module';



@NgModule({
  declarations: [
    CreateReservationComponent
  ],
  imports: [
    SharedModule,
    ReactiveFormsModule,
    CommonModule
  ]
})
export class ReservationModule { }
