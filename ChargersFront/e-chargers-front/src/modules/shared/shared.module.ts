import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SnackBarService } from './service/snack-bar.service';
import { UtilService } from './service/utils-service';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import { Interceptor } from './interceptors/interceptor.interceptor';
import { MapPageComponent } from './pages/map-page/map-page.component';
import { MaterialExampleModule } from 'src/material.module';
import { ChargerInfoComponent } from './components/charger-info/charger-info.component';
import { SearchChargersComponent } from './components/search-chargers/search-chargers.component';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    ChargerInfoComponent,
    SearchChargersComponent,
    MapPageComponent
  ],
  imports: [
    CommonModule,
    MaterialExampleModule,
    ReactiveFormsModule
  ],
  exports: [
    MapPageComponent,
    MaterialExampleModule,
    SearchChargersComponent,
    ChargerInfoComponent
  ],
  providers: [
    SnackBarService,
    UtilService,
    { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
  ]
})
export class SharedModule { }
