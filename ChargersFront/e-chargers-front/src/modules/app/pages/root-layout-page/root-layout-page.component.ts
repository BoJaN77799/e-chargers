import { Component, OnInit } from '@angular/core';
import { UtilService } from 'src/modules/shared/service/utils-service';

@Component({
  selector: 'app-root-layout-page',
  templateUrl: './root-layout-page.component.html',
  styleUrls: ['./root-layout-page.component.scss']
})
export class RootLayoutPageComponent {

  constructor(public utilService: UtilService) { }

  isRole(role: string): boolean {
    return this.utilService.isRoleInUserRoles(role);
  }

}
