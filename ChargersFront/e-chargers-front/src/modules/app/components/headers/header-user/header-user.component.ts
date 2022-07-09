import { Component, OnInit } from '@angular/core';
import { AuthService } from 'src/modules/auth/services/auth.service';

@Component({
  selector: 'app-header-user',
  templateUrl: './header-user.component.html',
  styleUrls: ['./header-user.component.scss']
})
export class HeaderUserComponent implements OnInit {


  constructor(private authService: AuthService) {
  }

  ngOnInit(): void {
  }

  logout(): void {
    this.authService.logout().subscribe((result) => {
      console.log(result);
    });

    sessionStorage.removeItem("user");
  }
}
