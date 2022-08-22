import { Injectable } from "@angular/core";
import { Router, CanActivate } from "@angular/router";
import { UtilService } from "src/modules/shared/service/utils-service";
import { AuthService } from "../../services/auth.service";

@Injectable({
  providedIn: "root",
})
export class LoginGuard implements CanActivate {
  constructor(public router: Router, public utilService: UtilService) { }

  canActivate(): boolean {
    if (this.utilService.isLoggedIn()) {
      if (this.utilService.isRole("Administrator")) {
        this.router.navigate(["myapp/admin"]);
      }
      return false;
    }
    return true;
  }
}
