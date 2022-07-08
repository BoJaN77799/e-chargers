import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { UserHomeComponent } from './peges/user-home/user-home.component';
import { RouterModule } from '@angular/router';
import { UserRoutes } from './admin.routes';

@NgModule({
  declarations: [
    UserHomeComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(UserRoutes)
  ]
})
export class UserModule { }
