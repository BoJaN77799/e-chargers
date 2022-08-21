import { Component, Input, OnInit } from '@angular/core';
import * as moment from 'moment';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { ReviewDTO } from '../../model/reviewDTO';
import { SnackBarService } from '../../service/snack-bar.service';

@Component({
  selector: 'app-recension',
  templateUrl: './recension.component.html',
  styleUrls: ['./recension.component.scss']
})
export class RecensionComponent implements OnInit {

  @Input()
  recension: ReviewDTO | undefined;

  constructor(private authSerice: AuthService, private snackBarService: SnackBarService) { }

  ngOnInit(): void {
  }

  formatDate(date: number | undefined) {
    return moment(date).format("YYYY-MM-DD HH:mm")
  }

  strikeUser(username: string | undefined) {
    if (username) {
      this.authSerice.strikeUser(username).subscribe(
        (response) => {
          if (response) {
            this.snackBarService.openSnackBar(response.body as string);
          }
        },
        (err) => {
          this.snackBarService.openSnackBarFast(err.error)
        }
      )
    }
  }

}
