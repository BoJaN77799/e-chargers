import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SnackBarService } from './service/snack-bar.service';
import { UtilService } from './service/utils-service';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { Interceptor } from './interceptors/interceptor.interceptor';
import { MapPageComponent } from './pages/map-page/map-page.component';
import { MaterialExampleModule } from 'src/material.module';

@NgModule({
  declarations: [


    MapPageComponent
  ],
  imports: [
    CommonModule,
    MaterialExampleModule
  ],
  exports: [
    MapPageComponent,
    MaterialExampleModule
  ],
  providers: [
    SnackBarService,
    UtilService,
    { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
  ]
})
export class SharedModule { }
