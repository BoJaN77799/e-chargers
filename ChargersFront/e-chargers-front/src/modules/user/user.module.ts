import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UserHomeComponent } from './peges/user-home/user-home.component';
import { RouterModule } from '@angular/router';
import { UserRoutes } from './admin.routes';
import { SharedModule } from '../shared/shared.module';
import { ChargerInfoComponent } from './components/charger-info/charger-info.component';
import { SearchChargersComponent } from './components/search-chargers/search-chargers.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MaterialExampleModule } from 'src/material.module';

@NgModule({
  declarations: [
    UserHomeComponent,
    ChargerInfoComponent,
    SearchChargersComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(UserRoutes),
    SharedModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class UserModule { }
