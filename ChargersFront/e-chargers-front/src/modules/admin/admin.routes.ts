import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { AdminHomeComponent } from "./pages/admin-home/admin-home.component";

export const AdminRoutes: Routes = [
  {
    path: "homepage",
    pathMatch: "full",
    component: AdminHomeComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "Administrator" },
  },
];