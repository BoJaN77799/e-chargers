import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HeaderAdminComponent } from './components/headers/header-admin/header-admin.component';
import { HeaderCommonComponent } from './components/headers/header-common/header-common.component';
import { HeaderUserComponent } from './components/headers/header-user/header-user.component';
import { NotFoundPageComponent } from './pages/not-found-page/not-found-page.component';
import { RootLayoutPageComponent } from './pages/root-layout-page/root-layout-page.component';
import { UnauthorizedPageComponent } from './pages/unauthorized-page/unauthorized-page.component';
import { HttpClientModule } from '@angular/common/http';
import { AuthModule } from '../auth/auth.module';
import { SharedModule } from '../shared/shared.module';

@NgModule({
  declarations: [
    AppComponent,
    HeaderCommonComponent,
    HeaderAdminComponent,
    HeaderUserComponent,
    NotFoundPageComponent,
    RootLayoutPageComponent,
    UnauthorizedPageComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    AuthModule,
    SharedModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
