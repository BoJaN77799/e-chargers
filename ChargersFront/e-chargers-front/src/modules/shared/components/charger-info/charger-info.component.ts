import { AfterViewInit, Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { CreateReservationComponent } from 'src/modules/reservation/components/create-reservation/create-reservation.component';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { ReviewDTO } from '../../model/reviewDTO';
import { ReviewService } from '../../service/reviewService';
import { UtilService } from '../../service/utils-service';

@Component({
  selector: 'app-charger-info',
  templateUrl: './charger-info.component.html',
  styleUrls: ['./charger-info.component.scss']
})
export class ChargerInfoComponent {

  @Input()
  charger: ChargerDTO | undefined

  reviewsVisible = false
  recensions: ReviewDTO[];

  constructor(
    public dialog: MatDialog,
    private utilService: UtilService,
    private reviewssService: ReviewService) {
    this.recensions = [];
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
    if (this.recensions.length === 0) {
      console.log(this.charger)
      if (this.charger)
        this.reviewssService.getAllReviewsOfCharger(this.charger.id).subscribe(
          (response) => {
            this.recensions = response.body as ReviewDTO[]
            console.log(this.recensions)
          },
          (err) => {
            console.log(err.error)
          }
        )
    }
  }

  toggleReviews() {
    this.loadRecensions();
    this.reviewsVisible = !this.reviewsVisible
  }

}
