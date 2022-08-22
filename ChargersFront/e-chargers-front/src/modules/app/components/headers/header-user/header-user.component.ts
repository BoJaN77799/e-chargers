import { Component, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { AuthService } from 'src/modules/auth/services/auth.service';
import { ProfileComponent } from 'src/modules/shared/components/profile/profile.component';
import { VehiclesComponent } from 'src/modules/user/compontents/vehicles/vehicles.component';

@Component({
  selector: 'app-header-user',
  templateUrl: './header-user.component.html',
  styleUrls: ['./header-user.component.scss']
})
export class HeaderUserComponent implements OnInit {


  constructor(private router: Router, private authService: AuthService, public dialog: MatDialog) {
  }

  ngOnInit(): void {
  }

  logout(): void {
    sessionStorage.removeItem("user");
    this.router.navigate(['myapp/user/homepage'])
  }

  openProfileDialog(): void {
    const dialogRef = this.dialog.open(ProfileComponent, {
      width: '30%',
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
    });
  }

  openVehiclesDialog(): void {
    const dialogRef = this.dialog.open(VehiclesComponent, {
      width: '40%',
      height: '80%',
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
    });
  }
}
