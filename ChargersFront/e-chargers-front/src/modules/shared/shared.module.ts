import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SnackBarService } from './service/snack-bar.service';
import { UtilService } from './service/utils-service';
import { MatSnackBarModule } from '@angular/material/snack-bar';
// import { PaginationComponent } from './components/pagination/pagination.component';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
// import { Interceptor } from './interceptors/interceptor.interceptor';
// import { SharedDatePickerService } from './services/shared-data-picker.service';
// import { DateFormatPipe } from './pipes/date-format.pipe';
import { MatTableModule } from '@angular/material/table';
import { Interceptor } from './interceptors/interceptor.interceptor';
import { MapPageComponent } from './pages/map-page/map-page.component';

@NgModule({
  declarations: [


    MapPageComponent
  ],
  imports: [
    CommonModule,
    MatSnackBarModule,
    MatTableModule
  ],
  exports: [
    MapPageComponent
  ],
  providers: [
    SnackBarService,
    UtilService,
    { provide: HTTP_INTERCEPTORS, useClass: Interceptor, multi: true },
  ]
})
export class SharedModule { }
