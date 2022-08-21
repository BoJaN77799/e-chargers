import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { ProfileComponent } from 'src/modules/shared/components/profile/profile.component';
import { VehiclesComponent } from 'src/modules/user/compontents/vehicles/vehicles.component';

@Component({
  selector: 'app-header-user',
  templateUrl: './header-user.component.html',
  styleUrls: ['./header-user.component.scss']
})
export class HeaderUserComponent implements OnInit {


  constructor(private authService: AuthService, public dialog: MatDialog) {
  }

  ngOnInit(): void {
  }

  logout(): void {
    // this.authService.logout().subscribe((result) => {
    //   console.log(result);
    // });

    sessionStorage.removeItem("user");
  }

  openProfileDialog(): void {
    const dialogRef = this.dialog.open(ProfileComponent, {
      width: '30%',
      panelClass: 'custom-dialog-container'
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
    });
  }

  openVehiclesDialog(): void {
    const dialogRef = this.dialog.open(VehiclesComponent, {
      width: '40%',
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
    });
  }
}
