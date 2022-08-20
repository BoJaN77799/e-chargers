import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { UserReservationsComponent } from "../reservation/pages/user-reservations/user-reservations.component";
import { UserHomeComponent } from "./pages/user-home/user-home.component";

export const UserRoutes: Routes = [
  {
    path: "homepage",
    pathMatch: "full",
    component: UserHomeComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "RegisteredUser" },
  },
  {
    path: "reservations",
    pathMatch: "full",
    component: UserReservationsComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "RegisteredUser" },
  },
];