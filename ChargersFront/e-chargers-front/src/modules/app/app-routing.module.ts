import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { NotFoundPageComponent } from './pages/not-found-page/not-found-page.component';
import { RootLayoutPageComponent } from './pages/root-layout-page/root-layout-page.component';
import { UnauthorizedPageComponent } from './pages/unauthorized-page/unauthorized-page.component';

const routes: Routes = [
  {
    path: "myapp",
    component: RootLayoutPageComponent,
    children: [
      {
        path: "auth",
        loadChildren: () =>
          import("./../auth/auth.module").then((m) => m.AuthModule),
      }
    ]
  },
  {
    path: "myapp/unauthorized",
    component: UnauthorizedPageComponent,
    pathMatch: "full"
  },
  {
    path: "",
    redirectTo: "myapp/auth/login",
    pathMatch: "full",
  },
  {
    path: "",
    redirectTo: "myapp/auth/registration",
    pathMatch: "full",
  },
  {
    path: "**",
    component: NotFoundPageComponent
  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
