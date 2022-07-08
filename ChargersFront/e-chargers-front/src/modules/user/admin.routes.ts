import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { UserHomeComponent } from "./peges/user-home/user-home.component";

export const UserRoutes: Routes = [
  {
    path: "home",
    pathMatch: "full",
    component: UserHomeComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "RegisteredUser" },
  },
];