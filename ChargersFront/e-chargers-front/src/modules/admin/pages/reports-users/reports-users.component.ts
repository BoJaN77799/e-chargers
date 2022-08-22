import { Component, OnInit } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import * as moment from 'moment';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { UserReportDTO } from '../../model/reportsDTO';
import { ReportsService } from '../../service/reportsService';

@Component({
  selector: 'app-reports-users',
  templateUrl: './reports-users.component.html',
  styleUrls: ['./reports-users.component.scss']
})
export class ReportsUsersComponent implements OnInit {

  users: UserReportDTO[]

  displayedColumns: string[] = ['username', 'email', 'firstname', 'lastname', 'user_role', 'strikes', 'banned', 'banned_at', 'banned_until'];
  dataSource: MatTableDataSource<UserReportDTO>;

  constructor(private reportsService: ReportsService, private snackBarService: SnackBarService) {
    this.users = [];
    this.dataSource = new MatTableDataSource(this.users);
  }

  ngOnInit(): void {
    this.reportsService.getUserReport().subscribe(
      (response) => {
        if (response) {
          this.users = response.body as UserReportDTO[]
          console.log(this.users)
          this.dataSource = new MatTableDataSource(this.users)
        }
      },
      (err) => {
        console.log(err.error)
        this.snackBarService.openSnackBarFast(err.error)
      }
    )
  }

  formatDate(date: number) {
    if (date === 0) {
      return "/"
    }
    return moment(date).format("YYYY-MM-DD HH:mm")
  }

}
