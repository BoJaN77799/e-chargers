import { Component, Input, OnInit } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { CreateReservationComponent } from 'src/modules/reservation/components/create-reservation/create-reservation.component';
import { ChargerDTO } from 'src/modules/shared/model/chargerDTO';
import { UtilService } from '../../service/utils-service';

@Component({
  selector: 'app-charger-info',
  templateUrl: './charger-info.component.html',
  styleUrls: ['./charger-info.component.scss']
})
export class ChargerInfoComponent implements OnInit {

  @Input()
  charger: ChargerDTO | undefined

  constructor(public dialog: MatDialog, private utilService: UtilService) { }

  ngOnInit(): void {
  }

  openReservationDialog(): void {
    const dialogRef = this.dialog.open(CreateReservationComponent, {
      data: this.charger,
      width: '60%',
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log(result);
    });
  }

  isRole(role: string): boolean {
    return this.utilService.isRole(role);
  }

}
