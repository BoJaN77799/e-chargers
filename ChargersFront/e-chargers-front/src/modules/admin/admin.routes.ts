import { Routes } from "@angular/router";
import { RoleGuard } from "../auth/guards/role/role.guard";
import { AdminHomeComponent } from "./pages/admin-home/admin-home.component";
import { CreateChargerComponent } from "./pages/create-charger/create-charger.component";

export const AdminRoutes: Routes = [
  {
    path: "homepage",
    pathMatch: "full",
    component: AdminHomeComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "Administrator" },
  },
  {
    path: "charger-creation",
    pathMatch: "full",
    component: CreateChargerComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "Administrator" },
  },
];