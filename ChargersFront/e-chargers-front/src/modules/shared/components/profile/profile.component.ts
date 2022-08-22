import { Component, OnInit } from '@angular/core';
import { UserProfileDTO, UserReportDTO } from 'src/modules/admin/model/reportsDTO';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { SnackBarService } from '../../service/snack-bar.service';
import { UtilService } from '../../service/utils-service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss']
})
export class ProfileComponent implements OnInit {

  user!: UserProfileDTO;

  constructor(private authService: AuthService, private utilService: UtilService, private snackBarService: SnackBarService) { }

  ngOnInit(): void {
    this.authService.getUserInfo(this.utilService.getLoggedUsername()).subscribe(
      (response) => {
        if (response) {
          this.user = response.body as UserProfileDTO
        }
      },
      (err) => {
        console.log(err)
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

  checkStrikes() {
    return this.user.strikes > 2
  }

}