import { LiveAnnouncer } from '@angular/cdk/a11y';
import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { UtilService } from 'src/modules/shared/service/utils-service';
import { VehicleDTO } from 'src/modules/user/model/vehicleDTO';
import { UserHomeComponent } from 'src/modules/user/pages/user-home/user-home.component';
import { VehicleService } from 'src/modules/user/service/vehicleService';
import { ReservationService } from '../../service/reservationService';
import * as moment from 'moment';
import { ReservationDateValidator } from '../../validators/ReservationDateValidator';
import { ReservationDTO } from '../../model/reservationDTO';

@Component({
  selector: 'app-create-reservation',
  templateUrl: './create-reservation.component.html',
  styleUrls: ['./create-reservation.component.scss']
})
export class CreateReservationComponent implements OnInit {

  vehicles: VehicleDTO[];

  reservationForm: FormGroup;

  constructor(public dialog: MatDialog,
    public dialogRef: MatDialogRef<UserHomeComponent>,
    @Inject(MAT_DIALOG_DATA) public charger: ChargerDTO,
    private _liveAnnouncer: LiveAnnouncer,
    private reservationService: ReservationService,
    private snackBarService: SnackBarService,
    private vehicleService: VehicleService,
    private utilService: UtilService,
    private fb: FormBuilder,
  ) {
    this.vehicles = [];
    let startDateTime = moment((new Date()).getTime() + 5460000).format('YYYY-MM-DDTHH:mm')
    this.reservationForm = this.fb.group({
      date_from: [startDateTime, [Validators.required, ReservationDateValidator]],
      duration: [15, [Validators.required, Validators.min(15), Validators.max(90)]],
      vehicle: [null, [Validators.required]]
    });
  }


  ngOnInit(): void {
    let username = this.utilService.getLoggedUsername();
    this.vehicleService.getAllVehiclesOfUser(username).subscribe(
      (response) => {
        if (response)
          this.vehicles = response.body as VehicleDTO[]
      },
      (err) => {
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

  submit() {

    let reservationDTO: ReservationDTO = {
      "id": 0,
      "username": this.utilService.getLoggedUsername(),
      "charger_id": this.charger.id,
      "charger_name": this.charger.name,
      "vehicle_id": this.reservationForm.get('vehicle')?.value,
      "vehicle_name": "",
      "date_from": moment(this.reservationForm.get('date_from')?.value).valueOf(),
      "duration": this.reservationForm.get('duration')?.value,
    }
    console.log(reservationDTO)

    this.reservationService.createReservation(reservationDTO).subscribe(
      (response) => {
        if (response.body)
          this.snackBarService.openSnackBar(response.body)
        this.dialogRef.close();
      },
      (err) => {
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

}
