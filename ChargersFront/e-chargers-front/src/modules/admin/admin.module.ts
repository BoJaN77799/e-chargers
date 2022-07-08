import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AdminHomeComponent } from './pages/admin-home/admin-home.component';
import { AdminRoutes } from './admin.routes';
import { RouterModule } from '@angular/router';

@NgModule({
  declarations: [
    AdminHomeComponent
  ],
  imports: [
    CommonModule,
    RouterModule.forChild(AdminRoutes),
  ]
})
export class AdminModule { }
