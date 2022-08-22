import { Injectable } from "@angular/core";
import { Router, ActivatedRouteSnapshot, CanActivate } from "@angular/router";
import { JwtHelperService } from "@auth0/angular-jwt";
import { UtilService } from "src/modules/shared/service/utils-service";
import { AuthService } from "../../services/auth.service";

@Injectable({
  providedIn: "root",
})
export class RoleGuard implements CanActivate {
  constructor(public auth: AuthService, public router: Router, public utilService: UtilService) { }

  canActivate(route: ActivatedRouteSnapshot): boolean {
    const expectedRoles: string = route.data["expectedRoles"];
    const token = sessionStorage.getItem("user");

    if (!token) {
      this.router.navigate(["myapp/user/home"]);
      return false;
    }

    // const info = jwt.decodeToken(token);
    const roles: string[] = expectedRoles.split("|");
    for (let i = 0; i < roles.length; i++) {
      if (this.utilService.isRole(roles[i]))
        return true;
    }

    this.router.navigate(["myapp/unauthorized"]);
    return false;
  }
}
