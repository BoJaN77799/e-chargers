import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AdminHomeComponent } from './pages/admin-home/admin-home.component';
import { AdminRoutes } from './admin.routes';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../shared/shared.module';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CreateChargerComponent } from './pages/create-charger/create-charger.component';

@NgModule({
  declarations: [
    AdminHomeComponent,
    CreateChargerComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(AdminRoutes),
    SharedModule,
    ReactiveFormsModule,
    FormsModule
  ]
})
export class AdminModule { }
