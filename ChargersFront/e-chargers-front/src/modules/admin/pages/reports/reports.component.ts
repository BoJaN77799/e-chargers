import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import * as moment from 'moment';
import { ReservationDateValidator } from 'src/modules/reservation/validators/ReservationDateValidator';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { Report, ReportItem } from '../../model/reportsDTO';
import { ReportsService } from '../../service/reportsService';

@Component({
  selector: 'app-reports',
  templateUrl: './reports.component.html',
  styleUrls: ['./reports.component.scss']
})
export class ReportsComponent implements OnInit {

  reservationForm: FormGroup;

  chargersReport!: Report;

  reportItems: ReportItem[]

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
    this.reportItems = []
  }

  ngOnInit(): void {
  }

  submit() {
    let startDateTime = moment(this.reservationForm.get("date_from")?.value)
    let endDateTime = moment(this.reservationForm.get("date_to")?.value)

    if (startDateTime > endDateTime) {
      this.snackBarService.openSnackBar("Invalid dates - start is after or equal end")
      return
    }

    this.reportsService.getReport(startDateTime.valueOf(), endDateTime.valueOf()).subscribe(
      (response) => {
        this.chargersReport = response.body as Report;

        const iterrableMap = new Map(Object.entries(this.chargersReport.chargers));
        for (let [key, value] of iterrableMap) {
          this.reportItems.push(value)
        }
        console.log(this.reportItems)
      },
      (err) => {
        this.snackBarService.openSnackBar("There is no report for this period.")
      }
    )

  }

}
