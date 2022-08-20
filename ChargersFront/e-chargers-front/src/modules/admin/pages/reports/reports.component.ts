import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import * as moment from 'moment';
import { ReservationDateValidator } from 'src/modules/reservation/validators/ReservationDateValidator';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { Report } from '../../model/reportsDTO';
import { ReportsService } from '../../service/reportsService';

@Component({
  selector: 'app-reports',
  templateUrl: './reports.component.html',
  styleUrls: ['./reports.component.scss']
})
export class ReportsComponent implements OnInit {

  reservationForm: FormGroup;

  chargersReport!: Report;

  constructor(
    private fb: FormBuilder,
    private snackBarService: SnackBarService,
    private reportsService: ReportsService) {
    let startDateTime = moment((new Date()).getTime() - 7776000000).format('YYYY-MM-DDTHH:mm')
    let endDateTime = moment((new Date()).getTime() + 7776000000).format('YYYY-MM-DDTHH:mm')

    this.reservationForm = this.fb.group({
      date_from: [startDateTime, [Validators.required]],
      date_to: [endDateTime, [Validators.required]],
    });
  }

  ngOnInit(): void {
  }

  submit() {
    let startDateTime = moment(this.reservationForm.get("date_from")?.value)
    let endDateTime = moment(this.reservationForm.get("date_from")?.value)

    if (startDateTime > endDateTime) {
      this.snackBarService.openSnackBar("Invalid dates - start is after or equal end")
      return
    }

    this.reportsService.getReport(startDateTime.unix(), endDateTime.unix()).subscribe(
      (response) => {
        this.chargersReport = response.body as Report;
      },
      (err) => {
        this.snackBarService.openSnackBar("There is no report for this period.")
      }
    )

  }

}
