import { Component, Input, OnInit } from '@angular/core';
import * as moment from 'moment';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { RecensionWithUserDTO } from '../../model/reviewDTO';
import { SnackBarService } from '../../service/snack-bar.service';
import { UtilService } from '../../service/utils-service';

@Component({
  selector: 'app-recension',
  templateUrl: './recension.component.html',
  styleUrls: ['./recension.component.scss']
})
export class RecensionComponent implements OnInit {

  @Input()
  recension: RecensionWithUserDTO;

  constructor(private authSerice: AuthService, private snackBarService: SnackBarService, private utilService: UtilService) {
    this.recension = {} as RecensionWithUserDTO
  }

  ngOnInit(): void {
  }

  formatDate(date: number) {
    return moment(date).format("YYYY-MM-DD HH:mm")
  }

  getFullName = () => `${this.recension.firstname} ${this.recension.lastname}`


  getToxicityClass(toxic: number): string {
    if (toxic > 0.75) {
      return 'strike-ai-red';
    } else if (toxic > 0.40 && toxic <= 0.75) {
      return 'strike-ai-orange';
    } else {
      return 'strike-ai-green';
    }
  }

  getToxicityPercentage = (toxic: number) => toxic * 100

  strikeUser(userId: string) {
    this.authSerice.strikeUser(userId, this.recension.id!).subscribe(
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

  isRole(role: string) {
    return this.utilService.isRole(role);
  }

}
