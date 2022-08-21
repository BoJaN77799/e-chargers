import { AfterContentChecked, AfterViewInit, Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatTableDataSource } from '@angular/material/table';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { UtilService } from 'src/modules/shared/service/utils-service';
import { VehicleDTO } from 'src/modules/user/model/vehicleDTO';
import { ChargerService } from 'src/modules/user/service/chargerService';
import { VehicleService } from 'src/modules/user/service/vehicleService';
import { ReservationDTO } from '../../model/reservationDTO';
import { TableData } from '../../model/tableData';
import { ReservationService } from '../../service/reservationService';
import { CreateReviewComponent } from 'src/modules/shared/components/create-review/create-review.component';

import * as moment from 'moment'

@Component({
  selector: 'app-user-reservations',
  templateUrl: './user-reservations.component.html',
  styleUrls: ['./user-reservations.component.scss']
})
export class UserReservationsComponent implements OnInit {

  reservations: ReservationDTO[]
  vehicles: VehicleDTO[]
  chargers: ChargerDTO[]

  loggedUsername: string

  displayedColumns: string[] = ['charger_name', 'vehicle_name', 'date_from', 'duration', 'button'];
  dataSource: MatTableDataSource<ReservationDTO>;


  constructor(private reservationService: ReservationService,
    private utilService: UtilService,
    private snackBarService: SnackBarService,
    public dialog: MatDialog) {
    this.reservations = [];
    this.vehicles = [];
    this.chargers = [];
    this.loggedUsername = utilService.getLoggedUsername();
    this.dataSource = new MatTableDataSource(this.reservations)
  }

  ngOnInit(): void {
    this.reservationService.getAllReservationsOfUser(this.utilService.getLoggedUsername()).subscribe(
      (response) => {
        if (response) {
          this.reservations = response.body as ReservationDTO[]
          this.dataSource = new MatTableDataSource(this.reservations)
          console.log(this.reservations);

        }
      },
      (err) => {
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

  cancelReservation(reservation: ReservationDTO) {

    let now = moment((new Date()).getTime());
    let startCharging = moment(reservation.date_from)

    if (now > startCharging) {
      this.snackBarService.openSnackBar("You can't cancel reservation during charing.")
      return
    }

    if (reservation.date_from + reservation.duration * 60 * 1000)

      this.reservationService.cancelReservation(reservation.id).subscribe(
        (response) => {
          if (response)
            this.snackBarService.openSnackBar(response.body as string);
          let index = this.reservations.findIndex((element) => element.id = reservation.id)
          this.reservations.splice(index, 1)
          this.dataSource = new MatTableDataSource(this.reservations)
        },
        (err) => {
          this.snackBarService.openSnackBarFast(err.error)
        }
      )
  }

  checkIfDisabled(reservation: ReservationDTO) {
    let now = moment((new Date()).getTime());
    let endCharging = moment(reservation.date_from + reservation.duration * 60 * 1000)

    if (now < endCharging) {
      return true;
    }

    return false;
  }

  formatDate(date: number) {
    return moment(date).format("YYYY-MM-DD HH:mm")
  }

  openReviewDialog(reservation: ReservationDTO): void {
    const dialogRef = this.dialog.open(CreateReviewComponent, {
      data: reservation,
      width: '30%',
    });

    dialogRef.afterClosed().subscribe(result => {
      // this.reservations.splice(this.reservations.findIndex((element) => element.id === reservation.id), 1)
      console.log(result);
    });
  }


}