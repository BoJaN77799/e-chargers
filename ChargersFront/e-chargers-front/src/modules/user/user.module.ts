import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UserHomeComponent } from './pages/user-home/user-home.component';
import { RouterModule } from '@angular/router';
import { UserRoutes } from './user.routes';
import { SharedModule } from '../shared/shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { ReservationModule } from '../reservation/reservation.module';

@NgModule({
  declarations: [
    UserHomeComponent,
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(UserRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule,
    ReservationModule
  ]
})
export class UserModule { }
