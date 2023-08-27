import { AfterViewInit, Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { CreateReservationComponent } from 'src/modules/reservation/components/create-reservation/create-reservation.component';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { RecensionDTO, RecensionWithUserDTO } from '../../model/reviewDTO';
import { ReviewService } from '../../service/reviewService';
import { UtilService } from '../../service/utils-service';
import { SnackBarService } from '../../service/snack-bar.service';

@Component({
  selector: 'app-charger-info',
  templateUrl: './charger-info.component.html',
  styleUrls: ['./charger-info.component.scss']
})
export class ChargerInfoComponent {

  @Input()
  charger: ChargerDTO

  reviewsVisible = false
  recensions: RecensionWithUserDTO[];

  constructor(
    public dialog: MatDialog,
    private utilService: UtilService,
    private reviewssService: ReviewService,
    private snackBarService: SnackBarService) {
    this.recensions = [];
    this.charger = {} as ChargerDTO
  }

  openReservationDialog(): void {
    const dialogRef = this.dialog.open(CreateReservationComponent, {
      data: this.charger,
      width: '30%',
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
    });
  }

  isRole(role: string): boolean {
    return this.utilService.isRole(role);
  }

  loadRecensions() {
    if (this.charger)
      this.reviewssService.getAllReviewsOfCharger(this.charger.id).subscribe(
        (response) => {
          if (response.body) {
            this.recensions = response.body as RecensionWithUserDTO[]
            this.reviewsVisible = true
          }
          else {
            this.snackBarService.openSnackBar(`Charger '${this.charger?.name}' has no reviews yet.`)
          }
        },
        (err) => {
          console.log(err.error)
        }
      )
  }

  openReviews() {
    if (!this.reviewsVisible) {
      this.loadRecensions();
    } else {
      this.reviewsVisible = false
    }
  }
}
