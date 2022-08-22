import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CreateReservationComponent } from './components/create-reservation/create-reservation.component';
import { ReactiveFormsModule } from '@angular/forms';
import { SharedModule } from '../shared/shared.module';
import { UserReservationsComponent } from './pages/user-reservations/user-reservations.component';



@NgModule({
  declarations: [
    CreateReservationComponent,
    UserReservationsComponent
  ],
  imports: [
    SharedModule,
    ReactiveFormsModule,
    CommonModule,
  ],
  exports: [
    UserReservationsComponent
  ]
})
export class ReservationModule { }
