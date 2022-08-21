import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../services/auth.service';
import { SnackBarService } from 'src/modules/shared/service/snack-bar.service';
import { UtilService } from 'src/modules/shared/service/utils-service';
import { UsernameValidator } from 'src/modules/shared/validators/UsernameValidator';
import { MinLengthValidator } from 'src/modules/shared/validators/MinLengthValidator';
import { MaxLengthValidator } from 'src/modules/shared/validators/MaxLengthValidator';
import { MinLengthPasswordValidator } from 'src/modules/shared/validators/MinLengthPasswordValidator';
import { PasswordValidator } from 'src/modules/shared/validators/PasswordValidator';
import { Login } from '../../models/login';
import * as moment from 'moment';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  form: FormGroup;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private snackBarService: SnackBarService,
    private utilService: UtilService
  ) {
    this.form = this.fb.group({
      username: [null, [Validators.required, UsernameValidator, MinLengthValidator, MaxLengthValidator]],
      password: [null, [Validators.required, PasswordValidator, MinLengthPasswordValidator, MaxLengthValidator]],
    });
  }

  ngOnInit(): void {
  }

  submit() {
    const auth: Login = {
      username: this.form.value.username,
      password: this.form.value.password,
    };

    this.authService.login(auth).subscribe((result: any) => {
      this.snackBarService.openSnackBar("Successful login!");

      const token = JSON.stringify(result);
      sessionStorage.setItem("user", token);

      if (this.utilService.isRole("Administrator")) {
        this.router.navigate(["myapp/admin/homepage"]);
      }
      if (this.utilService.isRole("RegisteredUser")) {
        this.router.navigate(["myapp/user/homepage"])
      }
    },
      (err: any) => {
        if ((err.error as string).includes("banned")) {
          let banned_until_index = (err.error as string).lastIndexOf('l')
          let banned_until = (err.error as string).substring(banned_until_index + 2)
          console.log(banned_until)
          this.snackBarService.openSnackBar("you are banned until " + moment(Number(banned_until)).format("YYYY-MM-DD HH:mm"))
        } else
          this.snackBarService.openSnackBar(err.error);
      }
    );
  }

  register() {
    this.router.navigate(["myapp/auth/registration"]);
  }
}
