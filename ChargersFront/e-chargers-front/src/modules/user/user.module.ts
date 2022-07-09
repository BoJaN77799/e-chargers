import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UserHomeComponent } from './peges/user-home/user-home.component';
import { RouterModule } from '@angular/router';
import { UserRoutes } from './admin.routes';
import { SharedModule } from '../shared/shared.module';
import { ChargerInfoComponent } from './components/charger-info/charger-info.component';

@NgModule({
  declarations: [
    UserHomeComponent,
    ChargerInfoComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(UserRoutes),
    SharedModule
  ]
})
export class UserModule { }
