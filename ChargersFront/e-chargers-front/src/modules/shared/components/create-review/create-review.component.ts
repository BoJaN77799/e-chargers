import { Component, Inject, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialog, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import * as moment from 'moment';
import { ReservationDTO } from 'src/modules/reservation/model/reservationDTO';
import { UserReservationsComponent } from 'src/modules/reservation/pages/user-reservations/user-reservations.component';
import { ReviewDTO } from '../../model/reviewDTO';
import { ReviewService } from '../../service/reviewService';
import { SnackBarService } from '../../service/snack-bar.service';
import { UtilService } from '../../service/utils-service';

@Component({
  selector: 'app-create-review',
  templateUrl: './create-review.component.html',
  styleUrls: ['./create-review.component.scss']
})
export class CreateReviewComponent implements OnInit {

  reviewForm: FormGroup

  rating: number = 0

  constructor(public dialog: MatDialog,
    public dialogRef: MatDialogRef<UserReservationsComponent>,
    @Inject(MAT_DIALOG_DATA) public reservation: ReservationDTO,
    private fb: FormBuilder,
    private utilService: UtilService,
    private reviewService: ReviewService,
    private snackBarService: SnackBarService) {
    this.reviewForm = this.fb.group({
      content: [null, [Validators.required]]
    });
  }

  ngOnInit(): void {
  }

  submit() {

    if (this.rating === 0) {
      this.snackBarService.openSnackBar("Please rate charger")
      return
    }

    let review: ReviewDTO = {
      "username": this.utilService.getLoggedUsername(),
      "charger_id": this.reservation.charger_id,
      "content": this.reviewForm.get("content")?.value,
      "date": moment().unix(),
      "rate": this.rating,
      "toxic": 0,
    }

    this.reviewService.create(review).subscribe(
      (response) => {
        if (response) {
          this.snackBarService.openSnackBar(response.body as string);
          this.dialogRef.close(this.reservation);
        }
      },
      (err) => {
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

  changeRatging(rating: number) {
    this.rating = rating
  }

}
